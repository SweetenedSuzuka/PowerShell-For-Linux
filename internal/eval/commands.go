package eval

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"powershell/internal/ast"
	"powershell/internal/builtin"
	"powershell/internal/external"
	"powershell/internal/object"
	"powershell/internal/parser"
	"powershell/internal/shell"
)

// RunSource 解析并执行一段源码，返回输出对象。
func (e *Evaluator) RunSource(src string) ([]*object.PSObject, error) {
	res := parser.Parse(src)
	if res.Error != nil {
		return nil, res.Error
	}
	var out []*object.PSObject
	for _, st := range res.List.Statements {
		out = append(out, e.EvalStatement(st)...)
	}
	return out, nil
}

// EvalStatements 是顶层求值入口：执行语句列表，返回输出对象。
func (e *Evaluator) EvalStatements(list *ast.StatementList) []*object.PSObject {
	out, sig := e.runStatements(list.Statements)
	if sig != nil && sig.kind == flowExit {
		e.ExitRequested = true
		e.ExitCode = sig.code
	}
	return out
}

// EvalStatement 执行单条语句并返回输出对象（顶层逐语句输出用，保证与直写命令顺序一致）。
func (e *Evaluator) EvalStatement(st ast.Node) []*object.PSObject {
	out, sig := e.runStatements([]ast.Node{st})
	if sig != nil && sig.kind == flowExit {
		e.ExitRequested = true
		e.ExitCode = sig.code
	}
	return out
}

func (e *Evaluator) evalPipeline(pipe *ast.Pipeline) []*object.PSObject {
	var cur []*object.PSObject
	if pipe.Expr != nil {
		if inc, ok := pipe.Expr.(*ast.Increment); ok {
			// $i++ 作为语句：仅副作用，不输出
			e.evalValue(inc)
		} else {
			cur = flattenOutput(e.evalValue(pipe.Expr))
		}
		e.Session.LastSuccess = true // 纯表达式语句求值成功
	}
	for i, cmd := range pipe.Commands {
		isLast := i == len(pipe.Commands)-1
		cur = flattenPipelineList(e.execCommand(cmd, cur, isLast))
	}
	return cur
}

// flattenPipelineList 在管道节点间把数组摊平（PowerShell 的枚举语义）。
func flattenPipelineList(in []*object.PSObject) []*object.PSObject {
	var out []*object.PSObject
	for _, o := range in {
		if o.IsArray() && o.TypeName == "Object[]" {
			for _, it := range o.ArrayItems() {
				out = append(out, flattenPipelineList([]*object.PSObject{it})...)
			}
		} else {
			out = append(out, o)
		}
	}
	return out
}

func flattenOutput(o *object.PSObject) []*object.PSObject {
	if o == nil || o.IsNull() {
		return nil
	}
	if o.IsArray() {
		return o.ArrayItems()
	}
	return []*object.PSObject{o}
}

// execCommand 调度一条命令：别名 → 函数 → 内置 → 脚本 → 外部。
func (e *Evaluator) execCommand(cmd *ast.Command, input []*object.PSObject, isLast bool) []*object.PSObject {
	name := cmd.Name
	// 递归解析别名链（New-Alias foo ls; foo 也生效）；带深度上限防环
	for i := 0; i < 8; i++ {
		resolved, ok := e.Session.ResolveAlias(name)
		if !ok || resolved == name {
			break
		}
		name = resolved
	}
	if fn, ok := e.findFunction(name); ok {
		return e.applyRedirects(cmd, e.callFunction(fn, cmd, input))
	}
	if fn, ok := builtin.Lookup(name); ok {
		spec := builtin.Spec(name)
		args, err := builtin.Bind(e, cmd, spec, nil)
		if err != nil {
			e.writeError(err)
			return nil
		}
		// 参数绑定后才定 $?：绑定过程可能读取 $?，不能被新命令提前覆盖
		e.Session.LastSuccess = true
		// 2> 重定向：执行期间把 stderr 指到目标（$null 则丢弃）
		if w, closer := e.stderrRedirectTarget(cmd); w != nil {
			oldErr := e.hostErr
			e.hostErr = w
			ctx := &builtin.Context{
				Shell: e.Session, Engine: e,
				Stdout: e.hostOut, Stderr: e.hostErr, Stdin: e.stdin,
				Args: args, Input: input,
			}
			out, err := fn(ctx)
			e.hostErr = oldErr
			if closer != nil {
				closer.Close()
			}
			if err != nil {
				e.writeError(err)
				return nil
			}
			return e.applyRedirects(cmd, out)
		}
		ctx := &builtin.Context{
			Shell: e.Session, Engine: e,
			Stdout: e.hostOut, Stderr: e.hostErr, Stdin: e.stdin,
			Args: args, Input: input,
		}
		out, err := fn(ctx)
		if err != nil {
			e.writeError(err)
			return nil
		}
		return e.applyRedirects(cmd, out)
	}
	if isScriptPath(name) {
		return e.applyRedirects(cmd, e.runScriptFile(name, input))
	}
	return e.runExternal(cmd, input, isLast)
}

func (e *Evaluator) findFunction(name string) (*shell.Function, bool) {
	if fn, ok := e.Session.Functions[name]; ok {
		return fn, true
	}
	for k, v := range e.Session.Functions {
		if strings.EqualFold(k, name) {
			return v, true
		}
	}
	return nil, false
}

func isScriptPath(name string) bool {
	return strings.HasSuffix(strings.ToLower(name), ".ps1")
}

// redirTargetPath 求值重定向目标并解析为绝对路径（~ 展开、相对基于会话当前目录）。
// $null/空或非法盘符返回 false（非法盘符会报错）。
func (e *Evaluator) redirTargetPath(node ast.Node) (string, bool) {
	target := e.evalValue(node)
	if target.IsNull() || target.String() == "" {
		return "", false
	}
	np, err := shell.ResolvePath(e.Session.Cwd, target.String())
	if err != nil {
		e.writeError(err)
		return "", false
	}
	return np, true
}

// stderrRedirectTarget 打开命令的 2> 重定向目标；目标为 $null 时返回 io.Discard。
func (e *Evaluator) stderrRedirectTarget(cmd *ast.Command) (io.Writer, io.Closer) {
	for _, r := range cmd.Redirs {
		if r.Kind == ast.RedirStderr {
			target, ok := e.redirTargetPath(r.Target)
			if !ok {
				return io.Discard, nil
			}
			flags := os.O_WRONLY | os.O_CREATE
			if r.Append {
				flags |= os.O_APPEND
			} else {
				flags |= os.O_TRUNC
			}
			if f, err := os.OpenFile(target, flags, 0o644); err == nil {
				return f, f
			}
			return io.Discard, nil
		}
	}
	return nil, nil
}

// applyRedirects 处理命令的 stdout 重定向（> / >>）。
func (e *Evaluator) applyRedirects(cmd *ast.Command, out []*object.PSObject) []*object.PSObject {
	if len(cmd.Redirs) == 0 {
		return out
	}
	for _, r := range cmd.Redirs {
		if r.Kind != ast.RedirStdout && r.Kind != ast.RedirAppend {
			continue
		}
		target, ok := e.redirTargetPath(r.Target)
		if !ok {
			continue // > $null 或非法盘符：丢弃
		}
		var buf bytes.Buffer
		_ = object.FormatOutput(&buf, out)
		flags := os.O_WRONLY | os.O_CREATE
		if r.Append {
			flags |= os.O_APPEND
		} else {
			flags |= os.O_TRUNC
		}
		f, err := os.OpenFile(target, flags, 0o644)
		if err != nil {
			e.writeError(fmt.Errorf("无法写入重定向目标 %s：%v", target, err))
			continue
		}
		_, _ = f.WriteString(buf.String())
		_ = f.Close()
	}
	return nil
}

// ---- 函数调用 ----

func (e *Evaluator) callFunction(fn *shell.Function, cmd *ast.Command, input []*object.PSObject) []*object.PSObject {
	e.pushScope()
	defer e.popScope()
	e.inCapture++
	defer func() { e.inCapture-- }()

	// 求值实参
	var posVals []*object.PSObject
	namedVals := map[string]*object.PSObject{}
	namedSwitch := map[string]bool{}
	for _, slot := range cmd.ArgOrder {
		switch slot.Kind {
		case ast.ArgPositional:
			posVals = append(posVals, e.evalValue(cmd.Positional[slot.Index]))
		case ast.ArgNamed:
			namedVals[slot.Name] = e.evalValue(cmd.Named[slot.Index].Value)
		case ast.ArgSwitch:
			namedSwitch[slot.Name] = true
		}
	}
	sc := e.scopes[len(e.scopes)-1]
	bound := map[string]bool{}
	for _, p := range fn.Params {
		if namedSwitch[p.Name] {
			sc[p.Name] = object.Bool(true)
			bound[p.Name] = true
		} else if v, ok := namedVals[p.Name]; ok {
			sc[p.Name] = v
			bound[p.Name] = true
		}
	}
	pi := 0
	for _, p := range fn.Params {
		if bound[p.Name] {
			continue
		}
		if pi < len(posVals) {
			sc[p.Name] = posVals[pi]
			pi++
		} else if p.Default != nil {
			sc[p.Name] = e.evalValue(p.Default)
		} else {
			sc[p.Name] = object.Null()
		}
	}
	var extra []*object.PSObject
	for ; pi < len(posVals); pi++ {
		extra = append(extra, posVals[pi])
	}
	sc["args"] = object.Array(extra)
	if len(input) > 0 {
		sc["input"] = object.Array(input)
	}

	out, sig := e.runStatements(fn.Body.Body.Statements)
	if sig != nil {
		if sig.kind == flowReturn {
			return unwrapOutput(sig.value)
		}
		if sig.kind == flowExit {
			panic(sig)
		}
	}
	return out
}

// ---- 外部命令 ----

// externalArgv 依源码顺序重建外部命令的 argv。
func (e *Evaluator) externalArgv(cmd *ast.Command) (string, []string) {
	program := cmd.Name
	var argv []string
	for _, slot := range cmd.ArgOrder {
		switch slot.Kind {
		case ast.ArgPositional:
			v := e.evalValue(cmd.Positional[slot.Index])
			argv = append(argv, v.String())
		case ast.ArgNamed:
			v := e.evalValue(cmd.Named[slot.Index].Value)
			argv = append(argv, "-"+slot.Name, v.String())
		case ast.ArgSwitch:
			argv = append(argv, "-"+slot.Name)
		}
	}
	return program, argv
}

func (e *Evaluator) notFound(name string) {
	fmt.Fprintf(e.hostErr, "%s : 无法将“%s”项识别为 cmdlet、函数、脚本文件或可运行程序的名称。\n", e.Session.StyleName(), name)
	e.Session.LastSuccess = false
}

func (e *Evaluator) runExternal(cmd *ast.Command, input []*object.PSObject, isLast bool) []*object.PSObject {
	program, argv := e.externalArgv(cmd)
	if _, ok := external.LookPath(program); !ok {
		e.notFound(cmd.Name)
		return nil
	}
	// stdin：管道输入渲染为文本；无输入则用会话 stdin
	var inSrc io.Reader = e.stdin
	if len(input) > 0 {
		var sb strings.Builder
		for _, o := range input {
			sb.WriteString(o.String())
			sb.WriteByte('\n')
		}
		inSrc = strings.NewReader(sb.String())
	}
	// stdout/stderr 目标：重定向文件或宿主
	stdout := io.Writer(e.hostOut)
	stderr := io.Writer(e.hostErr)
	var outFile, errFile *os.File
	for _, r := range cmd.Redirs {
		target, ok := e.redirTargetPath(r.Target)
		if !ok {
			// > $null 或非法盘符：丢弃
			if r.Kind == ast.RedirStdout || r.Kind == ast.RedirAppend {
				stdout = io.Discard
			} else {
				stderr = io.Discard
			}
			continue
		}
		flags := os.O_WRONLY | os.O_CREATE
		if r.Append {
			flags |= os.O_APPEND
		} else {
			flags |= os.O_TRUNC
		}
		switch r.Kind {
		case ast.RedirStdout, ast.RedirAppend:
			if f, err := os.OpenFile(target, flags, 0o644); err == nil {
				outFile = f
			}
		case ast.RedirStderr:
			if f, err := os.OpenFile(target, flags, 0o644); err == nil {
				errFile = f
			}
		}
	}
	if outFile != nil {
		stdout = outFile
		defer outFile.Close()
	}
	if errFile != nil {
		stderr = errFile
		defer errFile.Close()
	}

	// 最后一条命令且不在捕获模式 → 直连宿主
	if isLast && e.inCapture == 0 {
		code := external.Exec(e.Session.Cwd, program, argv, inSrc, stdout, stderr)
		e.Session.LastExit = code
		e.Session.LastSuccess = code == 0
		return nil
	}
	// 捕获输出：有显式重定向目标（文件/丢弃）就直写目标，否则按行转字符串对象
	var outBuf, errBuf bytes.Buffer
	outW := io.Writer(&outBuf)
	errW := io.Writer(&errBuf)
	captureOut, captureErr := true, true
	if outFile != nil {
		outW, captureOut = outFile, false
	} else if stdout == io.Discard {
		outW, captureOut = io.Discard, false
	}
	if errFile != nil {
		errW, captureErr = errFile, false
	} else if stderr == io.Discard {
		errW, captureErr = io.Discard, false
	}
	code := external.Exec(e.Session.Cwd, program, argv, inSrc, outW, errW)
	e.Session.LastExit = code
	e.Session.LastSuccess = code == 0
	if captureErr && errBuf.Len() > 0 {
		_, _ = io.Copy(e.hostErr, &errBuf)
	}
	var lines []*object.PSObject
	if captureOut {
		text := strings.TrimSuffix(outBuf.String(), "\n")
		if text != "" {
			for _, ln := range strings.Split(text, "\n") {
				lines = append(lines, object.Str(ln))
			}
		}
	}
	return lines
}

// ---- 脚本文件 ----

// RunScriptFile 读取并执行一个 .ps1 脚本，返回输出对象。
func (e *Evaluator) RunScriptFile(path string, args []*object.PSObject) []*object.PSObject {
	return e.runScript(path, args, nil)
}

// RunScriptFileStreaming 读取并逐语句执行脚本，每语句的输出交给 emit（保证直写命令顺序）。
func (e *Evaluator) RunScriptFileStreaming(path string, args []*object.PSObject, emit func(objs []*object.PSObject)) {
	e.runScript(path, args, emit)
}

func (e *Evaluator) runScript(path string, args []*object.PSObject, emit func(objs []*object.PSObject)) []*object.PSObject {
	data, err := os.ReadFile(path)
	if err != nil {
		e.writeError(fmt.Errorf("无法读取脚本 %s：%v", path, err))
		return nil
	}
	res := parser.Parse(string(data))
	if res.Error != nil {
		e.writeError(fmt.Errorf("脚本 %s 解析错误：%v", path, res.Error))
		return nil
	}
	abs, _ := filepath.Abs(path)
	oldPSCommandPath := e.Session.PSCommandPath
	oldArgs := e.Session.Args
	e.Session.PSCommandPath = abs
	e.Session.Args = args
	defer func() {
		e.Session.PSCommandPath = oldPSCommandPath
		e.Session.Args = oldArgs
	}()
	e.inCapture++
	defer func() { e.inCapture-- }()
	var all []*object.PSObject
	for _, st := range res.List.Statements {
		objs := e.EvalStatement(st)
		all = append(all, objs...)
		if emit != nil {
			emit(objs)
		}
		if e.ExitRequested {
			break
		}
	}
	return all
}

// runScriptFile 作为命令调用脚本（.\foo.ps1）。
func (e *Evaluator) runScriptFile(name string, input []*object.PSObject) []*object.PSObject {
	path := name
	if !filepath.IsAbs(path) {
		path = filepath.Join(e.Session.Cwd, path)
	}
	path = filepath.Clean(path)
	var args []*object.PSObject
	if len(input) > 0 {
		args = input
	}
	return e.RunScriptFile(path, args)
}

// ---- 过滤表达式（Where-Object 裸属性） ----

// EvalFilterExpr 在 $_ = obj 上下文求值过滤表达式（裸字视为对象属性）。
func (e *Evaluator) EvalFilterExpr(node ast.Node, obj *object.PSObject) (bool, error) {
	rewritten := rewriteBarewords(node)
	val, err := e.EvalExpr(rewritten, map[string]*object.PSObject{"_": obj, "PSItem": obj})
	if err != nil {
		return false, err
	}
	return val.Truthy(), nil
}

// rewriteBarewords 把二元/一元表达式中的裸字替换为 $_.<name>。
func rewriteBarewords(n ast.Node) ast.Node {
	switch v := n.(type) {
	case *ast.Binary:
		return &ast.Binary{Op: v.Op, L: rewriteBarewords(v.L), R: rewriteBarewords(v.R)}
	case *ast.Unary:
		return &ast.Unary{Op: v.Op, Operand: rewriteBarewords(v.Operand)}
	case *ast.BareWord:
		return &ast.MemberAccess{Base: &ast.VarRef{Name: "_"}, Prop: v.Value}
	case *ast.ScriptBlock:
		return v
	}
	return n
}

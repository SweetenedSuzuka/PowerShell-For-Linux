package eval

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"powershell/internal/ast"
	"powershell/internal/builtin"
	"powershell/internal/external"
	"powershell/internal/lang"
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
	// 一次性入口拒绝执行不完整语句
	if res.Incomplete {
		return nil, fmt.Errorf("%s", lang.T(lang.MsgIncompleteInput))
	}
	var out []*object.PSObject
	for _, st := range res.List.Statements {
		out = append(out, e.EvalStatement(st)...)
	}
	return out, nil
}

// EvalStatements 是顶层求值入口：执行语句列表，返回输出对象。
// 未捕获的终止错误（throw）在此打印并标记 $?，不中断会话。
func (e *Evaluator) EvalStatements(list *ast.StatementList) []*object.PSObject {
	out, sig := e.runStatements(list.Statements)
	if sig != nil {
		switch sig.kind {
		case flowExit:
			e.ExitRequested = true
			e.ExitCode = sig.code
		case flowError:
			e.printError(fmt.Errorf("%s", sig.value.String()))
		case flowBreak, flowContinue:
			// 无所属循环的 break/continue 终止当前语句序列（与 PowerShell 一致），已产生的输出保留
		}
	}
	return out
}

// EvalStatement 执行单条语句并返回输出对象（顶层逐语句输出用，保证与直写命令顺序一致）。
func (e *Evaluator) EvalStatement(st ast.Node) []*object.PSObject {
	out, sig := e.runStatements([]ast.Node{st})
	if sig != nil {
		switch sig.kind {
		case flowExit:
			e.ExitRequested = true
			e.ExitCode = sig.code
		case flowError:
			e.printError(fmt.Errorf("%s", sig.value.String()))
		case flowBreak, flowContinue:
			// 无所属循环的 break/continue 只终止本条语句（与 PowerShell 一致）
		}
	}
	return out
}

func (e *Evaluator) evalPipeline(pipe *ast.Pipeline) []*object.PSObject {
	var cur []*object.PSObject
	if pipe.Expr != nil {
		// 纯表达式语句求值前置位 $?（对齐命令路径：求值中出错由 writeError 覆盖为 false）
		e.Session.LastSuccess = true
		if inc, ok := pipe.Expr.(*ast.Increment); ok {
			// $i++ 作为语句：仅副作用，不输出
			e.evalValue(inc)
		} else {
			cur = flattenOutput(e.evalValue(pipe.Expr))
		}
	}
	for i, cmd := range pipe.Commands {
		isLast := i == len(pipe.Commands)-1
		hasInput := pipe.Expr != nil || i > 0
		if hasInput {
			// defer 配对递减：命令以 panic 上抛控制流信号时计数随栈展开归位
			e.inPipeline++
			defer func() { e.inPipeline-- }()
		}
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

// builtinError 处理内置 cmdlet 返回的错误。
// errf 返回的终止错误直接上抛，其余按显式动作与首选项分发。
func (e *Evaluator) builtinError(args *builtin.BoundArgs, err error) {
	var term *builtin.TerminatingError
	if errors.As(err, &term) && term.Record != nil {
		panic(&flowSignal{kind: flowError, value: term.Record})
	}
	e.dispatchError(args.ErrorAction, err)
}

// execCommand 调度一条命令：别名 → 函数 → 内置 → 脚本 → 外部。
// Name 为 "&" 的是调用命令：目标求值为脚本块时执行脚本块，否则按名字走常规分发。
func (e *Evaluator) execCommand(cmd *ast.Command, input []*object.PSObject, isLast bool) []*object.PSObject {
	if cmd.Name == "&" {
		return e.execInvoke(cmd, input, isLast)
	}
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
			e.reportError(err)
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
				e.builtinError(args, err)
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
			e.builtinError(args, err)
			return nil
		}
		return e.applyRedirects(cmd, out)
	}
	if isScriptPath(name) {
		// 显式位置实参（如 .\s.ps1 1 2 3）优先作为脚本实参；
		// 无显式实参时沿用管道输入（保持原有近似行为）
		var args []*object.PSObject
		for _, slot := range cmd.ArgOrder {
			if slot.Kind == ast.ArgPositional {
				args = append(args, e.evalValue(cmd.Positional[slot.Index]))
			}
		}
		if len(args) == 0 && len(input) > 0 {
			args = input
		}
		return e.applyRedirects(cmd, e.runScriptFile(name, args))
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
		e.reportError(err)
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
// stdout 被重定向时输出不进管道（返回 nil）；只有 stderr 重定向时输出照常返回。
func (e *Evaluator) applyRedirects(cmd *ast.Command, out []*object.PSObject) []*object.PSObject {
	hasStdout := false
	for _, r := range cmd.Redirs {
		if r.Kind != ast.RedirStdout && r.Kind != ast.RedirAppend {
			continue
		}
		hasStdout = true
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
			e.reportError(fmt.Errorf("%s", lang.T(lang.MsgRedirectWrite, target, err)))
			continue
		}
		_, _ = f.WriteString(buf.String())
		_ = f.Close()
	}
	if hasStdout {
		return nil
	}
	return out
}

// ---- 函数调用 ----

func (e *Evaluator) callFunction(fn *shell.Function, cmd *ast.Command, input []*object.PSObject) []*object.PSObject {
	e.pushScope()
	defer e.popScope()
	e.inCapture++
	defer func() { e.inCapture-- }()

	ca := e.evalCallArgs(cmd, cmd.ArgOrder)
	extra, ok := e.bindParams(fn.Params, ca)
	if !ok {
		return nil
	}
	sc := e.scopes[len(e.scopes)-1]
	sc["args"] = object.Array(extra)
	if len(input) > 0 {
		sc["input"] = object.Array(input)
	}

	if fn.Begin != nil || fn.Process != nil || fn.End != nil || fn.Filter {
		return e.callFunctionNamedBlocks(fn, input)
	}
	out, sig := e.runStatements(fn.Body.Body.Statements)
	if sig != nil {
		switch sig.kind {
		case flowReturn:
			// return 前的输出（如 try/finally 沿途写入的）一并保留
			return append(out, unwrapOutput(sig.value)...)
		case flowBreak, flowContinue:
			// 函数体内没有所属循环，break/continue 沿调用栈上抛（与 PowerShell 一致）
			sig.out = out
			panic(sig)
		case flowExit, flowError:
			panic(sig)
		}
	}
	return out
}

// callFunctionNamedBlocks 按 begin/process/end 语义执行带命名块（或 filter）的函数：
// begin 先跑一次；有管道输入时 process 对每项执行并把 $_ 绑到该项，无输入时以 $null 跑一次；
// filter 无 Process 块时 Body 即 process。end 最后跑一次。
// process 里的 return 只结束当前项的处理，继续下一项；begin/end 里的 return 结束整个函数。
// break/continue 在命名块里没有所属循环，沿调用栈上抛（与 PowerShell 一致），由外层循环捕获。
func (e *Evaluator) callFunctionNamedBlocks(fn *shell.Function, input []*object.PSObject) []*object.PSObject {
	var out []*object.PSObject
	// runBlock 执行一个块，stop 为 true 表示不再继续后续阶段
	runBlock := func(body *ast.Block, inProcess bool) (stop bool) {
		o, sig := e.runStatements(body.Body.Statements)
		out = append(out, o...)
		if sig == nil {
			return false
		}
		switch sig.kind {
		case flowReturn:
			// return 在 process 里只结束本次；在 begin/end 里结束函数
			return !inProcess
		case flowBreak, flowContinue:
			sig.out = out
			panic(sig)
		case flowExit, flowError:
			panic(sig)
		}
		return false
	}
	if fn.Begin != nil {
		if runBlock(fn.Begin, false) {
			return out
		}
	}
	process := fn.Process
	if process == nil && fn.Filter {
		process = fn.Body
	}
	if process != nil {
		items := input
		if e.inPipeline == 0 {
			// 直调（无管道）：process 以 $null 跑一次
			e.scopes[len(e.scopes)-1]["_"] = object.Null()
			runBlock(process, true)
		} else {
			for _, item := range items {
				e.scopes[len(e.scopes)-1]["_"] = item
				if runBlock(process, true) {
					break
				}
			}
		}
	}
	if fn.End != nil {
		runBlock(fn.End, false)
	}
	return out
}

// callArgs 是求值后的调用实参：位置值、命名值与开关名。
type callArgs struct {
	posVals     []*object.PSObject
	namedVals   map[string]*object.PSObject
	namedSwitch map[string]bool
}

// evalCallArgs 按源码顺序求值一条命令的实参。
func (e *Evaluator) evalCallArgs(cmd *ast.Command, slots []ast.ArgItem) callArgs {
	ca := callArgs{namedVals: map[string]*object.PSObject{}, namedSwitch: map[string]bool{}}
	for _, slot := range slots {
		switch slot.Kind {
		case ast.ArgPositional:
			ca.posVals = append(ca.posVals, e.evalValue(cmd.Positional[slot.Index]))
		case ast.ArgNamed:
			ca.namedVals[slot.Name] = e.evalValue(cmd.Named[slot.Index].Value)
		case ast.ArgSwitch:
			ca.namedSwitch[slot.Name] = true
		}
	}
	return ca
}

// bindParams 按形参声明把调用实参落位到当前作用域：命名与开关优先，其余使用位置实参，缺的用默认值或 $null。
// 返回剩余位置实参与绑定是否成功；失败时错误已写出，调用方不再执行被调方。
func (e *Evaluator) bindParams(params []ast.FunctionParam, ca callArgs) ([]*object.PSObject, bool) {
	sc := e.scopes[len(e.scopes)-1]
	bound := map[string]bool{}
	bind := func(p ast.FunctionParam, v *object.PSObject) bool {
		cv, ok := e.bindParamValue(p, v)
		if !ok {
			return false
		}
		sc[p.Name] = cv
		bound[p.Name] = true
		return true
	}
	for _, p := range params {
		var val *object.PSObject
		have := false
		if ca.namedSwitch[p.Name] {
			val, have = object.Bool(true), true
		} else if v, ok := ca.namedVals[p.Name]; ok {
			val, have = v, true
		}
		if have && !bind(p, val) {
			return nil, false
		}
	}
	pi := 0
	for _, p := range params {
		if bound[p.Name] {
			continue
		}
		if pi < len(ca.posVals) {
			if !bind(p, ca.posVals[pi]) {
				return nil, false
			}
			pi++
		} else if p.Default != nil {
			if !bind(p, e.evalValue(p.Default)) {
				return nil, false
			}
		} else {
			sc[p.Name] = object.Null()
		}
	}
	return ca.posVals[pi:], true
}

// execInvoke 执行 & 调用命令：首个位置实参是调用目标。
// 目标为脚本块时按函数语义执行（param 形参、$args、$input、动态作用域）；其余目标转成名字后走常规命令分发。
func (e *Evaluator) execInvoke(cmd *ast.Command, input []*object.PSObject, isLast bool) []*object.PSObject {
	targetIdx := -1
	var rest []ast.ArgItem
	for _, slot := range cmd.ArgOrder {
		if slot.Kind == ast.ArgPositional && targetIdx < 0 {
			targetIdx = slot.Index
			continue
		}
		rest = append(rest, slot)
	}
	if targetIdx < 0 {
		e.reportError(fmt.Errorf("%s", lang.T(lang.MsgInvokeTargetMissing)))
		return nil
	}
	target := e.evalValue(cmd.Positional[targetIdx])
	if node, ok := target.Value.(*ast.ScriptBlock); ok {
		ca := e.evalCallArgs(cmd, rest)
		return e.applyRedirects(cmd, e.invokeScriptBlock(node, ca, input))
	}
	// 非脚本块目标按名字分发：改写成以目标字符串为名的普通命令，剔除已消费的目标实参
	nameCmd := rewriteWithoutPositional(cmd, targetIdx)
	nameCmd.Name = target.String()
	return e.execCommand(nameCmd, input, isLast)
}

// rewriteWithoutPositional 复制命令并剔除指定下标的位置实参，其余实参重建子列表与 ArgOrder 下标。
func rewriteWithoutPositional(cmd *ast.Command, drop int) *ast.Command {
	nc := &ast.Command{Name: cmd.Name, Redirs: cmd.Redirs}
	var pos []ast.Node
	var named []ast.NamedArg
	var sws []string
	for _, slot := range cmd.ArgOrder {
		switch slot.Kind {
		case ast.ArgPositional:
			if slot.Index == drop {
				continue
			}
			nc.ArgOrder = append(nc.ArgOrder, ast.ArgItem{Kind: ast.ArgPositional, Index: len(pos)})
			pos = append(pos, cmd.Positional[slot.Index])
		case ast.ArgNamed:
			nc.ArgOrder = append(nc.ArgOrder, ast.ArgItem{Kind: ast.ArgNamed, Name: slot.Name, Index: len(named), Inline: slot.Inline})
			named = append(named, cmd.Named[slot.Index])
		case ast.ArgSwitch:
			nc.ArgOrder = append(nc.ArgOrder, ast.ArgItem{Kind: ast.ArgSwitch, Name: slot.Name, Index: len(sws)})
			sws = append(sws, cmd.Switches[slot.Index])
		}
	}
	nc.Positional = pos
	nc.Named = named
	nc.Switches = sws
	return nc
}

// invokeScriptBlock 以函数语义执行脚本块：实参绑定、$args、$input 与函数调用同一套规则；
// 带命名块时按 begin/process/end 语义执行（直调 process 以 $null 跑一次），与函数一致。
func (e *Evaluator) invokeScriptBlock(node *ast.ScriptBlock, ca callArgs, input []*object.PSObject) []*object.PSObject {
	e.pushScope()
	defer e.popScope()
	e.inCapture++
	defer func() { e.inCapture-- }()

	var params []ast.FunctionParam
	// 块体开头的 param() 声明从体里提取
	stmts := node.Body.Statements
	if len(stmts) > 0 {
		if pb, ok := stmts[0].(*ast.ParamBlock); ok {
			params = pb.Params
			stmts = stmts[1:]
		}
	}
	sc := e.scopes[len(e.scopes)-1]
	extra, ok := e.bindParams(params, ca)
	if !ok {
		return nil
	}
	sc["args"] = object.Array(extra)
	if len(input) > 0 {
		sc["input"] = object.Array(input)
	}

	if node.Begin != nil || node.Process != nil || node.End != nil {
		fn := &shell.Function{Params: params, Body: &ast.Block{Body: node.Body}, Begin: node.Begin, Process: node.Process, End: node.End}
		return e.callFunctionNamedBlocks(fn, input)
	}
	out, sig := e.runStatements(stmts)
	if sig != nil {
		switch sig.kind {
		case flowReturn:
			return append(out, unwrapOutput(sig.value)...)
		case flowBreak, flowContinue:
			// 脚本块体内没有所属循环，break/continue 沿调用栈上抛（与 PowerShell 一致）
			sig.out = out
			panic(sig)
		case flowExit, flowError:
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
	fmt.Fprintf(e.hostErr, "%s : %s\n", e.Session.StyleName(), lang.T(lang.MsgCommandNotFound, name))
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
		e.Session.LastExit = 1
		e.reportError(fmt.Errorf("%s", lang.T(lang.MsgScriptReadFail, path, err)))
		return nil
	}
	res := parser.Parse(string(data))
	if res.Error != nil {
		// 脚本没有执行任何语句视作失败：置失败退出码，让 -File 与脚本调用方凭退出码感知
		e.Session.LastExit = 1
		e.reportError(fmt.Errorf("%s", lang.T(lang.MsgScriptParseFail, path, res.Error)))
		return nil
	}
	// 截断的脚本在此拒绝执行，前半段语句不会运行
	if res.Incomplete {
		e.Session.LastExit = 1
		e.reportError(fmt.Errorf("%s : %s", path, lang.T(lang.MsgIncompleteInput)))
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
	stmts := res.List.Statements
	// param() 块：脚本开头的参数声明，按实参绑定后从语句里剔除
	if len(stmts) > 0 {
		if pb, ok := stmts[0].(*ast.ParamBlock); ok {
			if !e.bindScriptParams(pb.Params, args) {
				// 实参与形参声明不符视作脚本没有执行任何语句，置失败退出码供调用方感知
				e.Session.LastExit = 1
				return nil
			}
			stmts = stmts[1:]
		}
	}
	var all []*object.PSObject
	for _, st := range stmts {
		objs, sig := e.runStatements([]ast.Node{st})
		all = append(all, objs...)
		if emit != nil {
			emit(objs)
		}
		if sig != nil {
			switch sig.kind {
			case flowExit:
				e.ExitRequested = true
				e.ExitCode = sig.code
			case flowError:
				// 未捕获的终止错误：中止脚本并向上传播，调用方的 try 可以捕获；
				// 传到最外层时由 EvalStatement/main.go 打印。
				// panic 前已产生的输出一并携带。
				sig.out = all
				panic(sig)
			case flowBreak, flowContinue:
				// 无所属循环的 break/continue 终止脚本的当前语句序列（与 PowerShell 一致）
				return all
			}
		}
		if e.ExitRequested {
			break
		}
	}
	return all
}

// bindParamValue 把实参转换成形参 [类型] 标注声明的类型；未标注原样返回。
// 数组类型把单值包成单元素数组后逐元素转换。
// ok 为 false 表示无法转换，错误已写出，调用方不再执行被调方。
func (e *Evaluator) bindParamValue(p ast.FunctionParam, v *object.PSObject) (*object.PSObject, bool) {
	if p.TypeName == "" {
		return v, true
	}
	norm := strings.ToLower(p.TypeName)
	if strings.HasSuffix(norm, "[]") {
		elem := strings.TrimSuffix(norm, "[]")
		items := make([]*object.PSObject, 0, len(v.ArrayItems()))
		for _, it := range v.ArrayItems() {
			out, err := convertTarget(it, elem)
			if err != nil {
				e.writeBindError(err, it, p.Name, elem)
				return nil, false
			}
			items = append(items, out)
		}
		return object.Array(items), true
	}
	out, err := convertTarget(v, norm)
	if err != nil {
		e.writeBindError(err, v, p.Name, norm)
		return nil, false
	}
	return out, true
}

// writeBindError 写参数绑定错误：类型未注册报"无法找到类型"，否则报实参转换失败。
func (e *Evaluator) writeBindError(err error, v *object.PSObject, param, target string) {
	if errors.Is(err, errTypeUnknown) {
		e.reportError(fmt.Errorf("%s", lang.T(lang.MsgTypeUnknown, target)))
		return
	}
	e.reportError(fmt.Errorf("%s", lang.T(lang.MsgBindConvertFail, v.String(), param, target)))
}

// bindScriptParams 按 param() 声明把脚本实参绑到当前作用域（脚本不推独立作用域，变量可见性等价调用点：控制台调用留在会话，函数内调用随函数销毁）。
// 位置实参依次落位，缺的用默认值或 $null，剩余实参保留在 $args。
// 返回 false 表示有实参无法转换成形参声明的类型（错误已写出），调用方不再执行脚本。
func (e *Evaluator) bindScriptParams(params []ast.FunctionParam, args []*object.PSObject) bool {
	sc := e.scopes[len(e.scopes)-1]
	bound := 0
	for _, p := range params {
		if bound < len(args) {
			cv, ok := e.bindParamValue(p, args[bound])
			if !ok {
				return false
			}
			sc[p.Name] = cv
			bound++
		} else if p.Default != nil {
			cv, ok := e.bindParamValue(p, e.evalValue(p.Default))
			if !ok {
				return false
			}
			sc[p.Name] = cv
		} else {
			sc[p.Name] = object.Null()
		}
	}
	e.Session.Args = args[bound:]
	return true
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

// Package eval 实现 PowerShell 求值器：表达式、语句、管道、命令调度、控制流、函数调用与外部命令。
package eval

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"powershell/internal/ast"
	"powershell/internal/builtin"
	"powershell/internal/lang"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// Evaluator 是一次求值会话。
type Evaluator struct {
	Session       *shell.Session
	stdout        io.Writer
	stderr        io.Writer
	stdin         io.Reader
	hostOut       io.Writer
	hostErr       io.Writer
	scopes        []map[string]*object.PSObject // 变量作用域栈，scopes[0] 为全局
	inCapture     int                           // 进入捕获模式（函数/脚本块/子表达式）计数
	inPipeline    int                           // 命令处于管道输入位的层数（>0 表示本次调用有管道输入，哪怕为零项）
	ExitRequested bool                          // 是否遇到 exit 语句
	ExitCode      int                           // exit 码
}

// New 创建求值器。
func New(sess *shell.Session, stdin io.Reader, stdout, stderr io.Writer) *Evaluator {
	return &Evaluator{
		Session: sess,
		stdout:  stdout,
		stderr:  stderr,
		stdin:   stdin,
		hostOut: stdout,
		hostErr: stderr,
		scopes:  []map[string]*object.PSObject{sess.Vars},
	}
}

// ---- Engine 接口实现（供内置 cmdlet 调用） ----

// EvalExpr 在 extra 变量作用域下求值表达式。
func (e *Evaluator) EvalExpr(node ast.Node, extra map[string]*object.PSObject) (*object.PSObject, error) {
	if len(extra) > 0 {
		e.pushScope()
		defer e.popScope()
		for k, v := range extra {
			e.scopes[len(e.scopes)-1][k] = v
		}
	}
	return e.evalValue(node), nil
}

// InvokeBlock 执行脚本块并返回其输出对象（不打印）。
func (e *Evaluator) InvokeBlock(block *ast.Block, extra map[string]*object.PSObject, stdout io.Writer) ([]*object.PSObject, error) {
	if block == nil {
		return nil, nil
	}
	e.inCapture++
	defer func() { e.inCapture-- }()
	if len(extra) > 0 {
		e.pushScope()
		defer e.popScope()
		for k, v := range extra {
			e.scopes[len(e.scopes)-1][k] = v
		}
	}
	out, sig := e.runStatements(block.Body.Statements)
	if sig != nil {
		switch sig.kind {
		case flowReturn:
			// return 前的输出一并保留（与函数调用一致：& { "a"; return "b" } 输出 a、b）
			return append(out, unwrapOutput(sig.value)...), nil
		case flowError:
			panic(sig) // 脚本块内终止错误向调用方传播（外层 try 可捕获，不在此落定）
		}
	}
	return out, nil
}

// LookupVar 按名字查变量，供内置 cmdlet 读取首选项这类作用域敏感变量。
func (e *Evaluator) LookupVar(name string) *object.PSObject {
	return e.lookupVar(name, "")
}

// ---- 作用域与变量 ----

func (e *Evaluator) pushScope() {
	e.scopes = append(e.scopes, map[string]*object.PSObject{})
}

func (e *Evaluator) popScope() {
	if len(e.scopes) > 1 {
		e.scopes = e.scopes[:len(e.scopes)-1]
	}
}

// lookupVar 按名字与作用域修饰符查变量，不区分大小写。
// scope 为空：自顶向下查（PowerShell 默认读语义）；"script"/"global"：只查全局（scopes[0]，即脚本作用域，本解释器脚本不推独立作用域）；"local"：只查当前（栈顶）作用域。
func (e *Evaluator) lookupVar(name, scope string) *object.PSObject {
	switch scope {
	case "script", "global":
		if v, ok := e.scopes[0][scopeVarKey(e.scopes[0], name)]; ok {
			return v
		}
	case "local":
		if v, ok := e.scopes[len(e.scopes)-1][scopeVarKey(e.scopes[len(e.scopes)-1], name)]; ok {
			return v
		}
		return object.Null()
	default:
		for i := len(e.scopes) - 1; i >= 0; i-- {
			if v, ok := e.scopes[i][scopeVarKey(e.scopes[i], name)]; ok {
				return v
			}
		}
	}
	if v, ok := e.Session.GetVar(name); ok {
		return v
	}
	return object.Null()
}

// scopeVarKey 取某层作用域的存储键：已存在（不区分大小写）沿用原大小写，否则用传入名。
// 写入沿用旧键，避免同名不同大小写并存。
func scopeVarKey(sc map[string]*object.PSObject, name string) string {
	if _, ok := sc[name]; ok {
		return name
	}
	for k := range sc {
		if strings.EqualFold(k, name) {
			return k
		}
	}
	return name
}

// setVar 按作用域修饰符写变量：script/global 写全局（scopes[0]），local 与默认写当前（栈顶）。
func (e *Evaluator) setVar(name, scope string, val *object.PSObject) error {
	if shell.IsReadOnlyVar(name) {
		return fmt.Errorf("%s", lang.T(lang.MsgReadonlyVar, name))
	}
	if strings.EqualFold(name, "ErrorActionPreference") {
		// 首选项名按规范大小写存储。
		// 空值视为恢复默认。
		name = "ErrorActionPreference"
		if val.IsNull() {
			if scope == "script" || scope == "global" {
				delete(e.scopes[0], name)
			} else {
				delete(e.scopes[len(e.scopes)-1], name)
			}
			return nil
		}
		if _, ok := shell.ParseErrorAction(val.String()); !ok {
			return fmt.Errorf("%s", lang.T(lang.MsgErrorActionPreferenceInvalid, val.String()))
		}
	}
	if scope == "script" || scope == "global" {
		e.scopes[0][scopeVarKey(e.scopes[0], name)] = val
	} else {
		top := e.scopes[len(e.scopes)-1]
		top[scopeVarKey(top, name)] = val
	}
	return nil
}

// writeError 把错误写到 stderr、累积进 $Error 并标记 $? 为 false。
func (e *Evaluator) writeError(err error) {
	if err == nil {
		return
	}
	msg := err.Error()
	fmt.Fprintf(e.hostErr, "%s : %s\n", e.Session.StyleName(), msg)
	e.Session.RecordError(msg)
	e.Session.LastSuccess = false
}

// printError 把未捕获的终止错误写到 stderr 并标记 $? 为 false。
// 错误记录已在抛出时累积，这里不再重复记录。
func (e *Evaluator) printError(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(e.hostErr, "%s : %s\n", e.Session.StyleName(), err.Error())
	e.Session.LastSuccess = false
}

// dispatchError 按生效动作分发错误。
// Stop 与 Inquire 记录后转为终止错误上抛，其余记为非终止错误。
func (e *Evaluator) dispatchError(action string, err error) {
	if err == nil {
		return
	}
	switch builtin.ResolveErrorAction(action, e.LookupVar) {
	case "stop", "inquire":
		rec := e.Session.RecordError(err.Error())
		e.Session.LastSuccess = false
		panic(&flowSignal{kind: flowError, value: rec})
	case "silentlycontinue":
		e.Session.RecordError(err.Error())
		e.Session.LastSuccess = false
	case "ignore":
		e.Session.LastSuccess = false
	default:
		e.writeError(err)
	}
}

// reportError 分发求值层错误。
// 求值层没有 -ErrorAction 上下文，只读 $ErrorActionPreference。
func (e *Evaluator) reportError(err error) {
	e.dispatchError("", err)
}

// throwError 抛出一个求值期终止错误：累积进 $Error 后以 flowError 上抛，外层 try 可捕获。
// 调用方已判定必须报错的场景用，不经首选项分发。
func (e *Evaluator) throwError(msg string) {
	rec := e.Session.RecordError(msg)
	exc := object.Object("System.RuntimeException", msg)
	exc.AddProp("Message", object.Str(msg))
	rec.AddProp("Exception", exc)
	panic(&flowSignal{kind: flowError, value: rec})
}

// ReportPanic 把顶层回收到的非控制流 panic 转为非终止错误。
// 控制流信号返回 false，交由调用方继续传播。
func (e *Evaluator) ReportPanic(r any) bool {
	if r == nil {
		return false
	}
	if _, ok := r.(*flowSignal); ok {
		return false
	}
	e.writeError(fmt.Errorf("%v", r))
	return true
}

// ---- 表达式求值 ----

func (e *Evaluator) evalValue(n ast.Node) *object.PSObject {
	switch v := n.(type) {
	case nil:
		return object.Null()
	case *ast.StrLit:
		return object.Str(v.Value)
	case *ast.StrTemplate:
		var sb strings.Builder
		for _, part := range v.Parts {
			sb.WriteString(e.evalValue(part).String())
		}
		return object.Str(sb.String())
	case *ast.Number:
		if v.IsInt {
			return object.Int(int64(v.Value))
		}
		return object.Float(v.Value)
	case *ast.BoolLit:
		return object.Bool(v.Value)
	case *ast.NullLit:
		return object.Null()
	case *ast.VarRef:
		return e.lookupVar(v.Name, v.Scope)
	case *ast.EnvRef:
		return object.Str(os.Getenv(v.Name))
	case *ast.BareWord:
		return object.Str(v.Value)
	case *ast.Paren:
		return e.evalValue(v.Inner)
	case *ast.Unary:
		val := e.evalValue(v.Operand)
		switch v.Op {
		case "-not", "!":
			return object.Bool(!val.Truthy())
		case "-":
			if n, ok := val.AsFloat(); ok {
				return object.Float(-n)
			}
			e.throwError(lang.T(lang.MsgArithmeticInvalid))
			return object.Null()
		}
		return val
	case *ast.Binary:
		return e.evalBinary(v)
	case *ast.Ternary:
		if e.evalValue(v.Cond).Truthy() {
			return e.evalValue(v.If)
		}
		return e.evalValue(v.Else)
	case *ast.ArrayLit:
		// @(...) ：
		//   唯一元素 → “确保是数组”：元素输出流直接作为结果（@($arr) 得到 $arr 的元素）。
		//   多个元素 → 各语句输出依次收集，数组值整体占一位（@(@("a"),"b") 内层数组保留）。
		//   管道元素 → 每项输出各占一位（@("x" | ForEach-Object { ... })）。
		if v.Flatten && len(v.Items) == 1 {
			if pl, ok := v.Items[0].(*ast.PipelineExpr); ok {
				return object.Array(e.evalPipeline(pl.Pipeline))
			}
			if cmd, ok := v.Items[0].(*ast.Command); ok {
				// 单个命令：按单元素管道执行，保留输出流原样（无输出即空流，$null 占一位）。
				return object.Array(e.evalPipeline(&ast.Pipeline{Commands: []*ast.Command{cmd}}))
			}
			// 单个纯表达式：直接包装（$null 占一位，@($null) 计 1，与 PowerShell 一致）。
			vv := e.evalValue(v.Items[0])
			if vv == nil {
				return object.Array(nil)
			}
			if vv.IsArray() {
				return object.Array(vv.ArrayItems())
			}
			return object.Array([]*object.PSObject{vv})
		}
		items := make([]*object.PSObject, 0, len(v.Items))
		for _, it := range v.Items {
			if v.Flatten {
				if pl, ok := it.(*ast.PipelineExpr); ok {
					items = append(items, e.evalPipeline(pl.Pipeline)...)
					continue
				}
			}
			items = append(items, e.evalValue(it))
		}
		return object.Array(items)
	case *ast.HashtableLit:
		entries := make([]object.HashEntry, 0, len(v.Pairs))
		for _, p := range v.Pairs {
			entries = append(entries, object.HashEntry{
				Key:   e.evalValue(p.Key).String(),
				Value: e.evalValue(p.Value),
			})
		}
		return object.Hashtable(entries)
	case *ast.TypeCast:
		// [pscustomobject]@{...}：哈希表条目变属性
		if strings.EqualFold(v.TypeName, "pscustomobject") {
			expr := e.evalValue(v.Expr)
			if entries, ok := expr.Value.([]object.HashEntry); ok {
				return object.PSCustomObject(entries)
			}
			return expr
		}
		// 无操作数的类型字面量：求值为类型名，供 -is/-as 与变量保存消费
		if v.Expr == nil {
			return object.Str(strings.ToLower(v.TypeName))
		}
		return e.convertValue(e.evalValue(v.Expr), v.TypeName)
	case *ast.StaticMember:
		var argv []*object.PSObject
		for _, a := range v.Args {
			argv = append(argv, e.evalValue(a))
		}
		if val, ok := e.staticMember(v.TypeName, v.Name, argv); ok {
			return val
		}
		e.reportError(fmt.Errorf("%s", lang.T(lang.MsgStaticMemberNotFound, strings.ToLower(strings.TrimPrefix(strings.ToLower(v.TypeName), "system.")), v.Name)))
		return object.Null()
	case *ast.MemberAccess:
		base := e.evalValue(v.Base)
		return e.memberProp(base, v.Prop)
	case *ast.MethodCall:
		return e.evalMethodCall(v)
	case *ast.Index:
		return e.evalIndex(v)
	case *ast.ScriptBlock:
		return object.ScriptBlock(v)
	case *ast.SubExpr:
		e.inCapture++
		out, sig := e.runStatements(v.Body.Statements)
		e.inCapture--
		if sig != nil {
			switch sig.kind {
			case flowReturn:
				// 子表达式整体作为输出流：return 前的输出保留（$( "a"; return "b" ) → a、b）
				return wrapSingle(append(out, unwrapOutput(sig.value)...))
			case flowError:
				panic(sig) // 子表达式里的终止错误向上传播（外层 try 可捕获）
			}
		}
		return wrapSingle(out)
	case *ast.Increment:
		if v.Target != nil {
			return e.incrTarget(v)
		}
		delta := int64(1)
		if v.Op == "--" {
			delta = -1
		}
		old, nv := e.incrOldNew(e.lookupVar(v.Var, v.Scope), delta)
		_ = e.setVar(v.Var, v.Scope, nv)
		return old
	case *ast.PipelineExpr:
		out := e.evalPipeline(v.Pipeline)
		return wrapSingle(out)
	case *ast.Command:
		// 表达式位的命令节点（& 调用进赋值右侧、子表达式等）：按单元素管道执行
		out := e.evalPipeline(&ast.Pipeline{Commands: []*ast.Command{v}})
		return wrapSingle(out)
	case *ast.PropertyRef:
		return object.Null()
	}
	return object.Null()
}

// incrTarget 对下标或属性目标做增减：读当前值、加减一、写回，返回新值。
// 数组整数下标、哈希表键、对象属性可写；其余报错。
func (e *Evaluator) incrTarget(v *ast.Increment) *object.PSObject {
	delta := int64(1)
	if v.Op == "--" {
		delta = -1
	}
	switch t := v.Target.(type) {
	case *ast.Index:
		return e.incrIndex(t, delta)
	case *ast.MemberAccess:
		return e.incrMember(t, delta)
	}
	e.reportError(fmt.Errorf("%s", lang.T(lang.MsgIncrTargetInvalid)))
	return object.Null()
}

// incrOldNew 返回加减前后的值：后缀自增取加一前的值，与 PowerShell 一致。
// 浮点保持浮点，$null 从 0 起，非数字报错。
func (e *Evaluator) incrOldNew(cur *object.PSObject, delta int64) (old, nv *object.PSObject) {
	if cur.TypeName == "Double" {
		f, _ := cur.AsFloat()
		return cur, object.Float(f + float64(delta))
	}
	if cur.IsNull() {
		return object.Int(0), object.Int(delta)
	}
	n, ok := cur.AsInt()
	if !ok {
		e.throwError(lang.T(lang.MsgArithmeticInvalid))
		return object.Null(), object.Null()
	}
	return object.Int(n), object.Int(n + delta)
}

// incrIndex 对数组整数下标或哈希表键做增减并写回。
// 负数下标从末尾数，与读取一致；多下标、字符串下标与越界报错。
func (e *Evaluator) incrIndex(t *ast.Index, delta int64) *object.PSObject {
	base := e.evalValue(t.Base)
	if base.TypeName == "Hashtable" {
		if entries, ok := base.Value.([]object.HashEntry); ok {
			key := e.evalValue(t.Index).String()
			for i := range entries {
				if strings.EqualFold(entries[i].Key, key) {
					old, nv := e.incrOldNew(entries[i].Value, delta)
					entries[i].Value = nv
					return old
				}
			}
		}
		e.throwError(lang.T(lang.MsgIncrTargetInvalid))
		return object.Null()
	}
	if !base.IsArray() {
		e.throwError(lang.T(lang.MsgIncrTargetInvalid))
		return object.Null()
	}
	items := base.ArrayItems()
	idx := e.evalValue(t.Index)
	if idx.IsArray() {
		e.throwError(lang.T(lang.MsgIncrTargetInvalid))
		return object.Null()
	}
	n, ok := idx.AsInt()
	if !ok {
		e.throwError(lang.T(lang.MsgIncrTargetInvalid))
		return object.Null()
	}
	if n < 0 {
		n = int64(len(items)) + n
	}
	if n < 0 || int(n) >= len(items) {
		e.throwError(lang.T(lang.MsgIndexOutOfRange))
		return object.Null()
	}
	old, nv := e.incrOldNew(items[n], delta)
	items[n] = nv
	return old
}

// incrMember 对对象属性做增减并写回；属性不存在报错。
func (e *Evaluator) incrMember(t *ast.MemberAccess, delta int64) *object.PSObject {
	base := e.evalValue(t.Base)
	if base == nil || base.IsNull() {
		e.throwError(lang.T(lang.MsgIncrTargetInvalid))
		return object.Null()
	}
	cur, ok := base.PropValue(t.Prop)
	if !ok {
		e.throwError(lang.T(lang.MsgIncrTargetInvalid))
		return object.Null()
	}
	old, nv := e.incrOldNew(cur, delta)
	base.SetProp(t.Prop, nv)
	return old
}

// wrapSingle 把输出列表包成单个值：0 → $null，1 → 该项，多 → 数组。
func wrapSingle(out []*object.PSObject) *object.PSObject {
	if len(out) == 0 {
		return object.Null()
	}
	if len(out) == 1 {
		return out[0]
	}
	return object.Array(out)
}

// unwrapOutput 把返回值展开为对象流（数组摊平）。
func unwrapOutput(v *object.PSObject) []*object.PSObject {
	if v == nil || v.IsNull() {
		return nil
	}
	if v.IsArray() {
		return v.ArrayItems()
	}
	return []*object.PSObject{v}
}

// ---- 属性与方法 ----

// memberProp 取对象的属性。
func (e *Evaluator) memberProp(base *object.PSObject, prop string) *object.PSObject {
	if base == nil {
		return object.Null()
	}
	if v, ok := base.PropValue(prop); ok {
		return v
	}
	// Count/Length：有值标量为 1，$null 为 0
	switch strings.ToLower(prop) {
	case "count", "length":
		if base.IsNull() {
			return object.Int(0)
		}
		return object.Int(1)
	}
	return object.Null()
}

// propertyOf 在对象上取属性（用于 Where-Object Length 等裸属性）。
func (e *Evaluator) propertyOf(obj *object.PSObject, name string) *object.PSObject {
	if obj == nil {
		return object.Null()
	}
	return e.memberProp(obj, name)
}

// evalMethodCall 调用字符串/数组/哈希表方法。
func (e *Evaluator) evalMethodCall(m *ast.MethodCall) *object.PSObject {
	base := e.evalValue(m.Base)
	args := make([]*object.PSObject, 0, len(m.Args))
	for _, a := range m.Args {
		args = append(args, e.evalValue(a))
	}
	arg := func(i int) *object.PSObject {
		if i < len(args) {
			return args[i]
		}
		return object.Null()
	}
	switch base.TypeName {
	case "String":
		s := base.String()
		switch strings.ToLower(m.Name) {
		case "toupper":
			return object.Str(strings.ToUpper(s))
		case "tolower":
			return object.Str(strings.ToLower(s))
		case "trim":
			return object.Str(strings.TrimSpace(s))
		case "trimstart":
			// 无参清前导空白（PowerShell 语义）；有参按字符集裁剪
			if len(args) == 0 {
				return object.Str(strings.TrimLeftFunc(s, unicode.IsSpace))
			}
			return object.Str(strings.TrimLeft(s, arg(0).String()))
		case "trimend":
			// 无参清尾随空白；有参按字符集裁剪
			if len(args) == 0 {
				return object.Str(strings.TrimRightFunc(s, unicode.IsSpace))
			}
			return object.Str(strings.TrimRight(s, arg(0).String()))
		case "contains":
			return object.Bool(strings.Contains(s, arg(0).String()))
		case "startswith":
			return object.Bool(strings.HasPrefix(s, arg(0).String()))
		case "endswith":
			return object.Bool(strings.HasSuffix(s, arg(0).String()))
		case "indexof":
			return object.Int(runeIndex(s, arg(0).String()))
		case "lastindexof":
			// 一参为子串；两参为 子串,起始下标（.NET 搜索范围含起始下标，向左找）
			if len(args) == 0 {
				return object.Null()
			}
			sub := arg(0).String()
			if len(args) >= 2 {
				from, ok := arg(1).AsInt()
				if !ok {
					return object.Null()
				}
				if from < 0 {
					return object.Int(-1)
				}
				r := []rune(s)
				if int(from) >= len(r) {
					return object.Int(runeLastIndex(s, sub))
				}
				return object.Int(runeLastIndex(string(r[:from+1]), sub))
			}
			return object.Int(runeLastIndex(s, sub))
		case "substring":
			start, ok := arg(0).AsInt()
			if !ok {
				return object.Null()
			}
			r := []rune(s)
			if len(args) >= 2 {
				ln, _ := args[1].AsInt()
				if start < 0 || ln < 0 || int(start) > len(r) || int(start)+int(ln) > len(r) {
					e.throwError(lang.T(lang.MsgSubstringOutOfRange))
					return object.Null()
				}
				return object.Str(string(r[start : start+ln]))
			}
			if start < 0 || int(start) > len(r) {
				e.throwError(lang.T(lang.MsgSubstringOutOfRange))
				return object.Null()
			}
			return object.Str(string(r[start:]))
		case "replace":
			return object.Str(strings.ReplaceAll(s, arg(0).String(), arg(1).String()))
		case "remove":
			// 一参从下标删到末尾；两参为 起始下标,删除长度（对齐 .NET Remove）
			start, ok := arg(0).AsInt()
			if !ok {
				return object.Null()
			}
			if start < 0 {
				return object.Null()
			}
			r := []rune(s)
			if len(args) >= 2 {
				ln, lok := arg(1).AsInt()
				if !lok || ln < 0 {
					return object.Null()
				}
				if int(start) >= len(r) {
					return object.Str(s)
				}
				end := start + ln
				if end > int64(len(r)) {
					end = int64(len(r))
				}
				return object.Str(string(r[:start]) + string(r[end:]))
			}
			if int(start) >= len(r) {
				return object.Str(s)
			}
			return object.Str(string(r[:start]))
		case "padleft":
			// 一参为总宽（空格补齐）；两参为 总宽,填充字符（只取首个字符，对齐 .NET char 参数）
			total, ok := arg(0).AsInt()
			if !ok || total < 0 {
				return object.Null()
			}
			pad := " "
			if len(args) >= 2 {
				ps := arg(1).String()
				if ps != "" {
					pad = string([]rune(ps)[0])
				}
			}
			if int(total) <= len([]rune(s)) {
				return object.Str(s)
			}
			return object.Str(strings.Repeat(pad, int(total)-len([]rune(s))) + s)
		case "padright":
			// 同 PadLeft，右对齐方向相反（右侧补齐）
			total, ok := arg(0).AsInt()
			if !ok || total < 0 {
				return object.Null()
			}
			pad := " "
			if len(args) >= 2 {
				ps := arg(1).String()
				if ps != "" {
					pad = string([]rune(ps)[0])
				}
			}
			if int(total) <= len([]rune(s)) {
				return object.Str(s)
			}
			return object.Str(s + strings.Repeat(pad, int(total)-len([]rune(s))))
		case "insert":
			// 在指定下标插入子串（对齐 .NET Insert）
			pos, ok := arg(0).AsInt()
			if !ok || pos < 0 {
				return object.Null()
			}
			if int(pos) > len([]rune(s)) {
				return object.Null()
			}
			r := []rune(s)
			return object.Str(string(r[:pos]) + arg(1).String() + string(r[pos:]))
		case "split":
			// 无参按任意空白分割且合并连续空白（.NET Split() 无参语义，与 strings.Fields 一致）
			if len(args) == 0 {
				parts := strings.Fields(s)
				items := make([]*object.PSObject, len(parts))
				for i, p := range parts {
					items[i] = object.Str(p)
				}
				return object.Array(items)
			}
			sep := arg(0).String()
			if sep == "" {
				sep = " "
			}
			parts := strings.Split(s, sep)
			items := make([]*object.PSObject, len(parts))
			for i, p := range parts {
				items[i] = object.Str(p)
			}
			return object.Array(items)
		}
	case "DateTime":
		if t, ok := base.Value.(time.Time); ok {
			switch strings.ToLower(m.Name) {
			case "addhours", "addminutes", "addseconds", "adddays":
				n, nok := arg(0).AsFloat()
				if !nok {
					return object.Null()
				}
				var d time.Duration
				switch strings.ToLower(m.Name) {
				case "addhours":
					d = time.Duration(n * float64(time.Hour))
				case "addminutes":
					d = time.Duration(n * float64(time.Minute))
				case "addseconds":
					d = time.Duration(n * float64(time.Second))
				case "adddays":
					d = time.Duration(n * 24 * float64(time.Hour))
				}
				return object.DateTime(t.Add(d))
			case "addyears", "addmonths":
				n, nok := arg(0).AsInt()
				if !nok {
					return object.Null()
				}
				if strings.ToLower(m.Name) == "addyears" {
					return object.DateTime(t.AddDate(int(n), 0, 0))
				}
				return object.DateTime(t.AddDate(0, int(n), 0))
			case "toshortdatestring":
				// 使用 en-US 格式
				return object.Str(t.Format("1/2/2006"))
			case "tolongdatestring":
				return object.Str(t.Format("Monday, January 2, 2006"))
			case "toshorttimestring":
				return object.Str(t.Format("3:04 PM"))
			case "tolongtimestring":
				return object.Str(t.Format("3:04:05 PM"))
			case "tostring":
				return object.Str(t.Format("1/2/2006 3:04:05 PM"))
			case "touniversaltime":
				return object.DateTime(t.UTC())
			case "tolocaltime":
				return object.DateTime(t.Local())
			case "tofiletime":
				// Windows 文件时间：1601-01-01 起 100 纳秒刻度数
				const epochOffset = 116444736000000000
				u := t.UTC()
				// 用秒数加纳秒余数计算，不用 UnixNano：后者在 1672 年前与 2262 年后会溢出
				return object.Int(u.Unix()*10000000 + int64(u.Nanosecond())/100 + epochOffset)
			}
		}
	case "Object[]":
		items := base.ArrayItems()
		switch strings.ToLower(m.Name) {
		case "count":
			return object.Int(int64(len(items)))
		case "contains":
			for _, it := range items {
				if compareEq(it, arg(0)) {
					return object.Bool(true)
				}
			}
			return object.Bool(false)
		case "clear":
			// $Error.Clear()：带动态视图标记的数组把清空落到会话的错误记录本体
			if _, ok := base.PropValue(shell.ErrorViewMarker); ok {
				e.Session.ClearErrorRecords()
			}
			return object.Null()
		case "removeat":
			// $Error.RemoveAt(n)：同样落到会话的错误记录本体；越界报错
			if _, ok := base.PropValue(shell.ErrorViewMarker); ok {
				if idx, ok := arg(0).AsInt(); ok {
					if !e.Session.RemoveErrorRecord(idx) {
						e.reportError(fmt.Errorf("%s", lang.T(lang.MsgIndexOutOfRange)))
					}
					return object.Null()
				}
				return object.Null()
			}
		}
	case "Hashtable":
		switch strings.ToLower(m.Name) {
		case "contains":
			if entries, ok := base.Value.([]object.HashEntry); ok {
				for _, en := range entries {
					if strings.EqualFold(en.Key, arg(0).String()) {
						return object.Bool(true)
					}
				}
			}
			return object.Bool(false)
		case "count":
			if entries, ok := base.Value.([]object.HashEntry); ok {
				return object.Int(int64(len(entries)))
			}
		}
	case "ScriptBlock":
		// .Invoke(实参...) 执行脚本块，输出收集成数组返回
		if strings.EqualFold(m.Name, "invoke") {
			if node, ok := base.Value.(*ast.ScriptBlock); ok {
				out := e.invokeScriptBlock(node, callArgs{posVals: args, namedVals: map[string]*object.PSObject{}, namedSwitch: map[string]bool{}}, nil)
				return object.Array(out)
			}
		}
	}
	// 未注册 ToString 的类型返回显示字符串。
	if strings.EqualFold(m.Name, "tostring") && len(args) == 0 {
		return object.Str(base.String())
	}
	return object.Null()
}

// evalIndex 索引访问。
func (e *Evaluator) evalIndex(i *ast.Index) *object.PSObject {
	base := e.evalValue(i.Base)
	idx := e.evalValue(i.Index)
	if base.IsNull() {
		return object.Null()
	}
	if base.IsArray() {
		items := base.ArrayItems()
		// 多下标（$a[1..2]、$a[0,2]、$a[1..2,0]）：下标表达式是数组时逐元素取值，嵌套数组展平，越界补 $null。
		if idx.IsArray() {
			return indexSelect(idx, func(n int64) *object.PSObject {
				return arrayItemAt(items, n)
			})
		}
		n, ok := idx.AsInt()
		if !ok {
			return object.Null()
		}
		return arrayItemAt(items, n)
	}
	if base.TypeName == "String" {
		s := base.String()
		if idx.IsArray() {
			return indexSelect(idx, func(n int64) *object.PSObject {
				return stringItemAt(s, n)
			})
		}
		n, ok := idx.AsInt()
		if !ok {
			return object.Null()
		}
		return stringItemAt(s, n)
	}
	if base.TypeName == "Hashtable" {
		if entries, ok := base.Value.([]object.HashEntry); ok {
			key := idx.String()
			for _, en := range entries {
				if strings.EqualFold(en.Key, key) {
					return en.Value
				}
			}
		}
		return object.Null()
	}
	// 标量（非集合）索引：[0] 返回自身，其余下标 $null
	if n, ok := idx.AsInt(); ok && n == 0 {
		return base
	}
	return object.Null()
}

// indexSelect 按下标数组（范围/逗号表达式）逐元素取值；
// 单个结果展开为标量，多个结果组成数组。
func indexSelect(idx *object.PSObject, pick func(n int64) *object.PSObject) *object.PSObject {
	var sel []*object.PSObject
	for _, id := range flattenIndices(idx) {
		sel = append(sel, pick(id))
	}
	if len(sel) == 1 {
		return sel[0]
	}
	return object.Array(sel)
}

// flattenIndices 把下标表达式递归展平为整数下标列表（$a[1..2,0] 的嵌套数组展开）。
func flattenIndices(idx *object.PSObject) []int64 {
	if !idx.IsArray() {
		if n, ok := idx.AsInt(); ok {
			return []int64{n}
		}
		return nil
	}
	var out []int64
	for _, it := range idx.ArrayItems() {
		out = append(out, flattenIndices(it)...)
	}
	return out
}

// arrayItemAt 取数组元素：负数从末尾数，越界返回 $null。
func arrayItemAt(items []*object.PSObject, n int64) *object.PSObject {
	if n < 0 {
		n = int64(len(items)) + n
	}
	if n >= 0 && int(n) < len(items) {
		return items[n]
	}
	return object.Null()
}

// stringItemAt 取字符串字符：负数从末尾数，越界返回空串；下标按字符计。
func stringItemAt(s string, n int64) *object.PSObject {
	r := []rune(s)
	if n < 0 {
		n = int64(len(r)) + n
	}
	if n >= 0 && int(n) < len(r) {
		return object.Str(string(r[n]))
	}
	return object.Str("")
}

// runeIndex 返回子串首现的字符下标，未找到返回 -1；字节下标转字符计数。
func runeIndex(s, sub string) int64 {
	b := strings.Index(s, sub)
	if b < 0 {
		return -1
	}
	return int64(utf8.RuneCountInString(s[:b]))
}

// runeLastIndex 返回子串末现的字符下标，未找到返回 -1；字节下标转字符计数。
func runeLastIndex(s, sub string) int64 {
	b := strings.LastIndex(s, sub)
	if b < 0 {
		return -1
	}
	return int64(utf8.RuneCountInString(s[:b]))
}

// ---- 二元运算 ----

func (e *Evaluator) evalBinary(b *ast.Binary) *object.PSObject {
	switch b.Op {
	case "-and":
		if !e.evalValue(b.L).Truthy() {
			return object.Bool(false)
		}
		return object.Bool(e.evalValue(b.R).Truthy())
	case "-or":
		if e.evalValue(b.L).Truthy() {
			return object.Bool(true)
		}
		return object.Bool(e.evalValue(b.R).Truthy())
	case "-xor":
		return object.Bool(e.evalValue(b.L).Truthy() != e.evalValue(b.R).Truthy())
	case "??":
		// 空合并：左值非 $null 取左值，否则取右值（短路）
		lv := e.evalValue(b.L)
		if lv.IsNull() {
			return e.evalValue(b.R)
		}
		return lv
	case "-f":
		return e.formatOp(e.evalValue(b.L), e.evalValue(b.R))
	}
	l := e.evalValue(b.L)
	r := e.evalValue(b.R)
	return e.binaryOp(b.Op, l, r)
}


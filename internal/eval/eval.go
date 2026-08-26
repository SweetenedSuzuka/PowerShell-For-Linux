// Package eval 实现 PowerShell 求值器：表达式、语句、管道、命令调度、控制流、函数调用与外部命令。
package eval

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"powershell/internal/ast"
	"powershell/internal/builtin"
	"powershell/internal/external"
	"powershell/internal/lang"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// flowKind 是控制流信号类型。
type flowKind int

const (
	flowBreak flowKind = iota
	flowContinue
	flowReturn
	flowExit
	flowError // 终止错误（throw / 未捕获错误）
)

// flowSignal 用 panic/recover 传递 break/continue/return/exit/error。
type flowSignal struct {
	kind  flowKind
	value *object.PSObject // return 值 / 错误记录
	code  int              // exit 码
	out   []*object.PSObject // panic 前已产生的输出（传播时保留，如 throw 前循环的输出）
}

// RecoverError 提取 panic 值里的终止错误，供 main.go 在 -File 顶层打印。
// 非错误 panic 返回 nil，交由调用方继续传播。
func RecoverError(r any) error {
	if fs, ok := r.(*flowSignal); ok && fs.kind == flowError {
		return fmt.Errorf("%s", fs.value.String())
	}
	return nil
}

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
			return nil, fmt.Errorf("%s", sig.value.String())
		}
	}
	return out, nil
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

// lookupVar 按名字与作用域修饰符查变量。
// scope 为空：自顶向下查（PowerShell 默认读语义）；"script"/"global"：只查全局（scopes[0]，即脚本作用域，本解释器脚本不推独立作用域）；"local"：只查当前（栈顶）作用域。
func (e *Evaluator) lookupVar(name, scope string) *object.PSObject {
	switch scope {
	case "script", "global":
		if v, ok := e.scopes[0][name]; ok {
			return v
		}
	case "local":
		if v, ok := e.scopes[len(e.scopes)-1][name]; ok {
			return v
		}
		return object.Null()
	default:
		for i := len(e.scopes) - 1; i >= 0; i-- {
			if v, ok := e.scopes[i][name]; ok {
				return v
			}
		}
	}
	if v, ok := e.Session.GetVar(name); ok {
		return v
	}
	return object.Null()
}

// setVar 按作用域修饰符写变量：script/global 写全局（scopes[0]），local 与默认写当前（栈顶）。
func (e *Evaluator) setVar(name, scope string, val *object.PSObject) error {
	if shell.IsReadOnlyVar(name) {
		return fmt.Errorf("%s", lang.T(lang.MsgReadonlyVar, name))
	}
	if scope == "script" || scope == "global" {
		e.scopes[0][name] = val
	} else {
		e.scopes[len(e.scopes)-1][name] = val
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
			return object.Float(0)
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
			return object.Array(unwrapOutput(e.evalValue(v.Items[0])))
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
		e.writeError(fmt.Errorf("%s", lang.T(lang.MsgStaticMemberNotFound, strings.ToLower(strings.TrimPrefix(strings.ToLower(v.TypeName), "system.")), v.Name)))
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
		cur := e.lookupVar(v.Var, v.Scope)
		// 浮点变量按浮点增减：$i = 0.5; $i++ → 1.5（AsInt 会把 0.5 截断成 0）
		if cur.TypeName == "Double" {
			f, _ := cur.AsFloat()
			if v.Op == "++" {
				f++
			} else {
				f--
			}
			_ = e.setVar(v.Var, v.Scope, object.Float(f))
			return object.Float(f)
		}
		n, _ := cur.AsInt()
		if v.Op == "++" {
			n++
		} else {
			n--
		}
		_ = e.setVar(v.Var, v.Scope, object.Int(n))
		return object.Int(n)
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
			return object.Int(int64(strings.Index(s, arg(0).String())))
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
				if int(from) >= len(s) {
					return object.Int(int64(strings.LastIndex(s, sub)))
				}
				return object.Int(int64(strings.LastIndex(s[:from+1], sub)))
			}
			return object.Int(int64(strings.LastIndex(s, sub)))
		case "substring":
			start, ok := arg(0).AsInt()
			if !ok {
				return object.Null()
			}
			if len(args) >= 2 {
				ln, _ := args[1].AsInt()
				if int(start)+int(ln) <= len(s) {
					return object.Str(s[start : start+ln])
				}
				return object.Str(s[start:])
			}
			if int(start) <= len(s) {
				return object.Str(s[start:])
			}
			return object.Str("")
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
			if len(args) >= 2 {
				ln, lok := arg(1).AsInt()
				if !lok || ln < 0 {
					return object.Null()
				}
				if int(start) >= len(s) {
					return object.Str(s)
				}
				end := start + ln
				if end > int64(len(s)) {
					end = int64(len(s))
				}
				return object.Str(s[:start] + s[end:])
			}
			if int(start) >= len(s) {
				return object.Str(s)
			}
			return object.Str(s[:start])
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
						e.writeError(fmt.Errorf("%s", lang.T(lang.MsgIndexOutOfRange)))
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
	// 未注册 ToString 的类型走通用兜底：返回显示字符串
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

// stringItemAt 取字符串字符：负数从末尾数，越界返回空串。
func stringItemAt(s string, n int64) *object.PSObject {
	if n < 0 {
		n = int64(len(s)) + n
	}
	if n >= 0 && int(n) < len(s) {
		return object.Str(string(s[n]))
	}
	return object.Str("")
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

func (e *Evaluator) binaryOp(op string, l, r *object.PSObject) *object.PSObject {
	switch op {
	case "+":
		return addOp(l, r)
	case "-":
		return numOp(l, r, func(a, b float64) float64 { return a - b })
	case "*":
		return mulOp(l, r)
	case "/":
		// 除数为零报错并置 $?=false
		if rf, ok := r.AsFloat(); ok && rf == 0 {
			e.writeError(fmt.Errorf("%s", lang.T(lang.MsgDivideByZero)))
			return object.Null()
		}
		return numOp(l, r, func(a, b float64) float64 { return a / b })
	case "%":
		// 模数为零同样报错
		if rf, ok := r.AsFloat(); ok && rf == 0 {
			e.writeError(fmt.Errorf("%s", lang.T(lang.MsgDivideByZero)))
			return object.Null()
		}
		return numOp(l, r, func(a, b float64) float64 { return math.Mod(a, b) })
	case "..":
		return rangeOp(l, r)
	case "-eq", "-ne", "-lt", "-le", "-gt", "-ge", "-like", "-notlike", "-match", "-notmatch",
		"-ceq", "-cne", "-clt", "-cle", "-cgt", "-cge", "-clike", "-cnotlike", "-cmatch", "-cnotmatch":
		// 数组左值 → 逐元素过滤
		if l.IsArray() {
			var matches []*object.PSObject
			for _, it := range l.ArrayItems() {
				if pairMatch(op, it, r) {
					matches = append(matches, it)
				}
			}
			if len(matches) == 0 {
				return object.Null()
			}
			if len(matches) == 1 {
				return matches[0]
			}
			return object.Array(matches)
		}
		if op == "-match" || op == "-cmatch" {
			// 标量 -match/-cmatch 匹配成功后填充 $Matches（数组左值在上面的过滤分支里不填充）
			return e.evalMatch(op, l, r)
		}
		return object.Bool(pairMatch(op, l, r))
	case "-contains":
		return object.Bool(arrayContains(l, r))
	case "-notcontains":
		return object.Bool(!arrayContains(l, r))
	case "-ccontains":
		return object.Bool(arrayContainsCase(l, r))
	case "-cnotcontains":
		return object.Bool(!arrayContainsCase(l, r))
	case "-in":
		return object.Bool(arrayContains(r, l))
	case "-notin":
		return object.Bool(!arrayContains(r, l))
	case "-join":
		return joinOp(l, r)
	case "-split":
		return splitOp(l, r)
	case "-replace":
		return replaceOp(l, r)
	case "-is":
		return object.Bool(typeIs(l, r))
	case "-isnot":
		return object.Bool(!typeIs(l, r))
	case "-as":
		return asOp(l, r)
	case "-band":
		return bitOp(l, r, func(a, b int64) int64 { return a & b })
	case "-bor":
		return bitOp(l, r, func(a, b int64) int64 { return a | b })
	case "-bxor":
		return bitOp(l, r, func(a, b int64) int64 { return a ^ b })
	case "-shl":
		return bitOp(l, r, func(a, b int64) int64 { return a << uint(b&63) })
	case "-shr":
		return bitOp(l, r, func(a, b int64) int64 { return a >> uint(b&63) })
	}
	return object.Bool(false)
}

// pairMatch 判断单个元素与右值是否满足比较运算符。
// PowerShell 默认比较大小写不敏感；-c* 前缀是大小写敏感变体。
func pairMatch(op string, l, r *object.PSObject) bool {
	switch op {
	case "-eq":
		return compareEq(l, r)
	case "-ne":
		return !compareEq(l, r)
	case "-ceq":
		return caseSensitiveEq(l, r)
	case "-cne":
		return !caseSensitiveEq(l, r)
	case "-lt":
		return compareOrder(l, r) < 0
	case "-le":
		return compareOrder(l, r) <= 0
	case "-gt":
		return compareOrder(l, r) > 0
	case "-ge":
		return compareOrder(l, r) >= 0
	case "-clt":
		return caseSensitiveOrder(l, r) < 0
	case "-cle":
		return caseSensitiveOrder(l, r) <= 0
	case "-cgt":
		return caseSensitiveOrder(l, r) > 0
	case "-cge":
		return caseSensitiveOrder(l, r) >= 0
	case "-like":
		return object.WildcardMatchFold(r.String(), l.String())
	case "-notlike":
		return !object.WildcardMatchFold(r.String(), l.String())
	case "-clike":
		return object.WildcardMatch(r.String(), l.String())
	case "-cnotlike":
		return !object.WildcardMatch(r.String(), l.String())
	case "-match", "-notmatch", "-cmatch", "-cnotmatch":
		re, err := compilePattern(op, r.String())
		if err != nil {
			return false
		}
		matched := re.MatchString(l.String())
		if op == "-notmatch" || op == "-cnotmatch" {
			return !matched
		}
		return matched
	}
	return false
}

// compilePattern 编译 -match/-cmatch 用的正则，并把 .NET 命名组语法转成 Go 语法。
// -match/-notmatch 在模式前加 (?i) 实现大小写不敏感。
func compilePattern(op, pattern string) (*regexp.Regexp, error) {
	p := translateNamedGroups(pattern)
	if op == "-match" || op == "-notmatch" {
		p = "(?i)" + p
	}
	return regexp.Compile(p)
}

// translateNamedGroups 把 .NET 的命名组语法 (?<name>...) 转成 Go 的 (?P<name>...)。
// (?<= 与 (?<! 是环视，Go 不支持，原样保留，交给编译阶段报错。
func translateNamedGroups(pattern string) string {
	var sb strings.Builder
	for i := 0; i < len(pattern); {
		if i+2 < len(pattern) && pattern[i] == '(' && pattern[i+1] == '?' && pattern[i+2] == '<' {
			if i+3 < len(pattern) && (pattern[i+3] == '=' || pattern[i+3] == '!') {
				sb.WriteString("(?<") // 环视，保留原样
				i += 3
				continue
			}
			sb.WriteString("(?P<") // 命名组
			i += 3
			continue
		}
		sb.WriteByte(pattern[i])
		i++
	}
	return sb.String()
}

// evalMatch 处理标量的 -match/-cmatch：匹配成功后把捕获组写入 $Matches。
// 不匹配时不动 $Matches（与真 PowerShell 一致）；数组左值不经过这里。
func (e *Evaluator) evalMatch(op string, l, r *object.PSObject) *object.PSObject {
	re, err := compilePattern(op, r.String())
	if err != nil {
		return object.Bool(false)
	}
	idx := re.FindStringSubmatchIndex(l.String())
	if idx == nil {
		return object.Bool(false)
	}
	e.Session.Matches = buildMatches(re, l.String(), idx)
	return object.Bool(true)
}

// buildMatches 由正则匹配结果构造 $Matches 哈希表。
// 键规则与真 PowerShell 一致："0" 是整体匹配；命名组用组名；
// 未命名组按其在未命名组中的序号（从 1 起）作键；未参与的组不写入。
func buildMatches(re *regexp.Regexp, s string, idx []int) *object.PSObject {
	names := re.SubexpNames()
	entries := []object.HashEntry{{Key: "0", Value: object.Str(s[idx[0]:idx[1]])}}
	unnamed := 0
	for i := 1; i < len(names); i++ {
		key := names[i]
		if key == "" {
			unnamed++ // 未命名组按序号计数，未参与的也占号（与 .NET 一致）
			key = strconv.Itoa(unnamed)
		}
		if idx[2*i] < 0 {
			continue // 未参与的组不写入 $Matches
		}
		entries = append(entries, object.HashEntry{Key: key, Value: object.Str(s[idx[2*i]:idx[2*i+1]])})
	}
	return object.Hashtable(entries)
}

// caseSensitiveEq 大小写敏感相等（-ceq）。右操作数按左操作数类型转换，左为数字则两边按数字，左为字符串则两边按字符串，左为 $null 则只与 $null 相等。
func caseSensitiveEq(l, r *object.PSObject) bool {
	switch l.TypeName {
	case "Int", "Double", "Boolean":
		if ln, ok := l.AsFloat(); ok {
			if rn, ok2 := r.AsFloat(); ok2 {
				return ln == rn
			}
		}
	case "String":
		if r.IsNull() {
			return false
		}
		return l.String() == r.String()
	case "Null":
		return r.IsNull()
	}
	return l.String() == r.String()
}

// caseSensitiveOrder 大小写敏感顺序（-clt 等）。右操作数按左操作数类型参与比较：
// 左为数字则按数字，否则按字符串（"5" -clt "10" 是字典序，结果为 False）。
func caseSensitiveOrder(l, r *object.PSObject) int {
	switch l.TypeName {
	case "Int", "Double", "Boolean":
		if ln, ok := l.AsFloat(); ok {
			if rn, ok2 := r.AsFloat(); ok2 {
				if ln < rn {
					return -1
				}
				if ln > rn {
					return 1
				}
				return 0
			}
		}
	}
	return strings.Compare(l.String(), r.String())
}

// compareEq 相等判断（-eq）。右操作数按左操作数类型转换，左为数字则两边按数字，左为字符串则两边按字符串，左为 $null 则只与 $null 相等。
func compareEq(l, r *object.PSObject) bool {
	switch l.TypeName {
	case "Int", "Double", "Boolean":
		if ln, ok := l.AsFloat(); ok {
			if rn, ok2 := r.AsFloat(); ok2 {
				return ln == rn
			}
		}
	case "String":
		if r.IsNull() {
			return false
		}
		return strings.EqualFold(l.String(), r.String())
	case "Null":
		return r.IsNull()
	}
	return strings.EqualFold(l.String(), r.String())
}

// compareOrder 判断 -lt/-le/-gt/-ge 的顺序。右操作数按左操作数类型参与比较——左为数字则两边按数字（5 -lt "10" 为 True），左为字符串则两边按字符串（"5" -lt "10" 是字典序比较，结果为 False）。
func compareOrder(l, r *object.PSObject) int {
	switch l.TypeName {
	case "Int", "Double", "Boolean":
		if ln, ok := l.AsFloat(); ok {
			if rn, ok2 := r.AsFloat(); ok2 {
				if ln < rn {
					return -1
				}
				if ln > rn {
					return 1
				}
				return 0
			}
		}
	}
	return strings.Compare(strings.ToLower(l.String()), strings.ToLower(r.String()))
}

func arrayContains(arr, elem *object.PSObject) bool {
	for _, it := range arr.ArrayItems() {
		if compareEq(it, elem) {
			return true
		}
	}
	return false
}

// arrayContainsCase 大小写敏感版（-ccontains）。
func arrayContainsCase(arr, elem *object.PSObject) bool {
	for _, it := range arr.ArrayItems() {
		if caseSensitiveEq(it, elem) {
			return true
		}
	}
	return false
}

func addOp(l, r *object.PSObject) *object.PSObject {
	if l.IsArray() {
		return object.Array(append(l.ArrayItems(), r))
	}
	if r.IsArray() {
		return object.Array(append([]*object.PSObject{l}, r.ArrayItems()...))
	}
	// 任一操作数是浮点 → 浮点加法（整型路径会把 0.5 截断成 0，如 2 + 1/2）
	if l.TypeName == "Double" || r.TypeName == "Double" {
		if lf, ok := l.AsFloat(); ok {
			if rf, ok2 := r.AsFloat(); ok2 {
				return object.Float(lf + rf)
			}
		}
	}
	if li, ok := l.AsInt(); ok {
		if ri, ok2 := r.AsInt(); ok2 {
			return object.Int(li + ri)
		}
	}
	if lf, ok := l.AsFloat(); ok {
		if rf, ok2 := r.AsFloat(); ok2 {
			return object.Float(lf + rf)
		}
	}
	return object.Str(l.String() + r.String())
}

func numOp(l, r *object.PSObject, fn func(a, b float64) float64) *object.PSObject {
	lf, ok := l.AsFloat()
	if !ok {
		return object.Null()
	}
	rf, ok2 := r.AsFloat()
	if !ok2 {
		return object.Null()
	}
	return object.Float(fn(lf, rf))
}

func mulOp(l, r *object.PSObject) *object.PSObject {
	if l.TypeName == "String" {
		if n, ok := r.AsInt(); ok && n >= 0 {
			return object.Str(strings.Repeat(l.String(), int(n)))
		}
	}
	if r.TypeName == "String" {
		if n, ok := l.AsInt(); ok && n >= 0 {
			return object.Str(strings.Repeat(r.String(), int(n)))
		}
	}
	lf, ok := l.AsFloat()
	if !ok {
		return object.Null()
	}
	rf, ok2 := r.AsFloat()
	if !ok2 {
		return object.Null()
	}
	return object.Float(lf * rf)
}

func rangeOp(l, r *object.PSObject) *object.PSObject {
	ls, ok := l.AsInt()
	if !ok {
		return object.Null()
	}
	rs, ok2 := r.AsInt()
	if !ok2 {
		return object.Null()
	}
	var items []*object.PSObject
	if ls <= rs {
		for i := ls; i <= rs; i++ {
			items = append(items, object.Int(i))
		}
	} else {
		for i := ls; i >= rs; i-- {
			items = append(items, object.Int(i))
		}
	}
	return object.Array(items)
}

// formatOp 实现 .NET 风格格式串："{模板}" -f 值[, 值...]。
// 支持 {N}、{N,宽度}（空格对齐）、{N:规格}（D 十进制补零、X/x 十六进制、F 定点小数、N 千分位）。
// {{ 与 }} 转义字面大括号；未知规格退化为原样字符串；下标越界 → 报错并置 $?=false。
func (e *Evaluator) formatOp(f, args *object.PSObject) *object.PSObject {
	format := f.String()
	items := flattenArgs(args)
	var sb strings.Builder
	for i := 0; i < len(format); i++ {
		c := format[i]
		if c != '{' {
			sb.WriteByte(c)
			continue
		}
		if i+1 < len(format) && format[i+1] == '{' {
			sb.WriteByte('{')
			i++
			continue
		}
		j := strings.IndexByte(format[i+1:], '}')
		if j < 0 {
			sb.WriteString(format[i:])
			break
		}
		inner := format[i+1 : i+1+j]
		idxPart := inner
		spec := ""
		if k := strings.IndexByte(inner, ':'); k >= 0 {
			idxPart, spec = inner[:k], inner[k+1:]
		}
		alignPart := idxPart
		widthStr := ""
		if k := strings.IndexByte(idxPart, ','); k >= 0 {
			alignPart, widthStr = idxPart[:k], idxPart[k+1:]
		}
		idx, err := strconv.Atoi(strings.TrimSpace(alignPart))
		if err != nil || idx < 0 || idx >= len(items) {
			e.writeError(fmt.Errorf("%s", lang.T(lang.MsgFormatIndexOut, inner, len(items))))
			return object.Null()
		}
		s := formatArg(items[idx], spec)
		if w, werr := strconv.Atoi(strings.TrimSpace(widthStr)); werr == nil {
			aw := int(math.Abs(float64(w)))
			if len(s) < aw {
				pad := strings.Repeat(" ", aw-len(s))
				if w < 0 {
					s += pad
				} else {
					s = pad + s
				}
			}
		}
		sb.WriteString(s)
		i += j + 1
	}
	return object.Str(sb.String())
}

// flattenArgs 把 -f 的参数数组展平为位置参数列表（范围/括号数组/变量数组都摊开）。
func flattenArgs(v *object.PSObject) []*object.PSObject {
	if !v.IsArray() {
		return []*object.PSObject{v}
	}
	var out []*object.PSObject
	for _, it := range v.ArrayItems() {
		if it.IsArray() {
			out = append(out, flattenArgs(it)...)
		} else {
			out = append(out, it)
		}
	}
	return out
}

// formatArg 按 .NET 格式规格格式化单个参数。
func formatArg(v *object.PSObject, spec string) string {
	if spec == "" {
		return v.String()
	}
	letter := spec[0]
	width := spec[1:]
	switch letter {
	case 'D', 'd':
		if n, ok := v.AsInt(); ok {
			w, _ := strconv.Atoi(width)
			digits := strconv.FormatInt(n, 10)
			neg := strings.HasPrefix(digits, "-")
			if neg {
				digits = digits[1:]
			}
			if w > len(digits) {
				digits = strings.Repeat("0", w-len(digits)) + digits
			}
			if neg {
				digits = "-" + digits
			}
			return digits
		}
	case 'X', 'x':
		if n, ok := v.AsInt(); ok {
			s := strconv.FormatInt(n, 16)
			if letter == 'X' {
				s = strings.ToUpper(s)
			}
			w, _ := strconv.Atoi(width)
			if w > len(s) {
				s = strings.Repeat("0", w-len(s)) + s
			}
			return s
		}
	case 'F', 'f':
		if fv, ok := v.AsFloat(); ok {
			dec := 2
			if d, derr := strconv.Atoi(width); derr == nil {
				dec = d
			}
			return strconv.FormatFloat(fv, 'f', dec, 64)
		}
	case 'N', 'n':
		if fv, ok := v.AsFloat(); ok {
			dec := 2
			if d, derr := strconv.Atoi(width); derr == nil {
				dec = d
			}
			return addThousands(strconv.FormatFloat(fv, 'f', dec, 64))
		}
	}
	// 未知规格：退化为普通字符串
	return v.String()
}

// addThousands 给数字串的整数部分加千分位逗号（.NET N 规格）。
func addThousands(s string) string {
	dot := strings.IndexByte(s, '.')
	intPart, fracPart := s, ""
	if dot >= 0 {
		intPart, fracPart = s[:dot], s[dot:]
	}
	sign := ""
	if strings.HasPrefix(intPart, "-") {
		sign, intPart = "-", intPart[1:]
	}
	var b strings.Builder
	b.WriteString(sign)
	n := len(intPart)
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteByte(intPart[i])
	}
	return b.String() + fracPart
}

func joinOp(l, r *object.PSObject) *object.PSObject {
	sep := r.String()
	var parts []string
	if l.IsArray() {
		for _, it := range l.ArrayItems() {
			parts = append(parts, it.String())
		}
	} else {
		parts = []string{l.String()}
	}
	return object.Str(strings.Join(parts, sep))
}

func splitOp(l, r *object.PSObject) *object.PSObject {
	// 右操作数可为 分隔符 或 分隔符, 最大子串数（"a,b,c" -split ",",2 → a、b,c）
	delim := r.String()
	maxN := -1
	if r.IsArray() {
		items := r.ArrayItems()
		if len(items) >= 1 {
			delim = items[0].String()
		}
		if len(items) >= 2 {
			if n, ok := items[1].AsInt(); ok {
				maxN = int(n)
			}
		}
	}
	re, err := regexp.Compile(delim)
	if err != nil {
		re = regexp.MustCompile(regexp.QuoteMeta(delim))
	}
	// n>0 时最多分 n 段、末段保留未分割剩余；n=0 或负数不限
	n := -1
	if maxN > 0 {
		n = maxN
	}
	parts := re.Split(l.String(), n)
	items := make([]*object.PSObject, len(parts))
	for i, p := range parts {
		items[i] = object.Str(p)
	}
	return object.Array(items)
}

func replaceOp(l, r *object.PSObject) *object.PSObject {
	items := r.ArrayItems()
	pattern := ""
	replacement := ""
	if len(items) >= 1 {
		pattern = items[0].String()
	}
	if len(items) >= 2 {
		replacement = items[1].String()
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return l
	}
	return object.Str(re.ReplaceAllString(l.String(), replacement))
}

func typeIs(l, r *object.PSObject) bool {
	name := strings.ToLower(r.String())
	switch name {
	case "int", "int32", "int64", "long":
		return l.TypeName == "Int"
	case "string":
		return l.TypeName == "String"
	case "bool", "boolean":
		return l.TypeName == "Boolean"
	case "double", "float", "single":
		return l.TypeName == "Double"
	case "array", "object[]":
		return l.IsArray()
	case "hashtable":
		return l.TypeName == "Hashtable"
	case "datetime":
		return l.TypeName == "DateTime"
	}
	return false
}

func asOp(l, r *object.PSObject) *object.PSObject {
	switch strings.ToLower(r.String()) {
	case "int", "int32", "int64", "long":
		if n, ok := l.AsInt(); ok {
			return object.Int(n)
		}
		return object.Null()
	case "double", "float", "single":
		if n, ok := l.AsFloat(); ok {
			return object.Float(n)
		}
		return object.Null()
	case "string":
		return object.Str(l.String())
	case "bool":
		return object.Bool(l.Truthy())
	case "datetime":
		if v, ok := parseDatetimeValue(l); ok {
			return v
		}
		return object.Null()
	case "array", "object[]":
		if l.IsArray() {
			return l
		}
		return object.Array(l.ArrayItems())
	}
	return l
}

// convertValue 把值转换为方括号里声明的类型；数组后缀对每个元素分别转换。
// 无法转换时写非终止错误并返回 $null（与除零的既有错误风格一致）。
func (e *Evaluator) convertValue(v *object.PSObject, typeName string) *object.PSObject {
	norm := strings.ToLower(typeName)
	if strings.HasSuffix(norm, "[]") {
		elem := strings.TrimSuffix(norm, "[]")
		items := make([]*object.PSObject, 0, len(v.ArrayItems()))
		for _, it := range v.ArrayItems() {
			items = append(items, e.convertScalar(it, elem))
		}
		return object.Array(items)
	}
	return e.convertScalar(v, norm)
}

// convertScalar 单值转换：target 为归一化小写类型名，未知类型报"无法找到类型"。
func (e *Evaluator) convertScalar(v *object.PSObject, target string) *object.PSObject {
	out, err := convertTarget(v, target)
	if err != nil {
		e.writeError(err)
		return object.Null()
	}
	return out
}

// convertTarget 单值转换尝试；无法转换或目标类型未注册时返回错误说明。
func convertTarget(v *object.PSObject, target string) (*object.PSObject, error) {
	switch target {
	case "int", "int32", "int64", "long":
		if n, ok := v.AsInt(); ok {
			return object.Int(n), nil
		}
	case "string":
		return object.Str(v.String()), nil
	case "double", "float", "single":
		if f, ok := v.AsFloat(); ok {
			return object.Float(f), nil
		}
	case "bool", "boolean":
		return object.Bool(v.Truthy()), nil
	case "hashtable":
		if _, ok := v.Value.([]object.HashEntry); ok {
			return v, nil
		}
	case "datetime":
		if dv, ok := parseDatetimeValue(v); ok {
			return dv, nil
		}
	case "void":
		return object.Null(), nil
	default:
		return nil, errTypeUnknown
	}
	return nil, fmt.Errorf("%s", lang.T(lang.MsgConvertFail, v.String(), target))
}

// errTypeUnknown 表示目标类型未注册，区别于值无法转换成已注册类型。
var errTypeUnknown = errors.New("type unknown")

// parseDatetimeValue 从 DateTime 原值或常见格式的字符串构造时间对象。
func parseDatetimeValue(v *object.PSObject) (*object.PSObject, bool) {
	if v.TypeName == "DateTime" {
		return v, true
	}
	s := v.String()
	for _, layout := range []string{time.RFC3339, "2006-01-02 15:04:05", "2006-01-02", "15:04:05"} {
		if t, err := time.Parse(layout, s); err == nil {
			return object.DateTime(t), true
		}
	}
	return nil, false
}

func bitOp(l, r *object.PSObject, fn func(a, b int64) int64) *object.PSObject {
	li, ok := l.AsInt()
	if !ok {
		return object.Null()
	}
	ri, ok2 := r.AsInt()
	if !ok2 {
		return object.Null()
	}
	return object.Int(fn(li, ri))
}

// wildcardMatch 委托给 object 包的通配符实现。
func wildcardMatch(pattern, s string) bool {
	return object.WildcardMatch(pattern, s)
}

// requireBuiltin 简化：确保包被引用（编译期占位）。
var _ = builtin.Lookup
var _ = external.LookPath

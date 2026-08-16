package eval

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"powershell/internal/ast"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// runStatements 执行一串语句，捕获 break/continue/return/exit/error 信号。
// 信号携带的 panic 前输出（sig.out）合并进返回输出，保证 throw/return 前的输出不丢。
func (e *Evaluator) runStatements(stmts []ast.Node) (out []*object.PSObject, sig *flowSignal) {
	defer func() {
		if r := recover(); r != nil {
			if fs, ok := r.(*flowSignal); ok {
				sig = fs
				out = append(out, fs.out...)
				return
			}
			panic(r)
		}
	}()
	for _, st := range stmts {
		out = append(out, e.execStatement(st)...)
	}
	return out, nil
}

// execStatement 执行一条语句，返回其输出的对象。
func (e *Evaluator) execStatement(n ast.Node) []*object.PSObject {
	switch v := n.(type) {
	case *ast.Chain:
		var out []*object.PSObject
		out = append(out, e.execStatement(v.Left)...)
		success := e.Session.LastSuccess
		if v.Op == "&&" && success {
			out = append(out, e.execStatement(v.Right)...)
		} else if v.Op == "||" && !success {
			out = append(out, e.execStatement(v.Right)...)
		}
		return out
	case *ast.Pipeline:
		return e.evalPipeline(v)
	case *ast.Assign:
		e.execAssign(v)
		return nil
	case *ast.If:
		return e.execIf(v)
	case *ast.ForEach:
		return e.execForEach(v)
	case *ast.While:
		return e.execWhile(v, false)
	case *ast.DoWhile:
		return e.execWhile(&ast.While{Cond: v.Cond, Body: v.Body}, true)
	case *ast.For:
		return e.execFor(v)
	case *ast.Switch:
		return e.execSwitch(v)
	case *ast.FunctionDef:
		e.Session.Functions[v.Name] = &shell.Function{Name: v.Name, Params: v.Params, Body: v.Body, Filter: v.Filter}
		return nil
	case *ast.ParamBlock:
		// 脚本/函数体开头的 param() 已被解析器提取；这里说明出现在别处的 param 不合法
		e.writeError(fmt.Errorf("param 语句只能在函数或脚本中使用"))
		return nil
	case *ast.Break:
		panic(&flowSignal{kind: flowBreak})
	case *ast.Continue:
		panic(&flowSignal{kind: flowContinue})
	case *ast.Return:
		var val *object.PSObject
		if v.Value != nil {
			val = e.evalValue(v.Value)
		}
		panic(&flowSignal{kind: flowReturn, value: val})
	case *ast.Exit:
		code := 0
		if v.Code != nil {
			if c, ok := e.evalValue(v.Code).AsInt(); ok {
				code = int(c)
			}
		}
		panic(&flowSignal{kind: flowExit, code: code})
	case *ast.Try:
		return e.execTry(v)
	case *ast.Throw:
		return e.execThrow(v)
	case *ast.PipelineExpr:
		return e.evalPipeline(v.Pipeline)
	}
	return nil
}

// execThrow 抛出一个终止错误：构造错误记录（Message + Exception）后以 flowError 信号上抛。
// 无表达式时消息为 ScriptHalted（对齐 PowerShell throw 的默认行为）。
func (e *Evaluator) execThrow(v *ast.Throw) []*object.PSObject {
	msg := "ScriptHalted"
	if v.Value != nil {
		val := e.evalValue(v.Value)
		if !val.IsNull() {
			msg = val.String()
		}
	}
	rec := object.Error(msg)
	exc := object.Object("System.RuntimeException", msg)
	exc.AddProp("Message", object.Str(msg))
	rec.AddProp("Exception", exc)
	panic(&flowSignal{kind: flowError, value: rec})
}

// execTry 执行 try/catch/finally：
// body 出错（flowError）→ 找第一个匹配的 catch（无类型全捕，[Exception]/[System.Exception] 基类全捕，
// 其余按异常类型名精确匹配）→ 把错误记录临时绑到 $_（块结束恢复，普通赋值穿透外层）执行 catch 体；
// finally 无论是否出错、是否被捕获都恒执行；catch/finally 自身的信号优先，
// 未捕获的错误在 finally 之后原样上抛（外层 try 可继续捕获）。
func (e *Evaluator) execTry(v *ast.Try) []*object.PSObject {
	var out []*object.PSObject
	bodyOut, sig := e.runStatements(v.Body.Body.Statements)
	out = append(out, bodyOut...)

	if sig != nil && sig.kind == flowError {
		handled := false
		for _, cc := range v.Catches {
			if !catchMatches(cc.TypeName, sig.value) {
				continue
			}
			// catch 块不推独立作用域（对齐 PowerShell）：普通变量赋值穿透，
			// 只有 $_ 是临时绑定，块结束恢复原值
			sc := e.scopes[len(e.scopes)-1]
			oldUS, hadUS := sc["_"]
			sc["_"] = sig.value
			co, csig := e.runStatements(cc.Body.Body.Statements)
			if hadUS {
				sc["_"] = oldUS
			} else {
				delete(sc, "_")
			}
			out = append(out, co...)
			sig = csig // catch 自身的信号接管（break/continue/return/exit/error）
			handled = true
			break
		}
		if !handled {
			// 无匹配 catch：错误保留，finally 之后原样上抛
		}
	}

	// finally 恒执行；其输出并入，其信号覆盖一切（含未捕获的错误）
	if v.Finally != nil {
		fo, fsig := e.runStatements(v.Finally.Body.Statements)
		out = append(out, fo...)
		if fsig != nil {
			sig = fsig
		}
	}

	if sig != nil {
		sig.out = out // 保留 try/catch/finally 已产生的输出（如 return 前的 finally 输出）
		panic(sig)
	}
	return out
}

// catchMatches 判断 catch 的类型过滤是否匹配错误记录。
// 空类型全捕；"Exception"/"System.Exception" 是全部异常的基类，视为全捕；
// 其余按错误异常的类型名做不区分大小写精确匹配。
func catchMatches(typeName string, errObj *object.PSObject) bool {
	if typeName == "" {
		return true
	}
	exc, ok := errObj.PropValue("Exception")
	if !ok {
		return false
	}
	if strings.EqualFold(typeName, exc.TypeName) {
		return true
	}
	if strings.EqualFold(typeName, "Exception") || strings.EqualFold(typeName, "System.Exception") {
		return true
	}
	return false
}

// execAssign 处理赋值（含 $env: 与复合赋值）。
// 右侧若是语句节点（$x = switch ... 等），执行语句并把输出包成单个值。
func (e *Evaluator) execAssign(a *ast.Assign) {
	e.Session.LastSuccess = true
	var val *object.PSObject
	switch a.Value.(type) {
	case *ast.If, *ast.Switch, *ast.ForEach, *ast.While, *ast.DoWhile, *ast.For, *ast.Try:
		val = e.execStatementValue(a.Value)
	default:
		val = e.evalValue(a.Value)
	}
	if strings.HasPrefix(a.Target, "env:") {
		if a.Op != "=" {
			cur := os.Getenv(a.Target[len("env:"):])
			val = e.binaryOp(a.Op[:len(a.Op)-1], object.Str(cur), val)
		}
		os.Setenv(a.Target[len("env:"):], val.String())
		return
	}
	if a.Op != "=" {
		cur := e.lookupVar(a.Target, a.Scope)
		val = e.binaryOp(a.Op[:len(a.Op)-1], cur, val)
	}
	if err := e.setVar(a.Target, a.Scope, val); err != nil {
		e.writeError(err)
	}
}

// execStatementValue 执行语句并把输出包成单个值（$x = switch ... 等场景）。
func (e *Evaluator) execStatementValue(n ast.Node) *object.PSObject {
	out := e.execStatement(n)
	return wrapSingle(out)
}

func (e *Evaluator) execIf(v *ast.If) []*object.PSObject {
	for _, b := range v.Branches {
		if e.evalValue(b.Cond).Truthy() {
			return e.runBlockOutput(b.Body)
		}
	}
	if v.Else != nil {
		return e.runBlockOutput(v.Else)
	}
	return nil
}

// runBlockOutput 执行块并返回输出对象；break/continue/return/exit/error 一律向上传播。
func (e *Evaluator) runBlockOutput(block *ast.Block) []*object.PSObject {
	if block == nil {
		return nil
	}
	out, sig := e.runStatements(block.Body.Statements)
	if sig != nil {
		sig.out = out // 保留块内 panic 前已产生的输出
		panic(sig)
	}
	return out
}

func (e *Evaluator) execForEach(v *ast.ForEach) []*object.PSObject {
	coll := e.evalValue(v.Coll)
	sc := e.scopes[len(e.scopes)-1]
	// 循环变量绑定在现有作用域，结束后恢复（保证体内对其它变量的赋值影响外层）
	oldVar, hadVar := sc[v.Var]
	oldUS, hadUS := sc["_"]
	defer func() {
		if hadVar {
			sc[v.Var] = oldVar
		} else {
			delete(sc, v.Var)
		}
		if hadUS {
			sc["_"] = oldUS
		} else {
			delete(sc, "_")
		}
	}()
	var out []*object.PSObject
	for _, item := range coll.ArrayItems() {
		sc[v.Var] = item
		sc["_"] = item
		o, sig := e.runStatements(v.Body.Body.Statements)
		out = append(out, o...)
		if sig != nil {
			switch sig.kind {
			case flowBreak:
				return out
			case flowContinue:
				continue
			case flowReturn, flowExit, flowError:
				sig.out = out // 保留 panic 前已收集的输出
				panic(sig)
			}
		}
	}
	return out
}

func (e *Evaluator) execWhile(v *ast.While, doFirst bool) []*object.PSObject {
	var out []*object.PSObject
	for {
		if !doFirst {
			if !e.evalValue(v.Cond).Truthy() {
				break
			}
		}
		o, sig := e.runStatements(v.Body.Body.Statements)
		out = append(out, o...)
		if sig != nil {
			switch sig.kind {
			case flowBreak:
				return out
			case flowContinue:
				// 继续判断条件
			case flowReturn, flowExit, flowError:
				sig.out = out // 保留 panic 前已收集的输出
				panic(sig)
			}
		}
		doFirst = false
		if !e.evalValue(v.Cond).Truthy() {
			break
		}
	}
	return out
}

func (e *Evaluator) execFor(v *ast.For) []*object.PSObject {
	var out []*object.PSObject
	if v.Init != nil {
		if a, ok := v.Init.(*ast.Assign); ok {
			e.execAssign(a)
		} else {
			e.evalValue(v.Init)
		}
	}
	for {
		if v.Cond != nil && !e.evalValue(v.Cond).Truthy() {
			break
		}
		o, sig := e.runStatements(v.Body.Body.Statements)
		out = append(out, o...)
		if sig != nil {
			switch sig.kind {
			case flowBreak:
				return out
			case flowContinue:
				// 执行 post 后继续
			case flowReturn, flowExit, flowError:
				sig.out = out // 保留 panic 前已收集的输出
				panic(sig)
			}
		}
		if v.Post != nil {
			e.evalValue(v.Post)
		}
	}
	return out
}

// execSwitch 执行 switch 语句。
// 值为数组时逐元素匹配：每个元素跑全部 case（可命中多个），default 按元素判断；
// break 退出整个 switch，continue 进入下一个元素（标量时二者等效，都退出）。
// 与 foreach 同机制：不推独立作用域（体内普通赋值穿透外层），只临时绑定 $_/PSItem，结束恢复。
func (e *Evaluator) execSwitch(v *ast.Switch) []*object.PSObject {
	val := e.evalValue(v.Value)
	sc := e.scopes[len(e.scopes)-1]
	oldItem, hadItem := sc["PSItem"]
	oldUS, hadUS := sc["_"]
	defer func() {
		if hadItem {
			sc["PSItem"] = oldItem
		} else {
			delete(sc, "PSItem")
		}
		if hadUS {
			sc["_"] = oldUS
		} else {
			delete(sc, "_")
		}
	}()
	sc["PSItem"] = val

	var items []*object.PSObject
	if val.IsArray() {
		items = val.ArrayItems()
	} else {
		items = []*object.PSObject{val}
	}

	var out []*object.PSObject
nextItem:
	for _, item := range items {
		sc["_"] = item
		sc["PSItem"] = item
		matched := false
		for _, c := range v.Cases {
			if c.Cond == nil {
				// default：仅在本元素无匹配时执行
				if !matched {
					o, sig := e.runStatements(c.Body.Body.Statements)
					out = append(out, o...)
					if sig != nil {
						switch sig.kind {
						case flowBreak:
							return out
						case flowContinue:
							continue nextItem
						case flowReturn, flowExit, flowError:
							sig.out = out // 保留 panic 前已收集的输出
							panic(sig)
						}
					}
				}
				continue
			}
			var isMatch bool
			switch cond := c.Cond.(type) {
			case *ast.ScriptBlock:
				isMatch = e.evalBlockValue(cond.Body).Truthy()
			default:
				cv := e.evalValue(c.Cond)
				switch v.Mode {
				case "regex":
					isMatch, _ = regexp.MatchString(cv.String(), item.String())
				case "wildcard":
					isMatch = wildcardMatch(cv.String(), item.String())
				default:
					isMatch = compareEq(cv, item)
				}
			}
			if isMatch {
				matched = true
				o, sig := e.runStatements(c.Body.Body.Statements)
				out = append(out, o...)
				if sig != nil {
					switch sig.kind {
					case flowBreak:
						return out
					case flowContinue:
						continue nextItem
					case flowReturn, flowExit, flowError:
						sig.out = out // 保留 panic 前已收集的输出
						panic(sig)
					}
				}
			}
		}
	}
	return out
}

// evalBlockValue 执行语句块并取其"返回值"（输出包装为单个值）。
func (e *Evaluator) evalBlockValue(body *ast.StatementList) *object.PSObject {
	e.inCapture++
	defer func() { e.inCapture-- }()
	out, sig := e.runStatements(body.Statements)
	if sig != nil {
		switch sig.kind {
		case flowReturn:
			// return 前的输出保留（switch 条件块与函数语义一致）
			return wrapSingle(append(out, unwrapOutput(sig.value)...))
		case flowError:
			panic(sig) // 块内的终止错误向上传播（外层 try 可捕获）
		}
	}
	return wrapSingle(out)
}

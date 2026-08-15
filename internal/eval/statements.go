package eval

import (
	"os"
	"regexp"
	"strings"

	"powershell/internal/ast"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// runStatements 执行一串语句，捕获 break/continue/return/exit 信号。
func (e *Evaluator) runStatements(stmts []ast.Node) (out []*object.PSObject, sig *flowSignal) {
	defer func() {
		if r := recover(); r != nil {
			if fs, ok := r.(*flowSignal); ok {
				sig = fs
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
	case *ast.PipelineExpr:
		return e.evalPipeline(v.Pipeline)
	}
	return nil
}

// execAssign 处理赋值（含 $env: 与复合赋值）。
func (e *Evaluator) execAssign(a *ast.Assign) {
	e.Session.LastSuccess = true
	val := e.evalValue(a.Value)
	if strings.HasPrefix(a.Target, "env:") {
		if a.Op != "=" {
			cur := os.Getenv(a.Target[len("env:"):])
			val = e.binaryOp(a.Op[:len(a.Op)-1], object.Str(cur), val)
		}
		os.Setenv(a.Target[len("env:"):], val.String())
		return
	}
	if a.Op != "=" {
		cur := e.lookupVar(a.Target)
		val = e.binaryOp(a.Op[:len(a.Op)-1], cur, val)
	}
	if err := e.setVar(a.Target, val); err != nil {
		e.writeError(err)
	}
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

// runBlockOutput 执行块并返回输出对象；break/continue/return/exit 一律向上传播。
func (e *Evaluator) runBlockOutput(block *ast.Block) []*object.PSObject {
	if block == nil {
		return nil
	}
	out, sig := e.runStatements(block.Body.Statements)
	if sig != nil {
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
			case flowReturn, flowExit:
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
			case flowReturn, flowExit:
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
			case flowReturn, flowExit:
				panic(sig)
			}
		}
		if v.Post != nil {
			e.evalValue(v.Post)
		}
	}
	return out
}

func (e *Evaluator) execSwitch(v *ast.Switch) []*object.PSObject {
	val := e.evalValue(v.Value)
	e.pushScope()
	e.scopes[len(e.scopes)-1]["_"] = val
	e.scopes[len(e.scopes)-1]["PSItem"] = val
	defer e.popScope()

	var out []*object.PSObject
	matched := false
	for _, c := range v.Cases {
		if c.Cond == nil {
			// default：仅在无匹配时执行
			if !matched {
				o, sig := e.runStatements(c.Body.Body.Statements)
				out = append(out, o...)
				if sig != nil {
					if sig.kind == flowBreak {
						return out
					}
					if sig.kind == flowReturn || sig.kind == flowExit {
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
				isMatch, _ = regexp.MatchString(cv.String(), val.String())
			case "wildcard":
				isMatch = wildcardMatch(cv.String(), val.String())
			default:
				isMatch = compareEq(cv, val)
			}
		}
		if isMatch {
			matched = true
			o, sig := e.runStatements(c.Body.Body.Statements)
			out = append(out, o...)
			if sig != nil {
				if sig.kind == flowBreak {
					return out
				}
				if sig.kind == flowReturn || sig.kind == flowExit {
					panic(sig)
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
	if sig != nil && sig.kind == flowReturn {
		return sig.value
	}
	return wrapSingle(out)
}

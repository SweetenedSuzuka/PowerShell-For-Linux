package parser

import (
	"strings"

	"powershell/internal/ast"
	"powershell/internal/lang"
	"powershell/internal/lexer"
)

// expression.go 实现表达式解析（二元优先级、一元、后缀、主表达式）。
// isValueStart 判断一个 token 是否能开始一个"值"（参数值或位置实参）。
func (p *Parser) isValueStart(t lexer.Token) bool {
	switch t.Type {
	case TkWord, TkNumber, TkString, TkVariable, TkBraceVar, TkDot:
		return true
	case TkPunct:
		return t.Text == "(" || t.Text == "@" || t.Text == "{" || t.Text == "["
	case TkOp:
		// 单独运算符可作值（/ 根路径、-、+、*、%、! 等字面量，与"裸运算符参数"一致）
		switch t.Text {
		case "-", "!", "/", "+", "%", "*":
			return true
		}
	}
	return false
}

// ---- 表达式 ----

// parseExpression 解析表达式（含逗号数组）。
// 非实参模式下，若以"后跟参数/终止符的裸字"开头，视为命令调用（赋值右侧、条件等场景）。
func (p *Parser) parseExpression(argMode bool) ast.Node {
	if p.err != nil {
		return nil
	}
	// 一元逗号：,1 → 数组
	if p.cur().Type == TkPunct && p.cur().Text == "," {
		p.advance()
		item := p.parseExpression(false)
		return &ast.ArrayLit{Items: []ast.Node{item}}
	}
	if !argMode && p.cur().Type == TkWord {
		nt := p.peekAt(1)
		_, dashPrec := p.binaryOpInfo(nt)
		isBinary := (nt.Type == TkDashWord && dashPrec >= 0) ||
			(nt.Type == TkOp && (nt.Text == "+" || nt.Text == "-" || nt.Text == "*" || nt.Text == "/" || nt.Text == "%"))
		if !isBinary {
			// 语句作表达式：$x = switch (...) {...} / $x = if (...) {...}
			if isStatementKeyword(p.cur().Text) {
				return p.parseStatement()
			}
			// 命令调用作为表达式（如 $x = Get-ChildItem、foreach 集合）
			return &ast.PipelineExpr{Pipeline: p.parsePipeline()}
		}
	}
	e := p.parseBinaryExpr(0, argMode)
	if p.err != nil {
		return e
	}
	// 表达式后接管道：如 $x = 1,2,3 | Measure-Object
	if !argMode && p.cur().Type == TkPunct && p.cur().Text == "|" {
		pipe := &ast.Pipeline{Expr: e}
		for p.cur().Type == TkPunct && p.cur().Text == "|" {
			p.advance()
			p.skipNewlines()
			elem := p.parsePipelineElement()
			if c, ok := elem.(*ast.Command); ok {
				pipe.Commands = append(pipe.Commands, c)
			} else {
				p.fail(lang.T(lang.MsgParsePipeRightCmd))
				break
			}
		}
		return &ast.PipelineExpr{Pipeline: pipe}
	}
	return e
}

func (p *Parser) parseBinaryExpr(minPrec int, argMode bool) ast.Node {
	return p.parseBinaryTail(p.parseUnary(argMode), minPrec, argMode)
}

// parseBinaryExprFrom 以已解析的 lhs 继续解析二元表达式。
func (p *Parser) parseBinaryExprFrom(lhs ast.Node, minPrec int, argMode bool) ast.Node {
	return p.parseBinaryTail(lhs, minPrec, argMode)
}

// parseBinaryTail 解析二元运算符的后缀：三元 ?:、格式 -f 参数列表、逗号数组、普通二元运算。
func (p *Parser) parseBinaryTail(lhs ast.Node, minPrec int, argMode bool) ast.Node {
	for {
		if p.err != nil {
			break
		}
		op, prec := p.binaryOpInfo(p.cur())
		if prec < 0 || prec < minPrec {
			break
		}
		switch op {
		case ",":
			// 逗号：构建数组（比比较运算绑定更紧，使 1,2,3 -eq 2 过滤整个数组）
			items := []ast.Node{lhs}
			for p.cur().Type == TkPunct && p.cur().Text == "," {
				p.advance()
				if !argMode {
					p.skipNewlines()
				}
				items = append(items, p.parseBinaryExpr(prec+1, argMode))
			}
			lhs = &ast.ArrayLit{Items: items}
		case "?":
			// 三元运算符 cond ? 真 : 假（右结合）
			p.advance()
			if !argMode {
				p.skipNewlines()
			}
			ifExpr := p.parseBinaryExpr(prec, argMode)
			sep := p.cur()
			if !(sep.Type == TkWord && sep.Text == ":") && sep.Type != TkColon {
				p.fail(lang.T(lang.MsgParseTernaryColon))
				break
			}
			p.advance()
			if !argMode {
				p.skipNewlines()
			}
			elseExpr := p.parseBinaryExpr(prec, argMode)
			lhs = &ast.Ternary{Cond: lhs, If: ifExpr, Else: elseExpr}
		case "-f":
			// 格式运算符：右侧是逗号分隔的参数列表（如 "{0} {1}" -f 1,2）。
			// 参数项含范围但排除算术（-f 比算术绑定更紧："{0}" -f 5 * 2 先格式化再乘）
			p.advance()
			if !argMode {
				p.skipNewlines()
			}
			items := []ast.Node{p.parseBinaryExpr(41, argMode)}
			for p.cur().Type == TkPunct && p.cur().Text == "," {
				p.advance()
				if !argMode {
					p.skipNewlines()
				}
				items = append(items, p.parseBinaryExpr(41, argMode))
			}
			lhs = &ast.Binary{Op: op, L: lhs, R: &ast.ArrayLit{Items: items}}
		default:
			p.advance()
			// 表达式模式下行尾是二元运算符时语句在下一行继续，与 PowerShell 排版一致；
			// 实参模式（命令参数位）的换行仍是参数分隔，不跨行
			if !argMode {
				p.skipNewlines()
			}
			rhs := p.parseBinaryExpr(prec+1, argMode)
			lhs = &ast.Binary{Op: op, L: lhs, R: rhs}
		}
	}
	return lhs
}

// binaryOpInfo 返回当前 token 作为二元运算符的信息；不是运算符则 prec 为 -1。
func (p *Parser) binaryOpInfo(t lexer.Token) (string, int) {
	switch t.Type {
	case TkPunct:
		if t.Text == "," {
			return ",", 28
		}
	case TkOp:
		switch t.Text {
		case "*", "/", "%":
			return t.Text, 40
		case "+", "-":
			return t.Text, 30
		case "??":
			return "??", 12
		}
	case TkWord:
		if t.Text == ".." {
			return "..", 45 // 范围绑定比算术更严格：1..3 -gt 1 先成范围再过滤
		}
		if t.Text == "?" {
			return "?", 5 // 三元运算符（Where-Object 别名的 ? 只在命令位置出现）
		}
	case TkDashWord:
		// 运算符名归一化为小写，大写输入同样识别。
		name := "-" + strings.ToLower(t.Text)
		if !lexer.IsComparisonOp(name) {
			return "", -1
		}
		switch name {
		case "-f":
			return name, 38 // 格式运算符，绑定比算术紧（"x" + "{0}" -f 5 先格式化）
		case "-and", "-or", "-xor":
			return name, 10
		case "-eq", "-ne", "-lt", "-gt", "-le", "-ge",
			"-like", "-notlike", "-match", "-notmatch",
			"-contains", "-notcontains", "-in", "-notin",
			"-ceq", "-cne", "-clt", "-cgt", "-cle", "-cge",
			"-clike", "-cnotlike", "-cmatch", "-cnotmatch",
			"-ccontains", "-cnotcontains",
			"-ieq", "-ine", "-ilt", "-igt", "-ile", "-ige",
			"-ilike", "-inotlike", "-imatch", "-inotmatch",
			"-icontains", "-inotcontains", "-iin", "-inotin":
			return name, 20
		case "-is", "-isnot", "-as":
			return name, 35
		case "-replace", "-join", "-split":
			return name, 15
		case "-shl", "-shr", "-band", "-bor", "-bxor":
			return name, 25
		}
	}
	return "", -1
}

func (p *Parser) parseUnary(argMode bool) ast.Node {
	t := p.cur()
	if t.Type == TkDashWord && strings.EqualFold(t.Text, "not") {
		p.advance()
		return &ast.Unary{Op: "-not", Operand: p.parseUnary(argMode)}
	}
	if t.Type == TkOp && (t.Text == "!" || t.Text == "-") {
		// 单独运算符作为参数（如 echo -）
		if isLoneOpArg(t.Text, p.peekAt(1)) {
			p.advance()
			return &ast.BareWord{Value: t.Text}
		}
		p.advance()
		return &ast.Unary{Op: t.Text, Operand: p.parseUnary(argMode)}
	}
	// & 调用进表达式：赋值右侧、子表达式等场景与语句位同一形态
	if t.Type == TkOp && t.Text == "&" && !isLoneOpArg("&", p.peekAt(1)) {
		return p.parseInvokeCommand()
	}
	return p.parsePostfix(argMode)
}

func (p *Parser) parsePostfix(argMode bool) ast.Node {
	e := p.parsePrimary(argMode)
	for {
		if p.err != nil {
			break
		}
		t := p.cur()
		if t.Type == TkDot {
			// 属性访问；但 ' .' 单独出现（当前目录）不合并
			nt := p.peekAt(1)
			if nt.Type != TkWord {
				break
			}
			p.advance() // .
			p.advance() // 属性名（可能含点，如 $h.a.b 被词法器并为 a.b）
			// 词法器把 '.' 视为裸字字符，因此 a.b.c 会作为单个词进入这里。
			// 按 PowerShell 语义，未加引号的点号只会是链式成员访问，故拆开逐段解析。
			segs := strings.Split(nt.Text, ".")
			for _, seg := range segs[:len(segs)-1] {
				e = &ast.MemberAccess{Base: e, Prop: seg}
			}
			last := segs[len(segs)-1]
			if p.cur().Type == TkPunct && p.cur().Text == "(" {
				// 方法调用（只能是最后一段）
				p.advance() // (
				var args []ast.Node
				for {
					if p.err != nil {
						break
					}
					p.skipNewlinesAndSemicolons()
					if p.cur().Type == TkPunct && p.cur().Text == ")" {
						p.advance()
						break
					}
					if p.cur().Type == TkEOF {
						p.incomplete = true
						break
					}
					args = append(args, p.parseBinaryExpr(29, false))
					if p.cur().Type == TkPunct && p.cur().Text == "," {
						p.advance()
					}
				}
				e = &ast.MethodCall{Base: e, Name: last, Args: args}
				continue
			}
			e = &ast.MemberAccess{Base: e, Prop: last}
			continue
		}
		if t.Type == TkPunct && t.Text == "[" {
			p.advance()
			p.skipNewlines()
			idx := p.parseExpression(false)
			p.skipNewlines()
			p.expectPunct("]")
			e = &ast.Index{Base: e, Index: idx}
			continue
		}
		if t.Type == TkOp && (t.Text == "++" || t.Text == "--") {
			if vr, ok := e.(*ast.VarRef); ok {
				p.advance()
				return &ast.Increment{Var: vr.Name, Scope: vr.Scope, Op: t.Text}
			}
			p.fail(lang.T(lang.MsgParseIncDecVar))
			break
		}
		break
	}
	return e
}

// atCommaAhead 报告 @() 当前元素内是否有顶层逗号：从当前位置按括号深度扫描 token。
// 深度归零处的逗号分隔下一个元素；`)`/`]`/`}`、管道、分号、赋值、换行与文件尾表示没有。
// 词法 token 原子计入，字符串与脚本块内部不影响深度；只读游标，不移动位置。
func (p *Parser) atCommaAhead() bool {
	depth := 0
	for i := p.pos; i < len(p.toks); i++ {
		t := p.toks[i]
		if t.Type == TkEOF || t.Type == TkNewline {
			return false
		}
		if t.Type == TkOp && (t.Text == "=" || t.Text == "&&" || t.Text == "||") {
			return false
		}
		if t.Type != TkPunct {
			continue
		}
		switch t.Text {
		case "(", "[", "{":
			depth++
		case ")", "]", "}":
			if depth == 0 {
				return false
			}
			depth--
		case ",":
			if depth == 0 {
				return true
			}
		case "|", ";":
			if depth == 0 {
				return false
			}
		}
	}
	return false
}

func (p *Parser) parsePrimary(argMode bool) ast.Node {
	t := p.cur()
	switch t.Type {
	case TkEOF:
		p.incomplete = true
		return &ast.BareWord{Value: ""}
	case TkNumber:
		if argMode && p.canMergeBareword() {
			return p.mergeBareword()
		}
		p.advance()
		return &ast.Number{Value: t.Num, IsInt: t.IsInt}
	case TkString:
		p.advance()
		return p.stringFromParts(t.Parts)
	case TkVariable:
		return p.parseVariable(t)
	case TkBraceVar:
		p.advance()
		if strings.HasPrefix(strings.ToLower(t.Text), "env:") {
			return &ast.EnvRef{Name: t.Text[len("env:"):]}
		}
		scope, name := splitScopeName(t.Text)
		return &ast.VarRef{Name: name, Scope: scope}
	case TkWord:
		if argMode && p.canMergeBareword() {
			return p.mergeBareword()
		}
		// 表达式模式下的裸字按字符串处理
		p.advance()
		return &ast.BareWord{Value: t.Text}
	case TkDot:
		p.advance()
		return &ast.BareWord{Value: "."}
	case TkDashWord:
		// 判定机制与 parseBinaryTail 一致：
		// 能作二元运算符的才报缺左操作数，其余（如只有一元用法的 -not）按意外参数处理。
		if _, prec := p.binaryOpInfo(t); prec >= 0 {
			p.fail(lang.T(lang.MsgParseOpMissingLeft, t.Text))
		} else {
			p.fail(lang.T(lang.MsgParseUnexpectedArg, t.Text))
		}
		p.advance()
		return &ast.BareWord{Value: t.Text}
	case TkOp:
		// 单独运算符作为参数字符串（如 Write-Output /、cd -）
		if isLoneOpArg(t.Text, p.peekAt(1)) {
			p.advance()
			return &ast.BareWord{Value: t.Text}
		}
		// 实参模式下 /、+、%、* 是字面量（-Path / -Recurse、Get-ChildItem /、-Filter *2）
		if argMode && (t.Text == "/" || t.Text == "+" || t.Text == "%" || t.Text == "*") {
			p.advance()
			return &ast.BareWord{Value: t.Text}
		}
		if t.Text == "-" || t.Text == "!" {
			p.advance()
			return &ast.Unary{Op: t.Text, Operand: p.parseUnary(argMode)}
		}
		if t.Text == "<" {
			p.fail(lang.T(lang.MsgParseInputRedirect))
		} else if t.Text == ">" || t.Text == ">>" {
			p.fail(lang.T(lang.MsgParseRedirectAfterCmd))
		} else {
			p.fail(lang.T(lang.MsgParseUnexpectedOp, t.Text))
		}
		p.advance()
		return &ast.BareWord{Value: ""}
	case TkPunct:
		switch t.Text {
		case "(":
			p.advance()
			// 括号内部换行自由，与 PowerShell 排版一致
			p.skipNewlines()
			inner := p.parseExpression(false)
			// 括号内可以是管道：expr | cmd | ...
			if p.cur().Type == TkPunct && p.cur().Text == "|" {
				pipe := &ast.Pipeline{Expr: inner}
				for p.cur().Type == TkPunct && p.cur().Text == "|" {
					p.advance()
					p.skipNewlines()
					elem := p.parsePipelineElement()
					if c, ok := elem.(*ast.Command); ok {
						pipe.Commands = append(pipe.Commands, c)
					} else {
						p.fail(lang.T(lang.MsgParsePipeRightCmd))
						break
					}
				}
				inner = &ast.PipelineExpr{Pipeline: pipe}
			}
			p.skipNewlines()
			p.expectPunct(")")
			return &ast.Paren{Inner: inner}
		case "@":
			nt := p.peekAt(1)
			if nt.Type == TkPunct && nt.Text == "(" {
				p.advance() // @
				p.advance() // (
				var items []ast.Node
				for {
					if p.err != nil {
						break
					}
					p.skipNewlinesAndSemicolons()
					if p.cur().Type == TkPunct && p.cur().Text == ")" {
						p.advance()
						break
					}
					if p.cur().Type == TkEOF {
						p.incomplete = true
						break
					}
					// 元素解析优先级低于逗号（28），使 "x","y" | cmd 先成数组再进管道
					var item ast.Node
					switch {
					case p.cur().Type == TkWord && isAtCommandWord(p.cur().Text) && !p.atCommaAhead():
						// @() 内裸字走命令位置（与原版 PowerShell 一致）：无顶层逗号的单命令元素按命令解析。
						item = p.parsePipelineElement()
					case p.atCommaAhead():
						// 含顶层逗号的多元素沿用原解析路径，保持原有行为。
						item = p.parseBinaryExpr(27, false)
					default:
						// 其余沿用表达式路径（含语句关键字与管道，判据同 parseExpression）。
						item = p.parseExpression(false)
					}
					// 元素内可以是管道：@("x" | ForEach-Object { ... })
					if p.cur().Type == TkPunct && p.cur().Text == "|" {
						pipe := &ast.Pipeline{Expr: item}
						for p.cur().Type == TkPunct && p.cur().Text == "|" {
							p.advance()
							p.skipNewlines()
							elem := p.parsePipelineElement()
							if c, ok := elem.(*ast.Command); ok {
								pipe.Commands = append(pipe.Commands, c)
							} else {
								p.fail(lang.T(lang.MsgParsePipeRightCmd))
								break
							}
						}
						item = &ast.PipelineExpr{Pipeline: pipe}
					}
					items = append(items, item)
					if p.cur().Type == TkPunct && p.cur().Text == "," {
						p.advance()
					}
				}
				return &ast.ArrayLit{Items: items, Flatten: true}
			}
			if nt.Type == TkPunct && nt.Text == "{" {
				p.advance() // @
				return p.parseHashtable()
			}
			// @name 展开：不支持，当作普通文本
			p.advance()
			return &ast.BareWord{Value: "@"}
		case "{":
			return p.parseScriptBlockExpr()
		case ",":
			p.advance()
			item := p.parseExpression(false)
			return &ast.ArrayLit{Items: []ast.Node{item}}
		case "[":
			// 类型字面量与强制转换：[int]、[int[]]、[System.Int32]、[int]$x、[pscustomobject]@{...}
			p.advance() // [
			name := p.cur()
			if name.Type != TkWord {
				p.fail(lang.T(lang.MsgParseTypeLiteralName))
				p.advance()
				return &ast.BareWord{Value: ""}
			}
			typeName := name.Text
			p.advance()
			// [T][] 的数组后缀可叠加（如 int[][]），逐对消费紧贴的方括号
			for p.cur().Type == TkPunct && p.cur().Text == "[" && p.peekAt(1).Type == TkPunct && p.peekAt(1).Text == "]" {
				p.advance()
				p.advance()
				typeName += "[]"
			}
			if p.cur().Type == TkEOF {
				p.incomplete = true
				return &ast.BareWord{Value: ""}
			}
			if !(p.cur().Type == TkPunct && p.cur().Text == "]") {
				p.fail(lang.T(lang.MsgParseTypeLiteralRbracket))
				return &ast.BareWord{Value: ""}
			}
			p.advance()
			// [pscustomobject]@{...} 构造自定义对象，两种模式都保留既有行为
			if strings.EqualFold(typeName, "pscustomobject") && p.cur().Type == TkPunct && p.cur().Text == "@" {
				p.advance() // @
				return &ast.TypeCast{TypeName: typeName, Expr: p.parseHashtable()}
			}
			// 实参模式：类型字面量与静态成员都按裸词字符串处理（如 Write-Output [int] 原样输出）
			if argMode {
				return &ast.BareWord{Value: "[" + typeName + "]"}
			}
			// 词法器把 :: 与成员名并成一个裸词（[math]::Sqrt → "math"、"]"、"::Sqrt"），在此拆出静态成员名；带括号为方法调用，否则为静态属性。
			if p.cur().Type == TkWord && strings.HasPrefix(p.cur().Text, "::") {
				return p.finishStaticMember(typeName, strings.TrimPrefix(p.cur().Text, "::"))
			}
			// 后面紧跟操作数时是强制转换，否则是类型字面量本身（求值为类型名，供 -is/-as 消费）
			operandStart := false
			switch p.cur().Type {
			case TkVariable, TkBraceVar, TkNumber, TkString:
				operandStart = true
			case TkPunct:
				operandStart = p.cur().Text == "(" || p.cur().Text == "@" || p.cur().Text == "{"
			}
			if !operandStart {
				return &ast.TypeCast{TypeName: typeName}
			}
			return &ast.TypeCast{TypeName: typeName, Expr: p.parseUnary(argMode)}
		}
	}
	p.fail(lang.T(lang.MsgParseUnexpectedToken, p.describe(t)))
	return &ast.BareWord{Value: ""}
}

// splitScopeName 把 $script:x 这类带作用域修饰符的名字拆成 (作用域, 名字)。
// 只认 script/global/local；env: 由 EnvRef 单独处理，其余含冒号名字按原样返回。
func splitScopeName(text string) (string, string) {
	idx := strings.Index(text, ":")
	if idx <= 0 {
		return "", text
	}
	scope := strings.ToLower(text[:idx])
	switch scope {
	case "script", "global", "local":
		return scope, text[idx+1:]
	}
	return "", text
}

func (p *Parser) parseVariable(t lexer.Token) ast.Node {
	// $() 子表达式
	if t.Text == "" {
		nt := p.peekAt(1)
		if nt.Type == TkPunct && nt.Text == "(" {
			p.advance() // $
			p.advance() // (
			body := p.parseStatementList(')')
			p.expectPunct(")")
			return &ast.SubExpr{Body: body}
		}
		p.advance()
		return &ast.VarRef{Name: ""}
	}
	p.advance()
	switch strings.ToLower(t.Text) {
	case "true":
		return &ast.BoolLit{Value: true}
	case "false":
		return &ast.BoolLit{Value: false}
	case "null":
		return &ast.NullLit{}
	}
	if strings.HasPrefix(strings.ToLower(t.Text), "env:") {
		return &ast.EnvRef{Name: t.Text[len("env:"):]}
	}
	scope, name := splitScopeName(t.Text)
	return &ast.VarRef{Name: name, Scope: scope}
}

func (p *Parser) parseScriptBlockExpr() ast.Node {
	p.advance() // {
	sb := &ast.ScriptBlock{}
	body := p.parseFunctionStatements(nil, sb)
	if p.cur().Type == TkPunct && p.cur().Text == "}" {
		p.advance()
	} else if p.cur().Type == TkEOF {
		p.incomplete = true
	}
	sb.Body = body
	return sb
}

// parseValueExpr 解析"值表达式"：不做裸字合并、不做命令检测（哈希表键等场景用）。
func (p *Parser) parseValueExpr() ast.Node {
	return p.parseBinaryExpr(0, false)
}

func (p *Parser) parseHashtable() ast.Node {
	p.advance() // {
	ht := &ast.HashtableLit{}
	for {
		if p.err != nil {
			break
		}
		p.skipNewlinesAndSemicolons()
		t := p.cur()
		if t.Type == TkPunct && t.Text == "}" {
			p.advance()
			break
		}
		if t.Type == TkEOF {
			p.incomplete = true
			break
		}
		key := p.parseValueExpr() // 键：值表达式（不合并、不作命令）
		sep := p.cur()
		if !(sep.Type == TkOp && (sep.Text == "=" || sep.Text == ":")) {
			p.fail(lang.T(lang.MsgParseHashEntry))
			break
		}
		p.advance()
		val := p.parseExpression(false)
		ht.Pairs = append(ht.Pairs, ast.HashPair{Key: key, Value: val})
		if p.cur().Type == TkPunct && p.cur().Text == ";" {
			p.advance()
		}
	}
	return ht
}

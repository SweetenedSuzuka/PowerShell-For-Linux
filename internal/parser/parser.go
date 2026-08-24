// Package parser 实现 PowerShell 语法分析（递归下降）。
//
// 解析策略要点：
//   - 语句（; 或换行分隔）→ 管道（| 分隔）→ 命令。
//   - 命令参数采用"实参模式"：裸字默认是字符串，且相邻 token 可合并为一个裸字（如 a=b、2+3）。
//     比较运算符（-eq 等）会把参数转为比较表达式，以便 Where-Object Length -gt 100 这类写法可用。
//   - 命名参数（-Name value / -Name:value）与开关（-Force）在 AST 中保留。
//     参数绑定阶段由求值器依据各 cmdlet 的参数定义裁决。
package parser

import (
	"fmt"
	"strings"

	"powershell/internal/ast"
	"powershell/internal/lang"
	"powershell/internal/lexer"
)

// 便捷别名，避免处处写 lexer.TkXXX。
const (
	TkEOF      = lexer.TkEOF
	TkNewline  = lexer.TkNewline
	TkWord     = lexer.TkWord
	TkNumber   = lexer.TkNumber
	TkString   = lexer.TkString
	TkVariable = lexer.TkVariable
	TkBraceVar = lexer.TkBraceVar
	TkDashWord = lexer.TkDashWord
	TkColon    = lexer.TkColon
	TkDot      = lexer.TkDot
	TkOp       = lexer.TkOp
	TkPunct    = lexer.TkPunct
)

// Result 是解析结果。
type Result struct {
	List       *ast.StatementList
	Incomplete bool // 输入不完整（未闭合括号/尾随管道/运算符等），REPL 用于续行
	Error      error
}

// Parse 解析一段 PowerShell 源码。
func Parse(src string) *Result {
	p := &Parser{src: src, toks: lexer.New(src).Tokens()}
	list := p.parseStatementList(0)
	if p.err != nil {
		return &Result{List: list, Incomplete: false, Error: p.err}
	}
	// 尾随反引号（行续行）也算不完整
	trimmed := strings.TrimRight(src, " \t\r")
	incomplete := p.incomplete
	if strings.HasSuffix(trimmed, "`") {
		incomplete = true
	}
	return &Result{List: list, Incomplete: incomplete, Error: nil}
}

// Parser 是递归下降解析器。
type Parser struct {
	src        string
	toks       []lexer.Token
	pos        int
	err        error
	incomplete bool
}

func (p *Parser) cur() lexer.Token { return p.toks[p.pos] }
func (p *Parser) peekAt(n int) lexer.Token {
	if p.pos+n < len(p.toks) {
		return p.toks[p.pos+n]
	}
	return p.toks[len(p.toks)-1]
}

func (p *Parser) advance() lexer.Token {
	t := p.toks[p.pos]
	if p.pos < len(p.toks)-1 {
		p.pos++
	}
	return t
}

// fail 记录第一条解析错误，后续错误不覆盖。消息已由 lang.T 填好参数。
func (p *Parser) fail(msg string) {
	if p.err == nil {
		p.err = fmt.Errorf("%s", msg)
	}
}

// expectPunct 期望一个指定文本的标点 token。
func (p *Parser) expectPunct(text string) lexer.Token {
	t := p.cur()
	if t.Type == TkEOF {
		p.incomplete = true
		return t
	}
	if !(t.Type == TkPunct && t.Text == text) {
		p.fail(lang.T(lang.MsgParseExpectText, text, p.describe(t)))
		return t
	}
	return p.advance()
}

// expectWord 期望一个裸字 token。
func (p *Parser) expectWord(what string) string {
	t := p.cur()
	if t.Type == TkEOF {
		p.incomplete = true
		return ""
	}
	if t.Type != TkWord {
		p.fail(lang.T(lang.MsgParseExpectWhat, what, p.describe(t)))
		return ""
	}
	p.advance()
	return t.Text
}

// describe 生成 token 的人类可读描述（用于错误信息）。
func (p *Parser) describe(t lexer.Token) string {
	switch t.Type {
	case TkEOF:
		return lang.T(lang.MsgTokEOF)
	case TkNewline:
		return lang.T(lang.MsgTokNewline)
	case TkWord:
		return lang.T(lang.MsgTokWord, t.Text)
	case TkNumber:
		return lang.T(lang.MsgTokNumber, t.Text)
	case TkString:
		return lang.T(lang.MsgTokString)
	case TkVariable:
		return fmt.Sprintf("$%s", t.Text)
	case TkBraceVar:
		return fmt.Sprintf("${%s}", t.Text)
	case TkDashWord:
		return fmt.Sprintf("-%s", t.Text)
	case TkOp:
		return lang.T(lang.MsgTokOp, t.Text)
	case TkPunct:
		return lang.T(lang.MsgTokWord, t.Text)
	}
	return "token"
}

func (p *Parser) isStatementEnd(t lexer.Token) bool {
	return t.Type == TkEOF || t.Type == TkNewline ||
		(t.Type == TkPunct && (t.Text == ";" || t.Text == "}"))
}

// ---- 语句 ----

// parseStatementList 解析一串语句，直到 EOF 或闭合字符 term（0 表示顶层）。
func (p *Parser) parseStatementList(term byte) *ast.StatementList {
	list := &ast.StatementList{}
	// 块/脚本开头的 param() 声明块：它是块的头部而非普通语句，先行解析，不受语句终止符检查约束（param($x) 后可直接跟语句）。
	if p.cur().Type == TkWord && strings.EqualFold(p.cur().Text, "param") {
		nt := p.peekAt(1)
		if nt.Type == TkPunct && nt.Text == "(" {
			list.Statements = append(list.Statements, p.parseParamBlock())
		}
	}
	for {
		p.skipNewlinesAndSemicolons()
		t := p.cur()
		if t.Type == TkEOF {
			if term != 0 {
				p.incomplete = true
			}
			break
		}
		if t.Type == TkPunct && len(t.Text) == 1 && t.Text[0] == term {
			break
		}
		stmt := p.parseStatement()
		if p.err != nil {
			break
		}
		if stmt == nil {
			break
		}
		list.Statements = append(list.Statements, stmt)
		// 语句后必须是终止符
		nt := p.cur()
		if nt.Type == TkEOF {
			break
		}
		if nt.Type == TkNewline || (nt.Type == TkPunct && nt.Text == ";") {
			continue
		}
		if nt.Type == TkPunct && len(nt.Text) == 1 && nt.Text[0] == term {
			continue
		}
		p.fail(lang.T(lang.MsgParseUnexpectedAfter, p.describe(nt)))
		break
	}
	return list
}

func (p *Parser) skipNewlinesAndSemicolons() {
	for {
		t := p.cur()
		if t.Type == TkNewline || (t.Type == TkPunct && t.Text == ";") {
			p.advance()
		} else {
			break
		}
	}
}

// isStatementKeyword 报告单词是否为可作表达式的语句关键字（赋值右侧等位置）。
func isStatementKeyword(text string) bool {
	switch strings.ToLower(text) {
	case "if", "switch", "foreach", "while", "do", "for", "try":
		return true
	}
	return false
}

func (p *Parser) parseStatement() ast.Node {
	t := p.cur()
	if t.Type == TkWord {
		switch {
		case strings.EqualFold(t.Text, "if"):
			return p.parseIf()
		case strings.EqualFold(t.Text, "foreach"):
			return p.parseForEach()
		case strings.EqualFold(t.Text, "while"):
			return p.parseWhile()
		case strings.EqualFold(t.Text, "do"):
			return p.parseDoWhile()
		case strings.EqualFold(t.Text, "for"):
			return p.parseFor()
		case strings.EqualFold(t.Text, "switch"):
			return p.parseSwitch()
		case strings.EqualFold(t.Text, "function"):
			return p.parseFunction(false)
		case strings.EqualFold(t.Text, "filter"):
			return p.parseFunction(true)
		case strings.EqualFold(t.Text, "param"):
			return p.parseParamBlock()
		case strings.EqualFold(t.Text, "break"):
			p.advance()
			return &ast.Break{}
		case strings.EqualFold(t.Text, "continue"):
			p.advance()
			return &ast.Continue{}
		case strings.EqualFold(t.Text, "return"):
			p.advance()
			r := &ast.Return{}
			if !p.isStatementEnd(p.cur()) {
				r.Value = p.parseExpression(false)
			}
			return r
		case strings.EqualFold(t.Text, "exit"):
			p.advance()
			e := &ast.Exit{}
			if !p.isStatementEnd(p.cur()) {
				e.Code = p.parseExpression(false)
			}
			return e
		case strings.EqualFold(t.Text, "try"):
			return p.parseTry()
		case strings.EqualFold(t.Text, "throw"):
			return p.parseThrow()
		case strings.EqualFold(t.Text, "catch"):
			p.fail(lang.T(lang.MsgParseCatchAfterTry))
			return nil
		case strings.EqualFold(t.Text, "finally"):
			p.fail(lang.T(lang.MsgParseFinallyAfterTry))
			return nil
		}
	}
	if (t.Type == TkVariable || t.Type == TkBraceVar) && p.isAssignOp(p.peekAt(1)) {
		return p.parseAssign()
	}
	return p.parseStatementPipeline()
}

// parseStatementPipeline 解析管道，并处理 PowerShell 7 的 && / || 链。
func (p *Parser) parseStatementPipeline() ast.Node {
	var stmt ast.Node = p.parsePipeline()
	for {
		t := p.cur()
		if t.Type == TkOp && (t.Text == "&&" || t.Text == "||") {
			op := t.Text
			p.advance()
			right := p.parsePipeline()
			stmt = &ast.Chain{Left: stmt, Right: right, Op: op}
			continue
		}
		break
	}
	return stmt
}

func (p *Parser) isAssignOp(t lexer.Token) bool {
	if t.Type != TkOp {
		return false
	}
	switch t.Text {
	case "=", "+=", "-=", "*=", "/=", "%=":
		return true
	}
	return false
}

func (p *Parser) parseAssign() *ast.Assign {
	t := p.cur()
	target := t.Text
	p.advance()
	opTok := p.cur()
	if opTok.Type != TkOp {
		p.fail(lang.T(lang.MsgParseAssignOp))
		return nil
	}
	p.advance()
	value := p.parseExpression(false)
	scope, name := splitScopeName(target)
	return &ast.Assign{Target: name, Scope: scope, Op: opTok.Text, Value: value}
}

func (p *Parser) parseBlock() *ast.Block {
	if p.err != nil {
		return nil
	}
	// 块语句的大括号允许另起一行书写，与 PowerShell 排版一致；分号不能代替大括号，故只跳过换行。
	p.skipNewlines()
	t := p.cur()
	if t.Type == TkEOF {
		// Body 必须非 nil：parseFunction 与求值端都会解引用 Body.Statements
		p.incomplete = true
		return &ast.Block{Body: &ast.StatementList{}}
	}
	if !(t.Type == TkPunct && t.Text == "{") {
		p.fail(lang.T(lang.MsgParseExpectBrace, p.describe(t)))
		return nil
	}
	p.advance()
	body := p.parseStatementList('}')
	if p.err != nil {
		return &ast.Block{Body: body}
	}
	if p.cur().Type == TkPunct && p.cur().Text == "}" {
		p.advance()
	} else if p.cur().Type == TkEOF {
		p.incomplete = true
	}
	return &ast.Block{Body: body}
}

func (p *Parser) parseIf() ast.Node {
	p.advance() // if
	br := ast.IfBranch{}
	p.expectPunct("(")
	br.Cond = p.parseExpression(false)
	p.expectPunct(")")
	br.Body = p.parseBlock()
	if p.err != nil {
		return &ast.If{Branches: []ast.IfBranch{br}}
	}
	node := &ast.If{Branches: []ast.IfBranch{br}}
	for {
		t := p.peekPastNewlines()
		if t.Type == TkWord && strings.EqualFold(t.Text, "elseif") {
			p.skipNewlines()
			p.advance()
			eb := ast.IfBranch{}
			p.expectPunct("(")
			eb.Cond = p.parseExpression(false)
			p.expectPunct(")")
			eb.Body = p.parseBlock()
			node.Branches = append(node.Branches, eb)
			continue
		}
		if t.Type == TkWord && strings.EqualFold(t.Text, "else") {
			p.skipNewlines()
			p.advance()
			node.Else = p.parseBlock()
		}
		break
	}
	return node
}

// parseTry 解析 try { } catch [类型]? { } finally { } 链。
// catch 可多个，可带可选的 [类型] 过滤；finally 至多一个，均可省略。
func (p *Parser) parseTry() ast.Node {
	p.advance() // try
	node := &ast.Try{}
	node.Body = p.parseBlock()
	if p.err != nil {
		return node
	}
	hasHandler := false
	for {
		t := p.peekPastNewlines()
		if t.Type == TkEOF {
			// catch/finally 可以写在后续行且可以有多个；一个都没有时输入视为未完，
			// 已有处理器则语句完整（EOF 只是不再有更多 catch）
			if !hasHandler {
				p.incomplete = true
			}
			break
		}
		if t.Type == TkWord && strings.EqualFold(t.Text, "catch") {
			p.skipNewlines()
			p.advance()
			cc := ast.CatchClause{}
			// 类型过滤允许另起一行书写
			p.skipNewlines()
			if p.cur().Type == TkPunct && p.cur().Text == "[" {
				p.advance()
				tn := p.cur()
				if tn.Type == TkWord {
					p.advance()
					cc.TypeName = tn.Text
				}
				if p.cur().Type == TkPunct && p.cur().Text == "]" {
					p.advance()
				}
			}
			cc.Body = p.parseBlock()
			if p.err != nil {
				return node
			}
			node.Catches = append(node.Catches, cc)
			hasHandler = true
			continue
		}
		if t.Type == TkWord && strings.EqualFold(t.Text, "finally") {
			p.skipNewlines()
			p.advance()
			hasHandler = true
			node.Finally = p.parseBlock()
			if p.err != nil {
				return node
			}
			break
		}
		break
	}
	return node
}

// parseThrow 解析 throw [表达式] 语句；无表达式时默认消息 ScriptHalted。
func (p *Parser) parseThrow() ast.Node {
	p.advance() // throw
	th := &ast.Throw{}
	if !p.isStatementEnd(p.cur()) {
		th.Value = p.parseExpression(false)
	}
	return th
}

func (p *Parser) parseForEach() ast.Node {
	p.advance() // foreach
	p.expectPunct("(")
	varTok := p.cur()
	if varTok.Type != TkVariable && varTok.Type != TkBraceVar {
		p.fail(lang.T(lang.MsgParseForeachVar))
		return nil
	}
	p.advance()
	// in 关键字
	inTok := p.cur()
	if !(inTok.Type == TkWord && strings.EqualFold(inTok.Text, "in")) {
		p.fail(lang.T(lang.MsgParseForeachIn))
		return nil
	}
	p.advance()
	coll := p.parseExpression(false)
	p.expectPunct(")")
	body := p.parseBlock()
	return &ast.ForEach{Var: varTok.Text, Coll: coll, Body: body}
}

func (p *Parser) parseWhile() ast.Node {
	p.advance() // while
	p.expectPunct("(")
	cond := p.parseExpression(false)
	p.expectPunct(")")
	body := p.parseBlock()
	return &ast.While{Cond: cond, Body: body}
}

func (p *Parser) parseDoWhile() ast.Node {
	p.advance() // do
	body := p.parseBlock()
	// while 允许写在 } 的下一行
	w := p.peekPastNewlines()
	if w.Type == TkEOF {
		// while 可能写在后续行：输入到此为止视为未完，交由流式入口继续等待或拒绝
		p.incomplete = true
		return &ast.DoWhile{Body: body}
	}
	if !(w.Type == TkWord && strings.EqualFold(w.Text, "while")) {
		p.fail(lang.T(lang.MsgParseDoWhile))
		return &ast.DoWhile{Body: body}
	}
	p.skipNewlines()
	p.advance()
	p.expectPunct("(")
	cond := p.parseExpression(false)
	p.expectPunct(")")
	return &ast.DoWhile{Cond: cond, Body: body}
}

func (p *Parser) parseFor() ast.Node {
	p.advance() // for
	p.expectPunct("(")
	node := &ast.For{}
	if p.cur().Type == TkPunct && p.cur().Text == ";" {
		p.advance()
	} else if (p.cur().Type == TkVariable || p.cur().Type == TkBraceVar) && p.isAssignOp(p.peekAt(1)) {
		node.Init = p.parseAssign()
	} else {
		node.Init = p.parseExpression(false)
	}
	p.expectPunct(";")
	if p.cur().Type == TkPunct && p.cur().Text == ";" {
		p.advance()
	} else {
		node.Cond = p.parseExpression(false)
	}
	p.expectPunct(";")
	if p.cur().Type == TkPunct && p.cur().Text == ")" {
		p.advance()
	} else {
		node.Post = p.parseExpression(false)
		p.expectPunct(")")
	}
	node.Body = p.parseBlock()
	return node
}

func (p *Parser) parseSwitch() ast.Node {
	p.advance() // switch
	// 可选开关：-regex / -wildcard / -exact / -casesensitive（记录但 MVP 仅影响匹配方式）
	mode := "exact"
	for {
		t := p.cur()
		if t.Type != TkDashWord {
			break
		}
		switch strings.ToLower(t.Text) {
		case "regex":
			mode = "regex"
		case "wildcard":
			mode = "wildcard"
		case "casesensitive":
			mode = "case"
		}
		p.advance()
	}
	node := &ast.Switch{Mode: mode}
	p.expectPunct("(")
	node.Value = p.parseExpression(false)
	p.expectPunct(")")
	// switch 的大括号允许另起一行书写，与 PowerShell 排版一致
	p.skipNewlines()
	p.expectPunct("{")
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
		c := ast.SwitchCase{}
		if t.Type == TkWord && strings.EqualFold(t.Text, "default") {
			p.advance()
		} else {
			c.Cond = p.parseExpression(false)
		}
		c.Body = p.parseBlock()
		node.Cases = append(node.Cases, c)
	}
	return node
}

// parseParamList 解析参数列表（param(...) 块与 function f(...) 括号形式共用）。
// 进入时当前位置已越过 '('；'$名' 后可选 '= 默认值'，逗号分隔。
func (p *Parser) parseParamList() []ast.FunctionParam {
	var params []ast.FunctionParam
	for {
		p.skipNewlinesAndSemicolons()
		t := p.cur()
		if t.Type == TkPunct && t.Text == ")" {
			p.advance()
			break
		}
		if t.Type == TkEOF {
			p.incomplete = true
			break
		}
		if t.Type != TkVariable && t.Type != TkBraceVar {
			p.fail(lang.T(lang.MsgParseFuncParam))
			break
		}
		param := ast.FunctionParam{Name: t.Text}
		p.advance()
		if p.cur().Type == TkOp && p.cur().Text == "=" {
			p.advance()
			param.Default = p.parseExpression(false)
		}
		params = append(params, param)
		if p.cur().Type == TkPunct && p.cur().Text == "," {
			p.advance()
		}
	}
	return params
}

// parseParamBlock 解析 param(...) 参数声明块。
func (p *Parser) parseParamBlock() ast.Node {
	p.advance() // param
	if p.cur().Type != TkPunct || p.cur().Text != "(" {
		p.fail(lang.T(lang.MsgParseParamList))
		return &ast.ParamBlock{}
	}
	p.advance() // (
	params := p.parseParamList()
	return &ast.ParamBlock{Params: params}
}

func (p *Parser) parseFunction(filter bool) ast.Node {
	p.advance() // function / filter
	name := p.expectWord(lang.T(lang.MsgParseFuncName))
	fn := &ast.FunctionDef{Name: name, Filter: filter}
	if p.cur().Type == TkPunct && p.cur().Text == "(" {
		p.advance()
		fn.Params = p.parseParamList()
	}
	fn.Body = p.parseBlock()
	// param() 块：函数体开头的参数声明，提取进 Params 并从体里剔除
	if fn.Body != nil && len(fn.Body.Body.Statements) > 0 {
		if pb, ok := fn.Body.Body.Statements[0].(*ast.ParamBlock); ok {
			fn.Params = append(fn.Params, pb.Params...)
			fn.Body.Body.Statements = fn.Body.Body.Statements[1:]
		}
	}
	return fn
}

// ---- 管道与命令 ----

func (p *Parser) parsePipeline() *ast.Pipeline {
	pipe := &ast.Pipeline{}
	elem := p.parsePipelineElement()
	if p.err != nil {
		return pipe
	}
	if cmd, ok := elem.(*ast.Command); ok {
		pipe.Commands = append(pipe.Commands, cmd)
	} else {
		pipe.Expr = elem
	}
	for {
		t := p.cur()
		if !(t.Type == TkPunct && t.Text == "|") {
			break
		}
		p.advance()
		p.skipNewlines()
		if p.cur().Type == TkEOF {
			p.incomplete = true
			break
		}
		cmd := p.parsePipelineElement()
		if c, ok := cmd.(*ast.Command); ok {
			pipe.Commands = append(pipe.Commands, c)
		} else {
			p.fail(lang.T(lang.MsgParsePipeRightCmd))
			break
		}
	}
	return pipe
}

// skipNewlines 跳过换行 token。
func (p *Parser) skipNewlines() {
	for p.cur().Type == TkNewline {
		p.advance()
	}
}

// peekPastNewlines 返回从当前位置跳过若干换行后的 token，不消费任何 token。
// 用于判断 } 之后的换行后面是不是 catch/else/elseif/finally 这类延续关键字；
// 只有命中时才真正消费换行，避免吃掉语句分隔符导致外层误判。
func (p *Parser) peekPastNewlines() lexer.Token {
	i := p.pos
	for i < len(p.toks)-1 && p.toks[i].Type == TkNewline {
		i++
	}
	return p.toks[i]
}

// parsePipelineElement 解析管道的一个元素：可能是命令，也可能是纯表达式。
// 裸字开头的元素一律按命令解析。
// 比较/算术运算符由命令参数内部的合并逻辑处理，例如 Where-Object Length -gt 100 中的 -gt 会与前置裸字参数合并成比较表达式。
func (p *Parser) parsePipelineElement() ast.Node {
	t := p.cur()
	if t.Type == TkEOF {
		p.incomplete = true
		return &ast.BareWord{Value: ""}
	}
	if t.Type == TkWord {
		return p.parseCommand()
	}
	return p.parseExpression(false)
}

// rawSpan 返回 [开始 token, 当前前一 token] 的原始源码片段。
func (p *Parser) rawSpan(begin int) string {
	if p.pos == begin || p.pos == 0 {
		return ""
	}
	start := p.toks[begin].Pos
	end := p.toks[p.pos-1].Pos + len(p.toks[p.pos-1].Raw)
	if start >= end || start >= len(p.src) {
		return ""
	}
	if end > len(p.src) {
		end = len(p.src)
	}
	return p.src[start:end]
}

func (p *Parser) parseCommand() *ast.Command {
	cmd := &ast.Command{}
	nameTok := p.cur()
	cmd.Name = nameTok.Text
	cmd.RawParts = append(cmd.RawParts, nameTok.Raw)
	p.advance()

	for {
		t := p.cur()
		if p.err != nil {
			break
		}
		// 终止
		if t.Type == TkEOF || t.Type == TkNewline {
			break
		}
		if t.Type == TkPunct && (t.Text == "|" || t.Text == ";" || t.Text == "}" || t.Text == ")") {
			break
		}
		if t.Type == TkOp && (t.Text == "&&" || t.Text == "||") {
			break // 管道链运算符结束当前命令
		}
		// 重定向
		if t.Type == TkOp && (t.Text == ">" || t.Text == ">>") {
			kind := ast.RedirStdout
			appendMode := t.Text == ">>"
			if appendMode {
				kind = ast.RedirAppend
			}
			p.advance()
			target := p.parseExpression(true)
			cmd.Redirs = append(cmd.Redirs, ast.Redirection{Kind: kind, Target: target, Append: appendMode})
			continue
		}
		// 2> / 2>> 错误重定向
		if t.Type == TkNumber && t.Num == 2 {
			nt := p.peekAt(1)
			if nt.Type == TkOp && (nt.Text == ">" || nt.Text == ">>") && nt.Adjacent {
				kind := ast.RedirStderr
				p.advance() // 2
				p.advance() // >
				target := p.parseExpression(true)
				cmd.Redirs = append(cmd.Redirs, ast.Redirection{Kind: kind, Target: target, Append: nt.Text == ">>"})
				continue
			}
		}
		// 命名参数 / 开关
		if t.Type == TkDashWord {
			// 二元运算符（比较、逻辑、成员测试等）：把最后一个位置实参并入运算表达式，
			// 后续 token 由 parseBinaryTail 消费。判定依据与 parseBinaryTail 一致，都是 binaryOpInfo；
			// 不在表里的词（如 -not，它只有一元用法）按普通命名参数处理。
			if _, prec := p.binaryOpInfo(t); prec >= 0 {
				if len(cmd.Positional) == 0 {
					p.fail(lang.T(lang.MsgParseCmpOp, t.Text))
					break
				}
				lhs := cmd.Positional[len(cmd.Positional)-1]
				cmd.Positional = cmd.Positional[:len(cmd.Positional)-1]
				expr := p.parseBinaryExprFrom(lhs, 0, true)
				cmd.Positional = append(cmd.Positional, expr)
				continue
			}
			name, val, hasVal := strings.Cut(t.Text, ":")
			if hasVal {
				p.advance() // 吃掉 -Name: 本体（含冒号）
				// -Name:$var / -Name:5 / -Name:"str"：$、引号不是 dash word 字符，词法器会把冒号后的内联值拆成独立 token（紧贴冒号）。
				// 用该 token 的原始文本作为内联值，保证 -Recurse:$true 等语义正确。
				// 仅合并单个 token 能完整表示的值；表达式（括号等）跨多 token，不合并。
				if val == "" && p.cur().Adjacent && p.isValueStart(p.cur()) {
					switch p.cur().Type {
					case TkVariable, TkBraceVar, TkNumber, TkString, TkWord:
						val = p.cur().Raw
						p.advance()
					}
				}
				value := parseInlineExpr(val)
				cmd.Named = append(cmd.Named, ast.NamedArg{Name: name, Value: value})
				cmd.ArgOrder = append(cmd.ArgOrder, ast.ArgItem{Kind: ast.ArgNamed, Name: name, Index: len(cmd.Named) - 1, Inline: true})
				cmd.RawParts = append(cmd.RawParts, "-"+t.Text)
				continue
			}
			p.advance()
			// 支持 -Name=Value（等号赋值）
			if p.cur().Type == TkOp && p.cur().Text == "=" {
				p.advance()
				begin := p.pos
				value := p.parseExpression(true)
				raw := p.rawSpan(begin)
				cmd.Named = append(cmd.Named, ast.NamedArg{Name: name, Value: value})
				cmd.ArgOrder = append(cmd.ArgOrder, ast.ArgItem{Kind: ast.ArgNamed, Name: name, Index: len(cmd.Named) - 1, Inline: true})
				cmd.RawParts = append(cmd.RawParts, "-"+name+"=")
				if raw != "" {
					cmd.RawParts = append(cmd.RawParts, raw)
				}
				continue
			}
			// 开关或命名参数：看下一个 token
			if p.isValueStart(p.cur()) {
				begin := p.pos
				value := p.parseExpression(true)
				raw := p.rawSpan(begin)
				cmd.Named = append(cmd.Named, ast.NamedArg{Name: name, Value: value})
				cmd.ArgOrder = append(cmd.ArgOrder, ast.ArgItem{Kind: ast.ArgNamed, Name: name, Index: len(cmd.Named) - 1})
				cmd.RawParts = append(cmd.RawParts, "-"+name)
				if raw != "" {
					cmd.RawParts = append(cmd.RawParts, raw)
				}
			} else {
				cmd.Switches = append(cmd.Switches, name)
				cmd.ArgOrder = append(cmd.ArgOrder, ast.ArgItem{Kind: ast.ArgSwitch, Name: name, Index: len(cmd.Switches) - 1})
				cmd.RawParts = append(cmd.RawParts, "-"+name)
			}
			continue
		}
		// 位置实参
		begin := p.pos
		arg := p.parseExpression(true)
		if p.err != nil {
			break
		}
		cmd.Positional = append(cmd.Positional, arg)
		cmd.ArgOrder = append(cmd.ArgOrder, ast.ArgItem{Kind: ast.ArgPositional, Index: len(cmd.Positional) - 1})
		if raw := p.rawSpan(begin); raw != "" {
			cmd.RawParts = append(cmd.RawParts, raw)
		}
	}
	return cmd
}

// parseInlineExpr 解析 -Name:Value 的内联值。
func parseInlineExpr(text string) ast.Node {
	p := &Parser{src: text, toks: lexer.New(text).Tokens()}
	node := p.parseExpression(true)
	if p.err != nil || node == nil {
		return &ast.BareWord{Value: text}
	}
	return node
}

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
// 非实参模式下，若以"后跟参数/终止符的裸字"开头，视为命令调用（赋值 RHS、条件等场景）。
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
		isBinary := (nt.Type == TkDashWord && lexer.IsComparisonOp("-"+nt.Text)) ||
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
				items = append(items, p.parseBinaryExpr(prec+1, argMode))
			}
			lhs = &ast.ArrayLit{Items: items}
		case "?":
			// 三元运算符 cond ? 真 : 假（右结合）
			p.advance()
			ifExpr := p.parseBinaryExpr(prec, argMode)
			sep := p.cur()
			if !(sep.Type == TkWord && sep.Text == ":") && sep.Type != TkColon {
				p.fail(lang.T(lang.MsgParseTernaryColon))
				break
			}
			p.advance()
			elseExpr := p.parseBinaryExpr(prec, argMode)
			lhs = &ast.Ternary{Cond: lhs, If: ifExpr, Else: elseExpr}
		case "-f":
			// 格式运算符：RHS 是逗号分隔的参数列表（如 "{0} {1}" -f 1,2）。
			// 参数项含范围但排除算术（-f 比算术绑定更紧："{0}" -f 5 * 2 先格式化再乘）
			p.advance()
			items := []ast.Node{p.parseBinaryExpr(41, argMode)}
			for p.cur().Type == TkPunct && p.cur().Text == "," {
				p.advance()
				items = append(items, p.parseBinaryExpr(41, argMode))
			}
			lhs = &ast.Binary{Op: op, L: lhs, R: &ast.ArrayLit{Items: items}}
		default:
			p.advance()
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
		name := "-" + t.Text
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
			"-ccontains", "-cnotcontains":
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
			idx := p.parseExpression(false)
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
		if strings.HasPrefix(t.Text, "env:") {
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
		if lexer.IsComparisonOp("-" + t.Text) {
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
					item := p.parseBinaryExpr(27, false)
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
			// 类型字面量：目前只支持 [pscustomobject]@{...} 构造自定义对象
			p.advance() // [
			name := p.cur()
			if name.Type != TkWord {
				p.fail(lang.T(lang.MsgParseTypeLiteralName))
				p.advance()
				return &ast.BareWord{Value: ""}
			}
			p.advance()
			if !(p.cur().Type == TkPunct && p.cur().Text == "]") {
				p.fail(lang.T(lang.MsgParseTypeLiteralRbracket))
				return &ast.BareWord{Value: ""}
			}
			p.advance()
			if strings.EqualFold(name.Text, "pscustomobject") {
				if p.cur().Type == TkPunct && p.cur().Text == "@" {
					p.advance() // @
					return &ast.TypeCast{TypeName: name.Text, Expr: p.parseHashtable()}
				}
				p.fail(lang.T(lang.MsgParsePSCustomObject))
				return &ast.BareWord{Value: ""}
			}
			p.fail(lang.T(lang.MsgParseTypeLiteral))
			return &ast.BareWord{Value: ""}
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
	if strings.HasPrefix(t.Text, "env:") {
		return &ast.EnvRef{Name: t.Text[len("env:"):]}
	}
	scope, name := splitScopeName(t.Text)
	return &ast.VarRef{Name: name, Scope: scope}
}

func (p *Parser) parseScriptBlockExpr() ast.Node {
	p.advance() // {
	body := p.parseStatementList('}')
	if p.cur().Type == TkPunct && p.cur().Text == "}" {
		p.advance()
	} else if p.cur().Type == TkEOF {
		p.incomplete = true
	}
	return &ast.ScriptBlock{Body: body}
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

func (p *Parser) stringFromParts(parts []lexer.StringPart) ast.Node {
	if len(parts) == 1 && parts[0].Kind == lexer.PartLit {
		return &ast.StrLit{Value: parts[0].Text}
	}
	var nodes []ast.Node
	for _, part := range parts {
		switch part.Kind {
		case lexer.PartLit:
			nodes = append(nodes, &ast.StrLit{Value: part.Text})
		case lexer.PartVar:
			scope, name := splitScopeName(part.Text)
			nodes = append(nodes, &ast.VarRef{Name: name, Scope: scope})
		case lexer.PartEnvVar:
			nodes = append(nodes, &ast.EnvRef{Name: part.Text})
		case lexer.PartBraceVar:
			if strings.HasPrefix(part.Text, "env:") {
				nodes = append(nodes, &ast.EnvRef{Name: part.Text[len("env:"):]})
			} else {
				scope, name := splitScopeName(part.Text)
				nodes = append(nodes, &ast.VarRef{Name: name, Scope: scope})
			}
		case lexer.PartSubexpr:
			sub := Parse(part.Text)
			nodes = append(nodes, &ast.SubExpr{Body: sub.List})
		}
	}
	return &ast.StrTemplate{Parts: nodes}
}

// canMergeBareword 判断当前位置能否把相邻 token 合并进裸字。
func (p *Parser) canMergeBareword() bool {
	// 检查相邻 token 是否为裸字延续
	nt := p.peekAt(1)
	switch nt.Type {
	case TkWord, TkNumber:
		return nt.Adjacent
	case TkOp:
		return nt.Adjacent && isBarewordOp(nt.Text) && !(nt.Text == "=" && p.isStrongValueStart(p.peekAt(2)))
	case TkDot, TkColon:
		return nt.Adjacent
	}
	return false
}

// isStrongValueStart 判断 token 是否以"强值"开头（引号串/变量/括号）。
// 此时前面的 '=' 是赋值分隔而非裸字延续。
func (p *Parser) isStrongValueStart(t lexer.Token) bool {
	switch t.Type {
	case TkString, TkVariable, TkBraceVar:
		return true
	case TkPunct:
		return t.Text == "(" || t.Text == "@"
	}
	return false
}

// mergeBareword 把相邻的裸字延续 token 合并为一个字符串。
func (p *Parser) mergeBareword() ast.Node {
	var sb strings.Builder
	sb.WriteString(p.cur().Raw)
	p.advance()
	for {
		nt := p.cur()
		if !nt.Adjacent {
			break
		}
		switch nt.Type {
		case TkWord, TkNumber:
			sb.WriteString(nt.Raw)
			p.advance()
		case TkOp:
			if !isBarewordOp(nt.Text) {
				return &ast.BareWord{Value: sb.String()}
			}
			if nt.Text == "=" && p.isStrongValueStart(p.peekAt(1)) {
				return &ast.BareWord{Value: sb.String()}
			}
			sb.WriteString(nt.Raw)
			p.advance()
		case TkDot, TkColon:
			sb.WriteString(nt.Raw)
			p.advance()
		default:
			return &ast.BareWord{Value: sb.String()}
		}
	}
	return &ast.BareWord{Value: sb.String()}
}

func isBarewordOp(text string) bool {
	switch text {
	case "+", "-", "*", "/", "%", "=", "!", "<", ">":
		return true
	}
	return false
}

// isLoneOpArg 判断运算符是否作为独立参数出现（后跟语句结束，如 Write-Output /）。
func isLoneOpArg(op string, next lexer.Token) bool {
	switch op {
	case "/", "+", "*", "%", "-":
	default:
		return false
	}
	if next.Type == TkEOF || next.Type == TkNewline {
		return true
	}
	if next.Type == TkPunct {
		switch next.Text {
		case ";", ")", "}", "|":
			return true
		}
	}
	return false
}

package parser

import (
	"strings"
	"powershell/internal/ast"
	"powershell/internal/lang"
	"powershell/internal/lexer"
)

// statement.go 实现语句分发、管道、赋值与块解析。
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
// 链式运算符写在行尾时右操作数从下一行继续；行首出现的 && / || 属语法错误，到此断句。
func (p *Parser) parseStatementPipeline() ast.Node {
	var stmt ast.Node = p.parsePipeline()
	for {
		t := p.cur()
		if t.Type == TkOp && (t.Text == "&&" || t.Text == "||") {
			p.advance()
			p.skipNewlines()
			right := p.parsePipeline()
			stmt = &ast.Chain{Left: stmt, Right: right, Op: t.Text}
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
	// 赋值号在行尾时语句在下一行继续，与 PowerShell 排版一致
	p.skipNewlines()
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


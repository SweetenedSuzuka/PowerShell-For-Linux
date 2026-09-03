package parser

import (
	"strings"

	"powershell/internal/ast"
	"powershell/internal/lang"
)

// controlflow.go 实现控制流语句解析（if/try/throw/foreach/while/do/for/switch）。
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
			// catch/finally 可以写在后续行且可以有多个。
			// 一个都没有时当作输入未完成处理，已有处理器则语句完整，EOF 只是不再有更多 catch。
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
		// try 语句必须有 catch 或 finally：后面跟的是其它内容而一个处理器都没有时按解析错误处理。
		if !hasHandler {
			p.fail(lang.T(lang.MsgParseTryMissingHandler))
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

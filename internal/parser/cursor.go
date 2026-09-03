package parser

import (
	"fmt"

	"powershell/internal/lang"
	"powershell/internal/lexer"
)

// cursor.go 实现解析游标移动、错误记录与 token 描述。
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


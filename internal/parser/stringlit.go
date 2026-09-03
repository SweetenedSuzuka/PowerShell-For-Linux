package parser

import (
	"strings"

	"powershell/internal/ast"
	"powershell/internal/lexer"
)

// stringlit.go 实现字符串模板与裸字合并。
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
			if strings.HasPrefix(strings.ToLower(part.Text), "env:") {
				nodes = append(nodes, &ast.EnvRef{Name: part.Text[len("env:"):]})
			} else {
				scope, name := splitScopeName(part.Text)
				nodes = append(nodes, &ast.VarRef{Name: name, Scope: scope})
			}
		case lexer.PartSubexpr:
			// 子表达式是独立解析过程，其错误与未完状态必须并入外层。
			// 否则插值会静默丢弃解析失败的语句或执行截断的残缺语句。
			sub := Parse(part.Text)
			if sub.Error != nil {
				p.fail(sub.Error.Error())
			}
			if sub.Incomplete {
				p.incomplete = true
			}
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

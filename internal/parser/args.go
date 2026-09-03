package parser

import (
	"strings"

	"powershell/internal/ast"
	"powershell/internal/lang"
	"powershell/internal/lexer"
)

// args.go 实现调用实参与静态成员解析工具。
// parseCallArgs 解析括号内的逗号分隔实参列表；进入时当前 token 是 "("。
// 零实参也返回非 nil 空切片，供调用方区分"无括号的属性形态"。
func (p *Parser) parseCallArgs() []ast.Node {
	p.advance()
	args := []ast.Node{}
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
	return args
}

// finishStaticMember 完成 [类型]::成员 的静态成员节点；进入时当前 token 是成员名。
// 成员名为空或以冒号开头按解析错误处理。
// 成员名可带点分尾链（如 ::Now.Year，词法器把点并入裸字），拆成逐段成员访问；尾链末段紧跟括号时按方法调用解析，实参挂到末段。
func (p *Parser) finishStaticMember(typeName, member string) ast.Node {
	p.advance() // 消费成员名
	if member == "" || strings.HasPrefix(member, ":") {
		p.fail(lang.T(lang.MsgParseTypeLiteralName))
		return &ast.BareWord{Value: ""}
	}
	segs := strings.Split(member, ".")
	sm := &ast.StaticMember{TypeName: typeName, Name: segs[0]}
	var node ast.Node = sm
	isMethod := false
	var methodArgs []ast.Node
	if p.cur().Type == TkPunct && p.cur().Text == "(" {
		isMethod = true
		methodArgs = p.parseCallArgs()
	}
	if len(segs) > 1 {
		for _, seg := range segs[1 : len(segs)-1] {
			node = &ast.MemberAccess{Base: node, Prop: seg}
		}
		last := segs[len(segs)-1]
		if isMethod {
			node = &ast.MethodCall{Base: node, Name: last, Args: methodArgs}
		} else {
			node = &ast.MemberAccess{Base: node, Prop: last}
		}
	} else if isMethod {
		sm.Args = methodArgs
	}
	return node
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


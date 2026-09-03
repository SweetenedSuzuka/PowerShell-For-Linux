package parser

import (
	"strings"

	"powershell/internal/ast"
	"powershell/internal/lang"
	"powershell/internal/lexer"
)

// function.go 实现函数定义与 param 块解析。

// parseParamList 解析参数列表（param(...) 块与 function f(...) 括号形式共用）。
// 进入时当前位置已越过 '('；每个形参为可选 '[类型]' 标注加 '$名'，其后可选 '= 默认值'，逗号分隔。
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
		param := ast.FunctionParam{}
		if t.Type == TkPunct && t.Text == "[" {
			p.advance()
			tn := p.cur()
			if tn.Type == TkEOF {
				p.incomplete = true
				break
			}
			if tn.Type != TkWord {
				p.fail(lang.T(lang.MsgParseParamTypeName))
				break
			}
			param.TypeName = tn.Text
			p.advance()
			// 数组后缀可叠加（如 int[][]），逐对消费紧贴的方括号
			for p.cur().Type == TkPunct && p.cur().Text == "[" && p.peekAt(1).Type == TkPunct && p.peekAt(1).Text == "]" {
				p.advance()
				p.advance()
				param.TypeName += "[]"
			}
			if p.cur().Type == TkEOF {
				p.incomplete = true
				break
			}
			if !(p.cur().Type == TkPunct && p.cur().Text == "]") {
				p.fail(lang.T(lang.MsgParseParamTypeRbracket))
				break
			}
			p.advance()
			t = p.cur()
			if t.Type == TkEOF {
				p.incomplete = true
				break
			}
		}
		if t.Type != TkVariable && t.Type != TkBraceVar {
			p.fail(lang.T(lang.MsgParseFuncParam))
			break
		}
		param.Name = t.Text
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
	fn.Body = p.parseBlockNamed(fn)
	return fn
}

// parseBlockNamed 解析函数体大括号，体头按 begin/process/end 命名块拆进 fn。
// 命名块必须连续出现在体开头，其后出现裸语句或重复块名报错。
func (p *Parser) parseBlockNamed(fn *ast.FunctionDef) *ast.Block {
	if p.err != nil {
		return nil
	}
	p.skipNewlines()
	t := p.cur()
	if t.Type == TkEOF {
		p.incomplete = true
		return &ast.Block{Body: &ast.StatementList{}}
	}
	if !(t.Type == TkPunct && t.Text == "{") {
		p.fail(lang.T(lang.MsgParseExpectBrace, p.describe(t)))
		return nil
	}
	p.advance()
	body := &ast.Block{Body: p.parseFunctionStatements(fn, nil)}
	if p.err != nil {
		return body
	}
	if p.cur().Type == TkPunct && p.cur().Text == "}" {
		p.advance()
	} else if p.cur().Type == TkEOF {
		p.incomplete = true
	}
	return body
}

// parseFunctionStatements 解析函数体/脚本块体的语句列表：开头连续的 begin/process/end 命名块拆进目标，其余语句照常。
// 命名块的形态是裸字 begin/process/end 后紧跟 { ... }；裸字后不是大括号时按普通命令处理。
func (p *Parser) parseFunctionStatements(fn *ast.FunctionDef, sb *ast.ScriptBlock) *ast.StatementList {
	list := &ast.StatementList{}
	if p.cur().Type == TkWord && strings.EqualFold(p.cur().Text, "param") {
		nt := p.peekAt(1)
		if nt.Type == TkPunct && nt.Text == "(" {
			pb := p.parseParamBlock().(*ast.ParamBlock)
			if fn != nil {
				fn.Params = append(fn.Params, pb.Params...)
			} else {
				list.Statements = append(list.Statements, pb)
			}
			p.skipNewlinesAndSemicolons()
		}
	}
	sawNamed := false
	for {
		p.skipNewlinesAndSemicolons()
		t := p.cur()
		if t.Type == TkEOF || (t.Type == TkPunct && t.Text == "}") {
			break
		}
		if name := namedBlockKeyword(t); name != "" {
			nt := p.peekAt(1)
			isBlock := nt.Type == TkPunct && nt.Text == "{"
			if !isBlock {
				// 裸字 begin 等不接大括号：按普通命令语句回主路径
				list.Statements = append(list.Statements, p.parseStatement())
				if p.err != nil {
					break
				}
				continue
			}
			slots := slotsOf(fn, sb)
			var slot **ast.Block
			switch name {
			case "begin":
				slot = slots.begin
			case "process":
				slot = slots.process
			default:
				slot = slots.end
			}
			if *slot != nil {
				p.fail(lang.T(lang.MsgParseNamedBlockDuplicate, name))
				break
			}
			p.advance() // 消费块名，parseBlock 从 { 开始
			*slot = p.parseBlock()
			sawNamed = true
			continue
		}
		if sawNamed {
			p.fail(lang.T(lang.MsgParseNamedBlockPosition))
			break
		}
		stmt := p.parseStatement()
		if p.err != nil || stmt == nil {
			break
		}
		list.Statements = append(list.Statements, stmt)
		nt := p.cur()
		if nt.Type == TkEOF {
			break
		}
		if nt.Type == TkNewline || (nt.Type == TkPunct && nt.Text == ";") {
			continue
		}
		if nt.Type == TkPunct && nt.Text == "}" {
			continue
		}
		p.fail(lang.T(lang.MsgParseUnexpectedAfter, p.describe(nt)))
		break
	}
	return list
}

// blockSlots 汇总函数或脚本块的三个命名块槽位。
type blockSlots struct{ begin, process, end **ast.Block }

func slotsOf(fn *ast.FunctionDef, sb *ast.ScriptBlock) blockSlots {
	if fn != nil {
		return blockSlots{&fn.Begin, &fn.Process, &fn.End}
	}
	return blockSlots{&sb.Begin, &sb.Process, &sb.End}
}

// namedBlockKeyword 返回 token 对应的命名块名（begin/process/end），其余返回空。
func namedBlockKeyword(t lexer.Token) string {
	if t.Type != TkWord {
		return ""
	}
	switch strings.ToLower(t.Text) {
	case "begin":
		return "begin"
	case "process":
		return "process"
	case "end":
		return "end"
	}
	return ""
}


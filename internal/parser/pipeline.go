package parser

import (
	"strings"

	"powershell/internal/ast"
	"powershell/internal/lang"
	"powershell/internal/lexer"
)

// pipeline.go 实现管道骨架与命令参数收集。
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
// 只有命中时才消费这些换行，语句分隔符不受影响。
func (p *Parser) peekPastNewlines() lexer.Token {
	i := p.pos
	for i < len(p.toks)-1 && p.toks[i].Type == TkNewline {
		i++
	}
	return p.toks[i]
}

// parsePipelineElement 解析管道的一个元素：可能是命令，也可能是纯表达式。
// 裸字开头的元素一律按命令解析；& 开头按调用命令解析。
// 比较/算术运算符由命令参数内部的合并逻辑处理，例如 Where-Object Length -gt 100 中的 -gt 会与前置裸字参数合并成比较表达式。
func (p *Parser) parsePipelineElement() ast.Node {
	t := p.cur()
	if t.Type == TkEOF {
		p.incomplete = true
		return &ast.BareWord{Value: ""}
	}
	if t.Type == TkOp && t.Text == "&" {
		return p.parseInvokeCommand()
	}
	if t.Type == TkWord {
		return p.parseCommand()
	}
	return p.parseExpression(false)
}

// parseInvokeCommand 解析 & 调用命令：& 目标 实参...。
// 目标与实参收进 Name 为 "&" 的命令节点，求值期按目标值分发到脚本块或按名字分发到命令。
func (p *Parser) parseInvokeCommand() *ast.Command {
	p.advance() // &
	cmd := &ast.Command{Name: "&"}
	cmd.RawParts = append(cmd.RawParts, "&")
	if p.cur().Type == TkEOF || p.cur().Type == TkNewline {
		// 只有 & 没有目标：语句在此截断
		p.incomplete = true
		return cmd
	}
	p.collectCommandArgs(cmd)
	return cmd
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
	p.collectCommandArgs(cmd)
	return cmd
}

// atStderrRedirect 报告当前位置是否为紧贴的 2> / 2>>（与重定向分支同条件）。
func (p *Parser) atStderrRedirect() bool {
	t := p.cur()
	if t.Type != TkNumber || t.Num != 2 {
		return false
	}
	nt := p.peekAt(1)
	return nt.Type == TkOp && (nt.Text == ">" || nt.Text == ">>") && nt.Adjacent
}

// collectCommandArgs 收集命令的实参、开关与重定向，直到语句终止符。
func (p *Parser) collectCommandArgs(cmd *ast.Command) {
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
			// 二元运算符（比较、逻辑、成员测试等）会把最后一个位置实参并入运算表达式，后续 token 由 parseBinaryTail 消费。
			// 判定依据与 parseBinaryTail 一致，都是 binaryOpInfo；不在表里的词（如只有一元用法的 -not）按普通命名参数处理。
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
				p.advance() // 消费 -Name: 本体（含冒号）
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
			// 例外：紧贴的 2> / 2>> 是错误重定向，不是开关的值（开关本身不取值）。
			if p.isValueStart(p.cur()) && !p.atStderrRedirect() {
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


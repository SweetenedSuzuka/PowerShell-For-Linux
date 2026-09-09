package parser

import (
	"strconv"
	"strings"
	"testing"

	"powershell/internal/ast"
)

// dump 把 AST 渲染成一行文本，便于断言结构。

func dump(n ast.Node) string {
	var sb strings.Builder
	writeNode(&sb, n)
	return sb.String()
}

func writeNode(sb *strings.Builder, n ast.Node) {
	switch v := n.(type) {
	case nil:
		sb.WriteString("nil")
	case *ast.StatementList:
		sb.WriteString("stmt[")
		for i, s := range v.Statements {
			if i > 0 {
				sb.WriteString("; ")
			}
			writeNode(sb, s)
		}
		sb.WriteString("]")
	case *ast.Pipeline:
		var parts []string
		if v.Expr != nil {
			var inner strings.Builder
			writeNode(&inner, v.Expr)
			parts = append(parts, "expr("+inner.String()+")")
		}
		for _, c := range v.Commands {
			var inner strings.Builder
			writeNode(&inner, c)
			parts = append(parts, inner.String())
		}
		for _, r := range v.Redirs {
			var inner strings.Builder
			writeNode(&inner, r.Target)
			parts = append(parts, "[redir>"+inner.String()+"]")
		}
		sb.WriteString(strings.Join(parts, " | "))
	case *ast.Command:
		sb.WriteString("cmd(")
		sb.WriteString(v.Name)
		for _, a := range v.Positional {
			sb.WriteString(" ")
			writeNode(sb, a)
		}
		for _, na := range v.Named {
			sb.WriteString(" -")
			sb.WriteString(na.Name)
			sb.WriteString(":")
			writeNode(sb, na.Value)
		}
		for _, sw := range v.Switches {
			sb.WriteString(" -")
			sb.WriteString(sw)
		}
		for _, r := range v.Redirs {
			sb.WriteString(" [redir>")
			writeNode(sb, r.Target)
			sb.WriteString("]")
		}
		sb.WriteString(")")
	case *ast.Assign:
		sb.WriteString("set(")
		sb.WriteString(v.Target)
		sb.WriteString(" ")
		sb.WriteString(v.Op)
		sb.WriteString(" ")
		writeNode(sb, v.Value)
		sb.WriteString(")")
	case *ast.Number:
		var s string
		if v.IsInt {
			s = strconv.FormatInt(int64(v.Value), 10)
		} else {
			s = strconv.FormatFloat(v.Value, 'g', -1, 64)
		}
		sb.WriteString("num(")
		sb.WriteString(s)
		sb.WriteString(")")
	case *ast.BareWord:
		sb.WriteString("word(")
		sb.WriteString(v.Value)
		sb.WriteString(")")
	case *ast.StrLit:
		sb.WriteString("str(")
		sb.WriteString(v.Value)
		sb.WriteString(")")
	case *ast.StrTemplate:
		sb.WriteString("tmpl[")
		for i, p := range v.Parts {
			if i > 0 {
				sb.WriteString("+")
			}
			writeNode(sb, p)
		}
		sb.WriteString("]")
	case *ast.VarRef:
		sb.WriteString("$")
		sb.WriteString(v.Name)
	case *ast.EnvRef:
		sb.WriteString("$env:")
		sb.WriteString(v.Name)
	case *ast.Binary:
		sb.WriteString("(")
		writeNode(sb, v.L)
		sb.WriteString(" ")
		sb.WriteString(v.Op)
		sb.WriteString(" ")
		writeNode(sb, v.R)
		sb.WriteString(")")
	case *ast.Ternary:
		sb.WriteString("tern(")
		writeNode(sb, v.Cond)
		sb.WriteString(" ? ")
		writeNode(sb, v.If)
		sb.WriteString(" : ")
		writeNode(sb, v.Else)
		sb.WriteString(")")
	case *ast.Unary:
		sb.WriteString("(u:")
		sb.WriteString(v.Op)
		sb.WriteString(" ")
		writeNode(sb, v.Operand)
		sb.WriteString(")")
	case *ast.MemberAccess:
		writeNode(sb, v.Base)
		sb.WriteString(".")
		sb.WriteString(v.Prop)
	case *ast.Index:
		sb.WriteString("idx[")
		writeNode(sb, v.Base)
		sb.WriteString(",")
		writeNode(sb, v.Index)
		sb.WriteString("]")
	case *ast.Paren:
		sb.WriteString("(")
		writeNode(sb, v.Inner)
		sb.WriteString(")")
	case *ast.ArrayLit:
		sb.WriteString("[")
		for i, it := range v.Items {
			if i > 0 {
				sb.WriteString(",")
			}
			writeNode(sb, it)
		}
		sb.WriteString("]")
	case *ast.HashtableLit:
		sb.WriteString("@{")
		for i, p := range v.Pairs {
			if i > 0 {
				sb.WriteString(";")
			}
			writeNode(sb, p.Key)
			sb.WriteString("=")
			writeNode(sb, p.Value)
		}
		sb.WriteString("}")
	case *ast.BoolLit:
		sb.WriteString("bool(")
		if v.Value {
			sb.WriteString("true")
		} else {
			sb.WriteString("false")
		}
		sb.WriteString(")")
	case *ast.NullLit:
		sb.WriteString("null")
	case *ast.If:
		sb.WriteString("if(")
		for i, b := range v.Branches {
			if i > 0 {
				sb.WriteString(" elif(")
			}
			writeNode(sb, b.Cond)
			sb.WriteString("){")
			writeNode(sb, b.Body)
			sb.WriteString("}")
		}
		if v.Else != nil {
			sb.WriteString(" else{")
			writeNode(sb, v.Else)
			sb.WriteString("}")
		}
		sb.WriteString(")")
	case *ast.Block:
		writeNode(sb, v.Body)
	case *ast.ForEach:
		sb.WriteString("foreach($")
		sb.WriteString(v.Var)
		sb.WriteString(" in ")
		writeNode(sb, v.Coll)
		sb.WriteString("){")
		writeNode(sb, v.Body)
		sb.WriteString("}")
	case *ast.FunctionDef:
		sb.WriteString("func(")
		sb.WriteString(v.Name)
		sb.WriteString("){")
		writeNode(sb, v.Body)
		sb.WriteString("}")
	case *ast.ScriptBlock:
		sb.WriteString("sb{")
		writeNode(sb, v.Body)
		sb.WriteString("}")
	case *ast.Return:
		sb.WriteString("return")
		if v.Value != nil {
			sb.WriteString(" ")
			writeNode(sb, v.Value)
		}
	case *ast.Exit:
		sb.WriteString("exit")
		if v.Code != nil {
			sb.WriteString(" ")
			writeNode(sb, v.Code)
		}
	case *ast.PipelineExpr:
		sb.WriteString("pipeexpr(")
		writeNode(sb, v.Pipeline)
		sb.WriteString(")")
	case *ast.PropertyRef:
		sb.WriteString("prop(")
		sb.WriteString(v.Name)
		sb.WriteString(")")
	case *ast.SubExpr:
		sb.WriteString("$(")
		writeNode(sb, v.Body)
		sb.WriteString(")")
	case *ast.StaticMember:
		sb.WriteString("static(")
		sb.WriteString(v.TypeName)
		sb.WriteString("::")
		sb.WriteString(v.Name)
		if v.Args != nil {
			sb.WriteString("(")
			for i, a := range v.Args {
				if i > 0 {
					sb.WriteString(",")
				}
				writeNode(sb, a)
			}
			sb.WriteString(")")
		}
		sb.WriteString(")")
	case *ast.TypeCast:
		if v.Expr == nil {
			sb.WriteString("type(")
			sb.WriteString(v.TypeName)
			sb.WriteString(")")
		} else {
			sb.WriteString("cast(")
			sb.WriteString(v.TypeName)
			sb.WriteString(",")
			writeNode(sb, v.Expr)
			sb.WriteString(")")
		}
	default:
		sb.WriteString("?")
	}
}

func parseOK(t *testing.T, src string) *ast.StatementList {
	t.Helper()
	res := Parse(src)
	if res.Error != nil {
		t.Fatalf("解析失败 %q: %v", src, res.Error)
	}
	return res.List
}

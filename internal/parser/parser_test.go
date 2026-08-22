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

func TestSimplePipeline(t *testing.T) {
	src := "Get-ChildItem -Force | Where-Object Length -gt 100 | Sort-Object Length -Descending"
	d := dump(parseOK(t, src))
	want := "stmt[cmd(Get-ChildItem -Force) | cmd(Where-Object (word(Length) -gt num(100))) | cmd(Sort-Object word(Length) -Descending)]"
	if d != want {
		t.Fatalf("解析 %q\n  得到 %s\n  想要 %s", src, d, want)
	}
}

func TestChainedMemberAccess(t *testing.T) {
	// 词法器把 a.b 并成一个词，解析器须拆成嵌套成员访问
	list := parseOK(t, "$h.a.b")
	pipe, ok := list.Statements[0].(*ast.Pipeline)
	if !ok || pipe.Expr == nil {
		t.Fatalf("$h.a.b 应解析为带表达式的管道")
	}
	outer, ok := pipe.Expr.(*ast.MemberAccess)
	if !ok {
		t.Fatalf("$h.a.b 外层应为成员访问")
	}
	if outer.Prop != "b" {
		t.Fatalf("外层属性应为 b，得到 %q", outer.Prop)
	}
	inner, ok := outer.Base.(*ast.MemberAccess)
	if !ok {
		t.Fatalf("$h.a.b 内层应为成员访问")
	}
	if inner.Prop != "a" {
		t.Fatalf("内层属性应为 a，得到 %q", inner.Prop)
	}
	if vr, ok := inner.Base.(*ast.VarRef); !ok || vr.Name != "h" {
		t.Fatalf("$h.a.b 基座应为变量 $h")
	}

	// 链式访问最后一段是方法调用时，应解析为 MethodCall
	list = parseOK(t, "$x.a.b()")
	pipe = list.Statements[0].(*ast.Pipeline)
	mc, ok := pipe.Expr.(*ast.MethodCall)
	if !ok {
		t.Fatalf("$x.a.b() 应解析为方法调用")
	}
	if mc.Name != "b" {
		t.Fatalf("方法名应为 b，得到 %q", mc.Name)
	}
	if ma, ok := mc.Base.(*ast.MemberAccess); !ok || ma.Prop != "a" {
		t.Fatalf("$x.a.b() 的基座应为 .a 成员访问")
	}
}

func TestAssignment(t *testing.T) {
	d := dump(parseOK(t, "$x = 1 + 2"))
	want := "stmt[set(x = (num(1) + num(2)))]"
	if d != want {
		t.Fatalf("得到 %s，想要 %s", d, want)
	}
	d = dump(parseOK(t, "$env:PATH = \"/usr/bin\""))
	want = "stmt[set(env:PATH = str(/usr/bin))]"
	if d != want {
		t.Fatalf("得到 %s，想要 %s", d, want)
	}
}

func TestStringExpansion(t *testing.T) {
	d := dump(parseOK(t, "Write-Output \"hi $name $(1+2)\""))
	want := "stmt[cmd(Write-Output tmpl[str(hi )+$name+str( )+$(stmt[expr((num(1) + num(2)))])])]"
	if d != want {
		t.Fatalf("得到 %s，想要 %s", d, want)
	}
}

func TestControlFlow(t *testing.T) {
	src := `if ($x -gt 5) {
  "big"
} elseif ($x -eq 5) {
  "five"
} else {
  "small"
}`
	d := dump(parseOK(t, src))
	if !strings.Contains(d, "if(") || !strings.Contains(d, "elif(") || !strings.Contains(d, "else{") {
		t.Fatalf("if 结构不完整: %s", d)
	}
	d = dump(parseOK(t, "foreach ($f in Get-ChildItem) { echo $f.Name }"))
	if !strings.Contains(d, "foreach($f in pipeexpr(cmd(Get-ChildItem)))") {
		t.Fatalf("foreach 解析失败: %s", d)
	}
}

func TestFunctionAndParams(t *testing.T) {
	d := dump(parseOK(t, "function Add($a, $b = 1) { return $a + $b }"))
	if !strings.Contains(d, "func(Add)") || !strings.Contains(d, "return") {
		t.Fatalf("函数解析失败: %s", d)
	}
}

func TestArraysAndHashtables(t *testing.T) {
	// @(...) 元素解析优先级低于逗号：1,2,3 在元素内先成数组（Flatten 摊平后语义与平铺一致）
	d := dump(parseOK(t, "$a = @(1, 2, 3)"))
	if !strings.Contains(d, "set(a = [[num(1),num(2),num(3)]") {
		t.Fatalf("数组解析失败: %s", d)
	}
	d = dump(parseOK(t, "$h = @{ Name = 'x'; Count = 2 }"))
	if !strings.Contains(d, "@{word(Name)=str(x);word(Count)=num(2)}") {
		t.Fatalf("哈希表解析失败: %s", d)
	}
	// 逗号数组
	d = dump(parseOK(t, "$b = 1,2,3"))
	if !strings.Contains(d, "set(b = [num(1),num(2),num(3)]") {
		t.Fatalf("逗号数组解析失败: %s", d)
	}
}

func TestNamedArgsAndSwitches(t *testing.T) {
	d := dump(parseOK(t, "Get-ChildItem -Path foo -Recurse -Filter *.txt"))
	want := "stmt[cmd(Get-ChildItem -Path:word(foo) -Filter:word(*.txt) -Recurse)]"
	if d != want {
		t.Fatalf("得到 %s，想要 %s", d, want)
	}
	// 内联值
	d = dump(parseOK(t, "Select-Object -First:5"))
	if d != "stmt[cmd(Select-Object -First:num(5))]" {
		t.Fatalf("内联值解析失败: %s", d)
	}
}

func TestRedirection(t *testing.T) {
	d := dump(parseOK(t, "Get-Content a > out.txt"))
	if !strings.Contains(d, "[redir>word(out.txt)]") {
		t.Fatalf("重定向解析失败: %s", d)
	}
	d = dump(parseOK(t, "Get-Content a 2> err.txt"))
	if !strings.Contains(d, "[redir>word(err.txt)]") {
		t.Fatalf("2> 重定向解析失败: %s", d)
	}
}

func TestBarewordMerging(t *testing.T) {
	d := dump(parseOK(t, "Write-Output a=b 2+3"))
	if !strings.Contains(d, "word(a=b)") || !strings.Contains(d, "word(2+3)") {
		t.Fatalf("裸字合并失败: %s", d)
	}
}

func TestTernaryAndOperators(t *testing.T) {
	d := dump(parseOK(t, `$x ? "y" : "n"`))
	if !strings.Contains(d, "tern($x ? str(y) : str(n))") {
		t.Fatalf("三元解析失败: %s", d)
	}
	d = dump(parseOK(t, `$null ?? "d"`))
	if !strings.Contains(d, "(null ?? ") {
		t.Fatalf("?? 解析失败: %s", d)
	}
	d = dump(parseOK(t, `"v={0}" -f 1,2`))
	if !strings.Contains(d, "-f [num(1),num(2)]") {
		t.Fatalf("-f 参数列表解析失败: %s", d)
	}
	// 范围绑定比比较运算更紧
	d = dump(parseOK(t, "1..3 -gt 1"))
	if !strings.Contains(d, "((num(1) .. num(3)) -gt num(1))") {
		t.Fatalf("范围优先级解析失败: %s", d)
	}
}

func TestRawParts(t *testing.T) {
	// 外部命令重建用
	list := parseOK(t, `grep "hello world" file.txt`)
	cmd := list.Statements[0].(*ast.Pipeline).Commands[0]
	if len(cmd.RawParts) != 3 {
		t.Fatalf("RawParts = %v", cmd.RawParts)
	}
	if cmd.RawParts[1] != `"hello world"` {
		t.Fatalf("RawParts[1] = %q", cmd.RawParts[1])
	}
}

func TestIncomplete(t *testing.T) {
	for _, src := range []string{
		"if ($x) {",
		"Get-ChildItem |",
		"1 +",
		"$x = (1+2",
		"echo hi `",
	} {
		res := Parse(src)
		if !res.Incomplete {
			t.Fatalf("%q 应判定为不完整", src)
		}
	}
	for _, src := range []string{
		"Get-ChildItem",
		"Get-ChildItem -Force",
		"$x = 1",
		"echo hi | echo world",
	} {
		res := Parse(src)
		if res.Error != nil || res.Incomplete {
			t.Fatalf("%q 应完整可解析: %v", src, res.Error)
		}
	}
}

func TestErrors(t *testing.T) {
	res := Parse("$x = (1 + 2")
	if res.Error == nil && !res.Incomplete {
		t.Fatal("未闭合括号应报错或标记不完整")
	}
}

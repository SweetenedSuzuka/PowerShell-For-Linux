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

// TestSwitchNewlineBraceTerminates 验证 switch 的分支解析在出错后必然终止，不失去响应。
// expectPunct 失败时不消费 token，分支循环若不检查 p.err 会无限追加 Cases。
func TestSwitchNewlineBraceTerminates(t *testing.T) {
	for _, src := range []string{
		"switch (2)\n{ default { \"d\" } }",
		"switch ($x)\n{",
		"switch ($x)\n{ \"a\" { 1 }",
	} {
		Parse(src) // 解析挂起时测试超时失败
	}
	for _, src := range []string{
		"switch (2) { default { \"d\" } }",
	} {
		res := Parse(src)
		if res.Error != nil || res.Incomplete || len(res.List.Statements) != 1 {
			t.Fatalf("%q 应可完整解析，实际 err=%v", src, res.Error)
		}
	}
}

// TestParserLoopsTerminate 验证命令实参、数组元素、方法实参、哈希表条目四个循环必然终止。
// 命令实参环的二元运算符分支与三个集合环都依赖子解析器推进，遇到不消费 token 的失败路径会原地空转。
func TestParserLoopsTerminate(t *testing.T) {
	// -not 只有一元用法，不在 binaryOpInfo 表里，按普通命名参数解析
	res := Parse("echo a -not")
	if res.Error != nil || res.Incomplete {
		t.Fatalf("echo a -not 应可解析（-not 按命名参数），实际 err=%v", res.Error)
	}
	res = Parse("echo a -not b")
	if res.Error != nil || res.Incomplete {
		t.Fatalf("echo a -not b 应可解析，实际 err=%v", res.Error)
	}
	// 三个集合环遇不消费 token 的失败路径应报错退出而非空转
	for _, src := range []string{
		"@( ] )",
		"@(1, }",
		"$s = \"x\"; $s.m(])",
		"@{ a = > }",
	} {
		res := Parse(src) // 解析挂起时测试超时失败
		if res.Error == nil {
			t.Fatalf("%q 应报解析错误", src)
		}
	}
}

// TestBinaryOpsInCommandArgs 验证命令实参位置的二元运算符合并不受 -not 处理调整影响。
func TestBinaryOpsInCommandArgs(t *testing.T) {
	res := Parse("Where-Object Length -gt 100")
	if res.Error != nil {
		t.Fatalf("Where-Object Length -gt 100 应可解析，实际 err=%v", res.Error)
	}
	res = Parse("Get-ChildItem | Where-Object Name -like \"a*\"")
	if res.Error != nil {
		t.Fatalf("管道加 -like 合并应可解析，实际 err=%v", res.Error)
	}
}

// TestBraceNewlineLayout 验证块语句的大括号允许另起一行书写，与 PowerShell 排版一致。
// 分号不能代替大括号；} 后的换行只有在其后是 catch/else/elseif/finally 时才被消费，不影响后续独立语句。
func TestBraceNewlineLayout(t *testing.T) {
	for _, src := range []string{
		"if ($true)\n{ \"x\" }",
		"if ($true)\n{ 1 }\nelseif ($false)\n{ 2 }\nelse\n{ 3 }",
		"foreach ($i in 1..2)\n{ $i }",
		"while ($false)\n{ 1 }",
		"do\n{ 1 }\nwhile ($false)",
		"for ($i = 0; $i -lt 1; $i++)\n{ }",
		"switch (2)\n{ default { \"d\" } }",
		"function f\n{ \"x\" }",
		"filter ff\n{ $_ }",
		"try\n{ throw \"e\" }\ncatch\n{ \"c\" }",
		"try { }\ncatch [System.Exception]\n{ }\nfinally\n{ 1 }",
	} {
		res := Parse(src)
		if res.Error != nil || res.Incomplete {
			t.Fatalf("%q 应可完整解析，实际 err=%v", src, res.Error)
		}
	}
	// if 块后换行再写独立语句应解析为两条语句，换行不被 else 检查误吃
	res := Parse("if ($true)\n{ 1 }\necho done")
	if res.Error != nil || len(res.List.Statements) != 2 {
		t.Fatalf("if 块后的独立语句应解析为两条语句，实际 err=%v 条数=%d", res.Error, len(res.List.Statements))
	}
	// 分号与缺大括号的写法仍拒绝
	for _, src := range []string{
		"if ($true); { \"x\" }",
		"if ($true)\necho x",
	} {
		if r := Parse(src); r.Error == nil {
			t.Fatalf("%q 应报解析错误（缺少块大括号）", src)
		}
	}
}

// TestIncompleteInputsFlagged 验证跨行构造在输入截断时标记不完整而非报错或崩溃。
// REPL 依赖该标记进入续行；一次性入口据此拒绝执行残缺语句。
func TestIncompleteInputsFlagged(t *testing.T) {
	for _, src := range []string{
		"function f\n",
		"filter ff\n",
		"try { \"x\" }\n",
		"do\n{ 1 }\n",
		"try\n{\nthrow \"e\"\n}\ncatch\n[System.Exception]\n",
	} {
		res := Parse(src)
		if res.Error != nil || !res.Incomplete {
			t.Fatalf("%q 应标记不完整，实际 err=%v incomplete=%v", src, res.Error, res.Incomplete)
		}
	}
}

// TestCatchTypeFilterNewline 验证 catch 的 [类型] 过滤允许另起一行书写。
func TestCatchTypeFilterNewline(t *testing.T) {
	res := Parse("try\n{\nthrow \"e\"\n}\ncatch\n[System.Exception]\n{ \"c\" }")
	if res.Error != nil || res.Incomplete || len(res.List.Statements) != 1 {
		t.Fatalf("catch 类型过滤换行应可完整解析，实际 err=%v", res.Error)
	}
}

// TestTryRequiresHandler 验证 try 语句必须有 catch 或 finally：
// 输入在 try 体后截断时标记不完整，后面跟其它内容时按解析错误处理。
func TestTryRequiresHandler(t *testing.T) {
	res := Parse("try { \"x\" }\n")
	if res.Error != nil || !res.Incomplete {
		t.Fatalf("try 体后截断应标记不完整，实际 err=%v incomplete=%v", res.Error, res.Incomplete)
	}
	for _, src := range []string{
		"try { echo X } \necho done",
		"try { echo X }; echo done",
		"if ($true) { try { 1 } }",
	} {
		if r := Parse(src); r.Error == nil {
			t.Fatalf("%q 应报解析错误（try 缺少 catch/finally）", src)
		}
	}
	for _, src := range []string{
		"try { 1 } catch { 2 }",
		"try { 1 } finally { 2 }",
		"try { 1 } catch [System.Exception] { 2 } finally { 3 }\necho done",
	} {
		r := Parse(src)
		if r.Error != nil || r.Incomplete {
			t.Fatalf("%q 应可完整解析，实际 err=%v incomplete=%v", src, r.Error, r.Incomplete)
		}
	}
}

// TestStringSubexprPropagatesState 验证双引号串内 $() 子表达式的解析状态并入外层：
// 子语句解析失败时报错、截断时标记不完整，不再静默丢弃或提前执行。
func TestStringSubexprPropagatesState(t *testing.T) {
	res := Parse(`Write-Output "a$(do x)b"`)
	if res.Error == nil {
		t.Fatal("子表达式内解析失败的语句应使外层报错")
	}
	res = Parse(`"$side$(try { "SIDE" })"`)
	if res.Error != nil || !res.Incomplete {
		t.Fatalf("子表达式截断应标记外层不完整，实际 err=%v incomplete=%v", res.Error, res.Incomplete)
	}
	res = Parse(`$s = "a$(echo hi)b"; $s`)
	if res.Error != nil || res.Incomplete {
		t.Fatalf("合法插值应不受影响，实际 err=%v", res.Error)
	}
}

// TestExpressionLineContinuation 验证行尾是二元运算符/逗号/赋值号时语句在下一行继续，
// 括号与下标内部换行自由，链式运算符可跨行；无运算符的换行仍是语句边界。
func TestExpressionLineContinuation(t *testing.T) {
	for _, src := range []string{
		"1 +\n2",
		"1 +\n\n2",
		"$x =\n5",
		"( 1 +\n2 )",
		"\n(1 + 2)\n",
		"$m = 1,\n2",
		`"{0}" -f\n1`,
		"$c = $true ?\n\"y\" :\n\"n\"",
		"$s = \"ab\"; $s[\n0]",
		"echo a &&\necho b",
		"echo a ||\necho b",
		"Get-ChildItem |\nMeasure-Object",
	} {
		res := Parse(src)
		if res.Error != nil {
			t.Fatalf("%q 应支持跨行续写，实际 err=%v", src, res.Error)
		}
	}
	// 行尾没有运算符时换行仍是语句边界：两条独立语句
	res := Parse("$x = 5\necho done")
	if res.Error != nil || len(res.List.Statements) != 2 {
		t.Fatalf("无运算符结尾的换行应分隔两条语句，实际 err=%v 条数=%d", res.Error, len(res.List.Statements))
	}
	// 截断输入仍标记不完整供 REPL 续行
	if res := Parse("1 +"); !res.Incomplete {
		t.Fatal("行尾悬挂运算符加 EOF 应标记不完整")
	}
}

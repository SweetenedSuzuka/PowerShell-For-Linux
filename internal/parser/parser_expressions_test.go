package parser

import (
	"powershell/internal/ast"
	"strings"
	"testing"
)

func TestChainedMemberAccess(t *testing.T) {
	// 词法分析器把 a.b 合并成一个词，解析器须拆成嵌套成员访问
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

func TestArraysAndHashtables(t *testing.T) {
	// @(...) 元素解析优先级低于逗号：1,2,3 在元素内先成数组（Flatten 展开后语义与平铺一致）
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

// TestTypeLiterals 验证类型字面量与强制转换的解析形态：
// 无操作数为类型字面量，紧跟操作数为强制转换，数组后缀合并入类型名，实参位置保持裸词字符串。
func TestTypeLiterals(t *testing.T) {
	cases := []struct {
		src  string
		want string
	}{
		{"[int]", "stmt[expr(type(int))]"},
		{"[System.Int32]", "stmt[expr(type(System.Int32))]"},
		{"[int[]]", "stmt[expr(type(int[]))]"},
		{"[int]\"42\"", `stmt[expr(cast(int,str(42)))]`},
		{"[int]$x", "stmt[expr(cast(int,$x))]"},
		{"[int](1 + 2)", "stmt[expr(cast(int,((num(1) + num(2)))))]"},
		{"[int[]](1, 2)", "stmt[expr(cast(int[],([num(1),num(2)])))]"},
		{"Write-Output [int]", "stmt[cmd(Write-Output word([int]))]"},
		{"[math]::Sqrt(4)", "stmt[expr(static(math::Sqrt(num(4))))]"},
		{"[datetime]::Now", "stmt[expr(static(datetime::Now))]"},
		{"[guid]::NewGuid()", "stmt[expr(static(guid::NewGuid()))]"},
	}
	for _, tc := range cases {
		d := dump(parseOK(t, tc.src))
		if d != tc.want {
			t.Fatalf("解析 %q 得到 %s 想要 %s", tc.src, d, tc.want)
		}
	}
	// 缺类型名与缺右括号仍报错/标记为不完整
	if res := Parse("[123]"); res.Error == nil {
		t.Fatal("[123] 应报错（类型字面量需要类型名）")
	}
	if res := Parse("[int"); !res.Incomplete {
		t.Fatal("[int 截断应标记不完整")
	}
}

// TestAtArrayCommaRules 验证 @() 逗号规则（与 PowerShell 一致）：首元素命令接纳逗号实参；值元素后裸命令词报告缺表达式错；裸逗号后裸字同样报错。
func TestAtArrayCommaRules(t *testing.T) {
	for _, src := range []string{
		`@(Write-Output "a", "b")`,
		`@(Get-ChildItem -Path /tmp, "x")`,
		`@(Get-Date)`,
		`@(1, 2)`,
		`@(Get-Date; Get-Date)`,
		`@($x; Get-Date)`,
	} {
		if res := Parse(src); res.Error != nil {
			t.Errorf("%q 应可解析，实际 err=%v", src, res.Error)
		}
	}
	for _, src := range []string{
		`@(1, abc)`,
		`@("a", Write-Output "b")`,
		`1, abc`,
	} {
		if res := Parse(src); res.Error == nil {
			t.Errorf("%q 应报告缺表达式错", src)
		}
	}
}

// TestFormatOperatorValue 验证 -f 后首项与逗号后项须是值表达式，裸字报错（与 PowerShell 一致）。
func TestFormatOperatorValue(t *testing.T) {
	for _, src := range []string{
		`"{0}{1}" -f 1, abc`,
		`"{0}{1}" -f 1, true`,
		`"{0}" -f abc`,
	} {
		if res := Parse(src); res.Error == nil {
			t.Errorf("%q 应报错，实际通过", src)
		}
	}
	for _, src := range []string{
		`"{0}{1}" -f 1, 2`,
		`"{0}" -f 5 * 2`,
		`"{0}" -f
1`,
	} {
		if res := Parse(src); res.Error != nil || res.Incomplete {
			t.Errorf("%q 应可解析，实际 err=%v", src, res.Error)
		}
	}
}

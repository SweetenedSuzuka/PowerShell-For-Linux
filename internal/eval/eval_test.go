package eval

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"powershell/internal/object"
	"powershell/internal/parser"
	"powershell/internal/shell"
)

// runEval 解析并执行一段源码，返回输出对象。
func runEval(t *testing.T, src string) []*object.PSObject {
	t.Helper()
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := New(sess, strings.NewReader(""), io.Discard, io.Discard)
	res := parser.Parse(src)
	if res.Error != nil {
		t.Fatalf("解析错误 %q: %v", src, res.Error)
	}
	var out []*object.PSObject
	for _, st := range res.List.Statements {
		out = append(out, ev.EvalStatement(st)...)
	}
	return out
}

func strs(objs []*object.PSObject) []string {
	var out []string
	for _, o := range objs {
		out = append(out, o.String())
	}
	return out
}

func wantStr(t *testing.T, src string, want ...string) {
	t.Helper()
	got := strs(runEval(t, src))
	if len(got) != len(want) {
		t.Fatalf("%q → %v，想要 %v", src, got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("%q → %v，想要 %v", src, got, want)
		}
	}
}

func TestArithmetic(t *testing.T) {
	wantStr(t, "1 + 2", "3")
	wantStr(t, "10 - 4", "6")
	wantStr(t, "6 * 7", "42")
	wantStr(t, "7 / 2", "3.5")
	wantStr(t, "10 % 3", "1")
	wantStr(t, "2 * 4", "8")
	// 范围运算符绑定比比较运算更紧（对齐 PowerShell）：1..3 -gt 1 先成范围再过滤
	wantStr(t, "1..3 -gt 1", "2", "3")
	wantStr(t, "(1..3) -gt 1", "2", "3")
}

func TestStringOps(t *testing.T) {
	wantStr(t, `"a" + "b"`, "ab")
	wantStr(t, `"ab" * 3`, "ababab")
	wantStr(t, `"abc".ToUpper()`, "ABC")
	wantStr(t, `"Hello".Length`, "5")
	wantStr(t, `"a,b,c".Split(",")[1]`, "b")
}

func TestComparisonAndLogic(t *testing.T) {
	wantStr(t, "5 -gt 3", "True")
	wantStr(t, "5 -lt 3", "False")
	wantStr(t, `"abc" -eq "ABC"`, "True")
	wantStr(t, `"apple" -like "a*"`, "True")
	wantStr(t, `"hello" -match "^h"`, "True")
	wantStr(t, "1 -lt 2 -and 3 -lt 4", "True")
	wantStr(t, "1 -gt 2 -or 3 -lt 4", "True")
	wantStr(t, "-not $false", "True")
}

func TestArrayOps(t *testing.T) {
	wantStr(t, "(1..3) -gt 1", "2", "3")
	wantStr(t, "1,2,3 -eq 2", "2")
	wantStr(t, "1,2,3 -contains 2", "True")
	wantStr(t, "2 -in 1,2,3", "True")
	wantStr(t, `"a","b" -join "-"`, "a-b")
	wantStr(t, `"a-b-c" -split "-"`, "a", "b", "c")
	wantStr(t, `"hello world" -replace "world","ps"`, "hello ps")
}

func TestRangeMembership(t *testing.T) {
	// 范围字面量与成员/比较运算符连用（.. 绑定比比较更紧，修复后与 PowerShell 一致）
	wantStr(t, "5 -in 1..10", "True")
	wantStr(t, "1..10 -contains 5", "True")
	wantStr(t, "1..10 -notcontains 99", "True")
	wantStr(t, "5 -notin 1..10", "False")
	wantStr(t, "1..3 -eq 2", "2")
}

func TestNullCoalescing(t *testing.T) {
	wantStr(t, `$null ?? "d"`, "d")
	wantStr(t, `0 ?? "x"`, "0")
	wantStr(t, `"" ?? "x"`, "")
	wantStr(t, `$null ?? 0 ?? "d"`, "0")
}

func TestTernary(t *testing.T) {
	wantStr(t, `$true ? "y" : "n"`, "y")
	wantStr(t, `$false ? "y" : "n"`, "n")
	wantStr(t, `3 -gt 2 ? "big" : "small"`, "big")
	wantStr(t, `$false ? 1 : $true ? 2 : 3`, "2") // 右结合
}

func TestFormatOperator(t *testing.T) {
	wantStr(t, `"v={0}" -f 42`, "v=42")
	wantStr(t, `"{0}:{1}" -f "a","b"`, "a:b")
	wantStr(t, `"{0:D3}" -f 7`, "007")
	wantStr(t, `"{0:X}" -f 255`, "FF")
	wantStr(t, `"{0:F1}" -f 3.14159`, "3.1")
	wantStr(t, `"a {0} b {1}" -f 1..2`, "a 1 b 2") // 范围展平为位置参数
	wantStr(t, `"x" + "{0}" -f 5`, "x5")           // -f 绑定比 + 紧
}

func TestVariablesAndScope(t *testing.T) {
	wantStr(t, "$x = 10; $x * 2", "20")
	wantStr(t, "$x = 5; $x += 3; $x", "8")
	wantStr(t, `$s = "name"; "$s!"`, "name!")
	wantStr(t, "$a = 1,2,3; $a.Count", "3")
	wantStr(t, "$h = @{k=1}; $h.k", "1")
	wantStr(t, "$env:DUMMY_VAR_XYZ", "")
}

func TestControlFlow(t *testing.T) {
	wantStr(t, "if (5 -gt 3) { 'big' } else { 'small' }", "big")
	wantStr(t, "foreach ($i in 1..3) { $i }", "1", "2", "3")
	wantStr(t, "$s = 0; foreach ($i in 1..4) { $s += $i }; $s", "10")
	wantStr(t, "$i = 0; while ($i -lt 3) { $i++; if ($i -eq 2) { continue }; $i }", "1", "3")
	wantStr(t, "switch (2) { 1 { '一' } 2 { '二' } default { '其他' } }", "二")
	wantStr(t, "for ($i=0; $i -lt 2; $i++) { $i }", "0", "1")
}

func TestPipeline(t *testing.T) {
	wantStr(t, "1..5 | Where-Object { $_ % 2 -eq 0 }", "2", "4")
	wantStr(t, "3,1,2 | Sort-Object", "1", "2", "3")
	wantStr(t, "1..5 | Select-Object -First 2", "1", "2")
	wantStr(t, "1..5 | Measure-Object -Sum | ForEach-Object { $_.Sum }", "15")
}

func TestFunctions(t *testing.T) {
	src := `
function Add($a, $b = 1) { return $a + $b }
Add 5
Add 5 10
`
	wantStr(t, src, "6", "15")
}

func TestHashtableAndSubexpr(t *testing.T) {
	wantStr(t, "$h = @{ Name = 'x'; Count = 2 }; $h.Count", "2")
	wantStr(t, "$(1 + 2)", "3")
	wantStr(t, `"total: $(6 * 7)"`, "total: 42")
}

func TestExternalCommand(t *testing.T) {
	// 未命中的命令应输出"未找到"错误到 stderr（这里仅验证不崩溃）
	objs := runEval(t, "thisCommandDoesNotExist123")
	_ = objs
}

func TestFormatting(t *testing.T) {
	// 文件对象走表格
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := New(sess, strings.NewReader(""), io.Discard, io.Discard)
	res := parser.Parse("Get-ChildItem -Name")
	if res.Error != nil {
		t.Fatal(res.Error)
	}
	var out []*object.PSObject
	for _, st := range res.List.Statements {
		out = append(out, ev.EvalStatement(st)...)
	}
	var buf bytes.Buffer
	if err := object.FormatOutput(&buf, out); err != nil {
		t.Fatal(err)
	}
	if buf.Len() == 0 {
		t.Error("Get-ChildItem 输出为空")
	}
}

func TestAddMemberInputObjectForms(t *testing.T) {
	// Add-Member 的 -InputObject 命名/位置/管道三形式都要生效
	wantStr(t, "(Add-Member -InputObject (Get-Date) -Name t -Value v).t", "v")
	wantStr(t, "(Add-Member (Get-Date) -Name t -Value v).t", "v")
	wantStr(t, "(Get-Date | Add-Member -Name t -Value v).t", "v")
}

func TestConvertFromJsonInputObjectNamed(t *testing.T) {
	// 命名 -InputObject 与位置、管道等价
	wantStr(t, "(ConvertFrom-Json -InputObject '{\"a\":1}').a", "1")
	wantStr(t, "(ConvertFrom-Json '{\"a\":1}').a", "1")
	wantStr(t, "'{\"a\":1}' | ConvertFrom-Json | ForEach-Object { $_.a }", "1")
}

func TestSortObjectPositionalMultiProps(t *testing.T) {
	// 位置多属性排序键逐个比较（如 Sort-Object Length,Name）
	wantStr(t, "\"n,v\n2,b\n1,z\n1,a\" | ConvertFrom-Csv | Sort-Object n,v | Select-Object v | ForEach-Object { $_.v }", "a", "z", "b")
	// 命名多属性同样生效
	wantStr(t, "\"n,v\n2,b\n1,z\n1,a\" | ConvertFrom-Csv | Sort-Object -Property n,v | Select-Object v | ForEach-Object { $_.v }", "a", "z", "b")
	// 单属性不受影响
	wantStr(t, "3,1,2 | Sort-Object", "1", "2", "3")
	// -Unique 按排序键去重
	wantStr(t, "1,1,2,2 | Sort-Object -Unique", "1", "2")
}

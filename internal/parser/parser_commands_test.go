package parser

import (
	"powershell/internal/ast"
	"strings"
	"testing"
)

func TestSimplePipeline(t *testing.T) {
	src := "Get-ChildItem -Force | Where-Object Length -gt 100 | Sort-Object Length -Descending"
	d := dump(parseOK(t, src))
	expected := "stmt[cmd(Get-ChildItem -Force) | cmd(Where-Object (word(Length) -gt num(100))) | cmd(Sort-Object word(Length) -Descending)]"
	if d != expected {
		t.Fatalf("解析 %q\n  得到 %s\n  想要 %s", src, d, expected)
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

func TestBarewordMerging(t *testing.T) {
	d := dump(parseOK(t, "Write-Output a=b 2+3"))
	if !strings.Contains(d, "word(a=b)") || !strings.Contains(d, "word(2+3)") {
		t.Fatalf("裸字合并失败: %s", d)
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

// TestBinaryOpsInCommandArgs 验证命令实参位置的二元运算符按 binaryOpInfo 合并，-not 按命名参数处理。
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

// TestInvokeCommand 验证 & 调用命令的解析形态：
// 目标与实参收进 Name 为 & 的命令节点，目标可为变量或字符串字面量。
func TestInvokeCommand(t *testing.T) {
	cases := map[string]string{
		`& { 1 }`:                   "cmd(& sb{stmt[expr(num(1))]})",
		"& $sb":                     "cmd(& $sb)",
		`& 'Get-ChildItem' -Name x`: "cmd(& str(Get-ChildItem) -Name:word(x))",
		`& $f 1 2`:                  "cmd(& $f num(1) num(2))",
		"1,2 | & { $_ }":            "cmd(& sb{stmt[expr($_)]})",
	}
	for src := range cases {
		res := Parse(src)
		if res.Error != nil || res.Incomplete {
			t.Fatalf("%q 应可完整解析，实际 err=%v incomplete=%v", src, res.Error, res.Incomplete)
		}
		d := dump(parseOK(t, src))
		if !strings.Contains(d, cases[src]) {
			t.Fatalf("%q 的形态应含 %q，实际 %s", src, cases[src], d)
		}
	}
	// & 后无目标标记为不完整
	if res := Parse("&"); !res.Incomplete {
		t.Fatal("& 截断应标记不完整")
	}
}

// TestCommandCommaMissingArg 验证命令名后直接跟逗号报告缺少参数；已有实参则按位置数组收集。
func TestCommandCommaMissingArg(t *testing.T) {
	for _, src := range []string{
		`Get-Date, Get-Date`,
		`Get-Date, "x"`,
		`Write-Output ,1`,
		`@(Get-Date, 1)`,
		`@(Get-Date, "x")`,
	} {
		res := Parse(src)
		if res.Error == nil || !strings.Contains(res.Error.Error(), "缺少参数") {
			t.Errorf("%q 应报告缺少参数，实际 err=%v", src, res.Error)
		}
	}
	for _, src := range []string{
		`Write-Output "a", Write-Output "b"`,
		`Get-ChildItem /tmp, "x"`,
		`Write-Output 1,2`,
		`Copy-Item a b, c`,
	} {
		if res := Parse(src); res.Error != nil || res.Incomplete {
			t.Errorf("%q 应可解析，实际 err=%v", src, res.Error)
		}
	}
}

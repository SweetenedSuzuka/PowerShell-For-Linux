package parser

import (
	"testing"
)

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
// 命令实参环的二元运算符分支与三个集合环都依赖子解析器推进，遇到不消费 token 的失败路径会死循环。
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
	// 三个集合环遇不消费 token 的失败路径应报错退出而非死循环
	for _, src := range []string{
		"@( ] )",
		"@(1, }",
		"$s = \"x\"; $s.m(])",
		"@{ a = > }",
	} {
		res := Parse(src) // 解析挂起时测试超时失败
		if res.Error == nil {
			t.Fatalf("%q 应报告解析错误", src)
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

// TestStringSubexprPropagatesState 验证双引号串内 $() 子表达式的解析状态合并入外层：
// 子语句解析失败时报错，截断时标记不完整，插值结果不含残缺语句。
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

// TestExpressionLineContinuation 验证行尾为二元运算符/逗号/赋值号时语句在下一行继续。
// 括号与下标内部换行自由，链式运算符可跨行；无运算符的换行仍是语句边界。
func TestExpressionLineContinuation(t *testing.T) {
	for _, src := range []string{
		"1 +\n2",
		"1 +\n\n2",
		"$x =\n5",
		"( 1 +\n2 )",
		"\n(1 + 2)\n",
		"$m = 1,\n2",
		`"{0}" -f
1`,
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

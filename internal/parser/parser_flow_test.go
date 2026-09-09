package parser

import (
	"powershell/internal/ast"
	"strings"
	"testing"
)

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
	// if 块后换行再写独立语句应解析为两条语句，换行不被 else 检查消耗
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
			t.Fatalf("%q 应报告解析错误（缺少块大括号）", src)
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
			t.Fatalf("%q 应报告解析错误（try 缺少 catch/finally）", src)
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

// TestParamTypeAnnotations 验证形参 [类型] 标注的解析：
// 标注写入 FunctionParam.TypeName 且数组后缀合并入名字，缺类型名或缺 ']' 报错，截断标记为不完整。
func TestParamTypeAnnotations(t *testing.T) {
	cases := map[string]string{
		"function f([int]$x) { $x }":          "int",
		"function f([System.Int32]$x) { $x }": "System.Int32",
		"function f([int[]]$a) { $a }":        "int[]",
		"function f([int[][]]$m) { $m }":      "int[][]",
		"function f([int]$x = 5) { $x }":      "int",
		"function g { param([string]$s) $s }": "string",
		"function g { param([int[]]$a) $a }":  "int[]",
	}
	for src, want := range cases {
		res := Parse(src)
		if res.Error != nil || res.Incomplete {
			t.Fatalf("%q 应完整解析，实际 err=%v incomplete=%v", src, res.Error, res.Incomplete)
		}
		var fd *ast.FunctionDef
		for _, st := range res.List.Statements {
			if f, ok := st.(*ast.FunctionDef); ok {
				fd = f
			}
		}
		if fd == nil || len(fd.Params) == 0 || fd.Params[0].TypeName != want {
			t.Fatalf("%q 首个形参的类型标注应为 %q", src, want)
		}
	}
	// 无标注的形参 TypeName 为空
	res := Parse("function h($a, [int]$b) { $a + $b }")
	if res.Error != nil {
		t.Fatalf("混合标注应可解析：%v", res.Error)
	}
	fd := res.List.Statements[0].(*ast.FunctionDef)
	if fd.Params[0].TypeName != "" || fd.Params[1].TypeName != "int" {
		t.Fatalf("混合标注解析错误：%q 与 %q", fd.Params[0].TypeName, fd.Params[1].TypeName)
	}
	// 标注缺类型名报错
	if r := Parse("function f([123]$x) { $x }"); r.Error == nil {
		t.Fatal("形参标注缺类型名应报错")
	}
	// 缺 ']' 报错
	if r := Parse("function f([int$x) { $x }"); r.Error == nil {
		t.Fatal("形参标注缺 ']' 应报错")
	}
	// 截断标记为不完整：类型名未完、']' 未到、变量未到三处
	for _, src := range []string{"function f([int", "function f([int]", "function g { param([int"} {
		if r := Parse(src); !r.Incomplete {
			t.Fatalf("%q 截断应标记不完整", src)
		}
	}
}

// TestNamedBlocks 验证函数体 begin/process/end 命名块的解析：
// 三块拆进 FunctionDef 对应字段、可乱序、filter 同样接受；裸语句与命名块混写、重复块名报错。
func TestNamedBlocks(t *testing.T) {
	res := Parse("function f { begin { 1 } process { 2 } end { 3 } }")
	if res.Error != nil || res.Incomplete {
		t.Fatalf("三块应完整解析，实际 err=%v incomplete=%v", res.Error, res.Incomplete)
	}
	fn := res.List.Statements[0].(*ast.FunctionDef)
	if fn.Begin == nil || fn.Process == nil || fn.End == nil {
		t.Fatalf("三块都应拆进对应字段：begin=%v process=%v end=%v", fn.Begin != nil, fn.Process != nil, fn.End != nil)
	}
	if len(fn.Body.Body.Statements) != 0 {
		t.Fatalf("命名块应从体里剔除，实际残留 %d 条", len(fn.Body.Body.Statements))
	}
	// 块乱序
	res = Parse("function h { end { 'e' } begin { 'b' } }")
	if res.Error != nil {
		t.Fatalf("块乱序应可解析：%v", res.Error)
	}
	fn = res.List.Statements[0].(*ast.FunctionDef)
	if fn.Begin == nil || fn.End == nil || fn.Process != nil {
		t.Fatalf("乱序两块的拆分结果错误：begin=%v process=%v end=%v", fn.Begin != nil, fn.Process != nil, fn.End != nil)
	}
	// filter 带命名块
	res = Parse("filter g { process { 9 } }")
	if res.Error != nil {
		t.Fatalf("filter 带命名块应可解析：%v", res.Error)
	}
	// 裸语句在命名块之后报错
	if r := Parse("function m { begin { 1 } 'bare' }"); r.Error == nil {
		t.Fatal("命名块后接裸语句应报错")
	}
	// 重复块名报错
	if r := Parse("function d { begin { 1 } begin { 2 } }"); r.Error == nil {
		t.Fatal("重复的 begin 块应报错")
	}
	// 裸字 begin 不接大括号时按普通命令处理（不报错）
	res = Parse("function n { begin -x }")
	if res.Error != nil {
		t.Fatalf("裸字 begin 不接大括号应按普通命令解析：%v", res.Error)
	}
}

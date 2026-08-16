package eval

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"powershell/internal/object"
	"powershell/internal/parser"
	"powershell/internal/shell"
)

// runEval 解析并执行一段源码，返回输出对象。
func runEval(t *testing.T, src string) []*object.PSObject {
	t.Helper()
	return runEvalWithStyle(t, shell.StyleCore, src)
}

// runEvalWithStyle 按指定风格（5.X/7.X）执行源码。
func runEvalWithStyle(t *testing.T, style shell.Style, src string) []*object.PSObject {
	t.Helper()
	sess := shell.New(style, io.Discard, io.Discard, strings.NewReader(""))
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

// TestNumericLiterals 验证数字字面量：紧贴除法（5/2、$x/2）、十六进制、KB 后缀。
func TestNumericLiterals(t *testing.T) {
	// 紧贴除法：/ 后跟数字且前一 token 是值 → 除法运算符
	wantStr(t, "5/2", "2.5")
	wantStr(t, "$x = 10; $x/2", "5")
	wantStr(t, "5 / 2", "2.5")
	wantStr(t, "(5)/2", "2.5")
	// 路径语义不破坏：/ 后跟数字但前面不是值 → 仍是路径参数
	wantStr(t, `Write-Output /2`, "/2")
	wantStr(t, `Write-Output a /2`, "a", "/2")
	wantStr(t, `Write-Output ./x`, "./x")
	// 十六进制字面量
	wantStr(t, "0x10", "16")
	wantStr(t, "0xff", "255")
	wantStr(t, "0x10 + 1", "17")
	// 数字后缀（二进制倍数）
	wantStr(t, "1KB", "1024")
	wantStr(t, "1MB", "1048576")
	wantStr(t, "2KB/2", "1024")
	wantStr(t, "2.5KB", "2560")
	wantStr(t, "1GB", "1073741824")
}

// TestFloatAddition 验证浮点加法不截断：整型路径按 TypeName 识别，
// 任一操作数是浮点就走浮点运算（2 + 1/2 = 2.5 而非 2）。
func TestFloatAddition(t *testing.T) {
	wantStr(t, "2 + 1/2", "2.5")
	wantStr(t, "1/2 + 1/2", "1")
	wantStr(t, "0.5 + 0.25", "0.75")
	wantStr(t, "5/2 + 1", "3.5")
	// 整型加法与字符串拼接不受影响
	wantStr(t, "1 + 2", "3")
	wantStr(t, `"a" + "b"`, "ab")
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

// TestScopeModifiers 验证 $script:/$global:/$local: 作用域修饰符（读写、复合、增量、插值）。
func TestScopeModifiers(t *testing.T) {
	// $script: 写回：函数内改脚本作用域变量
	wantStr(t, "$sf = 0; function Set { $script:sf = 5 }; Set; $sf", "5")
	// $script: 读：函数内读脚本作用域变量
	wantStr(t, "$sf = 7; function Get { $script:sf }; Get", "7")
	// $global: 写：函数内写全局变量
	wantStr(t, "$gg = 0; function SetG { $global:gg = 9 }; SetG; $gg", "9")
	// $local: 只读写当前作用域，不落到外层
	wantStr(t, "$lv = 1; function L { $local:lv = 2; $lv }; L; $lv", "2", "1")
	// $script: 复合赋值
	wantStr(t, "$sf = 3; function A { $script:sf += 2 }; A; $sf", "5")
	// $script: 增量
	wantStr(t, "$sf = 1; function I { $script:sf++ }; I; $sf", "2")
	// 字符串插值里的 $script:
	wantStr(t, `$sf = 8; function S { "v=$script:sf" }; S`, "v=8")
	// ${} 括号形式
	wantStr(t, "$sf = 6; function B { ${script:sf} = 6 }; B; $sf", "6")
}

func TestControlFlow(t *testing.T) {
	wantStr(t, "if (5 -gt 3) { 'big' } else { 'small' }", "big")
	wantStr(t, "foreach ($i in 1..3) { $i }", "1", "2", "3")
	wantStr(t, "$s = 0; foreach ($i in 1..4) { $s += $i }; $s", "10")
	wantStr(t, "$i = 0; while ($i -lt 3) { $i++; if ($i -eq 2) { continue }; $i }", "1", "3")
	wantStr(t, "switch (2) { 1 { '一' } 2 { '二' } default { '其他' } }", "二")
	wantStr(t, "for ($i=0; $i -lt 2; $i++) { $i }", "0", "1")
}

// TestSwitchArray 验证 switch 对数组逐元素匹配与 break/continue 语义。
func TestSwitchArray(t *testing.T) {
	// 数组逐元素：每个元素跑全部 case，default 按元素判断
	wantStr(t, "switch (1,2,3) { 2 { 'hit2' } default { '无匹配' } }", "无匹配", "hit2", "无匹配")
	// 逐元素 default 输出 $_（元素值）
	wantStr(t, "switch (1,2,3) { default { $_ } }", "1", "2", "3")
	// continue 进入下一元素（命中后的尾巴不执行）
	wantStr(t, "switch (1,2,3) { 2 { 'hit'; continue; 'x' } default { 'd' } }", "d", "hit", "d")
	// break 退出整个 switch
	wantStr(t, "switch (1,2,3) { 2 { 'hit'; break; 'x' } default { 'd' } }", "d", "hit")
	// 标量时多个 case 都检查（PowerShell 语义）
	wantStr(t, "switch (2) { 1 { '一' } 2 { '二' } 2 { '二2' } }", "二", "二2")
	// 标量 continue 退出 switch，不再检查后续 case
	wantStr(t, "switch (5) { 5 { 'a'; continue } 5 { 'b' } }", "a")
	// 字符串标量只按整体匹配一次
	wantStr(t, `switch ("abc") { "abc" { "yes" } default { "no" } }`, "yes")
	// regex 模式数组逐元素
	wantStr(t, `switch -regex ("a1","b2","c3") { "a\d" { "A" } "\d" { "N" } }`, "A", "N", "N", "N")
}

// TestStatementAsExpression 验证语句可作赋值右侧（$x = switch / if / foreach ...）。
func TestStatementAsExpression(t *testing.T) {
	// switch 作表达式：数组逐元素输出
	wantStr(t, "$swr = switch (1,2,3) { default { $_ } }; $swr -join ','", "1,2,3")
	// switch 命中 case
	wantStr(t, "$v = switch (5) { 5 { 'five' } default { 'other' } }; $v", "five")
	// if 作表达式
	wantStr(t, "$iv = if ($true) { 'yes' } else { 'no' }; $iv", "yes")
	// foreach 作表达式
	wantStr(t, "$fv = foreach ($i in 1..3) { $i * 2 }; $fv -join ','", "2,4,6")
	// while 作表达式
	wantStr(t, "$i = 0; $wv = while ($i -lt 3) { $i; $i++ }; $wv -join ','", "0,1,2")
}

// TestStderrRedirectKeepsOutput 验证 2> 重定向不影响 stdout 输出（此前 applyRedirects 一律吞掉）。
func TestStderrRedirectKeepsOutput(t *testing.T) {
	// 2>$null 只丢弃错误流，输出照常（内置与函数各一例）
	wantStr(t, `Write-Output "w" 2>$null`, "w")
	wantStr(t, "function F { 'f-out' }; F 2>$null", "f-out")
	// > 文件：输出进文件，不进管道
	dir := t.TempDir()
	outFile := filepath.Join(dir, "o.txt")
	src := "Write-Output 'file' > " + outFile
	if got := strs(runEval(t, src)); len(got) != 0 {
		t.Fatalf("> 文件应无输出，实际 %v", got)
	}
	data, err := os.ReadFile(outFile)
	if err != nil || string(data) != "file\n" {
		t.Fatalf("> 文件应写入 file，err=%v data=%q", err, data)
	}
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

// TestParamBlock 验证函数 param() 声明块（默认值、多余实参进 $args、旧括号语法共存）。
func TestParamBlock(t *testing.T) {
	// param() 块形式
	wantStr(t, "function Fp { param($x) $x * 2 }; Fp 21", "42")
	// 默认值
	wantStr(t, "function Fd { param($x, $y = 100) $x + $y }; Fd 5; Fd 5 7", "105", "12")
	// 多余实参进 $args
	wantStr(t, `function Fa { param($x) "x=$x rest=$args" }; Fa 1 2 3`, "x=1 rest=2 3")
	// 旧括号语法仍可用
	wantStr(t, "function Old($a, $b) { $a + $b }; Old 1 2", "3")
	// 顶层 param() 不是脚本/函数开头时执行报错（无输出）
	wantStr(t, "param($z)")
}

// TestTryCatchFinally 验证 try/catch/finally + throw：
// 基本捕获、$_ 绑定、finally 恒执行、类型过滤、函数/循环传播、return 顺序、try 作表达式。
func TestTryCatchFinally(t *testing.T) {
	// 基本捕获与 $_ 绑定（错误记录的 Message 属性）
	wantStr(t, `try { throw "boom" } catch { "已捕获" }`, "已捕获")
	wantStr(t, `try { throw "msg1" } catch { $_.Message }`, "msg1")
	// catch 块不推独立作用域：普通变量赋值穿透到外层（对齐 PowerShell）
	wantStr(t, `$tc = "未执行"; try { throw "boom" } catch { $tc = "已捕获" }; $tc`, "已捕获")
	// catch 的 $_ 是临时绑定：块结束后外层 $_ 不受影响
	wantStr(t, `$old = $_; try { throw "x" } catch { $tmp = $_ }; $_ -eq $old`, "True")
	// 无异常时 catch 不执行，finally 恒执行
	wantStr(t, `try { "正常体" } catch { "不会到" } finally { "finally" }`, "正常体", "finally")
	// 捕获后继续执行后续语句
	wantStr(t, `try { throw "x" } catch { "c" }; "继续"`, "c", "继续")
	// catch [System.Exception] 基类全捕
	wantStr(t, `try { throw "e1" } catch [System.Exception] { "基类捕获" }`, "基类捕获")
	// catch 精确类型不匹配 → 未捕获，顶层打印（无输出）
	wantStr(t, `try { throw "e2" } catch [System.ArgumentException] { "不会到" }`)
	// 多 catch 顺序：第一个匹配生效
	wantStr(t, `try { throw "x" } catch [System.ArgumentException] { "A" } catch { "全捕" }`, "全捕")
	// 函数内 throw 被调用方 try 捕获
	wantStr(t, `function f { throw "函数错" }; try { f } catch { "捕获" }`, "捕获")
	// 循环内 throw 传播到外层 try，且 throw 前循环已输出保留
	wantStr(t, `try { foreach ($i in 1..3) { if ($i -eq 2) { throw "循环错" }; "i=$i" } } catch { "循环捕获" }`, "i=1", "循环捕获")
	// return 时 finally 恒执行，且 finally 输出在返回值之前（对齐 PowerShell）
	wantStr(t, `function g { try { return "r" } finally { "fin" } }; g`, "fin", "r")
	// 嵌套 try：内层类型不匹配 → 外层捕获
	wantStr(t, `try { try { throw "内层错" } catch [System.ArgumentException] { "不会到" } } catch { "外层捕获" }`, "外层捕获")
	// catch 内重抛，外层再捕获
	wantStr(t, `try { try { throw "a" } catch { "内捕"; throw } } catch { "外捕" }`, "内捕", "外捕")
	// finally 里 throw 覆盖原错误
	wantStr(t, `try { throw "原错" } catch { "会捕获" } finally { throw "finally错" }`, "会捕获")
	// catch 里的 break 传播到外层循环
	wantStr(t, `for ($i = 0; $i -lt 2; $i++) { try { throw "t" } catch { "i=$i"; break } }; "结束"`, "i=0", "结束")
	// try 作表达式（赋值右侧）
	wantStr(t, `$r = try { "成功值" } catch { "失败值" }; "r=$r"`, "r=成功值")
	wantStr(t, `$r2 = try { throw "bad" } catch { "捕获值" }; "r2=$r2"`, "r2=捕获值")
	// 无 catch 只有 finally：finally 仍执行并输出，错误继续上抛（顶层打印，无输出）
	wantStr(t, `try { throw "只有finally" } finally { "清理" }`, "清理")
	// return 前与 finally 的输出在函数返回值之前（调用点统一语义）
	wantStr(t, `function f { "前"; return "r" }; f`, "前", "r")
	// 脚本块 return 沿途输出保留（与函数一致）
	wantStr(t, `"x" | ForEach-Object { "块前"; return "块后" }`, "块前", "块后")
	// 子表达式 return 沿途输出保留
	wantStr(t, `$("a"; return "b")`, "a", "b")
	// switch 不推独立作用域：case 块内普通赋值穿透外层（与 foreach 同机制）
	wantStr(t, `switch (1) { 1 { $sv = 5 } }; $sv`, "5")
}

// TestTryScriptPropagation 验证脚本内 throw 跨脚本传播与调用方 try 捕获。
func TestTryScriptPropagation(t *testing.T) {
	dir := t.TempDir()
	srcPath := filepath.Join(dir, "src.ps1")
	if err := os.WriteFile(srcPath, []byte("param($who)\n\"脚本运行中\"\nthrow \"脚本抛错 $who\"\n\"脚本尾部\"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	callPath := filepath.Join(dir, "call.ps1")
	callSrc := "try { " + srcPath + " \"X\" } catch { \"调用方捕获: $($_.Message)\" }\n\"调用方继续\"\n"
	if err := os.WriteFile(callPath, []byte(callSrc), 0644); err != nil {
		t.Fatal(err)
	}
	// 跨脚本：被调脚本 throw 前输出保留，调用方捕获后继续
	wantStr(t, callSrc, "脚本运行中", "调用方捕获: 脚本抛错 X", "调用方继续")
	// 顶层逐语句（REPL 语义）：未捕获错误打印到 stderr（io.Discard），会话继续执行后续语句
	wantStr(t, "\"前\"; throw \"停\"; \"后\"", "前", "后")
}


func TestScriptParam(t *testing.T) {
	dir := t.TempDir()
	scriptPath := filepath.Join(dir, "s.ps1")
	content := "param($a, $b = 10)\n\"a=$a b=$b args=$args\""
	if err := os.WriteFile(scriptPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := New(sess, strings.NewReader(""), io.Discard, io.Discard)

	// 显式实参按位置绑定
	out := ev.RunScriptFile(scriptPath, []*object.PSObject{object.Int(5), object.Int(6)})
	if got := strs(out); len(got) != 1 || got[0] != "a=5 b=6 args=" {
		t.Fatalf("显式实参 → %v", got)
	}
	// 缺参用默认值
	out = ev.RunScriptFile(scriptPath, nil)
	if got := strs(out); len(got) != 1 || got[0] != "a= b=10 args=" {
		t.Fatalf("默认值 → %v", got)
	}
	// 多余实参保留在 $args
	out = ev.RunScriptFile(scriptPath, []*object.PSObject{object.Int(1), object.Int(2), object.Int(3)})
	if got := strs(out); len(got) != 1 || got[0] != "a=1 b=2 args=3" {
		t.Fatalf("剩余实参 → %v", got)
	}
}

func TestHashtableAndSubexpr(t *testing.T) {
	wantStr(t, "$h = @{ Name = 'x'; Count = 2 }; $h.Count", "2")
	wantStr(t, "$(1 + 2)", "3")
	wantStr(t, `"total: $(6 * 7)"`, "total: 42")
}

func TestHashtableProps(t *testing.T) {
	// 对照真实 PowerShell 语义：Count 返回条目数，键优先于属性
	wantStr(t, "@{a=1;b=2}.Count", "2")
	wantStr(t, "@{}.Count", "0")
	wantStr(t, "@{Count=5;x=1}.Count", "5")
	wantStr(t, "@{a=1;b=2}.Length", "1")
	wantStr(t, "@{a=1;b=2}.Keys", "a", "b")
	wantStr(t, "@{a=1;b=2}.Values", "1", "2")
	wantStr(t, "@{a=1;b=2}.Keys.Count", "2")
	wantStr(t, "@{a=1;b=2}.Keys[0]", "a")
}

func TestPSVersionTableCore(t *testing.T) {
	// 7.X 风格：PSVersion 只标 7（不对齐具体小版本），可读 .Major
	wantStr(t, "$PSVersionTable.PSVersion.Major", "7")
	wantStr(t, "$PSVersionTable.PSVersion.Minor", "0")
	wantStr(t, "$PSVersionTable.PSVersion", "7")
	wantStr(t, "$PSVersionTable.PSEdition", "Core")
	wantStr(t, "$PSVersionTable.OS", "Linux")
	wantStr(t, "$PSVersionTable.GitCommitId", "0000000000000000000000000000000000000000")
}

func TestPSVersionTableDesktop(t *testing.T) {
	check := func(src string, want ...string) {
		t.Helper()
		got := strs(runEvalWithStyle(t, shell.StyleDesktop, src))
		if len(got) != len(want) {
			t.Fatalf("%q → %v，想要 %v", src, got, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Fatalf("%q → %v，想要 %v", src, got, want)
			}
		}
	}
	check("$PSVersionTable.PSVersion.Major", "5")
	check("$PSVersionTable.PSVersion.Minor", "1")
	check("$PSVersionTable.PSVersion", "5.1")
	check("$PSVersionTable.PSEdition", "Desktop")
}

func TestMatchesCapture(t *testing.T) {
	// 标量 -match 成功后填充 $Matches：0 是整体匹配，1.. 是捕获组
	wantStr(t, `$x = "abc123" -match "(\d+)"; $Matches[0]`, "123")
	wantStr(t, `$x = "abc123" -match "(\d+)"; $Matches[1]`, "123")
	wantStr(t, `$x = "2024-08-16" -match "^(\d+)-(\d+)-(\d+)$"; $Matches[3]`, "16")
	// 命名组用组名取
	wantStr(t, `$x = "abc123" -match "(?<letters>[a-z]+)(?<digits>\d+)"; $Matches.letters`, "abc")
	wantStr(t, `$x = "abc123" -match "(?<letters>[a-z]+)(?<digits>\d+)"; $Matches.digits`, "123")
	// 未参与的可选组不写入，且未命名组序号按全部未命名组计
	wantStr(t, `$x = "abc" -match "(\d+)?(a.*)"; $Matches.Keys -join ","`, "0,2")
	// 不匹配不清空旧值
	wantStr(t, `$x = "abc123" -match "(\d+)"; $y = "x" -match "(\d+)"; $Matches[1]`, "123")
	// 数组左值不设置 $Matches
	wantStr(t, `$x = "a","b1" -match "(\d)"; $Matches -eq $null`, "True")
	// 未匹配过时 $Matches 为 $null
	wantStr(t, "$Matches -eq $null", "True")
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

package eval

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"runtime"
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

// osName 复刻 shell.OSName 的平台映射，供测试构造跨平台的 OS 期望值。
func osName() string {
	switch runtime.GOOS {
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	case "darwin":
		return "Darwin"
	}
	return runtime.GOOS
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

// TestFloatAddition 验证浮点加法不截断：整型路径按 TypeName 识别，任一操作数是浮点就走浮点运算（2 + 1/2 = 2.5 而非 2）。
func TestFloatAddition(t *testing.T) {
	wantStr(t, "2 + 1/2", "2.5")
	wantStr(t, "1/2 + 1/2", "1")
	wantStr(t, "0.5 + 0.25", "0.75")
	wantStr(t, "5/2 + 1", "3.5")
	// 整型加法与字符串拼接不受影响
	wantStr(t, "1 + 2", "3")
	wantStr(t, `"a" + "b"`, "ab")
}

// TestFloatIncrement 验证浮点变量 ++/-- 不截断（$i = 0.5; $i++ → 1.5），整型与未定义变量行为不变。
func TestFloatIncrement(t *testing.T) {
	wantStr(t, "$i = 0.5; $i++; $i", "1.5")
	wantStr(t, "$i = 1.5; $i--; $i", "0.5")
	wantStr(t, "$i = 0.5; $i++; $i++; $i", "2.5")
	// 整型增量不变
	wantStr(t, "$i = 1; $i++; $i", "2")
	wantStr(t, "$i = 3; $i--; $i", "2")
	// 未定义变量从 0 起增
	wantStr(t, "$i++; $i", "1")
	// 复合赋值 += 浮点（走 addOp）
	wantStr(t, "$i = 0.5; $i += 1; $i", "1.5")
}

// TestDivideByZero 验证除零报错（对齐 PowerShell）：
// 结果置 $null、$? 置 false，REPL 语义下后续语句继续执行；正常除法与取模不受影响。
func TestDivideByZero(t *testing.T) {
	// 除零与模零：报错无输出（错误写到 stderr）
	wantStr(t, "5/0")
	wantStr(t, "5.0/0")
	wantStr(t, "1 % 0")
	// 赋值得到 $null
	wantStr(t, "$x = 5/0; $x -eq $null", "True")
	wantStr(t, "$x = 1 % 0; $x -eq $null", "True")
	// 除零后 $? 为 false（用 if 读取，赋值语句会先置位 $?）
	wantStr(t, "5/0; if ($?) { \"ok\" } else { \"err\" }", "err")
	// 后续语句继续执行
	wantStr(t, "1/0; \"继续\"", "继续")
	// 正常除法与取模不受影响
	wantStr(t, "5/2", "2.5")
	wantStr(t, "6/3", "2")
	wantStr(t, "5 % 2", "1")
}

// TestSplitMaxSubstrings 验证 -split 的最大子串数参数（"a,b,c" -split ",",2 → a、b,c），末段保留未分割剩余，0/负数/超上限不限段数。
func TestSplitMaxSubstrings(t *testing.T) {
	wantStr(t, `"a,b,c" -split ",",2`, "a", "b,c")
	wantStr(t, `"a,b,c" -split ",",1`, "a,b,c")
	wantStr(t, `"a,b,c" -split ",",5`, "a", "b", "c")
	wantStr(t, `"a,b,c" -split ",",0`, "a", "b", "c")
	wantStr(t, `"a,b,c" -split ",",-1`, "a", "b", "c")
	wantStr(t, `"a,b," -split ",",2`, "a", "b,")
	// 无最大子串数参数行为不变
	wantStr(t, `"a-b-c" -split "-"`, "a", "b", "c")
	// 正则分隔符同样生效
	wantStr(t, `"a1b22c" -split "\d+",2`, "a", "b22c")
}

func TestStringOps(t *testing.T) {
	wantStr(t, `"a" + "b"`, "ab")
	wantStr(t, `"ab" * 3`, "ababab")
	wantStr(t, `"abc".ToUpper()`, "ABC")
	wantStr(t, `"Hello".Length`, "5")
	wantStr(t, `"a,b,c".Split(",")[1]`, "b")
}

// TestRangeIndex 验证范围/逗号多下标索引的取值形态。
// 覆盖 $a[1..2]、$a[0,2]、负数从末尾数、嵌套下标展平、越界补 $null、字符串范围索引。
func TestRangeIndex(t *testing.T) {
	// 范围下标：逐元素取值
	wantStr(t, "$a = 1,2,3,4; $a[1..2]", "2", "3")
	wantStr(t, "$a = 1,2,3,4; $a[0..2]", "1", "2", "3")
	// 单个结果展开为标量
	wantStr(t, "$a = 1,2,3,4; $a[1..1]", "2")
	// 负数从末尾数（1..-1 = 1,0,-1）
	wantStr(t, "$a = 1,2,3,4; $a[1..-1]", "2", "1", "4")
	wantStr(t, "$a = 1,2,3,4; $a[-1..-3]", "4", "3", "2")
	// 逗号多下标
	wantStr(t, "$a = 1,2,3,4; $a[0,2]", "1", "3")
	// 嵌套下标数组展平（$a[1..2,0] = 下标 1,2,0）
	wantStr(t, "$a = 1,2,3,4; $a[1..2,0]", "2", "3", "1")
	// 变量范围
	wantStr(t, "$a = 1,2,3,4; $x = 1; $y = 2; $a[$x..$y]", "2", "3")
	// 降序范围
	wantStr(t, "$a = 1,2,3,4; $a[2..0]", "3", "2", "1")
	// 越界补 $null（显示为空；$a[1..9] 共 9 个位置，超出部分逐位补）
	wantStr(t, "$a = 1,2,3,4; $a[1..9]", "2", "3", "4", "", "", "", "", "", "")
	// 字符串范围索引：返回字符数组
	wantStr(t, `"abcdef"[1..3]`, "b", "c", "d")
	wantStr(t, `"abcdef"[1..-2]`, "b", "a", "f", "e")
	// 单个字符展开为标量
	wantStr(t, `"abcdef"[1..1]`, "b")
	// 标量下标行为不变
	wantStr(t, "$a = 1,2,3,4; $a[1]", "2")
	wantStr(t, "$a = 1,2,3,4; $a[-1]", "4")
	wantStr(t, `"abcdef"[2]`, "c")
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

// TestStringNumberComparison 验证字符串与数字混合比较按左操作数类型转换（对齐 PowerShell）：两个字符串按字典序（"5" -lt "10" 为 False），数字对字符串按数字（5 -lt "10" 为 True），$null 只与 $null 相等。
func TestStringNumberComparison(t *testing.T) {
	// 字符串-字符串：字典序
	wantStr(t, `"5" -lt "10"`, "False")
	wantStr(t, `"10" -lt "5"`, "True")
	wantStr(t, `"5" -gt "10"`, "True")
	wantStr(t, `"abc" -lt "abd"`, "True")
	wantStr(t, `"a" -lt "B"`, "True") // 大小写不敏感
	// 大小写敏感变体同样按字符串
	wantStr(t, `"5" -clt "10"`, "False")
	// 数字-字符串：右操作数转数字
	wantStr(t, `5 -lt "10"`, "True")
	wantStr(t, `1 -lt "2"`, "True")
	wantStr(t, `2 -gt "10"`, "False")
	// 字符串-数字：右操作数转字符串
	wantStr(t, `"5" -lt 10`, "False")
	// 相等：双向转换一致
	wantStr(t, `5 -eq "5"`, "True")
	wantStr(t, `"5" -eq 5`, "True")
	wantStr(t, `1 -eq 1.0`, "True")
	wantStr(t, `$true -eq 1`, "True")
	// 布尔-数字顺序：$true=1、$false=0 参与数字比较
	wantStr(t, `$true -lt 2`, "True")
	wantStr(t, `$false -lt 1`, "True")
	wantStr(t, `$true -gt 1`, "False")
	wantStr(t, `$true -ge 1`, "True")
	wantStr(t, `$true -clt 2`, "True")
	wantStr(t, `$false -clt 1`, "True")
	// $null 只与 $null 相等
	wantStr(t, `$null -eq $null`, "True")
	wantStr(t, `$null -eq ""`, "False")
	wantStr(t, `"" -eq $null`, "False")
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

// TestStderrRedirectKeepsOutput 验证 2> 重定向不影响 stdout 输出。
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

// TestParamTypeAnnotations 验证形参 [类型] 标注：
// 位置与命名实参、默认值都经类型转换，数组标注把单值包成单元素数组，无法转换时不执行被调方。
func TestParamTypeAnnotations(t *testing.T) {
	// 位置实参转换成声明类型
	wantStr(t, `function Fti([int]$x) { $x -is [int] }; Fti '42'`, "True")
	// 命名实参同样转换
	wantStr(t, `function Ftn([int]$n) { $n -is [int] }; Ftn -n '33'`, "True")
	// 默认值也经类型转换
	wantStr(t, `function Ftd([int]$y = '5') { $y -is [int] }; Ftd`, "True")
	// double 标注
	wantStr(t, `function Ftw([double]$v) { $v }; Ftw '2.5'`, "2.5")
	// 数组标注把单值包成单元素数组
	wantStr(t, `function Fta([int[]]$a) { $a.Count; $a[0] -is [int] }; Fta 5`, "1", "True")
	// 数组实参逐元素转换
	wantStr(t, `function Ftb([string[]]$s) { $s.Count; $s[1] }; Ftb 1,2,3`, "3", "2")
	// 未标注的形参保持原值不转换
	wantStr(t, "function Ftu($x) { $x -is [string] }; Ftu '42'", "True")
	// 无法转换时不执行函数体，后续语句继续
	wantStr(t, `function Ftf([int]$x) { "body" }; Ftf 'abc'; "after"`, "after")
	// 绑定失败置 $? 为 false
	wantStr(t, `function Ftq([int]$x) { "b" }; Ftq 'abc'; if ($?) { "yes" } else { "no" }`, "no")
}

// TestScriptParamTypeAnnotations 验证脚本 param() 的类型标注：
// 可转换实参转成声明类型，无法转换时脚本主体不再执行并置失败退出码。
func TestScriptParamTypeAnnotations(t *testing.T) {
	dir := t.TempDir()
	okPath := filepath.Join(dir, "ok.ps1")
	if err := os.WriteFile(okPath, []byte("param([int]$n)\n$n\n$n -is [int]\n"), 0644); err != nil {
		t.Fatal(err)
	}
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := New(sess, strings.NewReader(""), io.Discard, io.Discard)
	out := ev.RunScriptFile(okPath, []*object.PSObject{object.Str("11")})
	if got := strs(out); len(got) != 2 || got[0] != "11" || got[1] != "True" {
		t.Fatalf("脚本类型标注 → %v，想要 [11 True]", got)
	}

	badPath := filepath.Join(dir, "bad.ps1")
	if err := os.WriteFile(badPath, []byte("param([int]$n)\n\"不应输出\"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	out = ev.RunScriptFile(badPath, []*object.PSObject{object.Str("abc")})
	if got := strs(out); len(got) != 0 {
		t.Fatalf("绑定失败脚本应无输出，实际 %v", got)
	}
	if sess.LastExit != 1 {
		t.Fatalf("绑定失败应置失败退出码 1，实际 %d", sess.LastExit)
	}
}

// TestInvokeOperator 验证 & 调用运算符：
// 脚本块直调、变量持块、param 形参、$args、return 顺序、作表达式、throw 可捕获、动态作用域、按名字调命令。
func TestInvokeOperator(t *testing.T) {
	// 基本执行与多输出
	wantStr(t, `& { "a"; "b" }`, "a", "b")
	// 变量持有脚本块后调用
	wantStr(t, `$sb = { param($x) $x * 2 }; & $sb 21`, "42")
	// 块体开头 param() 提取为形参，多余实参进 $args
	wantStr(t, `$sb2 = { param($x) "x=$x args=$args" }; & $sb2 1 2 3`, "x=1 args=2 3")
	// 无 param 时实参全进 $args
	wantStr(t, `& { "n=$($args.Count)" } e1 e2`, "n=2")
	// return 前的输出保留，return 后不再执行
	wantStr(t, `& { "before"; return "val"; "after" }`, "before", "val")
	// 作表达式（赋值右侧）
	wantStr(t, `$v = & { 6 * 7 }; $v`, "42")
	// 块内 throw 可被调用方捕获
	wantStr(t, `try { & { throw "blk" } } catch { "caught: $($_.Message)" }`, "caught: blk")
	// 动态作用域：块内可见外层变量
	wantStr(t, `$outer = 10; & { $outer }`, "10")
	// 管道输入进 $input（$_ 不绑定），输出时逐项枚举
	wantStr(t, `1,2 | & { $input }`, "1", "2")
	// 变量持命令名按名字分发
	wantStr(t, `$c = "Write-Output"; & $c hi`, "hi")
	// 字符串字面量目标
	wantStr(t, `& 'Write-Output' hi`, "hi")
	// .Invoke 执行并把输出收集成数组
	wantStr(t, `$sb3 = { 1; 2 }; $r = $sb3.Invoke(); $r.Count`, "2")
	wantStr(t, `$sb4 = { param($a,$b) $a + $b }; $sb4.Invoke(5,6)`, "11")
	// 目标缺失报错且后续语句继续
	wantStr(t, `&; "after"`, "after")
}

// TestNamedBlocks 验证函数的 begin/process/end 命名块与 filter：
// 三段顺序、直调 process 以 $null 跑一次、process 逐项绑 $_、return 只结束本次、零输入跳过 process、块可乱序。
func TestNamedBlocks(t *testing.T) {
	// 三块齐全：begin 一次 → process 逐项 → end 一次
	wantStr(t, `function T1 { begin { "b" } process { "p:$_" } end { "e" } }; 1,2,3 | T1`, "b", "p:1", "p:2", "p:3", "e")
	// 直调：无管道输入，process 以 $null 跑一次
	wantStr(t, `function T2 { begin { "b" } process { "p" } end { "e" } }; T2`, "b", "p", "e")
	// 只有 process：管道逐项、直调跑一次
	wantStr(t, `function P1 { process { "p:$_" } }; 1,2 | P1`, "p:1", "p:2")
	wantStr(t, `function P2 { process { "p" } }; P2`, "p")
	// 块乱序仍按 begin/process/end 顺序执行
	wantStr(t, `function O1 { end { "e" } begin { "b" } process { "x:$_" } }; 9 | O1`, "b", "x:9", "e")
	// 零输入：begin/end 跑，process 不跑
	wantStr(t, `$a = @(); function Z1 { begin { "zb" } process { "zp:$_" } end { "ze" } }; $r = @($a | Z1); $r.Count`, "2")
	// param 与命名块共存，形参在 process 内可用
	wantStr(t, `function Q1([int]$mul) { process { $_ * $mul } }; 1,2,3 | Q1 -mul 10`, "10", "20", "30")
	// filter 的 Body 即 process
	wantStr(t, `filter F1 { "f:$_" }; 1,2 | F1`, "f:1", "f:2")
	wantStr(t, `filter F2 { "f" }; F2`, "f")
	// filter 也可带命名块
	wantStr(t, `filter F3 { begin { "fb" } process { "fp:$_" } }; 1 | F3`, "fb", "fp:1")
}

// TestNamedBlocksControlFlow 验证命名块内的 return/break/continue：
// process 里 return 只结束本次；break/continue 无所属循环时沿调用栈上抛终止当前语句序列。
func TestNamedBlocksControlFlow(t *testing.T) {
	// process 内 return 跳过本次剩余语句，继续下一项
	wantStr(t, `function R1 { process { if ($_ -eq 2) { return }; "r:$_" } }; 1,2,3 | R1`, "r:1", "r:3")
	// begin 内 return 结束整个函数（end 不再跑）
	wantStr(t, `function R2 { begin { "b"; return } process { "p" } end { "e" } }; R2`, "b")
	// process 内 break 上抛：中断外层循环体，循环后的语句继续
	wantStr(t, `function B1 { process { if ($_ -eq 2) { break }; "b:$_" } }; foreach ($i in 1..2) { 1,2,3 | B1; "inner" }; "after"`, "b:1", "after")
	// process 内 continue 上抛：终止当前语句（REPL 逐语句模式下后续语句照常）
	wantStr(t, `"start"; function C1 { process { if ($_ -eq 1) { continue }; "c:$_" } }; (1,2,3 | C1); "next"`, "start", "next")
}

// TestErrorVariable 验证 $Error 自动变量：
// 非终止错误与 throw 都累积、最新在 [0]、被 catch 的错误也进、Clear/RemoveAt 落到记录本体。
func TestErrorVariable(t *testing.T) {
	// 非终止错误累积，[0] 是最新
	wantStr(t, `[int]"abc"; $Error.Count`, "1")
	wantStr(t, `[int]"abc"; [int]"def"; $Error.Count; $Error[0].Message`,
		"2", `无法将值“def”转换为类型“int”。`)
	// throw 与被捕获的错误都进 $Error
	wantStr(t, `try { throw "boom" } catch { "in=$($Error.Count)" }; "after=$($Error.Count)"; $Error[0].Message`,
		"in=1", "after=1", "boom")
	// 无错误时为空数组
	wantStr(t, `$Error.Count`, "0")
	// Clear 清空本体
	wantStr(t, `[int]"abc"; $Error.Clear(); $Error.Count`, "0")
	// RemoveAt 删除指定下标
	wantStr(t, `[int]"abc"; [int]"def"; $Error.RemoveAt(0); $Error.Count; $Error[0].Message`,
		"1", `无法将值“abc”转换为类型“int”。`)
	// 参数绑定失败同样累积
	wantStr(t, `function F([int]$k) { "never" }; F 'bad'; $Error.Count; $Error[0].Message`,
		"1", `无法把实参“bad”转换成形参 k 声明的类型“int”。`)
}

// TestErrorAction 验证 -ErrorAction：默认继续、SilentlyContinue 与 Ignore 压住显示、Stop 转终止错误可被捕获、无效取值报绑定错误。
func TestErrorAction(t *testing.T) {
	// 默认 Continue：记录并置 $? 为 false，后续语句继续（$? 用 if 读取，直接读会被表达式语句先置位）
	wantStr(t, `Get-Item 不存在XYZ123; $Error.Count`, "1")
	wantStr(t, `Get-Item 不存在XYZ123; if ($?) { "ok" } else { "fail" }`, "fail")
	// SilentlyContinue：仍记录仍置失败，只是不显示
	wantStr(t, `Get-Item 不存在XYZ123 -ErrorAction SilentlyContinue; $Error.Count`, "1")
	wantStr(t, `Get-Item 不存在XYZ123 -ErrorAction SilentlyContinue; if ($?) { "ok" } else { "fail" }`, "fail")
	// Ignore：不记录，只置失败
	wantStr(t, `Get-Item 不存在XYZ123 -ErrorAction Ignore; $Error.Count`, "0")
	wantStr(t, `Get-Item 不存在XYZ123 -ErrorAction Ignore; if ($?) { "ok" } else { "fail" }`, "fail")
	// Inquire 在非交互场景按终止错误处理
	wantStr(t, `try { Get-Item 不存在XYZ123 -ErrorAction Inquire } catch { "caught" }; $Error.Count`, "caught", "1")
	// Stop：转为终止错误，可被 try/catch 捕获，只记一次
	wantStr(t, `try { Get-Item 不存在XYZ123 -ErrorAction Stop } catch { "caught" }; $Error.Count`, "caught", "1")
	// Stop 中断 try 体后续语句
	wantStr(t, `try { Get-Item 不存在XYZ123 -ErrorAction Stop; "after" } catch { "caught" }`, "caught")
	// 取值大小写不敏感
	wantStr(t, `try { Get-Item 不存在XYZ123 -ErrorAction stop } catch { "caught" }`, "caught")
	// 无效取值按绑定错误报告
	wantStr(t, `Get-Item foo -ErrorAction Bogus; $Error.Count`, "1")
	wantStr(t, `Get-Item foo -ErrorAction Bogus; if ($?) { "ok" } else { "fail" }`, "fail")
	// 不带值的 -ErrorAction 按缺值报绑定错误
	wantStr(t, `Get-Item foo -ErrorAction; $Error.Count`, "1")
}

// TestErrorActionOutput 验证显示侧：默认 Continue 写 stderr，SilentlyContinue 与 Ignore 不写。
func TestErrorActionOutput(t *testing.T) {
	runWithStderr := func(src string) string {
		var errBuf bytes.Buffer
		sess := shell.New(shell.StyleCore, io.Discard, &errBuf, strings.NewReader(""))
		ev := New(sess, strings.NewReader(""), io.Discard, &errBuf)
		res := parser.Parse(src)
		if res.Error != nil {
			t.Fatalf("解析错误 %q: %v", src, res.Error)
		}
		for _, st := range res.List.Statements {
			ev.EvalStatement(st)
		}
		return errBuf.String()
	}
	if out := runWithStderr(`Get-Item 不存在XYZ123`); out == "" {
		t.Error("默认 Continue 应写 stderr")
	}
	if out := runWithStderr(`Get-Item 不存在XYZ123 -ErrorAction SilentlyContinue`); out != "" {
		t.Errorf("SilentlyContinue 不应写 stderr，得到 %q", out)
	}
	if out := runWithStderr(`Get-Item 不存在XYZ123 -ErrorAction Ignore`); out != "" {
		t.Errorf("Ignore 不应写 stderr，得到 %q", out)
	}
	if out := runWithStderr(`$ErrorActionPreference = 'SilentlyContinue'; Get-Item 不存在XYZ123`); out != "" {
		t.Errorf("首选项 SilentlyContinue 不应写 stderr，得到 %q", out)
	}
}

// TestErrorActionPreference 验证 $ErrorActionPreference：默认值、首选项分发、显式参数覆盖、函数内局部生效、无效赋值报错且不生效。
func TestErrorActionPreference(t *testing.T) {
	// 未赋值时读到默认值
	wantStr(t, `$ErrorActionPreference`, "Continue")
	// 首选项 SilentlyContinue：记录但不显示
	wantStr(t, `$ErrorActionPreference = 'SilentlyContinue'; Get-Item 不存在XYZ123; $Error.Count`, "1")
	// 首选项 Stop：转为终止错误，可捕获且只记一次
	wantStr(t, `$ErrorActionPreference = 'Stop'; try { Get-Item 不存在XYZ123 } catch { "caught" }; $Error.Count`, "caught", "1")
	// 显式 -ErrorAction 覆盖首选项
	wantStr(t, `$ErrorActionPreference = 'Stop'; Get-Item 不存在XYZ123 -ErrorAction Continue; $Error.Count`, "1")
	// 函数内赋值只在局部生效
	wantStr(t, `function Pref { $ErrorActionPreference = 'Stop'; try { Get-Item 不存在XYZ123 } catch { "caught" } }; Pref; $ErrorActionPreference`, "caught", "Continue")
	// 无效赋值报错且不生效
	wantStr(t, `$ErrorActionPreference = 'Bogus'; $Error.Count`, "1")
	wantStr(t, `$ErrorActionPreference = 'Bogus'; $ErrorActionPreference`, "Continue")
	// 大小写混写仍落到同一首选项
	wantStr(t, `$erroractionpreference = 'Stop'; try { Get-Item 不存在XYZ123 } catch { "caught" }`, "caught")
	// 空值恢复默认
	wantStr(t, `$ErrorActionPreference = 'Stop'; $ErrorActionPreference = $null; $ErrorActionPreference`, "Continue")
	// Ignore 同样置失败
	wantStr(t, `Get-Item 不存在XYZ123 -ErrorAction Ignore; if ($?) { "ok" } else { "fail" }`, "fail")
	// Set-Variable 与 Clear-Variable 使用同样的校验
	wantStr(t, `Set-Variable -Name ErrorActionPreference -Value 'Bogus'; $Error.Count`, "1")
	wantStr(t, `$ErrorActionPreference = 'Stop'; Clear-Variable ErrorActionPreference; $ErrorActionPreference`, "Continue")
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
	// 对照 PowerShell 语义：Count 返回条目数，键优先于属性
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
	// OS 按 runtime.GOOS 报告，期望值跟随运行平台（Linux 上为 Linux、Windows 上为 Windows）
	wantStr(t, "$PSVersionTable.OS", osName())
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

// TestTrimNoArgs 验证 TrimStart/TrimEnd 无参清空白（有参按字符集裁剪行为不变）。
func TestTrimNoArgs(t *testing.T) {
	wantStr(t, `"  x  ".TrimStart()`, "x  ")
	wantStr(t, `"  x  ".TrimEnd()`, "  x")
	wantStr(t, `"  x  ".Trim()`, "x")
	wantStr(t, `"007".TrimStart("0")`, "7")
	wantStr(t, `"007".TrimEnd("7")`, "00")
}

// TestSplitNoArgs 验证 Split 无参按任意空白分割（含 tab/换行，连续空白合并），有参行为不变。
func TestSplitNoArgs(t *testing.T) {
	wantStr(t, "\"a\tb c\".Split()", "a", "b", "c")
	wantStr(t, "\"a  b   c\".Split()", "a", "b", "c")
	wantStr(t, `"a,b".Split(",")`, "a", "b")
	wantStr(t, `"a,b,".Split(",")`, "a", "b", "")
}

// TestStringMethodFamily 验证字符串方法族：LastIndexOf/Remove/PadLeft/PadRight/Insert。
func TestStringMethodFamily(t *testing.T) {
	wantStr(t, `"abcabc".LastIndexOf("b")`, "4")
	wantStr(t, `"abcabc".LastIndexOf("b", 2)`, "1")
	wantStr(t, `"abc".LastIndexOf("x")`, "-1")
	wantStr(t, `"abc".Remove(1)`, "a")
	wantStr(t, `"abc".Remove(1, 1)`, "ac")
	wantStr(t, `"7".PadLeft(3, "0")`, "007")
	wantStr(t, `"7".PadLeft(3)`, "  7")
	wantStr(t, `"7".PadRight(3, "0")`, "700")
	wantStr(t, `"abc".Insert(1, "X")`, "aXbc")
	wantStr(t, `"abc".Insert(3, "!")`, "abc!")
}

// TestDateTimeMethods 验证 DateTime 方法族：ToShortDateString/ToLongDateString/ToShortTimeString/ToLongTimeString/ToString/ToFileTime。
func TestDateTimeMethods(t *testing.T) {
	wantStr(t, `(Get-Date -Date "2020-01-15").ToShortDateString()`, "1/15/2020")
	wantStr(t, `(Get-Date -Date "2020-01-15").ToLongDateString()`, "Wednesday, January 15, 2020")
	wantStr(t, `(Get-Date -Date "2020-01-15 08:30:00").ToShortTimeString()`, "8:30 AM")
	wantStr(t, `(Get-Date -Date "2020-01-15 08:30:00").ToLongTimeString()`, "8:30:00 AM")
	wantStr(t, `(Get-Date -Date "2020-01-15").ToString()`, "1/15/2020 12:00:00 AM")
	// ToFileTime 取该时刻的 Windows 文件时间刻度；用带时区的 ISO 串固定为 UTC 零点，使期望值与时区无关。
	wantStr(t, `(Get-Date -Date "2020-01-15T00:00:00Z").ToFileTime()`, "132235200000000000")
	// 文件时间纪元 1601-01-01 UTC 的刻度为 0，验证远古日期不再溢出
	wantStr(t, `(Get-Date -Date "1601-01-01T00:00:00Z").ToFileTime()`, "0")
}

// TestTypeCasts 验证方括号强制转换的基础类型与数组后缀、失败报错、[void] 丢弃结果。
func TestTypeCasts(t *testing.T) {
	wantStr(t, `[int]"42"`, "42")
	wantStr(t, `[int]$true`, "1")
	wantStr(t, `[double]"1.5"`, "1.5")
	wantStr(t, `[string]42`, "42")
	wantStr(t, `[bool]""`, "False")
	wantStr(t, `[bool]"a"`, "True")
	wantStr(t, `[void](1 + 1)`)
	wantStr(t, `$d = [datetime]"2020-01-02"; $d.Year`, "2020")
	wantStr(t, `$h = @{a = 1}; ([hashtable]$h)["a"]`, "1")
	wantStr(t, `$a = [int[]](1, 2, "3"); $a -join ","`, "1,2,3")
	// -is / -as 直接消费类型字面量
	wantStr(t, `1 -is [int]`, "True")
	wantStr(t, `"a" -is [int]`, "False")
	wantStr(t, `1 -isnot [string]`, "True")
	wantStr(t, `"7" -as [int]`, "7")
	// 变量保存类型字面量后再用
	wantStr(t, `$t = [double]; "1.5" -as $t`, "1.5")
	// 转换失败：写错误、返回空，后续语句继续执行
	wantStr(t, `[int]"abc"`)
	wantStr(t, `[int]"abc"; "继续"`, "继续")
}

// TestStaticMembers 验证 [类型]::成员 的静态属性与静态方法分派。
func TestStaticMembers(t *testing.T) {
	wantStr(t, "[math]::Sqrt(4)", "2")
	wantStr(t, "[math]::Floor(1.9)", "1")
	wantStr(t, "[math]::Ceiling(1.1)", "2")
	wantStr(t, "[math]::Abs(-3)", "3")
	wantStr(t, "[math]::Round(2.5)", "2")
	wantStr(t, "[math]::Pow(2, 10)", "1024")
	wantStr(t, "[math]::Max(1, 2)", "2")
	wantStr(t, "[math]::Min(1, 2)", "1")
	wantStr(t, `[string]::IsNullOrEmpty("")`, "True")
	wantStr(t, `[string]::IsNullOrEmpty("a")`, "False")
	wantStr(t, `[string]::IsNullOrWhiteSpace("  ")`, "True")
	wantStr(t, `[string]::Join("-", 1, 2)`, "1-2")
	wantStr(t, `$arr = 1,2; [string]::Join(",", $arr)`, "1,2")
	wantStr(t, `[string]::Concat("a", 1, "b")`, "a1b")
	wantStr(t, `[string]::Format("{0}+{1}", 1, 2)`, "1+2")
	wantStr(t, `[datetime]::Now.Year -gt 2000`, "True")
	wantStr(t, "$g = [guid]::NewGuid(); $g.ToString().Length", "36")
	// 未注册成员报非终止错误，后续语句继续执行
	wantStr(t, "[math]::NoSuch(1); \"继续\"", "继续")
	// 未注册成员报错且后续语句继续执行
}

// TestTypeLiteralValue 类型字面量本身求值为类型名。
func TestTypeLiteralValue(t *testing.T) {
	wantStr(t, "[int]", "int")
	wantStr(t, "$t = [datetime]; $t", "datetime")
}

// TestPSCustomObjectLiteral 验证 [pscustomobject]@{...} 构造自定义对象（条目变属性）。
func TestPSCustomObjectLiteral(t *testing.T) {
	wantStr(t, `$p = [pscustomobject]@{a = 1; b = "x"}; $p.a`, "1")
	wantStr(t, `$p = [pscustomobject]@{a = 1; b = "x"}; $p.b`, "x")
	wantStr(t, `$p = [pscustomobject]@{n = 5}; $p.n`, "5")
}

// TestNewObject 验证 New-Object PSObject/PSCustomObject 与 -Property 哈希表。
func TestNewObject(t *testing.T) {
	wantStr(t, `$o = New-Object PSObject -Property @{a = 1; b = "x"}; $o.a`, "1")
	wantStr(t, `$o = New-Object PSObject -Property @{a = 1; b = "x"}; $o.b`, "x")
	wantStr(t, `$o = New-Object pscustomobject -Property @{k = "v"}; $o.k`, "v")
}

// TestTestPathType 验证 Test-Path -PathType 按类型过滤（Leaf 文件 / Container 目录）。
func TestTestPathType(t *testing.T) {
	// 用 /etc 与 /etc/hostname，仅 Linux 存在这些路径
	if runtime.GOOS != "linux" {
		t.Skip("跳过：测试依赖 Linux 专属路径 /etc、/etc/hostname")
	}
	wantStr(t, `Test-Path /etc/hostname -PathType Leaf`, "True")
	wantStr(t, `Test-Path /etc/hostname -PathType Container`, "False")
	wantStr(t, `Test-Path /etc -PathType Container`, "True")
	wantStr(t, `Test-Path /etc -PathType Leaf`, "False")
}

// TestGetMemberMemberType 验证 Get-Member -MemberType 过滤成员类型。
func TestGetMemberMemberType(t *testing.T) {
	// 用 /etc/hostname，仅 Linux 存在该路径
	if runtime.GOOS != "linux" {
		t.Skip("跳过：测试依赖 Linux 专属路径 /etc/hostname")
	}
	wantStr(t, `(Get-Item /etc/hostname | Get-Member -MemberType TypeName).Name`, "System.IO.FileInfo")
	wantStr(t, `$gm = Get-Item /etc/hostname | Get-Member -MemberType Property; ($gm.Count -gt 0) -and ($gm[0].MemberType -eq "Property")`, "True")
}

// TestJsonDepth 验证 ConvertTo-Json -Depth 截断嵌套展开。
func TestJsonDepth(t *testing.T) {
	wantStr(t, `@{ a = @{ b = 1 } } | ConvertTo-Json -Depth 1`, "{\n  \"a\": {}\n}")
	wantStr(t, `@{ a = @{ b = 1 } } | ConvertTo-Json -Depth 2`, "{\n  \"a\": {\n    \"b\": 1\n  }\n}")
}

// TestEncodingParams 验证 -Encoding 参数生效（BOM 与 ascii 替换字节数）。
func TestEncodingParams(t *testing.T) {
	// 用 /tmp 路径与字节计数，Windows 的 /tmp 映射会改变字节数
	if runtime.GOOS != "linux" {
		t.Skip("跳过：测试依赖 Linux 的 /tmp 路径与字节计数")
	}
	wantStr(t, `Set-Content /tmp/psl-e1.txt "hi" -Encoding utf8BOM; (Get-Item /tmp/psl-e1.txt).Length`, "6")
	wantStr(t, `Set-Content /tmp/psl-e2.txt "héllo" -Encoding ascii; (Get-Item /tmp/psl-e2.txt).Length`, "6")
	wantStr(t, `"x" | Out-File /tmp/psl-e3.txt -Encoding unicode; (Get-Item /tmp/psl-e3.txt).Length`, "6")
}

// TestSelectStringCase 验证 Select-String 默认大小写不敏感，-CaseSensitive 才敏感。
func TestSelectStringCase(t *testing.T) {
	// 默认不敏感：匹配大小写变体
	wantStr(t, `"Hello" | Select-String "hello" | ForEach-Object { $_.Line }`, "Hello")
	wantStr(t, `"HELLO" | Select-String "hello" | ForEach-Object { $_.Line }`, "HELLO")
	// -CaseSensitive：仅精确大小写匹配
	wantStr(t, `"hello" | Select-String "hello" -CaseSensitive | ForEach-Object { $_.Line }`, "hello")
	wantStr(t, `"Hello" | Select-String "hello" -CaseSensitive | ForEach-Object { $_.Line }`)
	// SimpleMatch 同样默认不敏感
	wantStr(t, `"HELLO" | Select-String "hello" -SimpleMatch | ForEach-Object { $_.Line }`, "HELLO")
}

// TestSelectStringLineNumber 验证 LineNumber 编号：管道输入按对象序号（未匹配也占号），文件输入逐行且空行计入。
func TestSelectStringLineNumber(t *testing.T) {
	// 管道单行对象：LineNumber 是对象在流中的序号
	wantStr(t, `"a","b","c" | Select-String "." | ForEach-Object { "$($_.LineNumber):$($_.Line)" }`, "1:a", "2:b", "3:c")
	// 未匹配的对象也占号
	wantStr(t, `"x","b","c" | Select-String "[bc]" | ForEach-Object { "$($_.LineNumber):$($_.Line)" }`, "2:b", "3:c")
	// 多行字符串整体作为一个匹配单位
	wantStr(t, "(\"line1`nlineX`nline3\" | Select-String \"lineX\").Count", "1")
}

// TestGroupObjectCase 验证 Group-Object 默认大小写不敏感合并，-CaseSensitive 才分组。
func TestGroupObjectCase(t *testing.T) {
	// 默认不敏感：apple/Apple/APPLE 合并为一组，Name 取首次原值
	wantStr(t, `"apple","Apple","APPLE" | Group-Object | ForEach-Object { $_.Name + ":" + $_.Count }`, "apple:3")
	// -CaseSensitive：按原值分组
	wantStr(t, `("apple","Apple","APPLE" | Group-Object -CaseSensitive | Measure-Object).Count`, "3")
}

// TestCompareObject 验证 Compare-Object 默认大小写不敏感、输出先右后左、IncludeEqual。
func TestCompareObject(t *testing.T) {
	// 默认大小写不敏感：B/b、c/C 视为相等，仅输出各自独有项，先右(=>)后左(<=)
	wantStr(t, `Compare-Object -ReferenceObject "a","B","c" -DifferenceObject "b","C","d" | ForEach-Object { $_.SideIndicator + $_.InputObject }`, "=>d", "<=a")
	// IncludeEqual：相等项(==)最先，然后右(=>)再左(<=)
	wantStr(t, `Compare-Object -ReferenceObject "a","b" -DifferenceObject "b","c" -IncludeEqual | ForEach-Object { $_.SideIndicator + $_.InputObject }`, "==b", "=>c", "<=a")
	// IncludeEqual 相等项显示参考集(ref)的值，非差集
	wantStr(t, `Compare-Object -ReferenceObject "A" -DifferenceObject "a" -IncludeEqual | ForEach-Object { $_.SideIndicator + $_.InputObject }`, "==A")
	// -CaseSensitive：a 与 A 不等
	wantStr(t, `Compare-Object -ReferenceObject "a","A" -DifferenceObject "a" -CaseSensitive | ForEach-Object { $_.SideIndicator + $_.InputObject }`, "<=A")
}

// TestSortObjectUniqueCase 验证 Sort-Object -Unique 默认折叠大小写去重，-CaseSensitive 才分。
func TestSortObjectUniqueCase(t *testing.T) {
	// 默认不敏感：apple/Apple/APPLE 折叠为 apple，排序后去重
	wantStr(t, `"apple","Apple","APPLE","banana" | Sort-Object -Unique | ForEach-Object { $_ }`, "apple", "banana")
	// -CaseSensitive：保留各大小写变体
	wantStr(t, `("apple","Apple","APPLE","banana" | Sort-Object -Unique -CaseSensitive | Measure-Object).Count`, "4")
}

// TestSelectObjectFirstLastZero 验证 Select-Object -First/-Last 显式 0 返回空。
func TestSelectObjectFirstLastZero(t *testing.T) {
	// -First 0 返回空（Count 为 0）
	wantStr(t, `("1","2","3" | Select-Object -First 0 | Measure-Object).Count`, "0")
	// -Last 0 返回空
	wantStr(t, `("1","2","3" | Select-Object -Last 0 | Measure-Object).Count`, "0")
	// -First 1 正常取首项
	wantStr(t, `"1","2","3" | Select-Object -First 1`, "1")
}

// TestMeasureObjectFields 验证 Measure-Object 字段总是补全，未开统计为 $null。
func TestMeasureObjectFields(t *testing.T) {
	// 未开开关：Count 有值，Sum/Average 等为空
	wantStr(t, `"1","2","3" | Measure-Object | ForEach-Object { $_.Count }`, "3")
	wantStr(t, `"1","2","3" | Measure-Object | ForEach-Object { $_.Sum -eq $null }`, "True")
	wantStr(t, `"1","2","3" | Measure-Object | ForEach-Object { $_.Average -eq $null }`, "True")
	// 开 Sum 有数字时 Sum 有值
	wantStr(t, `"1","2","3" | Measure-Object -Sum | ForEach-Object { $_.Sum }`, "6")
	// 开 Sum 遇非数字：累加统计作废，Sum 为空（对齐真 PowerShell）
	wantStr(t, `"a","b" | Measure-Object -Sum | ForEach-Object { $_.Sum -eq $null }`, "True")
	// 混合输入(含数字与非数字)开 Sum：仍作废
	wantStr(t, `"1","a","2" | Measure-Object -Sum | ForEach-Object { $_.Sum -eq $null }`, "True")
	// 混合输入开 Average：作废
	wantStr(t, `"2","a" | Measure-Object -Average | ForEach-Object { $_.Average -eq $null }`, "True")
	// -Property 模式：Count 只数有该属性的对象
	wantStr(t, `@{a=1},@{a=2},@{b=3} | Measure-Object -Property a -Sum | ForEach-Object { $_.Count }`, "2")
	wantStr(t, `@{a=1},@{a=2},@{b=3} | Measure-Object -Property a -Sum | ForEach-Object { $_.Sum }`, "3")
}

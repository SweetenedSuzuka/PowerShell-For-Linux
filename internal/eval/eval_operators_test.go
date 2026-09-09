package eval

import (
	"testing"
)

func TestArithmetic(t *testing.T) {
	wantStr(t, "1 + 2", "3")
	wantStr(t, "10 - 4", "6")
	wantStr(t, "6 * 7", "42")
	wantStr(t, "7 / 2", "3.5")
	wantStr(t, "10 % 3", "1")
	wantStr(t, "2 * 4", "8")
	// 范围运算符绑定比比较运算更紧（与 PowerShell 一致）：1..3 -gt 1 先成范围再过滤
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

// TestFloatAddition 验证浮点加法不截断：整型路径按 TypeName 识别，任一操作数是浮点就按浮点运算（2 + 1/2 = 2.5 而非 2）。
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
	// 复合赋值 += 浮点（调用 addOp）
	wantStr(t, "$i = 0.5; $i += 1; $i", "1.5")
}

// TestDivideByZero 验证除零报错（与 PowerShell 一致）：
// 结果置为 $null、$? 置为 false，REPL 语义下后续语句继续执行；正常除法与取模不受影响。
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

// TestStringNumberComparison 验证字符串与数字混合比较按左操作数类型转换（与 PowerShell 一致）：两个字符串按字典序（"5" -lt "10" 为 False），数字对字符串按数字（5 -lt "10" 为 True），$null 只与 $null 相等。
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
	// 范围字面量与成员/比较运算符连用（.. 绑定比比较更紧）
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

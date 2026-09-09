package eval

import (
	"bytes"
	"io"
	"os"
	"powershell/internal/object"
	"powershell/internal/parser"
	"powershell/internal/shell"
	"runtime"
	"strings"
	"testing"
)

func TestPipeline(t *testing.T) {
	wantStr(t, "1..5 | Where-Object { $_ % 2 -eq 0 }", "2", "4")
	wantStr(t, "3,1,2 | Sort-Object", "1", "2", "3")
	wantStr(t, "1..5 | Select-Object -First 2", "1", "2")
	wantStr(t, "1..5 | Measure-Object -Sum | ForEach-Object { $_.Sum }", "15")
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
	// 文件对象按表格渲染
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

// TestTrimNoArgs 验证 TrimStart/TrimEnd 无参数清空白（有参数按字符集裁剪行为不变）。
func TestTrimNoArgs(t *testing.T) {
	wantStr(t, `"  x  ".TrimStart()`, "x  ")
	wantStr(t, `"  x  ".TrimEnd()`, "  x")
	wantStr(t, `"  x  ".Trim()`, "x")
	wantStr(t, `"007".TrimStart("0")`, "7")
	wantStr(t, `"007".TrimEnd("7")`, "00")
}

// TestSplitNoArgs 验证 Split 无参数按任意空白分割（含 tab/换行，连续空白合并），有参数行为不变。
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
	// 文件时间纪元 1601-01-01 UTC 的刻度为 0，远古日期刻度正常
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

// TestCopyItemPerm 验证 Copy-Item 保留源文件权限位。
func TestCopyItemPerm(t *testing.T) {
	// 权限位仅 Linux 有效；源文件 755，目标须同样 755 而非固定值
	if runtime.GOOS != "linux" {
		t.Skip("跳过：测试依赖 Linux 文件权限位")
	}
	src := "/tmp/psl-cpp-src.sh"
	dst := "/tmp/psl-cpp-dst.sh"
	if err := os.WriteFile(src, []byte("#!/bin/sh\n"), 0o755); err != nil {
		t.Fatalf("写源文件失败：%v", err)
	}
	defer os.Remove(src)
	defer os.Remove(dst)
	runEval(t, `Copy-Item /tmp/psl-cpp-src.sh /tmp/psl-cpp-dst.sh`)
	info, err := os.Stat(dst)
	if err != nil {
		t.Fatalf("目标文件不存在：%v", err)
	}
	if info.Mode().Perm() != 0o755 {
		t.Errorf("目标权限位 %o，想要 755", info.Mode().Perm())
	}
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

// TestSelectStringQuiet 验证 -Quiet：命中输出单个 $true，未命中无输出。
func TestSelectStringQuiet(t *testing.T) {
	// 命中输出单个 True
	wantStr(t, `"aXb" | Select-String -Pattern "X" -Quiet`, "True")
	// 多条命中仍只输出一个
	wantStr(t, `("aXb","cXd" | Select-String -Pattern "X" -Quiet).Count`, "1")
	// 未命中无输出
	wantStr(t, `"abc" | Select-String -Pattern "Z" -Quiet`)
	// 未命中结果判空为真
	wantStr(t, `if ("abc" | Select-String -Pattern "Z" -Quiet) { "hit" } else { "miss" }`, "miss")
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

// TestSortObjectUniqueCase 验证 Sort-Object -Unique 默认不区分大小写去重，-CaseSensitive 才分。
func TestSortObjectUniqueCase(t *testing.T) {
	// 默认不敏感：apple/Apple/APPLE 都视为小写 apple，排序后去重
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

// TestSelectObjectSkip 验证 -Skip：从头部扣除、与 -First/-Last 组合、超界置为空、负数报错。
func TestSelectObjectSkip(t *testing.T) {
	// 跳过前 3 条
	wantStr(t, `1..10 | Select-Object -Skip 3`, "4", "5", "6", "7", "8", "9", "10")
	// 先跳过再取前 3 条
	wantStr(t, `1..10 | Select-Object -Skip 2 -First 3`, "3", "4", "5")
	// -Last 在时从尾部扣除
	wantStr(t, `1..10 | Select-Object -Skip 2 -Last 3`, "6", "7", "8")
	// 跳过 0 条与超界
	wantStr(t, `(1..10 | Select-Object -Skip 0 | Measure-Object).Count`, "10")
	wantStr(t, `(1..10 | Select-Object -Skip 20 | Measure-Object).Count`, "0")
	// 负数报错并记录
	wantStr(t, `1..5 | Select-Object -Skip -1; $Error.Count`, "1")
	wantStr(t, `1..5 | Select-Object -Skip -1; if ($?) { "ok" } else { "fail" }`, "fail")
}

// TestSelectObjectUniqueLast 验证 -Unique 在投影/展开之后去重。
func TestSelectObjectUniqueLast(t *testing.T) {
	// 按投影后的属性值去重
	wantStr(t, `([pscustomobject]@{n="a";m=1},[pscustomobject]@{n="a";m=2} | Select-Object -Property n -Unique | Measure-Object).Count`, "1")
	// 按展开后的值去重
	wantStr(t, `([pscustomobject]@{n="a"},[pscustomobject]@{n="a"},[pscustomobject]@{n="b"} | Select-Object -ExpandProperty n -Unique)`, "a", "b")
	// -First 之后去重不变
	wantStr(t, `(1,1,2,2,3 | Select-Object -First 2 -Unique | Measure-Object).Count`, "1")
}

// TestMeasureObjectFields 验证 Measure-Object 字段总是补全，未指定统计为 $null。
func TestMeasureObjectFields(t *testing.T) {
	// 未指定开关：Count 有值，Sum/Average 等为空
	wantStr(t, `"1","2","3" | Measure-Object | ForEach-Object { $_.Count }`, "3")
	wantStr(t, `"1","2","3" | Measure-Object | ForEach-Object { $_.Sum -eq $null }`, "True")
	wantStr(t, `"1","2","3" | Measure-Object | ForEach-Object { $_.Average -eq $null }`, "True")
	// 指定 -Sum 有数字时 Sum 有值
	wantStr(t, `"1","2","3" | Measure-Object -Sum | ForEach-Object { $_.Sum }`, "6")
	// 指定 -Sum 遇非数字：累加统计作废，Sum 为空（与原版 PowerShell 一致）
	wantStr(t, `"a","b" | Measure-Object -Sum | ForEach-Object { $_.Sum -eq $null }`, "True")
	// 混合输入(含数字与非数字)指定 -Sum：仍作废
	wantStr(t, `"1","a","2" | Measure-Object -Sum | ForEach-Object { $_.Sum -eq $null }`, "True")
	// 混合输入指定 -Average：作废
	wantStr(t, `"2","a" | Measure-Object -Average | ForEach-Object { $_.Average -eq $null }`, "True")
	// -Property 模式：Count 只数有该属性的对象
	wantStr(t, `@{a=1},@{a=2},@{b=3} | Measure-Object -Property a -Sum | ForEach-Object { $_.Count }`, "2")
	wantStr(t, `@{a=1},@{a=2},@{b=3} | Measure-Object -Property a -Sum | ForEach-Object { $_.Sum }`, "3")
}

# 类型字面量。
Write-Output "== 类型字面量 =="

# 118. 强制转换基础形态
$results += T "类型转换 int/string/double" (([int]"42" -eq 42) -and ([string]42 -eq "42") -and ([double]"1.5" -eq 1.5))
# 119. datetime 转换取字段
$d = [datetime]"2020-01-02"
$results += T "datetime 转换取 Year" (($d.Year -eq 2020))
# 120. -is/-as 直接接类型字面量
$results += T "-is/-as 接类型字面量" ((1 -is [int]) -and ("a" -isnot [int]) -and ("7" -as [int] -eq 7))
# 121. 转换失败报错且后续语句继续执行
$rTl = ""
[int]"abc"
if ($?) { $rTl = "no" } else { $rTl = "yes" }
Set-Content test/tmp/tl-after.txt "ran"
$results += T "转换失败报错且继续执行" (($rTl -eq "yes") -and ((Get-Content test/tmp/tl-after.txt) -eq "ran"))
Remove-Item test/tmp/tl-after.txt -Force 2>$null

# 122. 形参类型标注：位置实参转换成声明类型
function Fpt1([int]$x) { $x -is [int] }
$results += T "形参类型标注位置实参" ((Fpt1 "42") -eq $true)
# 123. 默认值也经类型转换
function Fpt2([int]$y = "5") { $y }
$results += T "形参默认值转换" ((Fpt2) -eq 5)
# 124. 数组标注把单值包成单元素数组
function Fpt3([int[]]$a) { $a.Count }
$results += T "数组标注单值包装" ((Fpt3 7) -eq 1)
# 125. 实参无法转换时不执行函数体，问号变量（$?）置为 false，并在后续语句继续
$rPf = ""
function Fpt4([int]$x) { Set-Content test/tmp/pb-body.txt "ran" }
Fpt4 "abc"
if ($?) { $rPf = "no" } else { $rPf = "yes" }
Set-Content test/tmp/pb-after.txt "ran"
$pbOk = (($rPf -eq "yes") -and ((Get-Content test/tmp/pb-after.txt) -eq "ran") -and (-not (Test-Path test/tmp/pb-body.txt)))
$results += T "形参绑定失败不执行函数体" ($pbOk)
Remove-Item test/tmp/pb-body.txt -Force 2>$null
Remove-Item test/tmp/pb-after.txt -Force 2>$null

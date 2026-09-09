# 虚拟属性、时间种类与直接输入。
# 215. Select * 含虚拟属性
$sa = Get-Date | Select-Object *
$results += T "星号选择" ((($sa.Year -gt 0)) -and (($sa.Ticks -gt 0)))
# 216. 过大、过小年份刻度
$tkMin = ([datetime]"0001-01-01").Ticks
$tkMax = ([datetime]"9999-12-31 23:59:59.9999999").Ticks
$results += T "极端年份刻度" ((($tkMin -eq 0)) -and (($tkMax -eq 3155378975999999999)))
# 217. 时间种类
$kdZone = ([datetime]"2020-01-01T00:00:00+08:00").Kind
$kdBare = ([datetime]"2020-01-01 12:00:00").Kind
$kdArg = (Get-Date "2020-01-01").Kind
$results += T "时间种类" ((($kdZone -eq "Local")) -and (($kdBare -eq "Unspecified")) -and (($kdArg -eq "Unspecified")))
# 218. InputObject 直接输入
$fi = ForEach-Object -InputObject "hi" -Process { $_ + "!" }
$fa = ForEach-Object -InputObject 1,2,3 -Process { $_.Count }
$fm = ForEach-Object -InputObject "m" -MemberName Length
$results += T "InputObject 直接输入" (((($fi -join ",") -eq "hi!")) -and ((($fa -join ",") -eq "3")) -and ((($fm -join ",") -eq "1")))
# 219. InputObject 与管道并存报错
$fe0 = $Error.Count
"pipe" | ForEach-Object -InputObject "arg" -Process { $_ } 2>$null | Out-Null
$results += T "InputObject 管道并存报错" (($Error.Count -eq ($fe0 + 1)))

# 选择投影：跳过、去重与静默。
# 163. Select-Object -Skip 跳过前 N 条
$sk1 = 1..10 | Select-Object -Skip 3
$results += T "Skip 跳过" ((($sk1 -join ",") -eq "4,5,6,7,8,9,10"))
# 164. -Skip 先扣除再取 -First
$sk2 = 1..10 | Select-Object -Skip 2 -First 3
$results += T "Skip 加 First" ((($sk2 -join ",") -eq "3,4,5"))
# 165. -Last 在时 -Skip 从尾部扣除
$sk3 = 1..10 | Select-Object -Skip 2 -Last 3
$results += T "Skip 加 Last 从尾扣除" ((($sk3 -join ",") -eq "6,7,8"))
# 166. -Skip 负数报错
$sk4 = $Error.Count
1..5 | Select-Object -Skip -1
if ($?) { $sk4ok = "ok" } else { $sk4ok = "fail" }
$results += T "Skip 负数报错" (($Error.Count -eq $sk4 + 1) -and ($sk4ok -eq "fail"))
# 167. -Unique 在投影之后去重
$sk5 = [pscustomobject]@{n="a";m=1},[pscustomobject]@{n="a";m=2} | Select-Object -Property n -Unique
$results += T "Unique 投影后去重" (($sk5.Count -eq 1) -and ($sk5[0].n -eq "a"))
# 168. Select-String -Quiet 命中输出单个 True
$ss1 = "aXb","cXd" | Select-String -Pattern "X" -Quiet
$results += T "Quiet 命中" (($ss1 -eq $true) -and (@($ss1).Count -eq 1))
# 169. Select-String -Quiet 未命中无输出
$ss2 = "abc" | Select-String -Pattern "Z" -Quiet
$results += T "Quiet 未命中无输出" ($ss2 -eq $null)
# 170. 自定义对象流合成一张表（首对象属性作列，后出属性不另起表）
[pscustomobject]@{n="x";m=1},[pscustomobject]@{n="y";z=2} | Out-File ft-merge.txt
$ftm = Get-Content ft-merge.txt
$results += T "自定义对象合成表" (($ftm.Count -eq 4) -and ($ftm[0] -like "n*") -and ((($ftm -join "")) -notlike "*z*"))

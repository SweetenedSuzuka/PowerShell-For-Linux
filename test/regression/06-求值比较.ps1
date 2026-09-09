# 求值与比较。
Write-Output "== 求值与比较 =="

# 70. 范围索引 $a[1..2]（多下标逐元素取值，负数从末尾数）
$idxArr = 1,2,3,4
$idx = $idxArr[1..2]
$results += T "范围索引" (($idx -join ",") -eq "2,3")
# 71. 字符串-字符串比较按字典序（"5" -lt "10" 应为 False）
$results += T "字符串比较 -lt" ((-not ("5" -lt "10")))
# 72. $i++ 浮点不截断（0.5 → 1.5）
$fi = 0.5
$fi++
$results += T '$i++ 浮点' (($fi -eq 1.5))
# 73. -split 最大子串参数（末段保留未分割剩余）
$sp = "a,b,c" -split ",",2
$results += T "-split 最大子串" (($sp -join "|") -eq "a|b,c")
# 74. 除零报错（结果 $null、$? 置为 false）
$zr = 5/0
if (-not $?) { $zd = "err" } else { $zd = "ok" }
$results += T "除零报错" (($zr -eq $null) -and ($zd -eq "err"))
# 75. 布尔-数字顺序比较（$true=1、$false=0 参与数字比较）
$results += T "布尔比较 -lt" (($true -lt 2) -and (-not ($true -gt 1)) -and ($false -lt 1))

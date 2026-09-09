# 语言结构：作用域、迭代、参数、异常。
Write-Output "== 语言结构（作用域/迭代/参数/异常） =="

# 62. try/catch 捕获
$tc = "未执行"
try { throw "boom" } catch { $tc = "已捕获" }
$results += T "try/catch 捕获" (($tc -eq "已捕获"))
# 63. param() 块（函数）
function Fp { param($x) $x * 2 }
$pv = $null
$pv = Fp 21 2>$null
$results += T "param() 块" (($pv -eq 42))
# 64. switch 对数组迭代
$swr = switch (1,2,3) { default { $_ } }
$results += T "switch 数组迭代" (($swr -join ",") -eq "1,2,3")
# 65. $script: 作用域写回（函数内修改脚本作用域变量）
$sf = 0
function SetSF { $script:sf = 5 }
SetSF
$results += T '$script: 作用域写回' (($sf -eq 5))

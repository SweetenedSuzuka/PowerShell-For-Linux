# begin、process、end 与 filter。
Write-Output "== begin/process/end 与 filter =="

# 131. 三块齐全：begin 一次 → process 逐项 → end 一次
function Nb1 { begin { "nb-b" } process { "nb-p:$_" } end { "nb-e" } }
$nb = 1,2 | Nb1
$results += T "begin/process/end 三段顺序" (($nb -join ",") -eq "nb-b,nb-p:1,nb-p:2,nb-e")
# 132. 直接调用：process 以 $null 跑一次
$nb2 = Nb1
$results += T "函数直调三段" (($nb2 -join ",") -eq "nb-b,nb-p:,nb-e")
# 133. 零输入：begin/end 跑，process 不跑
$zero = @()
$nb3 = @($zero | Nb1)
$results += T "零输入跳过 process" (($nb3.Count -eq 2) -and ($nb3[0] -eq "nb-b") -and ($nb3[1] -eq "nb-e"))
# 134. filter 的 Body 即 process，逐项绑定 $_
filter Nbf { "nf:$_" }
$nb4 = 1,2 | Nbf
$results += T "filter 逐项绑定" (($nb4 -join ",") -eq "nf:1,nf:2")
# 135. process 内 return 只结束本次
function Nb2 { process { if ($_ -eq 2) { return }; "nr:$_" } }
$nb5 = 1,2,3 | Nb2
$results += T "process 内 return 继续" (($nb5 -join ",") -eq "nr:1,nr:3")
# 136. param 形参在 process 内可用
function Nb3([int]$m) { process { $_ * $m } }
$nb6 = 1,2,3 | Nb3 -m 10
$results += T "param 与 process 共存" (($nb6 -join ",") -eq "10,20,30")

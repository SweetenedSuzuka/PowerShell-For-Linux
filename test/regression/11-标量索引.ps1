# 标量索引。
Write-Output "== 标量索引 =="

# 103. 标量变量 [0] 返回自身，非零下标返回 $null
$scalar = [pscustomobject]@{ tag = "t" }
$results += T "标量索引自身" (($scalar[0].tag -eq "t") -and ($scalar[5] -eq $null))
# 104. GroupInfo 不被当集合：[0] 是组对象本身，成员在 Group 属性
$gi = "x","y","x" | Group-Object
$results += T "GroupInfo 索引取到组本身" ((@($gi).Count -eq 2) -and ($gi[0].Name -eq "x") -and ($gi[0].Group -join ",") -eq "x,x")
# 105. Join-String 默认分隔符是空串（与 PowerShell 一致）
$results += T "Join-String 默认分隔符" (((1..3 | Join-String) -eq "123") -and ((1..3 | Join-String -Separator ",") -eq "1,2,3"))
# 106. @() 元素按输出流展开：@($arr) 得元素本身；@() 内可含管道
$aa = 1,2,3
$results += T "@() 展开与管道" ((@($aa).Count -eq 3) -and (@(1,2 | ForEach-Object { $_ * 10 }) -join ",") -eq "10,20")

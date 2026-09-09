# 数组路径与展开属性。
Write-Output "== 数组路径与 ExpandProperty（fab7df6 修复回归） =="

# 37. Get-Content 数组路径
"g1" | Set-Content g1.txt; "g2" | Set-Content g2.txt
$gc2 = Get-Content g1.txt, g2.txt
$results += T "Get-Content 数组路径" (($gc2 -join ",") -eq "g1,g2")
# 38. Get-Item 数组路径
$gi = Get-Item g1.txt, g2.txt
$results += T "Get-Item 数组路径" (($gi.Count -eq 2))
# 39. 数组 + TotalCount/Tail 按文件生效
"l1" | Set-Content m1.txt; "l2" | Add-Content m1.txt; "l3" | Add-Content m1.txt
"k1" | Set-Content m2.txt; "k2" | Add-Content m2.txt
$gt = Get-Content m1.txt, m2.txt -Tail 1
$results += T "多文件 -Tail 按文件" (($gt -join ",") -eq "l3,k2")
$gto = Get-Content m1.txt, m2.txt -TotalCount 1
$results += T "多文件 -TotalCount 按文件" (($gto -join ",") -eq "l1,k1")
# 40. Get-Content 多位置实参
$g3 = Get-Content g1.txt g2.txt
$results += T "Get-Content 多位置" (($g3 -join ",") -eq "g1,g2")
# 41. Get-Item 通配数组（g?.txt 避开本脚本早前创建的 g.txt）
$gw = Get-Item g?.txt, m*.txt
$results += T "Get-Item 通配数组" (($gw.Count -eq 4))
# 42. Get-Item 单路径不受影响
$gs = Get-Item g1.txt
$results += T "Get-Item 单路径" (($gs.Name -eq "g1.txt"))
# 43. Test-Path 数组路径
$results += T "Test-Path 数组" (((Test-Path g1.txt, m2.txt) -eq $true))
# 44. -ExpandProperty 标量
$h = @{ Name = "hn"; Length = 5 }
$ep = $h | Select-Object -ExpandProperty Name
$results += T "ExpandProperty 标量" (($ep -eq "hn"))
# 45. -ExpandProperty 数组展开
$arr = @{ Items = @(1,2,3) }
$epa = $arr | Select-Object -ExpandProperty Items
$results += T "ExpandProperty 数组展开" (($epa -join ",") -eq "1,2,3")
# 46. -Property 与 -ExpandProperty 并存时 Property 优先
$both = $h | Select-Object -Property Name -ExpandProperty Name
$results += T "Property 优先于 ExpandProperty" (($both.Name -eq "hn"))
# 47. 脚本块槽位：命名优先
$sw = 1,2,3 | Where-Object -FilterScript { $_ -gt 1 } { $_ -eq 2 }
$results += T "脚本块命名优先" (($sw -join ",") -eq "2,3")

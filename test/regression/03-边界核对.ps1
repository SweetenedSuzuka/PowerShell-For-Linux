# 边界核对：多源与槽位映射。
Write-Output "== 边界核对（Copy-Item 多源 / 槽位映射） =="

# 48. Copy-Item 四位置实参
"c1" | Set-Content a1.txt; "c2" | Set-Content a2.txt; "c3" | Set-Content a3.txt
New-Item -ItemType Directory -Path d4 | Out-Null
Copy-Item a1.txt a2.txt a3.txt d4
$results += T "Copy-Item 四位置" (((Get-ChildItem d4 -Name) -join ",") -eq "a1.txt,a2.txt,a3.txt")
# 49. Move-Item 三位置（用独立文件）
"m1" | Set-Content mv1.txt; "m2" | Set-Content mv2.txt
New-Item -ItemType Directory -Path d5 | Out-Null
Move-Item mv1.txt mv2.txt d5
$results += T "Move-Item 三位置" (((Get-ChildItem d5 -Name) -join ",") -eq "mv1.txt,mv2.txt")
# 50. Copy-Item 命名目标在前
New-Item -ItemType Directory -Path d6 | Out-Null
Copy-Item -Destination d6 a3.txt
$results += T "命名目标在前" (((Get-ChildItem d6 -Name) -join ",") -eq "a3.txt")
# 51. Copy-Item 数组 + 位置目标
New-Item -ItemType Directory -Path d7 | Out-Null
Copy-Item a1.txt,a2.txt d7 2>$null
$results += T "数组+位置目标" (((Get-ChildItem d7 -Name) -join ",") -eq "a1.txt,a2.txt")
# 52. Set-Content 管道 + 命名 Path 缺值（命名 Path 缺值时不写文件）
"x" | Set-Content -Path np.txt
$results += T "管道+命名缺值不写" ((-not (Test-Path np.txt)))
# 53. Add-Content 位置缺值追加空
Add-Content ac.txt
$results += T "Add-Content 位置缺值写空" (((Test-Path ac.txt)))
# 54. Get-ChildItem 通配多路径（多个通配符并列，结果包含目标文件）
"t1" | Set-Content f1.log; "t2" | Set-Content f2.md
$wc = Get-ChildItem *.txt, *.log -Name
$results += T "Get-ChildItem 通配数组" (($wc -contains "a1.txt") -and ($wc -contains "f1.log") -and ($wc.Count -ge 4))
# 55. Get-ChildItem 多目录
New-Item -ItemType Directory -Path dirA | Out-Null
New-Item -ItemType Directory -Path dirB | Out-Null
"x" | Set-Content dirA/fa.txt; "y" | Set-Content dirB/fb.txt
$gc = Get-ChildItem dirA dirB -Name
$results += T "Get-ChildItem 多目录" (($gc -join ",") -eq "fa.txt,fb.txt")
# 56. Select-Object -First 组合
$sf = Select-Object -First 2 a b c
$results += T "Select-Object -First 组合" (($sf -join ",") -eq "a,b")
# 57. Get-Unique 数组 + 位置
$gu2 = Get-Unique 3,1 3 2
$results += T "Get-Unique 数组+位置" (($gu2 -join ",") -eq "3,1,2")
# 58. Compare-Object 命名 + 位置映射
$co = Compare-Object -DifferenceObject b a
$results += T "Compare-Object 位置映射" (($co.Count -eq 2))
# 59. Set-Alias 两位置
Set-Alias al1 cv1 2>$null
$results += T "Set-Alias 两位置" (((Get-Alias al1).Definition -eq "cv1"))
# 60. 管道输入 Set-Content 写文件，再用路径读回来
"p1" | Set-Content p.txt
$gp = Get-Content p.txt
$results += T "Get-Content 单文件" (($gp -join ",") -eq "p1")
# 61. Select-Object 无管道位置当数据（回归）
$so2 = Select-Object x y
$results += T "Select-Object 无管道回归" (($so2 -join ",") -eq "x,y")

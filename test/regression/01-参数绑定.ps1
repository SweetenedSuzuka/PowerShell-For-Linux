# 参数绑定：位置绑定中心化回归。
Write-Output "== 参数绑定（位置绑定中心化回归） =="

# 1. Set-Content 全位置
Set-Content a.txt "val-a"
$results += T "Set-Content 全位置" (((Get-Content a.txt) -eq "val-a"))
# 2. Set-Content 命名 Path + 位置值（命名 Path 占位后，位置值映射到 Value）
Set-Content -Path b.txt bval
$results += T "命名Path+位置值" (((Get-Content b.txt) -eq "bval"))
# 3. Set-Content 全命名
Set-Content -Path c.txt -Value cval
$results += T "Set-Content 全命名绑定" (((Get-Content c.txt) -eq "cval"))
# 4. Set-Content 位置 Path + 命名 Value
Set-Content d.txt -Value dval
$results += T "位置Path+命名Value" (((Get-Content d.txt) -eq "dval"))
# 5. Add-Content 追加
Set-Content e.txt "e1"; Add-Content e.txt "e2"
$results += T "Add-Content" (((Get-Content e.txt) -join ",") -eq "e1,e2")
# 6. Copy-Item 三位置多源（末位目标）
"x1" | Set-Content f1.txt; "x2" | Set-Content f2.txt
New-Item -ItemType Directory -Path dir1 | Out-Null
Copy-Item f1.txt f2.txt dir1
$results += T "Copy-Item 三位置多源" (((Get-ChildItem dir1 -Name) -join ",") -eq "f1.txt,f2.txt")
# 7. Copy-Item 数组源
New-Item -ItemType Directory -Path dir2 | Out-Null
Copy-Item f1.txt,f2.txt dir2
$results += T "Copy-Item 数组源" (((Get-ChildItem dir2 -Name) -join ",") -eq "f1.txt,f2.txt")
# 8. Copy-Item 命名 Path + 位置
New-Item -ItemType Directory -Path dir3 | Out-Null
Copy-Item -Path f1.txt f2.txt dir3
$results += T "Copy-Item 命名Path+位置" (((Get-ChildItem dir3 -Name) -join ",") -eq "f1.txt,f2.txt")
# 9. Get-ChildItem 多路径
$results += T "Get-ChildItem 多路径" (((Get-ChildItem f1.txt, f2.txt -Name) -join ",") -eq "f1.txt,f2.txt")
# 10. Get-Content 行数
"line1" | Set-Content g.txt; "line2" | Add-Content g.txt
$results += T "Get-Content 位置Path" (((Get-Content g.txt).Count -eq 2))
# 11. Get-Content -Tail 命名
$results += T "Get-Content 命名Tail" (((Get-Content g.txt -Tail 1) -join ",") -eq "line2")
# 12. Join-Path 双位置
$results += T "Join-Path" (((Join-Path a.txt b.txt) -eq "a.txt/b.txt"))
# 13. Rename-Item 双位置
Rename-Item a.txt a2.txt
$results += T "Rename-Item" (((Test-Path a2.txt) -and -not (Test-Path a.txt)))
# 14. Compare-Object 双位置
$cmp = Compare-Object a b
$results += T "Compare-Object" (($cmp.Count -eq 2))
# 15. Select-Object 位置属性（管道）
Get-ChildItem f1.txt | Select-Object Name, Length | Out-Null
$results += T "Select-Object 管道属性" (($?))
# 16. Select-Object 无管道位置当数据
$so = Select-Object a b c
$results += T "Select-Object 无管道数据" (($so.Count -eq 3))
# 17. Where-Object 位置脚本块
$wf = 1,2,3,4 | Where-Object { $_ -gt 2 }
$results += T "Where-Object 脚本块" (($wf -join ",") -eq "3,4")
# 18. Where-Object 命名 FilterScript
$wf2 = 1,2,3 | Where-Object -FilterScript { $_ -eq 2 }
$results += T "Where-Object 命名" (($wf2 -join ",") -eq "2")
# 19. ForEach-Object 位置脚本块
$fe = "a","b" | ForEach-Object { $_ + "!" }
$results += T "ForEach-Object" (($fe -join ",") -eq "a!,b!")
# 20. ForEach-Object 命名 Process
$fe2 = 1,2 | ForEach-Object -Process { $_ * 10 }
$results += T "ForEach-Object 命名" (($fe2 -join ",") -eq "10,20")
# 21. Sort-Object 位置
$so2 = 2,1,3 | Sort-Object
$results += T "Sort-Object" (($so2 -join ",") -eq "1,2,3")
# 22. Measure-Command 位置脚本块
$mc = Measure-Command { Start-Sleep -Milliseconds 10 }
$results += T "Measure-Command" (($mc.TotalMilliseconds -gt 0))
# 23. Get-Unique 多位置
$gu = Get-Unique 3 1 3 2 2
$results += T "Get-Unique" (($gu -join ",") -eq "3,1,2")
# 24. Set-Content 命名缺值静默
Set-Content -Path nope.txt
$results += T "命名缺值静默" ((-not (Test-Path nope.txt)))
# 25. 位置缺值写空
Set-Content empty.txt
$results += T "位置缺值写空" (((Get-Item empty.txt).Length -eq 0))
# 26. Select-String 双位置
"needle here" | Set-Content pat.txt
$ss = Select-String needle pat.txt
$results += T "Select-String" (($ss.Count -eq 1))
# 27. Set-ItemProperty 三位置
Set-Content prop.txt "p"
Set-ItemProperty prop.txt LastWriteTime 2020-01-01
$results += T "Set-ItemProperty" (($?))
# 28. 管道 + Set-Content
"pipe-content" | Set-Content piped.txt
$results += T "管道写文件" (((Get-Content piped.txt) -eq "pipe-content"))
# 29. 管道组合
$combo = Get-ChildItem *.txt | Where-Object { $_.Length -gt 0 } | ForEach-Object { $_.Name.Length }
$results += T "文件过滤管道" (($combo.Count -gt 0))
# 30. Get-ItemProperty 双位置
$gip = Get-ItemProperty pat.txt Length
$results += T "Get-ItemProperty" (($gip.Length -gt 0))
# 31. 命名优先占位
$results += T "命名优先占位" (((Join-Path -ChildPath child /tmp) -eq "/tmp/child"))
# 32. Get-Content -Path + -Tail 混合
$gc2 = Get-Content g.txt -Tail 1
$results += T "命名+位置混合" (($gc2 -join ",") -eq "line2")
# 33. Set-PSVersion 位置
Set-PSVersion 5
$results += T "Set-PSVersion 位置" (($PSVersionTable.PSVersion.Major -eq 5))
Set-PSVersion 7
# 34. Get-ChildItem -Filter 命名不受位置干扰
$gf = Get-ChildItem *.txt -Filter "*.txt"
$results += T "Get-ChildItem -Filter" (($gf.Count -gt 0))
# 35. Set-Item 两个位置参数（只改写已存在的项）
"si-old" | Set-Content si.txt
Set-Item si.txt "si-val"
$results += T "Set-Item" (((Get-Content si.txt) -eq "si-val"))
# 36. Get-Command 位置
$gc = Get-Command Set-Content
$results += T "Get-Command 位置" (($gc.Name -eq "Set-Content"))

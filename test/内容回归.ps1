# 内容回归：对照已修复问题做内容级验证。
# 核验.ps1 只查 $? 不查实际内容（位置绑定槽位缺陷正是被它漏掉），本脚本检查命令输出的实际内容，作为正式回归资产。
# 用法：./powershell -NoLogo -NoProfile -File test/内容回归.ps1。
# 可重复运行：脚本在 test/tmp/reg 子目录工作，开头清空重建，结尾保留供排查。

$results = @()
# T 断言：返回 PASS/FAIL 字符串，顶层收集后统一打印与统计。
function T($name, $cond) { if ($cond) { "PASS  $name" } else { "FAIL  $name" } }

# 工作目录：清空重建，保证可重复运行
Remove-Item -Recurse -Force test/tmp/reg 2>$null >$null
New-Item -ItemType Directory -Force test/tmp/reg 2>$null >$null
Set-Location test/tmp/reg

Write-Output "== 参数绑定（位置绑定中心化回归） =="

# 1. Set-Content 全位置
Set-Content a.txt "val-a"
$results += T "Set-Content 全位置" (((Get-Content a.txt) -eq "val-a"))
# 2. Set-Content 命名 Path + 位置值（核心场景，此前被 Encoding 抢走）
Set-Content -Path b.txt bval
$results += T "命名Path+位置值" (((Get-Content b.txt) -eq "bval"))
# 3. Set-Content 全命名
Set-Content -Path c.txt -Value cval
$results += T "全命名" (((Get-Content c.txt) -eq "cval"))
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
$results += T "管道写内容" (((Get-Content piped.txt) -eq "pipe-content"))
# 29. 管道组合
$combo = Get-ChildItem *.txt | Where-Object { $_.Length -gt 0 } | ForEach-Object { $_.Name.Length }
$results += T "管道组合" (($combo.Count -gt 0))
# 30. Get-ItemProperty 双位置
$gip = Get-ItemProperty pat.txt Length
$results += T "Get-ItemProperty" (($gip.Length -gt 0))
# 31. 命名优先跳槽
$results += T "命名优先跳槽" (((Join-Path -ChildPath child /tmp) -eq "/tmp/child"))
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
# 35. Set-Item 位置双参
Set-Item si.txt "si-val"
$results += T "Set-Item" (((Get-Content si.txt) -eq "si-val"))
# 36. Get-Command 位置
$gc = Get-Command Set-Content
$results += T "Get-Command 位置" (($gc.Name -eq "Set-Content"))

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
# 45. -ExpandProperty 数组摊平
$arr = @{ Items = @(1,2,3) }
$epa = $arr | Select-Object -ExpandProperty Items
$results += T "ExpandProperty 数组摊平" (($epa -join ",") -eq "1,2,3")
# 46. -Property 与 -ExpandProperty 并存时 Property 优先
$both = $h | Select-Object -Property Name -ExpandProperty Name
$results += T "Property 优先于 ExpandProperty" (($both.Name -eq "hn"))
# 47. 脚本块槽位：命名优先
$sw = 1,2,3 | Where-Object -FilterScript { $_ -gt 1 } { $_ -eq 2 }
$results += T "脚本块命名优先" (($sw -join ",") -eq "2,3")

Write-Output "== 边界核对（Copy-Item 多源 / 槽位跳槽） =="

# 48. Copy-Item 四位置实参
"c1" | Set-Content a1.txt; "c2" | Set-Content a2.txt; "c3" | Set-Content a3.txt
New-Item -ItemType Directory -Path d4 | Out-Null
Copy-Item a1.txt a2.txt a3.txt d4
$results += T "Copy-Item 四位置" (((Get-ChildItem d4 -Name) -join ",") -eq "a1.txt,a2.txt,a3.txt")
# 49. Move-Item 三位置（用独立文件，不影响后续数组源测试）
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
# 52. Set-Content 管道 + 命名 Path 缺值（旧行为：不写）
"x" | Set-Content -Path np.txt
$results += T "管道+命名缺值不写" ((-not (Test-Path np.txt)))
# 53. Add-Content 位置缺值追加空
Add-Content ac.txt
$results += T "Add-Content 位置缺值写空" (((Test-Path ac.txt)))
# 54. Get-ChildItem 通配多路径（目录含其它 txt，断言改为包含性检查）
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
# 58. Compare-Object 命名 + 位置跳槽
$co = Compare-Object -DifferenceObject b a
$results += T "Compare-Object 跳槽" (($co.Count -eq 2))
# 59. Set-Alias 两位置
Set-Alias al1 cv1 2>$null
$results += T "Set-Alias 两位置" (((Get-Alias al1).Definition -eq "cv1"))
# 60. 管道输入 Get-Content（无路径）
"p1" | Set-Content p.txt
$gp = Get-Content p.txt
$results += T "Get-Content 单文件" (($gp -join ",") -eq "p1")
# 61. Select-Object 无管道位置当数据（回归）
$so2 = Select-Object x y
$results += T "Select-Object 无管道回归" (($so2 -join ",") -eq "x,y")

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
# 65. $script: 作用域写回（函数内改脚本作用域变量）
$sf = 0
function SetSF { $script:sf = 5 }
SetSF
$results += T '$script: 作用域写回' (($sf -eq 5))

Write-Output "== 词法与数字 =="

# 66. 除法 5/2（数字后的 / 是运算符，不是路径）
$q = 5/2
$results += T "除法 5/2" (($q -eq 2.5))
# 67. 0x10 十六进制
$hx = 0x10
$results += T "0x10 十六进制" (($hx -eq 16))
# 68. 1KB 后缀（二进制倍数）
$kb = 1KB
$results += T "1KB 后缀" (($kb -eq 1024))
# 69. 浮点加法不截断（2 + 1/2 = 2.5）
$fa = 2 + 1/2
$results += T "浮点加法" (($fa -eq 2.5))

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
# 74. 除零报错（结果 $null、$? 置 false）
$zr = 5/0
if (-not $?) { $zd = "err" } else { $zd = "ok" }
$results += T "除零报错" (($zr -eq $null) -and ($zd -eq "err"))
# 75. 布尔-数字顺序比较（$true=1、$false=0 参与数字比较）
$results += T "布尔比较 -lt" (($true -lt 2) -and (-not ($true -gt 1)) -and ($false -lt 1))

Write-Output "== 对象方法 =="

# 76. DateTime ToShortDateString（Get-Date -Date 指定日期）
$dts = (Get-Date -Date "2020-01-15").ToShortDateString()
$results += T "DateTime ToShortDateString" (($dts -eq "1/15/2020"))
# 77. FileInfo 虚拟属性 Extension（Get-Item 返回对象的路径派生属性）
"x" | Set-Content ov.txt
$fie = (Get-Item ov.txt).Extension
$results += T "FileInfo Extension" (($fie -eq ".txt"))
# 78. 字符串 PadLeft（总宽,填充字符）
$results += T "字符串 PadLeft" (("7".PadLeft(3,"0")) -eq "007")
# 79. TrimStart 无参清前导空白
$ts = "  x  ".TrimStart()
$results += T "TrimStart 空白" (($ts -eq "x  "))
# 80. Split 无参按任意空白分割（tab/空格/换行，连续空白合并）
$sp3 = ("a`tb c").Split()
$results += T "Split 无参空白" (($sp3.Count -eq 3))

Write-Output "== 功能补全 =="

# 81. Profile 启动加载（写启动脚本 → 子进程验证 → 清理；标记文件相对子进程 CWD=reg）
$root = Split-Path (Split-Path $PSCommandPath)
$profDir = "$HOME/.config/powershell"
New-Item -ItemType Directory -Force $profDir | Out-Null
$profMark = "profile-mark.txt"
Remove-Item $profMark -Force 2>$null >$null
Set-Content "$profDir/profile.ps1" "Set-Content $profMark 'loaded'"
sh -c "$root/powershell -NoLogo -Command '1+1'" | Out-Null
$profOk = Test-Path $profMark
Remove-Item -Recurse -Force $profDir
$results += T "Profile 启动加载" ($profOk)
# 82. Set-Content -Encoding（utf8BOM 3 字节 BOM + 内容；ascii 非 ASCII 变单字节）
Set-Content enc_bom.txt "hi" -Encoding utf8BOM
Set-Content enc_ascii.txt "héllo" -Encoding ascii
$e1 = (Get-Item enc_bom.txt).Length
$e2 = (Get-Item enc_ascii.txt).Length
$results += T "Set-Content -Encoding" (($e1 -eq 6) -and ($e2 -eq 6))
# 83. Out-File -Encoding unicode（UTF-16LE：BOM 2 字节 + 每字符 2 字节）
"x" | Out-File enc_u16.txt -Encoding unicode
$u16 = (Get-Item enc_u16.txt).Length
$results += T "Out-File -Encoding" (($u16 -eq 6))
# 84. Test-Path -PathType（Leaf 文件 / Container 目录）
$results += T "Test-Path -PathType" ((Test-Path ov.txt -PathType Leaf) -and -not (Test-Path ov.txt -PathType Container) -and (Test-Path dir1 -PathType Container))
# 85. Get-Member -MemberType（Property 过滤后全部是 Property）
$gm2 = Get-Item ov.txt | Get-Member -MemberType Property
$results += T "Get-Member -MemberType" (($gm2.Count -gt 0) -and ($gm2[0].MemberType -eq "Property"))
# 86. ConvertTo-Json -Depth（嵌套展开截断）
$jd = @{ a = @{ b = @{ c = 1 } } } | ConvertTo-Json -Depth 1
$results += T "ConvertTo-Json -Depth" (($jd -like '*"a": {}*'))
# 87. New-Object（PSObject -Property 哈希表变属性）
$no = New-Object PSObject -Property @{ aa = 1; bb = "z" }
$results += T "New-Object" (($no.aa -eq 1) -and ($no.bb -eq "z"))
# 88. [pscustomobject] 类型字面量（条目变属性）
$pc = [pscustomobject]@{ nn = 7; ss = "q" }
$results += T "[pscustomobject]" (($pc.nn -eq 7) -and ($pc.ss -eq "q"))

Write-Output "== 对象管道行为对齐 =="

# 89. Select-String 默认大小写不敏感（匹配大小写变体）
$ssn = "Hello","HELLO" | Select-String "hello"
$results += T "Select-String 默认不敏感" (($ssn.Count -eq 2))
# 90. Select-String -CaseSensitive 只认精确大小写
$ssc = "Hello","hello" | Select-String "hello" -CaseSensitive
$results += T "Select-String CaseSensitive" (($ssc.Count -eq 1) -and ($ssc[0].Line -eq "hello"))
# 91. Select-String 管道编号按对象序号，未匹配对象占号（x 不匹配，b/c 编号 2/3）
$ssl = "x","b","c" | Select-String "[bc]" | ForEach-Object { "$($_.LineNumber):$($_.Line)" }
$results += T "Select-String 管道编号" (($ssl -join ",") -eq "2:b,3:c")
# 92. Group-Object 默认大小写不敏感合并，Name 取首次原值；组数用 @() 包裹取集合计数
$gn = "apple","Apple" | Group-Object
$results += T "Group-Object 默认合并" ((@($gn).Count -eq 1) -and ($gn[0].Name -eq "apple") -and ($gn[0].Count -eq 2))
# 93. Group-Object -CaseSensitive 分开分组
$gc3 = ("apple","Apple" | Group-Object -CaseSensitive)
$results += T "Group-Object CaseSensitive" (($gc3.Count -eq 2))
# 94. Compare-Object 默认大小写不敏感，先右后左输出（b 两边共有被排除）
$co2 = Compare-Object -ReferenceObject a,B -DifferenceObject b,C
$coSeq = $co2 | ForEach-Object { $_.SideIndicator + $_.InputObject }
$results += T "Compare-Object 默认顺序" (($coSeq -join ",") -eq "=>C,<=a")
# 95. Compare-Object -IncludeEqual 相等项显示基准集的值（A 与 a 相等显示 A）
$cie = Compare-Object -ReferenceObject A -DifferenceObject a -IncludeEqual
$results += T "Compare-Object IncludeEqual值" (($cie[0].SideIndicator -eq "==") -and ($cie[0].InputObject -eq "A"))
# 96. Sort-Object -Unique 折叠大小写去重（apple/Apple 合并取排序后首个）
$su = "banana","apple","Apple" | Sort-Object -Unique
$results += T "Sort-Object Unique折叠" (($su -join ",") -eq "apple,banana")
# 97. Sort-Object -Unique -CaseSensitive 全保留
$suc = ("apple","Apple","banana" | Sort-Object -Unique -CaseSensitive).Count
$results += T "Sort-Object Unique CS" (($suc -eq 3))
# 98. Select-Object -First 0 返回空
$f0 = 1,2,3 | Select-Object -First 0
$l0 = 1,2,3 | Select-Object -Last 0
$results += T "Select-Object First/Last 0" ((($f0 -eq $null)) -and ($l0 -eq $null))
# 99. Measure-Object 未开的统计字段为 $null（字段总是补全）
$m9 = 1,2,3 | Measure-Object
$results += T "Measure-Object 字段补全" (($m9.Sum -eq $null) -and ($m9.Average -eq $null) -and ($m9.Count -eq 3))
# 100. Measure-Object Sum/Average 遇非数字作废置 $null
$mn = "a","b" | Measure-Object -Sum
$mx = 1,"a",2 | Measure-Object -Sum
$results += T "Measure-Object 非数字作废" ((($mn.Sum -eq $null)) -and ($mx.Sum -eq $null))
# 101. Measure-Object -Property 模式 Count 只数有该属性的对象
$mp = @{a=1},@{a=2},@{b=3} | Measure-Object -Property a -Sum
$results += T "Measure-Object Property Count" (($mp.Count -eq 2) -and ($mp.Sum -eq 3))

Write-Output "== Get-Process 真值（仅 Linux） =="

# 102. 进程 CPU/Memory 来自 /proc，pid 1 常驻进程两者应大于 0
$p1 = Get-Process | Where-Object { $_.Id -eq 1 }
$results += T "Get-Process 真值" (($p1.CPU -gt 0) -and ($p1.Memory -gt 0))

Write-Output "== 标量索引 =="

# 103. 标量变量 [0] 返回自身，非零下标返回 $null
$scalar = [pscustomobject]@{ tag = "t" }
$results += T "标量索引自身" (($scalar[0].tag -eq "t") -and ($scalar[5] -eq $null))
# 104. GroupInfo 不被当集合：[0] 是组对象本身，成员在 Group 属性
$gi = "x","y","x" | Group-Object
$results += T "GroupInfo 索引不穿透" ((@($gi).Count -eq 2) -and ($gi[0].Name -eq "x") -and ($gi[0].Group -join ",") -eq "x,x")
# 105. Join-String 默认分隔符是空串（与 PowerShell 一致）
$results += T "Join-String 默认分隔符" (((1..3 | Join-String) -eq "123") -and ((1..3 | Join-String -Separator ",") -eq "1,2,3"))
# 106. @() 元素按输出流摊平：@($arr) 得元素本身；@() 内可含管道
$aa = 1,2,3
$results += T "@() 摊平与管道" ((@($aa).Count -eq 3) -and (@(1,2 | ForEach-Object { $_ * 10 }) -join ",") -eq "10,20")

Write-Output "== 解析终止性 =="

# 107. switch 与 { 之间换行的输入能解析或报错后正常返回，不失去响应
# 该写法当前报"期望 '{'"，进程以退出码 1 结束；断言只保证退出而非卡死（退出码 124 是超时标记）
$root = Split-Path (Split-Path $PSCommandPath)
$nlSrc = "switch (2)`n{ default { 'd' } }"
sh -c "echo '$nlSrc' | $root/powershell -NoLogo -NoProfile -Command -" 2>$null >$null
$results += T "switch 换行大括号不失去响应" ($LASTEXITCODE -eq 1)

# == 结尾统计 ==
Write-Output ""
$failN = 0
foreach ($r in $results) {
    Write-Output $r
    if ($r -like "FAIL  *") { $failN++ }
}
Write-Output ""
Write-Output ("结果: 通过 " + ($results.Count - $failN) + "  失败 " + $failN)

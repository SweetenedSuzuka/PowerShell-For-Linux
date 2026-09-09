# 属性取值与表格读写。
# 228. 按名取单个属性值
Set-Content ipv.txt "hello!"
$ivLen = Get-ItemPropertyValue ipv.txt -Name Length
$ivPos = Get-ItemPropertyValue ipv.txt Length
$ivDt = Get-ItemPropertyValue ipv.txt -Name LastWriteTime
$results += T "按名取单个属性值" ((($ivLen -eq 7)) -and (($ivPos -eq 7)) -and (($ivDt -is [datetime])))
# 229. 缺失路径与缺失属性记错误后继续
$ie0 = $Error.Count
Get-ItemPropertyValue zzz-no-such-ipv-123 -Name Length 2>$null | Out-Null
Get-ItemPropertyValue ipv.txt -Name NoSuchPropXYZ 2>$null | Out-Null
$results += T "取属性值报错继续" (($Error.Count -eq ($ie0 + 2)))
# 230. 导出表格写文件
$eo1 = [pscustomobject]@{ Name = "x"; N = 1 }
$eo2 = [pscustomobject]@{ Name = "y"; N = 2 }
$eo1, $eo2 | Export-Csv eo.csv
$eoBack = Get-Content eo.csv
$eoConv = $eo1, $eo2 | ConvertTo-Csv
$results += T "导出表格写文件" ((($eoBack.Count -eq 3)) -and (($eoBack[0] -eq $eoConv[0])) -and (($eoBack[2] -eq $eoConv[2])))
# 231. 追加不重复表头，防覆盖报错
$eo1 | Export-Csv eoapp.csv
$eo2 | Export-Csv eoapp.csv -Append
$ea0 = $Error.Count
$eo1 | Export-Csv eoapp.csv -NoClobber 2>$null | Out-Null
$results += T "追加与防覆盖" (((((Get-Content eoapp.csv).Count) -eq 3)) -and (($Error.Count -eq ($ea0 + 1))))
# 232. 分隔符与属性筛选列，缺输入不建文件
$eo1, $eo2 | Export-Csv eosemi.csv -Delimiter ";"
$eo1 | Export-Csv eoprop.csv -Property N
Export-Csv eonoinput.csv
$results += T "分隔符与筛选列" ((((Get-Content eosemi.csv)[0] -eq "Name;N")) -and (((Get-Content eoprop.csv)[0] -eq "N")) -and ((-not (Test-Path eonoinput.csv))))
# 233. 读表格文件
"Name,N`nx,1`ny,2" | Set-Content ic.csv
$icBack = Import-Csv ic.csv
$results += T "读表格文件" ((($icBack.Count -eq 2)) -and (($icBack[0].Name -eq "x")) -and (($icBack[1].N -eq "2")))
# 234. 分隔符与自定义表头，类型行跳过
"Name;N`nx;1" | Set-Content icsemi.csv
$icSemi = Import-Csv icsemi.csv -Delimiter ";"
"1,2" | Set-Content ichead.csv
$icHead = Import-Csv ichead.csv -Header H1,H2
"#TYPE System.Management.Automation.PSCustomObject" | Set-Content ictype.csv
'"Name","N"' | Add-Content ictype.csv
'"x","1"' | Add-Content ictype.csv
$icType = Import-Csv ictype.csv
$results += T "分隔符表头类型行" ((($icSemi[0].Name -eq "x")) -and (($icHead.Count -eq 1)) -and (($icHead[0].H1 -eq "1")) -and (($icHead[0].H2 -eq "2")) -and (($icType[0].Name -eq "x")) -and (($icType[0].N -eq "1")))
# 235. 缺失路径报错继续，空文件与只有表头无输出
$ie1 = $Error.Count
Import-Csv ic-no-such-999.csv 2>$null
$icMissOk = $?
"" | Set-Content icempty.csv
"a,b" | Set-Content ichonly.csv
$results += T "缺失路径与空文件" ((($Error.Count -eq ($ie1 + 1))) -and ((-not $icMissOk)) -and (($null -eq (Import-Csv icempty.csv))) -and (($null -eq (Import-Csv ichonly.csv))))
# 236. 集合属性增删项
$ul1 = [pscustomobject]@{ L = @(1,2,3) }
$ul1r = $ul1 | Update-List -Property L -Add 4
$ul2 = [pscustomobject]@{ L = @(1,2,3,2) }
$ul2 | Update-List -Property L -Remove 2 | Out-Null
$results += T "集合属性增删项" ((($ul1.L -join ",") -eq "1,2,3,4") -and ((($ul1r.L -join ",")) -eq "1,2,3,4") -and (($ul2.L -join ",") -eq "1,3,2"))
# 237. 整列替换与加删同用，参数集冲突报错
$ul3 = [pscustomobject]@{ L = @(1,2,3) }
$ul3 | Update-List -Property L -Replace @(7,8) | Out-Null
$ul4 = [pscustomobject]@{ L = @(1,2,3) }
$ul4 | Update-List -Property L -Add 4 -Remove 2 | Out-Null
$ue0 = $Error.Count
$ul4 | Update-List -Property L -Add 4 -Replace @(9) 2>$null | Out-Null
$results += T "整列替换与参数集" (((($ul3.L -join ",")) -eq "7,8") -and ((($ul4.L -join ",")) -eq "1,3,4") -and (($Error.Count -eq ($ue0 + 1))))
# 238. 非集合与缺属性报错继续，无输入缺参数返回空
$ul5 = [pscustomobject]@{ L = 5 }
$ue1 = $Error.Count
$ul5 | Update-List -Property L -Add 4 2>$null
if ($?) { $ul5ok = $true } else { $ul5ok = $false }
$ul6 = [pscustomobject]@{ A = 1 }
$ul6 | Update-List -Property L -Add 4 2>$null
$results += T "非集合缺属性与空输入" ((($Error.Count -eq ($ue1 + 2))) -and ((-not $ul5ok)) -and (($null -eq (Update-List -Property L -Add 4))) -and (($null -eq ($ul6 | Update-List))))

# 默认出口、新指令与严格模式。
# 239. 默认出口直通显示并截断管道
$odc = 9 | Out-Default | Measure-Object | Select-Object -ExpandProperty Count
$odr2 = Out-Default -InputObject @()
$results += T "默认出口截断" ((($odc -eq 0)) -and (($null -eq $odr2)))
# 240. 输入冲突与超量位置报错继续，记录开关接受忽略
$oe0 = $Error.Count
1 | Out-Default -InputObject 2 2>$null
if ($?) { $odcok = $true } else { $odcok = $false }
Out-Default foo 2>$null
if ($?) { $odpok = $true } else { $odpok = $false }
8 | Out-Default -Transcript 2>$null
if ($?) { $odotok = $true } else { $odotok = $false }
$results += T "输入冲突位置报错" ((($Error.Count -eq ($oe0 + 2))) -and ((-not $odcok)) -and ((-not $odpok)) -and ($odotok))
# 241. 安全随机范围与个数
$sr1 = Get-SecureRandom -Maximum 100
$sr2 = Get-SecureRandom -Minimum 10 -Maximum 20
$sr3 = Get-SecureRandom 50
$sr4 = Get-SecureRandom -Maximum 10 -Count 3
$results += T "安全随机范围个数" ((($sr1 -ge 0) -and ($sr1 -lt 100)) -and (($sr2 -ge 10) -and ($sr2 -lt 20)) -and (($sr3 -ge 0) -and ($sr3 -lt 50)) -and (($sr4.Count -eq 3)))
# 242. 取样洗牌与参数集报错
$ss1 = Get-SecureRandom -InputObject (1,2,3,4,5) -Count 3
$ss2 = Get-SecureRandom -InputObject (1,2,3,4,5) -Shuffle
$se0 = $Error.Count
Get-SecureRandom -InputObject (1,2) -Maximum 5 2>$null | Out-Null
Get-SecureRandom -Minimum 10 -Maximum 5 2>$null | Out-Null
Get-SecureRandom -InputObject (1,2,3) -Count -1 2>$null | Out-Null
Get-SecureRandom -Maximum 10 -Count 0 2>$null | Out-Null
$results += T "取样洗牌参数集" ((($ss1.Count -eq 3)) -and ((($ss1 | Sort-Object -Unique).Count -eq 3)) -and (($ss2.Count -eq 5)) -and ((($ss2 | Sort-Object) -join ",") -eq "1,2,3,4,5") -and (($Error.Count -eq ($se0 + 4))))
# 243. 读取最新错误记录
Get-Content ge-no-such-aaa.txt 2>$null | Out-Null
Get-Content ge-no-such-bbb.txt 2>$null | Out-Null
$ge1 = Get-Error
$ge2 = Get-Error -Newest 2
$results += T "读取最新错误记录" ((($ge1.Count -eq 1)) -and (($ge2.Count -eq 2)) -and (($ge1[0].Message -eq $ge2[0].Message)) -and (($ge2[0].Message -ne $ge2[1].Message)))
# 244. 条数边界与输入对象
$ge3 = Get-Error -Newest 99
$ge4 = Get-Error -InputObject $Error[0]
$ge5 = $Error[0..1] | Get-Error
$ge0 = $Error.Count
Get-Error -Newest 0 2>$null | Out-Null
Get-Error -Newest 1 -InputObject $Error[0] 2>$null | Out-Null
$results += T "条数边界输入对象" ((($ge3.Count -eq $ge0)) -and (($ge4.Count -eq 1)) -and (($ge5.Count -eq 2)) -and (($null -eq (Get-Error -InputObject ($Error[0],$Error[1])))) -and (($Error.Count -eq ($ge0 + 2))))
# 245. 本地解析与位置类型
$dn1 = Resolve-DnsName localhost
$dn2 = Resolve-DnsName localhost A
$dn3 = "localhost" | Resolve-DnsName
$results += T "本地解析位置类型" ((($dn1.Data -contains "127.0.0.1")) -and (($dn2.Data -contains "127.0.0.1")) -and (($dn3.Data -contains "127.0.0.1")))
# 246. 解析失败与参数报错继续
$de0 = $Error.Count
Resolve-DnsName invalid.invalid 2>$null | Out-Null
Resolve-DnsName localhost -Type SOA 2>$null | Out-Null
Resolve-DnsName invalid.invalid -Server notahost.invalid 2>$null | Out-Null
Resolve-DnsName localhost A extra 2>$null | Out-Null
$results += T "解析失败参数报错" ((($Error.Count -eq ($de0 + 4))) -and (($null -eq (Resolve-DnsName))))
# 247. 凭据直接返回与组合报错
$gc0 = [pscustomobject]@{ UserName = "u"; Password = "p" }
$gc1 = Get-Credential -Credential $gc0
$ge0 = $Error.Count
Get-Credential -Credential $gc0 -UserName x 2>$null | Out-Null
Get-Credential a b 2>$null | Out-Null
$results += T "凭据直接返回组合报错" ((($gc1.UserName -eq "u")) -and (($gc1.Password -eq "p")) -and (($Error.Count -eq ($ge0 + 2))))
# 248. 提示输入凭据识别与非交互
$gcInner = "$root/test/tmp/reg/gcpinner.ps1"
'$c = Get-Credential' | Set-Content $gcInner
'"^u=" + $c.UserName' | Add-Content $gcInner
'$c' | Add-Content $gcInner
'$d = Get-Credential $c' | Add-Content $gcInner
'"^s=" + $d.UserName' | Add-Content $gcInner
'$e0 = $Error.Count' | Add-Content $gcInner
'Get-Credential -UserName $c 2>$null | Out-Null' | Add-Content $gcInner
'"^de=" + ($Error.Count - $e0)' | Add-Content $gcInner
$gcPipe = sh -c "printf 'testu1\ns3cr3tzz\n' | $root/powershell -NoLogo -NoProfile -File $gcInner" 2>$null
$gcNi = sh -c "$root/powershell -NoLogo -NoProfile -NonInteractive -Command 'Get-Credential'" 2>$null
$results += T "提示输入凭据识别非交互" ((($gcPipe -join "") -like "*^u=testu1*") -and ((($gcPipe -join "") -notlike "*s3cr3tzz*")) -and (($gcPipe -contains "^s=testu1")) -and (($gcPipe -contains "^de=1")) -and ((($gcNi -join "") -eq "")))
# 249. 严格模式未定义检查
Set-StrictMode -Version Latest
$se0 = $Error.Count
$smCaught = try { $nosuchvar_sm249 } catch { "yes" }
$smDef = 42
Set-StrictMode -Off
$nosuchvar_sm249b
$results += T "严格模式未定义检查" ((($Error.Count -eq ($se0 + 1))) -and (($smCaught -eq "yes")) -and (($smDef -eq 42)) -and (($null -eq $nosuchvar_sm249b)))
# 250. 版本写法参数集作用域
Set-StrictMode -Version 1.0
Set-StrictMode -Version 2.0
Set-StrictMode -Version 3
function SMF250 { Set-StrictMode -Off; $nosuchvar_sm250 }
$smfOut = SMF250 2>$null
try { $nosuchvar_sm250b } catch { }
$se2 = $Error.Count
Set-StrictMode -Version 9.9 -ErrorAction SilentlyContinue | Out-Null
Set-StrictMode -Version Latest -Off 2>$null | Out-Null
Set-StrictMode 2>$null | Out-Null
Set-StrictMode 1.0 extra 2>$null | Out-Null
Set-StrictMode -Off
$results += T "版本写法参数集作用域" ((($null -eq $smfOut)) -and (($Error.Count -eq ($se2 + 4))))
# 251. 表达式语句尾随重定向
$rvx = $null
$rvx 2>rv_err.txt
"world" > rv_out.txt
"more" >> rv_out.txt
$results += T "表达式尾随重定向" ((((Get-Content rv_out.txt) -join ",") -eq "world,more") -and ((Test-Path rv_err.txt)) -and ((@(Get-Content rv_err.txt).Count -eq 0)))

# 等待、数组字面、日期与问号变量。
# 199. Wait-Process 缺失 PID 报错
$Error.Clear()
Wait-Process -Id 999991
$results += T "等待缺失进程" ((($? -eq $false)) -and (($Error.Count -ge 1)))
# 200. Wait-Process 缺失进程名报错
$Error.Clear()
Wait-Process -Name "zzz-no-such-proc-abc"
$results += T "等待缺失进程名" ((($? -eq $false)) -and (($Error.Count -ge 1)))
# 201. @() 首命令接纳逗号实参
$at1 = @(Write-Output "a", "b")
$results += T "数组首命令" ((($at1.Count -eq 2)) -and ((($at1 -join "|") -eq "a|b")))
# 202. @() 命名首命令可解析
$at2 = @(Write-Output -InputObject "a", "b")
$results += T "数组命名首命令" ((($at2.Count -eq 2)) -and ((($at2 -join "|") -eq "a|b")))
# 203. @() 分号多语句不受逗号分隔影响
$at3 = @(Get-Date; Get-Date)
$results += T "数组分号多语句" (($at3.Count -eq 2))
# 204. 字符串内 $? 展开
Write-Output "x" > $null
$qi = "AFTER=$?"
Get-Content zzz-no-such-xyz-123 2>$null >$null
$qj = "AFTER=$?+$?x"
$results += T "字符串问号展开" ((($qi -eq "AFTER=True")) -and (($qj -eq "AFTER=False+Falsex")))
# 205. Start-Sleep 超量位置实参报错
$Error.Clear()
Start-Sleep 1 2
$results += T "睡眠超量位置" ((($? -eq $false)) -and (($Error.Count -ge 1)))
# 206. [version] 类型转换与未知类型本地化
$ver = [version]"1.2"
$Error.Clear()
$zz = [zzznope]"a"
$results += T "版本转换" ((($ver.Major -eq 1)) -and (($ver.Minor -eq 2)) -and (($ver.Build -eq -1)) -and (("$ver" -eq "1.2")) -and (($? -eq $false)) -and (($Error.Count -ge 1)) -and (("$($Error[0])" -like "*无法找到类型*")))
# 207. Get-Date 位置日期
$gd = Get-Date 2020-1-1
$results += T "位置日期" (($gd.Year -eq 2020))
# 208. 开关后 2> 重定向生效
Get-ChildItem | Format-Table -AutoSize 2>$null > redir22.out
$rr = @(Get-Content redir22.out)
Remove-Item redir22.out
$results += T "开关重定向" ((($rr.Count -gt 1)) -and (($rr[0] -like "*Mode*")))
# 209. Get-ChildItem 字面缺失路径报错
$Error.Clear()
$gm = @(Get-ChildItem zzz-no-such-xyz-123 2>$null)
$results += T "缺失路径报错" ((($gm.Count -eq 0)) -and (($Error.Count -ge 1)))
# 210. 缺失路径继续其余路径
$Error.Clear()
$gn = @(Get-ChildItem ., zzz-no-such-xyz-123 2>$null)
$results += T "缺失路径继续" ((($gn.Count -gt 0)) -and (($Error.Count -ge 1)))
# 211. Get-Date 非法日期报错
$Error.Clear()
Get-Date "zzz-nope" 2>$null
$results += T "非法日期报错" ((($? -eq $false)) -and (($Error.Count -ge 1)))
# 212. 日期对象链式成员
$dd = Get-Date 2024-01-02
$results += T "日期链式成员" ((($dd.Date.Year -eq 2024)) -and (($dd.Ticks -gt 0)))
# 213. 管道内错误粘滞
Get-Content zzz-no-such-xyz-123 2>$null | Out-Null
$results += T "管道错误粘滞" (($? -eq $false))
# 214. 成员读取重置
$qm = @(Get-Content zzz-no-such-xyz-123 2>$null).Count
$results += T "成员读取重置" (($? -eq $true))

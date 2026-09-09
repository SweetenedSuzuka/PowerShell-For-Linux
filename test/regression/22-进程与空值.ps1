# 进程、睡眠与空值管道。
# 192. Stop-Process 多 Id 逐个停止
$sp1 = Start-Process /bin/sleep -ArgumentList "30"
$sp2 = Start-Process /bin/sleep -ArgumentList "30"
Stop-Process -Id $sp1.Id, $sp2.Id
Start-Sleep -Milliseconds 800
$results += T "多 Id 停止" (((Test-Path "/proc/$($sp1.Id)") -eq $false) -and ((Test-Path "/proc/$($sp2.Id)") -eq $false))
Stop-Process -Id $sp1.Id, $sp2.Id
# 193. Get-Process -Name 通配
$sp3 = Start-Process /bin/sleep -ArgumentList "30"
$results += T "进程名通配" (((@(Get-Process -Name "slee*").Count) -ge 1) -and ((@(Get-Process -Name "zzz-no-such-*").Count) -eq 0))
Stop-Process -Id $sp3.Id
# 194. Start-Sleep 两个时间参数互斥报错
$Error.Clear()
Start-Sleep -Seconds 1 -Milliseconds 100
$results += T "睡眠双参互斥" ((($? -eq $false)) -and (($Error.Count -ge 1)))
# 195. $null 进管道跑一次
$nl1 = @($null | ForEach-Object { "x" })
$results += T "管道空值" ((($nl1.Count -eq 1)) -and (($nl1[0] -eq "x")))
# 196. $null 不参与计数选择
$nlm = $null | Measure-Object
$results += T "空值计数" (($nlm.Count -eq 0))
$results += T "空值选择" (((@($null | Select-Object -First 1).Count) -eq 0))
# 197. $null 不参与排序分组
$results += T "空值排序" (((@($null | Sort-Object).Count) -eq 0))
$results += T "空值分组" (((@($null | Group-Object).Count) -eq 0))
# 198. $null 写文件零字节
$null | Set-Content test/tmp/nullbyte.txt
$results += T "空值写文件" ((((Get-Item test/tmp/nullbyte.txt).Length) -eq 0))

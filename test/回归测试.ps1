# 回归测试入口：按顺序载入共用前置项与各分组，末尾汇总打印结果。
# 核验.ps1 只检查 $? 而不检查实际内容，此脚本会检查命令输出的实际内容。
# 用法：./powershell -NoLogo -NoProfile -File test/回归测试.ps1。
# 可重复运行：共用前置项在 test/tmp/reg 子目录工作，开头清空重建，结尾保留供排查。

$root = Split-Path (Split-Path $PSCommandPath)
& "$root/test/lib/TestCommon.ps1"
& "$root/test/regression/01-参数绑定.ps1"
& "$root/test/regression/02-数组路径展开.ps1"
& "$root/test/regression/03-边界核对.ps1"
& "$root/test/regression/04-语言结构.ps1"
& "$root/test/regression/05-词法数字.ps1"
& "$root/test/regression/06-求值比较.ps1"
& "$root/test/regression/07-对象方法.ps1"
& "$root/test/regression/08-功能补全.ps1"
& "$root/test/regression/09-对象管道.ps1"
& "$root/test/regression/10-进程真值.ps1"
& "$root/test/regression/11-标量索引.ps1"
& "$root/test/regression/12-解析终止.ps1"
& "$root/test/regression/13-排版兼容.ps1"
& "$root/test/regression/14-WhatIf数据安全.ps1"
& "$root/test/regression/15-类型字面量.ps1"
& "$root/test/regression/16-调用运算符.ps1"
& "$root/test/regression/17-块与过滤器.ps1"
& "$root/test/regression/18-错误体系.ps1"
& "$root/test/regression/19-选择投影.ps1"
& "$root/test/regression/20-数组与运算符.ps1"
& "$root/test/regression/21-异常与大小写.ps1"
& "$root/test/regression/22-进程与空值.ps1"
& "$root/test/regression/23-等待与日期.ps1"
& "$root/test/regression/24-虚拟属性与输入.ps1"
& "$root/test/regression/25-数组运算与抛出.ps1"
& "$root/test/regression/26-属性与表格.ps1"
& "$root/test/regression/27-出口与新指令.ps1"

# 事项 252 至 254 在入口顶层执行：记录文件只收入口顶层输出，载入文件内的输出进不了记录文件。
# 252. 开始记录写入追加防覆盖
Start-Transcript trrec.log | Out-Null
"rec-marker-252"
$trc = @(Get-Content trrec.log)
Start-Transcript trrec.log -Append | Out-Null
Start-Transcript trrec2.log | Out-Null
$trc2 = @(Get-Content trrec.log)
$te0 = $Error.Count
Start-Transcript trrec.log -NoClobber 2>$null | Out-Null
$results += T "开始记录写入追加防覆盖" ((($trc -contains "PowerShell transcript start")) -and (($trc -contains "rec-marker-252")) -and ((@($trc2 | Where-Object { $_ -eq "PowerShell transcript start" }).Count -eq 2)) -and (($trc2 -contains "PowerShell transcript end")) -and (($Error.Count -eq ($te0 + 1))))
# 253. 停止记录结束块空闲报错
$leftover = Stop-Transcript
$stp0 = $Error.Count
Stop-Transcript 2>$null | Out-Null
$stpIdle = ($Error.Count -eq ($stp0 + 1))
Stop-Transcript -WhatIf | Out-Null
$stpWhatifIdle = ($Error.Count -eq ($stp0 + 1))
Start-Transcript trstop.log | Out-Null
"stop-marker-253"
$sst = Stop-Transcript
$stc = @(Get-Content trstop.log)
$results += T "停止记录结束块空闲报错" ((($leftover -like "*trrec2.log*")) -and ($stpIdle) -and ($stpWhatifIdle) -and (($sst -like "*trstop.log*")) -and (($stc -contains "PowerShell transcript end")) -and (($stc -contains "stop-marker-253")))
# 254. 替换进程缺命令报错、无参数直接返回（成功替换会结束进程，只断言非替换路径）
$swp0 = $Error.Count
Switch-Process 2>$null | Out-Null
$swpNoop = ($Error.Count -eq $swp0)
Switch-Process -WithCommand "definitely-missing-xyz-123" 2>$null | Out-Null
$swpMiss = ($Error.Count -eq ($swp0 + 1))
Switch-Process -WithCommand "/bin/echo" -WhatIf 2>$null | Out-Null
$swpWI = ($Error.Count -eq ($swp0 + 2))
$results += T "替换进程缺命令报错无参空转" (($swpNoop) -and ($swpMiss) -and ($swpWI))
$Error.Clear()
$ErrorActionPreference = 'Continue'

# == 结尾统计 ==
Write-Output ""
$failN = 0
foreach ($r in $results) {
    Write-Output $r
    if ($r -like "FAIL  *") { $failN++ }
}
Write-Output ""
Write-Output ("结果: 通过 " + ($results.Count - $failN) + "  失败 " + $failN)

# WhatIf 与数据安全。
Write-Output "== WhatIf 与数据安全 =="

# 113. Remove-Item -WhatIf 打印预演且不删除
"whatif-body" | Set-Content test/tmp/wi.txt
$wiOut = Remove-Item test/tmp/wi.txt -WhatIf
$results += T "Remove-Item -WhatIf 不删除" ((Test-Path test/tmp/wi.txt) -and (($wiOut -join "").Length -gt 0))
Remove-Item test/tmp/wi.txt -Force
# 114. Rename-Item 目标已存在时报错且原文件保留
"a1" | Set-Content test/tmp/rn-a.txt
"b1" | Set-Content test/tmp/rn-b.txt
Rename-Item test/tmp/rn-a.txt test/tmp/rn-b.txt 2>test/tmp/rn.err >$null
$rnOk = (((Get-Content test/tmp/rn-a.txt) -eq "a1") -and ((Get-Content test/tmp/rn-b.txt) -eq "b1") -and ((Get-Content test/tmp/rn.err).Count -gt 0))
$results += T "Rename-Item 目标已存在报错" ($rnOk)
Remove-Item test/tmp/rn-a.txt,test/tmp/rn-b.txt -Force 2>$null
Remove-Item test/tmp/rn.err -Force 2>$null
# 115. Clear 系对不存在路径报错而非创建空文件
Clear-Content test/tmp/no-such-cc.txt 2>test/tmp/cc.err >$null
$rCc = ((-not (Test-Path test/tmp/no-such-cc.txt)) -and ((Get-Content test/tmp/cc.err).Count -gt 0))
Clear-Item test/tmp/no-such-ci.txt 2>test/tmp/ci.err >$null
$rCi = ((-not (Test-Path test/tmp/no-such-ci.txt)) -and ((Get-Content test/tmp/ci.err).Count -gt 0))
Set-Item test/tmp/no-such-si.txt "v" 2>test/tmp/si.err >$null
$rSi = ((-not (Test-Path test/tmp/no-such-si.txt)) -and ((Get-Content test/tmp/si.err).Count -gt 0))
$results += T "Clear 系不存在路径报错" ($rCc -and $rCi)
$results += T "Set-Item 不存在路径报错" ($rSi)
Remove-Item test/tmp/cc.err,test/tmp/ci.err,test/tmp/si.err -Force 2>$null
# 116. Remove-Item -Confirm 应答 y 后执行删除
"cfa" | Set-Content test/tmp/cfa.txt
sh -c "echo y | $root/powershell -NoLogo -NoProfile -Command 'Remove-Item test/tmp/cfa.txt -Confirm'" 2>$null >$null
$results += T "-Confirm 应答 y 后删除" ((Test-Path test/tmp/cfa.txt) -eq $false)
# 117. Remove-Item -Confirm 应答 n 后保留
"cf-b" | Set-Content test/tmp/cfb.txt
sh -c "echo n | $root/powershell -NoLogo -NoProfile -Command 'Remove-Item test/tmp/cfb.txt -Confirm'" 2>$null >$null
$results += T "-Confirm 应答 n 后保留" ((Test-Path test/tmp/cfb.txt) -eq $true)
Remove-Item test/tmp/cfb.txt -Force 2>$null

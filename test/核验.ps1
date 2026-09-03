# 指令核验脚本：对照 docs/指令参考.md 逐条运行示例，用 $? 判定通过/失败。
# 用法：./powershell -NoLogo -NoProfile -File test/核验.ps1。
# 平台相关项（systemctl/ping/xclip/sudo）在非对应环境会失败，属预期，见参考文档说明。

$pass = 0
$fail = 0

# 确保临时目录存在（脚本要在 test/tmp 下建文件）
New-Item -ItemType Directory -Force test/tmp 2>$null >$null

Write-Output "== 交互与风格 =="
Set-PSVersion 7 2>$null >$null; if ($?) { $pass++; "PASS  Set-PSVersion" } else { $fail++; "FAIL  Set-PSVersion" }
Get-Host 2>$null >$null; if ($?) { $pass++; "PASS  Get-Host" } else { $fail++; "FAIL  Get-Host" }
pwd 2>$null >$null; if ($?) { $pass++; "PASS  Get-Location/pwd" } else { $fail++; "FAIL  Get-Location/pwd" }
Push-Location /; Pop-Location 2>$null >$null; if ($?) { $pass++; "PASS  Push/Pop-Location" } else { $fail++; "FAIL  Push/Pop-Location" }
./powershell -NoProfile -ExecutionPolicy Bogus -Command "Write-Output x" 2>$null >$null; if ($LASTEXITCODE -eq 2) { $pass++; "PASS  -ExecutionPolicy 非法值" } else { $fail++; "FAIL  -ExecutionPolicy 非法值" }

Write-Output "== 文件与目录 =="
Get-ChildItem -Name 2>$null >$null; if ($?) { $pass++; "PASS  Get-ChildItem" } else { $fail++; "FAIL  Get-ChildItem" }
Get-ChildItem -Directory -Name 2>$null >$null; if ($?) { $pass++; "PASS  Get-ChildItem -Directory" } else { $fail++; "FAIL  Get-ChildItem -Directory" }
Get-ChildItem -Recurse -Name 2>$null >$null; if ($?) { $pass++; "PASS  Get-ChildItem -Recurse" } else { $fail++; "FAIL  Get-ChildItem -Recurse" }
Get-Item test/核验.ps1 2>$null >$null; if ($?) { $pass++; "PASS  Get-Item" } else { $fail++; "FAIL  Get-Item" }
Get-ItemProperty test/核验.ps1 2>$null >$null; if ($?) { $pass++; "PASS  Get-ItemProperty" } else { $fail++; "FAIL  Get-ItemProperty" }
New-Item -ItemType Directory test/tmp/n1 2>$null >$null; if ($?) { $pass++; "PASS  New-Item 目录" } else { $fail++; "FAIL  New-Item 目录" }
New-Item test/tmp/n2.txt -Force 2>$null >$null; if ($?) { $pass++; "PASS  New-Item 文件" } else { $fail++; "FAIL  New-Item 文件" }
Copy-Item test/tmp/n2.txt test/tmp/n3.txt 2>$null >$null; if ($?) { $pass++; "PASS  Copy-Item" } else { $fail++; "FAIL  Copy-Item" }
Move-Item test/tmp/n3.txt test/tmp/n4.txt 2>$null >$null; if ($?) { $pass++; "PASS  Move-Item" } else { $fail++; "FAIL  Move-Item" }
Rename-Item test/tmp/n4.txt test/tmp/n5.txt 2>$null >$null; if ($?) { $pass++; "PASS  Rename-Item" } else { $fail++; "FAIL  Rename-Item" }
Remove-Item test/tmp/n5.txt 2>$null >$null; Remove-Item -Recurse -Force test/tmp/n1 2>$null >$null; if ($?) { $pass++; "PASS  Remove-Item" } else { $fail++; "FAIL  Remove-Item" }
Get-PSDrive 2>$null >$null; if ($?) { $pass++; "PASS  Get-PSDrive" } else { $fail++; "FAIL  Get-PSDrive" }

Write-Output "== 内容读写 =="
Set-Content test/tmp/c1.txt "hello" 2>$null >$null; if ($?) { $pass++; "PASS  Set-Content" } else { $fail++; "FAIL  Set-Content" }
Get-Content test/tmp/c1.txt 2>$null >$null; if ($?) { $pass++; "PASS  Get-Content" } else { $fail++; "FAIL  Get-Content" }
Add-Content test/tmp/c1.txt "world" 2>$null >$null; if ($?) { $pass++; "PASS  Add-Content" } else { $fail++; "FAIL  Add-Content" }
Add-Content -Encoding utf8BOM test/tmp/enc.txt "a" 2>$null >$null; if ($?) { $pass++; "PASS  Add-Content -Encoding" } else { $fail++; "FAIL  Add-Content -Encoding" }
if ((Get-Content test/tmp/enc.txt -TotalCount 1) -eq "a") { $pass++; "PASS  读去 BOM" } else { $fail++; "FAIL  读去 BOM" }
Clear-Content test/tmp/c1.txt 2>$null >$null; if ($?) { $pass++; "PASS  Clear-Content" } else { $fail++; "FAIL  Clear-Content" }
Set-Item env:VERIFY_VAR "v" 2>$null >$null; if ($?) { $pass++; "PASS  Set-Item 环境变量" } else { $fail++; "FAIL  Set-Item 环境变量" }
Get-FileHash test/tmp/c1.txt -Algorithm MD5 2>$null >$null; if ($?) { $pass++; "PASS  Get-FileHash" } else { $fail++; "FAIL  Get-FileHash" }
Select-String -Path test/tmp/c1.txt "hello" 2>$null >$null; if ($?) { $pass++; "PASS  Select-String" } else { $fail++; "FAIL  Select-String" }

Write-Output "== 路径处理 =="
Test-Path /etc 2>$null >$null; if ($?) { $pass++; "PASS  Test-Path" } else { $fail++; "FAIL  Test-Path" }
Resolve-Path /etc 2>$null >$null; if ($?) { $pass++; "PASS  Resolve-Path" } else { $fail++; "FAIL  Resolve-Path" }
Convert-Path /etc 2>$null >$null; if ($?) { $pass++; "PASS  Convert-Path" } else { $fail++; "FAIL  Convert-Path" }
Split-Path /etc/hostname -Leaf 2>$null >$null; if ($?) { $pass++; "PASS  Split-Path" } else { $fail++; "FAIL  Split-Path" }
Join-Path /tmp test.txt 2>$null >$null; if ($?) { $pass++; "PASS  Join-Path" } else { $fail++; "FAIL  Join-Path" }

Write-Output "== 进程与服务 =="
Get-Process 2>$null >$null; if ($?) { $pass++; "PASS  Get-Process" } else { $fail++; "FAIL  Get-Process" }
Start-Sleep -Milliseconds 10 2>$null >$null; if ($?) { $pass++; "PASS  Start-Sleep" } else { $fail++; "FAIL  Start-Sleep" }
Get-Service 2>$null >$null; if ($?) { "PASS  Get-Service" } else { "SKIP  Get-Service (需 systemd)" }
Test-Connection -Count 1 127.0.0.1 2>$null >$null; if ($?) { "PASS  Test-Connection" } else { "SKIP  Test-Connection (需 ping)" }

Write-Output "== 系统信息 =="
Get-Date 2>$null >$null; if ($?) { $pass++; "PASS  Get-Date" } else { $fail++; "FAIL  Get-Date" }
Get-Date -Format "yyyy-MM-dd" 2>$null >$null; if ($?) { $pass++; "PASS  Get-Date -Format" } else { $fail++; "FAIL  Get-Date -Format" }
Get-Uptime 2>$null >$null; if ($?) { $pass++; "PASS  Get-Uptime" } else { $fail++; "FAIL  Get-Uptime" }
Get-ComputerInfo 2>$null >$null; if ($?) { $pass++; "PASS  Get-ComputerInfo" } else { $fail++; "FAIL  Get-ComputerInfo" }
Get-TimeZone 2>$null >$null; if ($?) { $pass++; "PASS  Get-TimeZone" } else { $fail++; "FAIL  Get-TimeZone" }
Get-Culture 2>$null >$null; if ($?) { $pass++; "PASS  Get-Culture" } else { $fail++; "FAIL  Get-Culture" }

Write-Output "== 输出与格式化 =="
Write-Output x 2>$null >$null; if ($?) { $pass++; "PASS  Write-Output" } else { $fail++; "FAIL  Write-Output" }
Write-Host -NoNewline x 2>$null >$null; if ($?) { $pass++; "PASS  Write-Host" } else { $fail++; "FAIL  Write-Host" }
Write-Warning w 2>$null >$null; if ($?) { $pass++; "PASS  Write-Warning" } else { $fail++; "FAIL  Write-Warning" }
Write-Verbose v 2>$null >$null; if ($?) { $pass++; "PASS  Write-Verbose" } else { $fail++; "FAIL  Write-Verbose" }
Write-Information i 2>$null >$null; if ($?) { $pass++; "PASS  Write-Information" } else { $fail++; "FAIL  Write-Information" }
Write-Debug d 2>$null >$null; if ($?) { $pass++; "PASS  Write-Debug" } else { $fail++; "FAIL  Write-Debug" }
Get-ChildItem -Name | Out-File test/tmp/o1.txt 2>$null >$null; if ($?) { $pass++; "PASS  Out-File" } else { $fail++; "FAIL  Out-File" }
1..100 | Out-Null 2>$null >$null; if ($?) { $pass++; "PASS  Out-Null" } else { $fail++; "FAIL  Out-Null" }
Get-Date | Out-String 2>$null >$null; if ($?) { $pass++; "PASS  Out-String" } else { $fail++; "FAIL  Out-String" }
Get-ChildItem | Format-Table -AutoSize 2>$null >$null; if ($?) { $pass++; "PASS  Format-Table" } else { $fail++; "FAIL  Format-Table" }
Get-Item test/核验.ps1 | Format-List 2>$null >$null; if ($?) { $pass++; "PASS  Format-List" } else { $fail++; "FAIL  Format-List" }
Get-ChildItem -Name | Format-Wide 2>$null >$null; if ($?) { $pass++; "PASS  Format-Wide" } else { $fail++; "FAIL  Format-Wide" }
"abc" | Format-Hex 2>$null >$null; if ($?) { $pass++; "PASS  Format-Hex" } else { $fail++; "FAIL  Format-Hex" }

Write-Output "== 对象管道 =="
1..10 | Where-Object { $_ -gt 5 } 2>$null >$null; if ($?) { $pass++; "PASS  Where-Object" } else { $fail++; "FAIL  Where-Object" }
Get-ChildItem | Where-Object Length -gt 0 2>$null >$null; if ($?) { $pass++; "PASS  Where-Object 裸属性" } else { $fail++; "FAIL  Where-Object 裸属性" }
1..3 | ForEach-Object { $_ * 2 } 2>$null >$null; if ($?) { $pass++; "PASS  ForEach-Object" } else { $fail++; "FAIL  ForEach-Object" }
1..10 | Select-Object -First 3 2>$null >$null; if ($?) { $pass++; "PASS  Select-Object" } else { $fail++; "FAIL  Select-Object" }
Get-ChildItem | Select-Object Name,Length 2>$null >$null; if ($?) { $pass++; "PASS  Select-Object 属性" } else { $fail++; "FAIL  Select-Object 属性" }
3,1,2 | Sort-Object 2>$null >$null; if ($?) { $pass++; "PASS  Sort-Object" } else { $fail++; "FAIL  Sort-Object" }
Get-ChildItem | Sort-Object Length -Descending 2>$null >$null; if ($?) { $pass++; "PASS  Sort-Object -Desc" } else { $fail++; "FAIL  Sort-Object -Desc" }
1,1,2,2,3 | Group-Object 2>$null >$null; if ($?) { $pass++; "PASS  Group-Object" } else { $fail++; "FAIL  Group-Object" }
Get-ChildItem | Measure-Object Length -Sum 2>$null >$null; if ($?) { $pass++; "PASS  Measure-Object" } else { $fail++; "FAIL  Measure-Object" }
Measure-Command { Start-Sleep -Milliseconds 5 } 2>$null >$null; if ($?) { $pass++; "PASS  Measure-Command" } else { $fail++; "FAIL  Measure-Command" }
Get-Date | Get-Member 2>$null >$null; if ($?) { $pass++; "PASS  Get-Member" } else { $fail++; "FAIL  Get-Member" }
1,1,2,2,3 | Get-Unique 2>$null >$null; if ($?) { $pass++; "PASS  Get-Unique" } else { $fail++; "FAIL  Get-Unique" }
Compare-Object (1,2,3) (2,3,4) 2>$null >$null; if ($?) { $pass++; "PASS  Compare-Object" } else { $fail++; "FAIL  Compare-Object" }
Get-ChildItem -Name | Tee-Object test/tmp/tee.txt 2>$null >$null; if ($?) { $pass++; "PASS  Tee-Object" } else { $fail++; "FAIL  Tee-Object" }
Get-Date | Add-Member -Name 标签 -Value 测试 2>$null >$null; if ($?) { $pass++; "PASS  Add-Member" } else { $fail++; "FAIL  Add-Member" }

Write-Output "== 运算符（7.X） =="
$r = 5 -in 1..10; if ($?) { $pass++; "PASS  范围 -in" } else { $fail++; "FAIL  范围 -in" }
$r = 1..10 -contains 5; if ($?) { $pass++; "PASS  范围 -contains" } else { $fail++; "FAIL  范围 -contains" }
$r = 1..3 -gt 1; if ($?) { $pass++; "PASS  范围优先于比较" } else { $fail++; "FAIL  范围优先于比较" }
$r = $null ?? "d"; if ($?) { $pass++; "PASS  空合并 ??" } else { $fail++; "FAIL  空合并 ??" }
$r = $true ? "y" : "n"; if ($?) { $pass++; "PASS  三元 ?:" } else { $fail++; "FAIL  三元 ?:" }
$r = "{0}:{1}" -f 1,2; if ($?) { $pass++; "PASS  格式 -f" } else { $fail++; "FAIL  格式 -f" }

Write-Output "== 数据转换 =="
ConvertTo-Json @{a=1} 2>$null >$null; if ($?) { $pass++; "PASS  ConvertTo-Json" } else { $fail++; "FAIL  ConvertTo-Json" }
ConvertFrom-Json '{"a":1}' 2>$null >$null; if ($?) { $pass++; "PASS  ConvertFrom-Json" } else { $fail++; "FAIL  ConvertFrom-Json" }
Get-ChildItem | ConvertTo-Csv 2>$null >$null; if ($?) { $pass++; "PASS  ConvertTo-Csv" } else { $fail++; "FAIL  ConvertTo-Csv" }
"a,b`n1,2" | ConvertFrom-Csv 2>$null >$null; if ($?) { $pass++; "PASS  ConvertFrom-Csv" } else { $fail++; "FAIL  ConvertFrom-Csv" }
ConvertFrom-StringData "a=1" 2>$null >$null; if ($?) { $pass++; "PASS  ConvertFrom-StringData" } else { $fail++; "FAIL  ConvertFrom-StringData" }
Test-Json '{"a":1}' 2>$null >$null; if ($?) { $pass++; "PASS  Test-Json" } else { $fail++; "FAIL  Test-Json" }
Get-Random -Minimum 1 -Maximum 10 2>$null >$null; if ($?) { $pass++; "PASS  Get-Random" } else { $fail++; "FAIL  Get-Random" }
New-Guid 2>$null >$null; if ($?) { $pass++; "PASS  New-Guid" } else { $fail++; "FAIL  New-Guid" }
New-TimeSpan -Minutes 5 2>$null >$null; if ($?) { $pass++; "PASS  New-TimeSpan" } else { $fail++; "FAIL  New-TimeSpan" }
$f = New-TemporaryFile; Remove-Item $f.FullName 2>$null >$null; if ($?) { $pass++; "PASS  New-TemporaryFile" } else { $fail++; "FAIL  New-TemporaryFile" }
1..3 | Join-String -Separator "," 2>$null >$null; if ($?) { $pass++; "PASS  Join-String" } else { $fail++; "FAIL  Join-String" }

Write-Output "== 变量/别名/历史 =="
Get-Variable HOME 2>$null >$null; if ($?) { $pass++; "PASS  Get-Variable" } else { $fail++; "FAIL  Get-Variable" }
Set-Variable -Name vv -Value 1 2>$null >$null; if ($?) { $pass++; "PASS  Set-Variable" } else { $fail++; "FAIL  Set-Variable" }
New-Variable -Name nv1 -Value 1 -Force 2>$null >$null; if ($?) { $pass++; "PASS  New-Variable" } else { $fail++; "FAIL  New-Variable" }
Remove-Variable nv1 2>$null >$null; if ($?) { $pass++; "PASS  Remove-Variable" } else { $fail++; "FAIL  Remove-Variable" }
Clear-Variable vv 2>$null >$null; if ($?) { $pass++; "PASS  Clear-Variable" } else { $fail++; "FAIL  Clear-Variable" }
Get-Alias ls 2>$null >$null; if ($?) { $pass++; "PASS  Get-Alias" } else { $fail++; "FAIL  Get-Alias" }
Set-Alias vtest Get-Content 2>$null >$null; if ($?) { $pass++; "PASS  Set-Alias" } else { $fail++; "FAIL  Set-Alias" }
Remove-Alias vtest 2>$null >$null; if ($?) { $pass++; "PASS  Remove-Alias" } else { $fail++; "FAIL  Remove-Alias" }
Export-Alias test/tmp/al.txt 2>$null >$null; if ($?) { $pass++; "PASS  Export-Alias" } else { $fail++; "FAIL  Export-Alias" }
Import-Alias test/tmp/al.txt 2>$null >$null; if ($?) { $pass++; "PASS  Import-Alias" } else { $fail++; "FAIL  Import-Alias" }
Get-History 2>$null >$null; if ($?) { $pass++; "PASS  Get-History" } else { $fail++; "FAIL  Get-History" }
Add-History "echo x" 2>$null >$null; if ($?) { $pass++; "PASS  Add-History" } else { $fail++; "FAIL  Add-History" }
Invoke-Expression "1 + 2" 2>$null >$null; if ($?) { $pass++; "PASS  Invoke-Expression" } else { $fail++; "FAIL  Invoke-Expression" }
Get-Command Get-Content 2>$null >$null; if ($?) { $pass++; "PASS  Get-Command" } else { $fail++; "FAIL  Get-Command" }
Get-Help Get-Content 2>$null >$null; if ($?) { $pass++; "PASS  Get-Help" } else { $fail++; "FAIL  Get-Help" }
Invoke-History "echo 历史重放" 2>$null >$null; if ($?) { $pass++; "PASS  Invoke-History" } else { $fail++; "FAIL  Invoke-History" }

Write-Output "== 链与重定向 =="
echo a && echo b 2>$null >$null; if ($?) { $pass++; "PASS  && 链" } else { $fail++; "FAIL  && 链" }
Test-Path /nonexist || echo fallback 2>$null >$null; if ($?) { $pass++; "PASS  || 链" } else { $fail++; "FAIL  || 链" }
Get-Content /nonexist 2>$null || echo ok 2>$null >$null; if ($?) { $pass++; "PASS  2>\$null" } else { $fail++; "FAIL  2>\$null" }
Write-Output x > test/tmp/r1.txt 2>$null >$null; if ($?) { $pass++; "PASS  > 重定向" } else { $fail++; "FAIL  > 重定向" }
Write-Output y >> test/tmp/r1.txt 2>$null >$null; if ($?) { $pass++; "PASS  >> 追加" } else { $fail++; "FAIL  >> 追加" }
Write-Output x > $null 2>$null >$null; if ($?) { $pass++; "PASS  > \$null" } else { $fail++; "FAIL  > \$null" }

Write-Output ""
Write-Output ("结果: 通过 " + $pass + "  失败 " + $fail)

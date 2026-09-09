# 功能补全。
Write-Output "== 功能补全 =="

# 81. Profile 启动加载（写启动脚本，子进程执行后检查标记文件，末尾删除配置目录）
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

# 7.X 命令格式测试（用 -Version 7 运行）
Write-Output "== PowerShell 7.X 命令格式 =="
Write-Output "PSVersionTable: PSEdition = $($PSVersionTable.PSEdition) / Platform = $($PSVersionTable.Platform)"
Write-Output "IsLinux 存在: $($null -ne (Get-Variable IsLinux -ErrorAction SilentlyContinue).Value)"
Write-Output "自动变量 IsWindows: $IsWindows"

Write-Output "== 对象管道 =="
Get-ChildItem -Name | Where-Object { $_ -match "\.go$" } | Select-Object -First 3

Write-Output "== 变量与字符串 =="
$greeting = "世界"
Write-Output "你好，$greeting！"

Write-Output "== 控制流 =="
foreach ($i in 1..3) { Write-Output "项 $i" }

Write-Output "== 退出码 =="
Write-Output "完成"

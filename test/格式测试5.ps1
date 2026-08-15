# 5.X 命令格式测试（用 -Version 5.1 运行）
Write-Output "== Windows PowerShell 5.X 命令格式 =="
Write-Output "PSVersionTable: PSEdition = $($PSVersionTable.PSEdition)"

Write-Output "== 5.X 独有别名 =="
Get-Alias sc, curl, wget | Format-Table -AutoSize

Write-Output "== 伪 Windows 路径提示符 =="
Write-Output "当前目录: $PWD"

Write-Output "== 5.X 无 \$IsLinux（应为空） =="
Write-Output "IsLinux: [$IsLinux]"

Write-Output "== 完成 =="

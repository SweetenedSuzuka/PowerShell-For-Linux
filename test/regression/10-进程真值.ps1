# Get-Process 真值，仅 Linux。
Write-Output "== Get-Process 真值（仅 Linux） =="

# 102. 进程 CPU/Memory 来自 /proc，pid 1 常驻进程两者应大于 0
$p1 = Get-Process | Where-Object { $_.Id -eq 1 }
$results += T "Get-Process 真值" (($p1.CPU -gt 0) -and ($p1.Memory -gt 0))

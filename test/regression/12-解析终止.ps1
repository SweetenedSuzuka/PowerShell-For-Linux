# 解析终止性。
Write-Output "== 解析终止性 =="

# 107. switch 与 { 之间换行的写法可正常解析执行
$nlSrc = "switch (2)`n{ default { 42 } }"
$nlOut = sh -c "echo '$nlSrc' | $root/powershell -NoLogo -NoProfile -Command -" 2>$null
$results += T "switch 换行大括号可解析" (($nlOut -join "") -eq "42")

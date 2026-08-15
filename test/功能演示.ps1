# 功能演示：把常用语言特性跑一遍，人工看输出是否正常
Write-Output "== 变量与环境 =="
$env:MY_TEST_VAR = "abc"
Write-Output "env: $env:MY_TEST_VAR"
Write-Output "PWD: $PWD"
Write-Output "PID: $PID"

Write-Output "== 字符串方法 =="
"Hello World".ToLower()
"  trim me  ".Trim()

Write-Output "== 比较与数组 =="
Write-Output (1,2,3,4,5 -gt 2)
Write-Output ("apple","banana","cherry" -like "b*")
Write-Output (5 -in 1..10)

Write-Output "== 7.X 运算符 =="
Write-Output ($null ?? "默认值")
Write-Output ($true ? "是" : "否")
Write-Output ("共 {0} 个文件" -f 3)

Write-Output "== 格式化 =="
Get-ChildItem -Name | Select-Object -First 3

Write-Output "== 哈希表显示 =="
@{ Name="样例"; Level=100 }

Write-Output "== 系统信息 =="
Get-Uptime
(Get-Process).Count

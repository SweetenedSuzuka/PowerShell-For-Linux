# 对象方法。
Write-Output "== 对象方法 =="

# 76. DateTime ToShortDateString（Get-Date -Date 指定日期）
$dts = (Get-Date -Date "2020-01-15").ToShortDateString()
$results += T "DateTime ToShortDateString" (($dts -eq "1/15/2020"))
# 77. FileInfo 虚拟属性 Extension（Get-Item 返回对象的路径派生属性）
"x" | Set-Content ov.txt
$fie = (Get-Item ov.txt).Extension
$results += T "FileInfo Extension" (($fie -eq ".txt"))
# 78. 字符串 PadLeft（总宽,填充字符）
$results += T "字符串 PadLeft" (("7".PadLeft(3,"0")) -eq "007")
# 79. TrimStart 无参清前导空白
$ts = "  x  ".TrimStart()
$results += T "TrimStart 空白" (($ts -eq "x  "))
# 80. Split 无参按任意空白分割（tab/空格/换行，连续空白合并）
$sp3 = ("a`tb c").Split()
$results += T "Split 无参空白" (($sp3.Count -eq 3))

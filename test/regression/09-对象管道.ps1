# 对象管道行为一致。
Write-Output "== 对象管道行为一致 =="

# 89. Select-String 默认大小写不敏感（匹配大小写变体）
$ssn = "Hello","HELLO" | Select-String "hello"
$results += T "Select-String 默认不敏感" (($ssn.Count -eq 2))
# 90. Select-String -CaseSensitive 只认精确大小写
$ssc = "Hello","hello" | Select-String "hello" -CaseSensitive
$results += T "Select-String CaseSensitive" (($ssc.Count -eq 1) -and ($ssc[0].Line -eq "hello"))
# 91. Select-String 管道编号按对象序号，未匹配对象占号（x 不匹配，b/c 编号 2/3）
$ssl = "x","b","c" | Select-String "[bc]" | ForEach-Object { "$($_.LineNumber):$($_.Line)" }
$results += T "Select-String 管道编号" (($ssl -join ",") -eq "2:b,3:c")
# 92. Group-Object 默认大小写不敏感合并，Name 取出首次原值，合并后只有一组
$gn = "apple","Apple" | Group-Object
$results += T "Group-Object 默认合并" ((@($gn).Count -eq 1) -and ($gn[0].Name -eq "apple") -and ($gn[0].Count -eq 2))
# 93. Group-Object -CaseSensitive 分开分组
$gc3 = ("apple","Apple" | Group-Object -CaseSensitive)
$results += T "Group-Object CaseSensitive" (($gc3.Count -eq 2))
# 94. Compare-Object 默认大小写不敏感，先右后左输出（b 两边共有被排除）
$co2 = Compare-Object -ReferenceObject a,B -DifferenceObject b,C
$coSeq = $co2 | ForEach-Object { $_.SideIndicator + $_.InputObject }
$results += T "Compare-Object 默认顺序" (($coSeq -join ",") -eq "=>C,<=a")
# 95. Compare-Object -IncludeEqual 相等项显示基准集的值（A 与 a 相等显示 A）
$cie = Compare-Object -ReferenceObject A -DifferenceObject a -IncludeEqual
$results += T "Compare-Object IncludeEqual值" (($cie[0].SideIndicator -eq "==") -and ($cie[0].InputObject -eq "A"))
# 96. Sort-Object -Unique 大小写不敏感去重（apple/Apple 合并为排序后首个）
$su = "banana","apple","Apple" | Sort-Object -Unique
$results += T "Sort-Object Unique去重" (($su -join ",") -eq "apple,banana")
# 97. Sort-Object -Unique -CaseSensitive 全保留
$suc = ("apple","Apple","banana" | Sort-Object -Unique -CaseSensitive).Count
$results += T "Sort-Object Unique CS" (($suc -eq 3))
# 98. Select-Object -First 0 返回空
$f0 = 1,2,3 | Select-Object -First 0
$l0 = 1,2,3 | Select-Object -Last 0
$results += T "Select-Object First/Last 0" ((($f0 -eq $null)) -and ($l0 -eq $null))
# 99. Measure-Object 未指定的统计字段为 $null（字段总是补全）
$m9 = 1,2,3 | Measure-Object
$results += T "Measure-Object 字段补全" (($m9.Sum -eq $null) -and ($m9.Average -eq $null) -and ($m9.Count -eq 3))
# 100. Measure-Object Sum/Average 遇非数字作废，字段置为 $null
$mn = "a","b" | Measure-Object -Sum
$mx = 1,"a",2 | Measure-Object -Sum
$results += T "Measure-Object 非数字作废" ((($mn.Sum -eq $null)) -and ($mx.Sum -eq $null))
# 101. Measure-Object -Property 模式 Count 只数有该属性的对象
$mp = @{a=1},@{a=2},@{b=3} | Measure-Object -Property a -Sum
$results += T "Measure-Object Property Count" (($mp.Count -eq 2) -and ($mp.Sum -eq 3))

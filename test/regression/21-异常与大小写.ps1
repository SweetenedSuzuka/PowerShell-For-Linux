# 异常抛出与大小写。
# 182. 无效算术抛出错误可捕获
$ec1 = try { "a" - "b" } catch { "caught" }
$results += T "无效算术抛出错误" ($ec1 -eq "caught")
# 183. 非数字取负抛出错误可捕获
$ec2 = try { -"abc" } catch { "caught" }
$results += T "取负抛出错误" ($ec2 -eq "caught")
# 184. 越界截取抛出错误可捕获
$ec3 = try { "abc".Substring(1,10) } catch { "caught" }
$results += T "截取越界抛出错误" ($ec3 -eq "caught")
# 185. 过滤器内抛出错误向外传播
$ec4 = try { 1,2 | ForEach-Object { throw "x" } } catch { "caught" }
$results += T "ForEach 抛出错误传播" ($ec4 -eq "caught")
# 186. while 条件每轮求值一次
$wl = 0
while ($wl++ -lt 2) { }
$results += T "while 每轮求值" ($wl -eq 3)
# 187. 后缀自增返回旧值
$po = 0
$pv = $po++
$results += T "后缀取旧值" ((($pv -eq 0) -and ($po -eq 1)))
# 188. 变量读取与写入不区分大小写
$vcFoo = 1
$VCFOO = 2
$results += T "变量大小写" ((($VCFOO -eq 2) -and ($VcFoo -eq 2)))
# 189. 自动变量大小写
$results += T "自动变量大小写" ((($pwd -ne $null) -and (($PID -gt 0) -and ($psedition -ne ""))))
# 190. Get-Variable 不区分大小写
New-Variable -Name GvC -Value 3
$results += T "取变量大小写" (((Get-Variable gvc).Value -eq 3))
Remove-Variable GVC
$results += T "删变量大小写" ((Get-Variable gvc) -eq $null)
# 191. 作用域修饰符不区分大小写
$global:SCV = 4
$results += T "作用域大小写" ($GLOBAL:SCV -eq 4)

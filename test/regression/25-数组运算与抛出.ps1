# 数组乘法、抛出前输出与问号变量。
# 220. 数组乘以整数时整体重复该次数
$am1 = @(1,2,3) * 2
$am2 = @("a","b") * 3
$am3 = @(1,2,3) * 1
$results += T "数组乘以整数整体重复" (((($am1 -join ",") -eq "1,2,3,1,2,3")) -and ((($am2 -join ",") -eq "a,b,a,b,a,b")) -and ((($am3 -join ",") -eq "1,2,3")))
# 221. 数组乘以零与空数组相乘得到空数组
$az = @(1,2,3) * 0
$ae = @() * 2
$results += T "数组乘零得到空数组" (((($null -eq $az) -eq $false)) -and (($az.Count -eq 0)) -and ((($null -eq $ae) -eq $false)) -and (($ae.Count -eq 0)))
# 222. 非法乘数抛出的错误可捕获
$an1 = try { @(1,2,3) * -1; "no" } catch { "caught" }
$an2 = try { 2 * @(1,2,3); "no" } catch { "caught" }
$results += T "数组非法乘数抛出错误" ((($an1 -eq "caught")) -and (($an2 -eq "caught")))
# 223. 脚本块内抛出错误前已产生的输出保留
$so1 = try { & { "s1a"; throw "s1x" } } catch { "s1caught" }
$results += T "脚本块抛出前输出保留" ((($so1 -join ",") -eq "s1a,s1caught"))
# 224. 函数内抛出错误前已产生的输出保留
function OFunc { "f2a"; throw "f2x" }
$so2 = try { OFunc } catch { "f2caught" }
$results += T "函数抛出前输出保留" ((($so2 -join ",") -eq "f2a,f2caught"))
# 225. 进入 catch 时问号变量保持失败
$qc = try { throw "t1" } catch { if ($?) { "true" } else { "false" } }
$results += T "进入catch时问号保持失败" (($qc -eq "false"))
# 226. 错误上抛经过 finally 时问号变量保持失败
$qf = try { try { throw "t2" } finally { if ($?) { "f-true" } else { "f-false" } } } catch { "outer" }
$results += T "经过finally时问号保持失败" ((($qf -join ",") -eq "f-false,outer"))
# 227. 空 catch 块之后问号变量保持失败
try { throw "t5" } catch { }
if ($?) { $qa = "true" } else { $qa = "false" }
$results += T "空catch之后问号保持失败" (($qa -eq "false"))

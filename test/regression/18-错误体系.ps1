# 错误体系：错误记录、错误动作与首选项。
Write-Output "== \$Error 自动变量 =="

# 137. 非终止错误累积，[0] 是最新
[int]"err-a"
[int]"err-b"
$results += T '$Error 累积与顺序' (($Error.Count -ge 2) -and ($Error[0].Message -like "*err-b*"))
# 138. throw 与被捕获的错误都进 $Error
$ec1 = $Error.Count
try { throw "ev-x" } catch {}
$results += T "throw 进 Error" (($Error.Count -eq $ec1 + 1) -and ($Error[0].Message -eq "ev-x"))
# 139. Clear 清空本体
$Error.Clear()
$results += T 'Error.Clear 清空' (($Error.Count -eq 0))
# 140. RemoveAt 删除指定下标
[int]"err-c"
[int]"err-d"
$Error.RemoveAt(0)
$results += T "Error.RemoveAt" (($Error.Count -eq 1) -and ($Error[0].Message -like "*err-c*"))
$Error.Clear()
# 141. -ErrorAction 默认继续：记录并继续
$ea0 = $Error.Count
Get-Item 不存在EA123
$results += T "ErrorAction 默认继续" (($Error.Count -eq $ea0 + 1))
# 142. SilentlyContinue：记录但不显示
$ea1 = $Error.Count
Get-Item 不存在EA123 -ErrorAction SilentlyContinue
$results += T "ErrorAction SilentlyContinue 记录" (($Error.Count -eq $ea1 + 1))
# 143. Ignore：不记录
$ea2 = $Error.Count
Get-Item 不存在EA123 -ErrorAction Ignore
$results += T "ErrorAction Ignore 不记录" (($Error.Count -eq $ea2))
# 144. Stop：转终止错误可捕获，只记一次
$ea3 = $Error.Count
$eaCaught = try { Get-Item 不存在EA123 -ErrorAction Stop; "no" } catch { "yes" }
$results += T "ErrorAction Stop 可捕获" (($eaCaught -eq "yes") -and ($Error.Count -eq $ea3 + 1))
# 145. Stop 中断 try 体后续语句
$eaAfter = try { Get-Item 不存在EA123 -ErrorAction Stop; "after" } catch { "caught" }
$results += T "ErrorAction Stop 中断后续" ($eaAfter -eq "caught")
# 146. 无效取值报绑定错误
$ea4 = $Error.Count
Get-Item foo -ErrorAction Bogus
$results += T "ErrorAction 无效值报错" (($Error.Count -eq $ea4 + 1))
# 147. Inquire 在非交互场景按终止错误处理
$ea5 = $Error.Count
$eaIn = try { Get-Item 不存在EA123 -ErrorAction Inquire; "no" } catch { "yes" }
$results += T "ErrorAction Inquire 终止" (($eaIn -eq "yes") -and ($Error.Count -eq $ea5 + 1))
# 148. $ErrorActionPreference 默认值
$results += T "ErrorActionPreference 默认值" ($ErrorActionPreference -eq "Continue")
# 149. 首选项 SilentlyContinue 记录
$ep0 = $Error.Count
$ErrorActionPreference = 'SilentlyContinue'
Get-Item 不存在EA123
$results += T "首选项 SilentlyContinue 记录" (($Error.Count -eq $ep0 + 1))
# 150. 首选项 Stop 可捕获
$ep1 = $Error.Count
$ErrorActionPreference = 'Stop'
$epCaught = try { Get-Item 不存在EA123; "no" } catch { "yes" }
$results += T "首选项 Stop 可捕获" (($epCaught -eq "yes") -and ($Error.Count -eq $ep1 + 1))
# 151. 显式参数覆盖首选项
$ep2 = $Error.Count
Get-Item 不存在EA123 -ErrorAction Continue
$results += T "显式参数覆盖首选项" (($Error.Count -eq $ep2 + 1))
# 152. 函数内首选项只在局部生效
$ErrorActionPreference = 'Continue'
function PrefScope { $ErrorActionPreference = 'Stop'; try { Get-Item 不存在EA123 } catch { "caught" } }
$epScope = PrefScope
$results += T "首选项函数局部生效" (($epScope -eq "caught") -and ($ErrorActionPreference -eq "Continue"))
# 153. 无效首选项赋值报错且不生效
$ep3 = $Error.Count
$ErrorActionPreference = 'Bogus'
$results += T "首选项无效值报错" (($Error.Count -eq $ep3 + 1) -and ($ErrorActionPreference -eq "Continue"))
# 154. 首选项名大小写混写仍生效
$ErrorActionPreference = 'Continue'
$erroractionpreference = 'Stop'
$epCase = try { Get-Item 不存在EA123; "no" } catch { "yes" }
$results += T "首选项大小写写入" ($epCase -eq "yes")
$ErrorActionPreference = 'Continue'
# 155. 空值恢复默认
$ErrorActionPreference = 'Stop'
$ErrorActionPreference = $null
$results += T "首选项空值恢复默认" ($ErrorActionPreference -eq "Continue")
# 156. 赋值右侧读取 $? 拿到旧状态
Get-Item 不存在EA123
$aqRead = $?
$results += T "赋值右侧读取旧状态" ($aqRead -eq $false)
# 157. 右侧无新错误时赋值后 $? 为真
$aqTmp = 5
if ($?) { $aqOk = "ok" } else { $aqOk = "fail" }
$results += T "失败后赋值 $? 为真" ($aqOk -eq "ok")
# 158. 赋值右侧出错保持失败
$aqBad = 1/0
if ($?) { $aqBadOk = "ok" } else { $aqBadOk = "fail" }
$results += T "赋值右侧出错保持失败" (($aqBad -eq $null) -and ($aqBadOk -eq "fail"))
# 159. 被捕获的错误不影响赋值成功
Get-Item 不存在EA123
$aqCatch = try { throw "aq-x" } catch { "got" }
if ($?) { $aqCatchOk = "ok" } else { $aqCatchOk = "fail" }
$results += T "捕获后赋值置真" (($aqCatch -eq "got") -and ($aqCatchOk -eq "ok"))
# 160. 首选项 Stop 让除零可捕获
$em0 = $Error.Count
$ErrorActionPreference = 'Stop'
$emDiv = try { 5/0; "no" } catch { "yes" }
$results += T "首选项 Stop 除零可捕获" (($emDiv -eq "yes") -and ($Error.Count -eq $em0 + 1))
# 161. 首选项 Stop 让类型转换可捕获
$em1 = $Error.Count
$emCast = try { [int]"abc"; "no" } catch { "yes" }
$results += T "首选项 Stop 转换可捕获" (($emCast -eq "yes") -and ($Error.Count -eq $em1 + 1))
# 162. 默认除零继续
$ErrorActionPreference = 'Continue'
$em2 = $Error.Count
$emDef = 1/0
$results += T "默认除零继续" (($emDef -eq $null) -and ($Error.Count -eq $em2 + 1))

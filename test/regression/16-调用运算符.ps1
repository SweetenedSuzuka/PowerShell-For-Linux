# 调用运算符。
Write-Output "== 调用运算符 =="

# 126. & 直接调用脚本块与变量中的脚本块
$ic1 = & { "a"; "b" }
$results += T "& 直调脚本块" (($ic1 -join ",") -eq "a,b")
$sb = { param($x) $x * 2 }
$results += T "& 变量持块带 param" ((& $sb 21) -eq 42)
# 127. 多余实参进 $args
$sb2 = { param($x) "x=$x args=$args" }
$results += T "& 多余实参进 args" ((& $sb2 1 2 3) -eq "x=1 args=2 3")
# 128. 作表达式（赋值右侧）与 return 顺序
$iv = & { 6 * 7 }
$results += T "& 作表达式" (($iv -eq 42))
$iw = & { "before"; return "val"; "after" }
$results += T "& return 顺序" (($iw -join ",") -eq "before,val")
# 129. 块内 throw 可被调用方捕获
$ic = ""
try { & { throw "blk" } } catch { $ic = "got:$($_.Message)" }
$results += T "& throw 可捕获" (($ic -eq "got:blk"))
# 130. 变量中的命令名按名字分发与 .Invoke
$cmdName = "Write-Output"
$results += T "& 按名字分发" ((& $cmdName hi) -eq "hi")
$sbi = { 1; 2 }
$rIv = $sbi.Invoke()
$results += T ".Invoke 收集输出" (($rIv.Count -eq 2) -and ($rIv[0] -eq 1) -and ($rIv[1] -eq 2))

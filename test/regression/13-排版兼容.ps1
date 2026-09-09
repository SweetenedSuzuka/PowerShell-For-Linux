# 排版兼容。
Write-Output "== 排版兼容 =="

# 108. 块大括号允许另起一行书写（if/elseif/else）
$mIf = ""
if ($true)
{
    $mIf = "then"
}
elseif ($false)
{
    $mIf = "elseif"
}
else
{
    $mIf = "else"
}
$results += T "大括号换行 if/elseif/else" (($mIf -eq "then"))
# 109. try/catch/finally 的大括号换行书写
$mCatch = ""
$mFinally = ""
try
{
    throw "e"
}
catch
{
    $mCatch = "caught"
}
finally
{
    $mFinally = "fin"
}
$results += T "大括号换行 try/catch/finally" (($mCatch -eq "caught") -and ($mFinally -eq "fin"))
# 110. do/while 的 while 允许写在 } 的下一行
$mDo = 0
do
{
    $mDo++
}
while ($mDo -lt 3)
$results += T "大括号换行 do/while" (($mDo -eq 3))
# 111. 行尾运算符与括号内的表达式跨行续写
$mSum = 1 +
    2 +
    3
$mParen = (
    10 +
    20
)
$results += T "表达式跨行续行" (($mSum -eq 6) -and ($mParen -eq 30))
# 112. 链式运算符 && 允许写在行尾跨行续写
$mFile = "test/tmp/chain.txt"
New-Item -ItemType Directory -Force test/tmp >$null
Remove-Item $mFile -Force 2>$null
Set-Content $mFile "L" &&
    Add-Content $mFile "R"
$results += T "链式运算符跨行" (((Get-Content $mFile) -join ",") -eq "L,R")
Remove-Item $mFile -Force 2>$null

# 回归测试共用前置项：结果收集、T 断言、工作目录准备。
$results = @()
function T($name, $cond) { if ($cond) { "PASS  $name" } else { "FAIL  $name" } }
Remove-Item -Recurse -Force test/tmp/reg 2>$null >$null
New-Item -ItemType Directory -Force test/tmp/reg 2>$null >$null
Set-Location test/tmp/reg

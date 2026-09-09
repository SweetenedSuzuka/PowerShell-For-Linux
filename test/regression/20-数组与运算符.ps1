# 数组、运算符与自增写回。
# 171. @() 内裸命令执行（单命令）
$at1 = @(Get-Command -Name Get-Process)
$results += T "@() 执行单命令" ((@($at1).Count -eq 1) -and ($at1[0].Name -eq "Get-Process"))
# 172. @() 全量与 $() 一致
$results += T "@() 全量一致" ((@(Get-Command).Count) -eq ($(Get-Command).Count))
# 173. @() 字面量不受影响
$results += T "@() 字面量" (((@(1,2,3) -join ",")) -eq "1,2,3")
# 174. 比较运算符大小写不敏感
$results += T "大写运算符" ((("a" -EQ "a") -and ("b" -GT "a") -and ("abc" -MATCH "b")))
# 175. i- 显式不敏感变体
$results += T "显式不敏感变体" ((("A" -ieq "a") -and ("AbC" -ilike "a*c") -and ("AbC" -imatch "b") -and ("A" -iin @("a","b"))))
# 176. $ENV: 前缀大小写不限，变量名保持原样
$env:TmVar904 = "v904"
$results += T "ENV 大小写" ((($ENV:TmVar904) -eq "v904") -and (($env:TmVar904) -eq "v904"))
# 177. 下标自增写回同一位置
$iv = 1,2
$iv[0]++
$results += T "下标自增" (($iv[0] -eq 2) -and ($iv[1] -eq 2))
# 178. 属性自增写回
$im = [pscustomobject]@{n=1}
$im.n++
$results += T "属性自增" ($im.n -eq 2)
# 179. 哈希表键自增写回
$ih = @{k=5}
$ih["k"]++
$results += T "哈希键自增" ($ih["k"] -eq 6)
# 180. 多字节字符串按字符下标与长度
$results += T "多字节字符" ((("你好"[1] -eq "好") -and (("你好".Length) -eq 2) -and (("a好b".IndexOf("好")) -eq 1)))
# 181. 多字节截取与查找
$results += T "多字节截取" ((("你好世界".Substring(2,2) -eq "世界") -and (("你好世界".Remove(2) -eq "你好")) -and (("好好".LastIndexOf("好")) -eq 1)))

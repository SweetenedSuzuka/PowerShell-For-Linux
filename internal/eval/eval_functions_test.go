package eval

import (
	"io"
	"os"
	"path/filepath"
	"powershell/internal/object"
	"powershell/internal/shell"
	"strings"
	"testing"
)

func TestFunctions(t *testing.T) {
	src := `
function Add($a, $b = 1) { return $a + $b }
Add 5
Add 5 10
`
	wantStr(t, src, "6", "15")
}

// TestParamBlock 验证函数 param() 声明块（默认值、多余实参进 $args、旧括号语法共存）。
func TestParamBlock(t *testing.T) {
	// param() 块形式
	wantStr(t, "function Fp { param($x) $x * 2 }; Fp 21", "42")
	// 默认值
	wantStr(t, "function Fd { param($x, $y = 100) $x + $y }; Fd 5; Fd 5 7", "105", "12")
	// 多余实参进 $args
	wantStr(t, `function Fa { param($x) "x=$x rest=$args" }; Fa 1 2 3`, "x=1 rest=2 3")
	// 旧括号语法仍可用
	wantStr(t, "function Old($a, $b) { $a + $b }; Old 1 2", "3")
	// 顶层 param() 不是脚本/函数开头时执行报错（无输出）
	wantStr(t, "param($z)")
}

// TestParamTypeAnnotations 验证形参 [类型] 标注：
// 位置与命名实参、默认值都经类型转换，数组标注把单值包装成单元素数组，无法转换时不执行被调方。
func TestParamTypeAnnotations(t *testing.T) {
	// 位置实参转换成声明类型
	wantStr(t, `function Fti([int]$x) { $x -is [int] }; Fti '42'`, "True")
	// 命名实参同样转换
	wantStr(t, `function Ftn([int]$n) { $n -is [int] }; Ftn -n '33'`, "True")
	// 默认值也经类型转换
	wantStr(t, `function Ftd([int]$y = '5') { $y -is [int] }; Ftd`, "True")
	// double 标注
	wantStr(t, `function Ftw([double]$v) { $v }; Ftw '2.5'`, "2.5")
	// 数组标注把单值包装成单元素数组
	wantStr(t, `function Fta([int[]]$a) { $a.Count; $a[0] -is [int] }; Fta 5`, "1", "True")
	// 数组实参逐元素转换
	wantStr(t, `function Ftb([string[]]$s) { $s.Count; $s[1] }; Ftb 1,2,3`, "3", "2")
	// 未标注的形参保持原值不转换
	wantStr(t, "function Ftu($x) { $x -is [string] }; Ftu '42'", "True")
	// 无法转换时不执行函数体，后续语句继续
	wantStr(t, `function Ftf([int]$x) { "body" }; Ftf 'abc'; "after"`, "after")
	// 绑定失败把 $? 置为 false
	wantStr(t, `function Ftq([int]$x) { "b" }; Ftq 'abc'; if ($?) { "yes" } else { "no" }`, "no")
}

// TestScriptParamTypeAnnotations 验证脚本 param() 的类型标注：
// 可转换实参转成声明类型，无法转换时脚本主体不再执行并设置失败退出码。
func TestScriptParamTypeAnnotations(t *testing.T) {
	dir := t.TempDir()
	okPath := filepath.Join(dir, "ok.ps1")
	if err := os.WriteFile(okPath, []byte("param([int]$n)\n$n\n$n -is [int]\n"), 0644); err != nil {
		t.Fatal(err)
	}
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := New(sess, strings.NewReader(""), io.Discard, io.Discard)
	out := ev.RunScriptFile(okPath, []*object.PSObject{object.Str("11")})
	if got := strs(out); len(got) != 2 || got[0] != "11" || got[1] != "True" {
		t.Fatalf("脚本类型标注 → %v，想要 [11 True]", got)
	}

	badPath := filepath.Join(dir, "bad.ps1")
	if err := os.WriteFile(badPath, []byte("param([int]$n)\n\"不应输出\"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	out = ev.RunScriptFile(badPath, []*object.PSObject{object.Str("abc")})
	if got := strs(out); len(got) != 0 {
		t.Fatalf("绑定失败脚本应无输出，实际 %v", got)
	}
	if sess.LastExit != 1 {
		t.Fatalf("绑定失败应设置失败退出码 1，实际 %d", sess.LastExit)
	}
}

// TestInvokeOperator 验证 & 调用运算符：
// 脚本块直接调用、变量中的脚本块调用、param 形参、$args、return 顺序、作为表达式、throw 可捕获、动态作用域、按名字调用命令。
func TestInvokeOperator(t *testing.T) {
	// 基本执行与多输出
	wantStr(t, `& { "a"; "b" }`, "a", "b")
	// 变量持有脚本块后调用
	wantStr(t, `$sb = { param($x) $x * 2 }; & $sb 21`, "42")
	// 块体开头 param() 提取为形参，多余实参进 $args
	wantStr(t, `$sb2 = { param($x) "x=$x args=$args" }; & $sb2 1 2 3`, "x=1 args=2 3")
	// 无 param 时实参全进 $args
	wantStr(t, `& { "n=$($args.Count)" } e1 e2`, "n=2")
	// return 前的输出保留，return 后不再执行
	wantStr(t, `& { "before"; return "val"; "after" }`, "before", "val")
	// 作为表达式（赋值右侧）
	wantStr(t, `$v = & { 6 * 7 }; $v`, "42")
	// 块内 throw 可被调用方捕获
	wantStr(t, `try { & { throw "blk" } } catch { "caught: $($_.Message)" }`, "caught: blk")
	// 动态作用域：块内可见外层变量
	wantStr(t, `$outer = 10; & { $outer }`, "10")
	// 管道输入进 $input（$_ 不绑定），输出时逐项枚举
	wantStr(t, `1,2 | & { $input }`, "1", "2")
	// 变量中的命令名按名字分发
	wantStr(t, `$c = "Write-Output"; & $c hi`, "hi")
	// 字符串字面量目标
	wantStr(t, `& 'Write-Output' hi`, "hi")
	// .Invoke 执行并把输出收集成数组
	wantStr(t, `$sb3 = { 1; 2 }; $r = $sb3.Invoke(); $r.Count`, "2")
	wantStr(t, `$sb4 = { param($a,$b) $a + $b }; $sb4.Invoke(5,6)`, "11")
	// 目标缺失报错且后续语句继续
	wantStr(t, `&; "after"`, "after")
}

// TestNamedBlocks 验证函数的 begin/process/end 命名块与 filter：
// 三段顺序、直接调用时 process 以 $null 跑一次、process 逐项绑定 $_、return 只结束本次、零输入跳过 process、块可乱序。
func TestNamedBlocks(t *testing.T) {
	// 三块齐全：begin 一次 → process 逐项 → end 一次
	wantStr(t, `function T1 { begin { "b" } process { "p:$_" } end { "e" } }; 1,2,3 | T1`, "b", "p:1", "p:2", "p:3", "e")
	// 直接调用：无管道输入，process 以 $null 跑一次
	wantStr(t, `function T2 { begin { "b" } process { "p" } end { "e" } }; T2`, "b", "p", "e")
	// 只有 process：管道逐项，直接调用跑一次
	wantStr(t, `function P1 { process { "p:$_" } }; 1,2 | P1`, "p:1", "p:2")
	wantStr(t, `function P2 { process { "p" } }; P2`, "p")
	// 块乱序仍按 begin/process/end 顺序执行
	wantStr(t, `function O1 { end { "e" } begin { "b" } process { "x:$_" } }; 9 | O1`, "b", "x:9", "e")
	// 零输入：begin/end 跑，process 不跑
	wantStr(t, `$a = @(); function Z1 { begin { "zb" } process { "zp:$_" } end { "ze" } }; $r = @($a | Z1); $r.Count`, "2")
	// param 与命名块共存，形参在 process 内可用
	wantStr(t, `function Q1([int]$mul) { process { $_ * $mul } }; 1,2,3 | Q1 -mul 10`, "10", "20", "30")
	// filter 的 Body 即 process
	wantStr(t, `filter F1 { "f:$_" }; 1,2 | F1`, "f:1", "f:2")
	wantStr(t, `filter F2 { "f" }; F2`, "f")
	// filter 也可带命名块
	wantStr(t, `filter F3 { begin { "fb" } process { "fp:$_" } }; 1 | F3`, "fb", "fp:1")
}

// TestNamedBlocksControlFlow 验证命名块内的 return/break/continue：
// process 里 return 只结束本次；break/continue 无所属循环时沿调用栈上抛终止当前语句序列。
func TestNamedBlocksControlFlow(t *testing.T) {
	// process 内 return 跳过本次剩余语句，继续下一项
	wantStr(t, `function R1 { process { if ($_ -eq 2) { return }; "r:$_" } }; 1,2,3 | R1`, "r:1", "r:3")
	// begin 内 return 结束整个函数（end 不再跑）
	wantStr(t, `function R2 { begin { "b"; return } process { "p" } end { "e" } }; R2`, "b")
	// process 内 break 上抛：中断外层循环体，循环后的语句继续
	wantStr(t, `function B1 { process { if ($_ -eq 2) { break }; "b:$_" } }; foreach ($i in 1..2) { 1,2,3 | B1; "inner" }; "after"`, "b:1", "after")
	// process 内 continue 上抛：终止当前语句（REPL 逐语句模式下后续语句照常）
	wantStr(t, `"start"; function C1 { process { if ($_ -eq 1) { continue }; "c:$_" } }; (1,2,3 | C1); "next"`, "start", "next")
}

func TestScriptParam(t *testing.T) {
	dir := t.TempDir()
	scriptPath := filepath.Join(dir, "s.ps1")
	content := "param($a, $b = 10)\n\"a=$a b=$b args=$args\""
	if err := os.WriteFile(scriptPath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := New(sess, strings.NewReader(""), io.Discard, io.Discard)

	// 显式实参按位置绑定
	out := ev.RunScriptFile(scriptPath, []*object.PSObject{object.Int(5), object.Int(6)})
	if got := strs(out); len(got) != 1 || got[0] != "a=5 b=6 args=" {
		t.Fatalf("显式实参 → %v", got)
	}
	// 缺少参数用默认值
	out = ev.RunScriptFile(scriptPath, nil)
	if got := strs(out); len(got) != 1 || got[0] != "a= b=10 args=" {
		t.Fatalf("默认值 → %v", got)
	}
	// 多余实参保留在 $args
	out = ev.RunScriptFile(scriptPath, []*object.PSObject{object.Int(1), object.Int(2), object.Int(3)})
	if got := strs(out); len(got) != 1 || got[0] != "a=1 b=2 args=3" {
		t.Fatalf("剩余实参 → %v", got)
	}
}

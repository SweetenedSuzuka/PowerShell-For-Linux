package eval

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"powershell/internal/parser"
	"powershell/internal/shell"
	"strings"
	"testing"
)

// TestErrorVariable 验证 $Error 自动变量：
// 非终止错误与 throw 都累积、最新在 [0]、被 catch 的错误也进、Clear/RemoveAt 作用到记录本体。
func TestErrorVariable(t *testing.T) {
	// 非终止错误累积，[0] 是最新
	wantStr(t, `[int]"abc"; $Error.Count`, "1")
	wantStr(t, `[int]"abc"; [int]"def"; $Error.Count; $Error[0].Message`,
		"2", `无法将值“def”转换为类型“int”。`)
	// throw 与被捕获的错误都进 $Error
	wantStr(t, `try { throw "boom" } catch { "in=$($Error.Count)" }; "after=$($Error.Count)"; $Error[0].Message`,
		"in=1", "after=1", "boom")
	// 无错误时为空数组
	wantStr(t, `$Error.Count`, "0")
	// Clear 清空本体
	wantStr(t, `[int]"abc"; $Error.Clear(); $Error.Count`, "0")
	// RemoveAt 删除指定下标
	wantStr(t, `[int]"abc"; [int]"def"; $Error.RemoveAt(0); $Error.Count; $Error[0].Message`,
		"1", `无法将值“abc”转换为类型“int”。`)
	// 参数绑定失败同样累积
	wantStr(t, `function F([int]$k) { "never" }; F 'bad'; $Error.Count; $Error[0].Message`,
		"1", `无法把实参“bad”转换成形参 k 声明的类型“int”。`)
}

// TestErrorAction 验证 -ErrorAction：默认继续、SilentlyContinue 与 Ignore 压住显示、Stop 转终止错误可被捕获、无效取值报绑定错误。
func TestErrorAction(t *testing.T) {
	// 默认 Continue：记录并把 $? 置为 false，后续语句继续（$? 用 if 读取，直接读取会被表达式语句先置位）
	wantStr(t, `Get-Item 不存在XYZ123; $Error.Count`, "1")
	wantStr(t, `Get-Item 不存在XYZ123; if ($?) { "ok" } else { "fail" }`, "fail")
	// SilentlyContinue：仍记录仍置为失败，只是不显示
	wantStr(t, `Get-Item 不存在XYZ123 -ErrorAction SilentlyContinue; $Error.Count`, "1")
	wantStr(t, `Get-Item 不存在XYZ123 -ErrorAction SilentlyContinue; if ($?) { "ok" } else { "fail" }`, "fail")
	// Ignore：不记录，只置为失败
	wantStr(t, `Get-Item 不存在XYZ123 -ErrorAction Ignore; $Error.Count`, "0")
	wantStr(t, `Get-Item 不存在XYZ123 -ErrorAction Ignore; if ($?) { "ok" } else { "fail" }`, "fail")
	// Inquire 在非交互场景按终止错误处理
	wantStr(t, `try { Get-Item 不存在XYZ123 -ErrorAction Inquire } catch { "caught" }; $Error.Count`, "caught", "1")
	// Stop：转为终止错误，可被 try/catch 捕获，只记一次
	wantStr(t, `try { Get-Item 不存在XYZ123 -ErrorAction Stop } catch { "caught" }; $Error.Count`, "caught", "1")
	// Stop 中断 try 体后续语句
	wantStr(t, `try { Get-Item 不存在XYZ123 -ErrorAction Stop; "after" } catch { "caught" }`, "caught")
	// 取值大小写不敏感
	wantStr(t, `try { Get-Item 不存在XYZ123 -ErrorAction stop } catch { "caught" }`, "caught")
	// 无效取值按绑定错误报告
	wantStr(t, `Get-Item foo -ErrorAction Bogus; $Error.Count`, "1")
	wantStr(t, `Get-Item foo -ErrorAction Bogus; if ($?) { "ok" } else { "fail" }`, "fail")
	// 不带值的 -ErrorAction 按缺少值报绑定错误
	wantStr(t, `Get-Item foo -ErrorAction; $Error.Count`, "1")
}

// TestErrorActionOutput 验证显示侧：默认 Continue 写 stderr，SilentlyContinue 与 Ignore 不写。
func TestErrorActionOutput(t *testing.T) {
	runWithStderr := func(src string) string {
		var errBuf bytes.Buffer
		sess := shell.New(shell.StyleCore, io.Discard, &errBuf, strings.NewReader(""))
		ev := New(sess, strings.NewReader(""), io.Discard, &errBuf)
		res := parser.Parse(src)
		if res.Error != nil {
			t.Fatalf("解析错误 %q: %v", src, res.Error)
		}
		for _, st := range res.List.Statements {
			ev.EvalStatement(st)
		}
		return errBuf.String()
	}
	if out := runWithStderr(`Get-Item 不存在XYZ123`); out == "" {
		t.Error("默认 Continue 应写 stderr")
	}
	if out := runWithStderr(`Get-Item 不存在XYZ123 -ErrorAction SilentlyContinue`); out != "" {
		t.Errorf("SilentlyContinue 不应写 stderr，得到 %q", out)
	}
	if out := runWithStderr(`Get-Item 不存在XYZ123 -ErrorAction Ignore`); out != "" {
		t.Errorf("Ignore 不应写 stderr，得到 %q", out)
	}
	if out := runWithStderr(`$ErrorActionPreference = 'SilentlyContinue'; Get-Item 不存在XYZ123`); out != "" {
		t.Errorf("首选项 SilentlyContinue 不应写 stderr，得到 %q", out)
	}
}

// TestPreferenceMigration 验证求值层错误同样受首选项分发：Stop 转终止可捕获，SilentlyContinue 只记不显，默认行为不变。
func TestPreferenceMigration(t *testing.T) {
	// 除零：Stop 可捕获且只记一次
	wantStr(t, `$ErrorActionPreference = 'Stop'; try { 5/0 } catch { "caught" }; $Error.Count`, "caught", "1")
	// 除零：SilentlyContinue 记录
	wantStr(t, `$ErrorActionPreference = 'SilentlyContinue'; 5/0; $Error.Count`, "1")
	// 类型转换失败：Stop 可捕获
	wantStr(t, `$ErrorActionPreference = 'Stop'; try { [int]"abc" } catch { "caught" }; $Error.Count`, "caught", "1")
	// 静态成员缺失：Stop 可捕获
	wantStr(t, `$ErrorActionPreference = 'Stop'; try { [math]::NoSuchMember() } catch { "caught" }; $Error.Count`, "caught", "1")
	// 函数形参转换失败：Stop 可捕获
	wantStr(t, `function FM([int]$k) { "never" }; $ErrorActionPreference = 'Stop'; try { FM 'bad' } catch { "caught" }; $Error.Count`, "caught", "1")
	// 读不到的脚本：Stop 可捕获
	wantStr(t, `$ErrorActionPreference = 'Stop'; try { ./不存在QS1.ps1 } catch { "caught" }; $Error.Count`, "caught", "1")
	// 默认行为不变：继续执行
	wantStr(t, `$x = 1/0; "after"`, "after")
	// 未捕获的终止错误只记一次
	wantStr(t, `throw "solo-boom"; $Error.Count`, "1")
	wantStr(t, `$ErrorActionPreference = 'Stop'; 5/0; $Error.Count`, "1")
}

// TestAssignSuccessFlag 验证赋值语句的 $?：右侧先求值（读取到旧状态），无新错误才置为 true。
func TestAssignSuccessFlag(t *testing.T) {
	// 右侧读取 $? 拿到上一条语句的状态
	wantStr(t, `Get-Item 不存在QW1; $x = $?; if ($x) { "ok" } else { "fail" }`, "fail")
	wantStr(t, `$y = 1; $x = $?; if ($x) { "ok" } else { "fail" }`, "ok")
	// 右侧无新错误时赋值后 $? 为真
	wantStr(t, `Get-Item 不存在QW1; $x = 5; if ($?) { "ok" } else { "fail" }`, "ok")
	// 右侧出错保持失败
	wantStr(t, `$x = 1/0; if ($?) { "ok" } else { "fail" }`, "fail")
	// 被捕获的错误不影响：赋值成功置为 true
	wantStr(t, `Get-Item 不存在QW1; $x = try { throw "a" } catch { "b" }; $x`, "b")
	wantStr(t, `Get-Item 不存在QW1; $x = try { throw "a" } catch { "b" }; if ($?) { "ok" } else { "fail" }`, "ok")
	// 裸 try/catch 捕获后同样置为 true
	wantStr(t, `Get-Item 不存在QW1; try { throw "a" } catch { "b" }; if ($?) { "ok" } else { "fail" }`, "b", "ok")
}

// TestErrorActionPreference 验证 $ErrorActionPreference：默认值、首选项分发、显式参数覆盖、函数内局部生效、无效赋值报错且不生效。
func TestErrorActionPreference(t *testing.T) {
	// 未赋值时读到默认值
	wantStr(t, `$ErrorActionPreference`, "Continue")
	// 首选项 SilentlyContinue：记录但不显示
	wantStr(t, `$ErrorActionPreference = 'SilentlyContinue'; Get-Item 不存在XYZ123; $Error.Count`, "1")
	// 首选项 Stop：转为终止错误，可捕获且只记一次
	wantStr(t, `$ErrorActionPreference = 'Stop'; try { Get-Item 不存在XYZ123 } catch { "caught" }; $Error.Count`, "caught", "1")
	// 显式 -ErrorAction 覆盖首选项
	wantStr(t, `$ErrorActionPreference = 'Stop'; Get-Item 不存在XYZ123 -ErrorAction Continue; $Error.Count`, "1")
	// 函数内赋值只在局部生效
	wantStr(t, `function Pref { $ErrorActionPreference = 'Stop'; try { Get-Item 不存在XYZ123 } catch { "caught" } }; Pref; $ErrorActionPreference`, "caught", "Continue")
	// 无效赋值报错且不生效
	wantStr(t, `$ErrorActionPreference = 'Bogus'; $Error.Count`, "1")
	wantStr(t, `$ErrorActionPreference = 'Bogus'; $ErrorActionPreference`, "Continue")
	// 大小写混写仍作用到同一首选项
	wantStr(t, `$erroractionpreference = 'Stop'; try { Get-Item 不存在XYZ123 } catch { "caught" }`, "caught")
	// 空值恢复默认
	wantStr(t, `$ErrorActionPreference = 'Stop'; $ErrorActionPreference = $null; $ErrorActionPreference`, "Continue")
	// Ignore 同样置为失败
	wantStr(t, `Get-Item 不存在XYZ123 -ErrorAction Ignore; if ($?) { "ok" } else { "fail" }`, "fail")
	// Set-Variable 与 Clear-Variable 使用同样的校验
	wantStr(t, `Set-Variable -Name ErrorActionPreference -Value 'Bogus'; $Error.Count`, "1")
	wantStr(t, `$ErrorActionPreference = 'Stop'; Clear-Variable ErrorActionPreference; $ErrorActionPreference`, "Continue")
}

// TestTryCatchFinally 验证 try/catch/finally + throw：
// 基本捕获、$_ 绑定、finally 恒执行、类型过滤、函数/循环传播、return 顺序、try 作为表达式。
func TestTryCatchFinally(t *testing.T) {
	// 基本捕获与 $_ 绑定（错误记录的 Message 属性）
	wantStr(t, `try { throw "boom" } catch { "已捕获" }`, "已捕获")
	wantStr(t, `try { throw "msg1" } catch { $_.Message }`, "msg1")
	// catch 块不新建独立作用域：普通变量赋值对外可见（与 PowerShell 一致）
	wantStr(t, `$tc = "未执行"; try { throw "boom" } catch { $tc = "已捕获" }; $tc`, "已捕获")
	// catch 的 $_ 是临时绑定：块结束后外层 $_ 不受影响
	wantStr(t, `$old = $_; try { throw "x" } catch { $tmp = $_ }; $_ -eq $old`, "True")
	// 无异常时 catch 不执行，finally 恒执行
	wantStr(t, `try { "正常体" } catch { "不会到" } finally { "finally" }`, "正常体", "finally")
	// 捕获后继续执行后续语句
	wantStr(t, `try { throw "x" } catch { "c" }; "继续"`, "c", "继续")
	// catch [System.Exception] 基类全捕
	wantStr(t, `try { throw "e1" } catch [System.Exception] { "基类捕获" }`, "基类捕获")
	// catch 精确类型不匹配 → 未捕获，顶层打印（无输出）
	wantStr(t, `try { throw "e2" } catch [System.ArgumentException] { "不会到" }`)
	// 多 catch 顺序：第一个匹配生效
	wantStr(t, `try { throw "x" } catch [System.ArgumentException] { "A" } catch { "全捕" }`, "全捕")
	// 函数内 throw 被调用方 try 捕获
	wantStr(t, `function f { throw "函数错" }; try { f } catch { "捕获" }`, "捕获")
	// 循环内 throw 传播到外层 try，且 throw 前循环已输出保留
	wantStr(t, `try { foreach ($i in 1..3) { if ($i -eq 2) { throw "循环错" }; "i=$i" } } catch { "循环捕获" }`, "i=1", "循环捕获")
	// return 时 finally 恒执行，且 finally 输出在返回值之前（与 PowerShell 一致）
	wantStr(t, `function g { try { return "r" } finally { "fin" } }; g`, "fin", "r")
	// 嵌套 try：内层类型不匹配 → 外层捕获
	wantStr(t, `try { try { throw "内层错" } catch [System.ArgumentException] { "不会到" } } catch { "外层捕获" }`, "外层捕获")
	// catch 内重抛，外层再捕获
	wantStr(t, `try { try { throw "a" } catch { "内捕"; throw } } catch { "外捕" }`, "内捕", "外捕")
	// finally 里 throw 覆盖原错误
	wantStr(t, `try { throw "原错" } catch { "会捕获" } finally { throw "finally错" }`, "会捕获")
	// catch 里的 break 传播到外层循环
	wantStr(t, `for ($i = 0; $i -lt 2; $i++) { try { throw "t" } catch { "i=$i"; break } }; "结束"`, "i=0", "结束")
	// try 作为表达式（赋值右侧）
	wantStr(t, `$r = try { "成功值" } catch { "失败值" }; "r=$r"`, "r=成功值")
	wantStr(t, `$r2 = try { throw "bad" } catch { "捕获值" }; "r2=$r2"`, "r2=捕获值")
	// 无 catch 只有 finally：finally 仍执行并输出，错误继续上抛（顶层打印，无输出）
	wantStr(t, `try { throw "只有finally" } finally { "清理" }`, "清理")
	// return 前与 finally 的输出在函数返回值之前（调用点统一语义）
	wantStr(t, `function f { "前"; return "r" }; f`, "前", "r")
	// 脚本块 return 沿途输出保留（与函数一致）
	wantStr(t, `"x" | ForEach-Object { "块前"; return "块后" }`, "块前", "块后")
	// 子表达式 return 沿途输出保留
	wantStr(t, `$("a"; return "b")`, "a", "b")
	// switch 不新建独立作用域：case 块内普通赋值对外可见（与 foreach 同机制）
	wantStr(t, `switch (1) { 1 { $sv = 5 } }; $sv`, "5")
}

// TestTryScriptPropagation 验证脚本内 throw 跨脚本传播与调用方 try 捕获。
func TestTryScriptPropagation(t *testing.T) {
	dir := t.TempDir()
	srcPath := filepath.Join(dir, "src.ps1")
	if err := os.WriteFile(srcPath, []byte("param($who)\n\"脚本运行中\"\nthrow \"脚本抛出错误 $who\"\n\"脚本尾部\"\n"), 0644); err != nil {
		t.Fatal(err)
	}
	callPath := filepath.Join(dir, "call.ps1")
	callSrc := "try { " + srcPath + " \"X\" } catch { \"调用方捕获: $($_.Message)\" }\n\"调用方继续\"\n"
	if err := os.WriteFile(callPath, []byte(callSrc), 0644); err != nil {
		t.Fatal(err)
	}
	// 跨脚本：被调脚本 throw 前输出保留，调用方捕获后继续
	wantStr(t, callSrc, "脚本运行中", "调用方捕获: 脚本抛出错误 X", "调用方继续")
	// 顶层逐语句（REPL 语义）：未捕获错误打印到 stderr（io.Discard），会话继续执行后续语句
	wantStr(t, "\"前\"; throw \"停\"; \"后\"", "前", "后")
}

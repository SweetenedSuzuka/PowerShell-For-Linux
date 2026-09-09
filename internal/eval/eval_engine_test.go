package eval

import (
	"io"
	"powershell/internal/parser"
	"powershell/internal/shell"
	"strings"
	"testing"
)

// TestReportPanic 验证顶层回收：普通 panic 转为报错，控制流信号继续传播。
func TestReportPanic(t *testing.T) {
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := New(sess, strings.NewReader(""), io.Discard, io.Discard)
	// 普通值处理：记录错误并置为失败
	if !ev.ReportPanic("boom") {
		t.Fatal("普通 panic 应处理")
	}
	if sess.LastSuccess {
		t.Fatal("处理后 $? 应为失败")
	}
	if len(sess.ErrorRecords) == 0 {
		t.Fatal("处理后 $Error 应有记录")
	}
	// 控制流信号继续传播
	rec := sess.RecordError("x")
	if ev.ReportPanic(&flowSignal{kind: flowError, value: rec}) {
		t.Fatal("控制流信号不应处理")
	}
	if ev.ReportPanic(nil) {
		t.Fatal("空值不应处理")
	}
}

// TestEvalStatementHalted 未捕获的终止错误返回 true（调用方中止后续语句），普通语句与非终止错误返回 false。
func TestEvalStatementHalted(t *testing.T) {
	runHalted := func(src string) []bool {
		t.Helper()
		sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
		ev := New(sess, strings.NewReader(""), io.Discard, io.Discard)
		res := parser.Parse(src)
		if res.Error != nil {
			t.Fatalf("解析错误 %q: %v", src, res.Error)
		}
		var halted []bool
		for _, st := range res.List.Statements {
			_, h := ev.EvalStatementHalted(st)
			halted = append(halted, h)
		}
		return halted
	}
	assertHalted := func(src string, expected ...bool) {
		t.Helper()
		actual := runHalted(src)
		if len(actual) != len(expected) {
			t.Fatalf("%q → %v，想要 %v", src, actual, expected)
		}
		for i := range expected {
			if actual[i] != expected[i] {
				t.Fatalf("%q → %v，想要 %v", src, actual, expected)
			}
		}
	}
	assertHalted(`"a"; "b"`, false, false)
	assertHalted(`throw "boom"`, true)
	assertHalted(`"a"; throw "boom"; "b"`, false, true, false)
	assertHalted(`5/0`, false)
	assertHalted(`try { throw "boom" } catch { "caught" }`, false)
}

// TestThrowKeepsPriorOutput 块内抛出错误前已产生的输出保留（与 PowerShell 一致）：脚本块、函数、子表达式。
func TestThrowKeepsPriorOutput(t *testing.T) {
	wantStr(t, `& { "a"; throw "x" }`, "a")
	wantStr(t, `function KP { "a"; throw "x" }; KP`, "a")
	wantStr(t, `$( "a"; throw "x" )`, "a")
	wantStr(t, `try { & { "a"; throw "x" } } catch { "caught" }`, "a", "caught")
}

// TestTryCatchQuestionMark 捕获错误后不修改问号变量：进入 catch 块时保持失败，空 catch 块之后仍为失败，无错误时照常成功。
func TestTryCatchQuestionMark(t *testing.T) {
	wantStr(t, `try { throw "x" } catch { if ($?) { "t" } else { "f" } }`, "f")
	wantStr(t, `try { throw "x" } catch { }; if ($?) { "t" } else { "f" }`, "f")
	wantStr(t, `try { try { throw "x" } finally { if ($?) { "t" } else { "f" } } } catch { "c" }`, "f", "c")
	wantStr(t, `try { "ok" } catch { "c" }; if ($?) { "t" } else { "f" }`, "ok", "t")
	wantStr(t, `try { throw "x" } catch { "ok" }; if ($?) { "t" } else { "f" }`, "ok", "t")
}

// TestStrictModeUndefinedVariable 严格模式未定义检查：打开后读取未定义变量并记录错误，跳过本句继续执行，已赋值的正常，函数内继承且向外不泄漏，非法版本记录错误，关闭后恢复。
func TestStrictModeUndefinedVariable(t *testing.T) {
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := New(sess, strings.NewReader(""), io.Discard, io.Discard)
	run := func(src string) []string {
		t.Helper()
		res := parser.Parse(src)
		if res.Error != nil {
			t.Fatalf("解析错误 %q: %v", src, res.Error)
		}
		var got []string
		for _, st := range res.List.Statements {
			got = append(got, strs(ev.EvalStatement(st))...)
		}
		return got
	}
	errs := func() int { return len(sess.ErrorRecords) }
	run(`$nosuch_sm_t1`)
	if errs() != 0 {
		t.Fatalf("默认关闭应无错，实际 %d 条", errs())
	}
	run(`Set-StrictMode -Version Latest`)
	run(`$nosuch_sm_t2`)
	if errs() != 1 {
		t.Fatalf("严格模式下读取未定义变量应记录 1 条，实际 %d 条", errs())
	}
	if sess.LastSuccess {
		t.Fatalf("读未定义后问号变量应为失败")
	}
	if got := run(`"after-strict"`); len(got) != 1 || got[0] != "after-strict" {
		t.Fatalf("出错后应继续执行，实际 %v", got)
	}
	run(`$smDef = 42`)
	if got := run(`$smDef`); len(got) != 1 || got[0] != "42" {
		t.Fatalf("已赋值变量应正常读出，实际 %v", got)
	}
	if errs() != 1 {
		t.Fatalf("正常读取不应记错，实际 %d 条", errs())
	}
	run(`function SMFunc { Set-StrictMode -Off; $nosuch_sm_t3 }`)
	run(`SMFunc`)
	if errs() != 1 {
		t.Fatalf("函数内关闭应生效，实际 %d 条", errs())
	}
	run(`$nosuch_sm_t4`)
	if errs() != 2 {
		t.Fatalf("函数内关闭不应泄漏到外层，实际 %d 条", errs())
	}
	run(`Set-StrictMode -Version 9.9`)
	if errs() != 3 {
		t.Fatalf("非法版本应记录错误，实际 %d 条", errs())
	}
	run(`Set-StrictMode -Off`)
	run(`$nosuch_sm_t5`)
	if errs() != 3 {
		t.Fatalf("关闭后应恢复，实际 %d 条", errs())
	}
}

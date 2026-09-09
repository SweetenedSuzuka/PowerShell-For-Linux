package eval

import (
	"testing"
)

func TestControlFlow(t *testing.T) {
	wantStr(t, "if (5 -gt 3) { 'big' } else { 'small' }", "big")
	wantStr(t, "foreach ($i in 1..3) { $i }", "1", "2", "3")
	wantStr(t, "$s = 0; foreach ($i in 1..4) { $s += $i }; $s", "10")
	wantStr(t, "$i = 0; while ($i -lt 3) { $i++; if ($i -eq 2) { continue }; $i }", "1", "3")
	wantStr(t, "switch (2) { 1 { '一' } 2 { '二' } default { '其他' } }", "二")
	wantStr(t, "for ($i=0; $i -lt 2; $i++) { $i }", "0", "1")
}

// TestSwitchArray 验证 switch 对数组逐元素匹配与 break/continue 语义。
func TestSwitchArray(t *testing.T) {
	// 数组逐元素：每个元素跑全部 case，default 按元素判断
	wantStr(t, "switch (1,2,3) { 2 { 'hit2' } default { '无匹配' } }", "无匹配", "hit2", "无匹配")
	// 逐元素 default 输出 $_（元素值）
	wantStr(t, "switch (1,2,3) { default { $_ } }", "1", "2", "3")
	// continue 进入下一元素（命中后剩余语句不执行）
	wantStr(t, "switch (1,2,3) { 2 { 'hit'; continue; 'x' } default { 'd' } }", "d", "hit", "d")
	// break 退出整个 switch
	wantStr(t, "switch (1,2,3) { 2 { 'hit'; break; 'x' } default { 'd' } }", "d", "hit")
	// 标量时多个 case 都检查（PowerShell 语义）
	wantStr(t, "switch (2) { 1 { '一' } 2 { '二' } 2 { '二2' } }", "二", "二2")
	// 标量 continue 退出 switch，不再检查后续 case
	wantStr(t, "switch (5) { 5 { 'a'; continue } 5 { 'b' } }", "a")
	// 字符串标量只按整体匹配一次
	wantStr(t, `switch ("abc") { "abc" { "yes" } default { "no" } }`, "yes")
	// regex 模式数组逐元素
	wantStr(t, `switch -regex ("a1","b2","c3") { "a\d" { "A" } "\d" { "N" } }`, "A", "N", "N", "N")
}

// TestStatementAsExpression 验证语句可作赋值右侧（$x = switch / if / foreach ...）。
func TestStatementAsExpression(t *testing.T) {
	// switch 作为表达式：数组逐元素输出
	wantStr(t, "$swr = switch (1,2,3) { default { $_ } }; $swr -join ','", "1,2,3")
	// switch 命中 case
	wantStr(t, "$v = switch (5) { 5 { 'five' } default { 'other' } }; $v", "five")
	// if 作为表达式
	wantStr(t, "$iv = if ($true) { 'yes' } else { 'no' }; $iv", "yes")
	// foreach 作为表达式
	wantStr(t, "$fv = foreach ($i in 1..3) { $i * 2 }; $fv -join ','", "2,4,6")
	// while 作为表达式
	wantStr(t, "$i = 0; $wv = while ($i -lt 3) { $i; $i++ }; $wv -join ','", "0,1,2")
}

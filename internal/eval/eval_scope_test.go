package eval

import (
	"testing"
)

func TestVariablesAndScope(t *testing.T) {
	wantStr(t, "$x = 10; $x * 2", "20")
	wantStr(t, "$x = 5; $x += 3; $x", "8")
	wantStr(t, `$s = "name"; "$s!"`, "name!")
	wantStr(t, "$a = 1,2,3; $a.Count", "3")
	wantStr(t, "$h = @{k=1}; $h.k", "1")
	wantStr(t, "$env:DUMMY_VAR_XYZ", "")
}

// TestScopeModifiers 验证 $script:/$global:/$local: 作用域修饰符（读写、复合、增量、插值）。
func TestScopeModifiers(t *testing.T) {
	// $script: 写回：函数内改脚本作用域变量
	wantStr(t, "$sf = 0; function Set { $script:sf = 5 }; Set; $sf", "5")
	// $script: 读取：函数内读取脚本作用域变量
	wantStr(t, "$sf = 7; function Get { $script:sf }; Get", "7")
	// $global: 写入：函数内写入全局变量
	wantStr(t, "$gg = 0; function SetG { $global:gg = 9 }; SetG; $gg", "9")
	// $local: 只读写当前作用域，不作用到外层
	wantStr(t, "$lv = 1; function L { $local:lv = 2; $lv }; L; $lv", "2", "1")
	// $script: 复合赋值
	wantStr(t, "$sf = 3; function A { $script:sf += 2 }; A; $sf", "5")
	// $script: 增量
	wantStr(t, "$sf = 1; function I { $script:sf++ }; I; $sf", "2")
	// 字符串插值里的 $script:
	wantStr(t, `$sf = 8; function S { "v=$script:sf" }; S`, "v=8")
	// ${} 括号形式
	wantStr(t, "$sf = 6; function B { ${script:sf} = 6 }; B; $sf", "6")
}

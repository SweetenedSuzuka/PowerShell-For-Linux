package builtin_test

import (
	"io"
	"strings"
	"testing"

	"powershell/internal/ast"
	"powershell/internal/builtin"
	"powershell/internal/eval"
	"powershell/internal/object"
	"powershell/internal/parser"
	"powershell/internal/shell"
)

// firstCommand 解析源码并取第一条命令（单命令管道或直接命令语句）。
func firstCommand(t *testing.T, src string) *ast.Command {
	t.Helper()
	res := parser.Parse(src)
	if res.Error != nil {
		t.Fatalf("解析错误 %q: %v", src, res.Error)
	}
	st := res.List.Statements[0]
	switch n := st.(type) {
	case *ast.Pipeline:
		if len(n.Commands) != 1 {
			t.Fatalf("%q 不是单命令管道", src)
		}
		return n.Commands[0]
	case *ast.Command:
		return n
	}
	t.Fatalf("%q 语句类型 %T，不是命令", src, st)
	return nil
}

// bind 用真实求值器做参数绑定（走 eval → builtin 的完整接口路径）。
func bind(t *testing.T, src string) (*builtin.BoundArgs, error) {
	t.Helper()
	cmd := firstCommand(t, src)
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := eval.New(sess, strings.NewReader(""), io.Discard, io.Discard)
	return builtin.Bind(ev, cmd, builtin.Spec(cmd.Name), nil)
}

// want 断言命名参数存在且字符串值相等。
func want(t *testing.T, args *builtin.BoundArgs, name, want string) {
	t.Helper()
	v := args.Get(name)
	if v == nil {
		t.Fatalf("缺少命名参数 %s（Named 里没有）", name)
	}
	if v.String() != want {
		t.Fatalf("%s = %q，想要 %q", name, v.String(), want)
	}
}

// TestBindPositionalMapping 位置实参按规格中心化映射到命名参数。
func TestBindPositionalMapping(t *testing.T) {
	args, err := bind(t, "Set-Content a.txt hello")
	if err != nil {
		t.Fatal(err)
	}
	want(t, args, "Path", "a.txt")
	want(t, args, "Value", "hello")
	if !args.PosMapped["Path"] || !args.PosMapped["Value"] {
		t.Errorf("位置映射应标记 PosMapped，实际 %v", args.PosMapped)
	}
	if len(args.Positional) != 0 {
		t.Errorf("映射后不应留位置实参，实际 %v", args.Positional)
	}
}

// TestBindNamedPrecedence 已命名参数不被位置实参覆盖，位置实参落到下一个空槽位。
func TestBindNamedPrecedence(t *testing.T) {
	args, err := bind(t, "Set-Content -Path a.txt hello")
	if err != nil {
		t.Fatal(err)
	}
	want(t, args, "Path", "a.txt")
	want(t, args, "Value", "hello")
	if args.PosMapped["Path"] {
		t.Error("显式命名 -Path 不应标记 PosMapped")
	}
	if !args.PosMapped["Value"] {
		t.Error("位置实参映射到 Value 应标记 PosMapped")
	}
}

// TestBindUnsetSlotNotPositional 位置绑定核心回归：未声明位置槽位的参数（如 -Encoding）
// 不参与位置绑定，位置实参只映射到声明过的槽位。
func TestBindUnsetSlotNotPositional(t *testing.T) {
	args, err := bind(t, "Set-Content -Path b.txt bval")
	if err != nil {
		t.Fatal(err)
	}
	want(t, args, "Path", "b.txt")
	want(t, args, "Value", "bval")
}

// TestBindSkipNamedSlot 已命名的槽位跳过，位置实参填入第一个空槽位。
func TestBindSkipNamedSlot(t *testing.T) {
	args, err := bind(t, "Join-Path -ChildPath child /tmp")
	if err != nil {
		t.Fatal(err)
	}
	want(t, args, "ChildPath", "child")
	want(t, args, "Path", "/tmp")
	if !args.PosMapped["Path"] {
		t.Error("位置实参映射到 Path 应标记 PosMapped")
	}
}

// TestBindScriptBlockLazy 脚本块只填 NamedNode（AST 节点），不预求值。
func TestBindScriptBlockLazy(t *testing.T) {
	args, err := bind(t, "Where-Object { $_ -gt 2 }")
	if err != nil {
		t.Fatal(err)
	}
	if args.GetNode("FilterScript") == nil {
		t.Error("脚本块应保留 AST 节点在 NamedNode")
	}
	if args.Get("FilterScript") != nil {
		t.Error("脚本块不应预求值进 Named")
	}
}

// TestBindExcessPositional 超出规格槽位数的实参留在 Positional 兜底。
func TestBindExcessPositional(t *testing.T) {
	args, err := bind(t, "Copy-Item a b c")
	if err != nil {
		t.Fatal(err)
	}
	want(t, args, "Path", "a")
	want(t, args, "Destination", "b")
	if len(args.Positional) != 1 || args.Positional[0].String() != "c" {
		t.Errorf("第 3 个实参应留 Positional，实际 %v", args.Positional)
	}
}

// TestBindSwitchNoSlot 开关参数不占位置槽位。
func TestBindSwitchNoSlot(t *testing.T) {
	args, err := bind(t, "Get-ChildItem -Recurse x")
	if err != nil {
		t.Fatal(err)
	}
	want(t, args, "Path", "x")
	if !args.Switches["Recurse"] {
		t.Error("开关 -Recurse 应置真")
	}
	if len(args.Positional) != 0 {
		t.Errorf("开关不占槽位，不应留位置实参，实际 %v", args.Positional)
	}
}

// TestBindPosMappedOnlyMapped 全命名调用不标记任何 PosMapped。
func TestBindPosMappedOnlyMapped(t *testing.T) {
	args, err := bind(t, "Set-Content -Path a -Value b")
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range args.PosMapped {
		if v {
			t.Errorf("全命名调用不应标记 PosMapped：%s", k)
		}
	}
}

// TestBindArrayPositional 数组位置实参按原样进 Named。
func TestBindArrayPositional(t *testing.T) {
	args, err := bind(t, "Copy-Item a1.txt,a2.txt d")
	if err != nil {
		t.Fatal(err)
	}
	v := args.Get("Path")
	if v == nil || !v.IsArray() {
		t.Fatalf("Path 应是数组，实际 %v", v)
	}
	if got := v.ArrayItems(); len(got) != 2 {
		t.Errorf("数组应有 2 项，实际 %d", len(got))
	}
	want(t, args, "Destination", "d")
}

// TestBindCommonParamConsumed 公共参数（-Encoding 等）被消费，不占槽位也不报错。
func TestBindCommonParamConsumed(t *testing.T) {
	args, err := bind(t, "Set-Content -Path a -Encoding utf8 v")
	if err != nil {
		t.Fatal(err)
	}
	want(t, args, "Path", "a")
	want(t, args, "Value", "v")
}

// TestBindUnknownNamed 未知命名参数报错。
func TestBindUnknownNamed(t *testing.T) {
	_, err := bind(t, "Set-Content -NoSuch x")
	if err == nil {
		t.Fatal("未知参数 -NoSuch 应报错")
	}
}

// TestBindScriptBlockNamedNode 命名脚本块与位置脚本块都进 NamedNode（GetNode 取到）。
func TestBindScriptBlockNamedNode(t *testing.T) {
	for _, src := range []string{
		"Where-Object { $_ -gt 2 }",
		"Where-Object -FilterScript { $_ -gt 2 }",
	} {
		args, err := bind(t, src)
		if err != nil {
			t.Fatalf("%q: %v", src, err)
		}
		if args.GetNode("FilterScript") == nil {
			t.Errorf("%q: FilterScript 应在 NamedNode", src)
		}
	}
}

// TestBindValueObjects 绑定值保持对象类型（字符串与数字不混淆）。
func TestBindValueObjects(t *testing.T) {
	args, err := bind(t, "Start-Sleep 2")
	if err != nil {
		t.Fatal(err)
	}
	v := args.Get("Seconds")
	if v == nil {
		t.Fatal("Seconds 应为 2")
	}
	if _, ok := v.Value.(int64); !ok {
		t.Errorf("Seconds 应为整数对象，实际 %T", v.Value)
	}
}

// TestBindSwitchValueFallback 开关被赋予值时开关置真、值退回位置参数
// （Bind 的 ArgNamed 分支：Get-ChildItem -Force foo 中 foo 是位置实参，再按槽位映射）。
func TestBindSwitchValueFallback(t *testing.T) {
	args, err := bind(t, "Get-ChildItem -Force foo")
	if err != nil {
		t.Fatal(err)
	}
	if !args.Switches["Force"] {
		t.Error("-Force 被赋予值时应置真")
	}
	want(t, args, "Path", "foo")
	if !args.PosMapped["Path"] {
		t.Error("退回的位置实参应参与槽位映射并标记 PosMapped")
	}
}

// TestBindInlineSwitch 内联开关 -Recurse:$false / -Recurse:true 按布尔求值赋给开关。
// $true/$false 与裸字 true/false 两种写法都覆盖：
// 裸字整段被词法器吃进 dash word（-Recurse:true）；
// $var 中 $ 不是 dash word 字符，词法器拆成独立 token，由解析器合并回内联值（-Recurse:$true）。
func TestBindInlineSwitch(t *testing.T) {
	cases := []struct {
		src  string
		want bool
	}{
		{"Get-ChildItem -Recurse:$false", false},
		{"Get-ChildItem -Recurse:$true", true},
		{"Get-ChildItem -Recurse:false", false},
		{"Get-ChildItem -Recurse:true", true},
		{"Get-ChildItem -Recurse:0", false},
		{"Get-ChildItem -Recurse:1", true},
	}
	for _, c := range cases {
		args, err := bind(t, c.src)
		if err != nil {
			t.Fatalf("%q 绑定出错: %v", c.src, err)
		}
		if args.Switches["Recurse"] != c.want {
			t.Errorf("%q 开关 = %v，想要 %v", c.src, args.Switches["Recurse"], c.want)
		}
	}
}

// TestBindInlineNamedValue 内联变量作命名参数值：-Path:$p 与 -Path:$env:HOME 形式。
func TestBindInlineNamedValue(t *testing.T) {
	cmd := firstCommand(t, "Set-Content -Path:$p -Value:$v")
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	if err := sess.SetVar("p", object.Str("/tmp/inline.txt")); err != nil {
		t.Fatal(err)
	}
	if err := sess.SetVar("v", object.Str("val")); err != nil {
		t.Fatal(err)
	}
	ev := eval.New(sess, strings.NewReader(""), io.Discard, io.Discard)
	args, err := builtin.Bind(ev, cmd, builtin.Spec(cmd.Name), nil)
	if err != nil {
		t.Fatal(err)
	}
	if args.Named["Path"].String() != "/tmp/inline.txt" {
		t.Errorf("Path = %q，想要 /tmp/inline.txt", args.Named["Path"].String())
	}
	if args.Named["Value"].String() != "val" {
		t.Errorf("Value = %q，想要 val", args.Named["Value"].String())
	}
}

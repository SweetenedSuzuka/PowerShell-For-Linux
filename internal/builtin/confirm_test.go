package builtin

import (
	"io"
	"strings"
	"testing"

	"powershell/internal/shell"
)

// confirmCtx 构造带 -Confirm 开关与指定应答流的调用上下文。
func confirmCtx(answer string, withConfirm bool) (*Context, *bool, *bool) {
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	switches := map[string]bool{}
	if withConfirm {
		switches["Confirm"] = true
	}
	args := &BoundArgs{Switches: switches}
	ctx := &Context{
		Shell:  sess,
		Stdin:  strings.NewReader(answer),
		Stdout: io.Discard,
		Args:   args,
	}
	return ctx, new(bool), new(bool)
}

// TestConfirmSkipWithoutSwitch 未指定 -Confirm 时直接执行，不读输入。
func TestConfirmSkipWithoutSwitch(t *testing.T) {
	ctx, yesAll, noAll := confirmCtx("", false)
	if confirmSkip(ctx, "Remove-Item", "a", yesAll, noAll) {
		t.Fatal("无 -Confirm 不应跳过")
	}
}

// TestConfirmSkipAnswers 各应答的跳过与执行判定。
func TestConfirmSkipAnswers(t *testing.T) {
	cases := []struct {
		answer string
		skip   bool
	}{
		{"y\n", false},
		{"yes\n", false},
		{"\n", false},
		{"n\n", true},
		{"no\n", true},
		{"l\n", true},
	}
	for _, tc := range cases {
		ctx, ya, na := confirmCtx(tc.answer, true)
		if got := confirmSkip(ctx, "Remove-Item", "a", ya, na); got != tc.skip {
			t.Errorf("应答 %q：skip=%v，想要 %v", tc.answer, got, tc.skip)
		}
	}
}

// TestConfirmSkipAllFlags A/L 的选择写入标记并作用于后续目标。
func TestConfirmSkipAllFlags(t *testing.T) {
	ctx, ya, na := confirmCtx("a\n", true)
	if confirmSkip(ctx, "Remove-Item", "a1", ya, na) {
		t.Fatal("A 应执行")
	}
	if !*ya || *na {
		t.Fatalf("A 应置 yesAll，实际 %v %v", *ya, *na)
	}
	ctx2, ya2, na2 := confirmCtx("l\n", true)
	if !confirmSkip(ctx2, "Remove-Item", "b1", ya2, na2) {
		t.Fatal("L 应跳过")
	}
	if !*na2 {
		t.Fatal("L 应置 noAll")
	}
}

// TestConfirmSkipEofRejects 输入结束没有回答时按拒绝处理。
func TestConfirmSkipEofRejects(t *testing.T) {
	ctx, ya, na := confirmCtx("", true)
	if !confirmSkip(ctx, "Remove-Item", "a", ya, na) {
		t.Fatal("EOF 应按拒绝处理")
	}
}

// TestConfirmSkipUnknownRetries 不认识的应答重新提示，直到出现可识别选项。
func TestConfirmSkipUnknownRetries(t *testing.T) {
	ctx, ya, na := confirmCtx("zz\nn\n", true)
	if !confirmSkip(ctx, "Remove-Item", "a", ya, na) {
		t.Fatal("未知应答重试后 n 应跳过")
	}
}

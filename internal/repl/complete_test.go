package repl

import (
	"io"
	"strings"
	"testing"

	"powershell/internal/shell"
)

// testREPL 构造仅带会话的补全测试体（补全不依赖终端）。
func testREPL() *REPL {
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	return &REPL{Session: sess}
}

// wantContain 断言补全结果包含预期项。
func wantContain(t *testing.T, buf string, want ...string) {
	t.Helper()
	got := testREPL().complete(buf)
	for _, w := range want {
		found := false
		for _, g := range got {
			if g == w {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("%q 补全缺 %q，实际 %v", buf, w, got)
		}
	}
}

// TestCompleteParam 验证 - 后补参数名：规格参数、通用参数、别名解析、未知命令无补全。
func TestCompleteParam(t *testing.T) {
	// 前缀匹配规格参数
	wantContain(t, "Get-ChildItem -F", "-Filter ", "-Force ", "-File ")
	// 通用参数
	wantContain(t, "Get-ChildItem -Error", "-ErrorAction ", "-ErrorVariable ")
	// 别名解析到原命令
	wantContain(t, "gci -R", "-Recurse ")
	// 未知命令无补全
	if got := testREPL().complete("NoSuchCmd -X"); len(got) != 0 {
		t.Fatalf("未知命令不应补全，实际 %v", got)
	}
	// 命令补全不受影响（既有返回小写名）
	wantContain(t, "Get-Ch", "get-childitem ")
}

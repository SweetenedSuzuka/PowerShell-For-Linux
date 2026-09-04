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

// assertCompleteContains 断言补全结果包含预期项。
func assertCompleteContains(t *testing.T, buf string, expected ...string) {
	t.Helper()
	actual := testREPL().complete(buf)
	for _, item := range expected {
		found := false
		for _, cand := range actual {
			if cand == item {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("%q 补全缺 %q，实际 %v", buf, item, actual)
		}
	}
}

// TestCompleteParam 验证 - 后补参数名：规格参数、通用参数、别名解析、未知命令无补全。
func TestCompleteParam(t *testing.T) {
	// 前缀匹配规格参数
	assertCompleteContains(t, "Get-ChildItem -F", "-Filter ", "-Force ", "-File ")
	// 通用参数
	assertCompleteContains(t, "Get-ChildItem -Error", "-ErrorAction ", "-ErrorVariable ")
	// 别名解析到原命令
	assertCompleteContains(t, "gci -R", "-Recurse ")
	// 未知命令无补全
	if actual := testREPL().complete("NoSuchCmd -X"); len(actual) != 0 {
		t.Fatalf("未知命令不应补全，实际 %v", actual)
	}
	// 命令补全不受影响（既有返回小写名）
	assertCompleteContains(t, "Get-Ch", "get-childitem ")
}

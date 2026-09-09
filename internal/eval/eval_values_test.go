package eval

import (
	"powershell/internal/shell"
	"testing"
)

func TestHashtableAndSubexpr(t *testing.T) {
	wantStr(t, "$h = @{ Name = 'x'; Count = 2 }; $h.Count", "2")
	wantStr(t, "$(1 + 2)", "3")
	wantStr(t, `"total: $(6 * 7)"`, "total: 42")
}

func TestHashtableProps(t *testing.T) {
	// 对照 PowerShell 语义：Count 返回条目数，键优先于属性
	wantStr(t, "@{a=1;b=2}.Count", "2")
	wantStr(t, "@{}.Count", "0")
	wantStr(t, "@{Count=5;x=1}.Count", "5")
	wantStr(t, "@{a=1;b=2}.Length", "1")
	wantStr(t, "@{a=1;b=2}.Keys", "a", "b")
	wantStr(t, "@{a=1;b=2}.Values", "1", "2")
	wantStr(t, "@{a=1;b=2}.Keys.Count", "2")
	wantStr(t, "@{a=1;b=2}.Keys[0]", "a")
}

func TestPSVersionTableCore(t *testing.T) {
	// 7.X 风格：PSVersion 只标 7（不精确到具体小版本），可读 .Major
	wantStr(t, "$PSVersionTable.PSVersion.Major", "7")
	wantStr(t, "$PSVersionTable.PSVersion.Minor", "0")
	wantStr(t, "$PSVersionTable.PSVersion", "7")
	wantStr(t, "$PSVersionTable.PSEdition", "Core")
	// OS 按 runtime.GOOS 报告，期望值跟随运行平台（Linux 上为 Linux、Windows 上为 Windows）
	wantStr(t, "$PSVersionTable.OS", osName())
	wantStr(t, "$PSVersionTable.GitCommitId", "0000000000000000000000000000000000000000")
}

func TestPSVersionTableDesktop(t *testing.T) {
	assertEvalOut := func(src string, expected ...string) {
		t.Helper()
		got := strs(runEvalWithStyle(t, shell.StyleDesktop, src))
		if len(got) != len(expected) {
			t.Fatalf("%q → %v，想要 %v", src, got, expected)
		}
		for i := range expected {
			if got[i] != expected[i] {
				t.Fatalf("%q → %v，想要 %v", src, got, expected)
			}
		}
	}
	assertEvalOut("$PSVersionTable.PSVersion.Major", "5")
	assertEvalOut("$PSVersionTable.PSVersion.Minor", "1")
	assertEvalOut("$PSVersionTable.PSVersion", "5.1")
	assertEvalOut("$PSVersionTable.PSEdition", "Desktop")
}

package eval

import (
	"io"
	"powershell/internal/object"
	"powershell/internal/parser"
	"powershell/internal/shell"
	"runtime"
	"strings"
	"testing"
)

// runEval 解析并执行一段源码，返回输出对象。

func runEval(t *testing.T, src string) []*object.PSObject {
	t.Helper()
	return runEvalWithStyle(t, shell.StyleCore, src)
}

// runEvalWithStyle 按指定风格（5.X/7.X）执行源码。
func runEvalWithStyle(t *testing.T, style shell.Style, src string) []*object.PSObject {
	t.Helper()
	sess := shell.New(style, io.Discard, io.Discard, strings.NewReader(""))
	ev := New(sess, strings.NewReader(""), io.Discard, io.Discard)
	res := parser.Parse(src)
	if res.Error != nil {
		t.Fatalf("解析错误 %q: %v", src, res.Error)
	}
	var out []*object.PSObject
	for _, st := range res.List.Statements {
		out = append(out, ev.EvalStatement(st)...)
	}
	return out
}

func strs(objs []*object.PSObject) []string {
	var out []string
	for _, o := range objs {
		out = append(out, o.String())
	}
	return out
}

func wantStr(t *testing.T, src string, expected ...string) {
	t.Helper()
	actual := strs(runEval(t, src))
	if len(actual) != len(expected) {
		t.Fatalf("%q → %v，想要 %v", src, actual, expected)
	}
	for i := range expected {
		if actual[i] != expected[i] {
			t.Fatalf("%q → %v，想要 %v", src, actual, expected)
		}
	}
}

// osName 复刻 shell.OSName 的平台映射，供测试构造跨平台的 OS 期望值。
func osName() string {
	switch runtime.GOOS {
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	case "darwin":
		return "Darwin"
	}
	return runtime.GOOS
}

package repl

import (
	"io"
	"strings"
	"testing"

	"powershell/internal/eval"
	"powershell/internal/parser"
	"powershell/internal/shell"
)

// evalPromptDef 定义函数并返回配好求值器的 REPL。
func evalPromptDef(t *testing.T, src string) *REPL {
	t.Helper()
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := eval.New(sess, strings.NewReader(""), io.Discard, io.Discard)
	res := parser.Parse(src)
	if res.Error != nil {
		t.Fatalf("解析错误：%v", res.Error)
	}
	for _, st := range res.List.Statements {
		ev.EvalStatement(st)
	}
	return &REPL{Session: sess, Eval: ev}
}

// TestPromptText 验证 prompt 函数定制提示符：未定义回默认，定义后取值，失败回默认。
func TestPromptText(t *testing.T) {
	// 未定义回默认
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	r := &REPL{Session: sess}
	if got := r.promptText(); got != sess.Prompt() {
		t.Fatalf("未定义应回默认，实际 %q", got)
	}
	// 定义后取值
	r = evalPromptDef(t, `function prompt { 'MYPROMPT> ' }`)
	if got := r.promptText(); got != "MYPROMPT> " {
		t.Fatalf("定义后应取值，实际 %q", got)
	}
	// 失败回默认
	r = evalPromptDef(t, `function prompt { throw 'x' }`)
	if got := r.promptText(); got != r.Session.Prompt() {
		t.Fatalf("失败应回默认，实际 %q", got)
	}
}

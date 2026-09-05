package repl

import (
	"io"
	"strings"
	"testing"

	"powershell/internal/eval"
	"powershell/internal/shell"
)

// seqReader 按步返回行或错误，步数耗尽返回 EOF。
type seqReader struct {
	steps []any // string 或 error
	i     int
}

func (s *seqReader) ReadLine(prompt string) (string, error) {
	if s.i >= len(s.steps) {
		return "", io.EOF
	}
	st := s.steps[s.i]
	s.i++
	if err, ok := st.(error); ok {
		return "", err
	}
	return st.(string), nil
}

// TestCancelClearsPending 验证续行中 Ctrl-C 丢弃累积：预置未闭合输入，一次中断后累积清空且不执行。
func TestCancelClearsPending(t *testing.T) {
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := eval.New(sess, strings.NewReader(""), io.Discard, io.Discard)
	var out strings.Builder
	r := &REPL{Session: sess, Eval: ev, out: &out}
	r.pending = "if ($true) {\n"
	r.reader = &seqReader{steps: []any{errLineCancelled}}
	r.loop()
	if r.pending != "" {
		t.Fatalf("中断后累积应清空，实际 %q", r.pending)
	}
	if strings.Contains(out.String(), "True") {
		t.Fatalf("中断后不应执行，输出 %q", out.String())
	}
}

package eval

import (
	"io"
	"strings"
	"testing"

	"powershell/internal/ast"
	"powershell/internal/parser"
	"powershell/internal/shell"
)

// externalCmd 解析源码取第一条命令，重建其外部 argv。
func externalCmd(t *testing.T, src string) (string, []string) {
	t.Helper()
	res := parser.Parse(src)
	if res.Error != nil {
		t.Fatalf("解析错误 %q: %v", src, res.Error)
	}
	st := res.List.Statements[0]
	var cmd *ast.Command
	switch n := st.(type) {
	case *ast.Pipeline:
		cmd = n.Commands[0]
	case *ast.Command:
		cmd = n
	default:
		t.Fatalf("%q 语句类型 %T", src, st)
	}
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := New(sess, strings.NewReader(""), io.Discard, io.Discard)
	return ev.externalArgv(cmd)
}

// TestExternalArgv 外部命令 argv 依源码顺序重建：位置实参原样、命名参数配对、开关独立。
func TestExternalArgv(t *testing.T) {
	cases := []struct {
		src     string
		wantProg string
		wantArgv []string
	}{
		{"echo hello world", "echo", []string{"hello", "world"}},
		// -n 后跟值：解析为命名参数，argv 里保持 "-n" + 值两元素
		{"grep -n foo f.txt", "grep", []string{"-n", "foo", "f.txt"}},
		// 开关无值：单独一个元素
		{"ls -la", "ls", []string{"-la"}},
		{"git commit -m msg", "git", []string{"commit", "-m", "msg"}},
		{"printf \"%s\" x", "printf", []string{"%s", "x"}},
		// 数字字面量保持文本
		{"sleep 2", "sleep", []string{"2"}},
	}
	for _, c := range cases {
		prog, argv := externalCmd(t, c.src)
		if prog != c.wantProg {
			t.Errorf("%q 程序名 = %q，想要 %q", c.src, prog, c.wantProg)
		}
		if len(argv) != len(c.wantArgv) {
			t.Errorf("%q argv = %v，想要 %v", c.src, argv, c.wantArgv)
			continue
		}
		for i := range c.wantArgv {
			if argv[i] != c.wantArgv[i] {
				t.Errorf("%q argv[%d] = %q，想要 %q（argv 整体 %v）", c.src, i, argv[i], c.wantArgv[i], argv)
			}
		}
	}
}

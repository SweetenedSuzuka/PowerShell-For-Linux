package eval

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"powershell/internal/parser"
	"powershell/internal/shell"
	"strings"
	"testing"
)

// TestStderrRedirectKeepsOutput 验证 2> 重定向不影响 stdout 输出。
func TestStderrRedirectKeepsOutput(t *testing.T) {
	// 2>$null 只丢弃错误流，输出照常（内置与函数各一例）
	wantStr(t, `Write-Output "w" 2>$null`, "w")
	wantStr(t, "function F { 'f-out' }; F 2>$null", "f-out")
	// > 文件：输出进文件，不进管道
	dir := t.TempDir()
	outFile := filepath.Join(dir, "o.txt")
	src := "Write-Output 'file' > " + outFile
	if got := strs(runEval(t, src)); len(got) != 0 {
		t.Fatalf("> 文件应无输出，实际 %v", got)
	}
	data, err := os.ReadFile(outFile)
	if err != nil || string(data) != "file\n" {
		t.Fatalf("> 文件应写入 file，err=%v data=%q", err, data)
	}
}

// TestStderrRedirectRestoredAfterExecutePanic 验证 2> 下命令执行期上抛后 stderr 归位。
func TestStderrRedirectRestoredAfterExecutePanic(t *testing.T) {
	t.Chdir(t.TempDir())
	var errBuf bytes.Buffer
	sess := shell.New(shell.StyleCore, io.Discard, &errBuf, strings.NewReader(""))
	ev := New(sess, strings.NewReader(""), io.Discard, &errBuf)
	// 脚本块内的越界调用在命令执行期上抛终止错误，穿过命令帧后由顶层语句打印
	res := parser.Parse(`"a" | ForEach-Object { $_.Substring(99, 5) } 2> e.txt`)
	if res.Error != nil {
		t.Fatalf("解析错误：%v", res.Error)
	}
	for _, st := range res.List.Statements {
		ev.EvalStatement(st)
	}
	// 恢复已执行，报错回到原缓冲，不滞留文件
	if !strings.Contains(errBuf.String(), "子字符串") {
		t.Fatalf("抛出错误后 stderr 应归位，实际 %q", errBuf.String())
	}
	// 后续命令的错误继续写原缓冲，不再进文件
	res = parser.Parse(`Get-Item 不存在XYZ123`)
	if res.Error != nil {
		t.Fatalf("解析错误：%v", res.Error)
	}
	for _, st := range res.List.Statements {
		ev.EvalStatement(st)
	}
	if !strings.Contains(errBuf.String(), "不存在XYZ123") {
		t.Fatalf("后续 stderr 应写原缓冲，实际 %q", errBuf.String())
	}
}

// TestReadStripsUTF8BOM 验证读文件与读脚本去掉 UTF-8 BOM。
func TestReadStripsUTF8BOM(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	// Get-Content：首行无 BOM
	txtPath := filepath.Join(dir, "b.txt")
	data := append([]byte{0xEF, 0xBB, 0xBF}, []byte("hi\nsecond\n")...)
	if err := os.WriteFile(txtPath, data, 0644); err != nil {
		t.Fatal(err)
	}
	wantStr(t, `Get-Content b.txt -TotalCount 1`, "hi")
	// 脚本：带 BOM 照常解析执行
	psPath := filepath.Join(dir, "b.ps1")
	psData := append([]byte{0xEF, 0xBB, 0xBF}, []byte("Write-Output 'ok'\n")...)
	if err := os.WriteFile(psPath, psData, 0644); err != nil {
		t.Fatal(err)
	}
	sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
	ev := New(sess, strings.NewReader(""), io.Discard, io.Discard)
	if got := strs(ev.RunScriptFile(psPath, nil)); len(got) != 1 || got[0] != "ok" {
		t.Fatalf("BOM 脚本 → %v，想要 [ok]", got)
	}
}

// TestAddContentEncoding 验证追加尊重 -Encoding，旧文件不重复写 BOM。
func TestAddContentEncoding(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)
	runEval(t, `Add-Content -Encoding utf8BOM e.txt 'a'`)
	runEval(t, `Add-Content -Encoding utf8BOM e.txt 'b'`)
	data, err := os.ReadFile(filepath.Join(dir, "e.txt"))
	if err != nil {
		t.Fatal(err)
	}
	want := append([]byte{0xEF, 0xBB, 0xBF}, []byte("a\nb\n")...)
	if string(data) != string(want) {
		t.Fatalf("追加两次应只有一个 BOM，实际 % x", data)
	}
}

// TestNonInteractive 验证 -NonInteractive 语义：读取输入报错，确认提示直接拒绝。
func TestNonInteractive(t *testing.T) {
	newEv := func() (*shell.Session, *Evaluator) {
		sess := shell.New(shell.StyleCore, io.Discard, io.Discard, strings.NewReader(""))
		return sess, New(sess, strings.NewReader(""), io.Discard, io.Discard)
	}
	run := func(ev *Evaluator, src string) {
		res := parser.Parse(src)
		if res.Error != nil {
			t.Fatalf("解析错误 %q：%v", src, res.Error)
		}
		for _, st := range res.List.Statements {
			ev.EvalStatement(st)
		}
	}
	// 读取输入报错
	sess, ev := newEv()
	sess.NonInteractive = true
	run(ev, `Read-Host`)
	if sess.LastSuccess {
		t.Fatal("非交互读取应失败")
	}
	// 确认提示直接拒绝：带 -Confirm 的删除不执行
	dir := t.TempDir()
	t.Chdir(dir)
	if err := os.WriteFile(filepath.Join(dir, "k.txt"), []byte("x"), 0644); err != nil {
		t.Fatal(err)
	}
	run(ev, `Remove-Item k.txt -Confirm`)
	if _, err := os.Stat(filepath.Join(dir, "k.txt")); err != nil {
		t.Fatal("非交互确认应拒绝删除")
	}
	// 默认会话不受影响：EOF 下读取无输出但不报错
	sess, ev = newEv()
	run(ev, `Read-Host`)
	if !sess.LastSuccess {
		t.Fatal("默认读取 EOF 不应失败")
	}
}

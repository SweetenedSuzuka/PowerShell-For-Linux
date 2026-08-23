// Package external 实现外部命令（原生程序）执行。
package external

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"powershell/internal/lang"
)

// LookPath 在 PATH 中查找可执行文件；带路径分隔符的名字直接判定。
func LookPath(name string) (string, bool) {
	if strings.ContainsAny(name, `/\`) {
		if _, err := os.Stat(name); err == nil {
			return name, true
		}
		return "", false
	}
	p, err := exec.LookPath(name)
	if err != nil {
		return "", false
	}
	return p, true
}

// Exec 执行外部命令，返回退出码。
// dir 是命令的工作目录（空则用进程当前目录）；stdin/stdout/stderr 决定 IO 接续。
func Exec(dir, program string, argv []string, stdin io.Reader, stdout, stderr io.Writer) int {
	cmd := exec.Command(program, argv...)
	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	err := cmd.Run()
	if err == nil {
		return 0
	}
	if ee, ok := err.(*exec.ExitError); ok {
		return ee.ExitCode()
	}
	fmt.Fprintf(stderr, "%s\n", lang.T(lang.MsgExternalExecFail, program, err))
	return 127
}

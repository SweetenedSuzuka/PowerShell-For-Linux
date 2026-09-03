//go:build !linux

package repl

import (
	"io"
	"os"
)

// newLineReader 在非 Linux 平台回退为简单行读取（无 raw 编辑）。
func newLineReader(in *os.File, out io.Writer, hist *[]string, complete func(string) []string) lineReader {
	return &simpleReader{in: in, out: out}
}

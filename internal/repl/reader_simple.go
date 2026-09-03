package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// simpleReader 是朴素行读取器（Windows/非终端回退）。
type simpleReader struct {
	in  io.Reader
	out io.Writer
	br  *bufio.Reader
}

// ReadLine 读取一行（含提示符输出）。
func (s *simpleReader) ReadLine(prompt string) (string, error) {
	fmt.Fprint(s.out, prompt)
	if s.br == nil {
		s.br = bufio.NewReader(s.in)
	}
	line, err := s.br.ReadString('\n')
	if err != nil {
		return strings.TrimRight(line, "\r\n"), err
	}
	return strings.TrimRight(line, "\r\n"), nil
}

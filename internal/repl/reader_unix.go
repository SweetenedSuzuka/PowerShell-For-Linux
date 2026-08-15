//go:build linux

package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// unixEditor 是基于 raw 终端模式的行编辑器（Linux）。
type unixEditor struct {
	in       *os.File
	out      io.Writer
	br       *bufio.Reader
	history  *[]string
	histIdx  int
	complete func(string) []string
}

// enterRaw 切换到 raw 终端模式，返回旧 termios。
func enterRaw(fd int) (syscall.Termios, bool) {
	var old syscall.Termios
	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TCGETS, uintptr(unsafe.Pointer(&old))); errno != 0 {
		return old, false
	}
	raw := old
	raw.Lflag &^= syscall.ECHO | syscall.ICANON
	raw.Iflag &^= syscall.IXON | syscall.ICRNL
	raw.Cc[syscall.VMIN] = 1
	raw.Cc[syscall.VTIME] = 0
	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TCSETS, uintptr(unsafe.Pointer(&raw))); errno != 0 {
		return old, false
	}
	return old, true
}

func restoreTermios(fd int, t syscall.Termios) {
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TCSETS, uintptr(unsafe.Pointer(&t)))
}

// newLineReader 在 Linux 上优先使用 raw 行编辑器；非终端回退简单读取。
func newLineReader(in *os.File, out io.Writer, hist *[]string, complete func(string) []string) lineReader {
	fd := int(in.Fd())
	var t syscall.Termios
	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TCGETS, uintptr(unsafe.Pointer(&t))); errno != 0 {
		return &simpleReader{in: in, out: out, history: hist}
	}
	return &unixEditor{
		in: in, out: out, br: bufio.NewReader(in),
		history: hist, complete: complete,
	}
}

// ReadLine 读取一行（raw 模式）。
func (e *unixEditor) ReadLine(prompt string) (string, error) {
	fd := int(e.in.Fd())
	old, ok := enterRaw(fd)
	if !ok {
		return e.readLineSimple(prompt)
	}
	defer restoreTermios(fd, old)

	buf := []rune{}
	pos := 0
	e.histIdx = len(*e.history)
	fmt.Fprint(e.out, prompt)
	for {
		r, _, err := e.br.ReadRune()
		if err != nil {
			if err == io.EOF && len(buf) > 0 {
				return string(buf), nil
			}
			return "", err
		}
		switch {
		case r == '\n' || r == '\r':
			fmt.Fprint(e.out, "\r\n")
			return string(buf), nil
		case r == 0x03: // Ctrl-C：清空当前行
			buf = buf[:0]
			pos = 0
			fmt.Fprint(e.out, "\r\n"+prompt)
		case r == 0x04: // Ctrl-D：空行退出
			if len(buf) == 0 {
				return "", io.EOF
			}
		case r == 0x0c: // Ctrl-L：清屏
			fmt.Fprint(e.out, "\x1b[2J\x1b[H"+prompt)
			e.redraw(prompt, buf, pos)
		case r == 0x7f || r == 0x08: // 退格
			if pos > 0 {
				buf = append(buf[:pos-1], buf[pos:]...)
				pos--
				e.redraw(prompt, buf, pos)
			}
		case r == '\t':
			e.doComplete(prompt, &buf, &pos)
		case r == 0x1b:
			e.handleEscape(prompt, &buf, &pos)
		default:
			if r >= 0x20 {
				buf = append(buf[:pos], append([]rune{r}, buf[pos:]...)...)
				pos++
				e.redraw(prompt, buf, pos)
			}
		}
	}
}

func (e *unixEditor) readLineSimple(prompt string) (string, error) {
	fmt.Fprint(e.out, prompt)
	line, err := e.br.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimRight(line, "\r\n"), nil
}

// redraw 重画当前行并把光标移回正确位置。
func (e *unixEditor) redraw(prompt string, buf []rune, pos int) {
	fmt.Fprint(e.out, "\r\x1b[K")
	text := prompt + string(buf)
	fmt.Fprint(e.out, text)
	back := runeWidth(text) - runeWidth(prompt+string(buf[:pos]))
	if back > 0 {
		fmt.Fprintf(e.out, "\x1b[%dD", back)
	}
}

func (e *unixEditor) handleEscape(prompt string, buf *[]rune, pos *int) {
	b, _ := e.br.ReadByte()
	if b != '[' {
		return
	}
	b2, _ := e.br.ReadByte()
	switch b2 {
	case 'A': // 上：历史前一条
		if e.histIdx > 0 {
			e.histIdx--
			*buf = []rune((*e.history)[e.histIdx])
			*pos = len(*buf)
			e.redraw(prompt, *buf, *pos)
		}
	case 'B': // 下：历史后一条
		if e.histIdx < len(*e.history) {
			e.histIdx++
			if e.histIdx == len(*e.history) {
				*buf = []rune{}
			} else {
				*buf = []rune((*e.history)[e.histIdx])
			}
			*pos = len(*buf)
			e.redraw(prompt, *buf, *pos)
		}
	case 'C': // 右
		if *pos < len(*buf) {
			*pos++
			e.redraw(prompt, *buf, *pos)
		}
	case 'D': // 左
		if *pos > 0 {
			*pos--
			e.redraw(prompt, *buf, *pos)
		}
	case 'H': // Home
		*pos = 0
		e.redraw(prompt, *buf, *pos)
	case 'F': // End
		*pos = len(*buf)
		e.redraw(prompt, *buf, *pos)
	case '3': // Delete
		_, _ = e.br.ReadByte() // 吃掉 '~'
		if *pos < len(*buf) {
			*buf = append((*buf)[:*pos], (*buf)[*pos+1:]...)
			e.redraw(prompt, *buf, *pos)
		}
	}
}

func (e *unixEditor) doComplete(prompt string, buf *[]rune, pos *int) {
	line := string(*buf)
	before := line[:*pos]
	ws := strings.LastIndexFunc(before, func(r rune) bool {
		return r == ' ' || r == '\t' || r == '|' || r == ';' || r == '('
	})
	ws++
	word := before[ws:]
	cands := e.complete(before)
	if len(cands) == 0 {
		return
	}
	if len(cands) == 1 {
		replacement := cands[0]
		newLine := line[:ws] + replacement + line[*pos:]
		*buf = []rune(newLine)
		*pos = runeLen(before[:ws] + replacement)
		e.redraw(prompt, *buf, *pos)
		return
	}
	common := commonPrefix(cands)
	if runeLen(common) > runeLen(word) {
		newLine := line[:ws] + common + line[*pos:]
		*buf = []rune(newLine)
		*pos = runeLen(before[:ws] + common)
		e.redraw(prompt, *buf, *pos)
		return
	}
	// 列出候选
	fmt.Fprint(e.out, "\r\n")
	fmt.Fprintln(e.out, strings.Join(cands, "  "))
	fmt.Fprint(e.out, prompt+string(*buf))
	back := runeWidth(prompt+string(*buf)) - runeWidth(prompt+string((*buf)[:*pos]))
	if back > 0 {
		fmt.Fprintf(e.out, "\x1b[%dD", back)
	}
}

func commonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs[1:] {
		for !strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix)) && len(prefix) > 0 {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func runeLen(s string) int { return len([]rune(s)) }

// runeWidth 计算显示宽度（CJK 等宽字符按 2）。
func runeWidth(s string) int {
	w := 0
	for _, r := range s {
		if r > 0x2E7F {
			w += 2
		} else {
			w++
		}
	}
	return w
}

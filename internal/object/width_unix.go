//go:build linux

package object

import (
	"syscall"
	"unsafe"
)

// winsize 是 TIOCGWINSZ 返回的终端尺寸。
type winsize struct {
	Row, Col, X, Y uint16
}

// terminalColumns 直读标准输出的终端宽度，失败返回 0。
func terminalColumns() int {
	var ws winsize
	if _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, 1, syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&ws))); errno != 0 {
		return 0
	}
	return int(ws.Col)
}

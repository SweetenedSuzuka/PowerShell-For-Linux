//go:build !linux

package object

// terminalColumns 在非 Linux 上恒返回 0，转而使用 COLUMNS 与默认分支。
func terminalColumns() int {
	return 0
}

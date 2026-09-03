package eval

import (
	"fmt"

	"powershell/internal/object"
)

// signals.go 定义控制流信号（break/continue/return/throw 沿调用栈上抛的载体）。
// flowKind 是控制流信号类型。
type flowKind int

const (
	flowBreak flowKind = iota
	flowContinue
	flowReturn
	flowExit
	flowError // 终止错误（throw / 未捕获错误）
)

// flowSignal 用 panic/recover 传递 break/continue/return/exit/error。
type flowSignal struct {
	kind  flowKind
	value *object.PSObject // return 值 / 错误记录
	code  int              // exit 码
	out   []*object.PSObject // panic 前已产生的输出（传播时保留，如 throw 前循环的输出）
}

// RecoverError 提取 panic 值里的终止错误，供 main.go 在 -File 顶层打印。
// 非错误 panic 返回 nil，交由调用方继续传播。
func RecoverError(r any) error {
	if fs, ok := r.(*flowSignal); ok && fs.kind == flowError {
		return fmt.Errorf("%s", fs.value.String())
	}
	return nil
}


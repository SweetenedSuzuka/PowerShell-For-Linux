// Package builtin 实现内置 cmdlet（命令）。
//
// 每个 cmdlet 声明参数规格（ParamSpec），求值器（eval）据此做参数绑定，生成 BoundArgs。
// cmdlet 实现从 BoundArgs 读取参数，消费 Input 对象并产出新对象。
package builtin

import (
	"fmt"
	"io"
	"strings"

	"powershell/internal/ast"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// Engine 由 eval 包实现，供内置 cmdlet 调用（求值表达式、执行脚本块）。
type Engine interface {
	// EvalExpr 在 extra 变量作用域下求值表达式。
	EvalExpr(node ast.Node, extra map[string]*object.PSObject) (*object.PSObject, error)
	// InvokeBlock 执行脚本块并返回其输出的对象（不打印）。
	InvokeBlock(block *ast.Block, extra map[string]*object.PSObject) ([]*object.PSObject, error)
	// EvalFilterExpr 在 $_ = obj 上下文求值过滤表达式（裸字视为对象属性）。
	EvalFilterExpr(node ast.Node, obj *object.PSObject) (bool, error)
	// RunSource 解析并执行一段源码，返回输出对象（Invoke-Expression / Invoke-History 用）。
	RunSource(src string) ([]*object.PSObject, error)
	// LookupVar 按 PowerShell 默认读语义查变量（自顶向下查作用域，再查自动变量）。
	LookupVar(name string) *object.PSObject
}

// Context 是内置 cmdlet 的调用上下文。
type Context struct {
	Shell  *shell.Session
	Engine Engine
	Stdout io.Writer
	Stderr io.Writer
	Stdin  io.Reader
	// Console 是不受重定向影响的主机输出（提示、Write-Host、Out-Host、清屏用）；为空时回退 Stdout。
	Console io.Writer
	Args    *BoundArgs
	Input   []*object.PSObject // 管道输入对象
}

// console 取主机输出；未设置时回退 Stdout。
func (c *Context) console() io.Writer {
	if c.Console != nil {
		return c.Console
	}
	return c.Stdout
}

// BoundArgs 是绑定后的命令参数。
type BoundArgs struct {
	Positional     []*object.PSObject
	PositionalNode []ast.Node
	Named          map[string]*object.PSObject
	NamedNode      map[string]ast.Node
	Switches       map[string]bool
	PosMapped      map[string]bool // 参数名 → 是否由位置实参中心化映射而来（cmdlet 区分"显式命名"与"位置"用）
	ErrorAction    string          // -ErrorAction 的归一化取值（小写），空为未指定，按 Continue 处理
}

// Pos 取第 i 个位置参数（越界返回 nil）。
func (a *BoundArgs) Pos(i int) *object.PSObject {
	if i >= 0 && i < len(a.Positional) {
		return a.Positional[i]
	}
	return nil
}

// PosNode 取第 i 个位置参数的原始 AST 节点。
func (a *BoundArgs) PosNode(i int) ast.Node {
	if i >= 0 && i < len(a.PositionalNode) {
		return a.PositionalNode[i]
	}
	return nil
}

// Get 取命名参数（不存在返回 nil）。
func (a *BoundArgs) Get(name string) *object.PSObject {
	return a.Named[name]
}

// GetNode 取命名参数的原始 AST 节点。
func (a *BoundArgs) GetNode(name string) ast.Node {
	return a.NamedNode[name]
}

// Str 取命名参数并转字符串。
func (a *BoundArgs) Str(name string) (string, bool) {
	if v := a.Named[name]; v != nil && !v.IsNull() {
		return v.String(), true
	}
	return "", false
}

// Int 取命名参数并转整数。
func (a *BoundArgs) Int(name string) (int64, bool) {
	if v := a.Named[name]; v != nil {
		return v.AsInt()
	}
	return 0, false
}

// Switch 判断开关是否打开。
func (a *BoundArgs) Switch(name string) bool { return a.Switches[name] }

// StringSlice 把参数解释为字符串数组（-Property Name,Length 等）。
func (a *BoundArgs) StringSlice(name string) []string {
	v := a.Named[name]
	if v == nil {
		return nil
	}
	items := v.ArrayItems()
	out := make([]string, 0, len(items))
	for _, it := range items {
		out = append(out, it.String())
	}
	return out
}

// Error 生成一条错误对象输出（写入 stderr）。
func (c *Context) Error(msg string) []*object.PSObject {
	fmt.Fprintf(c.Stderr, "%s: %s\n", c.Shell.StyleName(), msg)
	return nil
}

// inlineSwitchBool 把内联开关值解释为布尔：$true/$false 原样。
// 裸字 true/false/1/0（PowerShell -Switch:value 写法）按布尔解析。
func inlineSwitchBool(v *object.PSObject) bool {
	if b, ok := v.Value.(bool); ok {
		return b
	}
	switch strings.ToLower(v.String()) {
	case "true", "1":
		return true
	case "false", "0", "":
		return false
	}
	return v.Truthy()
}

// ParamSpec 描述一个命令参数。
type ParamSpec struct {
	Name        string
	Switch      bool   // 开关参数（不取值）
	Position    int    // 位置参数序号；仅 PositionSet 为 true 时参与位置绑定
	PositionSet bool   // 是否显式声明了位置槽位（未声明的参数只接受命名，默认不参与位置绑定）
	Type        string // 提示：string/int/path/bool/object[]/scriptblock
}

// CmdFunc 是内置 cmdlet 实现签名。
type CmdFunc func(c *Context) ([]*object.PSObject, error)

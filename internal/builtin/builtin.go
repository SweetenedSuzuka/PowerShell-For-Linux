// Package builtin 实现内置 cmdlet（命令）。
//
// 每个 cmdlet 声明参数规格（ParamSpec），求值器（eval）据此做参数绑定，生成 BoundArgs。
// cmdlet 实现从 BoundArgs 读取参数，消费 Input 对象并产出新对象。
package builtin

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"powershell/internal/ast"
	"powershell/internal/lang"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// Engine 由 eval 包实现，供内置 cmdlet 调用（求值表达式、执行脚本块）。
type Engine interface {
	// EvalExpr 在 extra 变量作用域下求值表达式。
	EvalExpr(node ast.Node, extra map[string]*object.PSObject) (*object.PSObject, error)
	// InvokeBlock 执行脚本块并返回其输出的对象（不打印）。
	InvokeBlock(block *ast.Block, extra map[string]*object.PSObject, stdout io.Writer) ([]*object.PSObject, error)
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
	Args   *BoundArgs
	Input  []*object.PSObject // 管道输入对象
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

var registry = map[string]CmdFunc{}
var specMap = map[string][]ParamSpec{}
var displayMap = map[string]string{} // 小写名 → 原始大小写名（Get-Help/Get-Command 显示用）

// Register 注册 cmdlet（名字转小写存储）。
func Register(name string, spec []ParamSpec, fn CmdFunc) {
	lower := strings.ToLower(name)
	registry[lower] = fn
	specMap[lower] = spec
	displayMap[lower] = name
}

// Lookup 查找 cmdlet。
func Lookup(name string) (CmdFunc, bool) {
	fn, ok := registry[strings.ToLower(name)]
	return fn, ok
}

// Spec 返回 cmdlet 的参数规格。
func Spec(name string) []ParamSpec {
	return specMap[strings.ToLower(name)]
}

// commonParams 是所有 cmdlet 通用的参数：-ErrorAction、-WhatIf、-Confirm 按各自语义生效，其余接受但忽略。
var commonParams = map[string]bool{
	"erroraction": true, "errorvariable": true, "verbose": true, "debug": true,
	"whatif": true, "confirm": true, "outvariable": true, "outbuffer": true,
	"warningaction": true, "warningvariable": true, "informationaction": true,
	"informationvariable": true, "pipelinevariable": true, "progressaction": true,
	"inputobject": true, "encoding": true,
}

// CommonParamNames 返回通用参数的规范名（与 commonParams 一一对应，供补全用）。
func CommonParamNames() []string {
	return []string{
		"ErrorAction", "ErrorVariable", "Verbose", "Debug",
		"WhatIf", "Confirm", "OutVariable", "OutBuffer",
		"WarningAction", "WarningVariable", "InformationAction",
		"InformationVariable", "PipelineVariable", "ProgressAction",
		"InputObject", "Encoding",
	}
}

// commonSwitchParams 是开关型常见参数：不取值，后跟的值仍按位置参数处理。
var commonSwitchParams = map[string]bool{
	"verbose": true, "debug": true, "whatif": true, "confirm": true,
}

// Bind 依据参数规格把命令实参绑定为 BoundArgs。
// 命名参数进入 Named/Switches；位置实参先按源码顺序收进 Positional。
// bindPositional 依据规格的 Position 序号把位置实参中心化映射到命名参数；未声明位置槽位的实参（如数组多源、超量实参）保留在 Positional 里。
// engine 用于求值参数表达式；extra 是求值时额外的变量作用域。
func Bind(engine Engine, cmd *ast.Command, spec []ParamSpec, extra map[string]*object.PSObject) (*BoundArgs, error) {
	ba := &BoundArgs{
		Named:     map[string]*object.PSObject{},
		NamedNode: map[string]ast.Node{},
		Switches:  map[string]bool{},
		PosMapped: map[string]bool{},
	}
	findSpec := func(name string) *ParamSpec {
		for i := range spec {
			if strings.EqualFold(spec[i].Name, name) {
				return &spec[i]
			}
		}
		return nil
	}
	// 按源码顺序绑定（保持 -Force foo 中 foo 回到位置参数的语义）
	for _, slot := range cmd.ArgOrder {
		switch slot.Kind {
		case ast.ArgPositional:
			node := cmd.Positional[slot.Index]
			val, err := engine.EvalExpr(node, extra)
			if err != nil {
				return nil, err
			}
			ba.Positional = append(ba.Positional, val)
			ba.PositionalNode = append(ba.PositionalNode, node)
		case ast.ArgNamed:
			sp := findSpec(slot.Name)
			node := cmd.Named[slot.Index].Value
			val, err := engine.EvalExpr(node, extra)
			if err != nil {
				return nil, err
			}
			if sp == nil && commonParams[strings.ToLower(slot.Name)] {
				// 开关型常见参数（-Verbose 等）的值退回位置参数（-Verbose foo 中 foo 是位置实参）。
				// 取值型常见参数的值被参数消费：-ErrorAction 记入绑定供错误分发，其余直接忽略。
				// WhatIf/Confirm 记录为开关供 cmdlet 读取；-WhatIf:$false 的内联布尔照常生效。
				if strings.EqualFold(slot.Name, "erroraction") {
					action, ok := shell.ParseErrorAction(val.String())
					if !ok {
						return nil, fmt.Errorf("%s", lang.T(lang.MsgBindErrorActionInvalid, val.String()))
					}
					ba.ErrorAction = action
					continue
				}
				if strings.EqualFold(slot.Name, "whatif") || strings.EqualFold(slot.Name, "confirm") {
					ba.Switches[slot.Name] = inlineSwitchBool(val)
				} else if commonSwitchParams[strings.ToLower(slot.Name)] {
					ba.Positional = append(ba.Positional, val)
					ba.PositionalNode = append(ba.PositionalNode, node)
				}
				continue
			}
			if sp == nil {
				return nil, fmt.Errorf("%s", lang.T(lang.MsgBindNoParam, slot.Name))
			}
			if sp.Switch {
				if slot.Inline {
					// -Force:value 内联形式：按布尔求值赋给开关（-Recurse:$false / -Recurse:false 关闭递归）
					ba.Switches[sp.Name] = inlineSwitchBool(val)
				} else {
					// 开关被赋予了值：开关置真，值退回位置参数（如 Get-ChildItem -Force foo）
					ba.Switches[sp.Name] = true
					ba.Positional = append(ba.Positional, val)
					ba.PositionalNode = append(ba.PositionalNode, node)
				}
			} else {
				ba.Named[sp.Name] = val
				ba.NamedNode[sp.Name] = node
			}
		case ast.ArgSwitch:
			sp := findSpec(slot.Name)
			if sp == nil && commonParams[strings.ToLower(slot.Name)] {
				// 不带值的 -ErrorAction 按缺值报绑定错误。
				if strings.EqualFold(slot.Name, "erroraction") {
					return nil, fmt.Errorf("%s", lang.T(lang.MsgBindSwitchNoValue, slot.Name))
				}
				// WhatIf/Confirm 记录进开关表，cmdlet 据此跳过实际变更
				ba.Switches[slot.Name] = true
				continue
			}
			if sp == nil {
				return nil, fmt.Errorf("%s", lang.T(lang.MsgBindNoParam, slot.Name))
			}
			if sp.Switch {
				ba.Switches[sp.Name] = true
			} else {
				return nil, fmt.Errorf("%s", lang.T(lang.MsgBindSwitchNoValue, slot.Name))
			}
		}
	}
	bindPositional(ba, spec)
	return ba, nil
}

// bindPositional 把位置实参按规格的 Position 序号中心化映射到命名参数。
// 规则：
//   - 只有显式声明了位置槽位（PositionSet）的参数参与，未声明的参数（如
//     -Encoding、-Filter 这类仅命名参数）不占槽位；
//   - 第 k 个位置实参（0 起）映射到 Position 序号第 k 大的参数；
//   - 已被显式命名赋值的槽位跳过（如 Set-Content -Path foo bar 中 bar 落到 Value）；
//   - 脚本块参数只映射 AST 节点（NamedNode），保持惰性求值；
//   - 超出规格声明范围的实参保留在 Positional，由 cmdlet 自行读取处理。
func bindPositional(ba *BoundArgs, spec []ParamSpec) {
	var slots []*ParamSpec
	for i := range spec {
		if spec[i].PositionSet && !spec[i].Switch {
			slots = append(slots, &spec[i])
		}
	}
	if len(slots) == 0 {
		return
	}
	sort.Slice(slots, func(a, b int) bool { return slots[a].Position < slots[b].Position })
	var rest []*object.PSObject
	var restNode []ast.Node
	next := 0
	for i := 0; i < len(ba.Positional); i++ {
		// 跳过已被显式命名赋值的槽位（脚本块只填 NamedNode，两者都要查）
		for next < len(slots) && (ba.Named[slots[next].Name] != nil || ba.NamedNode[slots[next].Name] != nil) {
			next++
		}
		if next < len(slots) {
			sp := slots[next]
			if sp.Type == "scriptblock" {
				// 脚本块保留原始 AST 节点（惰性求值），值不预求值
				ba.NamedNode[sp.Name] = ba.PositionalNode[i]
			} else {
				ba.Named[sp.Name] = ba.Positional[i]
				ba.NamedNode[sp.Name] = ba.PositionalNode[i]
			}
			ba.PosMapped[sp.Name] = true
			next++
		} else {
			rest = append(rest, ba.Positional[i])
			restNode = append(restNode, ba.PositionalNode[i])
		}
	}
	ba.Positional = rest
	ba.PositionalNode = restNode
}

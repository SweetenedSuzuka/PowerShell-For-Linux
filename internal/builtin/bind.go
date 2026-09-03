package builtin

import (
	"fmt"
	"sort"
	"strings"

	"powershell/internal/ast"
	"powershell/internal/lang"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// bind.go 依据参数规格把命令实参绑定为 BoundArgs。

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

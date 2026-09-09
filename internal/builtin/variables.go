package builtin

import (
	"strings"

	"powershell/internal/lang"
	"powershell/internal/object"
)

// variables.go 实现变量类 cmdlet。

func cmdGetVariable(c *Context) ([]*object.PSObject, error) {
	names := c.Args.StringSlice("Name")
	match := func(n string) bool {
		if len(names) == 0 {
			return true
		}
		for _, pat := range names {
			if object.WildcardMatchFold(pat, n) {
				return true
			}
		}
		return false
	}
	var out []*object.PSObject
	for _, n := range c.Shell.AllVarNames() {
		if !match(n) {
			continue
		}
		if v, ok := c.Shell.GetVar(n); ok {
			o := object.Object("System.Management.Automation.PSVariable", nil)
			o.AddProp("Name", n)
			o.AddProp("Value", v)
			out = append(out, o)
		}
	}
	return out, nil
}

func cmdSetVariable(c *Context) ([]*object.PSObject, error) {
	name, val := pairArgs(c, "Name")
	if name == "" {
		return nil, nil
	}
	if val == nil && len(c.Input) > 0 {
		val = object.Array(c.Input)
	}
	if val != nil {
		if err := c.Shell.SetVar(name, val); err != nil {
			return errf(c, "Set-Variable : %v", err)
		}
	}
	return nil, nil
}

func cmdNewVariable(c *Context) ([]*object.PSObject, error) {
	name, val := pairArgs(c, "Name")
	if name == "" {
		return nil, nil
	}
	if c.Shell.HasVar(name) && !c.Args.Switch("Force") {
		return errf(c, "%s", lang.T(lang.MsgVarExists, name))
	}
	if val == nil {
		val = object.Null()
	}
	if err := c.Shell.SetVar(name, val); err != nil {
		return errf(c, "New-Variable : %v", err)
	}
	return nil, nil
}

func cmdRemoveVariable(c *Context) ([]*object.PSObject, error) {
	for _, n := range c.Args.StringSlice("Name") {
		c.Shell.DeleteVar(n)
	}
	return nil, nil
}

func cmdClearVariable(c *Context) ([]*object.PSObject, error) {
	if name, ok := c.Args.Str("Name"); ok && name != "" {
		_ = c.Shell.SetVar(name, object.Null())
	}
	return nil, nil
}

// cmdSetStrictMode 设置当前作用域的严格模式（只做未定义变量检查，各版本行为一致）。
func cmdSetStrictMode(c *Context) ([]*object.PSObject, error) {
	// 超量位置实参无槽位可接（Version 只占位置 0），报错而非静默忽略。
	if len(c.Args.Positional) > 0 {
		return errf(c, "%s", lang.T(lang.MsgPositionalParamNotFound, c.Args.Positional[0].String()))
	}
	ver, _ := c.Args.Str("Version")
	off := c.Args.Switch("Off")
	// -Version 与 -Off 分属不同参数集，不可同用；都不给也无法解析。
	if (ver == "" && !off) || (ver != "" && off) {
		return errf(c, "%s", lang.T(lang.MsgParamSetUnresolvable))
	}
	if off {
		c.Engine.SetStrictMode(false)
		return nil, nil
	}
	switch strings.ToLower(strings.TrimSpace(ver)) {
	case "latest", "1", "2", "3", "1.0", "2.0", "3.0":
		c.Engine.SetStrictMode(true)
		return nil, nil
	default:
		return errf(c, "%s", lang.T(lang.MsgStrictVersionBad, ver))
	}
}

// ---- 注册 ----

func init() {
	Register("Get-Variable", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdGetVariable)
	Register("Set-Variable", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Value", Position: 1, PositionSet: true, Type: "object"},
	}, cmdSetVariable)
	Register("New-Variable", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Value", Position: 1, PositionSet: true, Type: "object"},
		{Name: "Force", Switch: true},
	}, cmdNewVariable)
	Register("Remove-Variable", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdRemoveVariable)
	Register("Clear-Variable", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdClearVariable)
	Register("Set-StrictMode", []ParamSpec{
		{Name: "Version", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Off", Switch: true},
	}, cmdSetStrictMode)
}

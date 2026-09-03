package builtin

import (
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
			if object.WildcardMatch(pat, n) {
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
	if _, exists := c.Shell.Vars[name]; exists && !c.Args.Switch("Force") {
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
		delete(c.Shell.Vars, n)
	}
	return nil, nil
}

func cmdClearVariable(c *Context) ([]*object.PSObject, error) {
	if name, ok := c.Args.Str("Name"); ok && name != "" {
		_ = c.Shell.SetVar(name, object.Null())
	}
	return nil, nil
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
}

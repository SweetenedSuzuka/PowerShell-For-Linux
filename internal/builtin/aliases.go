package builtin

import (
	"fmt"
	"os"
	"strings"

	"powershell/internal/lang"
	"powershell/internal/object"
)

// aliases.go 实现别名类 cmdlet。

func cmdGetAlias(c *Context) ([]*object.PSObject, error) {
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
	for _, n := range c.Shell.AllAliasNames() {
		if !match(n) {
			continue
		}
		o := object.Object("AliasInfo", nil)
		o.AddProp("Name", n)
		o.AddProp("Definition", c.Shell.Aliases[n])
		o.Table = []object.Column{
			{Label: "Name", Align: "left"},
			{Label: "Definition", Align: "left"},
		}
		out = append(out, o)
	}
	return out, nil
}

func cmdSetAlias(c *Context) ([]*object.PSObject, error) {
	name, _ := c.Args.Str("Name")
	value, _ := c.Args.Str("Value")
	if name != "" && value != "" {
		c.Shell.SetAlias(name, value)
	}
	return nil, nil
}

func cmdNewAlias(c *Context) ([]*object.PSObject, error) {
	name, _ := c.Args.Str("Name")
	value, _ := c.Args.Str("Value")
	if name == "" || value == "" {
		return nil, nil
	}
	if _, exists := c.Shell.Aliases[name]; exists && !c.Args.Switch("Force") {
		return errf(c, "%s", lang.T(lang.MsgAliasExists, name))
	}
	c.Shell.SetAlias(name, value)
	return nil, nil
}

func cmdRemoveAlias(c *Context) ([]*object.PSObject, error) {
	name, _ := c.Args.Str("Name")
	for n := range c.Shell.Aliases {
		if strings.EqualFold(n, name) {
			delete(c.Shell.Aliases, n)
		}
	}
	return nil, nil
}

func cmdExportAlias(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	var sb strings.Builder
	for _, n := range c.Shell.AllAliasNames() {
		fmt.Fprintf(&sb, "%s=%s\n", n, c.Shell.Aliases[n])
	}
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	if err := os.WriteFile(full, []byte(sb.String()), 0o644); err != nil {
		return errf(c, "Export-Alias : %v", err)
	}
	return nil, nil
}

func cmdImportAlias(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	data, err := os.ReadFile(full)
	if err != nil {
		return errf(c, "Import-Alias : %v", err)
	}
	for _, ln := range strings.Split(string(data), "\n") {
		if k, v, ok := strings.Cut(strings.TrimSpace(ln), "="); ok {
			c.Shell.SetAlias(k, v)
		}
	}
	return nil, nil
}

// ---- 注册 ----

func init() {
	Register("Get-Alias", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdGetAlias)
	Register("Set-Alias", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Value", Position: 1, PositionSet: true, Type: "string"},
	}, cmdSetAlias)
	Register("New-Alias", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Value", Position: 1, PositionSet: true, Type: "string"},
		{Name: "Force", Switch: true},
	}, cmdNewAlias)
	Register("Remove-Alias", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdRemoveAlias)
	Register("Export-Alias", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
	}, cmdExportAlias)
	Register("Import-Alias", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
	}, cmdImportAlias)
}

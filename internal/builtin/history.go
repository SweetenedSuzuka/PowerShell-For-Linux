package builtin

import (
	"fmt"

	"powershell/internal/object"
)

// history.go 实现历史类 cmdlet。

func cmdGetHistory(c *Context) ([]*object.PSObject, error) {
	var out []*object.PSObject
	for i, h := range c.Shell.History {
		o := object.Object("HistoryInfo", nil)
		o.AddProp("Id", int64(i+1))
		o.AddProp("CommandLine", h)
		o.Table = []object.Column{
			{Label: "Id", Align: "right"},
			{Label: "CommandLine", Align: "left"},
		}
		out = append(out, o)
	}
	return out, nil
}

func cmdClearHistory(c *Context) ([]*object.PSObject, error) {
	c.Shell.History = nil
	return nil, nil
}

func cmdAddHistory(c *Context) ([]*object.PSObject, error) {
	if len(c.Input) > 0 {
		for _, o := range c.Input {
			c.Shell.History = append(c.Shell.History, o.String())
		}
	} else if v := c.Args.Get("InputObject"); v != nil {
		for _, it := range v.ArrayItems() {
			c.Shell.History = append(c.Shell.History, it.String())
		}
	}
	return nil, nil
}

func cmdInvokeHistory(c *Context) ([]*object.PSObject, error) {
	var cmdText string
	if id, ok := c.Args.Int("Id"); ok && id > 0 {
		if int(id) <= len(c.Shell.History) {
			cmdText = c.Shell.History[id-1]
		}
	} else if v := c.Args.Get("InputObject"); v != nil {
		cmdText = v.String()
	} else if len(c.Shell.History) > 0 {
		cmdText = c.Shell.History[len(c.Shell.History)-1]
	}
	if cmdText == "" {
		return nil, nil
	}
	fmt.Fprintf(c.Stdout, "%s\n", cmdText)
	return c.Engine.RunSource(cmdText)
}

// ---- 注册 ----

func init() {
	Register("Get-History", nil, cmdGetHistory)
	Register("Clear-History", nil, cmdClearHistory)
	Register("Add-History", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
	}, cmdAddHistory)
	Register("Invoke-History", []ParamSpec{
		{Name: "Id", Type: "int"},
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
	}, cmdInvokeHistory)
}

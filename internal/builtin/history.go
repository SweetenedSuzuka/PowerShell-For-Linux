package builtin

import (
	"fmt"

	"powershell/internal/lang"
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

// cmdGetError 读最新错误记录展示（默认最新 1 条，-Newest 取多条，-InputObject 只接受单个）。
func cmdGetError(c *Context) ([]*object.PSObject, error) {
	// 超量位置实参无槽位可接（Newest 只占位置 0），报错而非静默忽略。
	if len(c.Args.Positional) > 0 {
		return errf(c, "%s", lang.T(lang.MsgPositionalParamNotFound, c.Args.Positional[0].String()))
	}
	newestArg := c.Args.Get("Newest")
	inArg := c.Args.Get("InputObject")
	hasPipe := len(c.Input) > 0
	// -Newest 与输入对象并存、管道与命名输入并存按参数集报错。
	if newestArg != nil && (hasPipe || inArg != nil) {
		return errf(c, "%s", lang.T(lang.MsgParamSetUnresolvable))
	}
	if hasPipe && inArg != nil {
		return errf(c, "%s", lang.T(lang.MsgInputObjectWithPipeline))
	}
	// 输入对象只接受单个，数组返回空。
	if inArg != nil {
		if inArg.IsArray() {
			return nil, nil
		}
		return []*object.PSObject{inArg}, nil
	}
	if hasPipe {
		return c.Input, nil
	}
	// 会话记录路径：条数默认 1，为 0 返回空，为负报错，超量取全集。
	n := int64(1)
	if newestArg != nil && !newestArg.IsNull() {
		var ok bool
		n, ok = newestArg.AsInt()
		if !ok {
			return errf(c, "%s", lang.T(lang.MsgBindConvertFail, newestArg.String(), "Newest", "int"))
		}
		if n < 1 {
			return errf(c, "%s", lang.T(lang.MsgNewestRange))
		}
	}
	recs := c.Shell.ErrorRecords
	if int64(len(recs)) > n {
		recs = recs[:n]
	}
	if len(recs) == 0 {
		return nil, nil
	}
	out := make([]*object.PSObject, len(recs))
	copy(out, recs)
	return out, nil
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
	Register("Get-Error", []ParamSpec{
		{Name: "Newest", Position: 0, PositionSet: true, Type: "int"},
		{Name: "InputObject", Type: "object"},
	}, cmdGetError)
}

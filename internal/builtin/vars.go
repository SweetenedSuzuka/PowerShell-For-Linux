package builtin

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"powershell/internal/object"
)

// ---- 变量 ----

func cmdNewVariable(c *Context) ([]*object.PSObject, error) {
	name, val := pairArgs(c, "Name")
	if name == "" {
		return nil, nil
	}
	if _, exists := c.Shell.Vars[name]; exists && !c.Args.Switch("Force") {
		return errf(c, "New-Variable : 变量 $%s 已存在。", name)
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
	if ns := c.Args.StringSlice("Name"); len(ns) > 0 {
		for _, n := range ns {
			delete(c.Shell.Vars, n)
		}
	} else if p := c.Args.Pos(0); p != nil {
		for _, it := range p.ArrayItems() {
			delete(c.Shell.Vars, it.String())
		}
	}
	return nil, nil
}

func cmdClearVariable(c *Context) ([]*object.PSObject, error) {
	name := ""
	if n, ok := c.Args.Str("Name"); ok {
		name = n
	} else if p := c.Args.Pos(0); p != nil {
		name = p.String()
	}
	if name != "" {
		_ = c.Shell.SetVar(name, object.Null())
	}
	return nil, nil
}

// ---- 别名 ----

func cmdNewAlias(c *Context) ([]*object.PSObject, error) {
	name := ""
	if n, ok := c.Args.Str("Name"); ok {
		name = n
	} else if p := c.Args.Pos(0); p != nil {
		name = p.String()
	}
	value := ""
	if v, ok := c.Args.Str("Value"); ok {
		value = v
	} else if p := c.Args.Pos(1); p != nil {
		value = p.String()
	}
	if name == "" || value == "" {
		return nil, nil
	}
	if _, exists := c.Shell.Aliases[name]; exists && !c.Args.Switch("Force") {
		return errf(c, "New-Alias : 别名 %s 已存在。", name)
	}
	c.Shell.SetAlias(name, value)
	return nil, nil
}

func cmdRemoveAlias(c *Context) ([]*object.PSObject, error) {
	name := ""
	if n, ok := c.Args.Str("Name"); ok {
		name = n
	} else if p := c.Args.Pos(0); p != nil {
		name = p.String()
	}
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

// ---- 历史 ----

func cmdAddHistory(c *Context) ([]*object.PSObject, error) {
	if len(c.Input) > 0 {
		for _, o := range c.Input {
			c.Shell.History = append(c.Shell.History, o.String())
		}
	} else if v := c.Args.Get("InputObject"); v != nil {
		for _, it := range v.ArrayItems() {
			c.Shell.History = append(c.Shell.History, it.String())
		}
	} else if p := c.Args.Pos(0); p != nil {
		c.Shell.History = append(c.Shell.History, p.String())
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
	} else if p := c.Args.Pos(0); p != nil {
		cmdText = p.String()
	} else if len(c.Shell.History) > 0 {
		cmdText = c.Shell.History[len(c.Shell.History)-1]
	}
	if cmdText == "" {
		return nil, nil
	}
	fmt.Fprintf(c.Stdout, "%s\n", cmdText)
	return c.Engine.RunSource(cmdText)
}

// ---- 主机与信息 ----

func cmdGetHost(c *Context) ([]*object.PSObject, error) {
	return []*object.PSObject{c.Shell.HostObject()}, nil
}

func cmdWriteVerbose(c *Context) ([]*object.PSObject, error) {
	msg := strings.Join(namedOrPosArgs(c, "Message"), " ")
	fmt.Fprintf(c.Stderr, "详细: %s\n", msg)
	return nil, nil
}

func cmdWriteWarning(c *Context) ([]*object.PSObject, error) {
	msg := strings.Join(namedOrPosArgs(c, "Message"), " ")
	fmt.Fprintf(c.Stderr, "警告: %s\n", msg)
	return nil, nil
}

func cmdWriteInformation(c *Context) ([]*object.PSObject, error) {
	msg := strings.Join(namedOrPosArgs(c, "MessageData"), " ")
	fmt.Fprintln(c.Stdout, msg)
	return nil, nil
}

func cmdWriteDebug(c *Context) ([]*object.PSObject, error) {
	msg := strings.Join(namedOrPosArgs(c, "Message"), " ")
	fmt.Fprintf(c.Stderr, "调试: %s\n", msg)
	return nil, nil
}

func cmdOutHost(c *Context) ([]*object.PSObject, error) {
	// 输出到主机（默认行为）：直接渲染，不进入管道
	objs := inputItems(c)
	_ = object.FormatOutput(c.Stdout, objs)
	return nil, nil
}

func cmdReadHost(c *Context) ([]*object.PSObject, error) {
	prompt := ""
	if p, ok := c.Args.Str("Prompt"); ok {
		prompt = p
	} else if p := c.Args.Pos(0); p != nil {
		prompt = p.String()
	}
	if prompt != "" {
		fmt.Fprint(c.Stdout, prompt)
	}
	reader := bufio.NewReader(c.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil && line == "" {
		return nil, nil
	}
	return []*object.PSObject{object.Str(strings.TrimRight(line, "\r\n"))}, nil
}

func cmdInvokeExpression(c *Context) ([]*object.PSObject, error) {
	var src string
	if v := c.Args.Get("Command"); v != nil {
		src = v.String()
	} else if p := c.Args.Pos(0); p != nil {
		src = p.String()
	} else if len(c.Input) > 0 {
		// 命名/位置都没给时才用管道输入（PowerShell 绑定顺序）
		var sb strings.Builder
		for _, o := range c.Input {
			sb.WriteString(o.String())
			sb.WriteByte('\n')
		}
		src = sb.String()
	}
	if src == "" {
		return nil, nil
	}
	return c.Engine.RunSource(src)
}

// ---- 注册 ----

func init() {
	Register("New-Variable", []ParamSpec{
		{Name: "Name", Position: 0, Type: "string"},
		{Name: "Value", Type: "object"},
		{Name: "Force", Switch: true},
	}, cmdNewVariable)
	Register("Remove-Variable", []ParamSpec{
		{Name: "Name", Position: 0, Type: "string"},
	}, cmdRemoveVariable)
	Register("Clear-Variable", []ParamSpec{
		{Name: "Name", Position: 0, Type: "string"},
	}, cmdClearVariable)
	Register("New-Alias", []ParamSpec{
		{Name: "Name", Position: 0, Type: "string"},
		{Name: "Value", Position: 1, Type: "string"},
		{Name: "Force", Switch: true},
	}, cmdNewAlias)
	Register("Remove-Alias", []ParamSpec{
		{Name: "Name", Position: 0, Type: "string"},
	}, cmdRemoveAlias)
	Register("Export-Alias", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
	}, cmdExportAlias)
	Register("Import-Alias", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
	}, cmdImportAlias)
	Register("Add-History", []ParamSpec{
		{Name: "InputObject", Position: 0, Type: "object"},
	}, cmdAddHistory)
	Register("Invoke-History", []ParamSpec{
		{Name: "Id", Type: "int"},
		{Name: "InputObject", Position: 0, Type: "object"},
	}, cmdInvokeHistory)
	Register("Get-Host", nil, cmdGetHost)
	Register("Write-Verbose", []ParamSpec{
		{Name: "Message", Position: 0, Type: "string"},
	}, cmdWriteVerbose)
	Register("Write-Warning", []ParamSpec{
		{Name: "Message", Position: 0, Type: "string"},
	}, cmdWriteWarning)
	Register("Write-Information", []ParamSpec{
		{Name: "MessageData", Position: 0, Type: "string"},
	}, cmdWriteInformation)
	Register("Write-Debug", []ParamSpec{
		{Name: "Message", Position: 0, Type: "string"},
	}, cmdWriteDebug)
	Register("Out-Host", []ParamSpec{
		{Name: "InputObject", Position: 0, Type: "object"},
	}, cmdOutHost)
	Register("Read-Host", []ParamSpec{
		{Name: "Prompt", Position: 0, Type: "string"},
	}, cmdReadHost)
	Register("Invoke-Expression", []ParamSpec{
		{Name: "Command", Position: 0, Type: "string"},
	}, cmdInvokeExpression)
}

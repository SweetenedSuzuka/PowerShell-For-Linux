package builtin

import (
	"bufio"
	"fmt"
	"strings"

	"powershell/internal/lang"
	"powershell/internal/object"
)

// host.go 实现主机与信息输出类 cmdlet。

func cmdGetHost(c *Context) ([]*object.PSObject, error) {
	return []*object.PSObject{c.Shell.HostObject()}, nil
}

func cmdWriteVerbose(c *Context) ([]*object.PSObject, error) {
	msg := strings.Join(namedOrPosArgs(c, "Message"), " ")
	fmt.Fprintf(c.Stderr, "%s %s\n", lang.T(lang.MsgWriteVerbosePrefix), msg)
	return nil, nil
}

func cmdWriteWarning(c *Context) ([]*object.PSObject, error) {
	msg := strings.Join(namedOrPosArgs(c, "Message"), " ")
	fmt.Fprintf(c.Stderr, "%s %s\n", lang.T(lang.MsgWriteWarningPrefix), msg)
	return nil, nil
}

func cmdWriteInformation(c *Context) ([]*object.PSObject, error) {
	msg := strings.Join(namedOrPosArgs(c, "MessageData"), " ")
	fmt.Fprintln(c.console(), msg)
	return nil, nil
}

func cmdWriteDebug(c *Context) ([]*object.PSObject, error) {
	msg := strings.Join(namedOrPosArgs(c, "Message"), " ")
	fmt.Fprintf(c.Stderr, "%s %s\n", lang.T(lang.MsgWriteDebugPrefix), msg)
	return nil, nil
}

func cmdOutHost(c *Context) ([]*object.PSObject, error) {
	// 输出到主机（默认行为）：直接渲染，不进入管道
	objs := inputItems(c)
	_ = object.FormatOutput(c.console(), objs)
	return nil, nil
}

func cmdClearHost(c *Context) ([]*object.PSObject, error) {
	fmt.Fprint(c.console(), "\x1b[2J\x1b[H")
	return nil, nil
}

func cmdReadHost(c *Context) ([]*object.PSObject, error) {
	// 非交互运行不能读输入，直接报错。
	if c.Shell.NonInteractive {
		return nil, fmt.Errorf("%s", lang.T(lang.MsgReadHostNonInteractive))
	}
	prompt, _ := c.Args.Str("Prompt")
	if prompt != "" {
		fmt.Fprint(c.console(), prompt)
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
	Register("Get-Host", nil, cmdGetHost)
	Register("Write-Verbose", []ParamSpec{
		{Name: "Message", Position: 0, PositionSet: true, Type: "string"},
	}, cmdWriteVerbose)
	Register("Write-Warning", []ParamSpec{
		{Name: "Message", Position: 0, PositionSet: true, Type: "string"},
	}, cmdWriteWarning)
	Register("Write-Information", []ParamSpec{
		{Name: "MessageData", Position: 0, PositionSet: true, Type: "string"},
	}, cmdWriteInformation)
	Register("Write-Debug", []ParamSpec{
		{Name: "Message", Position: 0, PositionSet: true, Type: "string"},
	}, cmdWriteDebug)
	Register("Out-Host", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
	}, cmdOutHost)
	Register("Clear-Host", nil, cmdClearHost)
	Register("Read-Host", []ParamSpec{
		{Name: "Prompt", Position: 0, PositionSet: true, Type: "string"},
	}, cmdReadHost)
	Register("Invoke-Expression", []ParamSpec{
		{Name: "Command", Position: 0, PositionSet: true, Type: "string"},
	}, cmdInvokeExpression)
}

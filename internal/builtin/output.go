package builtin

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"powershell/internal/object"
)

// ---- 输出与文本 ----

func cmdWriteOutput(c *Context) ([]*object.PSObject, error) {
	return inputItems(c), nil
}

func cmdWriteHost(c *Context) ([]*object.PSObject, error) {
	parts := namedOrPosArgs(c, "Object")
	noNewline := c.Args.Switch("NoNewline")
	text := strings.Join(parts, " ")
	if noNewline {
		fmt.Fprint(c.Stdout, text)
	} else {
		fmt.Fprintln(c.Stdout, text)
	}
	return nil, nil
}

func cmdWriteError(c *Context) ([]*object.PSObject, error) {
	parts := namedOrPosArgs(c, "Message")
	if len(c.Input) > 0 {
		for _, o := range c.Input {
			parts = append(parts, o.String())
		}
	}
	fmt.Fprintf(c.Stderr, "错误: %s\n", strings.Join(parts, " "))
	c.Shell.LastSuccess = false
	return nil, nil
}

func cmdOutNull(c *Context) ([]*object.PSObject, error) {
	return nil, nil
}

func cmdOutFile(c *Context) ([]*object.PSObject, error) {
	path := ""
	if p, ok := c.Args.Str("FilePath"); ok {
		path = p
	} else if p := c.Args.Pos(0); p != nil {
		path = p.String()
	}
	if path == "" {
		return nil, nil
	}
	appendMode := c.Args.Switch("Append")
	objs := c.Input
	var buf strings.Builder
	_ = object.FormatOutput(&buf, objs)
	flags := os.O_WRONLY | os.O_CREATE
	if appendMode {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	f, err := os.OpenFile(full, flags, 0o644)
	if err != nil {
		return errf(c, "Out-File : 无法写入 %s。", path)
	}
	defer f.Close()
	_, _ = f.WriteString(buf.String())
	return nil, nil
}

func cmdFormatTable(c *Context) ([]*object.PSObject, error) {
	props := c.Args.StringSlice("Property")
	if len(props) == 0 {
		// 位置参数当属性列表（如 ft Name,Length）
		if p := c.Args.Pos(0); p != nil {
			for _, it := range p.ArrayItems() {
				props = append(props, it.String())
			}
		}
	}
	objs := c.Input
	if len(objs) == 0 {
		if p := c.Args.Pos(1); p != nil {
			objs = []*object.PSObject{p}
		}
	}
	_ = object.FormatTableTo(c.Stdout, objs, props)
	return nil, nil
}

func cmdFormatList(c *Context) ([]*object.PSObject, error) {
	props := c.Args.StringSlice("Property")
	if len(props) == 0 {
		if p := c.Args.Pos(0); p != nil {
			for _, it := range p.ArrayItems() {
				props = append(props, it.String())
			}
		}
	}
	objs := c.Input
	if len(objs) == 0 {
		if p := c.Args.Pos(1); p != nil {
			objs = []*object.PSObject{p}
		}
	}
	_ = object.FormatListTo(c.Stdout, objs, props)
	return nil, nil
}

func cmdFormatWide(c *Context) ([]*object.PSObject, error) {
	objs := c.Input
	width := 0
	if w, ok := c.Args.Int("Column"); ok {
		width = int(w)
	}
	// -Property <string>（位置 0）：取对象的该属性显示
	prop := ""
	if p, ok := c.Args.Str("Property"); ok {
		prop = p
	} else if p := c.Args.Pos(0); p != nil {
		prop = p.String()
	}
	_ = object.FormatWideTo(c.Stdout, objs, width, prop)
	return nil, nil
}

func cmdSelectString(c *Context) ([]*object.PSObject, error) {
	pattern := firstArg(c, "Pattern")
	if pattern == "" {
		return nil, nil
	}
	path := ""
	if p, ok := c.Args.Str("Path"); ok {
		path = p
	} else if len(c.Args.Positional) >= 2 {
		path = c.Args.Positional[1].String()
	}
	var out []*object.PSObject
	scan := func(name string, text string) {
		for i, line := range strings.Split(strings.TrimSuffix(text, "\n"), "\n") {
			if line == "" {
				continue
			}
			matched := false
			if c.Args.Switch("SimpleMatch") {
				matched = strings.Contains(line, pattern)
			} else if re, err := regexp.Compile(pattern); err == nil {
				matched = re.MatchString(line)
			}
			if matched {
				o := object.Object("Microsoft.PowerShell.Commands.MatchInfo", nil)
				o.AddProp("LineNumber", int64(i+1))
				o.AddProp("Line", line)
				o.AddProp("Path", name)
				o.AddProp("Pattern", pattern)
				out = append(out, o)
			}
		}
	}
	if path != "" {
		full, derr := resolvePath(c, path)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		data, err := os.ReadFile(full)
		if err != nil {
			return errf(c, "Select-String : 找不到路径 %s。", path)
		}
		scan(path, string(data))
	} else if v := c.Args.Get("InputObject"); v != nil {
		for _, it := range v.ArrayItems() {
			scan("", it.String())
		}
	} else {
		for _, o := range c.Input {
			scan("", o.String())
		}
	}
	return out, nil
}

// ---- 注册 ----

func init() {
	Register("Write-Output", []ParamSpec{
		{Name: "InputObject", Position: 0, Type: "object"},
	}, cmdWriteOutput)
	Register("Write-Host", []ParamSpec{
		{Name: "Object", Position: 0, Type: "object"},
		{Name: "NoNewline", Switch: true},
	}, cmdWriteHost)
	Register("Write-Error", []ParamSpec{
		{Name: "Message", Position: 0, Type: "string"},
	}, cmdWriteError)
	Register("Out-Null", []ParamSpec{
		{Name: "InputObject", Position: 0, Type: "object"},
	}, cmdOutNull)
	Register("Out-File", []ParamSpec{
		{Name: "FilePath", Position: 0, Type: "path"},
		{Name: "Append", Switch: true},
		{Name: "Encoding", Type: "string"},
	}, cmdOutFile)
	Register("Format-Table", []ParamSpec{
		{Name: "Property", Position: 0, Type: "string[]"},
		{Name: "AutoSize", Switch: true},
	}, cmdFormatTable)
	Register("Format-List", []ParamSpec{
		{Name: "Property", Position: 0, Type: "string[]"},
	}, cmdFormatList)
	Register("Format-Wide", []ParamSpec{
		{Name: "Property", Position: 0, Type: "string"},
		{Name: "Column", Type: "int"},
	}, cmdFormatWide)
	Register("Select-String", []ParamSpec{
		{Name: "Pattern", Position: 0, Type: "string"},
		{Name: "Path", Position: 1, Type: "path"},
		{Name: "SimpleMatch", Switch: true},
		{Name: "InputObject", Type: "object"},
	}, cmdSelectString)
}

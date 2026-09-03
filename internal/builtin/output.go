package builtin

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"os"
	"regexp"
	"strings"

	"powershell/internal/lang"
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
	fmt.Fprintf(c.Stderr, "%s %s\n", lang.T(lang.MsgWriteErrorPrefix), strings.Join(parts, " "))
	c.Shell.LastSuccess = false
	return nil, nil
}

func cmdOutNull(c *Context) ([]*object.PSObject, error) {
	return nil, nil
}

func cmdOutFile(c *Context) ([]*object.PSObject, error) {
	path, _ := c.Args.Str("FilePath")
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
		return errf(c, "%s", lang.T(lang.MsgCannotWrite, path))
	}
	defer f.Close()
	enc, _ := c.Args.Str("Encoding")
	// 追加模式不写 BOM，避免重复追加时多次写文件头
	_, _ = f.Write(encodeText(enc, buf.String(), !appendMode))
	return nil, nil
}

func cmdFormatTable(c *Context) ([]*object.PSObject, error) {
	props := c.Args.StringSlice("Property")
	objs := c.Input
	if len(objs) == 0 {
		// 无管道输入：剩余位置实参（超量/未声明槽位）当数据
		if len(c.Args.Positional) > 0 {
			objs = c.Args.Positional
		}
	}
	_ = object.FormatTableTo(c.Stdout, objs, props)
	return nil, nil
}

func cmdFormatList(c *Context) ([]*object.PSObject, error) {
	props := c.Args.StringSlice("Property")
	objs := c.Input
	if len(objs) == 0 {
		if len(c.Args.Positional) > 0 {
			objs = c.Args.Positional
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
	// -Property（命名或位置）：取对象的该属性显示
	prop, _ := c.Args.Str("Property")
	_ = object.FormatWideTo(c.Stdout, objs, width, prop)
	return nil, nil
}

func cmdSelectString(c *Context) ([]*object.PSObject, error) {
	pattern := firstArg(c, "Pattern")
	if pattern == "" {
		return nil, nil
	}
	path, _ := c.Args.Str("Path")
	caseSensitive := c.Args.Switch("CaseSensitive")
	simple := c.Args.Switch("SimpleMatch")
	quiet := c.Args.Switch("Quiet")
	var out []*object.PSObject
	// 正则预编译一次：默认大小写不敏感（加 (?i)），-CaseSensitive 时原样
	var re *regexp.Regexp
	if !simple {
		p := pattern
		if !caseSensitive {
			p = "(?i)" + pattern
		}
		re, _ = regexp.Compile(p)
	}
	matchLine := func(line string) bool {
		if simple {
			if caseSensitive {
				return strings.Contains(line, pattern)
			}
			return strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		}
		return re != nil && re.MatchString(line)
	}
	emit := func(name string, num int64, line string) {
		o := object.Object("Microsoft.PowerShell.Commands.MatchInfo", nil)
		o.AddProp("LineNumber", num)
		o.AddProp("Line", line)
		o.AddProp("Path", name)
		o.AddProp("Pattern", pattern)
		out = append(out, o)
	}
	if path != "" {
		full, derr := resolvePath(c, path)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		data, err := os.ReadFile(full)
		if err != nil {
			return errf(c, "%s", lang.T(lang.MsgPathNotFoundFmt, path))
		}
		// 文件输入逐行扫描，空行也计入行号
		for i, line := range strings.Split(strings.TrimSuffix(string(data), "\n"), "\n") {
			if matchLine(line) {
				emit(path, int64(i+1), line)
			}
		}
	} else if v := c.Args.Get("InputObject"); v != nil {
		// 管道/参数输入：每个对象整体作为一行参与匹配，LineNumber 是对象在流中的序号
		for n, it := range v.ArrayItems() {
			line := it.String()
			if matchLine(line) {
				emit("", int64(n+1), line)
			}
		}
	} else {
		for n, o := range c.Input {
			line := o.String()
			if matchLine(line) {
				emit("", int64(n+1), line)
			}
		}
	}
	// -Quiet 只报告是否命中：命中输出单个 $true，未命中无输出（但是是 $null，不是 $false）。
	if quiet {
		if len(out) == 0 {
			return nil, nil
		}
		return []*object.PSObject{object.Bool(true)}, nil
	}
	return out, nil
}

func cmdOutString(c *Context) ([]*object.PSObject, error) {
	var buf bytes.Buffer
	_ = object.FormatOutput(&buf, c.Input)
	text := buf.String()
	if c.Args.Switch("Stream") {
		var out []*object.PSObject
		for _, ln := range strings.Split(strings.TrimRight(text, "\n"), "\n") {
			out = append(out, object.Str(ln))
		}
		return out, nil
	}
	return []*object.PSObject{object.Str(text)}, nil
}

func cmdTeeObject(c *Context) ([]*object.PSObject, error) {
	path := firstArg(c, "FilePath")
	appendMode := c.Args.Switch("Append")
	if path != "" {
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
		if err == nil {
			for _, o := range c.Input {
				fmt.Fprintln(f, o.String())
			}
			f.Close()
		}
	}
	return c.Input, nil
}

// hexTypeLabel 是 Format-Hex 标签里的类型名（短名与全名）。
func hexTypeLabel(o *object.PSObject) (string, string) {
	switch o.TypeName {
	case "String":
		return "String", "System.String"
	case "Int":
		return "Int32", "System.Int32"
	case "Double":
		return "Double", "System.Double"
	case "Boolean":
		return "Boolean", "System.Boolean"
	}
	return o.TypeName, o.TypeName
}

// formatHexSegment 把一段字节按 PowerShell 格式渲染：16 位偏移、固定宽字节区、ASCII 对照列。
func formatHexSegment(sb *strings.Builder, data []byte) {
	sb.WriteString("          Offset Bytes                                           Ascii\n")
	sb.WriteString("                 00 01 02 03 04 05 06 07 08 09 0A 0B 0C 0D 0E 0F\n")
	sb.WriteString("          ------ ----------------------------------------------- -----\n")
	for i := 0; i < len(data); i += 16 {
		end := i + 16
		if end > len(data) {
			end = len(data)
		}
		fmt.Fprintf(sb, "%016X ", i)
		for j := i; j < i+16; j++ {
			if j < end {
				fmt.Fprintf(sb, "%02X ", data[j])
			} else {
				sb.WriteString("   ")
			}
		}
		for j := i; j < end; j++ {
			if data[j] >= 0x20 && data[j] < 0x7F {
				sb.WriteByte(data[j])
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
}

func cmdFormatHex(c *Context) ([]*object.PSObject, error) {
	var sb strings.Builder
	emit := func(label string, short, full string, data []byte) {
		if full == "" {
			fmt.Fprintf(&sb, "\n   Label: %s\n\n", label)
		} else {
			sum := fmt.Sprintf("%08X", crc32.ChecksumIEEE([]byte(label+full)))
			fmt.Fprintf(&sb, "\n   Label: %s (%s) <%s>\n\n", label, full, sum)
		}
		formatHexSegment(&sb, data)
	}
	if path, _ := c.Args.Str("Path"); path != "" {
		full, derr := resolvePath(c, path)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		data, err := os.ReadFile(full)
		if err != nil {
			return errf(c, "%s", lang.T(lang.MsgFormatHexNotFound, path))
		}
		emit(path, "", "", data)
		return []*object.PSObject{object.Str(strings.Trim(sb.String(), "\n"))}, nil
	}
	for _, o := range inputItems(c) {
		short, full := hexTypeLabel(o)
		emit(short, short, full, []byte(o.String()))
	}
	return []*object.PSObject{object.Str(strings.Trim(sb.String(), "\n"))}, nil
}

func cmdJoinString(c *Context) ([]*object.PSObject, error) {
	// 默认分隔符是空串
	sep := ""
	if s, ok := c.Args.Str("Separator"); ok {
		sep = s
	}
	prop, _ := c.Args.Str("Property")
	format, _ := c.Args.Str("FormatString")
	prefix, _ := c.Args.Str("OutputPrefix")
	suffix, _ := c.Args.Str("OutputSuffix")
	doubleQuote := c.Args.Switch("DoubleQuote")
	singleQuote := c.Args.Switch("SingleQuote")
	var parts []string
	for _, o := range inputItems(c) {
		v := o
		// -Property 取该属性的值转文本
		if prop != "" {
			if pv, ok := o.PropValue(prop); ok {
				v = pv
			}
		}
		text := v.String()
		switch {
		case doubleQuote:
			text = `"` + text + `"`
		case singleQuote:
			text = "'" + text + "'"
		case format != "":
			// 单遍解析复合格式：{{ 与 }} 为字面花括号，{0} 为当前对象
			var fb strings.Builder
			for i := 0; i < len(format); {
				if format[i] == '{' && i+1 < len(format) && format[i+1] == '{' {
					fb.WriteByte('{')
					i += 2
					continue
				}
				if format[i] == '}' && i+1 < len(format) && format[i+1] == '}' {
					fb.WriteByte('}')
					i += 2
					continue
				}
				if format[i] == '{' {
					if end := strings.IndexByte(format[i:], '}'); end > 0 {
						fb.WriteString(v.String())
						i += end + 1
						continue
					}
				}
				fb.WriteByte(format[i])
				i++
			}
			text = fb.String()
		}
		parts = append(parts, text)
	}
	return []*object.PSObject{object.Str(prefix + strings.Join(parts, sep) + suffix)}, nil
}

// ---- 注册 ----

func init() {
	Register("Write-Output", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
	}, cmdWriteOutput)
	Register("Write-Host", []ParamSpec{
		{Name: "Object", Position: 0, PositionSet: true, Type: "object"},
		{Name: "NoNewline", Switch: true},
	}, cmdWriteHost)
	Register("Write-Error", []ParamSpec{
		{Name: "Message", Position: 0, PositionSet: true, Type: "string"},
	}, cmdWriteError)
	Register("Out-Null", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
	}, cmdOutNull)
	Register("Out-File", []ParamSpec{
		{Name: "FilePath", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Append", Switch: true},
		{Name: "Encoding", Type: "string"},
	}, cmdOutFile)
	Register("Format-Table", []ParamSpec{
		{Name: "Property", Position: 0, PositionSet: true, Type: "string[]"},
		{Name: "AutoSize", Switch: true},
	}, cmdFormatTable)
	Register("Format-List", []ParamSpec{
		{Name: "Property", Position: 0, PositionSet: true, Type: "string[]"},
	}, cmdFormatList)
	Register("Format-Wide", []ParamSpec{
		{Name: "Property", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Column", Type: "int"},
	}, cmdFormatWide)
	Register("Select-String", []ParamSpec{
		{Name: "Pattern", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Path", Position: 1, PositionSet: true, Type: "path"},
		{Name: "SimpleMatch", Switch: true},
		{Name: "CaseSensitive", Switch: true},
		{Name: "Quiet", Switch: true},
		{Name: "InputObject", Type: "object"},
	}, cmdSelectString)
	Register("Out-String", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
		{Name: "Stream", Switch: true},
	}, cmdOutString)
	Register("Tee-Object", []ParamSpec{
		{Name: "FilePath", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Append", Switch: true},
	}, cmdTeeObject)
	Register("Format-Hex", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "InputObject", Type: "object"},
	}, cmdFormatHex)
	Register("Join-String", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
		{Name: "Separator", Position: 1, PositionSet: true, Type: "string"},
		{Name: "Property", Type: "string"},
		{Name: "FormatString", Type: "string"},
		{Name: "OutputPrefix", Type: "string"},
		{Name: "OutputSuffix", Type: "string"},
		{Name: "DoubleQuote", Switch: true},
		{Name: "SingleQuote", Switch: true},
	}, cmdJoinString)
}

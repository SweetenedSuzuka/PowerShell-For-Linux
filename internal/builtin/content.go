package builtin

import (
	"os"
	"strings"

	"powershell/internal/lang"
	"powershell/internal/object"
)

// content.go 实现文件内容读写类 cmdlet。

func cmdGetContent(c *Context) ([]*object.PSObject, error) {
	// 路径：-Path（命名或位置，数组摊平）加超量位置实参，逐个文件读取
	var paths []string
	if v := c.Args.Get("Path"); v != nil {
		for _, it := range v.ArrayItems() {
			paths = append(paths, it.String())
		}
	}
	for _, p := range c.Args.Positional {
		for _, it := range p.ArrayItems() {
			paths = append(paths, it.String())
		}
	}
	if len(paths) == 0 {
		if len(c.Input) > 0 {
			return c.Input, nil
		}
		return nil, nil
	}
	raw := c.Args.Switch("Raw")
	var out []*object.PSObject
	for _, path := range paths {
		full, derr := resolvePath(c, path)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		data, err := os.ReadFile(full)
		if err != nil {
			return errf(c, "%s", lang.T(lang.MsgPathNotFoundFmt, path))
		}
		// 去掉 UTF-8 BOM，避免首行带头码。
		text := StripUTF8BOM(string(data))
		if raw {
			out = append(out, object.Str(text))
			continue
		}
		lines := strings.Split(strings.TrimSuffix(text, "\n"), "\n")
		if len(lines) == 1 && lines[0] == "" {
			lines = nil
		}
		// -TotalCount / -Tail 对每个文件分别生效
		if total, ok := c.Args.Int("TotalCount"); ok && total >= 0 {
			if int(total) < len(lines) {
				lines = lines[:total]
			}
		}
		if tail, ok := c.Args.Int("Tail"); ok && tail >= 0 {
			if int(tail) < len(lines) {
				lines = lines[len(lines)-int(tail):]
			}
		}
		for _, l := range lines {
			out = append(out, object.Str(l))
		}
	}
	return out, nil
}

func cmdSetContent(c *Context) ([]*object.PSObject, error) {
	path, val := pathAndValue(c)
	if path == "" {
		return nil, nil
	}
	var content []*object.PSObject
	if len(c.Input) > 0 {
		content = c.Input
	} else if val != nil {
		content = []*object.PSObject{val}
	}
	var sb strings.Builder
	for _, o := range content {
		sb.WriteString(o.String())
		sb.WriteByte('\n')
	}
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	var wi whatIfCollector
	wi.cmdlet = "Set-Content"
	wi.c = c
	var yesAll, noAll bool
	if wi.hit(full) {
		out, _ := wi.result()
		return out, nil
	}
	if confirmSkip(c, "Set-Content", full, &yesAll, &noAll) {
		return nil, nil
	}
	enc, _ := c.Args.Str("Encoding")
	if err := os.WriteFile(full, encodeText(enc, sb.String(), true), 0o644); err != nil {
		return errf(c, "%s", lang.T(lang.MsgCannotWrite, path))
	}
	return nil, nil
}

func cmdAddContent(c *Context) ([]*object.PSObject, error) {
	path, val := pathAndValue(c)
	if path == "" {
		return nil, nil
	}
	var content []*object.PSObject
	if len(c.Input) > 0 {
		content = c.Input
	} else if val != nil {
		content = []*object.PSObject{val}
	}
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	var wi whatIfCollector
	wi.cmdlet = "Add-Content"
	wi.c = c
	var yesAll, noAll bool
	if wi.hit(full) {
		out, _ := wi.result()
		return out, nil
	}
	if confirmSkip(c, "Add-Content", full, &yesAll, &noAll) {
		return nil, nil
	}
	f, err := os.OpenFile(full, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return errf(c, "%s", lang.T(lang.MsgCannotOpen, path))
	}
	defer f.Close()
	// 已有内容的旧文件不再写 BOM，与新建文件行为一致。
	var sb strings.Builder
	for _, o := range content {
		sb.WriteString(o.String())
		sb.WriteByte('\n')
	}
	enc, _ := c.Args.Str("Encoding")
	bom := true
	if fi, serr := os.Stat(full); serr == nil && fi.Size() > 0 {
		bom = false
	}
	_, _ = f.Write(encodeText(enc, sb.String(), bom))
	return nil, nil
}

// ---- 注册 ----

func init() {
	Register("Get-Content", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Raw", Switch: true},
		{Name: "TotalCount", Type: "int"},
		{Name: "Tail", Type: "int"},
	}, cmdGetContent)
	Register("Set-Content", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Value", Position: 1, PositionSet: true, Type: "object"},
		{Name: "Encoding", Type: "string"},
	}, cmdSetContent)
	Register("Add-Content", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Value", Position: 1, PositionSet: true, Type: "object"},
		{Name: "Encoding", Type: "string"},
	}, cmdAddContent)
}

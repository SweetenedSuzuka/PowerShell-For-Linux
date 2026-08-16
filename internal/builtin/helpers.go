package builtin

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"powershell/internal/object"
	"powershell/internal/shell"
)

// timeSpanObj 构造字段齐全的 TimeSpan 对象（New-TimeSpan/Measure-Command/Get-Uptime 共用）。
func timeSpanObj(d time.Duration) *object.PSObject {
	o := object.Object("System.TimeSpan", int64(d))
	o.AddProp("Days", int64(d/(24*time.Hour)))
	o.AddProp("Hours", int64(d/time.Hour%24))
	o.AddProp("Minutes", int64(d/time.Minute%60))
	o.AddProp("Seconds", int64(d/time.Second%60))
	o.AddProp("TotalMilliseconds", float64(d.Milliseconds()))
	o.AddProp("TotalSeconds", d.Seconds())
	o.AddProp("TotalMinutes", d.Minutes())
	o.AddProp("TotalHours", d.Hours())
	return o
}

// resolvePath 把相对路径解析为基于会话当前目录的绝对路径（盘符/~/相对都在 shell.ResolvePath）。
func resolvePath(c *Context, p string) (string, error) {
	return shell.ResolvePath(c.Shell.Cwd, p)
}

// firstPathArg 取 -Path 或首个位置参数。
func firstPathArg(c *Context) string {
	return firstArg(c, "Path")
}

// firstArg 取命名参数（位置实参已由 Bind 按规格中心化映射到同名命名参数）。
func firstArg(c *Context, name string) string {
	if p, ok := c.Args.Str(name); ok && p != "" {
		return p
	}
	return ""
}

// namedOrPosArgs 取命名参数（可能是数组）的值；未给出则取全部位置参数。
// 用于 -Object/-Message/-MessageData 这类既接受命名也接受位置的参数。
func namedOrPosArgs(c *Context, name string) []string {
	if v := c.Args.Get(name); v != nil {
		var out []string
		for _, it := range v.ArrayItems() {
			out = append(out, it.String())
		}
		return out
	}
	var out []string
	for _, p := range c.Args.Positional {
		out = append(out, p.String())
	}
	return out
}

// pathAndValue 解析"路径 + 值"类命令（Set-Content/Add-Content/Set-Item）。
// 位置实参已由 Bind 中心化映射到 -Path/-Value（跳过已命名的槽位）。
// 缺值时的旧行为：显式命名 -Path 而无值视为命令没写完，静默不动；
// 位置形式（如 Set-Content foo）允许写空文件。用 PosMapped 区分两者。
func pathAndValue(c *Context) (string, *object.PSObject) {
	path, _ := c.Args.Str("Path")
	val := c.Args.Get("Value")
	if path == "" {
		return "", nil
	}
	if val == nil && !c.Args.PosMapped["Path"] {
		return "", nil
	}
	return path, val
}

// pairArgs 解析"名称 + 值"类命令（New-Variable/Set-Variable）。
// 位置实参已由 Bind 中心化映射到 -Name/-Value（跳过已命名的槽位）。
func pairArgs(c *Context, keyName string) (string, *object.PSObject) {
	key, _ := c.Args.Str(keyName)
	val := c.Args.Get("Value")
	if key == "" && val == nil {
		return "", nil
	}
	return key, val
}

// expandWildcard 展开路径通配符；无通配符则原样（相对路径基于 cwd）。
// 通配无匹配时返回空（各命令自然无操作），避免把原始模式当字面路径处理。
func expandWildcard(c *Context, pattern string) ([]string, error) {
	if !strings.ContainsAny(pattern, "*?[") {
		full, err := resolvePath(c, pattern)
		if err != nil {
			return nil, err
		}
		return []string{full}, nil
	}
	p := pattern
	// 通配模式也可能带盘符前缀（如 C:\*.txt），先归一化
	if np, err := shell.DrivePath(p); err != nil {
		return nil, err
	} else if np != p {
		p = np
	}
	if !filepath.IsAbs(p) {
		p = filepath.Join(c.Shell.Cwd, p)
	}
	matches, err := filepath.Glob(p)
	if err != nil || len(matches) == 0 {
		return nil, nil
	}
	return matches, nil
}

// errf 构造一条格式化错误（写 stderr）并标记 $? 为 false。
func errf(c *Context, format string, args ...any) ([]*object.PSObject, error) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(c.Stderr, "%s : %s\n", c.Shell.StyleName(), msg)
	c.Shell.LastSuccess = false
	return nil, nil
}

// pathList 汇总管道输入或 -Path / 位置参数中的路径列表（数组摊平，支持"可多个"）。
func pathList(c *Context) []string {
	var paths []string
	if len(c.Input) > 0 {
		for _, o := range c.Input {
			paths = append(paths, o.String())
		}
	}
	if v := c.Args.Get("Path"); v != nil {
		for _, it := range v.ArrayItems() {
			paths = append(paths, it.String())
		}
	} else if len(c.Args.Positional) > 0 {
		for _, it := range c.Args.Positional[0].ArrayItems() {
			paths = append(paths, it.String())
		}
	}
	return paths
}

// inputItems 取输入对象：优先管道输入，其次 -InputObject（含位置映射），再补齐剩余位置实参（未声明位置槽位的实参，数组摊平）。
func inputItems(c *Context) []*object.PSObject {
	if len(c.Input) > 0 {
		return c.Input
	}
	var out []*object.PSObject
	if v := c.Args.Get("InputObject"); v != nil {
		out = append(out, v.ArrayItems()...)
	}
	for _, p := range c.Args.Positional {
		if p.IsArray() {
			out = append(out, p.ArrayItems()...)
		} else {
			out = append(out, p)
		}
	}
	return out
}

package builtin

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"powershell/internal/lang"
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

// whatIfCollector 累计 -WhatIf 的预演输出；多目标命令跨循环收集后一并返回。
type whatIfCollector struct {
	c      *Context
	cmdlet string
	out    []*object.PSObject
}

// hit 报告本目标是否命中 -WhatIf；命中时登记预演行并指示跳过实际变更。
func (w *whatIfCollector) hit(target string) bool {
	if !w.c.Args.Switch("WhatIf") {
		return false
	}
	w.out = append(w.out, object.Str(lang.T(lang.MsgWhatIfPerform, target, w.cmdlet)))
	return true
}

// result 返回收集到的预演输出；没有任何命中时 ok 为 false。
func (w *whatIfCollector) result() ([]*object.PSObject, bool) {
	return w.out, len(w.out) > 0
}

// confirmSkip 处理 -Confirm：需要确认时提示用户选择，返回是否应跳过本次变更。
// yesAll 与 noAll 由调用方在同一命令的多目标循环间传递，使 A/L 的选择覆盖后续目标。
// 输入结束（EOF）按拒绝处理；未指定 -Confirm 且此前没有全部性选择时直接执行。
func confirmSkip(c *Context, operation string, target string, yesAll, noAll *bool) bool {
	if *noAll {
		return true
	}
	if *yesAll || !c.Args.Switch("Confirm") {
		return false
	}
	for {
		fmt.Fprint(c.Stdout, lang.T(lang.MsgConfirmPrompt, operation, target))
		answer, ok := readLineBytes(c.Stdin)
		if !ok {
			// 流结束没有等到回答：按拒绝处理
			return true
		}
		switch answer {
		case "", "y", "yes":
			return false
		case "a", "all":
			*yesAll = true
			return false
		case "n", "no":
			return true
		case "l":
			*noAll = true
			return true
		}
		// 其余输入不进行识别，重新提示
	}
}

// readLineBytes 从 r 逐字节读到换行为止；ok 为 false 表示流已结束且没有读到换行。
// 读取不做缓冲，后续提示所需的输入不会被提前消费。
func readLineBytes(r io.Reader) (string, bool) {
	var sb []byte
	b := make([]byte, 1)
	for {
		n, err := r.Read(b)
		if n == 0 || err != nil {
			return strings.TrimSpace(string(sb)), false
		}
		if b[0] == '\n' {
			return strings.TrimSpace(string(sb)), true
		}
		if b[0] != '\r' {
			sb = append(sb, b[0])
		}
	}
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

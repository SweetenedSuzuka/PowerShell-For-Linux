package builtin

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"powershell/internal/lang"
	"powershell/internal/object"
)

// ---- 路径与导航 ----

func cmdTestPath(c *Context) ([]*object.PSObject, error) {
	// 路径：-Path（命名或位置，数组摊平）加超量位置实参，命中任意一个即 True
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
		return []*object.PSObject{object.Bool(false)}, nil
	}
	// -PathType 过滤：Leaf 只认文件、Container 只认目录、Any/缺省不限制
	pt, _ := c.Args.Str("PathType")
	typeMatch := func(p string) bool {
		if pt == "" || strings.EqualFold(pt, "Any") {
			return true
		}
		st, err := os.Stat(p)
		if err != nil {
			return false
		}
		if strings.EqualFold(pt, "Leaf") {
			return !st.IsDir()
		}
		if strings.EqualFold(pt, "Container") {
			return st.IsDir()
		}
		return true
	}
	for _, path := range paths {
		expanded, derr := expandWildcard(c, path)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		for _, p := range expanded {
			if _, err := os.Stat(p); err == nil && typeMatch(p) {
				return []*object.PSObject{object.Bool(true)}, nil
			}
		}
	}
	return []*object.PSObject{object.Bool(false)}, nil
}

func cmdResolvePath(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	real, err := filepath.EvalSymlinks(full)
	if err == nil {
		full = real
	}
	return []*object.PSObject{object.Str(full)}, nil
}

func cmdConvertPath(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	return []*object.PSObject{object.Str(filepath.Clean(full))}, nil
}

func cmdSplitPath(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	if c.Args.Switch("Qualifier") {
		// 盘符限定符：本程序只有 C 盘，绝对路径（含 C: 前缀或 / 开头）返回 C:
		if strings.HasPrefix(path, "C:") || strings.HasPrefix(path, "c:") || strings.HasPrefix(path, "/") {
			return []*object.PSObject{object.Str("C:")}, nil
		}
		return []*object.PSObject{object.Str("")}, nil
	}
	if c.Args.Switch("Leaf") {
		return []*object.PSObject{object.Str(filepath.Base(full))}, nil
	}
	if c.Args.Switch("Parent") {
		return []*object.PSObject{object.Str(filepath.Dir(full))}, nil
	}
	// 无开关：输出父目录（PS 默认）
	return []*object.PSObject{object.Str(filepath.Dir(full))}, nil
}

func cmdJoinPath(c *Context) ([]*object.PSObject, error) {
	base, _ := c.Args.Str("Path")
	child, _ := c.Args.Str("ChildPath")
	if base == "" || child == "" {
		return nil, nil
	}
	return []*object.PSObject{object.Str(filepath.Join(base, child))}, nil
}

func cmdPushLocation(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	target := c.Shell.Cwd
	if path != "" {
		newPath, derr := resolvePath(c, path)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		if info, err := os.Stat(newPath); err != nil || !info.IsDir() {
			return errf(c, "%s", lang.T(lang.MsgPathNotFoundFmt, path))
		}
		target = filepath.Clean(newPath)
	}
	c.Shell.DirStack = append(c.Shell.DirStack, c.Shell.Cwd)
	c.Shell.Cwd = target
	return nil, nil
}

func cmdPopLocation(c *Context) ([]*object.PSObject, error) {
	if len(c.Shell.DirStack) == 0 {
		return nil, nil
	}
	last := c.Shell.DirStack[len(c.Shell.DirStack)-1]
	c.Shell.DirStack = c.Shell.DirStack[:len(c.Shell.DirStack)-1]
	c.Shell.Cwd = last
	return nil, nil
}

func cmdClearContent(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	paths, derr := expandWildcard(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	var wi whatIfCollector
	wi.cmdlet = "Clear-Content"
	wi.c = c
	var yesAll, noAll bool
	for _, p := range paths {
		// Clear-Content 只清空已存在的文件，不存在时报错而非创建空文件
		if _, err := os.Stat(p); err != nil {
			return errf(c, "%s", lang.T(lang.MsgPathNotFoundForSet, p))
		}
		if wi.hit(p) {
			continue
		}
		if confirmSkip(c, "Clear-Content", p, &yesAll, &noAll) {
			continue
		}
		if err := os.WriteFile(p, nil, 0o644); err != nil {
			return errf(c, "%s", lang.T(lang.MsgCannotClear, p))
		}
	}
	if out, ok := wi.result(); ok {
		return out, nil
	}
	return nil, nil
}

func cmdSetItem(c *Context) ([]*object.PSObject, error) {
	path, value := pathAndValue(c)
	if path == "" {
		return nil, nil
	}
	if value == nil {
		return nil, nil
	}
	// env: 驱动器
	if strings.HasPrefix(path, "env:") {
		os.Setenv(strings.TrimPrefix(path, "env:"), value.String())
		return nil, nil
	}
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	info, err := os.Stat(full)
	if err != nil {
		// Set-Item 只改写已存在的项，不存在时报错而非创建空文件
		return errf(c, "%s", lang.T(lang.MsgPathNotFoundForSet, path))
	}
	if info.IsDir() {
		return nil, nil
	}
	var wi whatIfCollector
	wi.cmdlet = "Set-Item"
	wi.c = c
	var yesAll, noAll bool
	if wi.hit(full) {
		out, _ := wi.result()
		return out, nil
	}
	if confirmSkip(c, "Set-Item", full, &yesAll, &noAll) {
		return nil, nil
	}
	if err := os.WriteFile(full, []byte(value.String()), 0o644); err != nil {
		return errf(c, "%s", lang.T(lang.MsgCannotWrite, path))
	}
	return nil, nil
}

func cmdClearItem(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	info, err := os.Stat(full)
	if err != nil {
		// Clear-Item 只清空已存在的项，不存在时报错而非创建空文件
		return errf(c, "%s", lang.T(lang.MsgPathNotFoundForSet, path))
	}
	if info.IsDir() {
		return nil, nil
	}
	var wi whatIfCollector
	wi.cmdlet = "Clear-Item"
	wi.c = c
	var yesAll, noAll bool
	if wi.hit(full) {
		out, _ := wi.result()
		return out, nil
	}
	if confirmSkip(c, "Clear-Item", full, &yesAll, &noAll) {
		return nil, nil
	}
	if err := os.WriteFile(full, nil, 0o644); err != nil {
		return errf(c, "%s", lang.T(lang.MsgCannotClear, path))
	}
	return nil, nil
}

func cmdGetItemProperty(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	info, err := os.Stat(full)
	if err != nil {
		return errf(c, "%s", lang.T(lang.MsgPathNotFoundFmt, path))
	}
	o := object.Object("System.Management.Automation.PSCustomObject", nil)
	o.AddProp("Name", info.Name())
	o.AddProp("FullName", full)
	o.AddProp("Length", info.Size())
	o.AddProp("LastWriteTime", info.ModTime())
	o.AddProp("Mode", object.UnixMode(info))
	// -Name：只保留指定属性（Windows 语义，如 Get-ItemProperty x -Name Length）
	nameFilter, _ := c.Args.Str("Name")
	if nameFilter != "" {
		var kept []object.Prop
		for _, p := range o.Props {
			if strings.EqualFold(p.Name, nameFilter) {
				kept = append(kept, p)
			}
		}
		if len(kept) == 0 {
			return errf(c, "%s", lang.T(lang.MsgPropNotFound, path, nameFilter))
		}
		o.Props = kept
	}
	return []*object.PSObject{o}, nil
}

func cmdSetItemProperty(c *Context) ([]*object.PSObject, error) {
	// 位置实参已由 Bind 中心化映射到 路径→属性名→值（跳过已命名的槽位）
	path, _ := c.Args.Str("Path")
	name, _ := c.Args.Str("Name")
	val := c.Args.Get("Value")
	if path == "" {
		return nil, nil
	}
	// MVP：支持 -Name LastWriteTime -Value <时间>；其余属性忽略
	if name != "" && val != nil && strings.EqualFold(name, "LastWriteTime") {
		full, derr := resolvePath(c, path)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		if err := os.Chtimes(full, time.Now(), time.Now()); err != nil {
			return errf(c, "%s", lang.T(lang.MsgPathNotFoundFmt, path))
		}
	}
	return nil, nil
}

// ---- 注册 ----

func init() {
	Register("Test-Path", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "PathType", Type: "string"},
	}, cmdTestPath)
	Register("Resolve-Path", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
	}, cmdResolvePath)
	Register("Convert-Path", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
	}, cmdConvertPath)
	Register("Split-Path", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Parent", Switch: true},
		{Name: "Leaf", Switch: true},
		{Name: "Qualifier", Switch: true},
	}, cmdSplitPath)
	Register("Join-Path", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "ChildPath", Position: 1, PositionSet: true, Type: "path"},
	}, cmdJoinPath)
	Register("Push-Location", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
	}, cmdPushLocation)
	Register("Pop-Location", nil, cmdPopLocation)
	Register("Clear-Content", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
	}, cmdClearContent)
	Register("Set-Item", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Value", Position: 1, PositionSet: true, Type: "object"},
	}, cmdSetItem)
	Register("Clear-Item", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
	}, cmdClearItem)
	Register("Get-ItemProperty", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Name", Position: 1, PositionSet: true, Type: "string"},
	}, cmdGetItemProperty)
	Register("Set-ItemProperty", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Name", Position: 1, PositionSet: true, Type: "string"},
		{Name: "Value", Position: 2, PositionSet: true, Type: "object"},
	}, cmdSetItemProperty)
}

package builtin

import (
	"os"
	"path/filepath"
	"strings"

	"powershell/internal/lang"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// navigation.go 实现导航定位类 cmdlet（目录列举、条目查看、位置切换、驱动器）。

func cmdGetChildItem(c *Context) ([]*object.PSObject, error) {
	// 路径：-Path（命名或位置，数组展开）加超量位置实参，全部当起始路径
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
		paths = []string{"."}
	}
	recurse := c.Args.Switch("Recurse")
	nameOnly := c.Args.Switch("Name")
	dirOnly := c.Args.Switch("Directory")
	fileOnly := c.Args.Switch("File")
	filter := ""
	if f, ok := c.Args.Str("Filter"); ok {
		filter = f
	}
	var out []*object.PSObject
	for _, path := range paths {
		expanded, derr := expandWildcard(c, path)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		isLiteral := !strings.ContainsAny(path, "*?[")
		for _, p := range expanded {
			info, err := os.Stat(p)
			if err != nil {
				// 字面缺失路径报错后继续其余路径（通配无匹配沿用静默；与 PowerShell 一致）。
				if isLiteral {
					full := p
					if resolved, rerr := resolvePath(c, path); rerr == nil {
						full = resolved
					}
					if _, stopErr := errf(c, "%s", lang.T(lang.MsgPathNotFoundForSet, full)); stopErr != nil {
						return nil, stopErr
					}
				}
				continue
			}
			out = append(out, listSinglePath(c, p, info, filter, nameOnly, recurse, dirOnly, fileOnly)...)
		}
	}
	return out, nil
}

func listSinglePath(c *Context, p string, info os.FileInfo, filter string, nameOnly bool, recurse bool, dirOnly bool, fileOnly bool) []*object.PSObject {
	var out []*object.PSObject
	if info.IsDir() {
		entries, err := os.ReadDir(p)
		if err != nil {
			return nil
		}
		for _, en := range entries {
			full := filepath.Join(p, en.Name())
			if filter != "" && !object.WildcardMatch(filter, en.Name()) {
				continue
			}
			fi, err := os.Stat(full)
			if err != nil {
				continue
			}
			if dirOnly && !fi.IsDir() {
				continue
			}
			if fileOnly && fi.IsDir() {
				continue
			}
			if nameOnly {
				out = append(out, object.Str(en.Name()))
			} else if fi.IsDir() {
				out = append(out, object.DirInfo(full, fi))
			} else {
				out = append(out, object.FileInfo(full, fi))
			}
			if recurse && fi.IsDir() {
				out = append(out, listSinglePath(c, full, fi, filter, nameOnly, recurse, dirOnly, fileOnly)...)
			}
		}
	} else {
		if dirOnly {
			return nil
		}
		if fileOnly && info.IsDir() {
			return nil
		}
		if filter != "" && !object.WildcardMatch(filter, filepath.Base(p)) {
			return nil
		}
		if nameOnly {
			out = append(out, object.Str(filepath.Base(p)))
		} else {
			out = append(out, object.FileInfo(p, info))
		}
	}
	return out
}

func cmdGetItem(c *Context) ([]*object.PSObject, error) {
	// 路径：-Path（命名或位置，数组展开）加超量位置实参
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
	var out []*object.PSObject
	for _, path := range paths {
		expanded, derr := expandWildcard(c, path)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		for _, p := range expanded {
			info, err := os.Stat(p)
			if err != nil {
				return errf(c, "%s", lang.T(lang.MsgPathNotFoundFmt, p))
			}
			if info.IsDir() {
				out = append(out, object.DirInfo(p, info))
			} else {
				out = append(out, object.FileInfo(p, info))
			}
		}
	}
	return out, nil
}

func cmdSetLocation(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		// 无参数时回主目录
		if h, err := os.UserHomeDir(); err == nil {
			c.Shell.Cwd = h
		}
		return nil, nil
	}
	if path == ".." {
		path = filepath.Join(c.Shell.Cwd, "..")
	}
	newPath, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	info, err := os.Stat(newPath)
	if err != nil || !info.IsDir() {
		return errf(c, "%s", lang.T(lang.MsgPathNotFoundForSet, path))
	}
	c.Shell.Cwd = filepath.Clean(newPath)
	return nil, nil
}

func cmdGetLocation(c *Context) ([]*object.PSObject, error) {
	return []*object.PSObject{object.Str(c.Shell.DisplayPath(c.Shell.Cwd))}, nil
}

func cmdGetPSDrive(c *Context) ([]*object.PSObject, error) {
	root := "/"
	driveName := "/"
	if c.Shell.Style == shell.StyleDesktop {
		driveName = "C"
	}
	d := object.Object("System.Management.Automation.PSDriveInfo", nil)
	d.AddProp("Name", driveName)
	d.AddProp("Root", root)
	d.AddProp("CurrentLocation", c.Shell.DisplayPath(c.Shell.Cwd))
	d.Table = []object.Column{
		{Label: "Name", Align: "left"},
		{Label: "Root", Align: "left"},
		{Label: "CurrentLocation", Align: "left"},
	}
	// 环境变量驱动器
	env := object.Object("PSDriveInfo", nil)
	env.AddProp("Name", "Env").AddProp("Root", "").AddProp("CurrentLocation", "")
	drives := []*object.PSObject{d, env}
	// -Name（命名或位置）：只列出指定名的驱动器，支持通配（如 Get-PSDrive -Name *）
	nameFilter, _ := c.Args.Str("Name")
	if nameFilter != "" {
		var out []*object.PSObject
		for _, dr := range drives {
			if v, ok2 := dr.PropValue("Name"); ok2 && object.WildcardMatchFold(nameFilter, v.String()) {
				out = append(out, dr)
			}
		}
		return out, nil
	}
	return drives, nil
}

// ---- 注册 ----

func init() {
	Register("Get-ChildItem", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Recurse", Switch: true},
		{Name: "Force", Switch: true},
		{Name: "Name", Switch: true},
		{Name: "Filter", Type: "string"},
		{Name: "Directory", Switch: true},
		{Name: "File", Switch: true},
	}, cmdGetChildItem)
	Register("Get-Item", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
	}, cmdGetItem)
	Register("Set-Location", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
	}, cmdSetLocation)
	Register("Get-Location", nil, cmdGetLocation)
	Register("Get-PSDrive", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdGetPSDrive)
}

package builtin

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"powershell/internal/object"
	"powershell/internal/shell"
)

// ---- 导航与文件 ----

func cmdGetChildItem(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		path = "."
	}
	recurse := c.Args.Switch("Recurse")
	nameOnly := c.Args.Switch("Name")
	dirOnly := c.Args.Switch("Directory")
	fileOnly := c.Args.Switch("File")
	filter := ""
	if f, ok := c.Args.Str("Filter"); ok {
		filter = f
	}
	paths, derr := expandWildcard(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	var out []*object.PSObject
	for _, p := range paths {
		info, err := os.Stat(p)
		if err != nil {
			continue
		}
		out = append(out, listOne(c, p, info, filter, nameOnly, recurse, dirOnly, fileOnly)...)
	}
	return out, nil
}

func listOne(c *Context, p string, info os.FileInfo, filter string, nameOnly bool, recurse bool, dirOnly bool, fileOnly bool) []*object.PSObject {
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
				out = append(out, listOne(c, full, fi, filter, nameOnly, recurse, dirOnly, fileOnly)...)
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
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	paths, derr := expandWildcard(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	var out []*object.PSObject
	for _, p := range paths {
		info, err := os.Stat(p)
		if err != nil {
			return errf(c, "Get-Item : 找不到路径 %s。", p)
		}
		if info.IsDir() {
			out = append(out, object.DirInfo(p, info))
		} else {
			out = append(out, object.FileInfo(p, info))
		}
	}
	return out, nil
}

func cmdSetLocation(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		// 无参数时回主目录（与 PowerShell 一致）
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
		return errf(c, "Set-Location : 找不到路径 %s，因为该路径不存在。", path)
	}
	c.Shell.Cwd = filepath.Clean(newPath)
	return nil, nil
}

func cmdGetLocation(c *Context) ([]*object.PSObject, error) {
	return []*object.PSObject{object.Str(c.Shell.DisplayPath(c.Shell.Cwd))}, nil
}

func cmdGetContent(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		if len(c.Input) > 0 {
			return c.Input, nil
		}
		return nil, nil
	}
	raw := c.Args.Switch("Raw")
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	data, err := os.ReadFile(full)
	if err != nil {
		return errf(c, "Get-Content : 找不到路径 %s。", path)
	}
	text := string(data)
	if raw {
		return []*object.PSObject{object.Str(text)}, nil
	}
	lines := strings.Split(strings.TrimSuffix(text, "\n"), "\n")
	if len(lines) == 1 && lines[0] == "" {
		lines = nil
	}
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
	items := make([]*object.PSObject, len(lines))
	for i, l := range lines {
		items[i] = object.Str(l)
	}
	return items, nil
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
	if err := os.WriteFile(full, []byte(sb.String()), 0o644); err != nil {
		return errf(c, "Set-Content : 无法写入 %s。", path)
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
	f, err := os.OpenFile(full, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return errf(c, "Add-Content : 无法打开 %s。", path)
	}
	defer f.Close()
	for _, o := range content {
		fmt.Fprintln(f, o.String())
	}
	return nil, nil
}

func cmdNewItem(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	itemType := "File"
	if t, ok := c.Args.Str("ItemType"); ok {
		itemType = t
	}
	force := c.Args.Switch("Force")
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	var err error
	if strings.EqualFold(itemType, "Directory") {
		if force {
			err = os.MkdirAll(full, 0o755)
		} else {
			err = os.Mkdir(full, 0o755)
		}
	} else {
		if force {
			if _, e := os.Stat(full); e != nil {
				err = os.WriteFile(full, nil, 0o644)
			}
		} else {
			f, e := os.OpenFile(full, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
			if e == nil {
				f.Close()
			} else {
				err = e
			}
		}
	}
	if err != nil {
		return errf(c, "New-Item : 无法创建 %s：%v", path, err)
	}
	if info, e := os.Stat(full); e == nil {
		if info.IsDir() {
			return []*object.PSObject{object.DirInfo(full, info)}, nil
		}
		return []*object.PSObject{object.FileInfo(full, info)}, nil
	}
	return nil, nil
}

func cmdRemoveItem(c *Context) ([]*object.PSObject, error) {
	recurse := c.Args.Switch("Recurse")
	force := c.Args.Switch("Force")
	for _, p := range pathList(c) {
		fulls, derr := expandWildcard(c, p)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		for _, full := range fulls {
			info, err := os.Stat(full)
			if err != nil {
				continue
			}
			if info.IsDir() {
				if recurse {
					err = os.RemoveAll(full)
				} else {
					err = os.Remove(full)
				}
			} else {
				if force {
					err = os.RemoveAll(full)
				} else {
					err = os.Remove(full)
				}
			}
			if err != nil && !force {
				return errf(c, "Remove-Item : 无法删除 %s：%v", p, err)
			}
		}
	}
	return nil, nil
}

func copyItem(c *Context, move bool) ([]*object.PSObject, error) {
	label := "Copy-Item"
	if move {
		label = "Move-Item"
	}
	recurse := c.Args.Switch("Recurse")
	// 源路径：命名 -Path（数组摊平）优先，其次管道输入，再次位置参数（末位为目标）
	var paths []string
	if v := c.Args.Get("Path"); v != nil {
		for _, it := range v.ArrayItems() {
			paths = append(paths, it.String())
		}
	} else if len(c.Input) > 0 {
		for _, o := range c.Input {
			paths = append(paths, o.String())
		}
	}
	// 目标：命名 -Destination 优先，其次位置参数末位
	dest := ""
	if d, ok := c.Args.Str("Destination"); ok {
		dest = d
	}
	posDest := ""
	if len(c.Args.Positional) > 0 && dest == "" {
		posDest = c.Args.Positional[len(c.Args.Positional)-1].String()
		dest = posDest
	}
	// 位置源：目标来自位置参数时排除末位，否则全部当源（数组摊平）
	if len(paths) == 0 {
		limit := len(c.Args.Positional)
		if posDest != "" {
			limit--
		}
		for i := 0; i < limit; i++ {
			for _, it := range c.Args.Positional[i].ArrayItems() {
				paths = append(paths, it.String())
			}
		}
	}
	if len(paths) == 0 {
		return nil, nil
	}
	// 展开全部源（多路径/通配），多源时目标必须是已存在目录（避免静默覆盖）
	var srcFulls []string
	for _, src := range paths {
		expanded, derr := expandWildcard(c, src)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		srcFulls = append(srcFulls, expanded...)
	}
	if len(srcFulls) > 1 {
		destFull, derr := resolvePath(c, dest)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		if info, err := os.Stat(destFull); err != nil || !info.IsDir() {
			return errf(c, "%s : 复制多个源时目标 %s 必须是已存在的目录。", label, dest)
		}
	}
	for _, srcFull := range srcFulls {
		info, err := os.Stat(srcFull)
		if err != nil {
			return errf(c, "%s : 找不到路径 %s。", label, srcFull)
		}
		destFull, derr := resolvePath(c, dest)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		if destInfo, err := os.Stat(destFull); err == nil && destInfo.IsDir() {
			destFull = filepath.Join(destFull, filepath.Base(srcFull))
		}
		if info.IsDir() && recurse {
			if err := copyDir(srcFull, destFull); err != nil {
				return errf(c, "%s : %v", label, err)
			}
		} else if info.IsDir() {
			return errf(c, "%s : 目录 %s 需要 -Recurse。", label, srcFull)
		} else {
			if err := copyFile(srcFull, destFull); err != nil {
				return errf(c, "%s : %v", label, err)
			}
		}
		if move {
			_ = os.RemoveAll(srcFull)
		}
	}
	return nil, nil
}

func copyFile(src, dest string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, data, 0o644)
}

func copyDir(src, dest string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dest, rel)
		if info.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		return copyFile(path, target)
	})
}

func cmdCopyItem(c *Context) ([]*object.PSObject, error) { return copyItem(c, false) }
func cmdMoveItem(c *Context) ([]*object.PSObject, error) { return copyItem(c, true) }

func cmdRenameItem(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	newName := ""
	// -NewName：只取叶子名，支持传完整路径；先过盘符校验（D:\ 报错，C:\ 归一到根再取叶子）
	if n, ok := c.Args.Str("NewName"); ok {
		np, derr := shell.DrivePath(n)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		newName = filepath.Base(np)
	}
	if len(c.Args.Positional) >= 2 {
		path = c.Args.Positional[0].String()
		np, derr := shell.DrivePath(c.Args.Positional[1].String())
		if derr != nil {
			return errf(c, "%v", derr)
		}
		newName = filepath.Base(np)
	}
	if path == "" || newName == "" {
		return nil, nil
	}
	old, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	newPath := filepath.Join(filepath.Dir(old), newName)
	if err := os.Rename(old, newPath); err != nil {
		return errf(c, "Rename-Item : 无法重命名 %s：%v", path, err)
	}
	return nil, nil
}

func cmdClearHost(c *Context) ([]*object.PSObject, error) {
	fmt.Fprint(c.Stdout, "\x1b[2J\x1b[H")
	return nil, nil
}

func cmdInvokeItem(c *Context) ([]*object.PSObject, error) {
	// 用系统默认方式打开（MVP：打印路径）
	path := firstPathArg(c)
	if path != "" {
		fmt.Fprintln(c.Stdout, path)
	}
	return nil, nil
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
	nameFilter := ""
	if n, ok := c.Args.Str("Name"); ok {
		nameFilter = n
	} else if p := c.Args.Pos(0); p != nil {
		nameFilter = p.String()
	}
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
		{Name: "Path", Position: 0, Type: "path"},
		{Name: "Recurse", Switch: true},
		{Name: "Force", Switch: true},
		{Name: "Name", Switch: true},
		{Name: "Filter", Type: "string"},
		{Name: "Directory", Switch: true},
		{Name: "File", Switch: true},
	}, cmdGetChildItem)
	Register("Get-Item", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
	}, cmdGetItem)
	Register("Set-Location", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
	}, cmdSetLocation)
	Register("Get-Location", nil, cmdGetLocation)
	Register("Get-Content", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
		{Name: "Raw", Switch: true},
		{Name: "TotalCount", Type: "int"},
		{Name: "Tail", Type: "int"},
	}, cmdGetContent)
	Register("Set-Content", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
		{Name: "Value", Position: 1, Type: "object"},
		{Name: "Encoding", Type: "string"},
	}, cmdSetContent)
	Register("Add-Content", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
		{Name: "Value", Position: 1, Type: "object"},
	}, cmdAddContent)
	Register("New-Item", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
		{Name: "ItemType", Type: "string"},
		{Name: "Force", Switch: true},
	}, cmdNewItem)
	Register("Remove-Item", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
		{Name: "Recurse", Switch: true},
		{Name: "Force", Switch: true},
	}, cmdRemoveItem)
	Register("Copy-Item", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
		{Name: "Destination", Position: 1, Type: "path"},
		{Name: "Recurse", Switch: true},
	}, cmdCopyItem)
	Register("Move-Item", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
		{Name: "Destination", Position: 1, Type: "path"},
		{Name: "Recurse", Switch: true},
		{Name: "Force", Switch: true},
	}, cmdMoveItem)
	Register("Rename-Item", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
		{Name: "NewName", Position: 1, Type: "string"},
	}, cmdRenameItem)
	Register("Clear-Host", nil, cmdClearHost)
	Register("Invoke-Item", []ParamSpec{
		{Name: "Path", Position: 0, Type: "path"},
	}, cmdInvokeItem)
	Register("Get-PSDrive", []ParamSpec{
		{Name: "Name", Position: 0, Type: "string"},
	}, cmdGetPSDrive)
}

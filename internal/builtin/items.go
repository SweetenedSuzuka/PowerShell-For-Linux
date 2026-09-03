package builtin

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"powershell/internal/lang"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// items.go 实现条目管理类 cmdlet（新建、删除、复制、移动、重命名、打开）。

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
	var wi whatIfCollector
	wi.cmdlet = "New-Item"
	wi.c = c
	var yesAll, noAll bool
	if wi.hit(full) {
		out, _ := wi.result()
		return out, nil
	}
	if confirmSkip(c, "New-Item", full, &yesAll, &noAll) {
		return nil, nil
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
		return errf(c, "%s", lang.T(lang.MsgCannotCreate, path, err))
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
	var wi whatIfCollector
	wi.cmdlet = "Remove-Item"
	wi.c = c
	var yesAll, noAll bool
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
			if wi.hit(full) {
				continue
			}
			if confirmSkip(c, "Remove-Item", full, &yesAll, &noAll) {
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
				return errf(c, "%s", lang.T(lang.MsgCannotDelete, p, err))
			}
		}
	}
	if out, ok := wi.result(); ok {
		return out, nil
	}
	return nil, nil
}

func copyItem(c *Context, move bool) ([]*object.PSObject, error) {
	label := "Copy-Item"
	if move {
		label = "Move-Item"
	}
	var wi whatIfCollector
	wi.cmdlet = label
	wi.c = c
	var yesAll, noAll bool
	recurse := c.Args.Switch("Recurse")
	// 源路径：命名/位置 -Path（数组摊平）优先，其次管道输入
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
	// 目标：-Destination（命名或位置）
	dest, _ := c.Args.Str("Destination")
	// 多源写法 Copy-Item a b c 复制 a、b 到 c，超出槽位的实参在此处处理：
	// 目标由位置映射时末位实参提升为目标，Path 与映射到 Destination 的值都并入源；
	// 目标显式命名时剩余实参全部并入源（如 Copy-Item a b -Destination d）。
	if len(c.Args.Positional) > 0 {
		rest := c.Args.Positional
		if c.Args.PosMapped["Destination"] {
			dest = rest[len(rest)-1].String()
			rest = rest[:len(rest)-1]
			if v := c.Args.Get("Destination"); v != nil {
				for _, it := range v.ArrayItems() {
					paths = append(paths, it.String())
				}
			}
		}
		for _, p := range rest {
			for _, it := range p.ArrayItems() {
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
			return errf(c, "%s", lang.T(lang.MsgCopyDestNotDir, dest))
		}
	}
	for _, srcFull := range srcFulls {
		info, err := os.Stat(srcFull)
		if err != nil {
			return errf(c, "%s : %s", label, lang.T(lang.MsgPathNotFoundFmt, srcFull))
		}
		if wi.hit(srcFull) {
			continue
		}
		if confirmSkip(c, label, srcFull, &yesAll, &noAll) {
			continue
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
			return errf(c, "%s : %s", label, lang.T(lang.MsgCopyNeedsRecurse, srcFull))
		} else {
			if err := copyFile(srcFull, destFull); err != nil {
				return errf(c, "%s : %v", label, err)
			}
		}
		if move {
			_ = os.RemoveAll(srcFull)
		}
	}
	if out, ok := wi.result(); ok {
		return out, nil
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
	// -NewName（命名或位置）：只取叶子名，支持传完整路径；先过盘符校验（D:\ 报错，C:\ 归一到根再取叶子）
	if n, ok := c.Args.Str("NewName"); ok {
		np, derr := shell.DrivePath(n)
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
	// 目标已存在时拒绝重命名（os.Rename 会静默覆盖，PowerShell 会抛出错误）
	if _, err := os.Lstat(newPath); err == nil {
		return errf(c, "%s", lang.T(lang.MsgRenameDestExists, newPath))
	}
	var wi whatIfCollector
	wi.cmdlet = "Rename-Item"
	wi.c = c
	var yesAll, noAll bool
	if wi.hit(old) {
		out, _ := wi.result()
		return out, nil
	}
	if confirmSkip(c, "Rename-Item", old, &yesAll, &noAll) {
		return nil, nil
	}
	if err := os.Rename(old, newPath); err != nil {
		return errf(c, "%s", lang.T(lang.MsgCannotRename, path, err))
	}
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

// ---- 注册 ----

func init() {
	Register("New-Item", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "ItemType", Type: "string"},
		{Name: "Force", Switch: true},
	}, cmdNewItem)
	Register("Remove-Item", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Recurse", Switch: true},
		{Name: "Force", Switch: true},
	}, cmdRemoveItem)
	Register("Copy-Item", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Destination", Position: 1, PositionSet: true, Type: "path"},
		{Name: "Recurse", Switch: true},
	}, cmdCopyItem)
	Register("Move-Item", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Destination", Position: 1, PositionSet: true, Type: "path"},
		{Name: "Recurse", Switch: true},
		{Name: "Force", Switch: true},
	}, cmdMoveItem)
	Register("Rename-Item", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "NewName", Position: 1, PositionSet: true, Type: "string"},
	}, cmdRenameItem)
	Register("Invoke-Item", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
	}, cmdInvokeItem)
}

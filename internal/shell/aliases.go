package shell

import (
	"sort"
	"strings"
)

// aliases.go 管理由风格决定的别名表。

// ResolveAlias 解析别名（不区分大小写）。
func (s *Session) ResolveAlias(name string) (string, bool) {
	if target, ok := s.Aliases[name]; ok {
		return target, true
	}
	// 别名表按小写存储，用户输入可能混大小写
	for k, v := range s.Aliases {
		if strings.EqualFold(k, name) {
			return v, true
		}
	}
	return "", false
}

// SetAlias 设置/覆盖别名（Set-Alias 用）。
func (s *Session) SetAlias(name, target string) {
	s.Aliases[name] = target
}

// AllAliasNames 返回全部别名名。
func (s *Session) AllAliasNames() []string {
	var names []string
	for n := range s.Aliases {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

// buildAliases 依风格构建别名表（键小写）。
func buildAliases(style Style) map[string]string {
	m := map[string]string{
		// 导航与文件
		"ls": "Get-ChildItem", "dir": "Get-ChildItem", "gci": "Get-ChildItem",
		"cd": "Set-Location", "sl": "Set-Location", "chdir": "Set-Location",
		"pwd": "Get-Location", "gl": "Get-Location",
		"cat": "Get-Content", "type": "Get-Content", "gc": "Get-Content",
		"echo": "Write-Output", "write": "Write-Output",
		"cls": "Clear-Host", "clear": "Clear-Host",
		"ps": "Get-Process", "gps": "Get-Process",
		"rm": "Remove-Item", "del": "Remove-Item", "erase": "Remove-Item",
		"ri": "Remove-Item", "rd": "Remove-Item", "rmdir": "Remove-Item",
		"cp": "Copy-Item", "copy": "Copy-Item",
		"mv": "Move-Item", "move": "Move-Item", "mi": "Move-Item",
		"ren": "Rename-Item", "rni": "Rename-Item",
		"ni": "New-Item", "gi": "Get-Item", "ii": "Invoke-Item", "cli": "Clear-Item",
		// 对象操作
		"?": "Where-Object", "where": "Where-Object",
		"%":       "ForEach-Object",
		"sort":    "Sort-Object",
		"select":  "Select-Object",
		"measure": "Measure-Object",
		"group":   "Group-Object",
		"fl":      "Format-List", "ft": "Format-Table", "fw": "Format-Wide",
		// 信息
		"gcm": "Get-Command", "gal": "Get-Alias", "gv": "Get-Variable",
		"help": "Get-Help", "man": "Get-Help",
		"date":        "Get-Date",
		"gdr":         "Get-PSDrive",
		"history":     "Get-History",
		"sls":         "Select-String",
		"switchstyle": "Set-PSVersion", // 切换风格别名
		// 更多常用别名
		"gp": "Get-ItemProperty", "sp": "Set-ItemProperty", "si": "Set-Item",
		"rv": "Remove-Variable", "clv": "Clear-Variable", "nv": "New-Variable", "sv": "Set-Variable",
		"na": "New-Alias", "sa": "Set-Alias",
		"gm": "Get-Member", "gu": "Get-Unique",
		"gsv": "Get-Service", "sasv": "Start-Service", "spsv": "Stop-Service",
		"iex": "Invoke-Expression", "iwr": "Invoke-WebRequest", "irm": "Invoke-RestMethod",
		"ihy": "Invoke-History", "ghy": "Get-History",
		"sleep": "Start-Sleep", "tee": "Tee-Object",
		"cpi": "Copy-Item", "gh": "Get-Help",
	}
	if style == StyleDesktop {
		// 5.X 独有：curl/wget 映射到 Invoke-WebRequest，sc 映射到 Set-Content
		m["curl"] = "Invoke-WebRequest"
		m["wget"] = "Invoke-WebRequest"
		m["sc"] = "Set-Content"
	}
	return m
}

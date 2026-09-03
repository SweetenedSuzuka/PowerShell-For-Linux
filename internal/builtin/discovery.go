package builtin

import (
	"fmt"
	"sort"
	"strings"

	"powershell/internal/lang"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// discovery.go 实现面向会话自身的 cmdlet（命令发现、帮助、版本切换）。

func cmdGetHelp(c *Context) ([]*object.PSObject, error) {
	name, _ := c.Args.Str("Name")
	if name == "" {
		fmt.Fprintln(c.Stdout, c.Shell.Usage())
		return nil, nil
	}
	// 支持通配符匹配（大小写不敏感）
	var matched []string
	for _, n := range AllCmdletNames() {
		display := canonicalName(n)
		if object.WildcardMatchFold(name, display) {
			matched = append(matched, n)
		}
	}
	if len(matched) == 0 {
		if alias, ok := c.Shell.ResolveAlias(name); ok {
			fmt.Fprintf(c.Stdout, "%s\n\n", lang.T(lang.MsgHelpAliasTo, name, canonicalName(alias)))
			return nil, nil
		}
		fmt.Fprintf(c.Stdout, "%s\n", lang.T(lang.MsgHelpNotFound, name))
		return nil, nil
	}
	for _, n := range matched {
		display := canonicalName(n)
		fmt.Fprintf(c.Stdout, "%s\n    %s\n\n", lang.T(lang.MsgHelpNameHeader), display)
		fmt.Fprintf(c.Stdout, "%s\n    %s", lang.T(lang.MsgHelpSyntaxHeader), display)
		for _, sp := range Spec(n) {
			fmt.Fprintf(c.Stdout, " [-%s", sp.Name)
			if !sp.Switch {
				fmt.Fprintf(c.Stdout, " <%s>", sp.Type)
			}
			fmt.Fprint(c.Stdout, "]")
		}
		fmt.Fprintln(c.Stdout)
		fmt.Fprintln(c.Stdout)
		fmt.Fprintf(c.Stdout, "%s\n    %s\n\n", lang.T(lang.MsgHelpAliasHeader), aliasesOf(c.Shell, display))
	}
	return nil, nil
}

func aliasesOf(s *shell.Session, cmdlet string) string {
	var out []string
	for _, n := range s.AllAliasNames() {
		if strings.EqualFold(s.Aliases[n], cmdlet) {
			out = append(out, n)
		}
	}
	return strings.Join(out, ", ")
}

// canonicalName 把注册表的小写 cmdlet 名转回原始大小写（get-childitem → Get-ChildItem）。
func canonicalName(name string) string {
	if orig, ok := displayMap[name]; ok {
		return orig
	}
	parts := strings.Split(name, "-")
	for i, p := range parts {
		if p == "" {
			continue
		}
		parts[i] = strings.ToUpper(p[:1]) + p[1:]
	}
	return strings.Join(parts, "-")
}

func cmdGetCommand(c *Context) ([]*object.PSObject, error) {
	names := c.Args.StringSlice("Name")
	var out []*object.PSObject
	match := func(n string) bool {
		if len(names) == 0 {
			return true
		}
		for _, pat := range names {
			if strings.EqualFold(pat, n) {
				return true
			}
		}
		return false
	}
	add := func(n, typ string) {
		display := canonicalName(n)
		if !match(display) && !match(n) {
			return
		}
		o := object.Object("System.Management.Automation.CommandInfo", nil)
		o.AddProp("Name", display)
		o.AddProp("CommandType", typ)
		out = append(out, o)
	}
	for n := range registry {
		add(n, "Cmdlet")
	}
	for n := range c.Shell.Aliases {
		if !match(n) {
			continue
		}
		o := object.Object("AliasInfo", nil)
		o.AddProp("Name", n)
		o.AddProp("CommandType", "Alias")
		o.AddProp("Definition", c.Shell.Aliases[n])
		o.Table = []object.Column{
			{Label: "CommandType", Align: "left"},
			{Label: "Name", Align: "left"},
		}
		out = append(out, o)
	}
	for n := range c.Shell.Functions {
		add(n, "Function")
	}
	for _, o := range out {
		if o.TypeName == "System.Management.Automation.CommandInfo" {
			o.Table = []object.Column{
				{Label: "CommandType", Align: "left"},
				{Label: "Name", Align: "left"},
			}
		}
	}
	return out, nil
}

// cmdSetPSVersion 切换 5.X/7.X 风格。
func cmdSetPSVersion(c *Context) ([]*object.PSObject, error) {
	ver, _ := c.Args.Str("Version")
	switch {
	case strings.HasPrefix(ver, "5"):
		c.Shell.SetStyle(shell.StyleDesktop)
	case strings.HasPrefix(ver, "7"):
		c.Shell.SetStyle(shell.StyleCore)
	default:
		return errf(c, "%s", lang.T(lang.MsgPSVersionBad, ver))
	}
	fmt.Fprintf(c.Stdout, "%s\n", lang.T(lang.MsgPSVersionSet, c.Shell.StyleName()))
	return nil, nil
}

// AllCmdletNames 列出全部内置 cmdlet 名（Get-Command / 补全用）。
func AllCmdletNames() []string {
	var names []string
	for n := range registry {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

// ---- 注册 ----

func init() {
	Register("Get-Help", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdGetHelp)
	Register("Get-Command", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdGetCommand)
	Register("Set-PSVersion", []ParamSpec{
		{Name: "Version", Position: 0, PositionSet: true, Type: "string"},
	}, cmdSetPSVersion)
}

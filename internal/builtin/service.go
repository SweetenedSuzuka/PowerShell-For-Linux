package builtin

import (
	"os"
	"os/exec"
	"strings"

	"powershell/internal/lang"
	"powershell/internal/object"
)

// service.go 实现服务类 cmdlet（systemd 映射）。

// runSystemctl 执行 systemctl；普通权限失败时自动用 sudo 重试。
func runSystemctl(c *Context, action, unit string) int {
	code := runExternalRaw(c, "systemctl", []string{action, unit})
	if code != 0 && os.Geteuid() != 0 {
		return runExternalRaw(c, "sudo", []string{"systemctl", action, unit})
	}
	return code
}

func cmdGetService(c *Context) ([]*object.PSObject, error) {
	if _, err := exec.LookPath("systemctl"); err != nil {
		return errf(c, "%s", lang.T(lang.MsgServiceNeedSystemd))
	}
	out, err := exec.Command("systemctl", "list-units", "--type=service", "--all", "--no-pager").Output()
	if err != nil {
		// 尝试 sudo
		out, err = exec.Command("sudo", "systemctl", "list-units", "--type=service", "--all", "--no-pager").Output()
		if err != nil {
			return errf(c, "%s", lang.T(lang.MsgServiceListFail))
		}
	}
	var result []*object.PSObject
	// -Name（命名或位置）：按服务名过滤（支持通配，Windows 语义）
	nameFilter, _ := c.Args.Str("Name")
	for _, ln := range strings.Split(string(out), "\n") {
		fields := strings.Fields(ln)
		if len(fields) < 3 || !strings.HasSuffix(fields[0], ".service") {
			continue
		}
		name := strings.TrimSuffix(fields[0], ".service")
		if nameFilter != "" && !object.WildcardMatch(nameFilter, name) {
			continue
		}
		status := fields[2]
		o := object.Object("System.ServiceProcess.ServiceController", nil)
		o.AddProp("Name", name)
		o.AddProp("Status", status)
		o.AddProp("DisplayName", name)
		o.Table = []object.Column{
			{Label: "Status", Align: "left"},
			{Label: "Name", Align: "left"},
			{Label: "DisplayName", Align: "left"},
		}
		result = append(result, o)
	}
	return result, nil
}

// serviceAction 生成对 systemd 服务执行某动作的 cmdlet 实现。
func serviceAction(action string) CmdFunc {
	return func(c *Context) ([]*object.PSObject, error) {
		name := firstArg(c, "Name")
		if name == "" {
			return nil, nil
		}
		unit := name
		if !strings.HasSuffix(unit, ".service") {
			unit += ".service"
		}
		runSystemctl(c, action, unit)
		return nil, nil
	}
}

func cmdSetService(c *Context) ([]*object.PSObject, error) {
	name := firstArg(c, "Name")
	if name == "" {
		return nil, nil
	}
	unit := name
	if !strings.HasSuffix(unit, ".service") {
		unit += ".service"
	}
	if status, ok := c.Args.Str("Status"); ok {
		switch strings.ToLower(status) {
		case "running", "started":
			runSystemctl(c, "start", unit)
		case "stopped":
			runSystemctl(c, "stop", unit)
		}
	}
	if startup, ok := c.Args.Str("StartupType"); ok {
		switch strings.ToLower(startup) {
		case "automatic", "auto":
			runSystemctl(c, "enable", unit)
		case "disabled":
			runSystemctl(c, "disable", unit)
		}
	}
	return nil, nil
}

// ---- 注册 ----

func init() {
	Register("Get-Service", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdGetService)
	Register("Start-Service", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, serviceAction("start"))
	Register("Stop-Service", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, serviceAction("stop"))
	Register("Restart-Service", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, serviceAction("restart"))
	Register("Resume-Service", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, serviceAction("start"))
	Register("Set-Service", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Status", Type: "string"},
		{Name: "StartupType", Type: "string"},
	}, cmdSetService)
}

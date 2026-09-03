package builtin

import (
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"powershell/internal/object"
)

// computer.go 实现本机类 cmdlet（机器信息、时区、区域、剪贴板、电源）。

func cmdGetComputerInfo(c *Context) ([]*object.PSObject, error) {
	o := object.Object("Microsoft.PowerShell.Commands.ComputerInfo", nil)
	hostname, _ := os.Hostname()
	o.AddProp("CsName", hostname)
	o.AddProp("OsName", readOSRelease("NAME"))
	o.AddProp("OsVersion", readOSRelease("VERSION_ID"))
	o.AddProp("OsArchitecture", runtime.GOARCH)
	o.AddProp("OsPlatform", runtime.GOOS)
	if data, err := os.ReadFile("/proc/meminfo"); err == nil {
		for _, ln := range strings.Split(string(data), "\n") {
			if strings.HasPrefix(ln, "MemTotal:") {
				fields := strings.Fields(ln)
				if len(fields) >= 2 {
					if kb, err := strconv.ParseInt(fields[1], 10, 64); err == nil {
						o.AddProp("CsTotalPhysicalMemory", kb*1024)
					}
				}
				break
			}
		}
	}
	o.AddProp("CsProcessors", runtime.NumCPU())
	return []*object.PSObject{o}, nil
}

func readOSRelease(key string) string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return ""
	}
	for _, ln := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(ln, key+"=") {
			return strings.Trim(strings.TrimPrefix(ln, key+"="), `"`)
		}
	}
	return ""
}

func cmdGetTimeZone(c *Context) ([]*object.PSObject, error) {
	id := "UTC"
	if data, err := os.ReadFile("/etc/timezone"); err == nil {
		id = strings.TrimSpace(string(data))
	}
	o := object.Object("System.TimeZoneInfo", nil)
	o.AddProp("Id", id)
	o.AddProp("DisplayName", id)
	return []*object.PSObject{o}, nil
}

func cmdSetTimeZone(c *Context) ([]*object.PSObject, error) {
	zone := firstArg(c, "Name")
	if zone == "" {
		return nil, nil
	}
	runExternalRaw(c, "sudo", []string{"timedatectl", "set-timezone", zone})
	return nil, nil
}

func cmdGetUptime(c *Context) ([]*object.PSObject, error) {
	var secs float64
	if runtime.GOOS == "linux" {
		if data, err := os.ReadFile("/proc/uptime"); err == nil {
			fields := strings.Fields(string(data))
			if len(fields) > 0 {
				secs, _ = strconv.ParseFloat(fields[0], 64)
			}
		}
	}
	d := time.Duration(secs * float64(time.Second))
	o := timeSpanObj(d)
	o.Table = []object.Column{
		{Label: "Days", Align: "right"},
		{Label: "Hours", Align: "right"},
		{Label: "Minutes", Align: "right"},
		{Label: "Seconds", Align: "right"},
	}
	return []*object.PSObject{o}, nil
}

func cmdGetCulture(c *Context) ([]*object.PSObject, error) {
	// 区域按界面语言的登记表返回，未登记的语言回退 zh-CN（与 $Host.CurrentCulture 一致）
	return []*object.PSObject{c.Shell.CultureObject()}, nil
}

func cmdGetClipboard(c *Context) ([]*object.PSObject, error) {
	if _, err := exec.LookPath("xclip"); err == nil {
		out, _ := exec.Command("xclip", "-o", "-selection", "clipboard").Output()
		return []*object.PSObject{object.Str(strings.TrimRight(string(out), "\n"))}, nil
	}
	if _, err := exec.LookPath("xsel"); err == nil {
		out, _ := exec.Command("xsel", "-b").Output()
		return []*object.PSObject{object.Str(strings.TrimRight(string(out), "\n"))}, nil
	}
	return []*object.PSObject{object.Str("")}, nil
}

func cmdSetClipboard(c *Context) ([]*object.PSObject, error) {
	var text string
	if len(c.Input) > 0 {
		for _, o := range c.Input {
			text += o.String() + "\n"
		}
	} else if v := c.Args.Get("Value"); v != nil {
		text = v.String()
	}
	text = strings.TrimRight(text, "\n")
	if _, err := exec.LookPath("xclip"); err == nil {
		cmd := exec.Command("xclip", "-selection", "clipboard")
		cmd.Stdin = strings.NewReader(text)
		_ = cmd.Run()
	} else if _, err := exec.LookPath("xsel"); err == nil {
		cmd := exec.Command("xsel", "-b")
		cmd.Stdin = strings.NewReader(text)
		_ = cmd.Run()
	}
	return nil, nil
}

func cmdRestartComputer(c *Context) ([]*object.PSObject, error) {
	runExternalRaw(c, "sudo", []string{"reboot"})
	return nil, nil
}

func cmdStopComputer(c *Context) ([]*object.PSObject, error) {
	runExternalRaw(c, "sudo", []string{"shutdown", "-h", "now"})
	return nil, nil
}

func cmdRenameComputer(c *Context) ([]*object.PSObject, error) {
	name, _ := c.Args.Str("NewName")
	if name != "" {
		runExternalRaw(c, "sudo", []string{"hostnamectl", "set-hostname", name})
	}
	return nil, nil
}

// ---- 注册 ----

func init() {
	Register("Get-ComputerInfo", nil, cmdGetComputerInfo)
	Register("Get-Uptime", nil, cmdGetUptime)
	Register("Get-TimeZone", nil, cmdGetTimeZone)
	Register("Set-TimeZone", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdSetTimeZone)
	Register("Get-Culture", nil, cmdGetCulture)
	Register("Get-Clipboard", nil, cmdGetClipboard)
	Register("Set-Clipboard", []ParamSpec{
		{Name: "Value", Position: 0, PositionSet: true, Type: "object"},
	}, cmdSetClipboard)
	Register("Restart-Computer", nil, cmdRestartComputer)
	Register("Stop-Computer", nil, cmdStopComputer)
	Register("Rename-Computer", []ParamSpec{
		{Name: "NewName", Position: 0, PositionSet: true, Type: "string"},
	}, cmdRenameComputer)
}

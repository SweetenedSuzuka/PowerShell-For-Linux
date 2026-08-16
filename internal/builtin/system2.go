package builtin

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"powershell/internal/object"
	"powershell/internal/shell"
)

// runExternalRaw 直接运行外部命令并透传 IO（服务/连接类 cmdlet 用）。
func runExternalRaw(c *Context, program string, args []string) int {
	cmd := exec.Command(program, args...)
	cmd.Stdin = c.Stdin
	cmd.Stdout = c.Stdout
	cmd.Stderr = c.Stderr
	if err := cmd.Run(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return ee.ExitCode()
		}
		return 127
	}
	return 0
}

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
		return errf(c, "Get-Service : 需要 systemd（systemctl）。")
	}
	out, err := exec.Command("systemctl", "list-units", "--type=service", "--all", "--no-pager").Output()
	if err != nil {
		// 尝试 sudo
		out, err = exec.Command("sudo", "systemctl", "list-units", "--type=service", "--all", "--no-pager").Output()
		if err != nil {
			return errf(c, "Get-Service : 无法读取服务列表。")
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

func cmdGetCulture(c *Context) ([]*object.PSObject, error) {
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = "en-US"
	}
	o := object.Object("System.Globalization.CultureInfo", nil)
	o.AddProp("Name", lang)
	o.AddProp("DisplayName", lang)
	return []*object.PSObject{o}, nil
}

func cmdTestConnection(c *Context) ([]*object.PSObject, error) {
	host := firstArg(c, "TargetName")
	if host == "" {
		return nil, nil
	}
	count := "4"
	if n, ok := c.Args.Int("Count"); ok && n > 0 {
		count = strconv.FormatInt(n, 10)
	}
	if _, err := exec.LookPath("ping"); err != nil {
		return errf(c, "Test-Connection : 找不到 ping 命令。")
	}
	code := runExternalRaw(c, "ping", []string{"-c", count, host})
	return []*object.PSObject{object.Bool(code == 0)}, nil
}

func cmdStartSleep(c *Context) ([]*object.PSObject, error) {
	var d time.Duration
	if sec, ok := c.Args.Int("Seconds"); ok {
		d = time.Duration(sec) * time.Second
	}
	if ms, ok := c.Args.Int("Milliseconds"); ok {
		d = time.Duration(ms) * time.Millisecond
	}
	if d > 0 {
		time.Sleep(d)
	}
	return nil, nil
}

func cmdStopProcess(c *Context) ([]*object.PSObject, error) {
	if id, ok := c.Args.Int("Id"); ok {
		_ = killProcess(int(id))
		return nil, nil
	}
	name := firstArg(c, "Name")
	if name == "" {
		if len(c.Input) > 0 {
			for _, o := range c.Input {
				if pid, ok := o.PropValue("Id"); ok {
					if p, ok2 := pid.AsInt(); ok2 {
						_ = killProcess(int(p))
					}
				}
			}
			return nil, nil
		}
		return nil, nil
	}
	if pid, err := strconv.Atoi(name); err == nil {
		_ = killProcess(pid)
		return nil, nil
	}
	// 按进程名
	if procs, err := listProcesses(); err == nil {
		for _, pr := range procs {
			if strings.EqualFold(pr.name, name) {
				_ = killProcess(pr.pid)
			}
		}
	}
	return nil, nil
}

func killProcess(pid int) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return p.Kill()
}

func cmdStartProcess(c *Context) ([]*object.PSObject, error) {
	path := firstArg(c, "FilePath")
	args := c.Args.StringSlice("ArgumentList")
	if len(args) == 0 && len(c.Args.Positional) > 0 {
		// 位置参数：首个（路径未命名时）当 FilePath，其余当 ArgumentList
		start := 0
		if path == "" {
			path = c.Args.Positional[0].String()
			start = 1
		}
		for i := start; i < len(c.Args.Positional); i++ {
			args = append(args, c.Args.Positional[i].String())
		}
	}
	if path == "" {
		return nil, nil
	}
	np, derr := shell.DrivePath(path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	path = np
	cmd := exec.Command(path, args...)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Start(); err != nil {
		return errf(c, "Start-Process : %v", err)
	}
	return []*object.PSObject{object.Process(cmd.Process.Pid, filepath.Base(path), 0, 0)}, nil
}

func cmdWaitProcess(c *Context) ([]*object.PSObject, error) {
	name := firstArg(c, "Name")
	if id, ok := c.Args.Int("Id"); ok {
		for processAlive(int(id)) {
			time.Sleep(100 * time.Millisecond)
		}
		return nil, nil
	}
	if name == "" {
		return nil, nil
	}
	if pid, err := strconv.Atoi(name); err == nil {
		for processAlive(pid) {
			time.Sleep(100 * time.Millisecond)
		}
		return nil, nil
	}
	for {
		found := false
		if procs, e := listProcesses(); e == nil {
			for _, pr := range procs {
				if strings.EqualFold(pr.name, name) {
					found = true
					break
				}
			}
		}
		if !found {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil, nil
}

// processAlive 判断进程是否存活（Linux 用 /proc，其它平台尽力而为）。
func processAlive(pid int) bool {
	if runtime.GOOS == "linux" {
		_, err := os.Stat(filepath.Join("/proc", strconv.Itoa(pid)))
		return err == nil
	}
	p, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	return p.Signal(syscall.Signal(0)) == nil
}

func cmdSetDate(c *Context) ([]*object.PSObject, error) {
	date := firstArg(c, "Date")
	if date == "" {
		return []*object.PSObject{object.DateTime(time.Now())}, nil
	}
	runExternalRaw(c, "sudo", []string{"date", "-s", date})
	return []*object.PSObject{object.DateTime(time.Now())}, nil
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
	Register("Get-ComputerInfo", nil, cmdGetComputerInfo)
	Register("Get-TimeZone", nil, cmdGetTimeZone)
	Register("Set-TimeZone", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdSetTimeZone)
	Register("Get-Culture", nil, cmdGetCulture)
	Register("Test-Connection", []ParamSpec{
		{Name: "TargetName", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Count", Type: "int"},
	}, cmdTestConnection)
	Register("Start-Sleep", []ParamSpec{
		{Name: "Seconds", Position: 0, PositionSet: true, Type: "int"},
		{Name: "Milliseconds", Type: "int"},
	}, cmdStartSleep)
	Register("Stop-Process", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Id", Type: "int"},
	}, cmdStopProcess)
	Register("Start-Process", []ParamSpec{
		{Name: "FilePath", Position: 0, PositionSet: true, Type: "path"},
		{Name: "ArgumentList", Type: "string[]"},
	}, cmdStartProcess)
	Register("Wait-Process", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Id", Type: "int"},
	}, cmdWaitProcess)
	Register("Set-Date", []ParamSpec{
		{Name: "Date", Position: 0, PositionSet: true, Type: "string"},
	}, cmdSetDate)
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

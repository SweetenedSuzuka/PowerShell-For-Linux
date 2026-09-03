package builtin

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"powershell/internal/lang"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// process.go 实现进程类 cmdlet。

func cmdGetProcess(c *Context) ([]*object.PSObject, error) {
	name, _ := c.Args.Str("Name")
	var ids []int
	if v := c.Args.Get("Id"); v != nil {
		// -Id 支持单个 PID、逗号分隔串与 PID 数组
		for _, it := range v.ArrayItems() {
			for _, f := range strings.Split(it.String(), ",") {
				if n, err := strconv.Atoi(strings.TrimSpace(f)); err == nil {
					ids = append(ids, n)
				}
			}
		}
	}
	procs, err := listProcesses()
	if err != nil {
		return []*object.PSObject{object.Process(os.Getpid(), "powershell", 0, 0)}, nil
	}
	var out []*object.PSObject
	for _, pr := range procs {
		if len(ids) > 0 {
			// -Id 按 PID 精确匹配；不存在的 PID 报非终止错误
			for _, id := range ids {
				if pr.pid == id {
					out = append(out, object.Process(pr.pid, pr.name, pr.cpu, pr.mem))
				}
			}
			continue
		}
		if name != "" && !strings.Contains(strings.ToLower(pr.name), strings.ToLower(name)) {
			continue
		}
		out = append(out, object.Process(pr.pid, pr.name, pr.cpu, pr.mem))
	}
	if len(ids) > 0 && len(out) < len(ids) {
		c.Shell.LastSuccess = false
		fmt.Fprintf(c.Stderr, "%s : %s\n", c.Shell.StyleName(), lang.T(lang.MsgProcIdNotFound, c.Args.Get("Id").String()))
	}
	// 无筛选却无结果时回退当前进程；带 -Id/-Name 查询无结果则报错并返回空
	if len(out) == 0 {
		if name != "" {
			c.Shell.LastSuccess = false
			fmt.Fprintf(c.Stderr, "%s : %s\n", c.Shell.StyleName(), lang.T(lang.MsgProcNameNotFound, name))
		}
		if name == "" && len(ids) == 0 {
			return []*object.PSObject{object.Process(os.Getpid(), "powershell", 0, 0)}, nil
		}
	}
	return out, nil
}

type procEntry struct {
	pid  int
	name string
	cpu  float64
	mem  int64
}

func listProcesses() ([]procEntry, error) {
	if runtime.GOOS == "linux" {
		return listLinuxProcs()
	}
	return nil, fmt.Errorf("%s", lang.T(lang.MsgUnsupported))
}

// clockTicks 是内核每秒的时钟滴答数（/proc/stat 时间字段单位），Linux 用户态 ABI 固定 100。
const clockTicks = 100

func listLinuxProcs() ([]procEntry, error) {
	entries, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}
	pageSize := int64(os.Getpagesize())
	var out []procEntry
	for _, en := range entries {
		if !en.IsDir() {
			continue
		}
		pid, err := strconv.Atoi(en.Name())
		if err != nil {
			continue
		}
		name := ""
		if data, err := os.ReadFile("/proc/" + en.Name() + "/comm"); err == nil {
			name = strings.TrimSpace(string(data))
		}
		cpu, mem := 0.0, int64(0)
		if data, err := os.ReadFile("/proc/" + en.Name() + "/stat"); err == nil {
			// stat 第 2 字段（comm）可含空格与括号，从最后一个 ')' 之后解析数字字段
			s := string(data)
			if i := strings.LastIndexByte(s, ')'); i >= 0 && i+2 <= len(s) {
				fields := strings.Fields(s[i+2:])
				// fields 从 state（原第 3 字段）起：utime 是原第 14 字段 → 下标 11；stime → 12；rss → 21
				if len(fields) > 22 {
					utime, _ := strconv.ParseFloat(fields[11], 64)
					stime, _ := strconv.ParseFloat(fields[12], 64)
					rss, _ := strconv.ParseInt(fields[21], 10, 64)
					cpu = (utime + stime) / clockTicks
					mem = rss * pageSize
				}
			}
		}
		out = append(out, procEntry{pid: pid, name: name, cpu: cpu, mem: mem})
	}
	return out, nil
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

// ---- 注册 ----

func init() {
	Register("Get-Process", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Id", Type: "string"},
	}, cmdGetProcess)
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
}

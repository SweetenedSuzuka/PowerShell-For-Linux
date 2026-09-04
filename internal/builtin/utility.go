package builtin

import (
	"os/exec"
	"strconv"
	"time"

	"powershell/internal/lang"
	"powershell/internal/object"
)

// utility.go 收纳不成体系的独立小指令（连接测试、暂停）。

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
		return errf(c, "%s", lang.T(lang.MsgPingNotFound))
	}
	code := runExternalRaw(c, "ping", []string{"-c", count, host})
	return []*object.PSObject{object.Bool(code == 0)}, nil
}

func cmdStartSleep(c *Context) ([]*object.PSObject, error) {
	// -Seconds 与 -Milliseconds 分属不同参数集，不可同用（按出现与否判定，与取值无关）。
	if c.Args.Get("Seconds") != nil && c.Args.Get("Milliseconds") != nil {
		return errf(c, "%s", lang.T(lang.MsgParamSetUnresolvable))
	}
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

// ---- 注册 ----

func init() {
	Register("Test-Connection", []ParamSpec{
		{Name: "TargetName", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Count", Type: "int"},
	}, cmdTestConnection)
	Register("Start-Sleep", []ParamSpec{
		{Name: "Seconds", Position: 0, PositionSet: true, Type: "int"},
		{Name: "Milliseconds", Type: "int"},
	}, cmdStartSleep)
}

package builtin

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"

	"powershell/internal/lang"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// ---- 系统与信息 ----

func cmdGetDate(c *Context) ([]*object.PSObject, error) {
	now := time.Now()
	// -Date 指定日期时间（本地时区，无时区信息时按本地解析）
	if d, ok := c.Args.Str("Date"); ok && d != "" {
		if t, err := parseDateArg(d); err == nil {
			now = t
		}
	}
	if f, ok := c.Args.Str("Format"); ok && f != "" {
		return []*object.PSObject{object.Str(now.Format(dotnetToGoLayout(f)))}, nil
	}
	return []*object.PSObject{object.DateTime(now)}, nil
}

// parseDateArg 按常见日期时间格式解析 -Date 参数（无时区信息按本地时区）。
func parseDateArg(s string) (time.Time, error) {
	layouts := []string{
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05",
		"2006-01-02",
		"2006/1/2",
		time.RFC3339,
	}
	for _, l := range layouts {
		if t, err := time.ParseInLocation(l, s, time.Local); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("%s", lang.T(lang.MsgDateParseFail, s))
}

// dotnetToGoLayout 把 .NET 日期格式串转换为 Go 布局（yyyy→2006、MM→01、dd→02、HH→15 等）。
func dotnetToGoLayout(f string) string {
	var sb strings.Builder
	for i := 0; i < len(f); i++ {
		c := f[i]
		if c == '\\' && i+1 < len(f) { // 转义
			sb.WriteByte(f[i+1])
			i++
			continue
		}
		// 收集同字符连续段
		j := i
		for j < len(f) && f[j] == c {
			j++
		}
		n := j - i
		switch c {
		case 'y':
			if n >= 4 {
				sb.WriteString("2006")
			} else {
				sb.WriteString("06")
			}
		case 'M':
			switch n {
			case 1:
				sb.WriteString("1")
			case 2:
				sb.WriteString("01")
			case 3:
				sb.WriteString("Jan")
			default:
				sb.WriteString("January")
			}
		case 'd':
			switch n {
			case 1:
				sb.WriteString("2")
			case 2:
				sb.WriteString("02")
			case 3:
				sb.WriteString("Mon")
			default:
				sb.WriteString("Monday")
			}
		case 'H':
			if n >= 2 {
				sb.WriteString("15")
			} else {
				sb.WriteString("15")
			}
		case 'h':
			if n >= 2 {
				sb.WriteString("03")
			} else {
				sb.WriteString("3")
			}
		case 'm':
			if n >= 2 {
				sb.WriteString("04")
			} else {
				sb.WriteString("4")
			}
		case 's':
			if n >= 2 {
				sb.WriteString("05")
			} else {
				sb.WriteString("5")
			}
		case 't':
			sb.WriteString("PM")
		case 'z':
			if n >= 3 {
				sb.WriteString("-07:00")
			} else {
				sb.WriteString("-07")
			}
		default:
			for k := 0; k < n; k++ {
				sb.WriteByte(c)
			}
		}
		i = j - 1
	}
	return sb.String()
}

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

func cmdGetAlias(c *Context) ([]*object.PSObject, error) {
	names := c.Args.StringSlice("Name")
	match := func(n string) bool {
		if len(names) == 0 {
			return true
		}
		for _, pat := range names {
			if object.WildcardMatch(pat, n) {
				return true
			}
		}
		return false
	}
	var out []*object.PSObject
	for _, n := range c.Shell.AllAliasNames() {
		if !match(n) {
			continue
		}
		o := object.Object("AliasInfo", nil)
		o.AddProp("Name", n)
		o.AddProp("Definition", c.Shell.Aliases[n])
		o.Table = []object.Column{
			{Label: "Name", Align: "left"},
			{Label: "Definition", Align: "left"},
		}
		out = append(out, o)
	}
	return out, nil
}

func cmdSetAlias(c *Context) ([]*object.PSObject, error) {
	name, _ := c.Args.Str("Name")
	value, _ := c.Args.Str("Value")
	if name != "" && value != "" {
		c.Shell.SetAlias(name, value)
	}
	return nil, nil
}

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

func cmdGetVariable(c *Context) ([]*object.PSObject, error) {
	names := c.Args.StringSlice("Name")
	match := func(n string) bool {
		if len(names) == 0 {
			return true
		}
		for _, pat := range names {
			if object.WildcardMatch(pat, n) {
				return true
			}
		}
		return false
	}
	var out []*object.PSObject
	for _, n := range c.Shell.AllVarNames() {
		if !match(n) {
			continue
		}
		if v, ok := c.Shell.GetVar(n); ok {
			o := object.Object("System.Management.Automation.PSVariable", nil)
			o.AddProp("Name", n)
			o.AddProp("Value", v)
			out = append(out, o)
		}
	}
	return out, nil
}

func cmdSetVariable(c *Context) ([]*object.PSObject, error) {
	name, val := pairArgs(c, "Name")
	if name == "" {
		return nil, nil
	}
	if val == nil && len(c.Input) > 0 {
		val = object.Array(c.Input)
	}
	if val != nil {
		if err := c.Shell.SetVar(name, val); err != nil {
			return errf(c, "Set-Variable : %v", err)
		}
	}
	return nil, nil
}

func cmdGetHistory(c *Context) ([]*object.PSObject, error) {
	var out []*object.PSObject
	for i, h := range c.Shell.History {
		o := object.Object("HistoryInfo", nil)
		o.AddProp("Id", int64(i+1))
		o.AddProp("CommandLine", h)
		o.Table = []object.Column{
			{Label: "Id", Align: "right"},
			{Label: "CommandLine", Align: "left"},
		}
		out = append(out, o)
	}
	return out, nil
}

func cmdClearHistory(c *Context) ([]*object.PSObject, error) {
	c.Shell.History = nil
	return nil, nil
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

func cmdInvokeWebRequest(c *Context) ([]*object.PSObject, error) {
	uri, _ := c.Args.Str("Uri")
	if uri == "" {
		return nil, nil
	}
	body, method := httpRequestArgs(c)
	resp, err := doHTTPRequest(method, uri, body)
	if err != nil {
		return errf(c, "Invoke-WebRequest : %v", err)
	}
	text := strings.TrimSuffix(resp, "\n")
	var lines []*object.PSObject
	if text != "" {
		for _, ln := range strings.Split(text, "\n") {
			lines = append(lines, object.Str(ln))
		}
	}
	return lines, nil
}

func cmdInvokeRestMethod(c *Context) ([]*object.PSObject, error) {
	uri, _ := c.Args.Str("Uri")
	if uri == "" {
		return nil, nil
	}
	body, method := httpRequestArgs(c)
	resp, err := doHTTPRequest(method, uri, body)
	if err != nil {
		return errf(c, "Invoke-RestMethod : %v", err)
	}
	// 尝试解析 JSON
	var v any
	if err := json.Unmarshal([]byte(resp), &v); err == nil {
		return []*object.PSObject{jsonToObject(v)}, nil
	}
	return []*object.PSObject{object.Str(strings.TrimSuffix(resp, "\n"))}, nil
}

// httpRequestArgs 取请求的 HTTP 方法与请求体（默认 GET）。
func httpRequestArgs(c *Context) (string, string) {
	method := "GET"
	if m, ok := c.Args.Str("Method"); ok && m != "" {
		method = m
	}
	body := ""
	if b := c.Args.Get("Body"); b != nil {
		body = b.String()
	}
	return body, method
}

// doHTTPRequest 发一次 HTTP 请求，返回响应文本；30 秒超时。
func doHTTPRequest(method, uri, body string) (string, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return "", err
	}
	if body != "" && (method == "POST" || method == "PUT" || method == "PATCH") {
		req.Body = io.NopCloser(strings.NewReader(body))
		req.ContentLength = int64(len(body))
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	return string(data), err
}

func cmdGetFileHash(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	algorithm := "SHA256"
	if a, ok := c.Args.Str("Algorithm"); ok {
		algorithm = a
	}
	var out []*object.PSObject
	paths, derr := expandWildcard(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	for _, p := range paths {
		data, err := os.ReadFile(p)
		if err != nil {
			continue
		}
		hash, err := computeHash(algorithm, data)
		if err != nil {
			return errf(c, "%s", lang.T(lang.MsgHashAlgoUnsupported, algorithm))
		}
		o := object.Object("Microsoft.PowerShell.Utility.FileHash", nil)
		o.AddProp("Algorithm", strings.ToUpper(algorithm))
		o.AddProp("Hash", strings.ToLower(hash))
		o.AddProp("Path", p)
		out = append(out, o)
	}
	return out, nil
}

func computeHash(algorithm string, data []byte) (string, error) {
	switch strings.ToUpper(algorithm) {
	case "SHA256", "SHA2_256":
		sum := sha256.Sum256(data)
		return fmt.Sprintf("%x", sum), nil
	case "SHA1":
		sum := sha1.Sum(data)
		return fmt.Sprintf("%x", sum), nil
	case "MD5":
		sum := md5.Sum(data)
		return fmt.Sprintf("%x", sum), nil
	case "SHA512", "SHA2_512":
		sum := sha512.Sum512(data)
		return fmt.Sprintf("%x", sum), nil
	}
	return "", fmt.Errorf("%s", lang.T(lang.MsgUnsupported))
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
	Register("Get-Date", []ParamSpec{
		{Name: "Date", Type: "string"},
		{Name: "Format", Type: "string"},
	}, cmdGetDate)
	Register("Get-Help", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdGetHelp)
	Register("Get-Command", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdGetCommand)
	Register("Get-Alias", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdGetAlias)
	Register("Set-Alias", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Value", Position: 1, PositionSet: true, Type: "string"},
	}, cmdSetAlias)
	Register("Get-Process", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Id", Type: "string"},
	}, cmdGetProcess)
	Register("Get-Uptime", nil, cmdGetUptime)
	Register("Get-Variable", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
	}, cmdGetVariable)
	Register("Set-Variable", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Value", Position: 1, PositionSet: true, Type: "object"},
	}, cmdSetVariable)
	Register("Get-History", nil, cmdGetHistory)
	Register("Clear-History", nil, cmdClearHistory)
	Register("Set-PSVersion", []ParamSpec{
		{Name: "Version", Position: 0, PositionSet: true, Type: "string"},
	}, cmdSetPSVersion)
	Register("Invoke-WebRequest", []ParamSpec{
		{Name: "Uri", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Method", Type: "string"},
		{Name: "Body", Type: "object"},
	}, cmdInvokeWebRequest)
	Register("Invoke-RestMethod", []ParamSpec{
		{Name: "Uri", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Method", Type: "string"},
		{Name: "Body", Type: "object"},
	}, cmdInvokeRestMethod)
	Register("Get-FileHash", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Algorithm", Type: "string"},
	}, cmdGetFileHash)
}

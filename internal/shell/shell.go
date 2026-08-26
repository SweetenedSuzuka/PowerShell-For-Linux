// Package shell 管理解释器会话状态：变量、别名、函数、当前目录、5.X/7.X 样式、$PSVersionTable、历史等。
package shell

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"powershell/internal/ast"
	"powershell/internal/lang"
	"powershell/internal/object"
)

// Style 表示两种 PowerShell 风格。
type Style int

const (
	StyleCore    Style = iota // PowerShell 7.X（Core）
	StyleDesktop              // Windows PowerShell 5.X（Desktop）
)

func (s Style) String() string {
	if s == StyleDesktop {
		return "Windows PowerShell"
	}
	return "PowerShell 7"
}

// Lang 是界面语言，直接沿用 lang 包的类型。
type Lang = lang.Lang

const (
	LangZh = lang.LangZh
	LangEn = lang.LangEn
)

// Function 是用户定义函数。
type Function struct {
	Name    string
	Params  []ast.FunctionParam
	Body    *ast.Block
	Filter  bool
	Begin   *ast.Block // begin 命名块，nil 为未声明
	Process *ast.Block // process 命名块，nil 为未声明
	End     *ast.Block // end 命名块，nil 为未声明
}

// Session 是一次解释器会话的全部状态。
type Session struct {
	Style         Style
	Lang          Lang                        // 界面语言（lang 包的语言码）
	Vars          map[string]*object.PSObject // 用户显式变量
	Aliases       map[string]string
	Functions     map[string]*Function
	Cwd           string
	DirStack      []string // Push-Location / Pop-Location 的目录栈
	History       []string
	HistoryFile   string // 历史文件路径（空则不持久化）
	LastExit      int
	LastSuccess   bool               // $?
	ErrorRecords  []*object.PSObject // $Error：本会话累积的错误记录，最新在前
	Matches       *object.PSObject   // $Matches：最近一次标量 -match 的捕获组，未匹配过为 nil
	PSCommandPath string
	Args          []*object.PSObject // 脚本/函数实参（$args）
	HostOut       io.Writer
	HostErr       io.Writer
	HostIn        io.Reader
}

// New 创建新会话。
func New(style Style, stdout, stderr io.Writer, stdin io.Reader) *Session {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = "."
	}
	s := &Session{
		Style:     style,
		Lang:      lang.Detect(),
		Vars:      map[string]*object.PSObject{},
		Functions: map[string]*Function{},
		Cwd:       cwd,
		HostOut:   stdout,
		HostErr:   stderr,
		HostIn:    stdin,
	}
	s.Aliases = buildAliases(style)
	// 会话语言即全局界面语言，提示文本与日期渲染都按它取
	lang.SetCurrent(s.Lang)
	return s
}

// SetStyle 切换 5.X/7.X 风格，联动别名与自动变量。
func (s *Session) SetStyle(style Style) {
	s.Style = style
	s.Aliases = buildAliases(style)
}

// StyleName 返回风格名。
func (s *Session) StyleName() string {
	if s.Style == StyleDesktop {
		return "Windows PowerShell"
	}
	return "PowerShell 7"
}

// ---- 变量 ----

// IsReadOnlyVar 报告变量是否为只读自动变量。
func IsReadOnlyVar(name string) bool {
	switch name {
	case "PSVersionTable", "PID", "PWD", "HOME", "Host", "IsLinux", "IsWindows", "IsMacOS",
		"PSCommandPath", "PSEdition", "PSHOME", "Matches":
		return true
	}
	return false
}

// maxErrorRecords 是 $Error 的容量上限（对应 PowerShell 的 $MaximumErrorCount 默认值）。
const maxErrorRecords = 256

// ErrorViewMarker 标记一个数组对象是 $Error 的动态视图（eval 层的 Clear/RemoveAt 据此落到记录本体）。
const ErrorViewMarker = "__ErrorView"

// RecordError 构造一条错误记录并累积进会话，最新在前；超出容量时丢弃最旧的。
func (s *Session) RecordError(msg string) *object.PSObject {
	rec := object.Error(msg)
	s.ErrorRecords = append([]*object.PSObject{rec}, s.ErrorRecords...)
	if len(s.ErrorRecords) > maxErrorRecords {
		s.ErrorRecords = s.ErrorRecords[:maxErrorRecords]
	}
	return rec
}

// ClearErrorRecords 清空已累积的错误记录。
func (s *Session) ClearErrorRecords() {
	s.ErrorRecords = nil
}

// RemoveErrorRecord 删除指定下标的错误记录；下标越界返回 false。
func (s *Session) RemoveErrorRecord(idx int64) bool {
	if idx < 0 || int(idx) >= len(s.ErrorRecords) {
		return false
	}
	s.ErrorRecords = append(s.ErrorRecords[:idx], s.ErrorRecords[idx+1:]...)
	return true
}

// SetVar 设置变量；只读自动变量拒绝修改。
func (s *Session) SetVar(name string, val *object.PSObject) error {
	if IsReadOnlyVar(name) {
		return fmt.Errorf("%s", lang.T(lang.MsgReadonlyVar, name))
	}
	s.Vars[name] = val
	return nil
}

// GetVar 读取变量：先查显式变量，再查自动变量。
func (s *Session) GetVar(name string) (*object.PSObject, bool) {
	if v, ok := s.Vars[name]; ok {
		return v, true
	}
	switch name {
	case "PWD":
		return object.Str(s.DisplayPath(s.Cwd)), true
	case "HOME":
		if h, err := os.UserHomeDir(); err == nil {
			return object.Str(h), true
		}
		return object.Null(), true
	case "PID":
		return object.Int(int64(os.Getpid())), true
	case "PSVersionTable":
		return s.VersionTable(), true
	case "LASTEXITCODE":
		return object.Int(int64(s.LastExit)), true
	case "?":
		return object.Bool(s.LastSuccess), true
	case "Matches":
		if s.Matches == nil {
			return object.Null(), true
		}
		return s.Matches, true
	case "Error":
		arr := object.Array(s.ErrorRecords)
		// 标记为 $Error 的动态视图：Clear/RemoveAt 等方法经 Session 落到 ErrorRecords 本体
		arr.AddProp(ErrorViewMarker, object.Str("1"))
		return arr, true
	case "PSCommandPath":
		if s.PSCommandPath != "" {
			return object.Str(s.PSCommandPath), true
		}
		return object.Null(), true
	case "args":
		if len(s.Args) > 0 {
			return object.Array(s.Args), true
		}
		return object.Array(nil), true
	case "Host":
		return s.HostObject(), true
	case "PSEdition":
		if s.Style == StyleDesktop {
			return object.Str("Desktop"), true
		}
		return object.Str("Core"), true
	case "IsLinux":
		if s.Style == StyleDesktop {
			return nil, false // 5.X 不存在此变量
		}
		return object.Bool(runtime.GOOS == "linux"), true
	case "IsWindows":
		if s.Style == StyleDesktop {
			return nil, false
		}
		return object.Bool(runtime.GOOS == "windows"), true
	case "IsMacOS":
		if s.Style == StyleDesktop {
			return nil, false
		}
		return object.Bool(runtime.GOOS == "darwin"), true
	case "IsCoreCLR":
		return object.Bool(s.Style != StyleDesktop), true
	case "PSHOME":
		if h, err := os.UserHomeDir(); err == nil {
			return object.Str(filepath.Join(h, ".local", "share", "powershell")), true
		}
		return object.Null(), true
	case "OFS":
		return object.Str(" "), true
	}
	return nil, false
}

// psVersion* 是 $PSVersionTable 里宣称的版本号，只标到 -Version 标志的粒度：
// 5.X 对应 Windows PowerShell 5.1，7.X 对应 PowerShell 7（Minor 为 0，渲染时省略）。
// 不宣称具体小版本；要改版本号，只改这里的常量，渲染与 $PSVersionTable 都会跟着变。
const (
	psVersionMajorDesktop = 5
	psVersionMinorDesktop = 1
	psVersionMajorCore    = 7
	psVersionMinorCore    = 0
)

// VersionTable 构造 $PSVersionTable 哈希表。
// 风格区分靠 PSEdition（Desktop/Core）。
func (s *Session) VersionTable() *object.PSObject {
	var entries []object.HashEntry
	add := func(k string, v any) {
		entries = append(entries, object.HashEntry{Key: k, Value: object.ToPS(v)})
	}
	if s.Style == StyleDesktop {
		add("PSVersion", object.Version(psVersionMajorDesktop, psVersionMinorDesktop, 0, 0))
		add("PSEdition", "Desktop")
	} else {
		add("PSVersion", object.Version(psVersionMajorCore, psVersionMinorCore, 0, 0))
		add("PSEdition", "Core")
		add("GitCommitId", "0000000000000000000000000000000000000000")
		add("OS", s.OSName())
		add("Platform", "Unix")
	}
	return object.VersionTable(entries)
}

func (s *Session) OSName() string {
	switch runtime.GOOS {
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	case "darwin":
		return "Darwin"
	}
	return runtime.GOOS
}

// cultureObject 构造 CultureInfo 对象（LCID/Name/DisplayName），区域取值由调用方按界面语言选定。
func cultureObject(name string, lcid int64, display string) *object.PSObject {
	c := object.Object("System.Globalization.CultureInfo", name)
	c.AddProp("LCID", lcid)
	c.AddProp("Name", name)
	c.AddProp("DisplayName", display)
	return c
}

// CultureObject 返回当前界面语言对应的 CultureInfo，Get-Culture 与 $Host.CurrentCulture 共用这一份。
// 各语言的区域数据在此登记：语言码 → 名称、LCID、显示名；未登记的语言回退默认语言的 zh-CN。
func (s *Session) CultureObject() *object.PSObject {
	switch s.Lang {
	case LangZh:
		return cultureObject("zh-CN", 2052, "中文（中国）")
	case LangEn:
		return cultureObject("en-US", 1033, "English (United States)")
	}
	return cultureObject("zh-CN", 2052, "中文（中国）")
}

// newUUID 生成随机 UUID v4。
func newUUID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return ""
	}
	b[6] = (b[6] & 0x0F) | 0x40
	b[8] = (b[8] & 0x3F) | 0x80
	return fmt.Sprintf("%02x%02x%02x%02x-%02x%02x-%02x%02x-%02x%02x-%02x%02x%02x%02x%02x%02x",
		b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15])
}

// HostObject 构造 $Host 对象：InstanceId 为本会话的随机 UUID，区域随界面语言（未登记的语言回退默认语言）。
func (s *Session) HostObject() *object.PSObject {
	h := object.Object("System.Management.Automation.Internal.Host.InternalHost", nil)
	h.AddProp("Name", "ConsoleHost")
	h.AddProp("InstanceId", newUUID())
	ui := object.Object("System.Management.Automation.Internal.Host.InternalHostUserInterface", nil)
	ui.AddProp("SupportsVirtualTerminal", true)
	h.AddProp("UI", ui)
	culture := s.CultureObject()
	h.AddProp("CurrentCulture", culture)
	h.AddProp("CurrentUICulture", culture)
	return h
}

// AllVarNames 列出全部可见变量名（用于 Get-Variable / 补全）。
func (s *Session) AllVarNames() []string {
	set := map[string]bool{}
	for n := range s.Vars {
		set[n] = true
	}
	for _, n := range []string{"PWD", "HOME", "PID", "PSVersionTable", "LASTEXITCODE", "?", "Matches", "PSCommandPath", "args", "Host", "PSEdition", "IsLinux", "IsWindows", "IsMacOS", "PSHOME", "OFS"} {
		set[n] = true
	}
	var names []string
	for n := range set {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

// ---- 别名与命令解析 ----

// ResolveAlias 解析别名（区分大小写不敏感）。
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

// ---- 路径显示 ----

// DrivePath 处理输入路径的盘符前缀：C: 开头视为根目录（C 盘 = 根），其它盘符报错。
// Linux 没有盘符概念，C: 只是给人看的系统盘表示；Windows 写法 C:\tmp、C:/tmp 都归一化到根路径。
func DrivePath(p string) (string, error) {
	if len(p) < 2 || p[1] != ':' {
		return p, nil
	}
	d := p[0]
	if !(d >= 'a' && d <= 'z') && !(d >= 'A' && d <= 'Z') {
		return p, nil
	}
	if !strings.EqualFold(p[:1], "C") {
		return "", fmt.Errorf("%s", lang.T(lang.MsgDriveUnsupported, strings.ToUpper(p[:1])))
	}
	rest := strings.TrimLeft(p[2:], `/\`)
	rest = strings.ReplaceAll(rest, `\`, `/`)
	if rest == "" {
		return "/", nil
	}
	return "/" + rest, nil
}

// ResolvePath 把输入路径解析为绝对路径：盘符归一化（C: 当根）、~ 展开、相对路径基于 cwd。
func ResolvePath(cwd, p string) (string, error) {
	if p == "" {
		return p, nil
	}
	if np, err := DrivePath(p); err != nil {
		return "", err
	} else if np != p {
		return np, nil
	}
	if p == "~" {
		if h, err := os.UserHomeDir(); err == nil {
			return h, nil
		}
		return p, nil
	}
	if strings.HasPrefix(p, "~/") {
		if h, err := os.UserHomeDir(); err == nil {
			return filepath.Join(h, p[2:]), nil
		}
	}
	if !filepath.IsAbs(p) {
		return filepath.Join(cwd, p), nil
	}
	return filepath.Clean(p), nil
}

// DisplayPath 把 Linux 路径显示为 Windows 风格（C:\，根目录即 C:\）。
// 这是本项目的身份标识：无论命令格式是 7 还是 5，提示符都显示 Windows 风格路径。
func (s *Session) DisplayPath(path string) string {
	if runtime.GOOS == "windows" {
		return path
	}
	return "C:" + strings.ReplaceAll(path, "/", "\\")
}

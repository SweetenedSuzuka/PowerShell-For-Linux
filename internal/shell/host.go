package shell

import (
	"crypto/rand"
	"fmt"
	"runtime"

	"powershell/internal/object"
)

// host.go 构造会话呈现对象（版本表、区域、主机）。

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

// HostObject 构造 $Host 对象：InstanceId 每次构造随机生成，区域随界面语言（未登记的语言回退默认语言）。
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

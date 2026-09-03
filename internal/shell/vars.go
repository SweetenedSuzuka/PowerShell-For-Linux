package shell

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"powershell/internal/lang"
	"powershell/internal/object"
)

// vars.go 实现会话变量读写（显式变量与自动变量）。

// IsReadOnlyVar 报告变量是否为只读自动变量，不区分大小写。
func IsReadOnlyVar(name string) bool {
	switch strings.ToLower(name) {
	case "psversiontable", "pid", "pwd", "home", "host", "islinux", "iswindows", "ismacos",
		"pscommandpath", "psedition", "pshome", "matches":
		return true
	}
	return false
}

// varKey 取显式变量的存储键：已存在（不区分大小写）沿用原大小写，否则用传入名。
func (s *Session) varKey(name string) string {
	if _, ok := s.Vars[name]; ok {
		return name
	}
	for k := range s.Vars {
		if strings.EqualFold(k, name) {
			return k
		}
	}
	return name
}

// HasVar 报告显式变量是否存在，不区分大小写。
func (s *Session) HasVar(name string) bool {
	_, ok := s.Vars[s.varKey(name)]
	return ok
}

// DeleteVar 删除显式变量，不区分大小写；没有返回 false。
func (s *Session) DeleteVar(name string) bool {
	k := s.varKey(name)
	if _, ok := s.Vars[k]; !ok {
		return false
	}
	delete(s.Vars, k)
	return true
}

// maxErrorRecords 是 $Error 的容量上限（对应 PowerShell 的 $MaximumErrorCount 默认值）。
const maxErrorRecords = 256

// ErrorViewMarker 标记一个数组对象是 $Error 的动态视图（eval 层的 Clear/RemoveAt 据此落到记录本体）。
const ErrorViewMarker = "__ErrorView"

// ParseErrorAction 把错误动作取值归一化为小写，未知取值返回 false。
func ParseErrorAction(s string) (string, bool) {
	switch strings.ToLower(s) {
	case "continue", "silentlycontinue", "stop", "inquire", "ignore":
		return strings.ToLower(s), true
	}
	return "", false
}

// RecordError 构造一条错误记录并累积进会话，最新在前；超出容量时丢弃最旧的。
func (s *Session) RecordError(msg string) *object.PSObject {
	rec := object.Error(msg)
	s.ErrorSeq++
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

// SetVar 设置变量；只读自动变量拒绝修改，$ErrorActionPreference 只接受有效取值。
func (s *Session) SetVar(name string, val *object.PSObject) error {
	if IsReadOnlyVar(name) {
		return fmt.Errorf("%s", lang.T(lang.MsgReadonlyVar, name))
	}
	if strings.EqualFold(name, "ErrorActionPreference") {
		// 首选项名按规范大小写存储。
		// 空值视为恢复默认。
		name = "ErrorActionPreference"
		if val.IsNull() {
			delete(s.Vars, name)
			return nil
		}
		if _, ok := ParseErrorAction(val.String()); !ok {
			return fmt.Errorf("%s", lang.T(lang.MsgErrorActionPreferenceInvalid, val.String()))
		}
	}
	s.Vars[s.varKey(name)] = val
	return nil
}

// GetVar 读取变量：先查显式变量，再查自动变量；两处都不区分大小写。
func (s *Session) GetVar(name string) (*object.PSObject, bool) {
	if v, ok := s.Vars[s.varKey(name)]; ok {
		return v, true
	}
	switch strings.ToLower(name) {
	case "pwd":
		return object.Str(s.DisplayPath(s.Cwd)), true
	case "home":
		if h, err := os.UserHomeDir(); err == nil {
			return object.Str(h), true
		}
		return object.Null(), true
	case "pid":
		return object.Int(int64(os.Getpid())), true
	case "psversiontable":
		return s.VersionTable(), true
	case "lastexitcode":
		return object.Int(int64(s.LastExit)), true
	case "?":
		return object.Bool(s.LastSuccess), true
	case "matches":
		if s.Matches == nil {
			return object.Null(), true
		}
		return s.Matches, true
	case "error":
		arr := object.Array(s.ErrorRecords)
		// 标记为 $Error 的动态视图：Clear/RemoveAt 等方法经 Session 落到 ErrorRecords 本体
		arr.AddProp(ErrorViewMarker, object.Str("1"))
		return arr, true
	case "erroractionpreference":
		// 首选项变量：未赋值或已清空时按 Continue 处理
		if v, ok := s.Vars[s.varKey(name)]; ok && !v.IsNull() {
			return v, true
		}
		return object.Str("Continue"), true
	case "pscommandpath":
		if s.PSCommandPath != "" {
			return object.Str(s.PSCommandPath), true
		}
		return object.Null(), true
	case "args":
		if len(s.Args) > 0 {
			return object.Array(s.Args), true
		}
		return object.Array(nil), true
	case "host":
		return s.HostObject(), true
	case "psedition":
		if s.Style == StyleDesktop {
			return object.Str("Desktop"), true
		}
		return object.Str("Core"), true
	case "islinux":
		if s.Style == StyleDesktop {
			return nil, false // 5.X 不存在此变量
		}
		return object.Bool(runtime.GOOS == "linux"), true
	case "iswindows":
		if s.Style == StyleDesktop {
			return nil, false
		}
		return object.Bool(runtime.GOOS == "windows"), true
	case "ismacos":
		if s.Style == StyleDesktop {
			return nil, false
		}
		return object.Bool(runtime.GOOS == "darwin"), true
	case "iscoreclr":
		return object.Bool(s.Style != StyleDesktop), true
	case "pshome":
		if h, err := os.UserHomeDir(); err == nil {
			return object.Str(filepath.Join(h, ".local", "share", "powershell")), true
		}
		return object.Null(), true
	case "ofs":
		return object.Str(" "), true
	}
	return nil, false
}

// AllVarNames 列出全部可见变量名（用于 Get-Variable / 补全），不区分大小写去重，显式变量优先。
// s.Vars中同时存在不同写法时，实际选择的写法取决于map遍历顺序，不过写入的时候已经进行过归一键，应该不会有这种情况。
func (s *Session) AllVarNames() []string {
	byLower := map[string]string{}
	for n := range s.Vars {
		if _, ok := byLower[strings.ToLower(n)]; !ok {
			byLower[strings.ToLower(n)] = n
		}
	}
	for _, n := range []string{"PWD", "HOME", "PID", "PSVersionTable", "LASTEXITCODE", "?", "Matches", "Error", "PSCommandPath", "args", "Host", "PSEdition", "IsLinux", "IsWindows", "IsMacOS", "IsCoreCLR", "PSHOME", "OFS", "ErrorActionPreference"} {
		if _, ok := byLower[strings.ToLower(n)]; !ok {
			byLower[strings.ToLower(n)] = n
		}
	}
	names := make([]string, 0, len(byLower))
	for _, n := range byLower {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

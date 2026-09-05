// Package shell 管理解释器会话状态：变量、别名、函数、当前目录、5.X/7.X 样式、$PSVersionTable、历史等。
package shell

import (
	"io"
	"os"

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
	Style          Style
	Lang           Lang                        // 界面语言（lang 包的语言码）
	Vars           map[string]*object.PSObject // 用户显式变量
	Aliases        map[string]string
	Functions      map[string]*Function
	Cwd            string
	DirStack       []string // Push-Location / Pop-Location 的目录栈
	History        []string
	HistoryFile    string // 历史文件路径（空则不持久化）
	LastExit       int
	LastSuccess    bool               // $?
	NonInteractive bool               // -NonInteractive：确认提示直接拒绝，读取输入报错
	ErrorRecords   []*object.PSObject // $Error：本会话累积的错误记录，最新在前
	ErrorSeq       uint64             // 错误记录序列号，每累积一条加一（求值器据此判断某段执行是否产生新错误）
	MemberReadSeq  uint64             // 成员读取未产生新错误时记下的错误序号（供赋值收尾判定）
	Matches        *object.PSObject   // $Matches：最近一次标量 -match 的捕获组，未匹配过为 nil
	PSCommandPath  string
	Args           []*object.PSObject // 脚本/函数实参（$args）
	HostOut        io.Writer
	HostErr        io.Writer
	HostIn         io.Reader
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

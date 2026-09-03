// Package parser 实现 PowerShell 语法分析（递归下降）。
//
// 解析策略要点：
//   - 语句（; 或换行分隔）→ 管道（| 分隔）→ 命令。
//   - 命令参数采用"实参模式"：裸字默认是字符串，且相邻 token 可合并为一个裸字（如 a=b、2+3）。
//     比较运算符（-eq 等）会把参数转为比较表达式，以便 Where-Object Length -gt 100 这类写法可用。
//   - 命名参数（-Name value / -Name:value）与开关（-Force）在 AST 中保留。
//     参数绑定阶段由求值器依据各 cmdlet 的参数定义裁决。
package parser

import (
	"strings"

	"powershell/internal/ast"
	"powershell/internal/lexer"
)

// 便捷别名，避免处处写 lexer.TkXXX。
const (
	TkEOF      = lexer.TkEOF
	TkNewline  = lexer.TkNewline
	TkWord     = lexer.TkWord
	TkNumber   = lexer.TkNumber
	TkString   = lexer.TkString
	TkVariable = lexer.TkVariable
	TkBraceVar = lexer.TkBraceVar
	TkDashWord = lexer.TkDashWord
	TkColon    = lexer.TkColon
	TkDot      = lexer.TkDot
	TkOp       = lexer.TkOp
	TkPunct    = lexer.TkPunct
)

// Result 是解析结果。
type Result struct {
	List       *ast.StatementList
	Incomplete bool // 输入不完整（未闭合括号/尾随管道/运算符等），REPL 用于续行
	Error      error
}

// Parse 解析一段 PowerShell 源码。
func Parse(src string) *Result {
	p := &Parser{src: src, toks: lexer.New(src).Tokens()}
	list := p.parseStatementList(0)
	if p.err != nil {
		return &Result{List: list, Incomplete: false, Error: p.err}
	}
	// 尾随反引号（行续行）也算不完整
	trimmed := strings.TrimRight(src, " \t\r")
	incomplete := p.incomplete
	if strings.HasSuffix(trimmed, "`") {
		incomplete = true
	}
	return &Result{List: list, Incomplete: incomplete, Error: nil}
}

// Parser 是递归下降解析器。
type Parser struct {
	src        string
	toks       []lexer.Token
	pos        int
	err        error
	incomplete bool
}


package shell

import (
	"fmt"
	"strings"
)

// Banner 返回启动横幅（按风格与界面语言）。
func (s *Session) Banner() string {
	if s.Style == StyleDesktop {
		if s.Lang == LangZh {
			return "Linux PowerShell by SweetenedSuzuka\n" +
				"版权没有(C) MacroHard Corporation。无法保留所有权利。\n\n" +
				"尝试新的跨平台 PowerShell https://aka.ms/pscore6\n" +
				"别想了，你试不了……等等，PowerShell 7好像还真支持Linux。"
		}
		return "Linux PowerShell by SweetenedSuzuka\n" +
			"Copyright (c) MacroHard Corporation. No rights reserved.\n\n" +
			"Try the new cross-platform PowerShell https://aka.ms/pscore6\n" +
			"Don't even try — you can't... wait, PowerShell 7 actually supports Linux."
	}
	if s.Lang == LangZh {
		return "PowerShell For Linux by SweetenedSuzuka\n" +
			"版权没有(C) MacroHard Corporation。无法保留所有权利。\n\n" +
			"输入 help 查看帮助。"
	}
	return "PowerShell For Linux by SweetenedSuzuka\n" +
		"Copyright (c) MacroHard Corporation. No rights reserved.\n\n" +
		"Type 'help' to get help."
}

// Prompt 返回主提示符（不含结尾空格）。
func (s *Session) Prompt() string {
	return "PS " + s.DisplayPath(s.Cwd) + ">"
}

// ContinuationPrompt 返回多行续行提示符。
func (s *Session) ContinuationPrompt() string {
	return ">>"
}

// PromptString 返回带颜色的提示符文本。
func (s *Session) PromptString() string {
	return s.Prompt()
}

// Usage 打印用法说明（-Help）。
func (s *Session) Usage() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s\n\n", s.Banner())
	b.WriteString(`用法:
  powershell [-Version 5.1|7] [-NoLogo] [-NoProfile] [-Command <命令>] [-File <脚本>]

参数:
  -Version <5.1|7>  选择 PowerShell 风格（5.X 或 7.X）
  -NoLogo           不显示启动横幅
  -NoProfile        不加载启动脚本
  -Command <命令>   执行命令后退出（- 表示从 stdin 读取）
  -File <脚本>      执行 .ps1 脚本后退出
  -Help, -?         显示本帮助

运行时切换风格:
  Set-PSVersion 5.1   切到 Windows PowerShell 5.X 风格
  Set-PSVersion 7     切到 PowerShell 7.X 风格
`)
	return b.String()
}

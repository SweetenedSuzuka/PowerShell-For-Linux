package shell

import (
	"fmt"
	"strings"

	"powershell/internal/lang"
)

// Banner 返回启动横幅（按风格与界面语言）。
func (s *Session) Banner() string {
	if s.Style == StyleDesktop {
		return lang.T(lang.MsgBannerDesktop)
	}
	return lang.T(lang.MsgBannerCore)
}

// Prompt 返回主提示符（不含结尾空格）。
func (s *Session) Prompt() string {
	return "PS " + s.DisplayPath(s.Cwd) + ">"
}

// ContinuationPrompt 返回多行续行提示符。
func (s *Session) ContinuationPrompt() string {
	return ">>"
}

// PromptString 返回提示符文本。
func (s *Session) PromptString() string {
	return s.Prompt()
}

// Usage 返回用法说明文本（-Help）。
func (s *Session) Usage() string {
	var b strings.Builder
	fmt.Fprintf(&b, "%s\n\n", s.Banner())
	b.WriteString(lang.T(lang.MsgUsage))
	return b.String()
}

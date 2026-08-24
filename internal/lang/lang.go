// Package lang 提供界面语言选择与提示文本。
// 界面语言按照优先顺序依次匹配环境变量的 LANGUAGE、LC_ALL、LC_MESSAGES、LANG 来决定，把每个变量的候选 locale 串成一条有序列表。
// 以第一个成功匹配的已注册语言为准，全部未命中时使用中文。
// 文本按语言分文件存放，每种语言一张表，如 zh.go、en.go。新增语言就新建一个文件定义表并在 catalogs 里注册，该语言暂缺的条目在取用时自动回退默认语言。
package lang

import (
	"fmt"
	"os"
	"strings"
)

// Lang 是界面语言代码，取 locale 的语言码部分（zh_CN.UTF-8 → zh）。
type Lang string

const (
	LangZh Lang = "zh"
	LangEn Lang = "en"
)

// current 是当前界面语言，包加载时按环境探测一次。
var current = Detect()

// Current 返回当前界面语言。
func Current() Lang { return current }

// SetCurrent 设置当前界面语言（shell.New 启动时调用，让会话与全局保持一致）。
func SetCurrent(l Lang) { current = l }

// catalogs 注册各语言的文本表。新增语言：新建对应文件定义一张表，这里加一行注册。
var catalogs = map[Lang]map[Msg]string{
	LangZh: zh,
	LangEn: en,
}

// defaultLang 是默认语言：探测不到已注册语言以及语言表缺条目时的归宿。
const defaultLang = LangZh

// localeVars 是探测顺序，遵循 GNU gettext 约定：LANGUAGE 优先于各 LC 变量。
var localeVars = []string{"LANGUAGE", "LC_ALL", "LC_MESSAGES", "LANG"}

// Detect 按环境变量选出界面语言。
// LANGUAGE 以冒号分隔多项，其余变量视为单项；语言码命中已注册语言的候选项立即采用，否则看下一项。全部看完没有命中时返回默认语言。
func Detect() Lang {
	for _, k := range localeVars {
		v := os.Getenv(k)
		if v == "" {
			continue
		}
		for _, cand := range strings.Split(v, ":") {
			cand = strings.TrimSpace(cand)
			if cand == "" || cand == "C" || strings.HasPrefix(cand, "C.") || cand == "POSIX" {
				continue
			}
			if l := Lang(languageCode(cand)); langRegistered(l) {
				return l
			}
		}
	}
	return defaultLang
}

// langRegistered 报告某语言码是否已注册文本表。
func langRegistered(l Lang) bool {
	_, ok := catalogs[l]
	return ok
}

// languageCode 取 locale 的语言码（zh_CN.UTF-8 → zh；fr_FR → fr）。
func languageCode(locale string) string {
	if i := strings.IndexAny(locale, "_@:-."); i >= 0 {
		return locale[:i]
	}
	return locale
}

// T 取当前语言的文本并填入参数；当前语言表没有这一条时回退默认语言表。
// 实参顺序以默认语言表为准，其它语言语序不同的改写措辞去对齐它，不用位置占位符。
func T(m Msg, args ...any) string {
	s, ok := catalogs[current][m]
	if !ok {
		s, ok = catalogs[defaultLang][m]
		if !ok {
			return ""
		}
	}
	if len(args) > 0 {
		return fmt.Sprintf(s, args...)
	}
	return s
}

package shell

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"powershell/internal/lang"
)

// drive.go 实现路径的盘符归一、解析与显示（C 盘即根目录，显示保持反斜杠风格）。

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

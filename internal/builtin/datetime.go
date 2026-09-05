package builtin

import (
	"fmt"
	"strings"
	"time"

	"powershell/internal/lang"
	"powershell/internal/object"
)

// datetime.go 实现日期时间类 cmdlet。

func cmdGetDate(c *Context) ([]*object.PSObject, error) {
	now := time.Now()
	// -Date 指定日期时间（本地时区，无时区信息时按本地解析）；解析失败报错，不回退当前时间。
	if d, ok := c.Args.Str("Date"); ok && d != "" {
		t, err := parseDateArg(d)
		if err != nil {
			return errf(c, "%v", err)
		}
		now = t
	}
	if f, ok := c.Args.Str("Format"); ok && f != "" {
		return []*object.PSObject{object.Str(now.Format(dotnetToGoLayout(f)))}, nil
	}
	// Get-Date 输出带 DisplayHint 属性（与 PowerShell 一致，类型转换来的时间对象没有）。
	o := object.DateTime(now)
	o.AddProp("DisplayHint", "DateTime")
	return []*object.PSObject{o}, nil
}

// parseDateArg 按常见日期时间格式解析 -Date 参数（无时区信息按本地时区）。
func parseDateArg(s string) (time.Time, error) {
	layouts := []string{
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05",
		"2006-01-02",
		"2006-1-2 15:04:05",
		"2006-1-2T15:04:05",
		"2006-1-2",
		"2006/1/2 15:04:05",
		"2006/1/2T15:04:05",
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

func cmdSetDate(c *Context) ([]*object.PSObject, error) {
	date := firstArg(c, "Date")
	if date == "" {
		return []*object.PSObject{object.DateTime(time.Now())}, nil
	}
	runExternalRaw(c, "sudo", []string{"date", "-s", date})
	return []*object.PSObject{object.DateTime(time.Now())}, nil
}

// ---- 注册 ----

func init() {
	Register("Get-Date", []ParamSpec{
		{Name: "Date", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Format", Type: "string"},
	}, cmdGetDate)
	Register("Set-Date", []ParamSpec{
		{Name: "Date", Position: 0, PositionSet: true, Type: "string"},
	}, cmdSetDate)
}

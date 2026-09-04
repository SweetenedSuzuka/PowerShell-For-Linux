// Package object 定义 PowerShell 对象模型（PSObject）与格式化输出。
//
// 内置 cmdlet 产出的是带类型与属性的对象，而非纯文本。
// 管道终点再按对象的形状（表格/列表/纯文本）渲染。
package object

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"powershell/internal/ast"
	"powershell/internal/lang"
)

// PSObject 是 PowerShell 对象。
type PSObject struct {
	TypeName string   // 类型名，如 "System.IO.FileInfo"、"String"
	Value    any      // 底层值：string/int64/float64/bool/[]*PSObject/...；nil 表示 $null
	Props    []Prop   // 属性列表（保持顺序）
	Table    []Column // 表格列定义；为空时按默认规则
}

// Prop 是一个具名属性。
type Prop struct {
	Name  string
	Value any // 标量（string/int64/...）或 *PSObject
}

// Column 是表格的一列。
type Column struct {
	Label string
	Width int    // 0 = 自适应
	Align string // "left" / "right"
}

// ---- 构造器 ----

// Str 创建字符串对象。
func Str(s string) *PSObject { return &PSObject{TypeName: "String", Value: s} }

// Int 创建整数对象。
func Int(v int64) *PSObject { return &PSObject{TypeName: "Int", Value: v} }

// Float 创建浮点对象。
func Float(v float64) *PSObject { return &PSObject{TypeName: "Double", Value: v} }

// Bool 创建布尔对象。
func Bool(b bool) *PSObject { return &PSObject{TypeName: "Boolean", Value: b} }

// Null 创建 $null 对象。
func Null() *PSObject { return &PSObject{TypeName: "Null", Value: nil} }

// Array 创建数组对象。
func Array(items []*PSObject) *PSObject {
	return &PSObject{TypeName: "Object[]", Value: items}
}

// HashEntry 是哈希表的一个键值对（保持顺序）。
type HashEntry struct {
	Key   string
	Value *PSObject
}

// Hashtable 创建哈希表对象。
func Hashtable(entries []HashEntry) *PSObject {
	return &PSObject{TypeName: "Hashtable", Value: entries}
}

// PSCustomObject 创建自定义对象：哈希表条目按顺序变成属性（[pscustomobject] 与 New-Object 共用）。
func PSCustomObject(entries []HashEntry) *PSObject {
	o := &PSObject{TypeName: "System.Management.Automation.PSCustomObject"}
	for _, en := range entries {
		o.AddProp(en.Key, en.Value)
	}
	return o
}

// DateTime 创建时间对象。
func DateTime(t time.Time) *PSObject {
	return &PSObject{TypeName: "DateTime", Value: t}
}

// versionParts 存放四段版本号（major.minor.build.revision）。
type versionParts struct {
	major, minor, build, revision int
}

// String 渲染版本号：
// - 从右向左裁剪所有值为 0 或 -1 的尾部段（-1 表示该段在解析时缺失）。
// - 至少保留主版本号（major），即使它为 0。
// - 非尾部段的 0 会正常显示。
//
// 示例：
//
//	[7,0,0,0]   -> "7"
//	[1,2,-1,-1] -> "1.2"
//	[1,0,3,0]   -> "1.0.3"  (中间的 0 保留)
//	[0,0,0,0]   -> "0"      (至少保留 major)
//	[1,-1,0,0]  -> "1"      (尾部 -1 和 0 均被裁剪)
func (v versionParts) String() string {
	segs := []int{v.major, v.minor, v.build, v.revision}
	end := len(segs) - 1
	for end > 0 && (segs[end] == 0 || segs[end] == -1) {
		end--
	}
	var sb strings.Builder
	for i := 0; i <= end; i++ {
		if i > 0 {
			sb.WriteByte('.')
		}
		sb.WriteString(strconv.Itoa(segs[i]))
	}
	return sb.String()
}

// Version 创建版本对象（TypeName "System.Version"），模拟 .NET 的 System.Version。
// 用于 $PSVersionTable.PSVersion，让脚本能读取 .Major/.Minor/.Build/.Revision。
func Version(major, minor, build, revision int) *PSObject {
	return &PSObject{TypeName: "System.Version", Value: versionParts{major, minor, build, revision}}
}

// Error 创建错误记录对象。
func Error(msg string) *PSObject {
	o := &PSObject{TypeName: "ErrorRecord", Value: msg}
	o.AddProp("Message", msg)
	return o
}

// ScriptBlock 创建脚本块对象，Value 存脚本块节点供 & 调用与 .Invoke 执行。
func ScriptBlock(sb *ast.ScriptBlock) *PSObject {
	return &PSObject{TypeName: "ScriptBlock", Value: sb}
}

// Object 创建带属性的普通对象。
func Object(typeName string, value any) *PSObject {
	return &PSObject{TypeName: typeName, Value: value}
}

// AddProp 添加属性。
func (o *PSObject) AddProp(name string, value any) *PSObject {
	o.Props = append(o.Props, Prop{Name: name, Value: value})
	return o
}

// SetTable 设置表格列。
func (o *PSObject) SetTable(cols ...Column) *PSObject {
	o.Table = cols
	return o
}

// Clone 复制对象（浅拷贝属性值）。
func (o *PSObject) Clone() *PSObject {
	n := &PSObject{TypeName: o.TypeName, Value: o.Value, Table: o.Table}
	n.Props = append(n.Props, o.Props...)
	return n
}

// IsNull 判断是否为 $null（带属性的对象即使 Value 为 nil 也不算空）。
func (o *PSObject) IsNull() bool {
	return o == nil || (o.Value == nil && len(o.Props) == 0)
}

// IsArray 判断是否为数组。
func (o *PSObject) IsArray() bool { _, ok := o.Value.([]*PSObject); return ok }

// ArrayItems 返回数组元素；非数组则返回单元素切片。
func (o *PSObject) ArrayItems() []*PSObject {
	if items, ok := o.Value.([]*PSObject); ok {
		return items
	}
	return []*PSObject{o}
}

// Unwrap 返回标量值（数组返回其首个元素；字符串原样）。
func (o *PSObject) Unwrap() any {
	if items, ok := o.Value.([]*PSObject); ok {
		if len(items) == 0 {
			return nil
		}
		return items[0].Unwrap()
	}
	if hs, ok := o.Value.([]HashEntry); ok {
		return hs
	}
	return o.Value
}

// PropValue 按名字取属性；不存在返回 nil, false。
// 先查 Props，再查虚拟属性（DateTime 字段、字符串 Length、数组 Count、哈希表键）。
func (o *PSObject) PropValue(name string) (*PSObject, bool) {
	for _, p := range o.Props {
		if strings.EqualFold(p.Name, name) {
			return ToPS(p.Value), true
		}
	}
	return o.virtualProp(name)
}

// virtualProp 解析不在 Props 里的虚拟属性（与 eval 的成员访问一致）。
// Select-Object/Sort-Object/Group-Object/Measure-Object 等按名取属性都要经过它。
func (o *PSObject) virtualProp(name string) (*PSObject, bool) {
	switch o.TypeName {
	case "String":
		if strings.EqualFold(name, "Length") {
			return Int(int64(utf8.RuneCountInString(o.String()))), true
		}
	case "Object[]":
		if strings.EqualFold(name, "Length") || strings.EqualFold(name, "Count") {
			return Int(int64(len(o.ArrayItems()))), true
		}
	case "Hashtable":
		if entries, ok := o.Value.([]HashEntry); ok {
			for _, en := range entries {
				if strings.EqualFold(en.Key, name) {
					return en.Value, true
				}
			}
			// 键未命中时才落到内置属性（PowerShell 键优先于属性）。
			// Count 返回条目数；Keys/Values 返回按插入顺序排列的数组。
			switch strings.ToLower(name) {
			case "count":
				return Int(int64(len(entries))), true
			case "keys":
				keys := make([]*PSObject, 0, len(entries))
				for _, en := range entries {
					keys = append(keys, Str(en.Key))
				}
				return Array(keys), true
			case "values":
				vals := make([]*PSObject, 0, len(entries))
				for _, en := range entries {
					vals = append(vals, en.Value)
				}
				return Array(vals), true
			}
		}
	case "System.IO.FileInfo", "System.IO.DirectoryInfo":
		// 虚拟属性从路径计算：Extension（目录恒空）、BaseName（文件去扩展名）、DirectoryName（父目录）。
		// 与 FullName 一致按传入路径原样计算，不解析为绝对路径（对齐现有简化）。
		if path, ok := o.Value.(string); ok {
			switch strings.ToLower(name) {
			case "extension":
				if o.TypeName == "System.IO.DirectoryInfo" {
					return Str(""), true
				}
				return Str(filepath.Ext(path)), true
			case "basename":
				base := filepath.Base(path)
				if o.TypeName == "System.IO.FileInfo" {
					base = strings.TrimSuffix(base, filepath.Ext(base))
				}
				return Str(base), true
			case "directoryname":
				return Str(filepath.Dir(path)), true
			}
		}
	case "DateTime":
		if t, ok := o.Value.(time.Time); ok {
			switch strings.ToLower(name) {
			case "year":
				return Int(int64(t.Year())), true
			case "month":
				return Int(int64(t.Month())), true
			case "day":
				return Int(int64(t.Day())), true
			case "hour":
				return Int(int64(t.Hour())), true
			case "minute":
				return Int(int64(t.Minute())), true
			case "second":
				return Int(int64(t.Second())), true
			case "dayofweek":
				return Str(t.Weekday().String()), true
			case "dayofyear":
				return Int(int64(t.YearDay())), true
			case "ticks":
				return Int(t.UnixNano()/100 + 621355968000000000), true
			case "millisecond":
				return Int(int64(t.Nanosecond() / 1e6)), true
			case "microsecond":
				return Int(int64(t.Nanosecond() / 1e3 % 1e3)), true
			case "nanosecond":
				return Int(int64(t.Nanosecond() % 1e3)), true
			case "kind":
				return Str(dateTimeKind(t)), true
			case "date":
				midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
				return Str(formatDateTime(midnight)), true
			case "timeofday":
				return Str(formatTimeOfDay(t)), true
			case "datetime":
				return Str(formatDateTime(t)), true
			}
		}
	case "System.Version":
		if v, ok := o.Value.(versionParts); ok {
			switch strings.ToLower(name) {
			case "major":
				return Int(int64(v.major)), true
			case "minor":
				return Int(int64(v.minor)), true
			case "build":
				return Int(int64(v.build)), true
			case "revision":
				return Int(int64(v.revision)), true
			}
		}
	}
	// Count/Length：有值标量为 1，$null 为 0
	if strings.EqualFold(name, "Count") || strings.EqualFold(name, "Length") {
		if o.IsNull() {
			return Int(0), true
		}
		return Int(1), true
	}
	return nil, false
}

// HasProp 判断是否有某属性。
func (o *PSObject) HasProp(name string) bool {
	for _, p := range o.Props {
		if strings.EqualFold(p.Name, name) {
			return true
		}
	}
	return false
}

// SetProp 设置属性值：已存在（不区分大小写）则替换，否则追加。
// 哈希表按条目键替换或追加；存入解包后的值，与 AddProp 口径一致。
func (o *PSObject) SetProp(name string, val *PSObject) *PSObject {
	var v any
	if val != nil {
		v = val.Value
	}
	if entries, ok := o.Value.([]HashEntry); ok && o.TypeName == "Hashtable" {
		for i := range entries {
			if strings.EqualFold(entries[i].Key, name) {
				entries[i].Value = val
				return o
			}
		}
		o.Value = append(entries, HashEntry{Key: name, Value: val})
		return o
	}
	for i := range o.Props {
		if strings.EqualFold(o.Props[i].Name, name) {
			o.Props[i].Value = v
			return o
		}
	}
	o.Props = append(o.Props, Prop{Name: name, Value: v})
	return o
}

// ToPS 把 Go 值包装成 PSObject（标量、切片、哈希条目、时间、错误等）。
func ToPS(v any) *PSObject {
	switch t := v.(type) {
	case *PSObject:
		return t
	case nil:
		return Null()
	case string:
		return Str(t)
	case bool:
		return Bool(t)
	case int:
		return Int(int64(t))
	case int64:
		return Int(t)
	case float64:
		return Float(t)
	case []*PSObject:
		return Array(t)
	case []string:
		items := make([]*PSObject, len(t))
		for i, s := range t {
			items[i] = Str(s)
		}
		return Array(items)
	case time.Time:
		return DateTime(t)
	case error:
		return Error(t.Error())
	case []HashEntry:
		return Hashtable(t)
	default:
		return Str(fmt.Sprintf("%v", v))
	}
}

// Truthy 实现 PowerShell 真值判断。
func (o *PSObject) Truthy() bool {
	if o == nil {
		return false
	}
	switch v := o.Value.(type) {
	case nil:
		return false
	case bool:
		return v
	case int64:
		return v != 0
	case float64:
		return v != 0
	case string:
		return v != ""
	case []*PSObject:
		return len(v) > 0
	case []HashEntry:
		return true // 空哈希表也是真值
	default:
		return true
	}
}

// AsInt 尝试转为整数。
func (o *PSObject) AsInt() (int64, bool) {
	switch v := o.Value.(type) {
	case int64:
		return v, true
	case float64:
		return int64(v), true
	case bool:
		if v {
			return 1, true
		}
		return 0, true
	case string:
		if n, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64); err == nil {
			return n, true
		}
	}
	return 0, false
}

// AsFloat 尝试转为浮点。
func (o *PSObject) AsFloat() (float64, bool) {
	switch v := o.Value.(type) {
	case int64:
		return float64(v), true
	case float64:
		return v, true
	case bool:
		if v {
			return 1, true
		}
		return 0, true
	case string:
		if n, err := strconv.ParseFloat(strings.TrimSpace(v), 64); err == nil {
			return n, true
		}
	}
	return 0, false
}

// String 返回显示字符串（PowerShell 风格）。
func (o *PSObject) String() string {
	if o == nil {
		return ""
	}
	switch v := o.Value.(type) {
	case nil:
		return ""
	case string:
		return v
	case bool:
		if v {
			return "True"
		}
		return "False"
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return formatFloat(v)
	case []*PSObject:
		parts := make([]string, 0, len(v))
		for _, it := range v {
			parts = append(parts, it.String())
		}
		return strings.Join(parts, " ")
	case []HashEntry:
		return hashtableString(v)
	case time.Time:
		return formatDateTime(v)
	case versionParts:
		return v.String()
	case *ast.ScriptBlock:
		// 脚本块对象：语句列表不还原源码，显示占位
		return "{ ... }"
	case *ast.StatementList:
		return "{ ... }"
	}
	return fmt.Sprintf("%v", o.Value)
}

func formatFloat(v float64) string {
	if v == float64(int64(v)) && v < 1e15 && v > -1e15 {
		return strconv.FormatInt(int64(v), 10)
	}
	s := strconv.FormatFloat(v, 'f', -1, 64)
	return s
}

// formatDateTime 按界面语言渲染 DateTime。
// 各语言的日期格式在此函数登记，尚未登记格式的语言回退默认语言的中文格式。
func formatDateTime(t time.Time) string {
	if lang.Current() == lang.LangEn {
		return t.Format("Monday, 2 January 2006 15:04:05")
	}
	return t.Format("2006年1月2日") + weekdayZh[t.Weekday()] + t.Format(" 15:04:05")
}

// dateTimeKind 按时区名报告时间的种类：本地为 Local，UTC 为 Utc，其余为 Unspecified。
func dateTimeKind(t time.Time) string {
	switch t.Location().String() {
	case "Local":
		return "Local"
	case "UTC":
		return "Utc"
	}
	return "Unspecified"
}

// formatTimeOfDay 渲染当日已过时间（时:分:秒，亚秒部分去尾零）。
func formatTimeOfDay(t time.Time) string {
	s := fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
	if ns := t.Nanosecond(); ns != 0 {
		s += "." + strings.TrimRight(fmt.Sprintf("%09d", ns), "0")
	}
	return s
}

// weekdayZh 是星期的中文名（Go 布局没有中文占位符，自行映射）。
var weekdayZh = map[time.Weekday]string{
	time.Sunday: "星期日", time.Monday: "星期一", time.Tuesday: "星期二",
	time.Wednesday: "星期三", time.Thursday: "星期四", time.Friday: "星期五",
	time.Saturday: "星期六",
}

func hashtableString(entries []HashEntry) string {
	parts := make([]string, 0, len(entries))
	for _, e := range entries {
		parts = append(parts, e.Key+"="+e.Value.String())
	}
	return "{" + strings.Join(parts, "; ") + "}"
}

// ---- 文件对象 ----

// FileInfo 创建文件信息对象（Get-ChildItem / Get-Item 用）。
func FileInfo(path string, info os.FileInfo) *PSObject {
	o := &PSObject{TypeName: "System.IO.FileInfo", Value: path}
	o.AddProp("Name", info.Name())
	o.AddProp("FullName", path)
	o.AddProp("Mode", UnixMode(info))
	o.AddProp("Length", info.Size())
	o.AddProp("LastWriteTime", info.ModTime())
	o.AddProp("PSIsContainer", false)
	o.Table = []Column{
		{Label: "Mode", Align: "left"},
		{Label: "LastWriteTime", Align: "left"},
		{Label: "Length", Align: "right"},
		{Label: "Name", Align: "left"},
	}
	return o
}

// DirInfo 创建目录信息对象。
func DirInfo(path string, info os.FileInfo) *PSObject {
	o := &PSObject{TypeName: "System.IO.DirectoryInfo", Value: path}
	o.AddProp("Name", info.Name())
	o.AddProp("FullName", path)
	o.AddProp("Mode", UnixMode(info))
	o.AddProp("Length", nil)
	o.AddProp("LastWriteTime", info.ModTime())
	o.AddProp("PSIsContainer", true)
	o.Table = []Column{
		{Label: "Mode", Align: "left"},
		{Label: "LastWriteTime", Align: "left"},
		{Label: "Length", Align: "right"},
		{Label: "Name", Align: "left"},
	}
	return o
}

// UnixMode 生成类 Unix 权限字符串（如 drwxr-xr-x）。
func UnixMode(info os.FileInfo) string {
	var sb strings.Builder
	if info.IsDir() {
		sb.WriteByte('d')
	} else {
		sb.WriteByte('-')
	}
	m := info.Mode().Perm()
	for _, group := range []os.FileMode{0o400, 0o040, 0o004} {
		r := m & group
		w := m & (group >> 1)
		x := m & (group >> 2)
		if r != 0 {
			sb.WriteByte('r')
		} else {
			sb.WriteByte('-')
		}
		if w != 0 {
			sb.WriteByte('w')
		} else {
			sb.WriteByte('-')
		}
		if x != 0 {
			sb.WriteByte('x')
		} else {
			sb.WriteByte('-')
		}
	}
	return sb.String()
}

// ---- 进程对象 ----

// Process 创建进程信息对象（Get-Process 用）。
func Process(pid int, name string, cpu float64, mem int64) *PSObject {
	o := &PSObject{TypeName: "System.Diagnostics.Process", Value: pid}
	o.AddProp("Id", int64(pid))
	o.AddProp("Name", name)
	o.AddProp("ProcessName", name)
	o.AddProp("CPU", cpu)
	o.AddProp("Memory", mem)
	o.Table = []Column{
		{Label: "Id", Align: "right"},
		{Label: "ProcessName", Align: "left"},
		{Label: "CPU", Align: "right"},
		{Label: "Memory", Align: "right"},
	}
	return o
}

// VersionTable 创建 $PSVersionTable 哈希表对象。
func VersionTable(entries []HashEntry) *PSObject {
	return Hashtable(entries)
}

// WildcardMatch 实现 PowerShell 风格通配符匹配（* ? [a-z] [!a-z]，* 可匹配任意字符）。
func WildcardMatch(pattern, s string) bool { return wildMatch(pattern, s) }

// WildcardMatchFold 大小写不敏感的通配符匹配（PowerShell 比较运算符默认语义）。
func WildcardMatchFold(pattern, s string) bool {
	return wildMatch(strings.ToLower(pattern), strings.ToLower(s))
}

func wildMatch(p, s string) bool {
	for len(p) > 0 {
		switch p[0] {
		case '*':
			for len(p) > 0 && p[0] == '*' {
				p = p[1:]
			}
			if len(p) == 0 {
				return true
			}
			for i := 0; i <= len(s); i++ {
				if wildMatch(p, s[i:]) {
					return true
				}
			}
			return false
		case '?':
			if len(s) == 0 {
				return false
			}
			p, s = p[1:], s[1:]
		case '[':
			end := strings.IndexByte(p, ']')
			if end < 0 {
				if len(s) == 0 || s[0] != '[' {
					return false
				}
				p, s = p[1:], s[1:]
				continue
			}
			if len(s) == 0 {
				return false
			}
			class := p[1:end]
			neg := false
			if len(class) > 0 && class[0] == '!' {
				neg = true
				class = class[1:]
			}
			matched := classContains(class, s[0])
			if neg {
				matched = !matched
			}
			if !matched {
				return false
			}
			p, s = p[end+1:], s[1:]
		case '`':
			if len(p) > 1 {
				p = p[1:]
				if len(s) == 0 || s[0] != p[0] {
					return false
				}
				p, s = p[1:], s[1:]
			} else {
				p = p[1:]
			}
		default:
			if len(s) == 0 || s[0] != p[0] {
				return false
			}
			p, s = p[1:], s[1:]
		}
	}
	return len(s) == 0
}

func classContains(class string, c byte) bool {
	i := 0
	for i < len(class) {
		if i+2 < len(class) && class[i+1] == '-' {
			lo, hi := class[i], class[i+2]
			if c >= lo && c <= hi {
				return true
			}
			i += 3
		} else {
			if class[i] == c {
				return true
			}
			i++
		}
	}
	return false
}

package builtin

import (
	"bytes"
	"crypto/rand"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"powershell/internal/ast"
	"powershell/internal/object"
)

// ---- 数据转换与对象 ----

func csvHeader(o *object.PSObject) []string {
	var names []string
	for _, p := range o.Props {
		names = append(names, p.Name)
	}
	return names
}

func cmdConvertToCsv(c *Context) ([]*object.PSObject, error) {
	items := inputItems(c)
	if len(items) == 0 {
		return nil, nil
	}
	var header []string
	if props := c.Args.StringSlice("Property"); len(props) > 0 {
		header = props
	} else {
		header = csvHeader(items[0])
		if len(header) == 0 {
			header = []string{"Value"}
		}
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	_ = w.Write(header)
	for _, it := range items {
		row := make([]string, len(header))
		for i, h := range header {
			if h == "Value" {
				row[i] = it.String()
				continue
			}
			if v, ok := it.PropValue(h); ok {
				row[i] = v.String()
			}
		}
		_ = w.Write(row)
	}
	w.Flush()
	text := strings.TrimRight(buf.String(), "\n")
	var out []*object.PSObject
	for _, ln := range strings.Split(text, "\n") {
		out = append(out, object.Str(ln))
	}
	return out, nil
}

func cmdConvertFromCsv(c *Context) ([]*object.PSObject, error) {
	var text string
	if len(c.Input) > 0 {
		var sb strings.Builder
		for _, o := range c.Input {
			sb.WriteString(o.String())
			sb.WriteByte('\n')
		}
		text = sb.String()
	} else if v := c.Args.Get("InputObject"); v != nil {
		text = v.String()
	}
	r := csv.NewReader(strings.NewReader(text))
	records, err := r.ReadAll()
	if err != nil {
		return errf(c, "ConvertFrom-Csv : %v", err)
	}
	if len(records) < 2 {
		return nil, nil
	}
	header := records[0]
	var out []*object.PSObject
	for _, rec := range records[1:] {
		o := object.Object("System.Management.Automation.PSCustomObject", nil)
		for i, h := range header {
			if i < len(rec) {
				o.AddProp(h, rec[i])
			}
		}
		out = append(out, o)
	}
	return out, nil
}

func cmdConvertToJson(c *Context) ([]*object.PSObject, error) {
	items := inputItems(c)
	// -Depth 限制嵌套展开深度（默认 2，对齐 PowerShell）
	depth := 2
	if d, ok := c.Args.Int("Depth"); ok && d >= 0 {
		depth = int(d)
	}
	// 单对象原样；多对象成数组
	var buf bytes.Buffer
	if len(items) == 1 {
		writeJSON(&buf, items[0], 0, depth)
	} else {
		buf.WriteByte('[')
		for i, it := range items {
			buf.WriteByte('\n')
			buf.WriteString(strings.Repeat("  ", 1))
			writeJSON(&buf, it, 1, depth)
			if i < len(items)-1 {
				buf.WriteByte(',')
			}
		}
		if len(items) > 0 {
			buf.WriteByte('\n')
		}
		buf.WriteByte(']')
	}
	return []*object.PSObject{object.Str(buf.String())}, nil
}

// writeJSON 把 PSObject 序列化为美化 JSON（2 空格缩进），保持对象/哈希表的键序。
// remaining 为剩余可展开深度；到 0 时数组/对象不再展开（对齐 PowerShell 的 -Depth 截断）。
func writeJSON(buf *bytes.Buffer, o *object.PSObject, indent, remaining int) {
	if o == nil || o.IsNull() {
		buf.WriteString("null")
		return
	}
	if o.IsArray() {
		if remaining <= 0 {
			buf.WriteString("[]")
			return
		}
		items := o.ArrayItems()
		if len(items) == 0 {
			buf.WriteString("[]")
			return
		}
		buf.WriteString("[\n")
		for i, it := range items {
			buf.WriteString(strings.Repeat("  ", indent+1))
			writeJSON(buf, it, indent+1, remaining-1)
			if i < len(items)-1 {
				buf.WriteByte(',')
			}
			buf.WriteByte('\n')
		}
		buf.WriteString(strings.Repeat("  ", indent))
		buf.WriteByte(']')
		return
	}
	if o.TypeName == "Hashtable" {
		if entries, ok := o.Value.([]object.HashEntry); ok {
			if remaining <= 0 || len(entries) == 0 {
				buf.WriteString("{}")
				return
			}
			buf.WriteString("{\n")
			for i, en := range entries {
				buf.WriteString(strings.Repeat("  ", indent+1))
				writeJSONString(buf, en.Key)
				buf.WriteString(": ")
				writeJSON(buf, en.Value, indent+1, remaining-1)
				if i < len(entries)-1 {
					buf.WriteByte(',')
				}
				buf.WriteByte('\n')
			}
			buf.WriteString(strings.Repeat("  ", indent))
			buf.WriteByte('}')
			return
		}
	}
	if len(o.Props) > 0 {
		if remaining <= 0 {
			buf.WriteString("{}")
			return
		}
		buf.WriteString("{\n")
		for i, p := range o.Props {
			buf.WriteString(strings.Repeat("  ", indent+1))
			writeJSONString(buf, p.Name)
			buf.WriteString(": ")
			writeJSON(buf, object.ToPS(p.Value), indent+1, remaining-1)
			if i < len(o.Props)-1 {
				buf.WriteByte(',')
			}
			buf.WriteByte('\n')
		}
		buf.WriteString(strings.Repeat("  ", indent))
		buf.WriteByte('}')
		return
	}
	writeJSONScalar(buf, o)
}

// writeJSONString 输出带转义的 JSON 字符串。
func writeJSONString(buf *bytes.Buffer, s string) {
	b, _ := json.Marshal(s)
	buf.Write(b)
}

// writeJSONScalar 输出 JSON 标量或 null。
func writeJSONScalar(buf *bytes.Buffer, o *object.PSObject) {
	var v any
	switch o.Value.(type) {
	case nil:
		v = nil
	case string:
		v = o.Value
	case bool:
		v = o.Value
	case int64:
		v = o.Value
	case float64:
		v = o.Value
	default:
		v = o.String()
	}
	b, err := json.Marshal(v)
	if err != nil {
		buf.WriteString("null")
		return
	}
	buf.Write(b)
}

func cmdConvertFromJson(c *Context) ([]*object.PSObject, error) {
	var text string
	if len(c.Input) > 0 {
		var sb strings.Builder
		for _, o := range c.Input {
			sb.WriteString(o.String())
		}
		text = sb.String()
	} else if v := c.Args.Get("InputObject"); v != nil {
		text = v.String()
	}
	var val any
	if err := json.Unmarshal([]byte(text), &val); err != nil {
		return errf(c, "ConvertFrom-Json : %v", err)
	}
	return []*object.PSObject{jsonToObject(val)}, nil
}

func jsonToObject(v any) *object.PSObject {
	switch t := v.(type) {
	case map[string]any:
		o := object.Object("System.Management.Automation.PSCustomObject", nil)
		keys := make([]string, 0, len(t))
		for k := range t {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			o.AddProp(k, jsonToObject(t[k]))
		}
		return o
	case []any:
		items := make([]*object.PSObject, 0, len(t))
		for _, it := range t {
			items = append(items, jsonToObject(it))
		}
		return object.Array(items)
	case string:
		return object.Str(t)
	case bool:
		return object.Bool(t)
	case float64:
		return object.Float(t)
	case nil:
		return object.Null()
	default:
		return object.Str(fmt.Sprintf("%v", v))
	}
}

func cmdConvertFromStringData(c *Context) ([]*object.PSObject, error) {
	var text string
	if v := c.Args.Get("StringData"); v != nil {
		text = v.String()
	} else if len(c.Input) > 0 {
		// 命名/位置都没给时才用管道输入
		var sb strings.Builder
		for _, o := range c.Input {
			sb.WriteString(o.String())
			sb.WriteByte('\n')
		}
		text = sb.String()
	}
	var entries []object.HashEntry
	for _, ln := range strings.Split(text, "\n") {
		ln = strings.TrimSpace(ln)
		if ln == "" || strings.HasPrefix(ln, "#") {
			continue
		}
		if k, v, ok := strings.Cut(ln, "="); ok {
			entries = append(entries, object.HashEntry{Key: strings.TrimSpace(k), Value: object.Str(strings.TrimSpace(v))})
		}
	}
	return []*object.PSObject{object.Hashtable(entries)}, nil
}

func cmdCompareObject(c *Context) ([]*object.PSObject, error) {
	ref := c.Args.Get("ReferenceObject")
	diff := c.Args.Get("DifferenceObject")
	var refItems, diffItems []*object.PSObject
	if ref != nil {
		refItems = ref.ArrayItems()
	}
	if diff != nil {
		diffItems = diff.ArrayItems()
	}
	// key 函数：默认小写折叠（大小写不敏感），-CaseSensitive 时原样
	caseSensitive := c.Args.Switch("CaseSensitive")
	keyOf := func(s string) string {
		if caseSensitive {
			return s
		}
		return strings.ToLower(s)
	}
	includeEqual := c.Args.Switch("IncludeEqual")
	// ref 与 diff 各自按 key 去重，保留首次原值
	type item struct {
		val  *object.PSObject
		key  string
		side string
	}
	inRef := map[string]bool{}
	var refUnique, diffUnique []item
	for _, r := range refItems {
		k := keyOf(r.String())
		if inRef[k] {
			continue
		}
		inRef[k] = true
		refUnique = append(refUnique, item{r, k, "<="})
	}
	inDiff := map[string]bool{}
	for _, d := range diffItems {
		k := keyOf(d.String())
		if inDiff[k] {
			continue
		}
		inDiff[k] = true
		diffUnique = append(diffUnique, item{d, k, "=>"})
	}
	// 输出顺序：IncludeEqual 时相等项最先，然后右侧独有（=>），再左侧独有（<=）
	// 相等项显示参考集（ref）的值，对齐真 PowerShell
	var equalOut, rightOut, leftOut []*object.PSObject
	seen := map[string]bool{}
	for _, d := range diffUnique {
		if inRef[d.key] {
			if includeEqual && !seen["=="+d.key] {
				seen["=="+d.key] = true
				// 相等项的值取 ref 侧首次出现
				var refVal *object.PSObject
				for _, r := range refUnique {
					if r.key == d.key {
						refVal = r.val
						break
					}
				}
				o := object.Object("System.Management.Automation.PSCustomObject", nil)
				o.AddProp("InputObject", refVal)
				o.AddProp("SideIndicator", "==")
				equalOut = append(equalOut, o)
			}
		} else if !seen["=>"+d.key] {
			seen["=>"+d.key] = true
			o := object.Object("System.Management.Automation.PSCustomObject", nil)
			o.AddProp("InputObject", d.val)
			o.AddProp("SideIndicator", "=>")
			rightOut = append(rightOut, o)
		}
	}
	for _, r := range refUnique {
		if !inDiff[r.key] && !seen["<="+r.key] {
			seen["<="+r.key] = true
			o := object.Object("System.Management.Automation.PSCustomObject", nil)
			o.AddProp("InputObject", r.val)
			o.AddProp("SideIndicator", "<=")
			leftOut = append(leftOut, o)
		}
	}
	out := append(equalOut, rightOut...)
	out = append(out, leftOut...)
	return out, nil
}

func cmdGetUnique(c *Context) ([]*object.PSObject, error) {
	var out []*object.PSObject
	seen := map[string]bool{}
	for _, o := range inputItems(c) {
		k := o.String()
		if !seen[k] {
			seen[k] = true
			out = append(out, o)
		}
	}
	return out, nil
}

func cmdGetRandom(c *Context) ([]*object.PSObject, error) {
	if items := inputItems(c); len(items) > 0 {
		// 从输入随机取（默认 1 个，-Count 可指定）
		n := 1
		if cnt, ok := c.Args.Int("Count"); ok && cnt > 0 {
			n = int(cnt)
		}
		var out []*object.PSObject
		perm := randPerm(len(items))
		for i := 0; i < n && i < len(items); i++ {
			out = append(out, items[perm[i]])
		}
		return out, nil
	}
	mn, mnOK := c.Args.Int("Minimum")
	mx, mxOK := c.Args.Int("Maximum")
	if mnOK || mxOK {
		// 只给一个端点时另一端取默认（PowerShell：Minimum 默认 0，Maximum 默认 int 上限）
		if !mnOK {
			mn = 0
		}
		if !mxOK {
			mx = 0x7fffffff
		}
		return []*object.PSObject{object.Int(randomInRange(mn, mx))}, nil
	}
	// 无参数：随机数
	return []*object.PSObject{object.Int(randomInRange(0, 0x7fffffff))}, nil
}

func randPerm(n int) []int {
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i
	}
	for i := n - 1; i > 0; i-- {
		j := randomInRange(0, int64(i+1))
		perm[i], perm[j] = perm[j], perm[i]
	}
	return perm
}

func randomInRange(lo, hi int64) int64 {
	var b [8]byte
	_, _ = rand.Read(b[:])
	v := int64(0)
	for _, c := range b {
		v = v*256 + int64(c)
	}
	if v < 0 {
		v = -v
	}
	if hi <= lo {
		return lo
	}
	return lo + v%(hi-lo)
}

func cmdMeasureCommand(c *Context) ([]*object.PSObject, error) {
	node := c.Args.GetNode("Expression")
	if sb, ok := node.(*ast.ScriptBlock); ok {
		start := time.Now()
		_, _ = c.Engine.InvokeBlock(&ast.Block{Body: sb.Body}, nil, c.Stdout)
		return []*object.PSObject{timeSpanObj(time.Since(start))}, nil
	}
	return nil, nil
}

func cmdOutString(c *Context) ([]*object.PSObject, error) {
	var buf bytes.Buffer
	_ = object.FormatOutput(&buf, c.Input)
	text := buf.String()
	if c.Args.Switch("Stream") {
		var out []*object.PSObject
		for _, ln := range strings.Split(strings.TrimRight(text, "\n"), "\n") {
			out = append(out, object.Str(ln))
		}
		return out, nil
	}
	return []*object.PSObject{object.Str(text)}, nil
}

func cmdTeeObject(c *Context) ([]*object.PSObject, error) {
	path := firstArg(c, "FilePath")
	appendMode := c.Args.Switch("Append")
	if path != "" {
		flags := os.O_WRONLY | os.O_CREATE
		if appendMode {
			flags |= os.O_APPEND
		} else {
			flags |= os.O_TRUNC
		}
		full, derr := resolvePath(c, path)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		f, err := os.OpenFile(full, flags, 0o644)
		if err == nil {
			for _, o := range c.Input {
				fmt.Fprintln(f, o.String())
			}
			f.Close()
		}
	}
	return c.Input, nil
}

func cmdFormatHex(c *Context) ([]*object.PSObject, error) {
	var sb strings.Builder
	for _, o := range inputItems(c) {
		data := []byte(o.String())
		// 16 字节一行
		for i := 0; i < len(data); i += 16 {
			end := i + 16
			if end > len(data) {
				end = len(data)
			}
			fmt.Fprintf(&sb, "%08x  ", i)
			for j := i; j < end; j++ {
				fmt.Fprintf(&sb, "%02x ", data[j])
			}
			sb.WriteString("\n")
		}
	}
	return []*object.PSObject{object.Str(strings.TrimRight(sb.String(), "\n"))}, nil
}

func cmdJoinString(c *Context) ([]*object.PSObject, error) {
	sep := " "
	if s, ok := c.Args.Str("Separator"); ok {
		sep = s
	}
	var parts []string
	for _, o := range inputItems(c) {
		parts = append(parts, o.String())
	}
	return []*object.PSObject{object.Str(strings.Join(parts, sep))}, nil
}

func cmdAddMember(c *Context) ([]*object.PSObject, error) {
	memberType := "NoteProperty"
	if t, ok := c.Args.Str("MemberType"); ok {
		memberType = t
	}
	name, _ := c.Args.Str("Name")
	val := c.Args.Get("Value")
	force := c.Args.Switch("Force")
	var out []*object.PSObject
	for _, o := range inputItems(c) {
		cp := o.Clone()
		if !strings.EqualFold(memberType, "NoteProperty") {
			// 只支持 NoteProperty
			cp = o
		}
		if name != "" {
			if !cp.HasProp(name) || force {
				if val != nil {
					cp.AddProp(name, val.Value)
				} else {
					cp.AddProp(name, nil)
				}
			}
		}
		out = append(out, cp)
	}
	return out, nil
}

func cmdNewGuid(c *Context) ([]*object.PSObject, error) {
	var b [16]byte
	_, _ = rand.Read(b[:])
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	guid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
	return []*object.PSObject{object.Str(guid)}, nil
}

func cmdNewTimeSpan(c *Context) ([]*object.PSObject, error) {
	d := time.Duration(0)
	if sec, ok := c.Args.Int("Seconds"); ok {
		d += time.Duration(sec) * time.Second
	}
	if mn, ok := c.Args.Int("Minutes"); ok {
		d += time.Duration(mn) * time.Minute
	}
	if h, ok := c.Args.Int("Hours"); ok {
		d += time.Duration(h) * time.Hour
	}
	return []*object.PSObject{timeSpanObj(d)}, nil
}

func cmdNewTemporaryFile(c *Context) ([]*object.PSObject, error) {
	ext := ""
	if e, ok := c.Args.Str("Extension"); ok {
		ext = e
	}
	f, err := os.CreateTemp("", "tmp*"+ext)
	if err != nil {
		return errf(c, "New-TemporaryFile : %v", err)
	}
	f.Close()
	if info, e := os.Stat(f.Name()); e == nil {
		return []*object.PSObject{object.FileInfo(f.Name(), info)}, nil
	}
	return []*object.PSObject{object.Str(f.Name())}, nil
}

func cmdTestJson(c *Context) ([]*object.PSObject, error) {
	var text string
	if v := c.Args.Get("Json"); v != nil {
		text = v.String()
	} else if len(c.Input) > 0 {
		// 命名/位置都没给时才用管道输入
		var sb strings.Builder
		for _, o := range c.Input {
			sb.WriteString(o.String())
		}
		text = sb.String()
	}
	var v any
	err := json.Unmarshal([]byte(text), &v)
	return []*object.PSObject{object.Bool(err == nil)}, nil
}

// ---- 注册 ----

func init() {
	Register("ConvertTo-Csv", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
		{Name: "Property", Type: "string[]"},
	}, cmdConvertToCsv)
	Register("ConvertFrom-Csv", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
	}, cmdConvertFromCsv)
	Register("ConvertTo-Json", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
		{Name: "Depth", Type: "int"},
	}, cmdConvertToJson)
	Register("ConvertFrom-Json", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
	}, cmdConvertFromJson)
	Register("ConvertFrom-StringData", []ParamSpec{
		{Name: "StringData", Position: 0, PositionSet: true, Type: "string"},
	}, cmdConvertFromStringData)
	Register("Compare-Object", []ParamSpec{
		{Name: "ReferenceObject", Position: 0, PositionSet: true, Type: "object"},
		{Name: "DifferenceObject", Position: 1, PositionSet: true, Type: "object"},
		{Name: "CaseSensitive", Switch: true},
		{Name: "IncludeEqual", Switch: true},
	}, cmdCompareObject)
	Register("Get-Unique", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
		{Name: "AsString", Switch: true},
	}, cmdGetUnique)
	Register("Get-Random", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
		{Name: "Minimum", Type: "int"},
		{Name: "Maximum", Type: "int"},
		{Name: "Count", Type: "int"},
	}, cmdGetRandom)
	Register("Measure-Command", []ParamSpec{
		{Name: "Expression", Position: 0, PositionSet: true, Type: "scriptblock"},
	}, cmdMeasureCommand)
	Register("Out-String", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
		{Name: "Stream", Switch: true},
	}, cmdOutString)
	Register("Tee-Object", []ParamSpec{
		{Name: "FilePath", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Append", Switch: true},
	}, cmdTeeObject)
	Register("Format-Hex", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
	}, cmdFormatHex)
	Register("Join-String", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
		{Name: "Separator", Type: "string"},
	}, cmdJoinString)
	Register("Add-Member", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
		{Name: "MemberType", Type: "string"},
		{Name: "Name", Type: "string"},
		{Name: "Value", Type: "object"},
		{Name: "Force", Switch: true},
	}, cmdAddMember)
	Register("New-Guid", nil, cmdNewGuid)
	Register("New-TimeSpan", []ParamSpec{
		{Name: "Seconds", Type: "int"},
		{Name: "Minutes", Type: "int"},
		{Name: "Hours", Type: "int"},
	}, cmdNewTimeSpan)
	Register("New-TemporaryFile", []ParamSpec{
		{Name: "Extension", Type: "string"},
	}, cmdNewTemporaryFile)
	Register("Test-Json", []ParamSpec{
		{Name: "Json", Position: 0, PositionSet: true, Type: "string"},
	}, cmdTestJson)
}

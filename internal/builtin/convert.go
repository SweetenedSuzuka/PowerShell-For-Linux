package builtin

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"powershell/internal/lang"
	"powershell/internal/object"
)

// convert.go 实现数据转换类 cmdlet（CSV/JSON/字符串数据互转与 JSON 校验）。

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
	var out []*object.PSObject
	for _, ln := range csvLines(items, c.Args.StringSlice("Property"), ',') {
		out = append(out, object.Str(ln))
	}
	return out, nil
}

// csvLines 按给定分隔符把对象列表转换为 CSV 文本行（首行为表头），ConvertTo-Csv 与 Export-Csv 共用。
func csvLines(items []*object.PSObject, props []string, delim rune) []string {
	if len(items) == 0 {
		return nil
	}
	var header []string
	if len(props) > 0 {
		header = props
	} else {
		header = csvHeader(items[0])
		if len(header) == 0 {
			header = []string{"Value"}
		}
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	w.Comma = delim
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
	return strings.Split(text, "\n")
}

// cmdExportCsv 把对象写成 CSV 文件（文本格式与 ConvertTo-Csv 一致）。
func cmdExportCsv(c *Context) ([]*object.PSObject, error) {
	path := firstPathArg(c)
	if path == "" {
		return nil, nil
	}
	// 分隔符：默认逗号，只接受单个字符；回车与换行不能作分隔符。
	delim := ','
	if d, _ := c.Args.Str("Delimiter"); d != "" {
		r := []rune(d)
		if len(r) != 1 || r[0] == '\r' || r[0] == '\n' {
			return errf(c, "%s", lang.T(lang.MsgConvertFail, d, "char"))
		}
		delim = r[0]
	}
	items := inputItems(c)
	if len(items) == 0 {
		return nil, nil
	}
	var rows []*object.PSObject
	for _, it := range items {
		if it == nil || it.IsNull() {
			continue
		}
		rows = append(rows, it)
	}
	lines := csvLines(rows, c.Args.StringSlice("Property"), delim)
	full, derr := resolvePath(c, path)
	if derr != nil {
		return errf(c, "%v", derr)
	}
	appendMode := c.Args.Switch("Append")
	exists := false
	nonEmpty := false
	if fi, serr := os.Stat(full); serr == nil {
		exists = true
		nonEmpty = fi.Size() > 0
		if c.Args.Switch("NoClobber") {
			return errf(c, "%s", lang.T(lang.MsgFileExists, path))
		}
	}
	var dryRun whatIfCollector
	dryRun.cmdlet = "Export-Csv"
	dryRun.c = c
	var yesAll, noAll bool
	if dryRun.reportWhatIf(full) {
		out, _ := dryRun.result()
		return out, nil
	}
	if confirmSkip(c, "Export-Csv", full, &yesAll, &noAll) {
		return nil, nil
	}
	// 追加到已有非空文件时只写数据行，不重复表头；其余情况写完整文本。
	payload := lines
	if appendMode && nonEmpty && len(lines) > 0 {
		payload = lines[1:]
	}
	text := ""
	if len(payload) > 0 {
		text = strings.Join(payload, "\n") + "\n"
	}
	enc, _ := c.Args.Str("Encoding")
	data := encodeText(enc, text, !nonEmpty)
	if appendMode && exists {
		f, err := os.OpenFile(full, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
		if err != nil {
			return errf(c, "%s", lang.T(lang.MsgCannotOpen, path))
		}
		defer f.Close()
		if _, err := f.Write(data); err != nil {
			return errf(c, "%s", lang.T(lang.MsgCannotWrite, path))
		}
		return nil, nil
	}
	if err := os.WriteFile(full, data, 0o644); err != nil {
		return errf(c, "%s", lang.T(lang.MsgCannotWrite, path))
	}
	return nil, nil
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
	return csvRowsToObjects(records[1:], records[0]), nil
}

// csvRowsToObjects 按表头把 CSV 记录转成表格对象（列缺失不建属性，多余列忽略）。
func csvRowsToObjects(records [][]string, header []string) []*object.PSObject {
	var out []*object.PSObject
	for _, rec := range records {
		o := object.Object("System.Management.Automation.PSCustomObject", nil)
		for i, h := range header {
			if i < len(rec) {
				o.AddProp(h, rec[i])
			}
		}
		out = append(out, o)
	}
	return out
}

// cmdImportCsv 从 CSV 文件读出表格对象（与 ConvertFrom-Csv 共用记录转换，#TYPE 首行跳过）。
func cmdImportCsv(c *Context) ([]*object.PSObject, error) {
	// 超量位置实参无槽位可接（Path 占位置 0，Delimiter 占位置 1），报错而非静默忽略。
	if len(c.Args.Positional) > 0 {
		return errf(c, "%s", lang.T(lang.MsgPositionalParamNotFound, c.Args.Positional[0].String()))
	}
	var paths []string
	if v := c.Args.Get("Path"); v != nil {
		for _, it := range v.ArrayItems() {
			paths = append(paths, it.String())
		}
	}
	if len(paths) == 0 {
		return nil, nil
	}
	// 分隔符：默认逗号，只接受单个字符；回车与换行不能作分隔符。
	delim := ','
	if d, _ := c.Args.Str("Delimiter"); d != "" {
		r := []rune(d)
		if len(r) != 1 || r[0] == '\r' || r[0] == '\n' {
			return errf(c, "%s", lang.T(lang.MsgConvertFail, d, "char"))
		}
		delim = r[0]
	}
	header := c.Args.StringSlice("Header")
	var out []*object.PSObject
	for _, path := range paths {
		full, derr := resolvePath(c, path)
		if derr != nil {
			return errf(c, "%v", derr)
		}
		data, err := os.ReadFile(full)
		if err != nil {
			return errf(c, "%s", lang.T(lang.MsgPathNotFoundFmt, path))
		}
		r := csv.NewReader(strings.NewReader(StripUTF8BOM(string(data))))
		r.Comma = delim
		r.FieldsPerRecord = -1
		rows, err := r.ReadAll()
		if err != nil {
			return errf(c, "Import-Csv : %v", err)
		}
		// 首行类型行不是数据，去掉后下一行才是表头（-Header 给出时剩下全是数据）。
		if len(rows) > 0 && len(rows[0]) > 0 && strings.HasPrefix(rows[0][0], "#TYPE") {
			rows = rows[1:]
		}
		if len(header) > 0 {
			out = append(out, csvRowsToObjects(rows, header)...)
			continue
		}
		if len(rows) < 2 {
			continue
		}
		out = append(out, csvRowsToObjects(rows[1:], rows[0])...)
	}
	if len(out) == 0 {
		return nil, nil
	}
	return out, nil
}

func cmdConvertToJson(c *Context) ([]*object.PSObject, error) {
	items := inputItems(c)
	// -Depth 限制嵌套展开深度（默认 2）
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
// remaining 为剩余可展开深度；到 0 时数组/对象不再展开。
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
	Register("Import-Csv", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "Delimiter", Position: 1, PositionSet: true, Type: "string"},
		{Name: "Header", Type: "string[]"},
	}, cmdImportCsv)
	Register("Export-Csv", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "InputObject", Type: "object"},
		{Name: "Property", Type: "string[]"},
		{Name: "Delimiter", Type: "string"},
		{Name: "Append", Switch: true},
		{Name: "NoClobber", Switch: true},
		{Name: "NoTypeInformation", Switch: true},
		{Name: "Encoding", Type: "string"},
	}, cmdExportCsv)
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
	Register("Test-Json", []ParamSpec{
		{Name: "Json", Position: 0, PositionSet: true, Type: "string"},
	}, cmdTestJson)
}

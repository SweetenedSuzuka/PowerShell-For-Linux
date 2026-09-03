package builtin

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

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

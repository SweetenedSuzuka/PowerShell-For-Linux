package object

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// formatShape 描述对象的渲染形状。
type formatShape int

const (
	shapeScalar      formatShape = iota // 标量：每行一个
	shapeTable                          // 表格（带 Table 定义或哈希表）
	shapeCustomTable                    // 自定义对象表格：与其它表格形状分组渲染，避免异构错列
	shapeList                           // 列表（Name : value）
)

func shapeOf(o *PSObject) formatShape {
	if o.TypeName == "Hashtable" || len(o.Table) > 0 {
		return shapeTable
	}
	// 自定义对象按表格渲染。
	// 连续同形对象合成一张表，首对象的属性作列。
	if o.TypeName == "System.Management.Automation.PSCustomObject" {
		return shapeCustomTable
	}
	if isScalarType(o) {
		return shapeScalar
	}
	return shapeList
}

// FormatOutput 渲染对象流：$null 不占位（与 PowerShell 一致）；按形状分组——标量每行一个、表格形状出表格、对象出列表。
func FormatOutput(w io.Writer, objs []*PSObject) error {
	objs = dropNulls(flatten(objs))
	if len(objs) == 0 {
		return nil
	}
	i := 0
	for i < len(objs) {
		shape := shapeOf(objs[i])
		j := i
		for j < len(objs) && shapeOf(objs[j]) == shape {
			j++
		}
		group := objs[i:j]
		switch shape {
		case shapeScalar:
			if err := writeStrings(w, group); err != nil {
				return err
			}
		case shapeTable, shapeCustomTable:
			// 单个自定义对象使用列表处理（与 PowerShell 一致），多个才合成表格
			if shape == shapeCustomTable && len(group) == 1 {
				if err := FormatListTo(w, group, nil); err != nil {
					return err
				}
				break
			}
			if err := FormatTableTo(w, group, nil); err != nil {
				return err
			}
		case shapeList:
			if err := FormatListTo(w, group, nil); err != nil {
				return err
			}
		}
		i = j
	}
	return nil
}

// dropNulls 丢弃空对象（$null 渲染时不占位）；结果是新的一组，不动入参。
func dropNulls(objs []*PSObject) []*PSObject {
	var out []*PSObject
	for _, o := range objs {
		if o == nil || o.IsNull() {
			continue
		}
		out = append(out, o)
	}
	return out
}

// flatten 把数组对象展开为元素列表（数组内嵌数组时递归）。
func flatten(objs []*PSObject) []*PSObject {
	var out []*PSObject
	for _, o := range objs {
		if items, ok := o.Value.([]*PSObject); ok && o.TypeName == "Object[]" {
			out = append(out, flatten(items)...)
		} else {
			out = append(out, o)
		}
	}
	return out
}

func writeStrings(w io.Writer, objs []*PSObject) error {
	for _, o := range objs {
		if _, err := fmt.Fprintln(w, o.String()); err != nil {
			return err
		}
	}
	return nil
}

// tableScalar 报告对象在表格中是否按标量直出：无可制表属性，-Property 忽略（DateTime 有属性，走表格）。
func tableScalar(o *PSObject) bool {
	switch o.TypeName {
	case "String", "Int", "Double", "Boolean", "Null", "ScriptBlock":
		return true
	}
	return false
}

// cellOf 取对象某列的显示值；哈希表特殊处理为 Name/Value 两列。
func cellOf(o *PSObject, label string) string {
	if o.TypeName == "Hashtable" {
		entries, _ := o.Value.([]HashEntry)
		switch label {
		case "Name":
			names := make([]string, 0, len(entries))
			for _, e := range entries {
				names = append(names, e.Key)
			}
			return strings.Join(names, ", ")
		case "Value":
			vals := make([]string, 0, len(entries))
			for _, e := range entries {
				vals = append(vals, e.Value.String())
			}
			return strings.Join(vals, ", ")
		}
	}
	if o.IsNull() {
		return ""
	}
	if v, ok := o.PropValue(label); ok {
		return v.String()
	}
	return ""
}

// dateTimeColumns 是 DateTime 默认表格列（顺序与 PowerShell 一致，不含 DateTime 成员列）。
var dateTimeColumns = []string{"DisplayHint", "Date", "Day", "DayOfWeek", "DayOfYear", "Hour", "Kind", "Millisecond", "Microsecond", "Nanosecond", "Minute", "Month", "Second", "Ticks", "TimeOfDay", "Year"}

// DateTimeSelectColumns 是 DateTime 全属性列（Select * 与 Format-List * 用，顺序与 PowerShell 一致）。
var DateTimeSelectColumns = []string{"DisplayHint", "DateTime", "Date", "Day", "DayOfWeek", "DayOfYear", "Hour", "Kind", "Millisecond", "Microsecond", "Nanosecond", "Minute", "Month", "Second", "Ticks", "TimeOfDay", "Year"}

// tableColumns 决定表格的列定义（标签 + 对齐）。
func tableColumns(objs []*PSObject) (labels []string, aligns []string) {
	for _, o := range objs {
		if o.TypeName == "Hashtable" {
			return []string{"Name", "Value"}, []string{"left", "left"}
		}
		if len(o.Table) > 0 {
			labels = make([]string, 0, len(o.Table))
			aligns = make([]string, 0, len(o.Table))
			for _, c := range o.Table {
				labels = append(labels, c.Label)
				aligns = append(aligns, c.Align)
			}
			return labels, aligns
		}
	}
	// DateTime 用默认列（虚拟属性经 PropValue 取值，顺序见 dateTimeColumns）。
	if len(objs) > 0 && objs[0].TypeName == "DateTime" {
		labels := make([]string, len(dateTimeColumns))
		copy(labels, dateTimeColumns)
		aligns := make([]string, len(dateTimeColumns))
		for i := range aligns {
			aligns[i] = "left"
		}
		return labels, aligns
	}
	// 对象没有表格列定义时用第一个对象的属性名作为列（如 Select-Object 的自定义对象）
	if len(objs) > 0 && len(objs[0].Props) > 0 {
		labels = make([]string, 0, len(objs[0].Props))
		aligns = make([]string, 0, len(objs[0].Props))
		for _, p := range objs[0].Props {
			labels = append(labels, p.Name)
			aligns = append(aligns, "left")
		}
		return labels, aligns
	}
	return []string{"Value"}, []string{"left"}
}

// TerminalWidth 返回输出宽度：终端→COLUMNS→默认 80。
func TerminalWidth() int {
	if c := terminalColumns(); c > 0 {
		return c
	}
	if s := os.Getenv("COLUMNS"); s != "" {
		if n, err := strconv.Atoi(strings.TrimSpace(s)); err == nil && n > 0 {
			return n
		}
	}
	return 80
}

// fitWidths 把列宽收进终端宽度：从右向左收窄到表头宽度为止。
func fitWidths(widths []int, labels []string, max int) {
	total := 2 * (len(widths) - 1)
	for _, wd := range widths {
		total += wd
	}
	for i := len(widths) - 1; i >= 0 && total > max; i-- {
		floor := displayWidth(labels[i])
		if widths[i] > floor {
			cut := widths[i] - floor
			if cut > total-max {
				cut = total - max
			}
			widths[i] -= cut
			total -= cut
		}
	}
}

// truncateDisplay 按显示宽度截断（中文等宽按 2 计）。
func truncateDisplay(s string, n int) string {
	if displayWidth(s) <= n {
		return s
	}
	w := 0
	for i, r := range s {
		rw := 1
		if r > 0x2E7F {
			rw = 2
		}
		if w+rw > n {
			return s[:i]
		}
		w += rw
	}
	return s
}

// FormatTableTo 以表格形式渲染对象；标量穿插直出（顺序：先落攒的表，再出标量行）。
func FormatTableTo(w io.Writer, objs []*PSObject, props []string) error {
	if len(objs) == 0 {
		return nil
	}
	var buf []*PSObject
	flush := func() error {
		if len(buf) == 0 {
			return nil
		}
		if err := writeTable(w, buf, props); err != nil {
			return err
		}
		buf = nil
		return nil
	}
	for _, o := range objs {
		if o == nil || o.IsNull() {
			continue
		}
		if tableScalar(o) {
			if err := flush(); err != nil {
				return err
			}
			if _, err := fmt.Fprintln(w, o.String()); err != nil {
				return err
			}
			continue
		}
		buf = append(buf, o)
	}
	return flush()
}

// writeTable 渲染一段非标量对象为一张表。
func writeTable(w io.Writer, objs []*PSObject, props []string) error {
	labels, aligns := tableColumns(objs)
	if len(props) > 0 {
		labels = props
		aligns = make([]string, len(props))
		for i := range aligns {
			aligns[i] = "left"
		}
	}
	// 计算各列宽度
	widths := make([]int, len(labels))
	for i, l := range labels {
		widths[i] = displayWidth(l)
	}
	rows := make([][]string, 0, len(objs))
	for _, o := range objs {
		if o.TypeName == "Hashtable" {
			// 哈希表：每个键值对占一行
			if entries, ok := o.Value.([]HashEntry); ok {
				for _, en := range entries {
					row := make([]string, len(labels))
					for i, l := range labels {
						switch l {
						case "Name":
							row[i] = en.Key
						case "Value":
							row[i] = en.Value.String()
						default:
							row[i] = ""
						}
						if wd := displayWidth(row[i]); wd > widths[i] {
							widths[i] = wd
						}
					}
					rows = append(rows, row)
				}
				continue
			}
		}
		row := make([]string, len(labels))
		for i, l := range labels {
			row[i] = cellOf(o, l)
			if wd := displayWidth(row[i]); wd > widths[i] {
				widths[i] = wd
			}
		}
		rows = append(rows, row)
	}
	// 总宽超限时收进终端宽度。
	fitWidths(widths, labels, TerminalWidth())
	// 表头
	writeRow(w, labels, widths, aligns)
	// 下划线
	dashes := make([]string, len(labels))
	for i := range labels {
		dashes[i] = strings.Repeat("-", widths[i])
	}
	writeRow(w, dashes, widths, aligns)
	// 数据行
	for _, row := range rows {
		writeRow(w, row, widths, aligns)
	}
	return nil
}

func writeRow(w io.Writer, cells []string, widths []int, aligns []string) {
	var sb strings.Builder
	for i, c := range cells {
		c = truncateDisplay(c, widths[i])
		pad := widths[i] - displayWidth(c)
		if pad < 0 {
			pad = 0
		}
		if i < len(aligns) && aligns[i] == "right" {
			sb.WriteString(strings.Repeat(" ", pad))
			sb.WriteString(c)
		} else {
			sb.WriteString(c)
			sb.WriteString(strings.Repeat(" ", pad))
		}
		if i < len(cells)-1 {
			sb.WriteString("  ")
		}
	}
	fmt.Fprintln(w, strings.TrimRight(sb.String(), " "))
}

// displayWidth 计算显示宽度（中文等宽字符按 2 计）。
func displayWidth(s string) int {
	w := 0
	for _, r := range s {
		if r > 0x2E7F { // 粗略：CJK 等宽按 2
			w += 2
		} else {
			w++
		}
	}
	return w
}

// FormatListTo 以列表形式渲染对象（Name : value）。
func FormatListTo(w io.Writer, objs []*PSObject, props []string) error {
	if len(objs) == 0 {
		return nil
	}
	first := true
	for _, o := range objs {
		if o == nil || o.IsNull() {
			continue
		}
		if !first {
			fmt.Fprintln(w)
		}
		first = false
		if isScalarType(o) && len(props) == 0 {
			fmt.Fprintln(w, o.String())
			continue
		}
		// 对齐属性名
		allProps := len(props) == 1 && props[0] == "*" // fl * 表示全部属性
		var names []string
		if len(props) > 0 && !allProps {
			names = props
		} else {
			for _, p := range o.Props {
				names = append(names, p.Name)
			}
			// DateTime 虚拟属性补齐（顺序见 DateTimeSelectColumns；其它类型沿用实属性）。
			if o.TypeName == "DateTime" {
				for _, vn := range DateTimeSelectColumns {
					dup := false
					for _, n := range names {
						if strings.EqualFold(n, vn) {
							dup = true
							break
						}
					}
					if !dup {
						names = append(names, vn)
					}
				}
			}
			if len(names) == 0 {
				fmt.Fprintln(w, o.String())
				continue
			}
		}
		max := 0
		for _, n := range names {
			if wd := displayWidth(n); wd > max {
				max = wd
			}
		}
		for _, n := range names {
			val := ""
			if len(props) > 0 && !allProps {
				if v, ok := o.PropValue(n); ok {
					val = v.String()
				}
			} else if o.TypeName == "DateTime" {
				// 虚拟属性经 PropValue 取值；实属性 PropValue 优先查 Props，结果一致。
				if v, ok := o.PropValue(n); ok {
					val = v.String()
				}
			} else {
				for _, p := range o.Props {
					if strings.EqualFold(p.Name, n) {
						val = ToPS(p.Value).String()
						break
					}
				}
			}
			fmt.Fprintf(w, "%s%s : %s\n", n, strings.Repeat(" ", max-displayWidth(n)), val)
		}
	}
	return nil
}

// FormatWideTo 以宽幅（多列）形式渲染字符串；prop 非空时取对象该属性显示。
func FormatWideTo(w io.Writer, objs []*PSObject, colWidth int, prop string) error {
	if colWidth <= 0 {
		colWidth = 40
	}
	var cells []string
	for _, o := range objs {
		if o == nil || o.IsNull() {
			continue
		}
		if prop != "" {
			if v, ok := o.PropValue(prop); ok {
				cells = append(cells, v.String())
			} else {
				cells = append(cells, "")
			}
		} else if len(o.Props) > 0 {
			// 无属性指定：优先显示第一个属性（FileInfo/PSCustomObject 显示 Name）
			cells = append(cells, ToPS(o.Props[0].Value).String())
		} else {
			cells = append(cells, o.String())
		}
	}
	// 确定列数（依据终端宽度）。
	cols := TerminalWidth() / colWidth
	if cols < 1 {
		cols = 1
	}
	if cols > len(cells) {
		cols = len(cells)
	}
	rows := (len(cells) + cols - 1) / cols
	for r := 0; r < rows; r++ {
		var sb strings.Builder
		for c := 0; c < cols; c++ {
			idx := c*rows + r
			if idx >= len(cells) {
				continue
			}
			sb.WriteString(cells[idx])
			if c < cols-1 {
				if pad := colWidth - displayWidth(cells[idx]); pad > 0 {
					sb.WriteString(strings.Repeat(" ", pad))
				}
			}
		}
		fmt.Fprintln(w, strings.TrimRight(sb.String(), " "))
	}
	return nil
}

func isScalarType(o *PSObject) bool {
	switch o.TypeName {
	case "String", "Int", "Double", "Boolean", "Null", "DateTime", "ScriptBlock":
		return true
	}
	return false
}

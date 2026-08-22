package builtin

import (
	"sort"
	"strings"

	"powershell/internal/ast"
	"powershell/internal/object"
)

// ---- 对象管道 ----

func filterMatches(c *Context, obj *object.PSObject) bool {
	// -FilterScript（命名或位置）：脚本块或比较表达式
	if node := c.Args.GetNode("FilterScript"); node != nil {
		return evalFilterNode(c, node, obj)
	}
	// -Property Length -gt 100 → Property 的值是一个比较表达式
	if node := c.Args.GetNode("Property"); node != nil {
		if _, isBare := node.(*ast.BareWord); isBare {
			return false // 缺比较运算符
		}
		return evalFilterNode(c, node, obj)
	}
	// 无参数：输入对象本身作为过滤器（输出为真者）
	return obj.Truthy()
}

func evalFilterNode(c *Context, node ast.Node, obj *object.PSObject) bool {
	if sb, ok := node.(*ast.ScriptBlock); ok {
		outs, _ := c.Engine.InvokeBlock(&ast.Block{Body: sb.Body}, map[string]*object.PSObject{"_": obj, "PSItem": obj}, c.Stdout)
		if len(outs) == 0 {
			return false
		}
		return outs[len(outs)-1].Truthy()
	}
	ok, err := c.Engine.EvalFilterExpr(node, obj)
	if err != nil {
		return false
	}
	return ok
}

func cmdWhereObject(c *Context) ([]*object.PSObject, error) {
	not := c.Args.Switch("Not")
	var out []*object.PSObject
	for _, obj := range c.Input {
		ok := filterMatches(c, obj)
		if not {
			ok = !ok
		}
		if ok {
			out = append(out, obj)
		}
	}
	return out, nil
}

func cmdSelectObject(c *Context) ([]*object.PSObject, error) {
	first, firstSet := c.Args.Int("First")
	last, lastSet := c.Args.Int("Last")
	unique := c.Args.Switch("Unique")
	expand, _ := c.Args.Str("ExpandProperty")

	items := c.Input
	props := c.Args.StringSlice("Property")
	if len(items) == 0 && c.Args.PosMapped["Property"] {
		// 无管道输入且属性来自位置映射：位置实参按数据原样输出（本项目语义，如 Select-Object a,b），不做属性选择（props 清空）。
		if v := c.Args.Get("Property"); v != nil {
			items = append(items, v.ArrayItems()...)
		}
		// 超量位置实参（第 2 个起的实参）也当数据
		items = append(items, c.Args.Positional...)
		props = nil
	}
	// -First/-Last 显式 0 时返回空（真 PowerShell 语义），未设置则不动
	if firstSet {
		if first <= 0 {
			items = nil
		} else if int(first) < len(items) {
			items = items[:first]
		}
	}
	if lastSet {
		if last <= 0 {
			items = nil
		} else if int(last) < len(items) {
			items = items[len(items)-int(last):]
		}
	}
	if unique {
		seen := map[string]bool{}
		var uniq []*object.PSObject
		for _, it := range items {
			if !seen[it.String()] {
				seen[it.String()] = true
				uniq = append(uniq, it)
			}
		}
		items = uniq
	}
	if len(props) > 0 {
		// * 表示选全部属性
		all := false
		for _, p := range props {
			if p == "*" {
				all = true
				break
			}
		}
		var out []*object.PSObject
		for _, it := range items {
			names := props
			if all {
				names = nil
				for _, pr := range it.Props {
					names = append(names, pr.Name)
				}
				if len(names) == 0 {
					out = append(out, it) // 无属性对象（标量等）原样保留
					continue
				}
			}
			// 只留选中的属性：生成 PSCustomObject（不保留原类型，避免按标量渲染漏掉属性）
			n := object.Object("System.Management.Automation.PSCustomObject", nil)
			for _, p := range names {
				if v, ok := it.PropValue(p); ok {
					n.AddProp(p, v.Value)
				} else {
					n.AddProp(p, nil)
				}
			}
			out = append(out, n)
		}
		return out, nil
	}
	// -ExpandProperty：取属性值本身输出（数组摊平），不做对象包装
	if expand != "" {
		var out []*object.PSObject
		for _, it := range items {
			if v, ok := it.PropValue(expand); ok {
				for _, e := range v.ArrayItems() {
					out = append(out, e)
				}
			}
		}
		return out, nil
	}
	return items, nil
}

func cmdSortObject(c *Context) ([]*object.PSObject, error) {
	props := c.Args.StringSlice("Property")
	desc := c.Args.Switch("Descending")
	unique := c.Args.Switch("Unique")
	caseSensitive := c.Args.Switch("CaseSensitive")

	// 取对象某属性值；缺失视为 $null（排序时排最前）。
	keyOf := func(o *object.PSObject, p string) *object.PSObject {
		if v, ok := o.PropValue(p); ok {
			return v
		}
		return object.Null()
	}
	// 多属性逐个比较：第一个非 0 结果决定顺序。
	compare := func(a, b *object.PSObject) int {
		if len(props) == 0 {
			return compareOrderBuiltin(a, b)
		}
		for _, p := range props {
			if c := compareOrderBuiltin(keyOf(a, p), keyOf(b, p)); c != 0 {
				return c
			}
		}
		return compareOrderBuiltin(a, b)
	}
	items := c.Input
	sort.SliceStable(items, func(i, j int) bool {
		ord := compare(items[i], items[j])
		if desc {
			return ord > 0
		}
		return ord < 0
	})
	if unique {
		seen := map[string]bool{}
		var uniq []*object.PSObject
		for _, it := range items {
			var sb strings.Builder
			if len(props) == 0 {
				sb.WriteString(it.String())
			} else {
				for _, p := range props {
					sb.WriteString(keyOf(it, p).String())
					sb.WriteByte(0)
				}
			}
			k := sb.String()
			// -Unique 默认按小写折叠去重，-CaseSensitive 时原样
			if !caseSensitive {
				k = strings.ToLower(k)
			}
			if !seen[k] {
				seen[k] = true
				uniq = append(uniq, it)
			}
		}
		items = uniq
	}
	return items, nil
}

func compareOrderBuiltin(a, b *object.PSObject) int {
	if an, ok := a.AsFloat(); ok {
		if bn, ok2 := b.AsFloat(); ok2 {
			if an < bn {
				return -1
			}
			if an > bn {
				return 1
			}
			return 0
		}
	}
	return strings.Compare(strings.ToLower(a.String()), strings.ToLower(b.String()))
}

func cmdForEachObject(c *Context) ([]*object.PSObject, error) {
	node := c.Args.GetNode("Process")
	var out []*object.PSObject
	run := func(n ast.Node, extra map[string]*object.PSObject) {
		if sb, ok := n.(*ast.ScriptBlock); ok {
			outs, _ := c.Engine.InvokeBlock(&ast.Block{Body: sb.Body}, extra, c.Stdout)
			out = append(out, outs...)
		}
	}
	// -Begin / -End：各执行一次（聚合写法）
	if begin := c.Args.GetNode("Begin"); begin != nil {
		run(begin, nil)
	}
	// -MemberName：对每个对象取该成员（如 ForEach-Object -MemberName Length）
	if mn, ok := c.Args.Str("MemberName"); ok && mn != "" {
		for _, obj := range c.Input {
			if v, ok := obj.PropValue(mn); ok {
				out = append(out, v)
			}
		}
	} else if sb, ok := node.(*ast.ScriptBlock); ok {
		for _, obj := range c.Input {
			outs, _ := c.Engine.InvokeBlock(&ast.Block{Body: sb.Body}, map[string]*object.PSObject{"_": obj, "PSItem": obj}, c.Stdout)
			out = append(out, outs...)
		}
	} else {
		out = append(out, c.Input...)
	}
	if end := c.Args.GetNode("End"); end != nil {
		run(end, nil)
	}
	return out, nil
}

func cmdGroupObject(c *Context) ([]*object.PSObject, error) {
	prop, _ := c.Args.Str("Property")
	caseSensitive := c.Args.Switch("CaseSensitive")
	groups := map[string][]*object.PSObject{}
	firstName := map[string]string{}
	var order []string
	for _, o := range c.Input {
		k := ""
		if prop != "" {
			if v, ok := o.PropValue(prop); ok {
				k = v.String()
			}
		} else {
			k = o.String()
		}
		// map key：默认按小写折叠（大小写不敏感），-CaseSensitive 时原样
		mk := k
		if !caseSensitive {
			mk = strings.ToLower(k)
		}
		if _, ok := groups[mk]; !ok {
			order = append(order, mk)
			firstName[mk] = k
		}
		groups[mk] = append(groups[mk], o)
	}
	var out []*object.PSObject
	for _, k := range order {
		// 组成员放 Group 属性（对应 PowerShell 的 GroupInfo.Group 属性），Value 保持空以免被当成集合。
		g := object.Object("GroupInfo", nil)
		g.AddProp("Group", groups[k])
		g.AddProp("Name", firstName[k])
		g.AddProp("Count", int64(len(groups[k])))
		g.Table = []object.Column{
			{Label: "Count", Align: "right"},
			{Label: "Name", Align: "left"},
		}
		out = append(out, g)
	}
	return out, nil
}

func cmdMeasureObject(c *Context) ([]*object.PSObject, error) {
	prop, _ := c.Args.Str("Property")
	lineFlag := c.Args.Switch("Line")
	sumFlag := c.Args.Switch("Sum")
	avgFlag := c.Args.Switch("Average")
	minFlag := c.Args.Switch("Minimum")
	maxFlag := c.Args.Switch("Maximum")

	items := c.Input
	// -Line：数输入里的行数（对应 wc -l）
	if lineFlag {
		var lines int64
		for _, o := range items {
			s := o.String()
			if s == "" {
				continue
			}
			lines += int64(strings.Count(s, "\n")) + 1
		}
		m := object.Object("MeasureInfo", nil)
		m.AddProp("Count", lines)
		return []*object.PSObject{m}, nil
	}
	var sum, avg, mn, mx float64
	var haveMin, haveMax bool
	var nums []float64
	// Count 按 -Property 过滤：只数能取到该属性的对象（无 -Property 时数全部）
	matchedCount := int64(0)
	// Sum/Average 遇非数字输入作废（对齐真 PowerShell：报错且字段置 $null）；Min/Max 仅统计数字
	sumAvgValid := true
	for _, o := range items {
		var v *object.PSObject
		if prop != "" {
			if pv, ok := o.PropValue(prop); ok {
				v = pv
			} else {
				continue // 对象没有该属性：不计入 Count
			}
		} else {
			v = o
		}
		matchedCount++
		if v == nil {
			continue // 属性值为 $null：计入 Count 但不参与统计
		}
		if n, ok := v.AsFloat(); ok {
			nums = append(nums, n)
			sum += n
			if !haveMin || n < mn {
				mn = n
				haveMin = true
			}
			if !haveMax || n > mx {
				mx = n
				haveMax = true
			}
		} else if sumFlag || avgFlag {
			// Sum/Average 开启后遇非数字：累加类统计作废
			sumAvgValid = false
		}
	}
	if len(nums) > 0 {
		avg = sum / float64(len(nums))
	}
	m := object.Object("MeasureInfo", nil)
	// 字段按真 PowerShell 顺序补全：Count 总有，统计字段未开或无数据或遇非数字时为 $null
	var sumVal, avgVal, minVal, maxVal any
	if sumFlag && sumAvgValid && len(nums) > 0 {
		sumVal = sum
	}
	if avgFlag && sumAvgValid && len(nums) > 0 {
		avgVal = avg
	}
	if minFlag && haveMin {
		minVal = mn
	}
	if maxFlag && haveMax {
		maxVal = mx
	}
	m.AddProp("Count", matchedCount)
	m.AddProp("Average", avgVal)
	m.AddProp("Sum", sumVal)
	m.AddProp("Maximum", maxVal)
	m.AddProp("Minimum", minVal)
	m.AddProp("StandardDeviation", nil)
	m.AddProp("Property", nil)
	return []*object.PSObject{m}, nil
}

func cmdGetMember(c *Context) ([]*object.PSObject, error) {
	// -MemberType 过滤成员类型（Property/TypeName 等，大小写不敏感，缺省返回全部）
	mt, _ := c.Args.Str("MemberType")
	typeMatch := func(t string) bool {
		return mt == "" || strings.EqualFold(t, mt)
	}
	seen := map[string]bool{}
	var out []*object.PSObject
	for _, o := range inputItems(c) {
		if !seen["type:"+o.TypeName] && typeMatch("TypeName") {
			seen["type:"+o.TypeName] = true
			t := object.Object("PSMemberInfo", nil)
			t.AddProp("Name", o.TypeName)
			t.AddProp("MemberType", "TypeName")
			t.AddProp("Definition", "")
			out = append(out, t)
		}
		for _, p := range o.Props {
			if seen[p.Name] || !typeMatch("Property") {
				continue
			}
			seen[p.Name] = true
			pm := object.Object("PSMemberInfo", nil)
			pm.AddProp("Name", p.Name)
			pm.AddProp("MemberType", "Property")
			pm.AddProp("Definition", object.ToPS(p.Value).String())
			out = append(out, pm)
		}
	}
	return out, nil
}

// cmdNewObject 构造对象：目前支持 PSObject/PSCustomObject（-Property 哈希表变属性）。
func cmdNewObject(c *Context) ([]*object.PSObject, error) {
	tn, _ := c.Args.Str("TypeName")
	if tn == "" {
		return nil, nil
	}
	switch strings.ToLower(tn) {
	case "psobject", "pscustomobject":
		var entries []object.HashEntry
		if p := c.Args.Get("Property"); p != nil && p.TypeName == "Hashtable" {
			if es, ok := p.Value.([]object.HashEntry); ok {
				entries = es
			}
		}
		return []*object.PSObject{object.PSCustomObject(entries)}, nil
	default:
		return errf(c, "New-Object : 不支持的类型 %s。", tn)
	}
}

// ---- 注册 ----

func init() {
	Register("Where-Object", []ParamSpec{
		{Name: "FilterScript", Position: 0, PositionSet: true, Type: "scriptblock"},
		{Name: "Property", Type: "string"},
		{Name: "Not", Switch: true},
		{Name: "SimpleMatch", Switch: true},
	}, cmdWhereObject)
	Register("Select-Object", []ParamSpec{
		{Name: "Property", Position: 0, PositionSet: true, Type: "string[]"},
		{Name: "ExpandProperty", Type: "string"},
		{Name: "First", Type: "int"},
		{Name: "Last", Type: "int"},
		{Name: "Unique", Switch: true},
	}, cmdSelectObject)
	Register("Sort-Object", []ParamSpec{
		{Name: "Property", Position: 0, PositionSet: true, Type: "string[]"},
		{Name: "Descending", Switch: true},
		{Name: "Unique", Switch: true},
		{Name: "CaseSensitive", Switch: true},
	}, cmdSortObject)
	Register("ForEach-Object", []ParamSpec{
		{Name: "Process", Position: 0, PositionSet: true, Type: "scriptblock"},
		{Name: "Begin", Type: "scriptblock"},
		{Name: "End", Type: "scriptblock"},
		{Name: "MemberName", Type: "string"},
	}, cmdForEachObject)
	Register("Group-Object", []ParamSpec{
		{Name: "Property", Position: 0, PositionSet: true, Type: "string"},
		{Name: "CaseSensitive", Switch: true},
	}, cmdGroupObject)
	Register("Measure-Object", []ParamSpec{
		{Name: "Property", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Sum", Switch: true},
		{Name: "Average", Switch: true},
		{Name: "Minimum", Switch: true},
		{Name: "Maximum", Switch: true},
		{Name: "Line", Switch: true},
	}, cmdMeasureObject)
	Register("Get-Member", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
		{Name: "MemberType", Type: "string"},
	}, cmdGetMember)
	Register("New-Object", []ParamSpec{
		{Name: "TypeName", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Property", Type: "hashtable"},
	}, cmdNewObject)
}

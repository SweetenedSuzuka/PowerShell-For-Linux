package builtin

import (
	"crypto/rand"
	"strings"

	"powershell/internal/object"
)

// sets.go 实现集合运算类 cmdlet（比较差异、去重、随机取样）。

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
	// 相等项显示参考集（ref）的值，对齐原版 PowerShell
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
	for _, byteVal := range b {
		v = v*256 + int64(byteVal)
	}
	if v < 0 {
		v = -v
	}
	if hi <= lo {
		return lo
	}
	return lo + v%(hi-lo)
}

// ---- 注册 ----

func init() {
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
}

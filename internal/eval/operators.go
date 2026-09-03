package eval

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"powershell/internal/builtin"
	"powershell/internal/external"
	"powershell/internal/lang"
	"powershell/internal/object"
)

// operators.go 实现纯运算符库（算术、比较、匹配、拼接、转换、位运算、通配）。
func (e *Evaluator) binaryOp(op string, l, r *object.PSObject) *object.PSObject {
	switch op {
	case "+":
		return addOp(l, r)
	case "-":
		return e.numOp(l, r, func(a, b float64) float64 { return a - b })
	case "*":
		return e.mulOp(l, r)
	case "/":
		// 除数为零报错并置 $?=false
		if rf, ok := r.AsFloat(); ok && rf == 0 {
			e.reportError(fmt.Errorf("%s", lang.T(lang.MsgDivideByZero)))
			return object.Null()
		}
		return e.numOp(l, r, func(a, b float64) float64 { return a / b })
	case "%":
		// 模数为零同样报错
		if rf, ok := r.AsFloat(); ok && rf == 0 {
			e.reportError(fmt.Errorf("%s", lang.T(lang.MsgDivideByZero)))
			return object.Null()
		}
		return e.numOp(l, r, func(a, b float64) float64 { return math.Mod(a, b) })
	case "..":
		return rangeOp(l, r)
	case "-eq", "-ne", "-lt", "-le", "-gt", "-ge", "-like", "-notlike", "-match", "-notmatch",
		"-ceq", "-cne", "-clt", "-cle", "-cgt", "-cge", "-clike", "-cnotlike", "-cmatch", "-cnotmatch",
		"-ieq", "-ine", "-ilt", "-ile", "-igt", "-ige", "-ilike", "-inotlike", "-imatch", "-inotmatch":
		// 数组左值 → 逐元素过滤
		if l.IsArray() {
			var matches []*object.PSObject
			for _, it := range l.ArrayItems() {
				if pairMatch(op, it, r) {
					matches = append(matches, it)
				}
			}
			if len(matches) == 0 {
				return object.Null()
			}
			if len(matches) == 1 {
				return matches[0]
			}
			return object.Array(matches)
		}
		if op == "-match" || op == "-cmatch" || op == "-imatch" {
			// 标量 -match/-cmatch 匹配成功后填充 $Matches（数组左值在上面的过滤分支里不填充）
			return e.evalMatch(op, l, r)
		}
		return object.Bool(pairMatch(op, l, r))
	case "-contains", "-icontains":
		return object.Bool(arrayContains(l, r))
	case "-notcontains", "-inotcontains":
		return object.Bool(!arrayContains(l, r))
	case "-ccontains":
		return object.Bool(arrayContainsCase(l, r))
	case "-cnotcontains":
		return object.Bool(!arrayContainsCase(l, r))
	case "-in", "-iin":
		return object.Bool(arrayContains(r, l))
	case "-notin", "-inotin":
		return object.Bool(!arrayContains(r, l))
	case "-join":
		return joinOp(l, r)
	case "-split":
		return splitOp(l, r)
	case "-replace":
		return replaceOp(l, r)
	case "-is":
		return object.Bool(typeIs(l, r))
	case "-isnot":
		return object.Bool(!typeIs(l, r))
	case "-as":
		return asOp(l, r)
	case "-band":
		return e.bitOp(l, r, func(a, b int64) int64 { return a & b })
	case "-bor":
		return e.bitOp(l, r, func(a, b int64) int64 { return a | b })
	case "-bxor":
		return e.bitOp(l, r, func(a, b int64) int64 { return a ^ b })
	case "-shl":
		return e.bitOp(l, r, func(a, b int64) int64 { return a << uint(b&63) })
	case "-shr":
		return e.bitOp(l, r, func(a, b int64) int64 { return a >> uint(b&63) })
	}
	return object.Bool(false)
}

// pairMatch 判断单个元素与右值是否满足比较运算符。
// PowerShell 默认比较大小写不敏感，-i* 前缀显式不敏感，-c* 前缀大小写敏感。
func pairMatch(op string, l, r *object.PSObject) bool {
	switch op {
	case "-eq", "-ieq":
		return compareEq(l, r)
	case "-ne", "-ine":
		return !compareEq(l, r)
	case "-ceq":
		return caseSensitiveEq(l, r)
	case "-cne":
		return !caseSensitiveEq(l, r)
	case "-lt", "-ilt":
		return compareOrder(l, r) < 0
	case "-le", "-ile":
		return compareOrder(l, r) <= 0
	case "-gt", "-igt":
		return compareOrder(l, r) > 0
	case "-ge", "-ige":
		return compareOrder(l, r) >= 0
	case "-clt":
		return caseSensitiveOrder(l, r) < 0
	case "-cle":
		return caseSensitiveOrder(l, r) <= 0
	case "-cgt":
		return caseSensitiveOrder(l, r) > 0
	case "-cge":
		return caseSensitiveOrder(l, r) >= 0
	case "-like", "-ilike":
		return object.WildcardMatchFold(r.String(), l.String())
	case "-notlike", "-inotlike":
		return !object.WildcardMatchFold(r.String(), l.String())
	case "-clike":
		return object.WildcardMatch(r.String(), l.String())
	case "-cnotlike":
		return !object.WildcardMatch(r.String(), l.String())
	case "-match", "-notmatch", "-cmatch", "-cnotmatch", "-imatch", "-inotmatch":
		re, err := compilePattern(op, r.String())
		if err != nil {
			return false
		}
		matched := re.MatchString(l.String())
		if op == "-notmatch" || op == "-cnotmatch" || op == "-inotmatch" {
			return !matched
		}
		return matched
	}
	return false
}

// compilePattern 编译 -match/-cmatch 用的正则，并把 .NET 命名组语法转成 Go 语法。
// -match/-notmatch 在模式前加 (?i) 实现大小写不敏感。
func compilePattern(op, pattern string) (*regexp.Regexp, error) {
	p := translateNamedGroups(pattern)
	if op == "-match" || op == "-notmatch" || op == "-imatch" || op == "-inotmatch" {
		p = "(?i)" + p
	}
	return regexp.Compile(p)
}

// translateNamedGroups 把 .NET 的命名组语法 (?<name>...) 转成 Go 的 (?P<name>...)。
// (?<= 与 (?<! 是环视，Go 不支持，原样保留，交给编译阶段报错。
func translateNamedGroups(pattern string) string {
	var sb strings.Builder
	for i := 0; i < len(pattern); {
		if i+2 < len(pattern) && pattern[i] == '(' && pattern[i+1] == '?' && pattern[i+2] == '<' {
			if i+3 < len(pattern) && (pattern[i+3] == '=' || pattern[i+3] == '!') {
				sb.WriteString("(?<") // 环视，保留原样
				i += 3
				continue
			}
			sb.WriteString("(?P<") // 命名组
			i += 3
			continue
		}
		sb.WriteByte(pattern[i])
		i++
	}
	return sb.String()
}

// evalMatch 处理标量的 -match/-cmatch：匹配成功后把捕获组写入 $Matches。
// 不匹配时不动 $Matches（与原版 PowerShell 一致）；数组左值不经过这里。
func (e *Evaluator) evalMatch(op string, l, r *object.PSObject) *object.PSObject {
	re, err := compilePattern(op, r.String())
	if err != nil {
		return object.Bool(false)
	}
	idx := re.FindStringSubmatchIndex(l.String())
	if idx == nil {
		return object.Bool(false)
	}
	e.Session.Matches = buildMatches(re, l.String(), idx)
	return object.Bool(true)
}

// buildMatches 由正则匹配结果构造 $Matches 哈希表。
// 键规则与原版 PowerShell 一致："0" 是整体匹配；命名组用组名；
// 未命名组按其在未命名组中的序号（从 1 起）作键；未参与的组不写入。
func buildMatches(re *regexp.Regexp, s string, idx []int) *object.PSObject {
	names := re.SubexpNames()
	entries := []object.HashEntry{{Key: "0", Value: object.Str(s[idx[0]:idx[1]])}}
	unnamed := 0
	for i := 1; i < len(names); i++ {
		key := names[i]
		if key == "" {
			unnamed++ // 未命名组按序号计数，未参与的也占号（与 .NET 一致）
			key = strconv.Itoa(unnamed)
		}
		if idx[2*i] < 0 {
			continue // 未参与的组不写入 $Matches
		}
		entries = append(entries, object.HashEntry{Key: key, Value: object.Str(s[idx[2*i]:idx[2*i+1]])})
	}
	return object.Hashtable(entries)
}

// caseSensitiveEq 大小写敏感相等（-ceq）。右操作数按左操作数类型转换，左为数字则两边按数字，左为字符串则两边按字符串，左为 $null 则只与 $null 相等。
func caseSensitiveEq(l, r *object.PSObject) bool {
	switch l.TypeName {
	case "Int", "Double", "Boolean":
		if ln, ok := l.AsFloat(); ok {
			if rn, ok2 := r.AsFloat(); ok2 {
				return ln == rn
			}
		}
	case "String":
		if r.IsNull() {
			return false
		}
		return l.String() == r.String()
	case "Null":
		return r.IsNull()
	}
	return l.String() == r.String()
}

// caseSensitiveOrder 大小写敏感顺序（-clt 等）。右操作数按左操作数类型参与比较：
// 左为数字则按数字，否则按字符串（"5" -clt "10" 是字典序，结果为 False）。
func caseSensitiveOrder(l, r *object.PSObject) int {
	switch l.TypeName {
	case "Int", "Double", "Boolean":
		if ln, ok := l.AsFloat(); ok {
			if rn, ok2 := r.AsFloat(); ok2 {
				if ln < rn {
					return -1
				}
				if ln > rn {
					return 1
				}
				return 0
			}
		}
	}
	return strings.Compare(l.String(), r.String())
}

// compareEq 相等判断（-eq）。右操作数按左操作数类型转换，左为数字则两边按数字，左为字符串则两边按字符串，左为 $null 则只与 $null 相等。
func compareEq(l, r *object.PSObject) bool {
	switch l.TypeName {
	case "Int", "Double", "Boolean":
		if ln, ok := l.AsFloat(); ok {
			if rn, ok2 := r.AsFloat(); ok2 {
				return ln == rn
			}
		}
	case "String":
		if r.IsNull() {
			return false
		}
		return strings.EqualFold(l.String(), r.String())
	case "Null":
		return r.IsNull()
	}
	return strings.EqualFold(l.String(), r.String())
}

// compareOrder 判断 -lt/-le/-gt/-ge 的顺序。右操作数按左操作数类型参与比较——左为数字则两边按数字（5 -lt "10" 为 True），左为字符串则两边按字符串（"5" -lt "10" 是字典序比较，结果为 False）。
func compareOrder(l, r *object.PSObject) int {
	switch l.TypeName {
	case "Int", "Double", "Boolean":
		if ln, ok := l.AsFloat(); ok {
			if rn, ok2 := r.AsFloat(); ok2 {
				if ln < rn {
					return -1
				}
				if ln > rn {
					return 1
				}
				return 0
			}
		}
	}
	return strings.Compare(strings.ToLower(l.String()), strings.ToLower(r.String()))
}

func arrayContains(arr, elem *object.PSObject) bool {
	for _, it := range arr.ArrayItems() {
		if compareEq(it, elem) {
			return true
		}
	}
	return false
}

// arrayContainsCase 大小写敏感版（-ccontains）。
func arrayContainsCase(arr, elem *object.PSObject) bool {
	for _, it := range arr.ArrayItems() {
		if caseSensitiveEq(it, elem) {
			return true
		}
	}
	return false
}

func addOp(l, r *object.PSObject) *object.PSObject {
	if l.IsArray() {
		return object.Array(append(l.ArrayItems(), r))
	}
	if r.IsArray() {
		return object.Array(append([]*object.PSObject{l}, r.ArrayItems()...))
	}
	// 任一操作数是浮点 → 浮点加法（整型路径会把 0.5 截断成 0，如 2 + 1/2）
	if l.TypeName == "Double" || r.TypeName == "Double" {
		if lf, ok := l.AsFloat(); ok {
			if rf, ok2 := r.AsFloat(); ok2 {
				return object.Float(lf + rf)
			}
		}
	}
	if li, ok := l.AsInt(); ok {
		if ri, ok2 := r.AsInt(); ok2 {
			return object.Int(li + ri)
		}
	}
	if lf, ok := l.AsFloat(); ok {
		if rf, ok2 := r.AsFloat(); ok2 {
			return object.Float(lf + rf)
		}
	}
	return object.Str(l.String() + r.String())
}

func (e *Evaluator) numOp(l, r *object.PSObject, fn func(a, b float64) float64) *object.PSObject {
	lf, ok := l.AsFloat()
	if !ok {
		e.throwError(lang.T(lang.MsgArithmeticInvalid))
		return object.Null()
	}
	rf, ok2 := r.AsFloat()
	if !ok2 {
		e.throwError(lang.T(lang.MsgArithmeticInvalid))
		return object.Null()
	}
	return object.Float(fn(lf, rf))
}

func (e *Evaluator) mulOp(l, r *object.PSObject) *object.PSObject {
	if l.TypeName == "String" {
		if n, ok := r.AsInt(); ok && n >= 0 {
			return object.Str(strings.Repeat(l.String(), int(n)))
		}
	}
	if r.TypeName == "String" {
		if n, ok := l.AsInt(); ok && n >= 0 {
			return object.Str(strings.Repeat(r.String(), int(n)))
		}
	}
	lf, ok := l.AsFloat()
	if !ok {
		e.throwError(lang.T(lang.MsgArithmeticInvalid))
		return object.Null()
	}
	rf, ok2 := r.AsFloat()
	if !ok2 {
		e.throwError(lang.T(lang.MsgArithmeticInvalid))
		return object.Null()
	}
	return object.Float(lf * rf)
}

func rangeOp(l, r *object.PSObject) *object.PSObject {
	ls, ok := l.AsInt()
	if !ok {
		return object.Null()
	}
	rs, ok2 := r.AsInt()
	if !ok2 {
		return object.Null()
	}
	var items []*object.PSObject
	if ls <= rs {
		for i := ls; i <= rs; i++ {
			items = append(items, object.Int(i))
		}
	} else {
		for i := ls; i >= rs; i-- {
			items = append(items, object.Int(i))
		}
	}
	return object.Array(items)
}

// formatOp 实现 .NET 风格格式串："{模板}" -f 值[, 值...]。
// 支持 {N}、{N,宽度}（空格对齐）、{N:规格}（D 十进制补零、X/x 十六进制、F 定点小数、N 千分位）。
// {{ 与 }} 转义字面大括号；未知规格退化为原样字符串；下标越界 → 报错并置 $?=false。
func (e *Evaluator) formatOp(f, args *object.PSObject) *object.PSObject {
	format := f.String()
	items := flattenArgs(args)
	var sb strings.Builder
	for i := 0; i < len(format); i++ {
		c := format[i]
		if c != '{' {
			sb.WriteByte(c)
			continue
		}
		if i+1 < len(format) && format[i+1] == '{' {
			sb.WriteByte('{')
			i++
			continue
		}
		j := strings.IndexByte(format[i+1:], '}')
		if j < 0 {
			sb.WriteString(format[i:])
			break
		}
		inner := format[i+1 : i+1+j]
		idxPart := inner
		spec := ""
		if k := strings.IndexByte(inner, ':'); k >= 0 {
			idxPart, spec = inner[:k], inner[k+1:]
		}
		alignPart := idxPart
		widthStr := ""
		if k := strings.IndexByte(idxPart, ','); k >= 0 {
			alignPart, widthStr = idxPart[:k], idxPart[k+1:]
		}
		idx, err := strconv.Atoi(strings.TrimSpace(alignPart))
		if err != nil || idx < 0 || idx >= len(items) {
			e.reportError(fmt.Errorf("%s", lang.T(lang.MsgFormatIndexOut, inner, len(items))))
			return object.Null()
		}
		s := formatArg(items[idx], spec)
		if w, werr := strconv.Atoi(strings.TrimSpace(widthStr)); werr == nil {
			aw := int(math.Abs(float64(w)))
			if len(s) < aw {
				pad := strings.Repeat(" ", aw-len(s))
				if w < 0 {
					s += pad
				} else {
					s = pad + s
				}
			}
		}
		sb.WriteString(s)
		i += j + 1
	}
	return object.Str(sb.String())
}

// flattenArgs 把 -f 的参数数组展平为位置参数列表（范围/括号数组/变量数组都摊开）。
func flattenArgs(v *object.PSObject) []*object.PSObject {
	if !v.IsArray() {
		return []*object.PSObject{v}
	}
	var out []*object.PSObject
	for _, it := range v.ArrayItems() {
		if it.IsArray() {
			out = append(out, flattenArgs(it)...)
		} else {
			out = append(out, it)
		}
	}
	return out
}

// formatArg 按 .NET 格式规格格式化单个参数。
func formatArg(v *object.PSObject, spec string) string {
	if spec == "" {
		return v.String()
	}
	letter := spec[0]
	width := spec[1:]
	switch letter {
	case 'D', 'd':
		if n, ok := v.AsInt(); ok {
			w, _ := strconv.Atoi(width)
			digits := strconv.FormatInt(n, 10)
			neg := strings.HasPrefix(digits, "-")
			if neg {
				digits = digits[1:]
			}
			if w > len(digits) {
				digits = strings.Repeat("0", w-len(digits)) + digits
			}
			if neg {
				digits = "-" + digits
			}
			return digits
		}
	case 'X', 'x':
		if n, ok := v.AsInt(); ok {
			s := strconv.FormatInt(n, 16)
			if letter == 'X' {
				s = strings.ToUpper(s)
			}
			w, _ := strconv.Atoi(width)
			if w > len(s) {
				s = strings.Repeat("0", w-len(s)) + s
			}
			return s
		}
	case 'F', 'f':
		if fv, ok := v.AsFloat(); ok {
			dec := 2
			if d, derr := strconv.Atoi(width); derr == nil {
				dec = d
			}
			return strconv.FormatFloat(fv, 'f', dec, 64)
		}
	case 'N', 'n':
		if fv, ok := v.AsFloat(); ok {
			dec := 2
			if d, derr := strconv.Atoi(width); derr == nil {
				dec = d
			}
			return addThousands(strconv.FormatFloat(fv, 'f', dec, 64))
		}
	}
	// 未知规格：退化为普通字符串
	return v.String()
}

// addThousands 给数字串的整数部分加千分位逗号（.NET N 规格）。
func addThousands(s string) string {
	dot := strings.IndexByte(s, '.')
	intPart, fracPart := s, ""
	if dot >= 0 {
		intPart, fracPart = s[:dot], s[dot:]
	}
	sign := ""
	if strings.HasPrefix(intPart, "-") {
		sign, intPart = "-", intPart[1:]
	}
	var b strings.Builder
	b.WriteString(sign)
	n := len(intPart)
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteByte(intPart[i])
	}
	return b.String() + fracPart
}

func joinOp(l, r *object.PSObject) *object.PSObject {
	sep := r.String()
	var parts []string
	if l.IsArray() {
		for _, it := range l.ArrayItems() {
			parts = append(parts, it.String())
		}
	} else {
		parts = []string{l.String()}
	}
	return object.Str(strings.Join(parts, sep))
}

func splitOp(l, r *object.PSObject) *object.PSObject {
	// 右操作数可为 分隔符 或 分隔符, 最大子串数（"a,b,c" -split ",",2 → a、b,c）
	delim := r.String()
	maxN := -1
	if r.IsArray() {
		items := r.ArrayItems()
		if len(items) >= 1 {
			delim = items[0].String()
		}
		if len(items) >= 2 {
			if n, ok := items[1].AsInt(); ok {
				maxN = int(n)
			}
		}
	}
	re, err := regexp.Compile(delim)
	if err != nil {
		re = regexp.MustCompile(regexp.QuoteMeta(delim))
	}
	// n>0 时最多分 n 段、末段保留未分割剩余；n=0 或负数不限
	n := -1
	if maxN > 0 {
		n = maxN
	}
	parts := re.Split(l.String(), n)
	items := make([]*object.PSObject, len(parts))
	for i, p := range parts {
		items[i] = object.Str(p)
	}
	return object.Array(items)
}

func replaceOp(l, r *object.PSObject) *object.PSObject {
	items := r.ArrayItems()
	pattern := ""
	replacement := ""
	if len(items) >= 1 {
		pattern = items[0].String()
	}
	if len(items) >= 2 {
		replacement = items[1].String()
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return l
	}
	return object.Str(re.ReplaceAllString(l.String(), replacement))
}

func typeIs(l, r *object.PSObject) bool {
	name := strings.ToLower(r.String())
	switch name {
	case "int", "int32", "int64", "long":
		return l.TypeName == "Int"
	case "string":
		return l.TypeName == "String"
	case "bool", "boolean":
		return l.TypeName == "Boolean"
	case "double", "float", "single":
		return l.TypeName == "Double"
	case "array", "object[]":
		return l.IsArray()
	case "hashtable":
		return l.TypeName == "Hashtable"
	case "datetime":
		return l.TypeName == "DateTime"
	}
	return false
}

func asOp(l, r *object.PSObject) *object.PSObject {
	switch strings.ToLower(r.String()) {
	case "int", "int32", "int64", "long":
		if n, ok := l.AsInt(); ok {
			return object.Int(n)
		}
		return object.Null()
	case "double", "float", "single":
		if n, ok := l.AsFloat(); ok {
			return object.Float(n)
		}
		return object.Null()
	case "string":
		return object.Str(l.String())
	case "bool":
		return object.Bool(l.Truthy())
	case "datetime":
		if v, ok := parseDatetimeValue(l); ok {
			return v
		}
		return object.Null()
	case "array", "object[]":
		if l.IsArray() {
			return l
		}
		return object.Array(l.ArrayItems())
	}
	return l
}

// convertValue 把值转换为方括号里声明的类型；数组后缀对每个元素分别转换。
// 无法转换时写非终止错误并返回 $null（与除零的既有错误风格一致）。
func (e *Evaluator) convertValue(v *object.PSObject, typeName string) *object.PSObject {
	norm := strings.ToLower(typeName)
	if strings.HasSuffix(norm, "[]") {
		elem := strings.TrimSuffix(norm, "[]")
		items := make([]*object.PSObject, 0, len(v.ArrayItems()))
		for _, it := range v.ArrayItems() {
			items = append(items, e.convertScalar(it, elem))
		}
		return object.Array(items)
	}
	return e.convertScalar(v, norm)
}

// convertScalar 单值转换：target 为归一化小写类型名，未知类型报"无法找到类型"。
func (e *Evaluator) convertScalar(v *object.PSObject, target string) *object.PSObject {
	out, err := convertTarget(v, target)
	if err != nil {
		e.reportError(err)
		return object.Null()
	}
	return out
}

// convertTarget 单值转换尝试；无法转换或目标类型未注册时返回错误说明。
func convertTarget(v *object.PSObject, target string) (*object.PSObject, error) {
	switch target {
	case "int", "int32", "int64", "long":
		if n, ok := v.AsInt(); ok {
			return object.Int(n), nil
		}
	case "string":
		return object.Str(v.String()), nil
	case "double", "float", "single":
		if f, ok := v.AsFloat(); ok {
			return object.Float(f), nil
		}
	case "bool", "boolean":
		return object.Bool(v.Truthy()), nil
	case "hashtable":
		if _, ok := v.Value.([]object.HashEntry); ok {
			return v, nil
		}
	case "datetime":
		if dv, ok := parseDatetimeValue(v); ok {
			return dv, nil
		}
	case "void":
		return object.Null(), nil
	default:
		return nil, errTypeUnknown
	}
	return nil, fmt.Errorf("%s", lang.T(lang.MsgConvertFail, v.String(), target))
}

// errTypeUnknown 表示目标类型未注册，区别于值无法转换成已注册类型。
var errTypeUnknown = errors.New("type unknown")

// parseDatetimeValue 从 DateTime 原值或常见格式的字符串构造时间对象。
func parseDatetimeValue(v *object.PSObject) (*object.PSObject, bool) {
	if v.TypeName == "DateTime" {
		return v, true
	}
	s := v.String()
	for _, layout := range []string{time.RFC3339, "2006-01-02 15:04:05", "2006-01-02", "15:04:05"} {
		if t, err := time.Parse(layout, s); err == nil {
			return object.DateTime(t), true
		}
	}
	return nil, false
}

func (e *Evaluator) bitOp(l, r *object.PSObject, fn func(a, b int64) int64) *object.PSObject {
	li, ok := l.AsInt()
	if !ok {
		e.throwError(lang.T(lang.MsgArithmeticInvalid))
		return object.Null()
	}
	ri, ok2 := r.AsInt()
	if !ok2 {
		e.throwError(lang.T(lang.MsgArithmeticInvalid))
		return object.Null()
	}
	return object.Int(fn(li, ri))
}

// wildcardMatch 委托给 object 包的通配符实现。
func wildcardMatch(pattern, s string) bool {
	return object.WildcardMatch(pattern, s)
}

// requireBuiltin 简化：确保包被引用（编译期占位）。
var _ = builtin.Lookup
var _ = external.LookPath

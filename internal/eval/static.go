// static.go 实现 [类型]::成员 的静态成员分派表。
// 覆盖 math/string/datetime/guid/int 等常用类型；未注册的类型或成员返回"不存在成员"错误。
package eval

import (
	"crypto/rand"
	"fmt"
	"math"
	"strings"
	"time"

	"powershell/internal/object"
)

// staticMember 分派静态属性与静态方法；ok 为 false 表示类型或成员未注册。
func (e *Evaluator) staticMember(typeName, name string, args []*object.PSObject) (*object.PSObject, bool) {
	typ := strings.ToLower(strings.TrimPrefix(strings.ToLower(typeName), "system."))
	norm := strings.ToLower(name)
	argN := len(args)
	f1 := func(i int) (float64, bool) {
		if argN > i {
			return args[i].AsFloat()
		}
		return 0, false
	}
	f2 := func(i int) (float64, float64, bool) {
		if argN > i+1 {
			a, ok1 := args[i].AsFloat()
			b, ok2 := args[i+1].AsFloat()
			return a, b, ok1 && ok2
		}
		return 0, 0, false
	}

	switch typ {
	case "math":
		switch norm {
		case "abs":
			if f, ok := f1(0); ok {
				return object.Float(math.Abs(f)), true
			}
		case "sqrt":
			if f, ok := f1(0); ok {
				return object.Float(math.Sqrt(f)), true
			}
		case "floor":
			if f, ok := f1(0); ok {
				return object.Float(math.Floor(f)), true
			}
		case "ceiling":
			if f, ok := f1(0); ok {
				return object.Float(math.Ceil(f)), true
			}
		case "truncate":
			if f, ok := f1(0); ok {
				return object.Float(math.Trunc(f)), true
			}
		case "exp":
			if f, ok := f1(0); ok {
				return object.Float(math.Exp(f)), true
			}
		case "log":
			if f, ok := f1(0); ok {
				if base, ok2 := f1(1); ok2 {
					return object.Float(math.Log(f) / math.Log(base)), true
				}
				return object.Float(math.Log(f)), true
			}
		case "log10":
			if f, ok := f1(0); ok {
				return object.Float(math.Log10(f)), true
			}
		case "sign":
			if f, ok := f1(0); ok {
				switch {
				case f > 0:
					return object.Int(1), true
				case f < 0:
					return object.Int(-1), true
				}
				return object.Int(0), true
			}
		case "pow":
			if a, b, ok := f2(0); ok {
				return object.Float(math.Pow(a, b)), true
			}
		case "max", "min":
			if a, b, ok := f2(0); ok {
				if (norm == "max") == (a >= b) {
					return object.Float(a), true
				}
				return object.Float(b), true
			}
		case "round":
			if f, ok := f1(0); ok {
				digits := 0.0
				if argN > 1 {
					if d, ok2 := args[1].AsInt(); ok2 && d > 0 && d < 15 {
						digits = float64(d)
					}
				}
				shift := math.Pow(10, digits)
				return object.Float(math.RoundToEven(f*shift) / shift), true
			}
		case "pi":
			if argN == 0 {
				return object.Float(math.Pi), true
			}
		}

	case "string":
		switch norm {
		case "isnullorempty":
			if argN >= 1 {
				return object.Bool(args[0].IsNull() || args[0].String() == ""), true
			}
		case "isnullorwhitespace":
			if argN >= 1 {
				return object.Bool(args[0].IsNull() || strings.TrimSpace(args[0].String()) == ""), true
			}
		case "join":
			if argN >= 2 {
				sep := args[0].String()
				var parts []string
				for _, v := range args[1:] {
					for _, it := range v.ArrayItems() {
						parts = append(parts, it.String())
					}
				}
				return object.Str(strings.Join(parts, sep)), true
			}
		case "concat":
			var sb strings.Builder
			for _, v := range args {
				sb.WriteString(v.String())
			}
			return object.Str(sb.String()), true
		case "format":
			if argN >= 1 {
				return e.formatOp(args[0], object.Array(args[1:])), true
			}
		}

	case "datetime":
		switch norm {
		case "now":
			return object.DateTime(time.Now()), true
		case "utcnow":
			return object.DateTime(time.Now().UTC()), true
		case "today":
			now := time.Now()
			return object.DateTime(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())), true
		case "parse":
			if argN >= 1 {
				if v, ok := parseDatetimeValue(args[0]); ok {
					return v, true
				}
			}
		}

	case "guid":
		if norm == "newguid" && argN == 0 {
			return object.Str(newGuidV4()), true
		}

	case "int", "int32":
		switch norm {
		case "maxvalue":
			return object.Int(int64(int32(^uint32(0) >> 1))), true
		case "minvalue":
			return object.Int(int64(-int32(^uint32(0)>>1) - 1)), true
		}
	case "int64", "long":
		switch norm {
		case "maxvalue":
			return object.Int(int64(^uint64(0) >> 1)), true
		case "minvalue":
			return object.Int(-int64(^uint64(0)>>1) - 1), true
		}
	}
	return nil, false
}

// newGuidV4 生成随机 UUID（版本 4 格式，crypto/rand 驱动）。
func newGuidV4() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

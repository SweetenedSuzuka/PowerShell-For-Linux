package object

import (
	"strings"
	"testing"
	"time"
)

func TestTruthy(t *testing.T) {
	cases := []struct {
		o    *PSObject
		want bool
	}{
		{Null(), false},
		{Bool(false), false},
		{Bool(true), true},
		{Int(0), false},
		{Int(1), true},
		{Float(0.0), false},
		{Float(0.5), true},
		{Str(""), false},
		{Str("0"), true},
		{Str("false"), true},
		{Array(nil), false},
		{Array([]*PSObject{Int(1)}), true},
		{Hashtable(nil), true},
	}
	for _, c := range cases {
		if got := c.o.Truthy(); got != c.want {
			t.Errorf("Truthy(%s) = %v, want %v", c.o.String(), got, c.want)
		}
	}
}

func TestStringConversion(t *testing.T) {
	if Str("hi").String() != "hi" {
		t.Error("string")
	}
	if Bool(true).String() != "True" {
		t.Error("bool")
	}
	if Int(-42).String() != "-42" {
		t.Error("int")
	}
	if Float(1.5).String() != "1.5" {
		t.Errorf("float: %s", Float(1.5).String())
	}
	if Float(2.0).String() != "2" {
		t.Errorf("float whole: %s", Float(2.0).String())
	}
}

func TestArrayAndProps(t *testing.T) {
	f := Object("File", "/tmp/a")
	f.AddProp("Name", "a")
	f.AddProp("Length", int64(3))
	if v, ok := f.PropValue("name"); !ok || v.String() != "a" {
		t.Error("prop lookup (case-insensitive)")
	}
	if v, ok := f.PropValue("Length"); !ok || v.String() != "3" {
		t.Error("prop length")
	}
	arr := Array([]*PSObject{Int(1), Str("x")})
	if arr.IsNull() || !arr.IsArray() {
		t.Error("array flags")
	}
	if len(arr.ArrayItems()) != 2 {
		t.Error("array items")
	}
}

func TestFormatTable(t *testing.T) {
	o1 := Object("File", "/a")
	o1.AddProp("Name", "one").AddProp("Length", int64(10))
	o1.Table = []Column{{Label: "Name", Align: "left"}, {Label: "Length", Align: "right"}}
	o2 := Object("File", "/b")
	o2.AddProp("Name", "longername").AddProp("Length", int64(12345))
	o2.Table = o1.Table
	var sb strings.Builder
	if err := FormatTableTo(&sb, []*PSObject{o1, o2}, nil); err != nil {
		t.Fatal(err)
	}
	out := sb.String()
	if !strings.Contains(out, "Name") || !strings.Contains(out, "longername") || !strings.Contains(out, "----") {
		t.Errorf("表格输出异常:\n%s", out)
	}
	if !strings.HasPrefix(out, "Name") {
		t.Errorf("表格应以表头开头:\n%s", out)
	}
}

func TestFormatList(t *testing.T) {
	o := Object("X", 1)
	o.AddProp("Name", "a").AddProp("Desc", "d")
	var sb strings.Builder
	if err := FormatListTo(&sb, []*PSObject{o}, nil); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(sb.String(), "Name : a") || !strings.Contains(sb.String(), "Desc : d") {
		t.Errorf("列表输出异常:\n%s", sb.String())
	}
}

func TestFormatListScalarProps(t *testing.T) {
	// 标量类型对象显式指定属性时，应列出属性而不被标量捷径忽略
	o := DateTime(time.Date(2026, 8, 15, 12, 0, 0, 0, time.UTC))
	o.AddProp("标签", "测试")
	var sb strings.Builder
	if err := FormatListTo(&sb, []*PSObject{o}, []string{"标签"}); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(sb.String(), "标签 : 测试") {
		t.Errorf("标量带属性列出异常:\n%s", sb.String())
	}
	// 不指定属性时仍按标量原样输出
	sb.Reset()
	if err := FormatListTo(&sb, []*PSObject{o}, nil); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(sb.String(), "2026") {
		t.Errorf("标量默认输出异常:\n%s", sb.String())
	}
}

func TestFormatOutputStrings(t *testing.T) {
	var sb strings.Builder
	if err := FormatOutput(&sb, []*PSObject{Str("a"), Str("b")}); err != nil {
		t.Fatal(err)
	}
	if sb.String() != "a\nb\n" {
		t.Errorf("字符串输出 = %q", sb.String())
	}
}

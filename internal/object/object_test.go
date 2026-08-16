package object

import (
	"os"
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

func TestVersionObject(t *testing.T) {
	// 末尾为 0 的段省略，至少保留 major
	if got := Version(7, 0, 0, 0).String(); got != "7" {
		t.Errorf("Version(7,0,0,0).String() = %q，想要 \"7\"", got)
	}
	if got := Version(5, 1, 0, 0).String(); got != "5.1" {
		t.Errorf("Version(5,1,0,0).String() = %q，想要 \"5.1\"", got)
	}
	if got := Version(1, 2, 3, 0).String(); got != "1.2.3" {
		t.Errorf("Version(1,2,3,0).String() = %q，想要 \"1.2.3\"", got)
	}
	if got := Version(1, 2, 3, 4).String(); got != "1.2.3.4" {
		t.Errorf("Version(1,2,3,4).String() = %q，想要 \"1.2.3.4\"", got)
	}
	v := Version(7, 0, 0, 0)
	if p, ok := v.PropValue("Major"); !ok || p.String() != "7" {
		t.Errorf("Version Major 取不到或值错：%v %v", p, ok)
	}
	if p, ok := v.PropValue("Minor"); !ok || p.String() != "0" {
		t.Errorf("Version Minor 取不到或值错：%v %v", p, ok)
	}
	if p, ok := v.PropValue("Build"); !ok || p.String() != "0" {
		t.Errorf("Version Build 取不到或值错：%v %v", p, ok)
	}
}

func TestHashtableVirtualProps(t *testing.T) {
	h := Hashtable([]HashEntry{
		{Key: "a", Value: Int(1)},
		{Key: "b", Value: Int(2)},
	})
	if v, ok := h.PropValue("Count"); !ok || v.String() != "2" {
		t.Errorf("Count 取不到或值错：%v %v", v, ok)
	}
	// 哈希表没有 Length，应走标量兜底返回 1
	if v, ok := h.PropValue("Length"); !ok || v.String() != "1" {
		t.Errorf("Length 取不到或值错：%v %v", v, ok)
	}
	if v, ok := h.PropValue("Keys"); !ok || v.String() != "a b" {
		t.Errorf("Keys 取不到或值错：%v %v", v, ok)
	}
	if v, ok := h.PropValue("Values"); !ok || v.String() != "1 2" {
		t.Errorf("Values 取不到或值错：%v %v", v, ok)
	}
	// 键优先于属性：存在 Count 键时返回键值而非条目数
	hc := Hashtable([]HashEntry{{Key: "Count", Value: Int(5)}})
	if v, ok := hc.PropValue("Count"); !ok || v.String() != "5" {
		t.Errorf("键 Count 应优先于属性：%v %v", v, ok)
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

// TestFileInfoVirtualProps 验证 FileInfo/DirectoryInfo 虚拟属性：Extension/BaseName/DirectoryName 从路径计算。
func TestFileInfoVirtualProps(t *testing.T) {
	fi, err := os.Stat("object.go")
	if err != nil {
		t.Fatal(err)
	}
	// 文件：Extension 含点、BaseName 去扩展名、DirectoryName 父目录
	f := FileInfo("/some/dir/report.tar.gz", fi)
	if v, ok := f.PropValue("Extension"); !ok || v.String() != ".gz" {
		t.Errorf("文件 Extension = %v", v)
	}
	if v, ok := f.PropValue("BaseName"); !ok || v.String() != "report.tar" {
		t.Errorf("文件 BaseName = %v", v)
	}
	if v, ok := f.PropValue("DirectoryName"); !ok || v.String() != "/some/dir" {
		t.Errorf("文件 DirectoryName = %v", v)
	}
	// 无扩展名文件：Extension 空串、BaseName 原名
	f2 := FileInfo("/tmp/README", fi)
	if v, _ := f2.PropValue("Extension"); v.String() != "" {
		t.Errorf("无扩展名 Extension = %q", v.String())
	}
	if v, _ := f2.PropValue("BaseName"); v.String() != "README" {
		t.Errorf("无扩展名 BaseName = %q", v.String())
	}
	// 目录：Extension 恒空、BaseName 是目录名（不去扩展名）
	d := DirInfo("/some/dir.v2/sub", fi)
	if v, _ := d.PropValue("Extension"); v.String() != "" {
		t.Errorf("目录 Extension = %q", v.String())
	}
	if v, _ := d.PropValue("BaseName"); v.String() != "sub" {
		t.Errorf("目录 BaseName = %q", v.String())
	}
	if v, _ := d.PropValue("DirectoryName"); v.String() != "/some/dir.v2" {
		t.Errorf("目录 DirectoryName = %q", v.String())
	}
}

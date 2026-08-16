package shell

import "testing"

// TestDrivePath 盘符前缀归一化：C: 当根目录，其它盘符报错，无盘符原样。
func TestDrivePath(t *testing.T) {
	cases := []struct {
		in   string
		want string
		err  bool
	}{
		{"C:\\tmp\\x", "/tmp/x", false},
		{"C:/tmp/x", "/tmp/x", false},
		{"C:x", "/x", false},
		{"C:", "/", false},
		{"C:\\", "/", false},
		{"c:\\tmp", "/tmp", false},
		{"/tmp/x", "/tmp/x", false},
		{"tmp/x", "tmp/x", false},
		{"D:\\x", "", true},
		{"d:/x", "", true},
		{"A:foo", "", true},
	}
	for _, c := range cases {
		got, err := DrivePath(c.in)
		if c.err {
			if err == nil {
				t.Errorf("%q 应报错，实际得 %q", c.in, got)
			}
			continue
		}
		if err != nil {
			t.Errorf("%q 不应报错：%v", c.in, err)
			continue
		}
		if got != c.want {
			t.Errorf("DrivePath(%q) = %q，想要 %q", c.in, got, c.want)
		}
	}
}

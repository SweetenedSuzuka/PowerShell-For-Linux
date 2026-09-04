package shell

import "testing"

// TestDrivePath 盘符前缀归一化：C: 当根目录，其它盘符报错，无盘符原样。
func TestDrivePath(t *testing.T) {
	cases := []struct {
		in   string
		expected string
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
	for _, tc := range cases {
		actual, err := DrivePath(tc.in)
		if tc.err {
			if err == nil {
				t.Errorf("%q 应报错，实际得 %q", tc.in, actual)
			}
			continue
		}
		if err != nil {
			t.Errorf("%q 不应报错：%v", tc.in, err)
			continue
		}
		if actual != tc.expected {
			t.Errorf("DrivePath(%q) = %q，想要 %q", tc.in, actual, tc.expected)
		}
	}
}

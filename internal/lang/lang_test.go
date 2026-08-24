package lang

import "testing"

// TestDetect 探测规则：候选语言码命中已注册语言即采用，其余语种与无值情形回退默认语言。
func TestDetect(t *testing.T) {
	cases := []struct {
		name string
		env  map[string]string
		want Lang
	}{
		{"已注册 en 命中", map[string]string{"LANG": "en_US.UTF-8"}, LangEn},
		{"LC_ALL 的 en 命中", map[string]string{"LC_ALL": "en_GB", "LANG": "zh_CN.UTF-8"}, LangEn},
		{"LC_MESSAGES 优先于 LANG", map[string]string{"LC_MESSAGES": "en_US", "LANG": "zh_CN.UTF-8"}, LangEn},
		{"已注册 zh 命中", map[string]string{"LANG": "zh_CN.UTF-8"}, LangZh},
		{"繁体同属 zh", map[string]string{"LANG": "zh_TW"}, LangZh},
		{"未注册语种看下一项直至默认", map[string]string{"LANG": "fr_FR.UTF-8"}, LangZh},
		{"未注册语种日语回退默认", map[string]string{"LANG": "ja_JP"}, LangZh},
		{"未设置用默认语言", nil, LangZh},
		{"C 与 POSIX 视为无值", map[string]string{"LANG": "C"}, LangZh},
		{"C.UTF-8 视为无值", map[string]string{"LANG": "C.UTF-8"}, LangZh},
		{"POSIX 视为无值", map[string]string{"LANG": "POSIX"}, LangZh},
		{"LANGUAGE 列表命中 en", map[string]string{"LANGUAGE": "fr:en:zh", "LANG": "zh_CN.UTF-8"}, LangEn},
		{"LANGUAGE 全未注册落到 LANG 的 en", map[string]string{"LANGUAGE": "fr:de", "LANG": "en_US.UTF-8"}, LangEn},
		{"LANGUAGE 优先于 LANG", map[string]string{"LANGUAGE": "zh_CN:en", "LANG": "en_US.UTF-8"}, LangZh},
		{"空 LANGUAGE 落到 LANG", map[string]string{"LANGUAGE": "", "LANG": "en_US.UTF-8"}, LangEn},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for _, k := range localeVars {
				t.Setenv(k, "")
			}
			for k, v := range c.env {
				t.Setenv(k, v)
			}
			if got := Detect(); got != c.want {
				t.Errorf("Detect() = %q，想要 %q", got, c.want)
			}
		})
	}
}

// TestLanguageCode 语言码截取。
func TestLanguageCode(t *testing.T) {
	cases := []struct{ in, want string }{
		{"zh_CN.UTF-8", "zh"},
		{"en_US", "en"},
		{"fr_FR@euro", "fr"},
		{"de-DE", "de"},
		{"C", "C"},
		{"plain", "plain"},
	}
	for _, c := range cases {
		if got := languageCode(c.in); got != c.want {
			t.Errorf("languageCode(%q) = %q，想要 %q", c.in, got, c.want)
		}
	}
}

// TestT 取文本：按当前语言查表填参，缺条目回退默认语言表。
func TestT(t *testing.T) {
	SetCurrent(LangZh)
	if got := T(MsgDivideByZero); got != "尝试除以零" {
		t.Errorf("默认语言文本 = %q", got)
	}
	SetCurrent(LangEn)
	if got := T(MsgPathNotFoundFmt, "/nope"); got != "Cannot find path '/nope'." {
		t.Errorf("英文文本 = %q", got)
	}
}

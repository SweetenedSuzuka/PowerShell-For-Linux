package lang

import (
	"testing"
)

// TestTablesComplete 两张语言表必须覆盖全部 Msg 编号，不多不少。
func TestTablesComplete(t *testing.T) {
	maxMsg := int(MsgFlagParseFail)
	for i := 0; i <= maxMsg; i++ {
		m := Msg(i)
		if _, ok := zh[m]; !ok {
			t.Errorf("zh 表缺条目 Msg(%d)", i)
		}
		if _, ok := en[m]; !ok {
			t.Errorf("en 表缺条目 Msg(%d)", i)
		}
	}
	if len(zh) != maxMsg+1 {
		t.Errorf("zh 表有 %d 条，应为 %d（可能含未定义编号的键）", len(zh), maxMsg+1)
	}
	if len(en) != maxMsg+1 {
		t.Errorf("en 表有 %d 条，应为 %d", len(en), maxMsg+1)
	}
}

// TestPlaceholderParity 同一条目两种语言的格式占位符序列必须一致（实参顺序对齐约定的机器检查）。
func TestPlaceholderParity(t *testing.T) {
	maxMsg := int(MsgFlagParseFail)
	for i := 0; i <= maxMsg; i++ {
		m := Msg(i)
		zs := placeholders(zh[m])
		es := placeholders(en[m])
		if len(zs) != len(es) {
			t.Errorf("Msg(%d) 占位符个数不一致：zh %q 有 %d 个，en %q 有 %d 个", i, zh[m], len(zs), en[m], len(es))
			continue
		}
		for k := range zs {
			if zs[k] != es[k] {
				t.Errorf("Msg(%d) 第 %d 个占位符不同：zh %q vs en %q", i, k, zs[k], es[k])
			}
		}
	}
}

// placeholders 按出现顺序抽出 fmt 格式串里的动词（%% 除外）。
func placeholders(s string) []string {
	var out []string
	for i := 0; i < len(s); i++ {
		if s[i] != '%' {
			continue
		}
		if i+1 < len(s) && s[i+1] == '%' {
			i++
			continue
		}
		// 动词主体：跳过宽度/精度标记直到动词字母
		j := i + 1
		for j < len(s) && s[j] != 'v' && s[j] != 's' && s[j] != 'd' && s[j] != 'q' && s[j] != 'x' {
			j++
		}
		if j < len(s) {
			out = append(out, s[i:j+1])
			i = j
		}
	}
	return out
}

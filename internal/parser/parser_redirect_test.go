package parser

import (
	"strings"
	"testing"
)

func TestRedirection(t *testing.T) {
	d := dump(parseOK(t, "Get-Content a > out.txt"))
	if !strings.Contains(d, "[redir>word(out.txt)]") {
		t.Fatalf("重定向解析失败: %s", d)
	}
	d = dump(parseOK(t, "Get-Content a 2> err.txt"))
	if !strings.Contains(d, "[redir>word(err.txt)]") {
		t.Fatalf("2> 重定向解析失败: %s", d)
	}
	// 开关后的 2> 是错误重定向，不是开关的值，也不是属性名。
	d = dump(parseOK(t, "Format-Table -AutoSize 2> err.txt"))
	if !strings.Contains(d, "-AutoSize") || !strings.Contains(d, "[redir>") {
		t.Fatalf("开关后重定向解析失败: %s", d)
	}
	if strings.Contains(d, "2>") {
		t.Fatalf("2> 被合并进实参: %s", d)
	}
}

// TestExprRedirect 验证纯表达式管道头的尾随重定向收进管道（如 $x 2>err.txt），同流两次报错。
func TestExprRedirect(t *testing.T) {
	d := dump(parseOK(t, "$x 2> err.txt"))
	if !strings.Contains(d, "[redir>word(err.txt)]") {
		t.Fatalf("表达式 2> 重定向解析失败: %s", d)
	}
	d = dump(parseOK(t, `"hi" > out.txt`))
	if !strings.Contains(d, "[redir>word(out.txt)]") {
		t.Fatalf("表达式 > 重定向解析失败: %s", d)
	}
	for _, src := range []string{
		`$x > a.txt > b.txt`,
		`$x 2> a.txt 2> b.txt`,
	} {
		res := Parse(src)
		if res.Error == nil {
			t.Errorf("%q 应报错，实际通过", src)
		}
	}
}

// TestDupRedirectError 验证同一命令同流重定向两次报错，混合流合法。
func TestDupRedirectError(t *testing.T) {
	for _, src := range []string{
		`Write-Output ok > a.txt > b.txt`,
		`Write-Output ok 2> a.txt 2> b.txt`,
		`Write-Output ok > a.txt >> b.txt`,
	} {
		res := Parse(src)
		if res.Error == nil {
			t.Errorf("%q 应报错，实际通过", src)
		}
	}
	for _, src := range []string{
		`Write-Output ok 2> a.txt > b.txt`,
		`Write-Output ok > a.txt`,
	} {
		if res := Parse(src); res.Error != nil || res.Incomplete {
			t.Errorf("%q 应可解析，实际 err=%v", src, res.Error)
		}
	}
}

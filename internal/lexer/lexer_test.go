package lexer

import "testing"

// tokenTypes 返回一串 token 的类型名，便于断言。
func tokenTypes(src string) []TokenType {
	toks := New(src).Tokens()
	var out []TokenType
	for _, t := range toks {
		out = append(out, t.Type)
	}
	return out
}

func hasType(types []TokenType, tt TokenType) bool {
	for _, t := range types {
		if t == tt {
			return true
		}
	}
	return false
}

func TestBasicTokens(t *testing.T) {
	types := tokenTypes("Get-ChildItem -Force | Where-Object")
	if !hasType(types, TkWord) || !hasType(types, TkDashWord) || !hasType(types, TkPunct) {
		t.Fatalf("basic tokens missing: %v", types)
	}
}

func TestNumbers(t *testing.T) {
	src := "42 3.14 -5 2-1 1.5e3"
	toks := New(src).Tokens()
	want := []float64{42, 3.14, -5}
	// 前三个应为独立数字
	for i, w := range want {
		if toks[i].Type != TkNumber || toks[i].Num != w {
			t.Fatalf("tok[%d] = %+v, want number %v", i, toks[i], w)
		}
	}
	// 2-1 应为 2 和 减法运算符
	if toks[3].Type != TkNumber || toks[3].Num != 2 {
		t.Fatalf("2-1 first: %+v", toks[3])
	}
	if toks[4].Type != TkOp || toks[4].Text != "-" {
		t.Fatalf("2-1 minus: %+v", toks[4])
	}
	if toks[5].Type != TkNumber || toks[5].Num != 1 {
		t.Fatalf("2-1 second: %+v", toks[5])
	}
	// 1.5e3
	if toks[6].Type != TkNumber || toks[6].Num != 1500 {
		t.Fatalf("1.5e3: %+v", toks[6])
	}
}

func TestNegativeVsMinus(t *testing.T) {
	// 行首 -5 是负数
	toks := New("-5").Tokens()
	if toks[0].Type != TkNumber || toks[0].Num != -5 {
		t.Fatalf("-5: %+v", toks[0])
	}
	// 空格后 -5 是负数
	toks = New("x -5").Tokens()
	if toks[1].Type != TkNumber || toks[1].Num != -5 {
		t.Fatalf("x -5: %+v", toks[1])
	}
	// $x-1 是减法
	toks = New("$x-1").Tokens()
	if toks[1].Type != TkOp || toks[1].Text != "-" {
		t.Fatalf("$x-1 minus: %+v", toks[1])
	}
}

func TestStrings(t *testing.T) {
	// 单引号字面
	toks := New("'hello'").Tokens()
	if toks[0].Type != TkString || toks[0].Parts[0].Text != "hello" {
		t.Fatalf("single quote: %+v", toks[0])
	}
	// 单引号内 '' 转义
	toks = New("'it''s'").Tokens()
	if toks[0].Parts[0].Text != "it's" {
		t.Fatalf("single quote escape: %+v", toks[0])
	}
	// 双引号变量展开分段
	toks = New(`"hi $name $env:PATH"`).Tokens()
	parts := toks[0].Parts
	if len(parts) != 4 {
		t.Fatalf("double quote parts = %d: %+v", len(parts), parts)
	}
	if parts[0].Text != "hi " || parts[0].Kind != PartLit {
		t.Fatalf("part0: %+v", parts[0])
	}
	if parts[1].Kind != PartVar || parts[1].Text != "name" {
		t.Fatalf("part1: %+v", parts[1])
	}
	if parts[2].Kind != PartLit || parts[2].Text != " " {
		t.Fatalf("part2: %+v", parts[2])
	}
	if parts[3].Kind != PartEnvVar || parts[3].Text != "PATH" {
		t.Fatalf("part3: %+v", parts[3])
	}
	// 反引号转义：源文本 "a`"b`tc`n" → a"b<TAB>c<LF>
	src := "\"a" + "`\"" + "b" + "`" + "tc" + "`" + "n\""
	toks = New(src).Tokens()
	if toks[0].Parts[0].Text != "a\"b\tc\n" {
		t.Fatalf("backtick escape: %q", toks[0].Parts[0].Text)
	}
	// 子表达式：源文本 "x$($a + 1)y"
	src = "\"x$($a + 1)y\""
	toks = New(src).Tokens()
	parts = toks[0].Parts
	if len(parts) != 3 || parts[1].Kind != PartSubexpr || parts[1].Text != "$a + 1" {
		t.Fatalf("subexpr: %+v", parts)
	}
}

func TestVariables(t *testing.T) {
	toks := New("$foo ${bar baz} $env:Path").Tokens()
	if toks[0].Type != TkVariable || toks[0].Text != "foo" {
		t.Fatalf("$foo: %+v", toks[0])
	}
	if toks[1].Type != TkBraceVar || toks[1].Text != "bar baz" {
		t.Fatalf("${bar baz}: %+v", toks[1])
	}
	if toks[2].Type != TkVariable || toks[2].Text != "env:Path" {
		t.Fatalf("$env:Path: %+v", toks[2])
	}
}

func TestDashWords(t *testing.T) {
	toks := New("-Force -eq -Name:Value").Tokens()
	if toks[0].Type != TkDashWord || toks[0].Text != "Force" {
		t.Fatalf("-Force: %+v", toks[0])
	}
	if toks[1].Type != TkDashWord || toks[1].Text != "eq" {
		t.Fatalf("-eq: %+v", toks[1])
	}
	// -Name:Value → DashWord(Name:Value)，解析器按 ':' 拆分
	if toks[2].Type != TkDashWord || toks[2].Text != "Name:Value" {
		t.Fatalf("-Name:Value: %+v", toks[2])
	}
}

func TestWordsAndPaths(t *testing.T) {
	// 路径作为裸字
	toks := New("./foo ../bar C:\\x .").Tokens()
	if toks[0].Type != TkWord || toks[0].Text != "./foo" {
		t.Fatalf("./foo: %+v", toks[0])
	}
	if toks[1].Type != TkWord || toks[1].Text != "../bar" {
		t.Fatalf("../bar: %+v", toks[1])
	}
	if toks[2].Type != TkWord || toks[2].Text != "C:\\x" {
		t.Fatalf("C:\\x: %+v", toks[2])
	}
	// 单独的 . 是 TkDot
	if toks[3].Type != TkDot {
		t.Fatalf("dot: %+v", toks[3])
	}
	// 通配符
	toks = New("*.txt * ?").Tokens()
	if toks[0].Type != TkWord || toks[0].Text != "*.txt" {
		t.Fatalf("*.txt: %+v", toks[0])
	}
	if toks[1].Type != TkWord || toks[1].Text != "*" {
		t.Fatalf("*: %+v", toks[1])
	}
	if toks[2].Type != TkWord || toks[2].Text != "?" {
		t.Fatalf("?: %+v", toks[2])
	}
	// 属性访问 $x.Length
	toks = New("$x.Length").Tokens()
	if toks[0].Type != TkVariable || toks[1].Type != TkDot || toks[2].Type != TkWord || toks[2].Text != "Length" {
		t.Fatalf("$x.Length: %+v", toks)
	}
}

func TestOperators(t *testing.T) {
	toks := New("1 + 2 * 3 / 4 % 2").Tokens()
	if toks[1].Type != TkOp || toks[1].Text != "+" {
		t.Fatalf("plus: %+v", toks[1])
	}
	if toks[3].Type != TkOp || toks[3].Text != "*" {
		t.Fatalf("star: %+v", toks[3])
	}
	// 赋值运算符
	toks = New("$x += 1").Tokens()
	if toks[1].Type != TkOp || toks[1].Text != "+=" {
		t.Fatalf("+=: %+v", toks[1])
	}
	// 重定向
	toks = New("> >> 2>").Tokens()
	if toks[0].Text != ">" || toks[1].Text != ">>" {
		t.Fatalf("redirect: %+v", toks)
	}
	// 2> 拆成 Number(2) 与 Op(>)，且相邻
	if toks[2].Type != TkNumber || toks[2].Num != 2 {
		t.Fatalf("2>: %+v", toks[2])
	}
	if toks[3].Type != TkOp || toks[3].Text != ">" || !toks[3].Adjacent {
		t.Fatalf("2> op: %+v", toks[3])
	}
}

func TestCommentAndNewline(t *testing.T) {
	src := "# 注释\nGet-Content a # 行尾注释\n1"
	toks := New(src).Tokens()
	// 注释不产出 token，出现两个换行和一个 Word、一个数字
	var newlines int
	for _, tk := range toks {
		if tk.Type == TkNewline {
			newlines++
		}
	}
	if newlines != 2 {
		t.Fatalf("newlines = %d, want 2", newlines)
	}
	if !hasType(tokenTypes(src), TkNumber) {
		t.Fatal("number missing")
	}
}

func TestAtForms(t *testing.T) {
	toks := New("@(1,2) @{a=1}").Tokens()
	if toks[0].Type != TkPunct || toks[0].Text != "@" {
		t.Fatalf("@(: %+v", toks[0])
	}
	if toks[1].Type != TkPunct || toks[1].Text != "(" {
		t.Fatalf("(: %+v", toks[1])
	}
	// @{ 的 @ 与 {
	if toks[6].Type != TkPunct || toks[6].Text != "@" {
		t.Fatalf("@{: %+v", toks[6])
	}
	if toks[7].Type != TkPunct || toks[7].Text != "{" {
		t.Fatalf("{: %+v", toks[7])
	}
}

func TestNullCoalesceToken(t *testing.T) {
	// ?? 连写是运算符
	toks := New(`$null ?? "d"`).Tokens()
	if toks[1].Type != TkOp || toks[1].Text != "??" {
		t.Fatalf("?? token: %+v", toks[1])
	}
	// 单个 ? 仍是裸字（Where-Object 别名 / 三元由解析器区分）
	toks = New("? { $_ }").Tokens()
	if toks[0].Type != TkWord || toks[0].Text != "?" {
		t.Fatalf("? word: %+v", toks[0])
	}
}

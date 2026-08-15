// Package lexer 实现 PowerShell 词法分析。
//
// 把 PowerShell 源码切成一串 Token，供解析器使用。支持：
//   - 裸字/标识符、数字
//   - 单引号字面串、双引号可展开串（含 $var / $env: / 反引号转义 / $() 子表达式）
//   - $var、${var}、$env:Name
//   - -xxx 参数/比较运算符（解析器按上下文区分）
//   - 算术/赋值运算符、标点、# 注释、换行
package lexer

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// TokenType 是词法单元的种类。
type TokenType int

const (
	TkEOF      TokenType = iota // 输入结束
	TkNewline                   // 换行
	TkWord                      // 裸字/标识符（命令名、位置实参、属性名等）
	TkNumber                    // 数字
	TkString                    // 字符串（单/双引号），内容在 Parts 中
	TkVariable                  // $name 或 $env:Name，Text 为名字
	TkBraceVar                  // ${name}，Text 为大括号内名字
	TkDashWord                  // -xxx，Text 不含前导 '-' 与可能的值部分（'-Name:' 时 Text 为 Name）
	TkColon                     // ':'
	TkDot                       // '.'（属性访问）
	TkOp                        // 算术/赋值/重定向运算符，Text 为其符号
	TkPunct                     // 其它标点：| ; , ( ) { } [ ] @
)

// PartKind 是双引号串内部"分段"的种类。
type PartKind int

const (
	PartLit      PartKind = iota // 字面文本
	PartVar                      // $name
	PartEnvVar                   // $env:Name
	PartBraceVar                 // ${name}
	PartSubexpr                  // $( ... ) 子表达式，Text 为内部原始文本
)

// StringPart 描述双引号串中的一段内容。
type StringPart struct {
	Kind PartKind
	Text string
}

// Token 是一个词法单元。
type Token struct {
	Type     TokenType
	Text     string       // 标识符名/运算符符号/变量名
	Num      float64      // 数字值
	IsInt    bool         // 数字是否为整数
	Parts    []StringPart // 字符串分段
	Adjacent bool         // 紧贴前一 Token（中间无空白），用于识别 '2>' 这类
	Line     int
	Col      int
	Pos      int    // 字节偏移（token 起始）
	Raw      string // 原始文本（含引号等），用于外部命令重建
}

// comparisonOps 是会被当作比较/逻辑运算符的 -xxx 集合。
// 解析器据此把 TkDashWord 判定为运算符而非参数。
var comparisonOps = map[string]bool{
	"-eq": true, "-ne": true, "-lt": true, "-gt": true, "-le": true,
	"-ge": true, "-like": true, "-notlike": true, "-match": true,
	"-notmatch": true, "-contains": true, "-notcontains": true,
	"-in": true, "-notin": true, "-is": true, "-isnot": true,
	"-as": true, "-shl": true, "-shr": true, "-band": true,
	"-bor": true, "-bxor": true, "-and": true, "-or": true,
	"-xor": true, "-not": true, "-replace": true, "-join": true,
	"-split": true,
	// 格式运算符 -f（二元运算符，识别路径与比较运算一致；无 cmdlet 用 -f 作参数名）
	"-f": true,
	// 大小写敏感变体（默认大小写不敏感）
	"-ceq": true, "-cne": true, "-clt": true, "-cgt": true, "-cle": true,
	"-cge": true, "-clike": true, "-cnotlike": true, "-cmatch": true,
	"-cnotmatch": true, "-ccontains": true, "-cnotcontains": true,
}

// IsComparisonOp 报告文本（含前导 '-'）是否为比较/逻辑运算符。
func IsComparisonOp(text string) bool { return comparisonOps[text] }

// Lexer 是词法分析器。
type Lexer struct {
	src      string
	pos      int
	line     int
	col      int
	tokStart int       // 当前 token 起始字节偏移（跳过空白后）
	prevType TokenType // 前一 token 类型（用于消解 *、% 的歧义）
	prevText string    // 前一 token 文本
}

// New 创建词法分析器。
func New(src string) *Lexer {
	return &Lexer{src: src, pos: 0, line: 1, col: 1}
}

func (l *Lexer) peek() byte {
	if l.pos >= len(l.src) {
		return 0
	}
	return l.src[l.pos]
}

func (l *Lexer) peekAt(off int) byte {
	if l.pos+off >= len(l.src) {
		return 0
	}
	return l.src[l.pos+off]
}

func (l *Lexer) next() byte {
	if l.pos >= len(l.src) {
		return 0
	}
	c := l.src[l.pos]
	l.pos++
	if c == '\n' {
		l.line++
		l.col = 1
	} else {
		l.col++
	}
	return c
}

// isWordChar 判断字符是否为裸字组成部分（含 - / \ : . * ? % ~ 等）。
func isWordChar(c byte) bool {
	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' {
		return true
	}
	switch c {
	case '_', '-', '/', '\\', ':', '.', '~', '*', '?', '%':
		return true
	}
	// 支持 Unicode 字母（如中文变量名/路径）
	return c >= 0x80
}

func isDigit(c byte) bool { return c >= '0' && c <= '9' }
func isLetter(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= 0x80
}

// nextIsLetter 判断当前位置开始的是否为 Unicode 字母（处理多字节）。
func (l *Lexer) nextIsLetter() bool {
	c := l.peek()
	if c >= 0x80 {
		r, _ := utf8.DecodeRuneInString(l.src[l.pos:])
		return unicode.IsLetter(r)
	}
	return isLetter(c)
}

// Tokens 一次性产出全部 Token。
func (l *Lexer) Tokens() []Token {
	var out []Token
	for {
		tok := l.nextTok()
		out = append(out, tok)
		if tok.Type == TkEOF {
			return out
		}
	}
}

// nextTok 产出下一个 Token，记录原始文本、起始偏移与前 token 类型。
func (l *Lexer) nextTok() Token {
	t := l.nextTok0()
	t.Raw = l.src[l.tokStart:l.pos]
	t.Pos = l.tokStart
	l.prevType = t.Type
	l.prevText = t.Text
	return t
}

// isAdjacent 报告当前位置是否紧贴前一 token（前一字符非空白）。
func (l *Lexer) isAdjacent() bool {
	if l.pos == 0 {
		return false
	}
	c := l.src[l.pos-1]
	return !(c == ' ' || c == '\t' || c == '\n' || c == '\r')
}

func (l *Lexer) nextTok0() Token {
	startLine, startCol := l.line, l.col

	// 跳过空白（空格/制表/回车），不跳过换行
	for {
		c := l.peek()
		if c == ' ' || c == '\t' || c == '\r' {
			l.next()
		} else {
			break
		}
	}

	// 相邻性：跳过空白后，token 起始位置的前一字符是否为空白
	adj := l.isAdjacent()
	l.tokStart = l.pos

	c := l.peek()
	if c == 0 {
		return Token{Type: TkEOF, Line: startLine, Col: startCol}
	}

	// 换行
	if c == '\n' {
		l.next()
		return Token{Type: TkNewline, Text: "\n", Line: startLine, Col: startCol, Adjacent: adj}
	}

	// 注释：'#' 到行尾（若 # 是裸字的一部分如 a#b，则由裸字扫描吃掉）
	if c == '#' {
		for l.peek() != 0 && l.peek() != '\n' {
			l.next()
		}
		return l.nextTok() // 递归跳过注释（不产出 token）
	}

	// 反引号：行内续行或转义（裸上下文下忽略）
	if c == '`' {
		l.next()
		if l.peek() == '\n' {
			l.next()
			return l.nextTok()
		}
		// 孤立的反引号：跳过（当作无意义字符）
		return l.nextTok()
	}

	// 数字
	if isDigit(c) || (c == '.' && isDigit(l.peekAt(1))) {
		return l.lexNumber(adj)
	}

	// 变量
	if c == '$' {
		return l.lexVariable(adj)
	}

	// 字符串
	if c == '\'' || c == '"' {
		return l.lexString(c, adj)
	}

	// -xxx 参数 / 运算符 / 负数
	if c == '-' {
		return l.lexDash(adj)
	}

	// 操作符
	if op := l.lexOperator(adj); op.Type != TkEOF {
		return op
	}

	// '.' 单独处理：属性访问点号；但 './'、'..'、'.\' 等路径前缀走裸字
	if c == '.' {
		nc := l.peekAt(1)
		if nc == '/' || nc == '\\' || nc == '.' {
			return l.lexWord(adj)
		}
		l.next()
		return Token{Type: TkDot, Text: ".", Line: startLine, Col: startCol, Adjacent: adj}
	}

	// 空合并运算符 '??'：两个问号连写 → 运算符；单个 '?' 仍走裸字（三元/Where-Object 别名由解析器区分）
	if c == '?' && l.peekAt(1) == '?' {
		l.next()
		l.next()
		return Token{Type: TkOp, Text: "??", Line: startLine, Col: startCol, Adjacent: adj}
	}

	// 裸字（含 $ 之外的一切字面符号：路径、命令名、中文等）
	if isWordChar(c) || unicode.IsLetter(rune(c)) {
		return l.lexWord(adj)
	}

	// 其它标点
	l.next()
	tt := TkPunct
	switch c {
	case '|':
		if l.peek() == '|' { // || 逻辑或链
			l.next()
			tt = TkOp
			return Token{Type: TkOp, Text: "||", Line: startLine, Col: startCol, Adjacent: adj}
		}
	case ';', ',', '(', ')', '{', '}', '[', ']', ':':
	case '@':
		// @( 与 @{ 与 @ 单独
		tt = TkPunct
	case '!':
		tt = TkOp
	}
	return Token{Type: tt, Text: string(c), Line: startLine, Col: startCol, Adjacent: adj}
}

func (l *Lexer) lexNumber(adj bool) Token {
	startLine, startCol := l.line, l.col
	var sb strings.Builder
	neg := false
	if l.peek() == '-' {
		neg = true
		sb.WriteByte(l.next())
	}
	isInt := true
	for {
		c := l.peek()
		if isDigit(c) {
			sb.WriteByte(l.next())
		} else if c == '.' && isDigit(l.peekAt(1)) && !strings.Contains(sb.String(), ".") {
			sb.WriteByte(l.next())
			isInt = false
		} else if c == 'e' || c == 'E' {
			// 科学计数法
			sb.WriteByte(l.next())
			if l.peek() == '+' || l.peek() == '-' {
				sb.WriteByte(l.next())
			}
			isInt = false
		} else {
			break
		}
	}
	s := sb.String()
	if neg && len(s) == 1 { // 只有负号，异常；回退为 '-'
		return Token{Type: TkOp, Text: "-", Line: startLine, Col: startCol, Adjacent: adj}
	}
	var f float64
	var isIntF bool
	if isInt {
		// 手动解析避免 strconv 对超大整数的限制
		f = parseInt(s)
		isIntF = true
	} else {
		f = parseFloat(s)
	}
	return Token{Type: TkNumber, Text: s, Num: f, IsInt: isIntF, Line: startLine, Col: startCol, Adjacent: adj}
}

func parseInt(s string) float64 {
	var v float64
	neg := false
	i := 0
	if i < len(s) && s[i] == '-' {
		neg = true
		i++
	}
	for ; i < len(s); i++ {
		v = v*10 + float64(s[i]-'0')
	}
	if neg {
		v = -v
	}
	return v
}

func parseFloat(s string) float64 {
	var v float64
	neg := false
	i := 0
	if i < len(s) && s[i] == '-' {
		neg = true
		i++
	}
	intPart := 0.0
	for i < len(s) && isDigit(s[i]) {
		intPart = intPart*10 + float64(s[i]-'0')
		i++
	}
	frac := 0.0
	if i < len(s) && s[i] == '.' {
		i++
		scale := 0.1
		for i < len(s) && isDigit(s[i]) {
			frac += float64(s[i]-'0') * scale
			scale /= 10
			i++
		}
	}
	v = intPart + frac
	// 指数
	if i < len(s) && (s[i] == 'e' || s[i] == 'E') {
		i++
		sign := 1.0
		if i < len(s) && (s[i] == '+' || s[i] == '-') {
			if s[i] == '-' {
				sign = -1
			}
			i++
		}
		exp := 0.0
		for i < len(s) && isDigit(s[i]) {
			exp = exp*10 + float64(s[i]-'0')
			i++
		}
		pow := 1.0
		for j := 0.0; j < exp; j++ {
			pow *= 10
		}
		if sign < 0 {
			v /= pow
		} else {
			v *= pow
		}
	}
	if neg {
		v = -v
	}
	return v
}

func (l *Lexer) lexVariable(adj bool) Token {
	startLine, startCol := l.line, l.col
	l.next()             // 吃掉 '$'
	if l.peek() == '?' { // $? 上次命令成功与否
		l.next()
		return Token{Type: TkVariable, Text: "?", Line: startLine, Col: startCol, Adjacent: adj}
	}
	if l.peek() == '{' {
		l.next()
		var sb strings.Builder
		for l.peek() != 0 && l.peek() != '}' {
			sb.WriteByte(l.next())
		}
		if l.peek() == '}' {
			l.next()
		}
		return Token{Type: TkBraceVar, Text: sb.String(), Line: startLine, Col: startCol, Adjacent: adj}
	}
	var sb strings.Builder
	for {
		c := l.peek()
		if l.nextIsLetter() || isDigit(c) || c == '_' || c == ':' {
			sb.WriteByte(l.next())
		} else {
			break
		}
	}
	return Token{Type: TkVariable, Text: sb.String(), Line: startLine, Col: startCol, Adjacent: adj}
}

func (l *Lexer) lexString(quote byte, adj bool) Token {
	startLine, startCol := l.line, l.col
	l.next() // 吃掉引号
	if quote == '\'' {
		// 单引号字面串，'' 转义为 '
		var sb strings.Builder
		for {
			c := l.next()
			if c == 0 {
				break // 未闭合，容错
			}
			if c == '\'' {
				if l.peek() == '\'' {
					sb.WriteByte('\'')
					l.next()
					continue
				}
				break
			}
			sb.WriteByte(c)
		}
		return Token{Type: TkString, Parts: []StringPart{{Kind: PartLit, Text: sb.String()}}, Line: startLine, Col: startCol, Adjacent: adj}
	}
	// 双引号可展开串
	parts := l.lexDoubleQuoted()
	return Token{Type: TkString, Parts: parts, Line: startLine, Col: startCol, Adjacent: adj}
}

func (l *Lexer) lexDoubleQuoted() []StringPart {
	var parts []StringPart
	var lit strings.Builder
	flush := func() {
		if lit.Len() > 0 {
			parts = append(parts, StringPart{Kind: PartLit, Text: lit.String()})
			lit.Reset()
		}
	}
	for {
		c := l.next()
		if c == 0 {
			break
		}
		switch c {
		case '"':
			flush()
			return parts
		case '`':
			e := l.next()
			switch e {
			case 'n':
				lit.WriteByte('\n')
			case 't':
				lit.WriteByte('\t')
			case 'r':
				lit.WriteByte('\r')
			case 'a':
				lit.WriteByte('\a')
			case '0':
				lit.WriteByte(0)
			case '`':
				lit.WriteByte('`')
			case '$':
				lit.WriteByte('$')
			case '"':
				lit.WriteByte('"')
			case '\'':
				lit.WriteByte('\'')
			case 0:
			default:
				// 未知转义：保留原字符
				lit.WriteByte(e)
			}
		case '$':
			flush()
			// 子表达式 $(...)
			if l.peek() == '(' {
				l.next()
				depth := 1
				var sb strings.Builder
				for l.peek() != 0 {
					ch := l.next()
					if ch == '(' {
						depth++
					} else if ch == ')' {
						depth--
						if depth == 0 {
							break
						}
					}
					sb.WriteByte(ch)
				}
				parts = append(parts, StringPart{Kind: PartSubexpr, Text: sb.String()})
				continue
			}
			// 变量
			if l.peek() == '{' {
				l.next()
				var sb strings.Builder
				for l.peek() != 0 && l.peek() != '}' {
					sb.WriteByte(l.next())
				}
				if l.peek() == '}' {
					l.next()
				}
				parts = append(parts, StringPart{Kind: PartBraceVar, Text: sb.String()})
				continue
			}
			var sb strings.Builder
			for {
				ch := l.peek()
				if l.nextIsLetter() || isDigit(ch) || ch == '_' || ch == ':' {
					sb.WriteByte(l.next())
				} else {
					break
				}
			}
			if sb.Len() > 0 {
				name := sb.String()
				if strings.HasPrefix(name, "env:") {
					parts = append(parts, StringPart{Kind: PartEnvVar, Text: name[len("env:"):]})
				} else {
					parts = append(parts, StringPart{Kind: PartVar, Text: name})
				}
			} else {
				lit.WriteByte('$')
			}
		default:
			lit.WriteByte(c)
		}
	}
	flush()
	return parts
}

func (l *Lexer) lexDash(adj bool) Token {
	startLine, startCol := l.line, l.col
	// '--'：后跟字母则视为单词（如 --force），否则为递减运算符
	if l.peekAt(1) == '-' {
		l.next()
		l.next()
		if isLetter(l.peek()) {
			var sb strings.Builder
			sb.WriteString("--")
			for {
				c := l.peek()
				if isWordChar(c) || unicode.IsLetter(rune(c)) {
					sb.WriteByte(l.next())
				} else {
					break
				}
			}
			return Token{Type: TkWord, Text: sb.String(), Line: startLine, Col: startCol, Adjacent: adj}
		}
		return Token{Type: TkOp, Text: "--", Line: startLine, Col: startCol, Adjacent: adj}
	}
	// '-' 后紧跟数字或 '.'：
	//   紧贴前一值 token（如 '2-1'、'$x-1'）→ 减法运算符
	//   否则（如 ' -1'、行首 '-1'）→ 负数
	if isDigit(l.peekAt(1)) || l.peekAt(1) == '.' {
		if adj {
			l.next()
			return Token{Type: TkOp, Text: "-", Line: startLine, Col: startCol, Adjacent: adj}
		}
		return l.lexNumber(adj)
	}
	if isLetter(l.peekAt(1)) || l.peekAt(1) == '?' {
		l.next() // 吃掉 '-'
		var sb strings.Builder
		for {
			c := l.peek()
			if isLetter(c) || isDigit(c) || c == '_' || c == '?' || c == '-' || c == ':' {
				sb.WriteByte(l.next())
			} else {
				break
			}
		}
		return Token{Type: TkDashWord, Text: sb.String(), Line: startLine, Col: startCol, Adjacent: adj}
	}
	// 否则是减号运算符
	l.next()
	return Token{Type: TkOp, Text: "-", Line: startLine, Col: startCol, Adjacent: adj}
}

// endsValue 判断前一 token 是否以"值"结尾（即 * / % 此时应视为运算符而非通配符）。
func (l *Lexer) endsValue(adj bool) bool {
	switch l.prevType {
	case TkNumber, TkString, TkVariable, TkBraceVar:
		return true
	case TkPunct:
		return l.prevText == ")" || l.prevText == "]" || l.prevText == "}"
	case TkWord:
		// 裸字后：若紧跟无空白则多为乘法（如 a*3），保守按值结尾
		return adj
	}
	return false
}

func (l *Lexer) lexOperator(adj bool) Token {
	startLine, startCol := l.line, l.col
	c := l.peek()
	switch c {
	case '=':
		l.next()
		if l.peek() == '=' { // '==' 少见，保留
			l.next()
			return Token{Type: TkOp, Text: "==", Line: startLine, Col: startCol, Adjacent: adj}
		}
		return Token{Type: TkOp, Text: "=", Line: startLine, Col: startCol, Adjacent: adj}
	case '+':
		l.next()
		if l.peek() == '+' { // 增量运算符
			l.next()
			return Token{Type: TkOp, Text: "++", Line: startLine, Col: startCol, Adjacent: adj}
		}
		if l.peek() == '=' {
			l.next()
			return Token{Type: TkOp, Text: "+=", Line: startLine, Col: startCol, Adjacent: adj}
		}
		return Token{Type: TkOp, Text: "+", Line: startLine, Col: startCol, Adjacent: adj}
	case '*':
		// '*'：前一 token 是值（如 2 * 3）或后跟数字（2*3）→ 运算符；否则是通配符裸字
		if isDigit(l.peekAt(1)) || l.endsValue(adj) {
			l.next()
			if l.peek() == '=' {
				l.next()
				return Token{Type: TkOp, Text: "*=", Line: startLine, Col: startCol, Adjacent: adj}
			}
			return Token{Type: TkOp, Text: "*", Line: startLine, Col: startCol, Adjacent: adj}
		}
		return l.lexWord(adj)
	case '/':
		// '/' 后跟字母/点/等 → 路径裸字；否则除法
		nc := l.peekAt(1)
		if isLetter(nc) || nc == '.' || isDigit(nc) {
			return l.lexWord(adj)
		}
		l.next()
		if l.peek() == '=' {
			l.next()
			return Token{Type: TkOp, Text: "/=", Line: startLine, Col: startCol, Adjacent: adj}
		}
		return Token{Type: TkOp, Text: "/", Line: startLine, Col: startCol, Adjacent: adj}
	case '%':
		// '%'：前一 token 是值（如 5 % 2）或后跟数字 → 模运算；否则是别名/通配符裸字
		if isDigit(l.peekAt(1)) || l.endsValue(adj) {
			l.next()
			if l.peek() == '=' {
				l.next()
				return Token{Type: TkOp, Text: "%=", Line: startLine, Col: startCol, Adjacent: adj}
			}
			return Token{Type: TkOp, Text: "%", Line: startLine, Col: startCol, Adjacent: adj}
		}
		return l.lexWord(adj)
	case '>':
		l.next()
		if l.peek() == '>' {
			l.next()
			return Token{Type: TkOp, Text: ">>", Line: startLine, Col: startCol, Adjacent: adj}
		}
		return Token{Type: TkOp, Text: ">", Line: startLine, Col: startCol, Adjacent: adj}
	case '<':
		l.next()
		if l.peek() == '<' {
			l.next()
		}
		return Token{Type: TkOp, Text: "<", Line: startLine, Col: startCol, Adjacent: adj}
	case '&':
		l.next()
		if l.peek() == '&' { // && 逻辑与链
			l.next()
			return Token{Type: TkOp, Text: "&&", Line: startLine, Col: startCol, Adjacent: adj}
		}
		return Token{Type: TkOp, Text: "&", Line: startLine, Col: startCol, Adjacent: adj}
	}
	return Token{Type: TkEOF}
}

func (l *Lexer) lexWord(adj bool) Token {
	startLine, startCol := l.line, l.col
	// ".." 后紧跟数字（如 1..5）→ 只取 ".." 作为范围运算符
	if l.peek() == '.' && l.peekAt(1) == '.' && isDigit(l.peekAt(2)) {
		l.next()
		l.next()
		return Token{Type: TkWord, Text: "..", Line: startLine, Col: startCol, Adjacent: adj}
	}
	var sb strings.Builder
	for {
		c := l.peek()
		if c == 0 {
			break
		}
		if isWordChar(c) || unicode.IsLetter(rune(c)) {
			// 数字开头的字面（如 '1.5x'）整体作为裸字保留
			sb.WriteByte(l.next())
		} else if c == '#' {
			// a#b 中 # 作为裸字一部分
			sb.WriteByte(l.next())
		} else {
			break
		}
	}
	return Token{Type: TkWord, Text: sb.String(), Line: startLine, Col: startCol, Adjacent: adj}
}

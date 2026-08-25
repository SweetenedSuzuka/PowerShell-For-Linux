// Package ast 定义 PowerShell 脚本的抽象语法树节点。
package ast

// Node 是所有 AST 节点的接口。
type Node interface{}

// ---- 语句层 ----

// StatementList 是一串语句（脚本主体或脚本块体）。
type StatementList struct {
	Statements []Node
}

// Pipeline 是管道：由 '|' 分隔的命令序列。
// Expr 非空表示管道以纯表达式开头（如 1,2,3 | ForEach-Object {...}）。
type Pipeline struct {
	Commands []*Command
	Expr     Node
}

// Chain 是管道链（PowerShell 7 的 && / ||）：左侧成功/失败决定是否执行右侧。
type Chain struct {
	Left  Node
	Right Node
	Op    string // && 或 ||
}

// RedirKind 是重定向的种类。
type RedirKind int

const (
	RedirStdout RedirKind = iota // >
	RedirAppend                  // >>
	RedirStderr                  // 2>
)

// Redirection 是一条重定向。
type Redirection struct {
	Kind   RedirKind
	Target Node
	Append bool // 追加（>>）
}

// NamedArg 是命名参数 -Name Value 或 -Name:Value。
type NamedArg struct {
	Name  string
	Value Node // 可为 nil（-Name 后无值）
}

// ArgKind 是命令参数的种类（用于按源码顺序重建外部命令 argv）。
type ArgKind int

const (
	ArgPositional ArgKind = iota
	ArgNamed
	ArgSwitch
)

// ArgItem 描述命令参数在源码中的顺序。
type ArgItem struct {
	Kind   ArgKind
	Name   string // 命名参数/开关的名字
	Index  int    // 在对应子列表中的下标
	Inline bool   // -Name:value / -Name=value 内联形式（开关用它做布尔值）
}

// Command 是命令调用（cmdlet/别名/函数/外部程序）。
type Command struct {
	Name       string
	Positional []Node
	Named      []NamedArg
	Switches   []string
	Redirs     []Redirection
	ArgOrder   []ArgItem // 参数出现顺序（外部命令传参用）
	RawParts   []string  // 各 token 原始文本（外部命令重建备用）
}

// Assign 是赋值：$x = 5、$env:X = "y"、$x += 1、$script:x = 1。
type Assign struct {
	Target string // 变量名，可为 "env:Name" 或 "$x"（去掉 $ 与 ${}）
	Scope  string // 作用域修饰符："" / "script" / "global" / "local"（$env: 单独走 Target）
	Op     string // =、+=、-=、*=、/=、%=
	Value  Node
}

// IfBranch 是 if 的一个分支（条件 + 主体）。
type IfBranch struct {
	Cond Node
	Body *Block
}

// If 是 if/elseif/else 语句。
type If struct {
	Branches []IfBranch
	Else     *Block
}

// ForEach 是 foreach ($x in $coll) { ... } 语句。
type ForEach struct {
	Var  string
	Coll Node
	Body *Block
}

// While 是 while 循环。
type While struct {
	Cond Node
	Body *Block
}

// DoWhile 是 do { ... } while (cond) 循环。
type DoWhile struct {
	Cond Node
	Body *Block
}

// For 是 for (init; cond; post) { ... } 循环。
type For struct {
	Init Node // 赋值语句或 nil
	Cond Node
	Post Node
	Body *Block
}

// SwitchCase 是 switch 的一个分支。
type SwitchCase struct {
	Cond Node // 匹配值；nil 表示 default
	Body *Block
}

// Switch 是 switch (expr) { ... } 语句。
type Switch struct {
	Value Node
	Cases []SwitchCase
	Mode  string // exact / regex / wildcard / case
}

// Break 是 break 语句。
type Break struct{}

// Continue 是 continue 语句。
type Continue struct{}

// Return 是 return [expr] 语句。
type Return struct {
	Value Node
}

// Exit 是 exit [code] 语句。
type Exit struct {
	Code Node
}

// FunctionParam 是函数参数（含默认值）。
type FunctionParam struct {
	Name    string
	Default Node
}

// ParamBlock 是 param(...) 参数声明块，只能出现在脚本或函数体开头。
type ParamBlock struct {
	Params []FunctionParam
}

// CatchClause 是 try 的一个 catch 分支。
type CatchClause struct {
	TypeName string // 可选 [类型] 过滤，空为全部捕获
	Body     *Block
}

// Try 是 try/catch/finally 语句。
type Try struct {
	Body    *Block
	Catches []CatchClause
	Finally *Block
}

// Throw 是 throw 语句。
type Throw struct {
	Value Node
}

// FunctionDef 是 function / filter 定义。
type FunctionDef struct {
	Name   string
	Params []FunctionParam
	Body   *Block
	Filter bool
}

// Block 是 { ... } 语句块。
type Block struct {
	Body *StatementList
}

// ---- 表达式层 ----

// StrLit 是单引号字面串。
type StrLit struct{ Value string }

// StrTemplate 是双引号可展开串，Parts 可为 StrLit/VarRef/EnvRef/任意表达式。
type StrTemplate struct{ Parts []Node }

// VarRef 是变量引用（$x / ${x}）。
type VarRef struct {
	Name  string
	Scope string // 作用域修饰符："" / "script" / "global" / "local"
}

// EnvRef 是环境变量引用（$env:Name）。
type EnvRef struct{ Name string }

// Number 是数字字面量。
type Number struct {
	Value float64
	IsInt bool
}

// BoolLit 是 $true / $false。
type BoolLit struct{ Value bool }

// NullLit 是 $null。
type NullLit struct{}

// ArrayLit 是数组 @(...) / ,(...) / 逗号表达式。
// Flatten 标记来自 @(...)：元素按输出流摊平（@($arr) 得到元素本身）；逗号列表不摊平（1,2,(3,4) 第 3 项是内嵌数组）。
type ArrayLit struct {
	Items   []Node
	Flatten bool
}

// HashPair 是哈希表的一个键值对。
type HashPair struct {
	Key   Node
	Value Node
}

// HashtableLit 是哈希表字面量 @{ k = v; ... }。
type HashtableLit struct{ Pairs []HashPair }

// TypeCast 是类型字面量：[pscustomobject]@{...} 构造自定义对象。
type TypeCast struct {
	TypeName string
	Expr     Node
}

// StaticMember 是静态成员访问：[类型]::名称。
// Args 为 nil 表示静态属性；非 nil（含空切片）表示带参的静态方法调用。
type StaticMember struct {
	TypeName string
	Name     string
	Args     []Node
}

// Paren 是括号表达式。
type Paren struct{ Inner Node }

// Unary 是一元运算：-x、-not x、!x。
type Unary struct {
	Op      string
	Operand Node
}

// Increment 是增量/减量运算符（$i++ / $i-- / $script:i++）。
type Increment struct {
	Var   string
	Scope string // 作用域修饰符："" / "script" / "global" / "local"
	Op    string // ++ 或 --
}

// Binary 是二元运算：算术、比较、逻辑。
type Binary struct {
	Op string
	L  Node
	R  Node
}

// Ternary 是三元运算符（PowerShell 7）：cond ? ifTrue : ifFalse。
type Ternary struct {
	Cond Node
	If   Node
	Else Node
}

// PropertyRef 是裸属性引用（Where-Object Length -gt 100 中的 Length）。
type PropertyRef struct{ Name string }

// MemberAccess 是属性访问（$x.Length）。
type MemberAccess struct {
	Base Node
	Prop string
}

// MethodCall 是方法调用（"abc".ToUpper()）。
type MethodCall struct {
	Base Node
	Name string
	Args []Node
}

// Index 是索引访问（$a[0]）。
type Index struct {
	Base  Node
	Index Node
}

// ScriptBlock 是脚本块 { ... }。
type ScriptBlock struct{ Body *StatementList }

// PipelineExpr 是作为表达式出现的管道（$(...) 内部）。
type PipelineExpr struct{ Pipeline *Pipeline }

// SubExpr 是 $( ... ) 子表达式（语句级）。
type SubExpr struct{ Body *StatementList }

// BareWord 是命令参数中的裸字（如 Get-Content a 中的 a）。
type BareWord struct{ Value string }

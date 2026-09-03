package builtin

import (
	"strings"
)

// registry.go 存 cmdlet 注册表（名字到实现与参数规格的映射）。

var registry = map[string]CmdFunc{}
var specMap = map[string][]ParamSpec{}
var displayMap = map[string]string{} // 小写名 → 原始大小写名（Get-Help/Get-Command 显示用）

// Register 注册 cmdlet（名字转小写存储）。
func Register(name string, spec []ParamSpec, fn CmdFunc) {
	lower := strings.ToLower(name)
	registry[lower] = fn
	specMap[lower] = spec
	displayMap[lower] = name
}

// Lookup 查找 cmdlet。
func Lookup(name string) (CmdFunc, bool) {
	fn, ok := registry[strings.ToLower(name)]
	return fn, ok
}

// Spec 返回 cmdlet 的参数规格。
func Spec(name string) []ParamSpec {
	return specMap[strings.ToLower(name)]
}

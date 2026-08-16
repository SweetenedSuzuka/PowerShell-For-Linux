package builtin_test

import (
	"testing"

	"powershell/internal/builtin"
)

// 检查所有注册的 ParamSpec：位置槽位不重复、序号连续、开关不带位置。
func TestSpecCheck(t *testing.T) {
	// 通过 Spec 无法枚举全部命令，直接扫已知可疑点：检查重复 Position 需遍历注册表。
	// 这里检查几个关键命令的槽位顺序。
	check := func(name string, want []string) {
		spec := builtin.Spec(name)
		var got []string
		for _, s := range spec {
			if s.PositionSet {
				got = append(got, s.Name)
			}
		}
		if len(got) != len(want) {
			t.Errorf("%s: 位置槽位 %v，期望 %v", name, got, want)
			return
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("%s: 位置槽位顺序 %v，期望 %v", name, got, want)
				return
			}
		}
	}
	check("Set-Content", []string{"Path", "Value"})
	check("Get-Content", []string{"Path"})
	check("Copy-Item", []string{"Path", "Destination"})
	check("Compare-Object", []string{"ReferenceObject", "DifferenceObject"})
	check("Where-Object", []string{"FilterScript"})
	check("Select-String", []string{"Pattern", "Path"})
	check("Select-Object", []string{"Property"})
	check("Join-Path", []string{"Path", "ChildPath"})
	check("Rename-Item", []string{"Path", "NewName"})
	check("Get-ChildItem", []string{"Path"})
	check("Set-Alias", []string{"Name", "Value"})
}

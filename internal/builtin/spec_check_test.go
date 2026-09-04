package builtin_test

import (
	"testing"

	"powershell/internal/builtin"
)

// 抽查关键命令的位置槽位顺序。
func TestSpecCheck(t *testing.T) {
	// 通过 Spec 无法枚举全部命令，直接扫已知可疑点：检查重复 Position 需遍历注册表。
	// 这里检查几个关键命令的槽位顺序。
	assertSlotOrder := func(name string, expected []string) {
		spec := builtin.Spec(name)
		var got []string
		for _, s := range spec {
			if s.PositionSet {
				got = append(got, s.Name)
			}
		}
		if len(got) != len(expected) {
			t.Errorf("%s: 位置槽位 %v，期望 %v", name, got, expected)
			return
		}
		for i := range got {
			if got[i] != expected[i] {
				t.Errorf("%s: 位置槽位顺序 %v，期望 %v", name, got, expected)
				return
			}
		}
	}
	assertSlotOrder("Set-Content", []string{"Path", "Value"})
	assertSlotOrder("Get-Content", []string{"Path"})
	assertSlotOrder("Copy-Item", []string{"Path", "Destination"})
	assertSlotOrder("Compare-Object", []string{"ReferenceObject", "DifferenceObject"})
	assertSlotOrder("Where-Object", []string{"FilterScript"})
	assertSlotOrder("Select-String", []string{"Pattern", "Path"})
	assertSlotOrder("Select-Object", []string{"Property"})
	assertSlotOrder("Join-Path", []string{"Path", "ChildPath"})
	assertSlotOrder("Rename-Item", []string{"Path", "NewName"})
	assertSlotOrder("Get-ChildItem", []string{"Path"})
	assertSlotOrder("Set-Alias", []string{"Name", "Value"})
}

package builtin

import (
	"crypto/rand"
	"fmt"
	"os"
	"strings"
	"time"

	"powershell/internal/object"
)

// new.go 实现构造类 cmdlet（新建对象与成员添加）。

func cmdAddMember(c *Context) ([]*object.PSObject, error) {
	memberType := "NoteProperty"
	if t, ok := c.Args.Str("MemberType"); ok {
		memberType = t
	}
	name, _ := c.Args.Str("Name")
	val := c.Args.Get("Value")
	force := c.Args.Switch("Force")
	var out []*object.PSObject
	for _, o := range inputItems(c) {
		cp := o.Clone()
		if !strings.EqualFold(memberType, "NoteProperty") {
			// 只支持 NoteProperty
			cp = o
		}
		if name != "" {
			if !cp.HasProp(name) || force {
				if val != nil {
					cp.AddProp(name, val.Value)
				} else {
					cp.AddProp(name, nil)
				}
			}
		}
		out = append(out, cp)
	}
	return out, nil
}

func cmdNewGuid(c *Context) ([]*object.PSObject, error) {
	var b [16]byte
	_, _ = rand.Read(b[:])
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	guid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
	return []*object.PSObject{object.Str(guid)}, nil
}

func cmdNewTimeSpan(c *Context) ([]*object.PSObject, error) {
	d := time.Duration(0)
	if sec, ok := c.Args.Int("Seconds"); ok {
		d += time.Duration(sec) * time.Second
	}
	if mn, ok := c.Args.Int("Minutes"); ok {
		d += time.Duration(mn) * time.Minute
	}
	if h, ok := c.Args.Int("Hours"); ok {
		d += time.Duration(h) * time.Hour
	}
	return []*object.PSObject{timeSpanObj(d)}, nil
}

func cmdNewTemporaryFile(c *Context) ([]*object.PSObject, error) {
	ext := ""
	if e, ok := c.Args.Str("Extension"); ok {
		ext = e
	}
	f, err := os.CreateTemp("", "tmp*"+ext)
	if err != nil {
		return errf(c, "New-TemporaryFile : %v", err)
	}
	f.Close()
	if info, e := os.Stat(f.Name()); e == nil {
		return []*object.PSObject{object.FileInfo(f.Name(), info)}, nil
	}
	return []*object.PSObject{object.Str(f.Name())}, nil
}

// ---- 注册 ----

func init() {
	Register("Add-Member", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
		{Name: "MemberType", Type: "string"},
		{Name: "Name", Type: "string"},
		{Name: "Value", Type: "object"},
		{Name: "Force", Switch: true},
	}, cmdAddMember)
	Register("New-Guid", nil, cmdNewGuid)
	Register("New-TimeSpan", []ParamSpec{
		{Name: "Seconds", Type: "int"},
		{Name: "Minutes", Type: "int"},
		{Name: "Hours", Type: "int"},
	}, cmdNewTimeSpan)
	Register("New-TemporaryFile", []ParamSpec{
		{Name: "Extension", Type: "string"},
	}, cmdNewTemporaryFile)
}

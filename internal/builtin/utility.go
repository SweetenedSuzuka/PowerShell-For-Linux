package builtin

import (
	"context"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"powershell/internal/lang"
	"powershell/internal/object"
)

// utility.go 收纳不成体系的独立小指令（连接测试、暂停）。

func cmdTestConnection(c *Context) ([]*object.PSObject, error) {
	host := firstArg(c, "TargetName")
	if host == "" {
		return nil, nil
	}
	count := "4"
	if n, ok := c.Args.Int("Count"); ok && n > 0 {
		count = strconv.FormatInt(n, 10)
	}
	if _, err := exec.LookPath("ping"); err != nil {
		return errf(c, "%s", lang.T(lang.MsgPingNotFound))
	}
	code := runExternalRaw(c, "ping", []string{"-c", count, host})
	return []*object.PSObject{object.Bool(code == 0)}, nil
}

func cmdStartSleep(c *Context) ([]*object.PSObject, error) {
	// -Seconds 与 -Milliseconds 分属不同参数集，不可同用（按出现与否判定，与取值无关）。
	if c.Args.Get("Seconds") != nil && c.Args.Get("Milliseconds") != nil {
		return errf(c, "%s", lang.T(lang.MsgParamSetUnresolvable))
	}
	// 超量位置实参无槽位可接（Seconds 只占位置 0），报错而非静默忽略。
	if len(c.Args.Positional) > 0 {
		return errf(c, "%s", lang.T(lang.MsgPositionalParamNotFound, c.Args.Positional[0].String()))
	}
	var d time.Duration
	if sec, ok := c.Args.Int("Seconds"); ok {
		d = time.Duration(sec) * time.Second
	}
	if ms, ok := c.Args.Int("Milliseconds"); ok {
		d = time.Duration(ms) * time.Millisecond
	}
	if d > 0 {
		time.Sleep(d)
	}
	return nil, nil
}

// dnsRecord 输出一条解析记录（统一 Name/Type/Data 三列，近似 dig +short 的行）。
func dnsRecord(name, typ, data string) *object.PSObject {
	o := object.Object("DnsRecord", nil)
	o.AddProp("Name", name)
	o.AddProp("Type", typ)
	o.AddProp("Data", data)
	o.Table = []object.Column{
		{Label: "Name", Align: "left"},
		{Label: "Type", Align: "left"},
		{Label: "Data", Align: "left"},
	}
	return o
}

// cmdResolveDnsName 域名解析（Go 内置解析，默认系统解析器，-Server 指定上游；默认查询 A 与 AAAA）。
func cmdResolveDnsName(c *Context) ([]*object.PSObject, error) {
	// 超量位置实参无槽位可接（Name 占位置 0，Type 占位置 1），报错而非静默忽略。
	if len(c.Args.Positional) > 0 {
		return errf(c, "%s", lang.T(lang.MsgPositionalParamNotFound, c.Args.Positional[0].String()))
	}
	typ, _ := c.Args.Str("Type")
	typ = strings.ToUpper(strings.TrimSpace(typ))
	switch typ {
	case "", "A", "AAAA", "CNAME", "MX", "TXT", "NS", "PTR":
	default:
		return errf(c, "%s", lang.T(lang.MsgDnsTypeUnsupported, typ))
	}
	resolver := net.DefaultResolver
	if server, _ := c.Args.Str("Server"); server != "" {
		addr := server
		if _, _, err := net.SplitHostPort(addr); err != nil {
			addr = net.JoinHostPort(addr, "53")
		}
		resolver = &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, _ string) (net.Conn, error) {
				d := net.Dialer{}
				return d.DialContext(ctx, network, addr)
			},
		}
	}
	// 名称来源：管道在前，命名在后；空名不参与解析。
	var names []string
	for _, o := range c.Input {
		if s := o.String(); s != "" {
			names = append(names, s)
		}
	}
	if v := c.Args.Get("Name"); v != nil {
		for _, it := range v.ArrayItems() {
			if s := it.String(); s != "" {
				names = append(names, s)
			}
		}
	}
	if len(names) == 0 {
		return nil, nil
	}
	ctx := context.Background()
	var out []*object.PSObject
	for _, name := range names {
		recs, err := lookupDnsRecord(resolver, ctx, name, typ)
		if err != nil {
			if _, terr := errf(c, "%s", lang.T(lang.MsgDnsResolveFail, name)); terr != nil {
				return nil, terr
			}
			continue
		}
		out = append(out, recs...)
	}
	if len(out) == 0 {
		return nil, nil
	}
	return out, nil
}

// lookupDnsRecord 按记录类型查询一个名称（空类型查询 A 与 AAAA；无答案返回空，不报错）。
func lookupDnsRecord(r *net.Resolver, ctx context.Context, name, typ string) ([]*object.PSObject, error) {
	trimDot := func(h string) string { return strings.TrimSuffix(h, ".") }
	addrs := func(want string) ([]*object.PSObject, error) {
		ips, err := r.LookupIP(ctx, "ip", name)
		if err != nil {
			return nil, err
		}
		var out []*object.PSObject
		for _, ip := range ips {
			t := "AAAA"
			if ip.To4() != nil {
				t = "A"
			}
			if want == "" || want == t {
				out = append(out, dnsRecord(name, t, ip.String()))
			}
		}
		return out, nil
	}
	switch typ {
	case "", "A", "AAAA":
		return addrs(typ)
	case "CNAME":
		host, err := r.LookupCNAME(ctx, name)
		if err != nil {
			return nil, err
		}
		if trimDot(host) == trimDot(name) {
			return nil, nil
		}
		return []*object.PSObject{dnsRecord(name, typ, trimDot(host))}, nil
	case "MX":
		mxs, err := r.LookupMX(ctx, name)
		if err != nil {
			return nil, err
		}
		var out []*object.PSObject
		for _, mx := range mxs {
			out = append(out, dnsRecord(name, typ, strconv.FormatInt(int64(mx.Pref), 10)+" "+trimDot(mx.Host)))
		}
		return out, nil
	case "TXT":
		txts, err := r.LookupTXT(ctx, name)
		if err != nil {
			return nil, err
		}
		var out []*object.PSObject
		for _, txt := range txts {
			out = append(out, dnsRecord(name, typ, txt))
		}
		return out, nil
	case "NS":
		nss, err := r.LookupNS(ctx, name)
		if err != nil {
			return nil, err
		}
		var out []*object.PSObject
		for _, ns := range nss {
			out = append(out, dnsRecord(name, typ, trimDot(ns.Host)))
		}
		return out, nil
	case "PTR":
		hosts, err := r.LookupAddr(ctx, name)
		if err != nil {
			return nil, err
		}
		var out []*object.PSObject
		for _, h := range hosts {
			out = append(out, dnsRecord(name, typ, trimDot(h)))
		}
		return out, nil
	}
	return nil, nil
}

// ---- 注册 ----

func init() {
	Register("Test-Connection", []ParamSpec{
		{Name: "TargetName", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Count", Type: "int"},
	}, cmdTestConnection)
	Register("Start-Sleep", []ParamSpec{
		{Name: "Seconds", Position: 0, PositionSet: true, Type: "int"},
		{Name: "Milliseconds", Type: "int"},
	}, cmdStartSleep)
	Register("Resolve-DnsName", []ParamSpec{
		{Name: "Name", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Type", Position: 1, PositionSet: true, Type: "string"},
		{Name: "Server", Type: "string"},
	}, cmdResolveDnsName)
}

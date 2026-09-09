package builtin

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"powershell/internal/lang"
	"powershell/internal/object"
	"powershell/internal/shell"
)

// host.go 实现主机与信息输出类 cmdlet。

func cmdGetHost(c *Context) ([]*object.PSObject, error) {
	return []*object.PSObject{c.Shell.HostObject()}, nil
}

func cmdWriteVerbose(c *Context) ([]*object.PSObject, error) {
	msg := strings.Join(namedOrPosArgs(c, "Message"), " ")
	fmt.Fprintf(c.Stderr, "%s %s\n", lang.T(lang.MsgWriteVerbosePrefix), msg)
	return nil, nil
}

func cmdWriteWarning(c *Context) ([]*object.PSObject, error) {
	msg := strings.Join(namedOrPosArgs(c, "Message"), " ")
	fmt.Fprintf(c.Stderr, "%s %s\n", lang.T(lang.MsgWriteWarningPrefix), msg)
	return nil, nil
}

func cmdWriteInformation(c *Context) ([]*object.PSObject, error) {
	msg := strings.Join(namedOrPosArgs(c, "MessageData"), " ")
	fmt.Fprintln(c.console(), msg)
	return nil, nil
}

func cmdWriteDebug(c *Context) ([]*object.PSObject, error) {
	msg := strings.Join(namedOrPosArgs(c, "Message"), " ")
	fmt.Fprintf(c.Stderr, "%s %s\n", lang.T(lang.MsgWriteDebugPrefix), msg)
	return nil, nil
}

func cmdOutHost(c *Context) ([]*object.PSObject, error) {
	// 输出到主机（默认行为）：直接渲染，不进入管道
	objs := inputItems(c)
	_ = object.FormatOutput(c.console(), objs)
	return nil, nil
}

// cmdOutDefault 默认格式化出口直通：输入按默认格式直接渲染，不进入管道（-Transcript 接受忽略，无会话记录设施）。
func cmdOutDefault(c *Context) ([]*object.PSObject, error) {
	// 超量位置实参无槽位可接（InputObject 只接受命名），报错而非静默忽略。
	if len(c.Args.Positional) > 0 {
		return errf(c, "%s", lang.T(lang.MsgPositionalParamNotFound, c.Args.Positional[0].String()))
	}
	// -InputObject 整体作一个输入，不展开数组；与管道输入并存时报错。
	input := c.Input
	if arg := c.Args.Get("InputObject"); arg != nil {
		if len(c.Input) > 0 {
			return errf(c, "%s", lang.T(lang.MsgInputObjectWithPipeline))
		}
		input = []*object.PSObject{arg}
	}
	if len(input) == 0 {
		return nil, nil
	}
	_ = object.FormatOutput(c.console(), input)
	return nil, nil
}

func cmdClearHost(c *Context) ([]*object.PSObject, error) {
	fmt.Fprint(c.console(), "\x1b[2J\x1b[H")
	return nil, nil
}

func cmdReadHost(c *Context) ([]*object.PSObject, error) {
	prompt, _ := c.Args.Str("Prompt")
	line, ok, err := promptLine(c, prompt, bufio.NewReader(c.Stdin))
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	return []*object.PSObject{object.Str(line)}, nil
}

// promptLine 按 Read-Host 链路读一行：非交互运行报错；输出提示后读行；流结束时 ok 为 false。
// 同一次输入源的多次读取共用一个读取器，逐次新建时前一个读取器的预读会占用后续行。
func promptLine(c *Context, prompt string, reader *bufio.Reader) (string, bool, error) {
	// 非交互运行不能读输入，直接报错。
	if c.Shell.NonInteractive {
		return "", false, fmt.Errorf("%s", lang.T(lang.MsgReadHostNonInteractive))
	}
	if prompt != "" {
		fmt.Fprint(c.console(), prompt)
	}
	line, err := reader.ReadString('\n')
	if err != nil && line == "" {
		return "", false, nil
	}
	return strings.TrimRight(line, "\r\n"), true, nil
}

// cmdGetCredential 终端提示输入用户名和密码（复用 Read-Host 链路；密码输入回显，输出表格只显示用户名）。
func cmdGetCredential(c *Context) ([]*object.PSObject, error) {
	// 超量位置实参无槽位可接（UserName 只占位置 0），报错而非静默忽略。
	if len(c.Args.Positional) > 0 {
		return errf(c, "%s", lang.T(lang.MsgPositionalParamNotFound, c.Args.Positional[0].String()))
	}
	// 非交互运行无法提示，直接返回空（Read-Host 此时报错，原版如此）。
	if c.Shell.NonInteractive {
		return nil, nil
	}
	credential := c.Args.Get("Credential")
	user, _ := c.Args.Str("UserName")
	// UserName 位置实参是凭据对象时按 -Credential 处理（绑定器不按类型分参数集）；命名写成凭据对象报转换错误。
	if userVal := c.Args.Get("UserName"); userVal != nil && userVal.TypeName == "PSCredential" {
		if c.Args.PosMapped["UserName"] {
			credential, user = userVal, ""
		} else {
			return errf(c, "%s", lang.T(lang.MsgBindConvertFail, userVal.TypeName, "UserName", "string"))
		}
	}
	message, _ := c.Args.Str("Message")
	title, _ := c.Args.Str("Title")
	// -Credential 单独给出时直接返回，与用户名、消息、标题同时给出时报错。
	if credential != nil {
		if user != "" || message != "" || title != "" {
			return errf(c, "%s", lang.T(lang.MsgParamSetUnresolvable))
		}
		return []*object.PSObject{credential}, nil
	}
	// 标题或消息给出时显示它们，否则显示默认两行。
	if title != "" || message != "" {
		if title != "" {
			fmt.Fprintln(c.console(), title)
		}
		if message != "" {
			fmt.Fprintln(c.console(), message)
		}
	} else {
		fmt.Fprintln(c.console(), lang.T(lang.MsgCredentialHeader))
		fmt.Fprintln(c.console(), lang.T(lang.MsgCredentialPrompt))
	}
	reader := bufio.NewReader(c.Stdin)
	if user == "" {
		var ok bool
		var err error
		user, ok, err = promptLine(c, "User: ", reader)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, nil
		}
	}
	password, ok, err := promptLine(c, lang.T(lang.MsgCredentialPassword, user), reader)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	o := object.Object("PSCredential", nil)
	o.AddProp("UserName", user)
	o.AddProp("Password", password)
	o.Table = []object.Column{
		{Label: "UserName", Align: "left"},
	}
	return []*object.PSObject{o}, nil
}

func cmdInvokeExpression(c *Context) ([]*object.PSObject, error) {
	var src string
	if v := c.Args.Get("Command"); v != nil {
		src = v.String()
	} else if len(c.Input) > 0 {
		// 命名/位置都没给时才用管道输入（PowerShell 绑定顺序）
		var sb strings.Builder
		for _, o := range c.Input {
			sb.WriteString(o.String())
			sb.WriteByte('\n')
		}
		src = sb.String()
	}
	if src == "" {
		return nil, nil
	}
	return c.Engine.RunSource(src)
}

// transcriptStars 是记录起止块的分隔行。
const transcriptStars = "**********************"

// writeTranscriptHeader 向记录文件写起始块（只写能取到的主机信息）。
func writeTranscriptHeader(f *os.File, c *Context) {
	startTime := time.Now().Format("20060102150405")
	userName := ""
	if currentUser, err := user.Current(); err == nil {
		userName = currentUser.Username
	}
	machine := ""
	if hostName, err := os.Hostname(); err == nil {
		machine = hostName
	}
	edition := "Core"
	if c.Shell.Style == shell.StyleDesktop {
		edition = "Desktop"
	}
	fmt.Fprintln(f, transcriptStars)
	fmt.Fprintln(f, "PowerShell transcript start")
	fmt.Fprintln(f, "Start time: "+startTime)
	fmt.Fprintln(f, "Username: "+userName)
	fmt.Fprintln(f, "Machine: "+machine)
	fmt.Fprintln(f, "Host Application: "+transcriptAppName())
	fmt.Fprintln(f, "Process ID: "+strconv.Itoa(os.Getpid()))
	fmt.Fprintln(f, "PSEdition: "+edition)
	fmt.Fprintln(f, "OS: "+c.Shell.OSName())
	fmt.Fprintln(f, transcriptStars)
}

// transcriptAppName 取宿主程序的文件名。
func transcriptAppName() string {
	name := filepath.Base(os.Args[0])
	if rest := strings.Join(os.Args[1:], " "); rest != "" {
		return name + " " + rest
	}
	return name
}

// writeTranscriptEnd 向记录文件写结束块（Stop-Transcript 与重复启动共用）。
func writeTranscriptEnd(f *os.File) {
	fmt.Fprintln(f, transcriptStars)
	fmt.Fprintln(f, "PowerShell transcript end")
	fmt.Fprintln(f, "End time: "+time.Now().Format("20060102150405"))
	fmt.Fprintln(f, transcriptStars)
}

// defaultTranscriptPath 按原版规则拼默认记录路径（家目录下 主机名.随机.时间戳）。
func defaultTranscriptPath(c *Context) string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = c.Shell.Cwd
	}
	machine := "localhost"
	if hostName, err := os.Hostname(); err == nil && hostName != "" {
		machine = hostName
	}
	var randBytes [6]byte
	if _, err := rand.Read(randBytes[:]); err != nil {
		copy(randBytes[:], "abcdef")
	}
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var sb strings.Builder
	for _, byteVal := range randBytes {
		sb.WriteByte(alphabet[int(byteVal)%len(alphabet)])
	}
	name := "PowerShell_transcript." + machine + "." + sb.String() + "." + time.Now().Format("20060102150405") + ".txt"
	return filepath.Join(home, name)
}

// cmdStartTranscript 开始会话记录，把主机输出同时写入文件（Stop-Transcript 配对停止）。
func cmdStartTranscript(c *Context) ([]*object.PSObject, error) {
	// 超量位置实参无槽位可接（Path 只占位置 0），报错而非静默忽略。
	if len(c.Args.Positional) > 0 {
		return errf(c, "%s", lang.T(lang.MsgPositionalParamNotFound, c.Args.Positional[0].String()))
	}
	path := firstPathArg(c)
	if path == "" {
		path, _ = c.Args.Str("LiteralPath")
	}
	full := path
	if full == "" {
		full = defaultTranscriptPath(c)
	} else {
		var derr error
		full, derr = resolvePath(c, path)
		if derr != nil {
			return errf(c, "%v", derr)
		}
	}
	if _, serr := os.Stat(full); serr == nil && c.Args.Switch("NoClobber") {
		return errf(c, "%s", lang.T(lang.MsgFileExists, path))
	}
	var dryRun whatIfCollector
	dryRun.cmdlet = "Start-Transcript"
	dryRun.c = c
	var yesAll, noAll bool
	if dryRun.reportWhatIf(full) {
		out, _ := dryRun.result()
		return out, nil
	}
	if confirmSkip(c, "Start-Transcript", full, &yesAll, &noAll) {
		return nil, nil
	}
	// 已在记录时先结束旧的，再开新的（与 PowerShell 一致，不报错）。
	if c.Shell.TranscriptFile != nil {
		writeTranscriptEnd(c.Shell.TranscriptFile)
		_ = c.Shell.TranscriptFile.Close()
		c.Shell.TranscriptFile = nil
	}
	flags := os.O_WRONLY | os.O_CREATE
	if c.Args.Switch("Append") {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}
	f, err := os.OpenFile(full, flags, 0o644)
	if err != nil {
		return errf(c, "%s", lang.T(lang.MsgCannotWrite, path))
	}
	writeTranscriptHeader(f, c)
	c.Shell.TranscriptFile = f
	c.Shell.TranscriptPath = full
	return []*object.PSObject{object.Str(lang.T(lang.MsgTranscriptStarted, full))}, nil
}

// cmdStopTranscript 停止会话记录，关闭记录文件并返回状态行（未记录时报错）。
func cmdStopTranscript(c *Context) ([]*object.PSObject, error) {
	// 超量位置实参无槽位可接（Stop-Transcript 不收路径），报错而非静默忽略。
	if len(c.Args.Positional) > 0 {
		return errf(c, "%s", lang.T(lang.MsgPositionalParamNotFound, c.Args.Positional[0].String()))
	}
	var dryRun whatIfCollector
	dryRun.cmdlet = "Stop-Transcript"
	dryRun.c = c
	var yesAll, noAll bool
	// 预演先行：未记录时加 -WhatIf 只打印预演行，不报空闲错误（与 PowerShell 一致）。
	if dryRun.reportWhatIf("") {
		out, _ := dryRun.result()
		return out, nil
	}
	if c.Shell.TranscriptFile == nil {
		return errf(c, "%s", lang.T(lang.MsgStopTranscriptIdle))
	}
	path := c.Shell.TranscriptPath
	if confirmSkip(c, "Stop-Transcript", path, &yesAll, &noAll) {
		return nil, nil
	}
	writeTranscriptEnd(c.Shell.TranscriptFile)
	_ = c.Shell.TranscriptFile.Close()
	c.Shell.TranscriptFile = nil
	c.Shell.TranscriptPath = ""
	return []*object.PSObject{object.Str(lang.T(lang.MsgTranscriptStopped, path))}, nil
}

// ---- 注册 ----

func init() {
	Register("Get-Host", nil, cmdGetHost)
	Register("Write-Verbose", []ParamSpec{
		{Name: "Message", Position: 0, PositionSet: true, Type: "string"},
	}, cmdWriteVerbose)
	Register("Write-Warning", []ParamSpec{
		{Name: "Message", Position: 0, PositionSet: true, Type: "string"},
	}, cmdWriteWarning)
	Register("Write-Information", []ParamSpec{
		{Name: "MessageData", Position: 0, PositionSet: true, Type: "string"},
	}, cmdWriteInformation)
	Register("Write-Debug", []ParamSpec{
		{Name: "Message", Position: 0, PositionSet: true, Type: "string"},
	}, cmdWriteDebug)
	Register("Out-Host", []ParamSpec{
		{Name: "InputObject", Position: 0, PositionSet: true, Type: "object"},
	}, cmdOutHost)
	Register("Out-Default", []ParamSpec{
		{Name: "Transcript", Switch: true},
		{Name: "InputObject", Type: "object"},
	}, cmdOutDefault)
	Register("Clear-Host", nil, cmdClearHost)
	Register("Read-Host", []ParamSpec{
		{Name: "Prompt", Position: 0, PositionSet: true, Type: "string"},
	}, cmdReadHost)
	Register("Get-Credential", []ParamSpec{
		{Name: "UserName", Position: 0, PositionSet: true, Type: "string"},
		{Name: "Credential", Type: "object"},
		{Name: "Message", Type: "string"},
		{Name: "Title", Type: "string"},
	}, cmdGetCredential)
	Register("Start-Transcript", []ParamSpec{
		{Name: "Path", Position: 0, PositionSet: true, Type: "path"},
		{Name: "LiteralPath", Type: "path"},
		{Name: "Append", Switch: true},
		{Name: "Force", Switch: true},
		{Name: "NoClobber", Switch: true},
		{Name: "IncludeInvocationHeader", Switch: true},
		{Name: "UseMinimalHeader", Switch: true},
	}, cmdStartTranscript)
	Register("Stop-Transcript", nil, cmdStopTranscript)
	Register("Invoke-Expression", []ParamSpec{
		{Name: "Command", Position: 0, PositionSet: true, Type: "string"},
	}, cmdInvokeExpression)
}

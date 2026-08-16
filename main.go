// powershell —— 在 Linux 上运行的 PowerShell 风格解释器。
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"powershell/internal/eval"
	"powershell/internal/object"
	"powershell/internal/parser"
	"powershell/internal/repl"
	"powershell/internal/shell"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	fs := flag.NewFlagSet("powershell", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // 错误提示由我们自行处理

	version := fs.String("Version", "7", "PowerShell 风格版本：5.1 或 7")
	command := fs.String("Command", "", "执行命令后退出（- 表示从标准输入读取）")
	file := fs.String("File", "", "执行 .ps1 脚本后退出")
	noLogo := fs.Bool("NoLogo", false, "不显示启动横幅")
	noProfile := fs.Bool("NoProfile", false, "不加载启动脚本")
	_ = noProfile // 目前无启动脚本，参数占位接受
	help := fs.Bool("?", false, "显示帮助")
	fs.BoolVar(help, "Help", false, "显示帮助")

	if err := fs.Parse(args); err != nil {
		fmt.Fprintln(os.Stderr, "参数解析失败。")
		return 2
	}

	style := shell.StyleCore
	switch {
	case strings.HasPrefix(*version, "5"):
		style = shell.StyleDesktop
	case strings.HasPrefix(*version, "7"):
		style = shell.StyleCore
	}

	sess := shell.New(style, os.Stdout, os.Stderr, os.Stdin)
	ev := eval.New(sess, os.Stdin, os.Stdout, os.Stderr)

	// 帮助
	if *help || fs.NArg() > 0 && (fs.Arg(0) == "-?" || fs.Arg(0) == "-Help" || fs.Arg(0) == "-h") {
		fmt.Fprint(os.Stdout, sess.Usage())
		return 0
	}

	// -File 脚本（-File 后的剩余位置参数作为脚本实参，供 param() 与 $args 使用）
	if *file != "" {
		var scriptArgs []*object.PSObject
		for _, a := range fs.Args() {
			scriptArgs = append(scriptArgs, object.Str(a))
		}
		func() {
			// 脚本顶层未捕获的终止错误（throw）：打印后以失败退出码结束
			defer func() {
				if r := recover(); r != nil {
					if err := eval.RecoverError(r); err != nil {
						fmt.Fprintf(os.Stderr, "%s : %v\n", sess.StyleName(), err)
						os.Exit(1)
					}
					panic(r)
				}
			}()
			ev.RunScriptFileStreaming(*file, scriptArgs, func(objs []*object.PSObject) {
				_ = object.FormatOutput(os.Stdout, objs)
			})
		}()
		return exitCode(ev, sess)
	}

	// -Command
	if *command != "" {
		var src string
		if *command == "-" {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				src = ""
			} else {
				src = string(data)
			}
		} else {
			src = *command
		}
		return executeOnce(sess, ev, src)
	}

	// 交互式 REPL：Linux 上从根目录进入（提示符 PS C:\>）
	if runtime.GOOS == "linux" {
		_ = os.Chdir("/")
	}
	sess.Cwd, _ = os.Getwd()
	repl.Run(sess, ev, !*noLogo, os.Stdin, os.Stdout, os.Stderr)
	return 0
}

// executeOnce 执行一段命令文本并输出结果，返回退出码。
func executeOnce(sess *shell.Session, ev *eval.Evaluator, src string) int {
	res := parser.Parse(src)
	if res.Error != nil {
		fmt.Fprintf(os.Stderr, "%s : %v\n", sess.StyleName(), res.Error)
		return 1
	}
	// 逐语句执行并格式化，保证与 Write-Host/Format-* 等直写命令的顺序一致
	for _, st := range res.List.Statements {
		objs := ev.EvalStatement(st)
		_ = object.FormatOutput(os.Stdout, objs)
		if ev.ExitRequested {
			return exitCode(ev, sess)
		}
	}
	return exitCode(ev, sess)
}

func exitCode(ev *eval.Evaluator, sess *shell.Session) int {
	if ev.ExitRequested {
		return ev.ExitCode
	}
	if sess.LastExit != 0 {
		return sess.LastExit
	}
	return 0
}

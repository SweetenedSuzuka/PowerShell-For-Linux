// powershell —— 在 Linux 上运行的 PowerShell 风格解释器。
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"powershell/internal/eval"
	"powershell/internal/lang"
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
	noExit := fs.Bool("NoExit", false, "执行后进入交互，不退出")
	nonInteractive := fs.Bool("NonInteractive", false, "非交互运行：确认提示直接拒绝，读取输入报错")
	executionPolicy := fs.String("ExecutionPolicy", "", "执行策略（只校验取值，不限制执行）")
	workingDirectory := fs.String("WorkingDirectory", "", "启动目录")
	help := fs.Bool("?", false, "显示帮助")
	fs.BoolVar(help, "Help", false, "显示帮助")

	if err := fs.Parse(args); err != nil {
		fmt.Fprintln(os.Stderr, lang.T(lang.MsgFlagParseFail))
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
	sess.NonInteractive = *nonInteractive

	// 帮助
	if *help || fs.NArg() > 0 && (fs.Arg(0) == "-?" || fs.Arg(0) == "-Help" || fs.Arg(0) == "-h") {
		fmt.Fprint(os.Stdout, sess.Usage())
		return 0
	}

	// 执行策略：只校验取值（Linux 上 PowerShell 同样不限制执行，未知取值报错）
	if *executionPolicy != "" {
		switch strings.ToLower(*executionPolicy) {
		case "allsigned", "bypass", "default", "remotesigned", "restricted", "unrestricted", "undefined":
		default:
			fmt.Fprintf(os.Stderr, "%s\n", lang.T(lang.MsgExecutionPolicyInvalid, *executionPolicy))
			return 2
		}
	}

	// 启动目录：不存在报错后继续（对齐 PowerShell）
	if *workingDirectory != "" {
		if err := os.Chdir(*workingDirectory); err != nil {
			fmt.Fprintf(os.Stderr, "%s : %s\n", sess.StyleName(), lang.T(lang.MsgPathNotFoundFmt, *workingDirectory))
		} else {
			sess.Cwd, _ = os.Getwd()
		}
	}

	// 启动脚本：默认加载 $HOME/.config/powershell/profile.ps1（-NoProfile 跳过）
	if !*noProfile {
		loadProfile(ev, sess, os.Stdout)
	}

	// -File 脚本（-File 后的剩余位置参数作为脚本实参，供 param() 与 $args 使用）
	if *file != "" {
		var scriptArgs []*object.PSObject
		for _, a := range fs.Args() {
			scriptArgs = append(scriptArgs, object.Str(a))
		}
		// 脚本失败标记：defer 内只标记不退出，展开走完再返回退出码。
		failed := false
		func() {
			// 脚本顶层未捕获的终止错误（throw）：打印后标记失败
			defer func() {
				if r := recover(); r != nil {
					if err := eval.RecoverError(r); err != nil {
						fmt.Fprintf(os.Stderr, "%s : %v\n", sess.StyleName(), err)
						failed = true
						return
					}
					panic(r)
				}
			}()
			ev.RunScriptFileStreaming(*file, scriptArgs, func(objs []*object.PSObject) {
				_ = object.FormatOutput(os.Stdout, objs)
			})
		}()
		code := exitCode(ev, sess)
		if failed {
			code = 1
		}
		if !*noExit {
			return code
		}
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
		// -NoExit：执行后进入交互，不直接返回退出码。
		if !*noExit {
			return executeOnce(sess, ev, src)
		}
		executeOnce(sess, ev, src)
	}

	// 交互式 REPL：Linux 上从根目录进入（提示符 PS C:\>，指定启动目录除外）
	if runtime.GOOS == "linux" && *workingDirectory == "" {
		_ = os.Chdir("/")
	}
	sess.Cwd, _ = os.Getwd()
	repl.Run(sess, ev, !*noLogo, os.Stdin, os.Stdout, os.Stderr)
	return 0
}

// loadProfile 加载用户启动脚本 $HOME/.config/powershell/profile.ps1（PowerShell 7 Linux 的 $PROFILE 路径）。
// 文件不存在跳过；脚本出错打印提示后继续启动（对齐 PowerShell 的宽容行为）。
func loadProfile(ev *eval.Evaluator, sess *shell.Session, out io.Writer) {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	path := filepath.Join(home, ".config", "powershell", "profile.ps1")
	if _, err := os.Stat(path); err != nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			if err := eval.RecoverError(r); err != nil {
				fmt.Fprintf(out, "%s : %v\n", sess.StyleName(), err)
				return
			}
			panic(r)
		}
	}()
	ev.RunScriptFileStreaming(path, nil, func(objs []*object.PSObject) {
		_ = object.FormatOutput(out, objs)
	})
}

// executeOnce 执行一段命令文本并输出结果，返回退出码。
func executeOnce(sess *shell.Session, ev *eval.Evaluator, src string) (code int) {
	// 单次执行顶层回收：普通 panic 转为报错并返回失败码。
	defer func() {
		if r := recover(); r != nil {
			if ev.ReportPanic(r) {
				code = 1
				return
			}
			panic(r)
		}
	}()
	res := parser.Parse(src)
	if res.Error != nil {
		fmt.Fprintf(os.Stderr, "%s : %v\n", sess.StyleName(), res.Error)
		return 1
	}
	// 非交互单次执行拒绝不完整的输入（交互续行允许）。
	if res.Incomplete {
		fmt.Fprintf(os.Stderr, "%s : %s\n", sess.StyleName(), lang.T(lang.MsgIncompleteInput))
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

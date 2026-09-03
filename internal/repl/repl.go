// Package repl 实现交互式 REPL：提示符、行编辑、历史、Tab 补全、多行续行。
package repl

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"powershell/internal/builtin"
	"powershell/internal/eval"
	"powershell/internal/object"
	"powershell/internal/parser"
	"powershell/internal/shell"
)

// lineReader 是行读取接口（Unix 为 raw 模式编辑器，其它平台回退为简单读取）。
type lineReader interface {
	ReadLine(prompt string) (string, error)
}

// REPL 是一次交互式会话。
type REPL struct {
	Session *shell.Session
	Eval    *eval.Evaluator
	in      *os.File
	out     io.Writer
	reader  lineReader
	pending string // 多行续行累积
}

// Run 启动 REPL 主循环。
func Run(sess *shell.Session, ev *eval.Evaluator, showBanner bool, in *os.File, out, errw io.Writer) {
	// 交互顶层回收：普通 panic 转为报错后返回，历史保存不受影响。
	defer func() {
		if rec := recover(); rec != nil {
			if ev.ReportPanic(rec) {
				return
			}
			panic(rec)
		}
	}()
	if showBanner {
		fmt.Fprintln(out, sess.Banner())
		fmt.Fprintln(out)
	}
	r := &REPL{
		Session: sess,
		Eval:    ev,
		in:      in,
		out:     out,
	}
	hist := loadHistory(historyPath())
	sess.History = hist
	defer func() {
		if err := saveHistory(historyPath(), sess.History); err != nil {
			// 历史保存失败静默忽略
		}
	}()
	r.reader = newLineReader(in, out, &sess.History, r.complete)
	r.loop()
}

func (r *REPL) loop() {
	for {
		prompt := r.promptText() + " "
		if r.pending != "" {
			prompt = r.Session.ContinuationPrompt() + " "
		}
		line, err := r.reader.ReadLine(prompt)
		if err != nil {
			if err == io.EOF {
				fmt.Fprintln(r.out)
			}
			return
		}
		r.pending += line + "\n"
		res := parser.Parse(r.pending)
		if res.Error != nil {
			fmt.Fprintf(r.out, "%s : %v\n", r.Session.StyleName(), res.Error)
			r.pending = ""
			continue
		}
		if res.Incomplete {
			continue
		}
		// 记录历史（整条合并）
		joined := strings.Join(strings.Split(strings.TrimRight(r.pending, "\n"), "\n"), "; ")
		if len(r.Session.History) == 0 || r.Session.History[len(r.Session.History)-1] != joined {
			r.Session.History = append(r.Session.History, joined)
		}
		// 逐语句执行并格式化，保证与直写命令顺序一致
		exited := false
		recovered := false
		func() {
			// 单轮回收：普通 panic 转为报错后继续下一轮。
			defer func() {
				if rec := recover(); rec != nil {
					if r.Eval.ReportPanic(rec) {
						recovered = true
						return
					}
					panic(rec)
				}
			}()
			for _, st := range res.List.Statements {
				objs := r.Eval.EvalStatement(st)
				_ = object.FormatOutput(r.out, objs)
				if r.Eval.ExitRequested {
					exited = true
					return
				}
			}
		}()
		if exited {
			r.pending = ""
			return
		}
		if recovered {
			r.pending = ""
			continue
		}
		r.pending = ""
	}
}

// promptText 返回主提示符：定义了 prompt 函数时调函数取值，失败回默认。
func (r *REPL) promptText() (text string) {
	text = r.Session.Prompt()
	if !hasPromptFunc(r.Session) {
		return text
	}
	res := parser.Parse("prompt\n")
	if res.Error != nil || res.Incomplete || len(res.List.Statements) == 0 {
		return text
	}
	var b strings.Builder
	func() {
		defer func() {
			if recover() == nil {
				return
			}
			b.Reset()
		}()
		for _, st := range res.List.Statements {
			for _, o := range r.Eval.EvalStatement(st) {
				b.WriteString(o.String())
			}
		}
	}()
	if b.Len() == 0 {
		return text
	}
	return b.String()
}

// hasPromptFunc 报告会话是否定义了 prompt 函数（名大小写不限）。
func hasPromptFunc(sess *shell.Session) bool {
	if _, ok := sess.Functions["prompt"]; ok {
		return true
	}
	for n := range sess.Functions {
		if strings.EqualFold(n, "prompt") {
			return true
		}
	}
	return false
}

// ---- Tab 补全 ----

func (r *REPL) complete(buf string) []string {
	if buf == "" {
		return nil
	}
	// 变量补全
	if strings.HasPrefix(buf, "$") {
		name := strings.TrimPrefix(buf, "$")
		var out []string
		for _, n := range r.Session.AllVarNames() {
			if strings.HasPrefix(strings.ToLower(n), strings.ToLower(name)) {
				out = append(out, "$"+n)
			}
		}
		return out
	}
	// 命令补全（第一个词）
	fields := strings.Fields(buf)
	if len(fields) == 0 {
		return nil // 纯空白行：无词可补
	}
	lastTok := fields[len(fields)-1]
	// 参数名补全（末词以 - 开头，按首词命令的规格参数加通用参数枚举）
	if strings.HasPrefix(lastTok, "-") {
		return r.completeParam(fields[0], lastTok)
	}
	if len(fields) == 1 {
		var out []string
		for _, n := range builtin.AllCmdletNames() {
			if strings.HasPrefix(strings.ToLower(n), strings.ToLower(lastTok)) {
				out = append(out, n+" ")
			}
		}
		for _, n := range r.Session.AllAliasNames() {
			if strings.HasPrefix(strings.ToLower(n), strings.ToLower(lastTok)) {
				out = append(out, n+" ")
			}
		}
		for n := range r.Session.Functions {
			if strings.HasPrefix(strings.ToLower(n), strings.ToLower(lastTok)) {
				out = append(out, n+" ")
			}
		}
		return out
	}
	// 文件路径补全
	return completePath(r.Session, lastTok)
}

// completeParam 补全指定命令的参数名（规格参数在前，通用参数在后，去重）。
func (r *REPL) completeParam(cmd, prefix string) []string {
	name := cmd
	if resolved, ok := r.Session.ResolveAlias(cmd); ok {
		name = resolved
	}
	if _, ok := builtin.Lookup(name); !ok {
		return nil
	}
	want := strings.ToLower(strings.TrimPrefix(prefix, "-"))
	seen := map[string]bool{}
	var out []string
	add := func(n string) {
		key := strings.ToLower(n)
		if strings.HasPrefix(key, want) && !seen[key] {
			seen[key] = true
			out = append(out, "-"+n+" ")
		}
	}
	for _, sp := range builtin.Spec(name) {
		add(sp.Name)
	}
	for _, n := range builtin.CommonParamNames() {
		add(n)
	}
	return out
}

func completePath(sess *shell.Session, tok string) []string {
	dir := filepath.Dir(tok)
	if dir == "." {
		dir = "."
	} else if !filepath.IsAbs(dir) {
		dir = filepath.Join(sess.Cwd, dir)
	}
	base := filepath.Base(tok)
	if base == tok && !strings.Contains(tok, string(filepath.Separator)) {
		base = tok
		dir = sess.Cwd
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var out []string
	for _, en := range entries {
		name := en.Name()
		if !strings.HasPrefix(strings.ToLower(name), strings.ToLower(base)) {
			continue
		}
		prefix := tok[:len(tok)-len(base)]
		if en.IsDir() {
			out = append(out, prefix+name+string(filepath.Separator))
		} else {
			out = append(out, prefix+name+" ")
		}
	}
	return out
}

// ---- 历史持久化 ----

func historyPath() string {
	if p := os.Getenv("POWERSHELL_HISTORY_FILE"); p != "" {
		return p
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ".powershell_history"
	}
	return filepath.Join(home, ".powershell_history")
}

func loadHistory(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var lines []string
	for _, ln := range strings.Split(string(data), "\n") {
		ln = strings.TrimRight(ln, "\r")
		if ln != "" {
			lines = append(lines, ln)
		}
	}
	return lines
}

func saveHistory(path string, hist []string) error {
	if len(hist) == 0 {
		return nil
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, h := range hist {
		fmt.Fprintln(f, h)
	}
	return nil
}

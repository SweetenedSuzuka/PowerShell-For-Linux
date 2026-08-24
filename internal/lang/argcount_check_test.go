package lang

// 静态核对：全仓库所有 lang.T 调用的实参个数必须与该 Msg 文本的占位符个数一致。
// 通过 go/ast 解析源码完成，避免 grep 误判。

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"

	"testing"
)

// msgIndexOf 把 Msg 常量名映射到编号值（解析本包常量声明自动生成）。
var msgIndexOf = buildMsgIndex()

func buildMsgIndex() map[string]int {
	out := map[string]int{}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "messages.go", nil, 0)
	if err != nil {
		return out
	}
	cur := -1
	for _, d := range f.Decls {
		gd, ok := d.(*ast.GenDecl)
		if !ok || gd.Tok != token.CONST {
			continue
		}
		for _, spec := range gd.Specs {
			vs := spec.(*ast.ValueSpec)
			if len(vs.Values) > 0 {
				// 唯一的显式值是块首的 MsgBannerDesktop = iota：从这里重新计数
				if id, ok := vs.Values[0].(*ast.Ident); ok && id.Name == "iota" {
					cur = 0
				}
			}
			if cur < 0 {
				continue
			}
			for _, name := range vs.Names {
				out[name.Name] = cur
				cur++
			}
		}
	}
	return out
}

var verbRe = regexp.MustCompile(`%[#+\- 0]*[0-9]*(?:\.[0-9]+)?[vsdqxX]`)

func placeholderCount(s string) int { return len(verbRe.FindAllString(s, -1)) }

func TestCallSitesArgCount(t *testing.T) {
	root := "../.."
	fset := token.NewFileSet()
	// 仓库根的 main 包与 internal 各子目录逐个解析
	var dirs []string
	dirs = append(dirs, root)
	internalEntries, err := os.ReadDir(filepath.Join(root, "internal"))
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range internalEntries {
		if e.IsDir() {
			dirs = append(dirs, filepath.Join(root, "internal", e.Name()))
		}
	}
	sort.Strings(dirs)
	for _, path := range dirs {
		pkg, err := parser.ParseDir(fset, path, func(fi os.FileInfo) bool {
			return !fi.IsDir() && filepath.Ext(fi.Name()) == ".go"
		}, 0)
		if err != nil {
			t.Fatalf("%s: %v", path, err)
		}
		for _, p := range pkg {
			for _, f := range p.Files {
				ast.Inspect(f, func(n ast.Node) bool {
					call, ok := n.(*ast.CallExpr)
					if !ok {
						return true
					}
					sel, ok := call.Fun.(*ast.SelectorExpr)
					if !ok || sel.Sel.Name != "T" {
						return true
					}
					ident, ok := sel.X.(*ast.Ident)
					if !ok || ident.Name != "lang" {
						return true
					}
					if len(call.Args) == 0 {
						return true
					}
					msgName, ok := call.Args[0].(*ast.SelectorExpr)
					if !ok {
						return true
					}
					args := len(call.Args) - 1
					want := placeholderCount(zh[Msg(msgIndexOf[msgName.Sel.Name])])
					if args != want {
						t.Errorf("%s: %s 期望 %d 个实参，实际 %d 个",
							fset.Position(call.Pos()), msgName.Sel.Name, want, args)
					}
					return true
				})
			}
		}
	}
}

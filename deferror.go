package deferror

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

const doc = "deferror is a Go linter that suggests a customly made defer function call"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name:             "deferror",
	Doc:              doc,
	Run:              run,
	RunDespiteErrors: false,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	configInit(pass)
	nodeFilter := []ast.Node{&ast.FuncDecl{}}
	a.inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			if ok, name := a.namedReturnErr(n); ok && !a.deferStart(n) {
				a.report(n, name)
			}
		}
	})

	return nil, nil
}

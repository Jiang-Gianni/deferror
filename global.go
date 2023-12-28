package deferror

import (
	"text/template"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ast/inspector"
)

var a = &A{}

type A struct {
	pass    *analysis.Pass
	ssa     *buildssa.SSA
	inspect *inspector.Inspector
	tmpl    *template.Template
}

type F struct {
	PkgName     string
	PkgPath     string
	FnName      string
	RecvName    string
	RecvType    string
	RecvPointer bool
	Params      []P
	ErrName     string
}

type P struct {
	Name    string
	Type    string
	Pointer bool
}

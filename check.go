package deferror

import (
	"go/ast"
)

func (a *A) namedReturnErr(fd *ast.FuncDecl) bool {
	if fd.Type == nil || fd.Type.Results == nil || fd.Type.Results.List == nil {
		return false
	}
	for _, field := range fd.Type.Results.List {
		ident, ok := field.Type.(*ast.Ident)
		if ok && ident.Name == "error" {
			for _, name := range field.Names {
				if name.Name == "err" {
					return true
				}
			}
		}
	}
	return false
}

func (a *A) deferStart(fd *ast.FuncDecl) bool {
	if len(fd.Body.List) == 0 {
		return false
	}
	_, ok := fd.Body.List[0].(*ast.DeferStmt)
	return ok
}

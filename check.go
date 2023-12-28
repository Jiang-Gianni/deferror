package deferror

import (
	"go/ast"
)

// Return true/false and the name of the error
func (a *A) namedReturnErr(fd *ast.FuncDecl) (bool, string) {
	if fd.Type == nil || fd.Type.Results == nil || fd.Type.Results.List == nil {
		return false, ""
	}
	for _, field := range fd.Type.Results.List {
		ident, ok := field.Type.(*ast.Ident)
		if ok && ident.Name == "error" && field.Names != nil && len(field.Names) > 0 && field.Names[0].Name != "" {
			return true, field.Names[0].Name
		}
	}
	return false, ""
}

// Function body start with a defer call
func (a *A) deferStart(fd *ast.FuncDecl) bool {
	if len(fd.Body.List) == 0 {
		return false
	}
	_, ok := fd.Body.List[0].(*ast.DeferStmt)
	return ok
}

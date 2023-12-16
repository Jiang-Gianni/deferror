package deferror

import (
	"bytes"
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func (a *A) report(fd *ast.FuncDecl) {

	f := &F{
		PkgName: a.pass.Pkg.Name(),
		PkgPath: a.pass.Pkg.Path(),
		FnName:  fd.Name.String(),
		Params:  params(fd),
	}
	f.RecvName, f.RecvType, f.RecvPointer = receiver(fd)
	// fmt.Printf("\n%+v\n", f)
	b, err := a.deferCall(f)
	if err != nil {
		panic(err)
	}
	a.pass.Report(analysis.Diagnostic{
		Pos:      fd.Body.Lbrace + 2,
		End:      fd.Body.End(),
		Category: "deferror",
		// Message:  "deferror",
		Message: fmt.Sprintf(
			"deferror suggests the following defer call for the function %s:\n %s",
			fd.Name.String(),
			b,
		),
		SuggestedFixes: []analysis.SuggestedFix{{TextEdits: []analysis.TextEdit{{
			Pos:     fd.Body.Lbrace + 2,
			End:     fd.Body.Lbrace + 2,
			NewText: b,
		}}}},
	})
}

func (a *A) deferCall(data any) ([]byte, error) {
	buffer := new(bytes.Buffer)
	if err := a.tmpl.ExecuteTemplate(buffer, "main", data); err != nil {
		return nil, fmt.Errorf("a.tmpl.ExecuteTemplate: %w", err)
	}
	return buffer.Bytes(), nil
}

func fieldNameType(field *ast.Field) (fieldName string, fieldType string, pointer bool) {
	// fmt.Printf("\n%T\n", field.Type)
	fieldName = field.Names[0].String()
	fieldType, pointer = exprType(field.Type)
	return
}

func exprType(expr ast.Expr) (string, bool) {
	var name, key, value string
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name, false
	case *ast.SelectorExpr:
		name, _ = exprType(t.X)
		return name + "." + t.Sel.Name, false
	case *ast.StarExpr:
		name, _ = exprType(t.X)
		return "*" + name, true
	case *ast.ArrayType:
		name, _ = exprType(t.Elt)
		return "[]" + name, false
	case *ast.MapType:
		key, _ = exprType(t.Key)
		value, _ = exprType(t.Value)
		return "map[" + key + "]" + value, false
	case *ast.ChanType:
		value, _ = exprType(t.Value)
		return "chan " + value, false
	case *ast.StructType:
		name = "struct{"
		for _, f := range t.Fields.List {
			fType, _ := exprType(f.Type)
			name += f.Names[0].Name + " " + fType + "\n"
		}
		return name + "}", false
	case *ast.FuncType:
		name = "func("
		for _, f := range t.Params.List {
			fType, _ := exprType(f.Type)
			name += fType + ","
		}
		name += ")"
		if t.Results != nil && len(t.Results.List) > 0 {
			name += "("
			for _, f := range t.Results.List {
				fType, _ := exprType(f.Type)
				name += fType + ","
			}
			name += ")"
		}
		return name, false
	}
	return "", false
}

func receiver(fd *ast.FuncDecl) (RecvName, RecvType string, RecvPointer bool) {
	if fd.Recv == nil || fd.Recv.List == nil || fd.Recv.List[0].Names == nil {
		return RecvName, RecvType, RecvPointer
	}
	return fieldNameType(fd.Recv.List[0])
}

func params(fd *ast.FuncDecl) []P {
	if len(fd.Type.Params.List) == 0 {
		return nil
	}
	pList := make([]P, len(fd.Type.Params.List))
	for i, p := range fd.Type.Params.List {
		pList[i].Name, pList[i].Type, pList[i].Pointer = fieldNameType(p)
	}
	return pList

}

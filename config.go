package deferror

import (
	"log"
	"os"
	"text/template"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var tmpFile = "./deferror.tmpl"

func configInit(pass *analysis.Pass) {

	a.pass = pass
	a.ssa = pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	a.inspect = pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	envDfrr := os.Getenv("DFRR_FILE")
	if envDfrr != "" {
		tmpFile = envDfrr
	}

	if _, err := os.Stat(tmpFile); err != nil {
		log.Fatalf("template file not provided or not readable %s: %s", tmpFile, err)
	}

	t, err := template.New("main").Funcs(funcMap).ParseFiles(tmpFile)
	if err != nil {
		log.Fatalf("template.ParseFiles: %s", err)
	}
	a.tmpl = t

}

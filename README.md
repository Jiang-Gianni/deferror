# deferror

Deferror is a Go linter that suggests a customly made `defer` function call when the function has a return named error and the function does not already start with a `defer` call.

## Install

```bash
go install github.com/Jiang-Gianni/deferror/cmd/deferror@latest
```


## How to use

A template file for the deferred function call is needed.

To generate a template file like [deferror.tmpl](./deferror.tmpl) with a custom name in the current directory:

```bash
deferror init -o "yourDeferror.tmpl"
```

If the `-o` flag is not specified then the file will be generated as `deferror.tmpl`.

To run the linter on a `go` file:

```bash
DFRR_FILE=yourDeferror.tmpl deferror yourGoFile.go
```

If the environment variable `DFRR_FILE` is not specified, the linter will look for `deferror.tmpl` as the template.

Add the `-fix` flag to make the linter write the suggested defer function call:

```bash
DFRR_FILE=yourDeferror.tmpl deferror -fix yourGoFile.go
```


## Template

The template must uses [text/template](https://pkg.go.dev/text/template) package. The functions in the [strings](https://pkg.go.dev/strings) package are made available (see [funcmap.go](./funcmap.go)).

The structure of the input data map is the following (type `F`):

```go
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
```


## Example

[readme.go](./example/readme.go)

```go
package example

import (
	"time"
)

type R struct{}

func (r *R) MyExample(now time.Duration, i *int) (a int, err error) {
	return 0, nil
}
```

The input data map for the template is:

```go
&{PkgName:example PkgPath:github.com/Jiang-Gianni/deferror/example FnName:MyExample RecvName:r RecvType:*R RecvPointer:true Params:[{Name:now Type:time.Duration Pointer:false} {Name:i Type:*int Pointer:true}] ErrName: err}
```

The following defer call is suggested with the default template:

```go
	defer func(now time.Duration, i *int) {
		if err != nil {
			err = fmt.Errorf("r.MyExample(%v, %v, ): %w", now, i, err)
		}
	}(now, i)
```


## Tips

### Wrap function [wrap.go](./dfrr/wrap.go)

With a `Wrap` function like [this](https://github.com/golang/pkgsite/blob/master/internal/derrors/derrors.go#L240):

```go
package dfrr

import "fmt"

func Wrap(errp *error, format string, args ...any) {
	if *errp != nil {
		*errp = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), *errp)
	}
}
```

the template can be simplified to:

```tmpl
{{define "main"}}   defer dfrr.Wrap(&err, "{{if .RecvName}}{{.RecvName}}.{{end}}{{.FnName}}({{template "percV" .}})", {{template "plist" .}})
{{end}}

{{define "plist"}}{{range $item := .Params}}{{$item.Name}}, {{end}}{{end}}

{{define "percV"}}{{range $item := .Params}}%v, {{end}}{{end}}
```

which suggests (`dfrr` module will need to be imported):

```go
defer dfrr.Wrap(&err, "r.MyExample(%v, %v, )", now, i)
```

You can clone this repository and try it yourself on the `example` folder.

## Subpackage templates

By using the [strings](https://pkg.go.dev/strings) package functions it is possible to determine if the analyzed function package name or path contains a substring which means it allows to define subpackage specific templates.

For example you may want to add the complete input list values to the wrapping error message only inside critical package.

```tmpl
{{if Contains .PkgPath "/api/"}}
{{template "apiDefer" .}}
{{end}}
```
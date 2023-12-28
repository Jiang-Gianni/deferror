/*
Copyright © 2024 Gianni Jiang

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate a deferror.tmpl file containing an example of the defer call.",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.WriteFile(viper.GetString("out"), tmplContents, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var tmplContents = []byte(`{{define "main"}}   defer func({{template "input" .}}) {
    if {{.ErrName}} != nil {
        {{.ErrName}} = fmt.Errorf("{{if .RecvName}}{{.RecvName}}.{{end}}{{.FnName}}({{template "percV" .}}): %w", {{template "plist" .}}{{if .Params}} ,{{end}}{{.ErrName}})
    }
}({{template "plist" .}})
{{end}}

{{define "plist"}}{{range $index, $item := .Params}}{{if $index}}, {{end}}{{$item.Name}}{{end}}{{end}}

{{define "percV"}}{{range $index, $item := .Params}}{{if $index}}, {{end}}%v{{end}}{{end}}

{{define "input"}}{{range $index, $item := .Params}}{{if $index}}, {{end}}{{$item.Name}} {{$item.Type}}{{end}}{{end}}`)

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.PersistentFlags().StringP("out", "o", "deferror.tmpl", "Output file for the template")
	viper.BindPFlag("out", rootCmd.PersistentFlags().Lookup("out"))
}

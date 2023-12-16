package deferror

import (
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"Clone":          strings.Clone,
	"Compare":        strings.Compare,
	"Contains":       strings.Contains,
	"ContainsAny":    strings.ContainsAny,
	"ContainsRune":   strings.ContainsRune,
	"Count":          strings.Count,
	"EqualFold":      strings.EqualFold,
	"Fields":         strings.Fields,
	"FieldsFunc":     strings.FieldsFunc,
	"HasPrefix":      strings.HasPrefix,
	"HasSuffix":      strings.HasSuffix,
	"Index":          strings.Index,
	"IndexAny":       strings.IndexAny,
	"IndexByte":      strings.IndexByte,
	"IndexFunc":      strings.IndexFunc,
	"IndexRune":      strings.IndexRune,
	"Join":           strings.Join,
	"LastIndex":      strings.LastIndex,
	"LastIndexAny":   strings.LastIndexAny,
	"LastIndexByte":  strings.LastIndexByte,
	"LastIndexFunc":  strings.LastIndexFunc,
	"Map":            strings.Map,
	"Repeat":         strings.Repeat,
	"Replace":        strings.Replace,
	"ReplaceAll":     strings.ReplaceAll,
	"Split":          strings.Split,
	"SplitAfter":     strings.SplitAfter,
	"SplitAfterN":    strings.SplitAfterN,
	"SplitN":         strings.SplitN,
	"ToLower":        strings.ToLower,
	"ToLowerSpecial": strings.ToLowerSpecial,
	"ToTitle":        strings.ToTitle,
	"ToTitleSpecial": strings.ToTitleSpecial,
	"ToUpper":        strings.ToUpper,
	"ToUpperSpecial": strings.ToUpperSpecial,
	"ToValidUTF8":    strings.ToValidUTF8,
	"Trim":           strings.Trim,
	"TrimFunc":       strings.TrimFunc,
	"TrimLeft":       strings.TrimLeft,
	"TrimLeftFunc":   strings.TrimLeftFunc,
	"TrimPrefix":     strings.TrimPrefix,
	"TrimRight":      strings.TrimRight,
	"TrimRightFunc":  strings.TrimRightFunc,
	"TrimSpace":      strings.TrimSpace,
	"TrimSuffix":     strings.TrimSuffix,
}
package plugin

import "text/template"

var headerTemplate = template.Must(template.New("header").Parse(`
// Code generated by protoc-gen-go-gorm. DO NOT EDIT.
// source: {{ .Proto.Name }}

package {{ .GoPackageName }}
`))

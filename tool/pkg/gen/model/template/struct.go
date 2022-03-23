package template

var Struct = `type {{.camelName}} struct {
{{range .fields}}	{{.}}
{{end}}
}`

package template

var Field = `{{.camelName}} {{.type}} {{.tag}} {{if .hasComment}}// {{.comment}}{{end}}`

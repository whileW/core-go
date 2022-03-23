package template

var Model = `package model

{{.imports}}

{{.struct}}

{{.selfMethods}}

{{.methods}}
`

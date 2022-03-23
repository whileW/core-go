package template

import "text/template"

var (
	ModelT,_ = template.New("model").Parse(Model)
	ImportT,_ = template.New("import").Parse(Import)
	FieldT,_ = template.New("field").Parse(Field)
	StructT,_ = template.New("struct").Parse(Struct)
)
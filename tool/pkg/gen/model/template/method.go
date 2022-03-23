package template

import "text/template"

var MethodT,_ = template.New("method").Parse(method)
var method = `func Search{{.camelName}}List(db *gorm.DB,opts ...orm.Option) ([]*{{.camelName}},error) {
	var data = []*{{.camelName}}{}
	opts = append(opts, orm.Option_TableName({{.camelName}}{}.TableName()))
	return data,orm.GetListRecord(db,&data,opts...)
}
func Search{{.camelName}}First(db *gorm.DB,opts ...orm.Option) (*{{.camelName}},error) {
	var data = &{{.camelName}}{}
	opts = append(opts, orm.Option_TableName({{.camelName}}{}.TableName()))
	return data,orm.GetFirstRecord(db,&data,opts...)
}`

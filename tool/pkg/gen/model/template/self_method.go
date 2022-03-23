package template

import "text/template"

var SelfMethodT,_ = template.New("selfMethod").Parse(selfMethod)
var selfMethod = `func ({{.camelName}})TableName() string {
	return "{{.name}}"
}
func (t *{{.camelName}})BeforeCreate(tx *gorm.DB) error {
	{{if .isNeedUUID}}if t.UUID == ""  {
		t.UUID = uuid.New().String()
	}{{end}}
	return nil
}`
package model

import "github.com/whileW/core-go/tool/pkg/util/stringx"

type(
	ModelJsonStruct struct {
		TableName 		string			`json:"tableName"`
		Fields 			[]*ModelFieldJsonStruct	`json:"fields"`
		IsDisableBaseModel bool		`json:"isDisableBaseModel"`
		IsNeedUUID 			bool	`json:"isNeedUUID"`
	}
	ModelFieldJsonStruct struct {
		Name 		string		`json:"name"`
		Type 		string		`json:"type"`
		Size 		int			`json:"size"`
		Comment 	string		`json:"comment"`
	}
)

func (m *ModelJsonStruct)CamelName() string {
	return stringx.ToCamel(m.TableName)
}
func (m *ModelFieldJsonStruct)CamelName() string {
	return stringx.ToCamel(m.Name)
}


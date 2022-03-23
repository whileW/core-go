package model

type (
	ModelTempStruct struct {
		IsDisabledBaseModel bool
		IsHaveUuid 			bool
		Name 				string
		CamelName			string
		Fields 				[]ModelFieldTempStruct
	}
	ModelFieldTempStruct struct {
		Name 				string
		Type 				string
		Tag 				string
		HasComment			bool
		Comment				string
	}
)


package model

import (
	"bytes"
	"fmt"
	"github.com/whileW/core-go/tool/pkg/gen/model/template"
	"github.com/whileW/core-go/tool/pkg/util/filex"
	"github.com/whileW/core-go/tool/pkg/util/templatex"
)

func Gen(m []*ModelJsonStruct) ([]*filex.FileOrDir,error) {
	fs := []*filex.FileOrDir{}
	for _,t := range m {
		f := &filex.FileOrDir{}
		if err := genModel(t,f);err != nil {
			return nil,err
		}
		fs = append(fs, f)
	}
	return fs,nil
}

func genModel(m *ModelJsonStruct,f *filex.FileOrDir) error {
	var (
		_imports = ""
		_struct = ""
		_selfMethods = ""
		_methods = ""
		_err error
	)

	_imports,_err = genImports(m.IsNeedUUID,m.IsDisableBaseModel)
	if _err != nil {
		return _err
	}

	_struct,_err = genStruct(m)
	if _err != nil {
		return _err
	}

	_selfMethods,_err = genSelfMethod(m)
	if _err != nil {
		return _err
	}

	output, err := templatex.ExecuteTemplate(template.ModelT,map[string]interface{}{
		"imports":_imports,
		"struct":_struct,
		"selfMethods":_selfMethods,
		"methods":_methods,
	})
	if err != nil {
		return err
	}
	f.Data = output.Bytes()
	return nil
}

func genImports(isNeedUUID ,isDisabledBaseModel bool) (string,error) {
	output, err := templatex.ExecuteTemplate(template.ImportT,map[string]interface{}{
		"isNeedUUID":isNeedUUID,
		"isNeedBaseModel":!isDisabledBaseModel,
	})
	if err != nil {
		return "",err
	}
	return output.String(),nil
}
func genStruct(m *ModelJsonStruct) (string,error) {
	var (
		_fields = []string{}
		_err error
	)

	_fields,_err = genFields(m.Fields,m.IsDisableBaseModel,m.IsNeedUUID)
	if _err != nil {
		return "",_err
	}

	output, err := templatex.ExecuteTemplate(template.StructT,map[string]interface{}{
		"camelName":m.CamelName(),
		"fields":_fields,
	})
	if err != nil {
		return "",err
	}
	return output.String(),nil
}
func genFields(fs []*ModelFieldJsonStruct,isDisabledBaseModel,isNeedUUID bool) ([]string,error) {
	var (
		_fields = []string{}
	)

	if !isDisabledBaseModel {
		_fields = append(_fields, "utils.BaseModel")
	}
	if isNeedUUID {
		_fields = append(_fields, fmt.Sprintf("UUID string `json:\"uuid\"`"))
	}

	var generate_tag = func(t *ModelFieldJsonStruct) string {
		return fmt.Sprintf("`json:\"%s\"`",t.Name)
	}
	var generate_type = func(t *ModelFieldJsonStruct) string {
		switch t.Type {
		case "uuid":
			return "string"
		case "string":
			fallthrough
		case "int":
			fallthrough
		case "int32":
			fallthrough
		case "int64":
			fallthrough
		case "float64":
			fallthrough
		case "float32":
			return t.Type
		default:
			return t.Type
		}
	}

	for _,t := range fs {
		buf := new(bytes.Buffer)
		err := template.FieldT.Execute(buf,map[string]interface{}{
			"camelName":t.CamelName(),
			"type":generate_type(t),
			"tag":generate_tag(t),
			"comment":t.Comment,
			"hasComment":t.Comment != "",
		})
		if err != nil {
			return _fields,err
		}
		_fields = append(_fields, buf.String())
	}
	return _fields,nil
}
func genSelfMethod(m *ModelJsonStruct) (string,error) {
	output, err := templatex.ExecuteTemplate(template.SelfMethodT,map[string]interface{}{
		"isNeedUUID":m.IsNeedUUID,
		"camelName":m.CamelName(),
		"name":m.TableName,
	})
	if err != nil {
		return "",err
	}
	return output.String(),nil
}
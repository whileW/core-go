package templatex

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
)

func ExecuteTemplate(tem *template.Template,data interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	if err := tem.Execute(buf, data); err != nil {
		return nil, errors.New(fmt.Sprintf("template execute error:%v",err))
	}
	//formatOutput, err := goformat.Source(buf.Bytes())
	//if err != nil {
	//	return nil, errors.New(fmt.Sprintf( "go format error:%v,%s",err, buf.String()))
	//}
	//
	//buf.Reset()
	//buf.Write(formatOutput)
	return buf, nil
}

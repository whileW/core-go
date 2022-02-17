package main

import (
	"github.com/whileW/core-go/example/api"
	"github.com/whileW/core-go/pkg/application"
	"github.com/whileW/core-go/pkg/orm"
	"github.com/whileW/core-go/pkg/system_variable"
)

func main()  {
	system_variable.SystemName = "example"

	application.AppendInitializeFns(orm.Initialize)
	application.StartUp(api.NewApiService())
	application.KeepAlive()
}
package datasource

import (
	"github.com/whileW/core-go/pkg/conf/datasource/file"
	"github.com/whileW/core-go/pkg/flag"
)

var (
	config_datasource_type = ""
	default_config_datasource_type = file.DataSourceFile
)


func init() {
	flag.Register(&flag.StringFlag{
		Name: "config_datasource_type",
		Usage: "--config_datasource_type=file",
		Default:default_config_datasource_type,
		Variable:&config_datasource_type,
		EnvVar:"CONFIG_DATASOURCE_TYPE",
	})

	// register
	Register(file.DataSourceFile, func() (DataSource, error) {
		return file.NewDataSource()
	})
}

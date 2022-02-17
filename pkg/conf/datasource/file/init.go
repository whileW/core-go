package file

import (
	"github.com/whileW/core-go/pkg/flag"
)

var (
	config_file_path 	string
)
func init() {
	flag.Register(&flag.StringFlag{
		Name:	"config_file_path",
		Usage: 	"--config_file_path=config.yaml",
		Default: "config.yaml",
		Variable: &config_file_path,
		EnvVar: "CONFIG_FILE_PATH",
	})
}

package template

var Import = `import (
	{{if .isNeedUUID}}"github.com/google/uuid"{{end}}
	"github.com/whileW/core-go/pkg/orm"
	{{if .isNeedBaseModel}}"github.com/whileW/core-go/utils"{{end}}
	"gorm.io/gorm"
)`

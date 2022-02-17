package datasource

import (
	"errors"
	"fmt"
)

var(
	//ErrConfigAddr not config addr
	//ErrEmptyConfigAddr = errors.New("parse config error, config addr is empty")
	// ErrInvalidDataSource defines an error that the scheme has been registered
	//ErrInvalidDataSource = errors.New("invalid config data source, please make sure the scheme has been registered")
	datasourceBuilders   = make(map[string]DataSourceCreatorFunc)
)

// DataSourceCreatorFunc represents a dataSource creator function
type DataSourceCreatorFunc func() (DataSource,error)

// DataSource ...
type DataSource interface {
	ReadConfig() (map[string]interface{}, error)
	//IsConfigChanged() <-chan struct{}
	//io.Closer
}

func Register(scheme string, creator DataSourceCreatorFunc) {
	datasourceBuilders[scheme] = creator
}

//NewDataSource ..
func NewDataSource() (DataSource, error) {
	creatorFunc, exist := datasourceBuilders[config_datasource_type]
	if !exist {
		return nil, errors.New(fmt.Sprintf("invalid config data source[%s], please make sure the scheme has been registered",config_datasource_type))
	}
	return creatorFunc()
}

package orm

import (
	"github.com/whileW/core-go/conf"
	"gorm.io/gorm"
)

var orm *Orm

type Orm struct {
	//IsInit			bool					//是否初始化
	dbs 			map[string]*gorm.DB		//db 链接
}
func NewOrm() *Orm {
	return &Orm{
		dbs: map[string]*gorm.DB{},
	}
}
func GetOrm() *Orm {
	return orm
}

func (m *Orm)Set(name string,db *gorm.DB)  {
	m.dbs[name] = db
}
func (m *Orm)Get(name string) *gorm.DB {
	if name == "" {
		name = conf.GetConf().SysSetting.SystemName
	}
	if v,ok := m.dbs[name];ok {
		if conf.GetConf().SysSetting.Env != "release" {
			v = v.Debug()
		}
		return v
	}else {
		panic("没有该DB实例")
		return nil
	}
}
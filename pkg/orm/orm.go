package orm

import (
	"errors"
	"fmt"
	"github.com/whileW/core-go/pkg/conf"
	"github.com/whileW/core-go/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//func init_by_flag(dsn string)  {
//	c := &Config{
//		Name: "",
//		DSN: dsn,
//	}
//	if err := initialize(c);err != nil {
//		fmt.Println(fmt.Sprintf("initialize orm by flag error. para:%v, err:%v",c,err))
//		return
//	}
//}

var (
	DSN		string
	// 最大活动连接数
	MaxOpenConns	int
	// 最大空闲连接数
	MaxIdleConns 	int
)
var defaultDb *gorm.DB
func Initialize() (err error) {
	fmt.Println("开始初始化 orm ...")
	DSN = conf.GetStringd("orm.dsn",DSN)
	if DSN == "" {
		return errors.New("orm 初始化异常：DSN is empty")
	}
	defaultDb, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		Logger:&ormLoger{Loger:log.WithModule("orm").Clone()},
	})
	if err != nil {
		return err
	}
	db,err := defaultDb.DB()
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(MaxOpenConns)
	db.SetMaxIdleConns(MaxIdleConns)
	return nil
}


func GetDB() *gorm.DB {
	if defaultDb != nil {
		return defaultDb
	}
	panic("please init db first")
}
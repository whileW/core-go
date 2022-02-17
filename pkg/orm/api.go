package orm

import (
	"gorm.io/gorm"
)

type Option func(*gorm.DB)

func SearchOption_UUID(uuid string) Option {
	return func(db *gorm.DB) {
		db.Where("uuid = ?",uuid)
	}
}

func Option_TableName(tableName string) Option {
	return func(db *gorm.DB) {
		db.Table(tableName)
	}
}
func GetFirstRecord(db *gorm.DB, in interface{}, options ...Option) error {
	db = db.Session(&gorm.Session{})
	for _, option := range options {
		option(db)
	}
	return db.First(in).Error
}
func GetListRecord(db *gorm.DB, in interface{}, options ...Option) error {
	db = db.Session(&gorm.Session{})
	for _, option := range options {
		option(db)
	}
	return db.Find(in).Error
}
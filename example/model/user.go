package model

import (
	"github.com/google/uuid"
	"github.com/whileW/core-go/pkg/orm"
	"github.com/whileW/core-go/utils"
	"gorm.io/gorm"
)

type User struct {
	utils.BaseModel
	UUID 			string			`json:"uuid"`
	Name 			string			`json:"name"`
	Role 			int				`json:"role"`
	Status 			int				`json:"status"`
}

func (User)TableName() string {
	return "user"
}
func (t *User)BeforeCreate(tx *gorm.DB) error {
	if t.UUID == ""  {
		t.UUID = uuid.New().String()
	}
	return nil
}
func SearchUserList(db *gorm.DB,opts ...orm.Option) ([]*User,error) {
	var data = []*User{}
	opts = append(opts, orm.Option_TableName(User{}.TableName()))
	return data,orm.GetListRecord(db,&data,opts...)
}
func SearchUserFirst(db *gorm.DB,opts ...orm.Option) (*User,error) {
	var data = &User{}
	opts = append(opts, orm.Option_TableName(User{}.TableName()))
	return data,orm.GetFirstRecord(db,&data,opts...)
}

func SearchUserOption_Name(name string) orm.Option {
	return func(db *gorm.DB) {
		db.Where("name = ?",name)
	}
}

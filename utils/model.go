package utils

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int            `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Deleted   gorm.DeletedAt `gorm:"index"`
}

type BaseModelV1 struct {
	CreatedAt time.Time      `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:datetime"`
	Deleted   gorm.DeletedAt `gorm:"index"`
}

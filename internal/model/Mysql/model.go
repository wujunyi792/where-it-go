package Mysql

import (
	"gorm.io/gorm"
)

type Example struct {
	gorm.Model
}

func (Example) TableName() string {
	return "examples"
}

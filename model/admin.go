package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20)"`
	Password string
	Avatar   string
}

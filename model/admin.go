package model

import "gorm.io/gorm"

// Admin
// @Description: admin model
type Admin struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20)"`
	Password string
	Avatar   string
}

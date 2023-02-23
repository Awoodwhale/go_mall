package model

import "gorm.io/gorm"

// Notice
// @Description: 通知model
type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}

package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	BossID    uint `gorm:"not null"`
	Num       uint `gorm:"not null"` // 数量
	MaxNum    uint `gorm:"not null"` // 限制购买数量
	Check     bool // 是否支付
}

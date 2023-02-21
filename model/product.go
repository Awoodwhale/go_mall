package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string
	Category      uint
	Title         string
	Info          string
	ImagePath     string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        uint
	BossName      string
	BossAvatar    string
}

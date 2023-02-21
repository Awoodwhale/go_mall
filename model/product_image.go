package model

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ProductID uint `gorm:"not null"`
	ImagePath string
}

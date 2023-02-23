package model

import "gorm.io/gorm"

// ProductImage
// @Description: 商品图片 model
type ProductImage struct {
	gorm.Model
	ProductID uint `gorm:"not null"`
	ImagePath string
}

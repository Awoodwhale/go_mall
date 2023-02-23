package model

import "gorm.io/gorm"

// Carousel
// @Description: 轮播图 model
type Carousel struct {
	gorm.Model
	ImagePath string
	ProductID uint `gorm:"not null"` // 产品ID
}

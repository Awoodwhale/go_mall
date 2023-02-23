package model

import "gorm.io/gorm"

// Category
// @Description: 商品分类 model
type Category struct {
	gorm.Model
	CategoryName string // 分类名称
}

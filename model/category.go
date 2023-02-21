package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string // 分类名称
}

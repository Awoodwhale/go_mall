package dao

import (
	"go_mall/model"
)

// migration
// @Description: 数据迁移
func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Address{},
			&model.Admin{},
			&model.Category{},
			&model.Carousel{},
			&model.Cart{},
			&model.Notice{},
			&model.Product{},
			&model.ProductImage{},
			&model.Order{},
			&model.Favorite{})
	if err != nil {
		panic(err)
	}
	return
}

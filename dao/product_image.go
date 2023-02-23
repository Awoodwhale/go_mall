package dao

import (
	"context"
	"go_mall/model"
	"gorm.io/gorm"
)

// ProductImageDao
// @Description: product image dao
type ProductImageDao struct {
	*gorm.DB
}

// NewProductImageDao
// @Description: 通过ctx获取product image dao
// @param c context.Context
// @return *ProductImageDao
func NewProductImageDao(c context.Context) *ProductImageDao {
	return &ProductImageDao{NewDBClient(c)}
}

// NewProductImageDaoByDB
// @Description: 通过db复用product image dao
// @param db *gorm.DB
// @return *ProductImageDao
func NewProductImageDaoByDB(db *gorm.DB) *ProductImageDao {
	return &ProductImageDao{db}
}

// CreateProductImage
// @Description: 创建product image到MySQL
// @receiver dao *ProductImageDao
// @param productImage *model.ProductImage
// @return error
func (dao *ProductImageDao) CreateProductImage(productImage *model.ProductImage) error {
	return dao.DB.Model(&model.ProductImage{}).Create(&productImage).Error
}

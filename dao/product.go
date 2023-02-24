package dao

import (
	"context"
	"go_mall/model"
	"gorm.io/gorm"
)

// ProductDao
// @Description: product dao
type ProductDao struct {
	*gorm.DB
}

// NewProductDao
// @Description: 获取product dao 通过ctx
// @param c context.Context
// @return *ProductDao
func NewProductDao(c context.Context) *ProductDao {
	return &ProductDao{NewDBClient(c)}
}

// NewProductDaoByDB
// @Description: 通过db复用product dao
// @param db *gorm.DB
// @return *ProductDao
func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// CreateProduct
// @Description: 创建product
// @receiver dao *ProductDao
// @param product *model.Product
// @return error
func (dao *ProductDao) CreateProduct(product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}

// GetProductByName
// @Description: 通过name获取product
// @receiver dao *ProductDao
// @param name string
// @return product *model.Product
// @return err error
func (dao *ProductDao) GetProductByName(name string) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("name=?", name).First(&product).Error
	return
}

// CountProductByCondition
// @Description: 通过condition获取product
// @receiver dao *ProductDao
// @param condition map[string]any
// @return total int64
// @return err error
func (dao *ProductDao) CountProductByCondition(condition map[string]any) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

// ListProductByCondition
// @Description: 通过condition获取product列表
// @receiver dao *ProductDao
// @param condition map[string]any
// @param page *model.BasePage
// @return products []*model.Product
// @return err error
func (dao *ProductDao) ListProductByCondition(condition map[string]any, page *model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Preload("Category").Where(condition).
		Offset(int((page.PageNum - 1) * page.PageSize)). // 分页
		Limit(int(page.PageSize)).Find(&products).Error
	return
}

func (dao *ProductDao) SearchProduct(info string, page *model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("title like ? or info like ?", "%"+info+"%", "%"+info+"%").
		Offset(int((page.PageNum - 1) * page.PageSize)). // 分页
		Limit(int(page.PageSize)).Find(&products).Error
	return
}

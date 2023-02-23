package dao

import (
	"context"
	"go_mall/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

// NewCarouselDao
// @Description: 通过ctx获取carouselDao
// @param ctx context.Context
// @return *UserDao
func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

// NewCarouselDaoByDB
// @Description: 通过db获取carouselDao
// @param db *gorm.DB
// @return *UserDao
func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// GetCarouselById
// @Description: 通过id获取carousel
// @receiver dao *CarouselDao
// @param id uint
// @return carousel *model.Carousel
// @return err error
func (dao *CarouselDao) GetCarouselById(id uint) (carousel *model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Where("id=?", id).First(&carousel).Error
	return
}

// ListCarousel
// @Description: 获取carousel的列表
// @receiver dao *CarouselDao
// @return carousels []model.Carousel
// @return err error
func (dao *CarouselDao) ListCarousel() (carousels []model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&carousels).Error
	return
}

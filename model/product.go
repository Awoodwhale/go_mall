package model

import (
	"go_mall/cache"
	"go_mall/pkg/utils"
	"gorm.io/gorm"
	"strconv"
)

// Product
// @Description: 商品 model
type Product struct {
	gorm.Model
	Name          string
	CategoryID    uint `gorm:"not null"`
	Category      Category
	Title         string
	Info          string `gorm:"size:1000"`
	ImagePath     string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        uint
	BossName      string
	BossAvatar    string
}

// GetView
// @Description: 获取product在redis中的view数量
// @receiver p *Product
// @return res uint64
func (p *Product) GetView() (res uint64) {
	viewCnt, err := cache.RedisClient.Get(cache.ProductViewKey(p.ID)).Result()
	if err != nil {
		utils.Logger.Errorln("model product get view count from redis, ", err)
		return 0
	}
	res, err = strconv.ParseUint(viewCnt, 10, 64)
	if err != nil {
		utils.Logger.Errorln("model product parse view count, ", err)
		return 0
	}
	return
}

// AddView
// @Description: 添加product在redis中view的cnt
// @receiver p *Product
func (p *Product) AddView() {
	cache.RedisClient.Incr(cache.ProductViewKey(p.ID))
}

package serializer

import (
	"go_mall/conf"
	"go_mall/model"
)

// Product
// @Description: product vo
type Product struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImagePath     string `json:"image_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint64 `json:"view"`
	CreateAt      int64  `json:"create_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossID        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

// BuildProduct
// @Description: 构建resp的product vo
// @param product *model.Product
// @return *Product
func BuildProduct(product *model.Product) *Product {
	var (
		productImagePath string
		avatarImagePath  string
	)
	if product.ImagePath != "" {
		productImagePath = conf.ImgHost + ":" + conf.ImgPort + conf.ProductPath + product.ImagePath
	}
	if product.BossAvatar != "" {
		avatarImagePath = conf.ImgHost + ":" + conf.ImgPort + conf.AvatarPath + product.BossAvatar
	}
	return &Product{
		ID:            product.ID,
		Name:          product.Name,
		CategoryId:    product.CategoryID,
		Title:         product.Title,
		Info:          product.Info,
		ImagePath:     productImagePath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		View:          product.GetView(),
		CreateAt:      product.CreatedAt.Unix(),
		Num:           product.Num,
		OnSale:        product.OnSale,
		BossID:        product.BossID,
		BossName:      product.BossName,
		BossAvatar:    avatarImagePath,
	}
}

// BuildProducts
// @Description: 构建前端展示的product list
// @param items []*model.Product
// @return products []*Product
func BuildProducts(items []*model.Product) (products []*Product) {
	for _, item := range items {
		products = append(products, BuildProduct(item))
	}
	return products
}

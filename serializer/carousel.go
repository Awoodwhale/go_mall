package serializer

import "go_mall/model"

// Carousel
// @Description: carousel vo
type Carousel struct {
	Id        uint   `json:"id"`
	ProductId uint   `json:"product_id"`
	ImagePath string `json:"image_path"`
	CreateAt  int64  `json:"create_at"`
}

// BuildCarousel
// @Description: 单个carousel
// @param item *model.Carousel
// @return Carousel
func BuildCarousel(item *model.Carousel) Carousel {
	return Carousel{
		Id:        item.ID,
		ProductId: item.ProductID,
		ImagePath: item.ImagePath,
		CreateAt:  item.CreatedAt.Unix(),
	}
}

// BuildCarousels
// @Description: carousel vo list
// @param items []model.Carousel
// @return carousels []Carousel
func BuildCarousels(items []model.Carousel) (carousels []Carousel) {
	for _, v := range items {
		carousels = append(carousels, BuildCarousel(&v))
	}
	return
}

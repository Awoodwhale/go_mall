package service

import (
	"context"
	"go_mall/dao"
	"go_mall/pkg/e"
	"go_mall/pkg/utils"
	"go_mall/serializer"
)

// CarouselService
// @Description: 轮播图service
type CarouselService struct {
}

// ListCarousel
// @Description: 获取轮播图列表
// @receiver service *CarouselService
// @param c context.Context
// @return serializer.Response
func (service *CarouselService) ListCarousel(c context.Context) serializer.Response {
	var (
		code        = e.Success
		err         error
		carouselDao = dao.NewCarouselDao(c)
	)

	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		code = e.ErrorWithSQL
		utils.Logger.Errorln("service carousel ListCarousel sql error,", err.Error())
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}

package v1

import (
	"github.com/gin-gonic/gin"
	"go_mall/pkg/utils"
	"go_mall/service"
	"net/http"
)

// ListCarousel
// @Description: 获取轮播图carousel列表
// @param c *gin.Context
func ListCarousel(c *gin.Context) {
	var listCarouselService service.CarouselService
	if err := c.ShouldBind(&listCarouselService); err == nil {
		res := listCarouselService.ListCarousel(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, &listCarouselService))
		utils.Logger.Errorln("ListCarousel api", err)
	}
}

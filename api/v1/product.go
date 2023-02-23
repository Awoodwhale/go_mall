package v1

import (
	"github.com/gin-gonic/gin"
	"go_mall/pkg/utils"
	"go_mall/service"
	"net/http"
)

// CreateProduct
// @Description: 创建product
// @param c *gin.Context
func CreateProduct(c *gin.Context) {
	var productService service.ProductService
	form, _ := c.MultipartForm()
	files := form.File["file"]
	if err := c.ShouldBind(&productService); err == nil {
		claims, _ := c.Get("claims")
		res := productService.CreateProduct(c.Request.Context(), claims.(*utils.Claims).ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, &productService))
		utils.Logger.Errorln("CreateProduct api", err)
	}
}

// ListProduct
// @Description: 获取product列表
// @param c *gin.Context
func ListProduct(c *gin.Context) {
	var productService service.ProductService
	if err := c.ShouldBind(&productService); err == nil {
		res := productService.ListProduct(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, &productService))
		utils.Logger.Errorln("CreateProduct api", err)
	}
}

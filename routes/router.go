package routes

import (
	"github.com/gin-gonic/gin"
	api "go_mall/api/v1"
	"go_mall/middleware"
	"go_mall/pkg/e"
	"go_mall/serializer"
	"net/http"
)

// NewRouter
// @Description: 设置gin router
// @return *gin.Engine
func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Cors())                    // cors中间件
	router.StaticFS("/static", http.Dir("./static")) // 设置fs路径

	v1 := router.Group("api/v1") // v1版本的api
	{
		// ping test
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, serializer.Response{
				Code:    e.Success,
				Message: "pong",
			})
		})

		v1.POST("user/register", api.UserRegister)              // 用户注册
		v1.POST("user/login", api.UserLogin)                    // 用户登录
		v1.GET("user/validate_email/:token", api.ValidateEmail) // 验证邮件有效
		v1.GET("carousel", api.ListCarousel)                    // 轮播图
		v1.GET("product", api.ListProduct)                      // 商品列表

		authed := v1.Group("/")      // 需要登录保护
		authed.Use(middleware.JWT()) // jwt校验中间件
		{
			// 用户操作
			authed.PUT("user/update", api.UserUpdate)        // 用户更新
			authed.POST("user/avatar", api.UserUploadAvatar) // 更新用户头像
			authed.POST("user/money", api.UserShowMoney)     // 显示用户金额
			authed.POST("user/send_email", api.SendEmail)    // 发送邮件

			// 商品操作
			authed.POST("product", api.CreateProduct) // 创建商品
		}
	}

	return router
}

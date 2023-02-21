package routes

import (
	"github.com/gin-gonic/gin"
	api "go_mall/api/v1"
	"go_mall/middleware"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors()) // 注册中间件
	router.StaticFS("/static", http.Dir("./static"))
	v1 := router.Group("api/v1") // v1版本的api
	{
		// ping test
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "pong")
		})

		v1.POST("user/register", api.UserRegister)
	}
	return router
}

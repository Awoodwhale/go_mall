package v1

import (
	"github.com/gin-gonic/gin"
	"go_mall/pkg/utils"
	"go_mall/service"
	"net/http"
)

func UserRegister(c *gin.Context) {
	/**
	 * UserRegister
	 * @Description: 用户注册
	 * @param c gin.Context
	 */
	var userRegisterService service.UserService
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UserLogin(c *gin.Context) {
	/**
	 * UserLogin
	 * @Description: 用户登录
	 * @param c	gin.Context
	 */
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UserUpdate(c *gin.Context) {
	/**
	 * UserUpdate
	 * @Description: 用户更新
	 * @param c gin.Context
	 */
	var userUpdateService service.UserService
	if err := c.ShouldBind(&userUpdateService); err == nil {
		claims, _ := c.Get("claims") // 从ctx拿出claims
		res := userUpdateService.Update(c.Request.Context(), claims.(*utils.Claims).ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func UserUploadAvatar(c *gin.Context) {
	var userAvatarService service.UserService
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize, fileName := fileHeader.Size, fileHeader.Filename
	if err := c.ShouldBind(&userAvatarService); err == nil {
		claims, _ := c.Get("claims")
		res := userAvatarService.UploadAvatar(c.Request.Context(), claims.(*utils.Claims).ID, file, fileSize, fileName)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

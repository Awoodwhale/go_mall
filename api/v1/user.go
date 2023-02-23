package v1

import (
	"github.com/gin-gonic/gin"
	"go_mall/pkg/e"
	"go_mall/pkg/utils"
	"go_mall/serializer"
	"go_mall/service"
	"net/http"
	"time"
)

// UserRegister
// @Description: 用户注册api
// @param c *gin.Context
func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    e.Error,
			Message: e.HandleBindingError(err, &userRegisterService),
		})
	}
}

// UserLogin
// @Description: 用户登录api
// @param c *gin.Context
func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    e.Error,
			Message: e.HandleBindingError(err, &userLoginService),
		})
	}
}

// UserUpdate
// @Description: 用户更新api
// @param c *gin.Context
func UserUpdate(c *gin.Context) {
	var userUpdateService service.UserService
	if err := c.ShouldBind(&userUpdateService); err == nil {
		claims, _ := c.Get("claims") // 从ctx拿出claims
		res := userUpdateService.Update(c.Request.Context(), claims.(*utils.Claims).ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    e.Error,
			Message: e.HandleBindingError(err, &userUpdateService),
		})
	}
}

// UserUploadAvatar
// @Description: 上传用户头像api
// @param c *gin.Context
func UserUploadAvatar(c *gin.Context) {
	var userAvatarService service.UserService
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize, fileName := fileHeader.Size, fileHeader.Filename
	if err := c.ShouldBind(&userAvatarService); err == nil {
		claims, _ := c.Get("claims")
		res := userAvatarService.UploadAvatar(c.Request.Context(), claims.(*utils.Claims).ID, file, fileSize, fileName)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    e.Error,
			Message: e.HandleBindingError(err, &userAvatarService),
		})
	}
}

// UserShowMoney
// @Description: 显示用户的money的api
// @param c *gin.Context
func UserShowMoney(c *gin.Context) {
	var showMoneyService service.ShowMoneyService
	if err := c.ShouldBind(&showMoneyService); err == nil {
		claims, _ := c.Get("claims")
		res := showMoneyService.ShowMoney(c.Request.Context(), claims.(*utils.Claims).ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    e.Error,
			Message: e.HandleBindingError(err, &showMoneyService),
		})
	}
}

// SendEmail
// @Description: 发送邮箱api
// @param c *gin.Context
func SendEmail(c *gin.Context) {
	var sendEmailService service.SendEmailService
	if err := c.ShouldBind(&sendEmailService); err == nil {
		claims, _ := c.Get("claims")
		res := sendEmailService.SendEmail(c.Request.Context(), claims.(*utils.Claims).ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    e.Error,
			Message: e.HandleBindingError(err, &sendEmailService),
		})
	}
}

// ValidateEmail
// @Description: 验证邮箱api
// @param c *gin.Context
func ValidateEmail(c *gin.Context) {
	var (
		code                 = e.Success
		token                = c.Param("token") // 获取token
		validateEmailService service.ValidateEmailService
	)
	claims, err := utils.ParseEmailToken(token) // 解析邮箱token
	if err != nil {
		// parse错误
		code = e.ErrorWithParseToken
	} else if time.Now().Unix() > claims.ExpiresAt {
		code = e.ErrorWithExpiredToken // token过期（其实不会走到这个分支，过期上面err会报错）
	}
	err = c.ShouldBind(&validateEmailService)
	if err != nil {
		code = e.Error
	}
	if code == e.Success {
		res := validateEmailService.ValidateEmail(c.Request.Context(), claims)
		c.JSON(http.StatusOK, res)
	} else {
		msg := e.HandleBindingError(err, &validateEmailService)
		if msg == "" {
			msg = e.GetMessageByCode(code)
		}
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    code,
			Message: msg,
		})
	}
}

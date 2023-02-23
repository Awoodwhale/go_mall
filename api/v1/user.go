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
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    e.Error,
			Message: e.HandleBindingError(err, &userRegisterService),
		})
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
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    e.Error,
			Message: e.HandleBindingError(err, &userLoginService),
		})
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
		c.JSON(http.StatusBadRequest, serializer.Response{
			Code:    e.Error,
			Message: e.HandleBindingError(err, &userUpdateService),
		})
	}
}

func UserUploadAvatar(c *gin.Context) {
	/**
	 * UserUploadAvatar
	 * @Description: 上传用户头像
	 * @param c
	 */
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

func UserShowMoney(c *gin.Context) {
	/**
	 * UserShowMoney
	 * @Description: 显示用户的money
	 * @param c
	 */
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

func SendEmail(c *gin.Context) {
	/**
	 * SendEmail
	 * @Description: 发送邮箱
	 * @param c
	 */
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

func ValidateEmail(c *gin.Context) {
	/**
	 * ValidateEmail
	 * @Description: 验证邮箱
	 * @param c
	 */
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

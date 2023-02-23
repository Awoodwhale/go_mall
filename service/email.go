package service

import (
	"context"
	"go_mall/conf"
	"go_mall/dao"
	"go_mall/model"
	"go_mall/pkg/e"
	"go_mall/pkg/utils"
	"go_mall/serializer"
	"gopkg.in/mail.v2"
)

// SendEmailService
// @Description: 发送邮箱的service
type SendEmailService struct {
	Email    string `json:"email" form:"email" binding:"required,email" msg:"邮箱格式错误"`
	Password string `json:"password" form:"password"`
	// 1 --> 绑定邮箱
	// 2 --> 解绑邮箱
	// 3 --> 修改密码
	OperationType uint `json:"operation_type" form:"operation_type" binding:"required" msg:"操作类型不能为空"`
}

// ValidateEmailService
// @Description: 验证邮箱的service
type ValidateEmailService struct {
	// 空结构体
}

// SendEmail
// @Description: 发送邮箱
// @receiver service *SendEmailService
// @param c context.Context
// @param uid uint
// @return serializer.Response
func (service *SendEmailService) SendEmail(c context.Context, uid uint) serializer.Response {
	var (
		code    = e.Success
		err     error
		address string
		notice  *model.Notice
		user    *model.User
		userDao = dao.NewUserDao(c)
	)

	// 操作类型判断
	if service.OperationType < 0 || service.OperationType > 3 {
		code = e.ErrorWithOperationType
	}

	// 读取用户数据
	user, err = userDao.GetUserById(uid)
	if err != nil {
		code = e.ErrorWithSQL
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	if service.OperationType == 1 { // 绑定邮箱check
		if user.Email == service.Email {
			// 绑定重复邮箱
			code = e.ErrorWithSameEmail
		}

	} else if service.OperationType == 2 { // 删除邮箱check
		if user.Email == "no binding email" {
			// 邮箱不存在，无需删除
			code = e.ErrorWithNotExistEmail
		}
	} else if service.OperationType == 3 { // 修改密码check
		if service.Password == "" || len(service.Password) > 20 {
			// 密码为空或者太长了
			code = e.ErrorWithPassword
		} else if user.Email != service.Email {
			// 邮箱校验不一致
			code = e.ErrorWithCheckEmail
		}
	}

	if code != e.Success {
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
		}
	}

	token, err := utils.GenerateEmailToken(uid, service.OperationType, service.Email, service.Password)
	if err != nil {
		code = e.ErrorWithGenToken
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	noticeDao := dao.NewNoticeDao(c)
	notice, err = noticeDao.GetNoticeById(service.OperationType)
	if err != nil {
		code = e.ErrorWithSQL
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	address = conf.ValidEmail + token // 发送方
	mailMsg := notice.Text + address
	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail) // 发送方
	m.SetHeader("To", service.Email)    // 接收方
	m.SetHeader("Subject", "go_mall 邮箱信息")
	m.SetBody("text/html", mailMsg)
	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpToken)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	// 发送邮箱
	if err = d.DialAndSend(m); err != nil {
		code = e.ErrorWithSendEmail
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	return serializer.Response{
		Code:    code,
		Message: e.GetMessageByCode(code),
	}
}

// ValidateEmail
// @Description: 验证邮箱
// @receiver service *ValidateEmailService
// @param c context.Context
// @param claims *utils.EmailClaims
// @return serializer.Response
func (service *ValidateEmailService) ValidateEmail(c context.Context, claims *utils.EmailClaims) serializer.Response {
	var (
		code          = e.Success
		err           error
		userID        = claims.UserID
		email         = claims.Email
		password      = claims.Password
		operationType = claims.OperationType
		userDao       = dao.NewUserDao(c)
	)

	// 获取用户信息
	user, err := userDao.GetUserById(userID)
	if err != nil {
		code = e.ErrorWithSQL
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	if operationType == 1 {
		// 1. 绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		// 2. 解绑邮箱
		user.Email = "no binding email"
	} else if operationType == 3 {
		err = user.SetPassword(password)
		if err != nil {
			code = e.ErrorWithEncryption
			return serializer.Response{
				Code:    code,
				Message: e.GetMessageByCode(code),
				Error:   err.Error(),
			}
		}
	}

	// 更新MySQL
	err = userDao.UpdateUserById(userID, user)
	if err != nil {
		code = e.ErrorWithSQL
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	// 成功就返回用户的信息
	return serializer.Response{
		Code:    code,
		Message: e.GetMessageByCode(code),
		Data:    serializer.BuildUser(user),
	}
}

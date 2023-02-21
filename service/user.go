package service

import (
	"context"
	"go_mall/dao"
	"go_mall/model"
	"go_mall/pkg/e"
	"go_mall/pkg/utils"
	"go_mall/serializer"
	"strconv"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` // 前端验证
}

func (service UserService) Register(ctx context.Context) serializer.Response {
	/**
	 * Register
	 * @Description: 用户注册
	 * @receiver service
	 * @param ctx
	 * @return serializer.Response
	 */
	code := e.Success
	if !utils.CheckKey(service.Key) { // 密钥检测
		code = e.ErrorWithKey
		return serializer.Response{
			Code: code, Message: e.GetMessageByCode(code),
		}
	}

	// 进行对称加密
	utils.Encrypt.SetKey(service.Key)

	userDao := dao.NewUserDao(ctx)
	// 查看这个名称的用户是否存在
	if exist := userDao.ExistOrNotByUserName(service.UserName); exist {
		// 用户已经存在了
		code = e.ErrorWithExistUser
		return serializer.Response{
			Code: code, Message: e.GetMessageByCode(code),
		}
	}

	user := model.User{
		UserName: service.UserName,
		NickName: service.NickName,
		Avatar:   "avatar.png",
		Status:   model.ActiveUser,
		Money:    utils.Encrypt.AesEncoding(strconv.Itoa(model.InitMoney)), // 初始金额加密
	}
	// 用户密码加密
	if err := user.SetPassword(service.Password); err != nil {
		code = e.ErrorWithFailedEncryption
		return serializer.Response{
			Code: code, Message: e.GetMessageByCode(code),
		}
	}

	// 创建用户，写入MySQL
	if err := userDao.CreateUser(&user); err != nil {
		code = e.Error
		return serializer.Response{
			Code:    code,
			Message: err.Error(),
		}
	}

	// 注册成功
	return serializer.Response{
		Code:    code,
		Message: "注册成功",
		Data:    user,
	}
}

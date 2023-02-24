package service

import (
	"context"
	"go_mall/dao"
	"go_mall/model"
	"go_mall/pkg/e"
	"go_mall/pkg/utils"
	"go_mall/serializer"
	"mime/multipart"
	"strconv"
)

// UserService
// @Description: 用户提交的表单service
type UserService struct {
	UserName string `json:"user_name" form:"user_name"`
	NickName string `json:"nick_name" form:"nick_name"`
	Password string `json:"password" form:"password"`
	Key      string `json:"key" form:"key"` // 前端验证
}

// ShowMoneyService
// @Description: 展示money的表单
type ShowMoneyService struct {
	Key string `json:"key" form:"key" binding:"required" msg:"校验密钥不可为空"` // 验证密码
}

// Register
// @Description: 用户注册
// @receiver service *UserService
// @param ctx context.Context
// @return serializer.Response
func (service *UserService) Register(ctx context.Context) serializer.Response {
	var (
		code = e.Success
		err  error
	)

	// 判断表单是否完整
	if service.UserName == "" || service.NickName == "" || service.Password == "" {
		code = e.InvalidParams
		utils.Logger.Debugln("service user Register empty username or nickname or password")
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
		}
	}

	if !utils.CheckKey(service.Key) { // 密钥检测
		code = e.ErrorWithKey
		return serializer.Response{
			Code: code, Message: e.GetMessageByCode(code),
		}
	}

	// 查看这个名称的用户是否存在
	userDao := dao.NewUserDao(ctx)
	if _, exist := userDao.ExistOrNotByUserName(service.UserName); exist {
		// 用户已经存在了
		code = e.ErrorWithExistUser
		return serializer.Response{
			Code: code, Message: e.GetMessageByCode(code),
		}
	}

	user := &model.User{
		UserName: service.UserName,
		NickName: service.NickName,
		Avatar:   "avatar.png",
		Status:   model.ActiveUser,
		Money:    utils.AesEncoding(strconv.Itoa(model.InitMoney), service.Key), // 初始金额加密
	}
	// 用户密码加密
	if err = user.SetPassword(service.Password); err != nil {
		code = e.ErrorWithEncryption
		return serializer.Response{
			Code: code, Message: e.GetMessageByCode(code), Error: err.Error(),
		}
	}

	// 创建用户，写入MySQL
	if err = userDao.CreateUser(user); err != nil {
		code = e.ErrorWithSQL
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	// 注册成功
	return serializer.Response{
		Code:    code,
		Message: "注册成功",
		Data:    serializer.BuildUser(user), // 传给前端build后的user信息
	}
}

// Login
// @Description: 用户登录
// @receiver service *UserService
// @param ctx context.Context
// @return serializer.Response
func (service *UserService) Login(ctx context.Context) serializer.Response {
	var (
		code    = e.Success
		user    *model.User
		err     error
		userDao = dao.NewUserDao(ctx)
	)

	user, exist := userDao.ExistOrNotByUserName(service.UserName)
	if !exist {
		// 用户不存在
		code = e.ErrorWithNotExistUser
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
		}
	}
	if !user.CheckPassword(service.Password) { // 密码错误
		code = e.ErrorWithPassword
		return serializer.Response{Code: code, Message: e.GetMessageByCode(code)}
	}

	// token签发
	token, err := utils.GenerateJWT(user.ID, service.UserName, 0)
	if err != nil {
		code = e.ErrorWithGenToken
		return serializer.Response{Code: code, Message: e.GetMessageByCode(code), Error: err.Error()}
	}

	// 没有发生错误那么就是成功登录
	return serializer.Response{
		Code:    code,
		Message: "登录成功",
		Data:    serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}
}

// Update
// @Description: 修改用户信息
// @receiver service *UserService
// @param ctx context.Context
// @param uid uint
// @return serializer.Response
func (service *UserService) Update(ctx context.Context, uid uint) serializer.Response {
	var (
		code       = e.Success
		user       *model.User
		err        error
		needUpdate = false
		userDao    = dao.NewUserDao(ctx)
	)

	user, err = userDao.GetUserById(uid)
	if err != nil {
		// MySQL错误
		code = e.ErrorWithSQL
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	if service.NickName != "" && service.NickName != user.NickName { // 修改nickname
		user.NickName = service.NickName
		needUpdate = true
	}
	// 需要修改再去MySQL修改
	if needUpdate {
		err = userDao.UpdateUserById(uid, user)
		if err != nil {
			code = e.ErrorWithSQL
			return serializer.Response{
				Code:    code,
				Message: e.GetMessageByCode(code),
				Error:   err.Error(),
			}
		}
	}

	// 信息修改成功
	return serializer.Response{
		Code:    code,
		Message: "更新用户信息成功",
		Data:    serializer.BuildUser(user),
	}

}

// UploadAvatar
// @Description: 上传头像到本地
// @receiver service *UserService
// @param ctx context.Context
// @param uid uint
// @param file multipart.File
// @param size int64
// @param filename string
// @param
// @return serializer.Response
func (service *UserService) UploadAvatar(
	ctx context.Context,
	uid uint,
	file multipart.File,
	size int64,
	filename string,
) serializer.Response {
	var (
		code = e.Success
		user *model.User
		err  error
	)

	if size > 5242880 {
		// 大于5MB就不允许上传
		code = e.ErrorWithFileSize
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
		}
	}

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uid)
	if err != nil {
		code = e.ErrorWithSQL
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	// 头像保存到本地
	filePath, err := UploadAvatarImage(file, filename, uid)
	if err != nil {
		code = e.ErrorWithUploadAvatar
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	// 更新头像到MySQL
	user.Avatar = filePath
	err = userDao.UpdateUserById(uid, user)
	if err != nil {
		code = e.ErrorWithSQL
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	return serializer.Response{
		Code:    code,
		Message: "上传头像成功",
		Data:    serializer.BuildUser(user),
	}
}

// ShowMoney
// @Description: 显示用户金额
// @receiver service *ShowMoneyService
// @param ctx context.Context
// @param uid uint
// @return serializer.Response
func (service *ShowMoneyService) ShowMoney(ctx context.Context, uid uint) serializer.Response {
	var (
		code    = e.Success
		err     error
		user    *model.User
		userDao = dao.NewUserDao(ctx)
	)

	// 检测key
	if !utils.CheckKey(service.Key) {
		code = e.ErrorWithKey
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
		}
	}

	// 读取user的money
	user, err = userDao.GetUserById(uid)
	if err != nil {
		code = e.ErrorWithSQL
		return serializer.Response{
			Code:    code,
			Message: e.GetMessageByCode(code),
			Error:   err.Error(),
		}
	}

	// 返回money
	return serializer.Response{
		Code:    code,
		Message: "获取用户余额成功",
		Data:    serializer.BuildMoney(user, service.Key),
	}
}

package dao

import (
	"context"
	"go_mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

// ExistOrNotByUserName
// @Description: 通过username查询用户是否存在
// @receiver dao *UserDao
// @param name string
// @return user *model.User
// @return exist bool
func (dao *UserDao) ExistOrNotByUserName(name string) (user *model.User, exist bool) {
	err := dao.DB.Model(&model.User{}).Where("user_name=?", name).First(&user).Error
	if user == nil || err == gorm.ErrRecordNotFound { // 不存在
		return nil, false
	} else if user != nil {
		return user, true
	}
	return nil, false // 不存在
}

// CreateUser
// @Description: 创建用户
// @receiver dao *UserDao
// @param user *model.User
// @return error
func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

// GetUserById
// @Description: 通过ID获取user
// @receiver dao *UserDao
// @param uid uint
// @return user *model.User
// @return err error
func (dao *UserDao) GetUserById(uid uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", uid).First(&user).Error
	return
}

// UpdateUserById
// @Description: 通过id修改user信息
// @receiver dao *UserDao
// @param uid uint
// @param user *model.User
// @return err error
func (dao *UserDao) UpdateUserById(uid uint, user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", uid).Updates(&user).Error
	return
}

// NewUserDao
// @Description: 通过ctx获取userDao
// @param ctx context.Context
// @return *UserDao
func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// NewUserDaoByDB
// @Description: 通过db获取userDao
// @param db *gorm.DB
// @return *UserDao
func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

package dao

import (
	"context"
	"go_mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func (dao *UserDao) ExistOrNotByUserName(name string) (exist bool) {
	/**
	 * ExistOrNotByUserName
	 * @Description: 通过username查询用户是否存在
	 * @receiver dao
	 * @param name
	 * @return user
	 * @return exist
	 * @return err
	 */
	var user *model.User
	err := dao.DB.Model(&model.User{}).Where("user_name=?", name).First(&user).Error
	if user == nil || err == gorm.ErrRecordNotFound { // 不存在
		return false
	} else if user != nil {
		return true
	}
	return false // 不存在
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

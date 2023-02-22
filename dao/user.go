package dao

import (
	"context"
	"go_mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func (dao *UserDao) ExistOrNotByUserName(name string) (user *model.User, exist bool) {
	/**
	 * ExistOrNotByUserName
	 * @Description: 通过username查询用户是否存在
	 * @receiver dao
	 * @param name
	 * @return user
	 * @return exist
	 * @return err
	 */
	err := dao.DB.Model(&model.User{}).Where("user_name=?", name).First(&user).Error
	if user == nil || err == gorm.ErrRecordNotFound { // 不存在
		return nil, false
	} else if user != nil {
		return user, true
	}
	return nil, false // 不存在
}

func (dao *UserDao) CreateUser(user *model.User) error {
	/**
	 * CreateUser
	 * @Description: 创建用户
	 * @receiver dao
	 * @param user
	 * @return error
	 */
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

func (dao *UserDao) GetUserById(uid uint) (user *model.User, err error) {
	/**
	 * GetUserById
	 * @Description: 通过ID获取user
	 * @receiver dao
	 * @param uid
	 * @return user
	 * @return err
	 */
	err = dao.DB.Model(&model.User{}).Where("id=?", uid).First(&user).Error
	return
}

func (dao *UserDao) UpdateUserById(uid uint, user *model.User) (err error) {
	/**
	 * UpdateUserById
	 * @Description: 通过id修改user信息
	 * @receiver dao
	 * @param uid
	 * @param user
	 * @return error
	 */
	err = dao.DB.Model(&model.User{}).Where("id=?", uid).Updates(&user).Error
	return
}

func NewUserDao(ctx context.Context) *UserDao {
	/**
	 * NewUserDao
	 * @Description: 通过ctx获取userDao
	 * @param ctx
	 * @return *UserDao
	 */
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	/**
	 * NewUserDaoByDB
	 * @Description: 通过db获取userDao
	 * @param db
	 * @return *UserDao
	 */
	return &UserDao{db}
}

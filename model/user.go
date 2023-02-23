package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	PasswordCost = 12       // 密码加密难度
	ActiveUser   = "active" // 激活的用户状态
	InitMoney    = 114514   // 注册后用户的初始金额
)

// User
// @Description: 用户 model
type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(50)"`
	Email    string `gorm:"type:varchar(50)"`
	Password string
	NickName string `gorm:"type:varchar(50)"`
	Status   string
	Avatar   string
	Money    string
}

// SetPassword
// @Description: 加密密码
// @receiver u *User
// @param password string
// @return error
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// CheckPassword
// @Description: 检测密码
// @receiver u *User
// @param pwd string
// @return bool
func (u *User) CheckPassword(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}

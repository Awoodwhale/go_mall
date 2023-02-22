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

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20) unique"`
	Email    string
	Password string
	NickName string
	Status   string
	Avatar   string
	Money    string
}

func (u *User) SetPassword(password string) error {
	/**
	 * SetPassword
	 * @Description: 加密密码
	 * @receiver u
	 * @param password
	 * @return error
	 */
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}

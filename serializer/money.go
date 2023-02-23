package serializer

import (
	"go_mall/model"
	"go_mall/pkg/utils"
)

type Money struct {
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	UserMoney string `json:"user_money"`
}

func BuildMoney(user *model.User, key string) *Money {
	utils.Encrypt.SetKey(key)
	money := utils.Encrypt.AesDecoding(user.Money)
	return &Money{
		UserID:    user.ID,
		UserName:  user.UserName,
		UserMoney: money,
	}
}

package serializer

import (
	"go_mall/model"
	"go_mall/pkg/utils"
)

// Money
// @Description: money vo
type Money struct {
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	UserMoney string `json:"user_money"`
}

// BuildMoney
// @Description: 构造一个money vo
// @param user *model.User
// @param key string
// @return *Money
func BuildMoney(user *model.User, key string) *Money {
	money := utils.AesDecoding(user.Money, key)
	return &Money{
		UserID:    user.ID,
		UserName:  user.UserName,
		UserMoney: money,
	}
}

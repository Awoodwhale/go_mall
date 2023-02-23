package serializer

import (
	"go_mall/conf"
	"go_mall/model"
)

// User
// @Description: vo view objective	传给前端的对象
type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Statue   string `json:"statue"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

// BuildUser
// @Description: 构造一个user vo
// @param user *model.User
// @return *User
func BuildUser(user *model.User) *User {
	return &User{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Statue:   user.Status,
		Avatar:   conf.ImgHost + ":" + conf.ImgPort + conf.AvatarPath + user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}

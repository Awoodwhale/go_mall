package e

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

var MessageFlags = map[int]string{
	Success:       "success",
	Error:         "failed",
	InvalidParams: "传入参数错误",

	ErrorWithKey:           "密钥长度错误",
	ErrorWithExistUser:     "用户已存在",
	ErrorWithNotExistUser:  "用户不存在",
	ErrorWithPassword:      "密码校验错误",
	ErrorWithGenToken:      "Token生成失败",
	ErrorWithParseToken:    "Token解析失败",
	ErrorWithExpiredToken:  "Token过期",
	ErrorWithEncryption:    "密码加密失败",
	ErrorWithUploadAvatar:  "头像上传失败",
	ErrorWithOperationType: "操作类型错误",
	ErrorWithCheckEmail:    "邮箱不一致",
	ErrorWithSendEmail:     "邮件发送失败",
	ErrorWithNotExistEmail: "邮箱不存在",
	ErrorWithSameEmail:     "绑定重复邮箱",

	ErrorWithSQL:      "SQL错误",
	ErrorWithFileSize: "文件大小错误,最大5MB",
}

// GetMessageByCode
// @Description: 获取code对应的message
// @param code int
// @return string
func GetMessageByCode(code int) string {
	if msg, ok := MessageFlags[code]; ok {
		return msg
	}
	return MessageFlags[Error]
}

// HandleBindingError
// @Description: 获取binding错误的msg tag
// @param err error
// @param obj any
// @return msg string
func HandleBindingError(err error, obj any) (msg string) {
	if err == nil {
		return
	}
	if errs, ok := err.(validator.ValidationErrors); ok {
		getObj := reflect.TypeOf(obj)
		for _, v := range errs {
			if f, exist := getObj.Elem().FieldByName(v.Field()); exist {
				msg = f.Tag.Get("msg")
				return
			}
		}
	}
	return
}

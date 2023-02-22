package e

var MessageFlags = map[int]string{
	Success:                   "success",
	Error:                     "failed",
	InvalidParams:             "传入参数错误",
	ErrorWithKey:              "密钥长度错误",
	ErrorWithExistUser:        "用户已存在",
	ErrorWithNotExistUser:     "用户不存在",
	ErrorWithPassword:         "密码校验错误",
	ErrorWithFailedToken:      "Token生成失败",
	ErrorWithFailedEncryption: "密码加密失败",
}

func GetMessageByCode(code int) string {
	if msg, ok := MessageFlags[code]; ok {
		return msg
	}
	return MessageFlags[Error]
}

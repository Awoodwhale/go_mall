package e

var MessageFlags = map[int]string{
	Success:                   "success",
	Error:                     "failed",
	InvalidParams:             "传入参数错误",
	ErrorWithExistUser:        "用户已经存在",
	ErrorWithFailedEncryption: "密码加密失败",
	ErrorWithKey:              "密钥长度错误",
}

func GetMessageByCode(code int) string {
	if msg, ok := MessageFlags[code]; ok {
		return msg
	}
	return MessageFlags[Error]
}

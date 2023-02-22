package e

var MessageFlags = map[int]string{
	Success:       "success",
	Error:         "failed",
	InvalidParams: "传入参数错误",

	ErrorWithKey:              "密钥长度错误",
	ErrorWithExistUser:        "用户已存在",
	ErrorWithNotExistUser:     "用户不存在",
	ErrorWithPassword:         "密码校验错误",
	ErrorWithFailedGenToken:   "Token生成失败",
	ErrorWithFailedParseToken: "Token解析失败",
	ErrorWithExpiredToken:     "Token过期",
	ErrorWithFailedEncryption: "密码加密失败",
	ErrorWithUploadAvatar:     "头像上传失败",

	ErrorWithSQL:      "SQL错误",
	ErrorWithFileSize: "文件大小错误,最大5MB",
}

func GetMessageByCode(code int) string {
	if msg, ok := MessageFlags[code]; ok {
		return msg
	}
	return MessageFlags[Error]
}

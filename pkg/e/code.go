package e

const (
	// 普遍错误
	Success       = 20000 // success
	Error         = 50000 // error
	InvalidParams = 40000 // 参数错误

	// user的错误
	ErrorWithExistUser        = 50001 // 用户存在
	ErrorWithNotExistUser     = 50002 // 用户不存在
	ErrorWithFailedEncryption = 50003 // 加密失败
	ErrorWithKey              = 50004 // 密钥key错误
	ErrorWithPassword         = 50005 // 密码错误
	ErrorWithFailedGenToken   = 50006 // token生成失败
	ErrorWithFailedParseToken = 50007 // token解析失败
	ErrorWithExpiredToken     = 50008 // token过期
	ErrorWithUploadAvatar     = 50009 // 头像上传失败

	// product的错误

	// SQL错误
	ErrorWithSQL      = 55555
	ErrorWithFileSize = 514514
)

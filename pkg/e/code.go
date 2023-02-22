package e

const (
	// 普遍错误
	Success       = 20000 // success
	Error         = 50000 // error
	InvalidParams = 40000 // 参数错误

	// user的错误
	ErrorWithExistUser     = 50001 // 用户存在
	ErrorWithNotExistUser  = 50002 // 用户不存在
	ErrorWithEncryption    = 50003 // 加密失败
	ErrorWithKey           = 50004 // 密钥key错误
	ErrorWithPassword      = 50005 // 密码错误
	ErrorWithGenToken      = 50006 // token生成失败
	ErrorWithParseToken    = 50007 // token解析失败
	ErrorWithExpiredToken  = 50008 // token过期
	ErrorWithUploadAvatar  = 50009 // 头像上传失败
	ErrorWithOperationType = 50010 // 操作类型错误
	ErrorWithCheckEmail    = 50011 // 邮箱不一致
	ErrorWithSendEmail     = 50012 // 邮箱发送错误
	ErrorWithNotExistEmail = 50013 // 邮箱不存在
	ErrorWithSameEmail     = 50014 // 绑定重复邮箱

	// product的错误

	// common错误
	ErrorWithSQL      = 504504 // SQL错误
	ErrorWithFileSize = 514514 // file大小错误
)

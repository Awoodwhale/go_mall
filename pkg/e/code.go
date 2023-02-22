package e

const (
	Success                   = 20000
	Error                     = 50000
	ErrorWithExistUser        = 50001 // 用户存在
	ErrorWithNotExistUser     = 50002 // 用户不存在
	ErrorWithFailedEncryption = 50003 // 加密失败
	ErrorWithKey              = 50004 // 密钥key错误
	ErrorWithPassword         = 50005 // 密码错误
	ErrorWithFailedToken      = 50006 // token生成失败
	InvalidParams             = 40000 // 参数错误
)

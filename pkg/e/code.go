package e

const (
	Success                   = 20000
	Error                     = 50000
	ErrorWithExistUser        = 50001 // 用户存在
	ErrorWithFailedEncryption = 50002
	ErrorWithKey              = 50003 // 密钥key错误
	InvalidParams             = 40000
)

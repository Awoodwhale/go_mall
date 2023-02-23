package serializer

// Response
// @Description: 响应
type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// TokenData
// @Description: token信息
type TokenData struct {
	User  any    `json:"user"`
	Token string `json:"token"`
}

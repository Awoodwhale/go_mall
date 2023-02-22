package serializer

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type TokenData struct {
	User  any    `json:"user"`
	Token string `json:"token"`
}

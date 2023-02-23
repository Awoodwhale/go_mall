package serializer

import "go_mall/pkg/e"

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

// ListData
// @Description: 列表数据
type ListData struct {
	Items any  `json:"items"`
	Total uint `json:"total"`
}

// BuildListResponse
// @Description: 返回列表数据resp
// @param item any
// @param total uint
// @return Response
func BuildListResponse(items any, total uint) Response {
	return Response{
		Code:    e.Success,
		Message: e.GetMessageByCode(e.Success),
		Data:    ListData{Items: items, Total: total},
	}
}

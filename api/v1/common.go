package v1

import (
	"go_mall/pkg/e"
	"go_mall/serializer"
)

// ErrorResponse
// @Description: 获取错误信息resp
// @param err error
// @param service any
// @return serializer.Response
func ErrorResponse(err error, service any) serializer.Response {
	return serializer.Response{
		Code:    e.Error,
		Message: e.GetMessageByCode(e.Error),
		Error:   e.HandleBindingError(err, service),
	}
}

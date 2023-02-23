package model

// BasePage
// @Description: 分页model
type BasePage struct {
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}

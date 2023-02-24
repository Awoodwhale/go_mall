package model

// BasePage
// @Description: 分页model
type BasePage struct {
	PageNum  uint `json:"page_num" form:"page_num"`
	PageSize uint `json:"page_size" form:"page_size"`
}

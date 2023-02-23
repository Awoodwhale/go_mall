package dao

import (
	"context"
	"go_mall/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

// GetNoticeById
// @Description: 通过id获取notice
// @receiver dao *NoticeDao
// @param uid uint
// @return notice *model.Notice
// @return err error
func (dao *NoticeDao) GetNoticeById(uid uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id=?", uid).First(&notice).Error
	return
}

// NewNoticeDao
// @Description: 通过ctx生成noticeDao
// @param ctx context.Context
// @return *NoticeDao
func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

// NewNoticeDaoByDB
// @Description: 通过db获取noticeDao
// @param db *gorm.DB
// @return *NoticeDao
func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

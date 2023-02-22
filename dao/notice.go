package dao

import (
	"context"
	"go_mall/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func (dao *NoticeDao) GetNoticeById(uid uint) (notice *model.Notice, err error) {
	/**
	 * GetNoticeById
	 * @Description: 通过id获取notice
	 * @receiver dao
	 * @param uid
	 * @return notice
	 * @return err
	 */
	err = dao.DB.Model(&model.Notice{}).Where("id=?", uid).First(&notice).Error
	return
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	/**
	 * NewNoticeDao
	 * @Description: 通过ctx生成dao
	 * @param ctx
	 * @return *NoticeDao
	 */
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	/**
	 * NewNoticeDaoByDB
	 * @Description: 通过db生成dao
	 * @param db
	 * @return *NoticeDao
	 */
	return &NoticeDao{db}
}

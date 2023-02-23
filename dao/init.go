package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var _db *gorm.DB

// InitDatabase
// @Description: 初始化database
// @param pathRead string
// @param pathWrite string
func InitDatabase(pathRead, pathWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       pathRead,
		DefaultStringSize:         256,  // string类型字段默认长度
		DisableDatetimePrecision:  true, // 禁止datetime精度
		DontSupportRenameIndex:    true, // 如果需要重命名索引，需要把索引删除后再重建
		DontSupportRenameColumn:   true, // 用change重命名列，mysql8之前的数据库不支持
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 空闲连接池
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	// 主从配置
	_db = db
	_ = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(pathWrite)},                      // 写操作
		Replicas: []gorm.Dialector{mysql.Open(pathRead), mysql.Open(pathRead)}, // 读操作
		Policy:   dbresolver.RandomPolicy{},
	}))

	// 数据迁移
	migration()
}

// NewDBClient
// @Description: 获取db对象
// @param ctx context.Context
// @return *gorm.DB
func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}

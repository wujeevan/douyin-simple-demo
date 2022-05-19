package repository

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Init() error {
	var err error
	dsn := "root:801925@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      newLogger,
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	db.DB()
	return err
}

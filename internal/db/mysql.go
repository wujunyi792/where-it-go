package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type MySQLCreator struct{}

func (m MySQLCreator) Create(ip string, port string, userName string, password string, dbName string) (*gorm.DB, error) {
	newLogger := logger2.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger2.Config{
			SlowThreshold:             time.Second,    // 慢 SQL 阈值
			LogLevel:                  logger2.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,           // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,          // 禁用彩色打印
		},
	)
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName, password, ip, port, dbName)
	//logger.Info.Printf("conn str: %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	return db, err
}

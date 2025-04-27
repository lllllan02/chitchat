package utils

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	var err error

	// 配置GORM日志
	gormLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢查询阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 彩色输出
		},
	)

	// 连接数据库
	DB, err = gorm.Open(mysql.Open(GetDSN()), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return err
	}

	// 配置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConn)
	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConn)
	// 设置连接最大生存时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}

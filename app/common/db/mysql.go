package db

import (
	"github.com/PittYao/gin_seed/app/common/global"
	"github.com/PittYao/gin_seed/app/common/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

//GormMySQL 连接mysql
func GormMySQL() *gorm.DB {
	if global.CONFIG.Mysql.Url == "" {
		global.LOG.Info("配置文件没有配置DB地址")
		return nil
	}

	// 配置zap为gorm的日志
	logger := zap.New(global.LOG)
	db, err := gorm.Open(mysql.Open(global.CONFIG.Mysql.Url), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		panic("MySQL启动异常: " + err.Error())
	}

	// config db coons
	sqlDB, err := db.DB()
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

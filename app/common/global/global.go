package global

import (
	"github.com/PittYao/gin_seed/app/etc"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// DB 数据库
	DB *gorm.DB

	// LOG 全局日志
	LOG *zap.Logger

	// CONFIG 全局系统配置
	CONFIG etc.Config
)

package global

import (
	"github.com/PittYao/gin_seed/app/etc"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	// DB 数据库
	DB *gorm.DB

	// REDIS 缓存数据库配置
	REDIS *redis.Client

	// LOG 全局日志
	LOG *zap.Logger

	// CONFIG 全局系统配置
	CONFIG etc.Config
)

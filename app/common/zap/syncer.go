package zap

import (
	"github.com/PittYao/gin_seed/app/common/global"
	"github.com/PittYao/gin_seed/app/common/globalkey"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"os"
)

// WriteSyncer 利用lumberjack库做日志分割
func WriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file,
		MaxSize:    global.CONFIG.Zap.MaxSize,    // 单文件最大
		MaxBackups: global.CONFIG.Zap.MaxBackups, //保留旧文件的最大个数
		MaxAge:     global.CONFIG.Zap.MaxAge,     // 保留旧文件的最大天数
		Compress:   global.CONFIG.Zap.Compress,
		LocalTime:  global.CONFIG.Zap.Localtime,
	}
	// 只在控制台输出日志
	if global.CONFIG.Zap.Format == globalkey.ZapConsole {
		return zapcore.AddSync(os.Stdout)
	}
	// 日志输出到控制台和文件
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
}

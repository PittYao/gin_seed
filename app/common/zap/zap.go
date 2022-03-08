package zap

import (
	"context"
	"errors"
	"fmt"
	"github.com/PittYao/gin_seed/app/common/global"
	"github.com/PittYao/gin_seed/app/common/globalkey"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

type Logger struct {
	ZapLogger                 *zap.Logger
	LogLevel                  gormlogger.LogLevel
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
	Colorful                  bool
}

func New(zapLogger *zap.Logger) Logger {
	return Logger{
		ZapLogger:                 zapLogger,
		LogLevel:                  gormlogger.Info,
		SlowThreshold:             2 * time.Second,
		IgnoreRecordNotFoundError: true,
		Colorful:                  false,
	}
}

//---  zap适配 gorm v2 实现gorm logger接口 --- //

func (l Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return Logger{
		ZapLogger:                 l.ZapLogger,
		LogLevel:                  level,
		SlowThreshold:             l.SlowThreshold,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
		Colorful:                  l.Colorful,
	}
}

func (l Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Info {
		return
	}
	l.ZapLogger.Sugar().Infof(str, args...)
}

func (l Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Warn {
		return
	}
	l.ZapLogger.Sugar().Warnf(str, args...)
}

func (l Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Error {
		return
	}
	l.ZapLogger.Sugar().Errorf(str, args...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	cost := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		l.ZapLogger.Error("sql_trace", zap.Error(err), zap.Duration("cost", cost), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.SlowThreshold != 0 && cost > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		l.ZapLogger.Warn("sql_trace", zap.Duration("cost", cost), zap.Int64("rows", rows), zap.String("sql", sql))
	case l.LogLevel >= gormlogger.Info:
		sql, rows := fc()
		l.ZapLogger.Info("sql_trace", zap.Duration("cost", cost), zap.Int64("rows", rows), zap.String("sql", sql))
	}

}

//---  适配 gorm v2 --- //

// Zap 初始化Logger
func Zap() (logger *zap.Logger) {

	// 调试级别
	debugLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorLevel := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/%s_debug.log", global.CONFIG.Zap.Dir, global.CONFIG.Name), debugLevel),
		getEncoderCore(fmt.Sprintf("./%s/%s_info.log", global.CONFIG.Zap.Dir, global.CONFIG.Name), infoLevel),
		getEncoderCore(fmt.Sprintf("./%s/%s_warn.log", global.CONFIG.Zap.Dir, global.CONFIG.Name), warnLevel),
		getEncoderCore(fmt.Sprintf("./%s/%s_error.log", global.CONFIG.Zap.Dir, global.CONFIG.Name), errorLevel),
	}

	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	// 对日志进行分割
	writer := WriteSyncer(fileName)
	return zapcore.NewCore(getEncoder(), writer, level)
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	// json格式日志输出到文件
	if global.CONFIG.Zap.Format == globalkey.ZapJson {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	// 日志输出到控制台
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

//getEncoderConfig 获取日志编码配置
func getEncoderConfig() (encoderConfig zapcore.EncoderConfig) {
	encoderConfig = zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return encoderConfig
}

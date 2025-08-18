package logc

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	"gorm.io/gorm/logger"
)

type CustomLogger struct {
	// 可以添加任何你需要的字段，例如日志输出的配置
}

func (l *CustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	// 返回一个新的 CustomLogger 实例（或返回本身），根据日志级别进行配置
	return &CustomLogger{}
}

func (l *CustomLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	// 实现 Info 日志记录
	logx.WithContext(ctx).Info("[INFO] %s: %v\n", msg, args)
}

func (l *CustomLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	// 实现 Warn 日志记录
	logx.WithContext(ctx).Info("[WARN] %s: %v\n", msg, args)
}

func (l *CustomLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	// 实现 Error 日志记录
	logx.WithContext(ctx).Error("[ERROR] %s: %v\n", msg, args)
}

func (l *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	// 实现 Trace 日志记录
	sql, rowsAffected := fc()
	duration := time.Since(begin)
	if err != nil {
		logx.WithContext(ctx).Debug("[TRACE] SQL: %s, Duration: %v, Rows: %d, Error: %v\n", sql, duration, rowsAffected, err)
	} else {
		logx.WithContext(ctx).Debug("[TRACE] SQL: %s, Duration: %v, Rows: %d\n", sql, duration, rowsAffected)
	}
}

package logger

import (
	"context"
	"fmt"
	"gorm.io/gorm/logger"
	"time"
)

// 实现gorm logger接口方法
//type Interface interface {
//	LogMode(LogLevel) Interface
//	Info(context.Context, string, ...interface{})
//	Warn(context.Context, string, ...interface{})
//	Error(context.Context, string, ...interface{})
//	Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
//}

type DbLog struct {
	LogLevel logger.LogLevel
}

func New() *DbLog {
	return new(DbLog)
}

func (l *DbLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *DbLog) Info(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	Info(msg, data)
}

func (l *DbLog) Warn(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	Warn(msg, data)
}

func (l *DbLog) Error(_ context.Context, msg string, data ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	Error(msg, data)
}

func (l *DbLog) Trace(_ context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//这块的逻辑可以自己根据业务情况修改
	fmt.Println(l.LogLevel)
	elapsed := time.Since(begin)
	sql, rows := fc()
	Warn(fmt.Sprintf("Trace sql [Duration: %s]: %v  row： %v  err: %v", elapsed, sql, rows, err))
}

package log

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	gl "gorm.io/gorm/logger"
)

// 实现gorm的日志接口，用来和gorm的日志结合起来

type GormLogger struct {
	LogLevel                  gl.LogLevel // 日志等级
	IgnoreRecordNotFoundError bool
	SlowThreshold             time.Duration // 慢sql的标准
}

func (l GormLogger) LogMode(level gl.LogLevel) gl.Interface {
	return &GormLogger{
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
	}
}

func (l GormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gl.Info {
		return
	}
	Infof(ctx, str, args...)
}

func (l GormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gl.Warn {
		return
	}
	Warnf(ctx, str, args...)
}

func (l GormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gl.Error {
		return
	}
	Errorf(ctx, str, args...)
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gl.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		Errorf(ctx, "err=%s elapsed=%s rows=%d sql=%s", err.Error(), elapsed.String(), rows, sql)
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gl.Warn:
		sql, rows := fc()
		var e string
		if err != nil {
			e = err.Error()
		}
		Warnf(ctx, "err=%s elapsed=%s rows=%d sql=%s", e, elapsed.String(), rows, sql)
	case l.LogLevel >= gl.Info:
		sql, rows := fc()
		var e string
		if err != nil {
			e = err.Error()
		}
		Infof(ctx, "err=%s elapsed=%s rows=%d sql=%s", e, elapsed.String(), rows, sql)
	}
}

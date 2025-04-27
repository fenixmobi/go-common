package log

import (
	"context"

	"go.uber.org/zap"
)

// 注入上下文

func ctxKv(ctx context.Context) *zap.Logger {
	var fieldList []zap.Field
	for _, field := range ctxFields {
		if ctx.Value(field) != nil {
			fieldList = append(fieldList, zap.String(field, ctx.Value(field).(string)))
		}
	}

	return logger.With(fieldList...)
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	ctxKv(ctx).Debug(msg, fields...)
}

func Debugw(ctx context.Context, msg string, args ...interface{}) {
	ctxKv(ctx).Sugar().Debugw(msg, args...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	ctxKv(ctx).Info(msg, fields...)
}

func Infow(ctx context.Context, msg string, args ...interface{}) {
	ctxKv(ctx).Sugar().Infow(msg, args...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	ctxKv(ctx).Warn(msg, fields...)
}

func Warnw(ctx context.Context, msg string, args ...interface{}) {
	ctxKv(ctx).Sugar().Warnw(msg, args...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	ctxKv(ctx).Error(msg, fields...)
}

func Errorw(ctx context.Context, msg string, args ...interface{}) {
	ctxKv(ctx).Sugar().Errorw(msg, args...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	ctxKv(ctx).Fatal(msg, fields...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	ctxKv(ctx).Panic(msg, fields...)
}

func With(fields ...zap.Field) *zap.Logger {
	return logger.With(fields...)
}

func WithOptions(fields ...zap.Option) *zap.Logger {
	return logger.WithOptions(fields...)
}

func Sync() {
	_ = logger.Sync()
}

func Debugf(ctx context.Context, template string, args ...interface{}) {
	ctxKv(ctx).Sugar().Debugf(template, args...)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	ctxKv(ctx).Sugar().Infof(template, args...)
}

func WithInfof(ctx context.Context, field zap.Field, template string, args ...interface{}) {
	ctxKv(ctx).Sugar().With(field).Infof(template, args...)
}

func Warnf(ctx context.Context, template string, args ...interface{}) {
	ctxKv(ctx).Sugar().Warnf(template, args...)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	ctxKv(ctx).Sugar().Errorf(template, args...)
}

func Fatalf(ctx context.Context, template string, args ...interface{}) {
	ctxKv(ctx).Sugar().Fatalf(template, args...)
}

func Panicf(ctx context.Context, template string, args ...interface{}) {
	ctxKv(ctx).Sugar().Panicf(template, args...)
}

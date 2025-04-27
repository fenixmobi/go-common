package log

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Options 日志配置
type Options struct {
	// 注入的上下文中的key
	CtxFields []string // 在上下文中传递的key

	// 日志文件的配置
	Filename     string // 文件名称
	MaxCount     uint   // 文件清理前的最大的保存个数
	CallerEnable bool   // 是否开启堆栈

	LogLevel     zapcore.Level // 日志等级
	CloseConsole bool          // 控制台输出
}

// 定义全局日志对象
var logger *zap.Logger

var ctxFields []string

// InitLog 初始化日志日志对象
func InitLog(opt Options, fields ...zap.Field) {
	ctxFields = opt.CtxFields

	// 初始化日志格式配置
	cfg := zapcore.EncoderConfig{
		MessageKey:   "message",
		TimeKey:      "timestamp",
		LevelKey:     "level",
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format(TimeFormat))
		},
	}

	// 创建日志实例
	if opt.CloseConsole {
		logger = zap.New(
			zapcore.NewTee(
				syncFile(cfg, opt),
			),
		)
	} else {
		logger = zap.New(
			zapcore.NewTee(
				syncConsole(cfg, opt),
				syncFile(cfg, opt),
			),
		)
	}

	if opt.CallerEnable {
		logger = logger.WithOptions(
			zap.AddCaller(),
			zap.AddCallerSkip(1),
		)
	}

	if len(fields) > 0 {
		logger = logger.With(fields...)
	}
}

// 设置输出对象
// 文件
func syncFile(cfg zapcore.EncoderConfig, opt Options) zapcore.Core {
	writer, err := rotatelogs.New(
		opt.Filename+".%Y%m%d.log",
		rotatelogs.WithLinkName(opt.Filename),
		rotatelogs.WithRotationCount(opt.MaxCount),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zapcore.AddSync(writer),
		opt.LogLevel,
	)
}

// 控制台
func syncConsole(cfg zapcore.EncoderConfig, opt Options) zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg),
		zapcore.AddSync(zapcore.Lock(os.Stdout)),
		opt.LogLevel,
	)
}

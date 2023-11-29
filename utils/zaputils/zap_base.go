package zaputils

import (
	"io"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Hook struct {
	*lumberjack.Logger
	writers []io.Writer
	enab    zapcore.LevelEnabler
	encoder zapcore.EncoderConfig
}

type HookOption func(*Hook)

func WithHookMaxsize(size int) HookOption {
	return func(hook *Hook) {
		hook.Logger.MaxSize = size
	}
}

func WithHookMaxBackups(num int) HookOption {
	return func(hook *Hook) {
		hook.Logger.MaxBackups = num
	}
}

func WithHookMaxAge(age int) HookOption {
	return func(hook *Hook) {
		hook.Logger.MaxAge = age
	}
}

func WithHookJackLogger(logger *lumberjack.Logger) HookOption {
	return func(hook *Hook) {
		if logger != nil {
			hook.Logger = logger
		}
	}
}

func WithHookJackLoggerNil() HookOption {
	return func(hook *Hook) {
		hook.Logger = nil
	}
}

func WithHookWriters(writers []io.Writer) HookOption {
	return func(hook *Hook) {
		hook.writers = writers
	}
}

func WithHookLevelEnab(enab zapcore.LevelEnabler) HookOption {
	return func(hook *Hook) {
		hook.enab = enab
	}
}

func WithHookEncoder(encoder zapcore.EncoderConfig) HookOption {
	return func(hook *Hook) {
		hook.encoder = encoder
	}
}

func NewHook(logPath string, opts ...HookOption) *Hook {
	if logPath == "" {
		logPath = "log"
	}

	res := &Hook{
		Logger: &lumberjack.Logger{
			Filename:   filepath.Join(logPath, filepath.Base(logPath)+".log"), // 日志文件路径，默认 os.TempDir()
			MaxSize:    200,                                                   // 每个日志文件保存10M，默认 200M
			MaxBackups: 10,                                                    // 保留10个备份，默认不限
			MaxAge:     10,                                                    // 保留10天，默认不限
			Compress:   true,                                                  // 是否压缩，默认不压缩
			LocalTime:  true,
		},
		enab: zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
			return lev >= zap.DebugLevel
		}),
		encoder: zap.NewDevelopmentEncoderConfig(),
	}

	for _, opt := range opts {
		opt(res)
	}

	return res
}

func NewZapLogCore(logPath string, opts ...HookOption) (core zapcore.Core, writer io.Writer) {

	hook := NewHook(logPath, opts...)

	var writers []zapcore.WriteSyncer
	if hook.Logger != nil {
		writers = []zapcore.WriteSyncer{zapcore.AddSync(hook.Logger)}
	}

	for _, ws := range hook.writers {
		writers = append(writers, zapcore.AddSync(ws))
	}

	consoleEncoder := zapcore.NewConsoleEncoder(hook.encoder)

	writer = zapcore.NewMultiWriteSyncer(writers...)

	// 设置日志级别
	return zapcore.NewCore(
		consoleEncoder,
		zapcore.NewMultiWriteSyncer(writers...),
		hook.enab,
	), writer
}

func NewZapStdoutCore(opts ...HookOption) (core zapcore.Core, writer io.Writer) {
	opts = append(opts, WithHookWriters([]io.Writer{
		os.Stdout,
	}), WithHookJackLoggerNil())
	return NewZapLogCore("", opts...)
}

func NewZapLog(logPath string, opts ...HookOption) (logger *zap.Logger) {
	logger, _ = NewZapLogExt(logPath, opts...)
	return logger
}

func NewZapLogExt(logPath string, opts ...HookOption) (logger *zap.Logger, writer io.Writer) {
	core, writer := NewZapLogCore(logPath, opts...)

	return zap.New(core, zap.AddCaller()), writer
}

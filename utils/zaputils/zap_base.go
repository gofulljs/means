package zaputils

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type HookOption func(**lumberjack.Logger)

func WithHookMaxsize(size int) HookOption {
	return func(logger **lumberjack.Logger) {
		(*logger).MaxSize = size
	}
}

func WithHookMaxBackups(num int) HookOption {
	return func(logger **lumberjack.Logger) {
		(*logger).MaxBackups = num
	}
}

func WithHookMaxAge(age int) HookOption {
	return func(logger **lumberjack.Logger) {
		(*logger).MaxAge = age
	}
}

func WithHookNew(hook *lumberjack.Logger) HookOption {
	return func(logger **lumberjack.Logger) {
		if hook != nil {
			*logger = hook
		}
	}
}

func NewZapLogCore(logPath string, isStdout bool, enab zapcore.LevelEnabler, opts ...HookOption) (core zapcore.Core) {
	hook := &lumberjack.Logger{
		Filename:   filepath.Join(logPath, filepath.Base(logPath)+".log"), // 日志文件路径，默认 os.TempDir()
		MaxSize:    200,                                                   // 每个日志文件保存10M，默认 100M
		MaxBackups: 10,                                                    // 保留30个备份，默认不限
		MaxAge:     10,                                                    // 保留7天，默认不限
		Compress:   true,                                                  // 是否压缩，默认不压缩
		LocalTime:  true,
	}

	for _, opt := range opts {
		opt(&hook)
	}

	writes := []zapcore.WriteSyncer{zapcore.AddSync(hook)}

	if isStdout {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// 设置日志级别
	return zapcore.NewCore(
		consoleEncoder,
		zapcore.NewMultiWriteSyncer(writes...),
		enab,
	)
}

func NewZapLogCoreEx(logPath string, isStdout, isColor bool, hook *lumberjack.Logger, enab zapcore.LevelEnabler) (core zapcore.Core, writer io.Writer) {
	base := path.Base(logPath)

	if hook == nil {
		hook = &lumberjack.Logger{
			Filename:   logPath + fmt.Sprintf("/%v.log", base), // 日志文件路径，默认 os.TempDir()
			MaxSize:    100,                                    // 每个日志文件保存10M，默认 100M
			MaxBackups: 30,                                     // 保留30个备份，默认不限
			MaxAge:     0,                                      // 保留7天，默认不限
			Compress:   true,                                   // 是否压缩，默认不压缩
			LocalTime:  true,
		}
	}

	writes := []zapcore.WriteSyncer{zapcore.AddSync(hook)}

	if isStdout {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}

	consoleConfig := zap.NewDevelopmentEncoderConfig()

	if isColor {
		consoleConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)

	writers := zapcore.NewMultiWriteSyncer(writes...)

	// 设置日志级别
	return zapcore.NewCore(
		consoleEncoder,
		writers,
		enab,
	), writers
}

func NewZapLog(logPath string, isStdout bool, enab zapcore.LevelEnabler, opts ...HookOption) (logger *zap.Logger) {
	core := NewZapLogCore(logPath, isStdout, enab, opts...)

	return zap.New(core, zap.AddCaller())
}

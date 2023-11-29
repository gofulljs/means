package zaputils

import (
	"io"
	"os"
	"path"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewPrettyLog(logPath string, isDebug, isColor bool, opts ...HookOption) *zap.Logger {
	logger, _ := NewPrettyLogExt(logPath, isDebug, isColor, opts...)
	return logger
}

// isDebug 仅针对文件
// isColor 仅针对terminal
func NewPrettyLogExt(logPath string, isDebug, isColor bool, opts ...HookOption) (logger *zap.Logger, writer io.Writer) {
	var stdoutOpts []HookOption
	copy(stdoutOpts, opts)
	if isColor {
		encoder := zap.NewDevelopmentEncoderConfig()
		encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder
		stdoutOpts = append(stdoutOpts, WithHookEncoder(encoder))
	}

	stdoutCore, stdoutWriter := NewZapStdoutCore(stdoutOpts...)

	enab := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		if isDebug {
			return lev >= zap.DebugLevel
		}

		return lev >= zap.InfoLevel
	})

	opts = append(opts, WithHookLevelEnab(enab))

	fileCore, fileWriter := NewZapLogCore(logPath, opts...)
	cores := []zapcore.Core{stdoutCore, fileCore}

	return zap.New(zapcore.NewTee(cores...), zap.AddCaller()), io.MultiWriter(stdoutWriter, fileWriter)
}

// debug以上均打印，只有一个文件
func NewZapOneLog(logPath string, isDebug, isStdout, isColor bool, opts ...HookOption) *zap.Logger {

	enab := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		if isDebug {
			return lev >= zap.DebugLevel
		}

		return lev >= zap.InfoLevel
	})

	opts = append(opts, WithHookLevelEnab(enab))

	if isStdout {
		opts = append(opts, WithHookWriters([]io.Writer{
			os.Stdout,
		}))
	}

	if isColor {
		encoder := zap.NewDevelopmentEncoderConfig()
		encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder
		opts = append(opts, WithHookEncoder(encoder))
	}

	return NewZapLog(logPath, opts...)
}

// 分3级显示日志信息 Info Error Debug 每级不同文件目录，便于查看需要信息
func NewMutil3ZapLog(logPath string, isDebug, isStdout, isColor bool, jackLoggerAll, jackLoggerError, jackLoggerInfo *lumberjack.Logger) *zap.Logger {
	logger, _, _ := NetMutil3ZapLogExt(logPath, isDebug, isStdout, isColor, jackLoggerAll, jackLoggerError, jackLoggerInfo)

	return logger
}

// 分3级导出主 logger 和 core
func NetMutil3ZapLogExt(logPath string, isDebug, isStdout, isColor bool, jackLoggerAll, jackLoggerError, jackLoggerInfo *lumberjack.Logger) (logger *zap.Logger, allWriter, errorWirter io.Writer) {
	// 级别
	pAll := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		if isDebug {
			return lev >= zap.DebugLevel
		}

		return lev >= zap.InfoLevel
	})

	// 记录大于错误等级的所有数据
	pErrors := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	// 等于Info的数据使用
	pInfo := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})

	allOpts := []HookOption{
		WithHookJackLogger(jackLoggerAll), WithHookLevelEnab(pAll),
	}
	if isStdout {
		allOpts = append(allOpts, WithHookWriters([]io.Writer{
			os.Stdout,
		}))
	}
	if isColor {
		encoder := zap.NewDevelopmentEncoderConfig()
		encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder
		allOpts = append(allOpts, WithHookEncoder(encoder))
	}

	coreAll, allWriter := NewZapLogCore(path.Join(logPath, "all"), allOpts...)
	coreError, errorWriter := NewZapLogCore(path.Join(logPath, "error"), WithHookJackLogger(jackLoggerError), WithHookLevelEnab(pErrors))
	coreInfo, _ := NewZapLogCore(path.Join(logPath, "info"), WithHookJackLogger(jackLoggerInfo), WithHookLevelEnab(pInfo))

	cores := []zapcore.Core{coreAll, coreError, coreInfo}

	return zap.New(zapcore.NewTee(cores...), zap.AddCaller()), allWriter, io.MultiWriter(allWriter, errorWriter)
}

package zaputils

import (
	"io"
	"path"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// debug以上均打印，只有一个文件
func NewZapOneLog(logPath string, isDebug, isStdout bool, opts ...HookOption) *zap.Logger {

	enab := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		if isDebug {
			return lev >= zap.DebugLevel
		}

		return lev >= zap.InfoLevel
	})

	return NewZapLog(logPath, isStdout, enab, opts...)
}

// 分3级显示日志信息 Info Error Debug 每级不同文件目录，便于查看需要信息
func NewMutil3ZapLog(logPath string, isDebug, isStdout bool, hookAll, hookError, hookInfo *lumberjack.Logger) *zap.Logger {
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

	coreAll := NewZapLogCore(path.Join(logPath, "all"), isStdout, pAll, WithHookNew(hookAll))
	coreError := NewZapLogCore(path.Join(logPath, "error"), false, pErrors, WithHookNew(hookError))
	coreInfo := NewZapLogCore(path.Join(logPath, "info"), false, pInfo, WithHookNew(hookInfo))

	cores := []zapcore.Core{coreAll, coreError, coreInfo}

	return zap.New(zapcore.NewTee(cores...), zap.AddCaller())
}

// 分3级导出主 logger 和 core
func NetMutil3ZapExLog(logPath string, isDebug, isStdout, isColor bool, hookAll, hookError, hookInfo *lumberjack.Logger) (logger *zap.Logger, allWriter, errorWirter io.Writer) {
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

	coreAll, allWriter := NewZapLogCoreEx(path.Join(logPath, "all"), isStdout, isColor, hookAll, pAll)
	coreError, errorWriter := NewZapLogCoreEx(path.Join(logPath, "error"), false, false, hookError, pErrors)
	coreInfo := NewZapLogCore(path.Join(logPath, "info"), false, pInfo, WithHookNew(hookInfo))

	cores := []zapcore.Core{coreAll, coreError, coreInfo}

	return zap.New(zapcore.NewTee(cores...), zap.AddCaller()), allWriter, errorWriter
}

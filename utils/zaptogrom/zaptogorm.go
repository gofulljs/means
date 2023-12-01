package zaptogrom

import (
	"context"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	gormlogger "gorm.io/gorm/logger"
)

type Logger struct {
	ZapLogger        *zap.Logger
	LogLevel         gormlogger.LogLevel
	SlowThreshold    time.Duration
	SkipCallerLookup bool
}

func New(zapLogger *zap.Logger) *Logger {
	return &Logger{
		ZapLogger:        zapLogger,
		LogLevel:         gormlogger.Warn,
		SlowThreshold:    100 * time.Millisecond,
		SkipCallerLookup: false,
	}
}

func (l *Logger) SetAsDefault() {
	gormlogger.Default = l
}

func (l *Logger) SetDebug(debug bool) {
	if debug {
		l.LogLevel = gormlogger.Info
	} else {
		l.LogLevel = gormlogger.Warn
	}
}

func (l Logger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return &Logger{
		ZapLogger:        l.ZapLogger,
		SlowThreshold:    l.SlowThreshold,
		LogLevel:         level,
		SkipCallerLookup: l.SkipCallerLookup,
	}
}

func (l *Logger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Info {
		return
	}
	l.logger().Sugar().Debugf(str, args...)
}

func (l *Logger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Warn {
		return
	}
	l.logger().Sugar().Warnf(str, args...)
}

func (l *Logger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Error {
		return
	}
	l.logger().Sugar().Errorf(str, args...)
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormlogger.Error:
		sql, rows := fc()
		l.logger().Sugar().Errorf("%v err:%v %v %vrows sql:%v", "trace", err, elapsed, rows, sql)
	case l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn:
		sql, rows := fc()
		l.logger().Sugar().Warnf("%v %v %vrows sql:%v", "trace", elapsed, rows, sql)
	case l.LogLevel >= gormlogger.Info:
		sql, rows := fc()
		l.logger().Sugar().Debugf("%v %v %vrows sql:%v", "trace", elapsed, rows, sql)
	}
}

var (
	gormPackage = "gorm.io/gorm"
)

func (l *Logger) logger() *zap.Logger {
	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.Contains(file, gormPackage):
		//case strings.Contains(file, zapgormPackage):
		default:
			return l.ZapLogger.WithOptions(zap.AddCallerSkip(i - 1))
		}
	}
	return l.ZapLogger
}

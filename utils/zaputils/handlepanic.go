package zaputils

import (
	"fmt"
	"runtime"

	"go.uber.org/zap"
)

func HandlePanic(logger *zap.Logger, prefix string, extra ...interface{}) {
	if err := recover(); err != nil {
		stackBuf := make([]byte, 4096)
		stackSize := runtime.Stack(stackBuf, false)
		stackTrace := string(stackBuf[:stackSize])

		logger.Error(prefix,
			zap.String("error", fmt.Sprintf("%v", err)),
			zap.String("stack_trace", stackTrace),
			zap.Any("any", func() interface{} {
				if extra != nil {
					return extra
				}
				return zap.Skip()
			}()),
		)
	}
}

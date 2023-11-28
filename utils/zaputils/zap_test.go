package zaputils

import "testing"

func Test_onelog(t *testing.T) {
	fFunc := func(logPath string, isDebug, isStdout bool) (f func(t *testing.T)) {
		return func(t *testing.T) {
			logger := NewZapOneLog(logPath, isDebug, isStdout)
			logger.Debug("hi Debug")
			logger.Error("hi Error")
			logger.Info("hi Info")
		}
	}

	t.Run("debug stdout", fFunc("log/oneall", true, true))

	t.Run("info stdout", fFunc("log/oneinfo", false, true))

	t.Run("debug", fFunc("log/onefile", true, false))
}

func Test_3log(t *testing.T) {

	fFunc := func(logPath string, isDebug, isStdout bool) (f func(t *testing.T)) {
		return func(t *testing.T) {
			logger := NewMutil3ZapLog(logPath, isDebug, isStdout, nil, nil, nil)
			logger.Debug("hi Debug")
			logger.Error("hi Error")
			logger.Info("hi Info")
			// logger.Panic("hi Panic") // Panic 底层调用 panic
		}
	}

	t.Run("debug stdout", fFunc("log/3all", true, true))

	t.Run("info stdout", fFunc("log/3info", false, true))

	t.Run("debug", fFunc("log/3file", true, false))
}

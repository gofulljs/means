package zaputils

import "testing"

func Test_onelog(t *testing.T) {
	fFunc := func(logPath string, isDebug, isStdout, isColor bool) (f func(t *testing.T)) {
		return func(t *testing.T) {
			logger := NewZapOneLog(logPath, isDebug, isStdout, isColor)
			logger.Debug("hi Debug")
			logger.Error("hi Error")
			logger.Info("hi Info")
		}
	}

	t.Run("debug stdout", fFunc("log/oneall", true, true, false))

	t.Run("info stdout", fFunc("log/oneinfo", false, true, false))

	t.Run("debug", fFunc("log/onefile", true, false, false))

	t.Run("debug stdout and color", fFunc("log/oneall_color", true, true, true))
}

func Test_3log(t *testing.T) {
	fFunc := func(logPath string, isDebug, isStdout, isColor bool) (f func(t *testing.T)) {
		return func(t *testing.T) {
			logger := NewMutil3ZapLog(logPath, isDebug, isStdout, isColor, nil, nil, nil)
			logger.Debug("hi Debug")
			logger.Error("hi Error")
			logger.Info("hi Info")
			// logger.Panic("hi Panic") // Panic 底层调用 panic
		}
	}

	t.Run("debug stdout", fFunc("log/3all", true, true, false))

	t.Run("info stdout", fFunc("log/3info", false, true, false))

	t.Run("debug", fFunc("log/3file", true, false, false))

	t.Run("debug stdout and color", fFunc("log/3all_color", true, true, true))

}

func Test_PrettyLog(t *testing.T) {
	fFunc := func(logPath string, isDebug, isColor bool) (f func(t *testing.T)) {
		return func(t *testing.T) {
			logger := NewPrettyLog(logPath, isDebug, isColor)
			logger.Debug("hi Debug")
			logger.Error("hi Error")
			logger.Info("hi Info")
			// logger.Panic("hi Panic") // Panic 底层调用 panic
		}
	}

	t.Run("debug stdout", fFunc("log/pretty_debug", true, false))

	t.Run("info stdout", fFunc("log/pretty_info", false, false))

	t.Run("debug", fFunc("log/pretty", true, false))

	t.Run("debug stdout and color", fFunc("log/pretty_color", true, true))

	t.Run("debug stdout and color info", fFunc("log/pretty_color_info", false, true))

}

package logger

type Logger interface {
	Printf(string, ...any)
	Print(...any)
}

var _log Logger

func Log(log ...Logger) Logger {
	if log == nil || len(log) == 0 {
		return _log
	}
	_log = log[0]
	return _log
}

func Printf(format string, args ...any) {
	_log.Printf(format, args...)
}

func Print(args ...any) {
	_log.Print(args...)
}

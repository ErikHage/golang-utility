package logging

type LogLevel int64

const (
	Trace LogLevel = 0
	Debug LogLevel = 1
	Info  LogLevel = 2
	Warn  LogLevel = 3
	Error LogLevel = 4
	Fatal LogLevel = 5
)

func (l LogLevel) String() string {
	switch l {
	case Trace:
		return "TRACE"
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	}
	return "unknown"
}

func LogLevelFromString(s string) LogLevel {
	switch s {
	case "TRACE":
		return Trace
	case "DEBUG":
		return Debug
	case "INFO":
		return Info
	case "WARN":
		return Warn
	case "ERROR":
		return Error
	case "FATAL":
		return Fatal
	}
	return Info
}

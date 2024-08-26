package v1

import "os"

func GetLogger(name string) Logger {
	logLevelVar := os.Getenv("LOG_LEVEL")
	logLevel := LogLevelFromString(logLevelVar)

	return &LoggerImpl{
		Name:     name,
		LogLevel: logLevel,
	}
}

package logging

import (
	"encoding/json"
	"fmt"
	"golang-utility/pkg/errors"
	"log"
	"net/http"
)

type Logger interface {
	Trace(message string)
	Debug(message string)
	Info(message string)
	Warn(message string, error errors.Error)
	Error(message string, error errors.Error)
	Fatal(message string, error errors.Error)
	LogIncomingRequest(req *http.Request)
}

type LoggerImpl struct {
	Name     string
	LogLevel LogLevel
}

func (logger *LoggerImpl) LogIncomingRequest(req *http.Request) {
	headersToLog := make(map[string]string)

	headersToLog["Content-Type"] = req.Header.Get("Content-Type")
	headersToLog["User-Agent"] = req.Header.Get("User-Agent")

	reqHeadersBytes, err := json.Marshal(headersToLog)

	if err != nil {
		logger.Info("Unable to parse request headers")
		return
	}

	message := fmt.Sprintf("[%s: %s] headers: %s", req.Method, req.URL.Path, string(reqHeadersBytes))

	logger.Info(message)
}

func (logger *LoggerImpl) Trace(message string) {
	if logger.LogLevel <= Trace {
		logger.log(Trace.String(), message, nil)
	}
}

func (logger *LoggerImpl) Debug(message string) {
	if logger.LogLevel <= Debug {
		logger.log(Debug.String(), message, nil)
	}
}

func (logger *LoggerImpl) Info(message string) {
	if logger.LogLevel <= Info {
		logger.log(Info.String(), message, nil)
	}
}

func (logger *LoggerImpl) Warn(message string, error errors.Error) {
	if logger.LogLevel <= Warn {
		logger.log(Warn.String(), message, error)
	}
}

func (logger *LoggerImpl) Error(message string, error errors.Error) {
	if logger.LogLevel <= Error {
		logger.log(Error.String(), message, error)
	}
}

func (logger *LoggerImpl) Fatal(message string, error errors.Error) {
	if logger.LogLevel <= Fatal {
		logger.log(Fatal.String(), message, error)
	}
}

func (logger *LoggerImpl) log(levelString, message string, error errors.Error) {
	logMessage := logger.formatLogMessage(levelString, message)
	log.Println(logMessage)

	if error != nil {
		errorLog := logger.formatLogMessage(levelString, errors.ToJsonString(error))
		log.Println(errorLog)

		if error.GetStackTrace() != nil {
			log.Println(*error.GetStackTrace())
		}
	}
}

func (logger *LoggerImpl) formatLogMessage(levelString string, message string) string {
	return fmt.Sprintf("[%s][%s] %s", levelString, logger.Name, message)
}

package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"golang-utility/pkg/errors"
)

type Logger interface {
	Debug(message string)
	Info(message string)
	Notice(message string)
	Warn(message string, error errors.Error)
	Error(message string, error errors.Error)
	Critical(message string, error errors.Error)
	Alert(message string, error errors.Error)
	Emergency(message string, error errors.Error)
	LogIncomingRequest(req *http.Request)
}

type LoggerImpl struct {
	logger *slog.Logger
}

func (logger *LoggerImpl) LogIncomingRequest(req *http.Request) {
	headersToLog := make(map[string]string)

	headersToLog["Content-Type"] = req.Header.Get("Content-Type")
	headersToLog["User-Agent"] = req.Header.Get("User-Agent")
	headersToLog["X-Amos-Client-Version"] = req.Header.Get("X-Amos-Client-Version")
	headersToLog["X-Amos-Device-Platform"] = req.Header.Get("X-Amos-Device-Platform")
	headersToLog["X-Amos-Device-Platform-Version"] = req.Header.Get("X-Amos-Device-Platform-Version")

	reqHeadersBytes, err := json.Marshal(headersToLog)

	if err != nil {
		logger.Info("Unable to parse request headers")
		return
	}

	message := fmt.Sprintf("[%s: %s] headers: %s", req.Method, req.URL.Path, string(reqHeadersBytes))

	logger.Info(message)
}

func (l *LoggerImpl) Debug(message string) {
	l.logger.Debug(message)
}
func (l *LoggerImpl) Info(message string) {
	l.logger.Info(message)
}
func (l *LoggerImpl) Notice(message string) {
	l.logger.Log(context.TODO(), LevelNotice, message)
}
func (l *LoggerImpl) Warn(message string, error errors.Error) {
	l.logger.Warn(message, "error", error)
}
func (l *LoggerImpl) Error(message string, error errors.Error) {
	l.logger.Error(message, "error", error)
}
func (l *LoggerImpl) Critical(message string, error errors.Error) {
	l.logger.Log(context.TODO(), LevelCritical, message, "error", error)
}

func (l *LoggerImpl) Alert(message string, error errors.Error) {
	l.logger.Log(context.TODO(), LevelAlert, message, "error", error)
}

func (l *LoggerImpl) Emergency(message string, error errors.Error) {
	l.logger.Log(context.TODO(), LevelEmergency, message, "error", error)
}

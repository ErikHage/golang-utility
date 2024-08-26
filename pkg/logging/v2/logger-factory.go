package logging

import (
	"context"
	"log/slog"
	"os"
	"runtime"
)

// Only exists to be able to have meaningful sourcelocations in the logs with a wrapped logger
type SourceHandler struct {
	slog.Handler
}

// look an extra stack frame up to avoid getting the logger file for everything
func (h *SourceHandler) Handle(ctx context.Context, r slog.Record) error {
	var pcs [1]uintptr
	// skip [runtime.Callers, this function, this function's caller, etc.]
	runtime.Callers(5, pcs[:])
	pc := pcs[0]
	r.PC = pc
	h.Handler.Handle(ctx, r)
	return nil
}

// Avoids overwriting with the base handler
func (h *SourceHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &SourceHandler{
		Handler: h.Handler.WithAttrs(attrs),
	}
}

// Avoids overwriting with the base handler
func (h *SourceHandler) WithGroup(name string) slog.Handler {
	return &SourceHandler{
		Handler: h.Handler.WithGroup(name),
	}
}

func replacer(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.MessageKey {
		a.Key = "message"
	} else if a.Key == slog.SourceKey {
		a.Key = "logging.googleapis.com/sourceLocation"
	} else if a.Key == slog.LevelKey {
		a.Key = "severity"
		level := a.Value.Any().(slog.Level)
		if severity, ok := LevelToSeverity[level]; ok {
			a.Value = slog.StringValue(severity)
		}
	}
	return a
}

func GetLogger(name string) Logger {
	logLevelVar := os.Getenv("LOG_LEVEL")
	logLevel := LogLevelFromString[logLevelVar]
	jsonHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource:   true,
		Level:       logLevel,
		ReplaceAttr: replacer,
	})

	sourceHandler := &SourceHandler{
		jsonHandler,
	}

	logger := slog.New(sourceHandler).With(slog.String("module", name))

	return &LoggerImpl{
		logger,
	}
}

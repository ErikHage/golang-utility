package logging

import "log/slog"

// Add GCP levels: https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#logseverity
const (
	LevelNotice    = slog.Level(2)
	LevelCritical  = slog.Level(12)
	LevelAlert     = slog.Level(14)
	LevelEmergency = slog.Level(16)
)

var LevelToSeverity = map[slog.Level]string{
	LevelNotice:    "NOTICE",
	LevelCritical:  "CRITICAL",
	LevelAlert:     "ALERT",
	LevelEmergency: "EMERGENCY",
}

var LogLevelFromString = map[string]slog.Level{
	"DEBUG":     slog.LevelDebug,
	"INFO":      slog.LevelInfo,
	"NOTICE":    LevelNotice,
	"WARN":      slog.LevelWarn,
	"ERROR":     slog.LevelError,
	"CRITICAL":  LevelCritical,
	"ALERT":     LevelAlert,
	"EMERGENCY": LevelEmergency,
}

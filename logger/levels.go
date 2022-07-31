package logger

import "github.com/rs/zerolog"

// Level is type for logging level constants.
type Level string

// List of available logging levels.
const (
	LevelDebug   Level = "debug"
	LevelInfo    Level = "info"
	LevelWarning Level = "warn"
	LevelError   Level = "error"
	LevelFatal   Level = "fatal"
)

// getZerologLevel convert level to zerolog format.
func getZerologLevel(level Level) zerolog.Level {
	switch level {
	case LevelDebug:
		return zerolog.DebugLevel
	case LevelInfo:
		return zerolog.InfoLevel
	case LevelWarning:
		return zerolog.WarnLevel
	case LevelError:
		return zerolog.ErrorLevel
	case LevelFatal:
		return zerolog.FatalLevel
	default:
		return zerolog.NoLevel
	}
}

package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type (
	// Logger struct.
	Logger struct {
		*zerolog.Logger
	}

	// loggerOptions is logger initial parameters.
	loggerOptions struct {
		format  Format
		level   Level
		outputs []io.Writer
	}
)

// New create logger instance.
func New(options ...loggerOption) *Logger {

	loggerOptions := loggerOptions{
		format:  FormatJSON,
		level:   LevelInfo,
		outputs: []io.Writer{os.Stderr},
	}
	for _, opt := range options {
		opt(&loggerOptions)
	}

	lw := zerolog.MultiLevelWriter(loggerOptions.outputs...)
	zl := zerolog.New(lw).Level(getZerologLevel(loggerOptions.level))

	return &Logger{
		Logger: &zl,
	}
}

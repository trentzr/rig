package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type (
	// Logger struct.
	Logger struct {
		*zerolog.Logger
	}
)

// New create logger instance.
func New(opts ...option) *Logger {

	lo := options{
		format:    FormatJSON,
		level:     LevelInfo,
		outputs:   []io.Writer{os.Stderr},
		tsEnabled: true,
		tsFormat:  time.RFC3339,
		tsName:    "ts",
	}
	for _, opt := range opts {
		opt(&lo)
	}

	if lo.format == FormatText || lo.format == FormatConsole {
		for i, o := range lo.outputs {
			switch o.(type) {
			case *os.File:
				lo.outputs[i] = &zerolog.ConsoleWriter{
					Out:        o,
					NoColor:    (lo.format != FormatConsole),
					TimeFormat: lo.tsFormat,
				}
			default:
				continue
			}
		}
	}

	lw := zerolog.MultiLevelWriter(lo.outputs...)
	zl := zerolog.New(lw).Level(getZerologLevel(lo.level))

	if lo.tsEnabled && lo.tsName != "" && lo.tsFormat != "" {
		zerolog.TimestampFieldName = lo.tsName
		zerolog.TimeFieldFormat = lo.tsFormat
		zl = zl.With().Timestamp().Logger()
	}

	return &Logger{
		Logger: &zl,
	}
}

// Debug implements Debug method for logger.
func (l *Logger) Debug(msg string) {
	l.Logger.Debug().Msg(msg)
}

// Info implements Info method for logger.
func (l *Logger) Info(msg string) {
	l.Logger.Info().Msg(msg)
}

// Warn implements Warn method for logger.
func (l *Logger) Warn(msg string) {
	l.Logger.Warn().Msg(msg)
}

// Error implements Error method for logger.
func (l *Logger) Error(msg string) {
	l.Logger.Error().Msg(msg)
}

// Fatal implements Fatal method for logger.
func (l *Logger) Fatal(msg string) {
	l.Logger.Fatal().Msg(msg)
}

// Debugf implements Debugf method for logger.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Logger.Debug().Msgf(format, args...)
}

// Infof implements Infof method for logger.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Logger.Info().Msgf(format, args...)
}

// Warnf implements Warnf method for logger.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Logger.Warn().Msgf(format, args...)
}

// Errorf implements Errorf method for logger.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Logger.Error().Msgf(format, args...)
}

// Fatalf implements Fatalf method for logger.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatal().Msgf(format, args...)
}

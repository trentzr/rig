package notifier

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/trentzr/rig/config/env"
	"github.com/trentzr/rig/notifier"
)

type (
	// Sentry notifier
	Sentry struct {
		hub        *sentry.Hub
		stacktrace bool
		timeout    time.Duration
		zlvl       zerolog.Level
	}
)

// NewNotifier create new sentry writer instance.
func NewNotifier(dsn string, opts ...option) (*Sentry, error) {

	if dsn == "" {
		return nil, ErrDsnIsEmpty
	}

	so := options{
		env:     env.GetString(env.FieldNameEnvironment, ""),
		distro:  env.GetString(env.FieldNameDistro, ""),
		level:   notifier.LevelError,
		timeout: 10 * time.Second,
		output:  os.Stderr,
		release: env.GetString(env.FieldNameRelease, ""),
	}
	for _, opt := range opts {
		opt(&so)
	}

	client, err := sentry.NewClient(sentry.ClientOptions{
		Dsn:              dsn,
		Debug:            so.level == notifier.LevelDebug,
		AttachStacktrace: so.stacktrace,
		DebugWriter:      so.output,
		Environment:      so.env,
		Dist:             so.distro,
		Release:          so.release,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to init sentry client: %w", err)
	}

	zlvl, err := zerolog.ParseLevel(string(so.level))
	if err != nil {
		zlvl = zerolog.ErrorLevel
	}

	return &Sentry{
		hub:        sentry.NewHub(client, sentry.NewScope()),
		stacktrace: so.stacktrace,
		timeout:    so.timeout,
		zlvl:       zlvl,
	}, nil
}

// Write implements io.Writer method.
func (s *Sentry) Write(data []byte) (int, error) {
	return s.WriteLevel(s.zlvl, data)
}

// WriteLevel is implements zerolog.LevelWriter method.
func (s *Sentry) WriteLevel(zlvl zerolog.Level, data []byte) (int, error) {

	n := len(data)
	if zlvl < s.zlvl {
		return n, nil
	}

	var extra Fields
	var err = json.Unmarshal(data, &extra)
	if err != nil {
		return 0, fmt.Errorf(" failed to unmarshal data: %w", err)
	}

	msgField, msg := s.getErrorMessage(extra)

	delete(extra, msgField)
	delete(extra, zerolog.MessageFieldName)
	delete(extra, zerolog.ErrorFieldName)
	delete(extra, zerolog.LevelFieldName)
	delete(extra, zerolog.TimestampFieldName)

	s.capture(zlvl, msg, extra)

	return n, nil
}

// capture create new sentry event.
func (s *Sentry) capture(level zerolog.Level, msg string, extra Fields) {

	var (
		e = sentry.NewEvent()
		t *sentry.Stacktrace
	)

	if s.stacktrace {
		t = s.getStacktrace(msg)
	}

	e.Message = msg
	e.Level = getSentryLevel(level)
	e.Timestamp = time.Now()
	e.Exception = []sentry.Exception{{
		Value:      msg,
		Stacktrace: t,
	}}

	_ = s.hub.CaptureEvent(e)
}

// Fields is map for incomming data fields.
type Fields map[string]interface{}

// getErrorMessage extract error message from fields.
func (s *Sentry) getErrorMessage(fields Fields) (fieldName, fieldData string) {

	// Possible err message fields.
	const (
		logFieldNameMessage = "message"
		logFieldNameErr     = "err"
		logFieldNameError   = "error"
	)

	var ok bool
	if fieldData, ok = fields[logFieldNameErr].(string); ok {
		return logFieldNameErr, fieldData
	}

	if fieldData, ok = fields[logFieldNameError].(string); ok {
		return logFieldNameError, fieldData
	}

	fieldData, _ = fields[logFieldNameMessage].(string)
	return logFieldNameMessage, fieldData
}

// getSentryLevel convert level to sentry format.
func getSentryLevel(level zerolog.Level) sentry.Level {
	switch level {
	case zerolog.DebugLevel:
		return sentry.LevelDebug
	case zerolog.InfoLevel:
		return sentry.LevelInfo
	case zerolog.WarnLevel:
		return sentry.LevelWarning
	case zerolog.ErrorLevel:
		return sentry.LevelError
	case zerolog.FatalLevel:
		return sentry.LevelFatal
	default:
		return sentry.LevelInfo
	}
}

// getStacktrace extract stacktrace from error.
func (s *Sentry) getStacktrace(msg string) (stacktrace *sentry.Stacktrace) {

	if stacktrace = sentry.ExtractStacktrace(errors.New(msg)); stacktrace == nil {
		return nil
	}

	var frames = make([]sentry.Frame, 0, len(stacktrace.Frames))
	for _, frame := range stacktrace.Frames {
		// Skip tracing into logger files.
		if strings.HasPrefix(frame.Module, "github.com/rs/zerolog") ||
			strings.HasSuffix(frame.Filename, "logger.go") ||
			strings.HasSuffix(frame.Filename, "sentry.go") {
			continue
		}

		frames = append(frames, frame)
	}

	stacktrace.Frames = frames

	return stacktrace
}

// Close method implements io.Closer
func (s *Sentry) Close() error {
	if !s.hub.Flush(s.timeout) {
		return ErrFlushTimeout
	}
	return nil
}

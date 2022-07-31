package notifier

import "errors"

// List of errors.
var (
	ErrDsnIsEmpty   = errors.New("sentry dsn is empty")
	ErrFlushTimeout = errors.New("sentry flush timeout")
)

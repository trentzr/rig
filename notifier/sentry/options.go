package notifier

import (
	"io"
	"time"

	"github.com/trentzr/rig/notifier"
)

type (
	// option func.
	option func(*options)

	// options is sentry initial params.
	options struct {
		distro     string
		env        string
		level      notifier.Level
		output     io.Writer
		release    string
		stacktrace bool
		timeout    time.Duration
	}
)

// WithDistro setup distro name.
func WithDistro(distro string) option {
	return func(so *options) {
		so.distro = distro
	}
}

// WithEnvironment setup env name.
func WithEnvironment(env string) option {
	return func(so *options) {
		so.env = env
	}
}

// WithLevel setup logging level (msg with level above it not been logged).
// Default is error
func WithLevel(level notifier.Level) option {
	return func(so *options) {
		so.level = level
	}
}

// WithDebugOutput setup sentry debug output destination.
// Default is os.Stderr.
func WithDebugOutput(output io.Writer) option {
	return func(so *options) {
		so.output = output
	}
}

// WithRelease setup release name.
func WithRelease(release string) option {
	return func(so *options) {
		so.release = release
	}
}

// WithStacktraceEnabled enable stacktrace extraction.
// Default is disabled.
func WithStacktraceEnabled(v bool) option {
	return func(so *options) {
		so.stacktrace = v
	}
}

// WithTimeout setup custom capture flush timeout.
// Default is 10 sec.
func WithTimeout(timeout time.Duration) option {
	return func(so *options) {
		so.timeout = timeout
	}
}

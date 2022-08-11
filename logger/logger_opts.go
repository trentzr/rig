package logger

import "io"

type (
	// option is logger modification func type.
	option func(*options)

	// options is logger initial parameters.
	options struct {
		format    Format
		level     Level
		outputs   []io.Writer
		tsEnabled bool
		tsFormat  string
		tsName    string
	}
)

// WithLevel setup severnity level for logger.
func WithLevel(level Level) option {
	return func(lo *options) {
		lo.level = level
	}
}

// WithFormat setup logging format.
func WithFormat(format Format) option {
	return func(lo *options) {
		lo.format = format
	}
}

// WithOutput setup logging output destination.
func WithOutputs(outputs ...io.Writer) option {
	return func(lo *options) {
		lo.outputs = outputs
	}
}

// WithTimestampEnabled setup timestamp enabled flag // Enabled by default.
func WithTimestampEnabled(v bool) option {
	return func(lo *options) {
		lo.tsEnabled = v
	}
}

// WithTimestampFieldName setup timestamp field name // "ts" by default.
func WithTimestampFieldName(name string) option {
	return func(lo *options) {
		lo.tsName = name
	}
}

// WithTimestampFormat setup timestamp format // time.RFC3339 by default.
func WithTimestampFormat(format string) option {
	return func(lo *options) {
		lo.tsFormat = format
	}
}

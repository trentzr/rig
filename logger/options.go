package logger

import "io"

// loggerOption is logger modification func type.
type loggerOption func(*loggerOptions)

// WithLevel setup severnity level for logger.
func WithLevel(level Level) loggerOption {
	return func(lo *loggerOptions) {
		lo.level = level
	}
}

// WithFormat setup logging format.
func WithFormat(format Format) loggerOption {
	return func(lo *loggerOptions) {
		lo.format = format
	}
}

// WithOutput setup logging output.
func WithOutput(output io.Writer) loggerOption {
	return func(lo *loggerOptions) {
		lo.outputs = append(lo.outputs, output)
	}
}

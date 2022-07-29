package logger

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

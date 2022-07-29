package logger

// Fields is custom type for WithFields method.
type Fields map[string]interface{}

// WithFields create new instance of logger with target fields.
func (l *Logger) WithFields(fields Fields) *Logger {
	zl := l.Logger.With().Fields(fields).Logger()
	return &Logger{Logger: &zl}
}

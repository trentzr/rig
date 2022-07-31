package logger

// Fields is custom type for WithFields method.
type Fields map[string]interface{}

// WithFields create new instance of logger with target fields.
func (l *Logger) WithFields(fields Fields) *Logger {
	zl := l.Logger.With().Fields(map[string]interface{}(fields)).Logger()
	return &Logger{Logger: &zl}
}

// WithField create new instance of logger with target field.
func (l *Logger) WithField(key string, value interface{}) *Logger {
	zl := l.Logger.With().Interface(key, value).Logger()
	return &Logger{Logger: &zl}
}

// Predefined logger field names.
const (
	FieldNameError = "error"
)

// WithError create new instance of logger with error field.
func (l *Logger) WithError(err error) *Logger {
	zl := l.Logger.With().Str(FieldNameError, err.Error()).Logger()
	return &Logger{Logger: &zl}
}

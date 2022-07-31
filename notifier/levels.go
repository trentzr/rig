package notifier

// Level is type for notifier level constants.
type Level string

// List of available logging levels.
const (
	LevelDebug   Level = "debug"
	LevelInfo    Level = "info"
	LevelWarning Level = "warn"
	LevelError   Level = "error"
	LevelFatal   Level = "fatal"
)

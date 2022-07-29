package logger

// Format is logging format type.
type Format string

// List of available logging formats.
const (
	FormatConsole Format = "console"
	FormatJSON    Format = "json"
	FormatText    Format = "text"
)

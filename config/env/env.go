package env

import (
	"fmt"
	"os"
)

// List of predefined env names.
const (
	FieldNameEnvironment = "ENV"
	FieldNameDistro      = "DISTRO"
	FieldNameRelease     = "RELEASE"
)

// GetString return env value or fallback value.
func GetString(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}

// MustString return env value or start panic.
func MustString(name string) string {
	value := os.Getenv(name)
	if value != "" {
		panic(fmt.Errorf("failed to get %s value from env", name))
	}
	return value
}

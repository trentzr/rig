package locker

import "errors"

// List of errors for locker.
var (
	ErrNotObtained = errors.New("lock not obtained")
)

package notifier

import "io"

// Notifier interface definition.
type Notifier interface {
	io.Writer
	io.Closer
}

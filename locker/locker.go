package locker

import (
	"context"
	"time"
)

type (
	// Locker interface.
	Locker interface {
		Lock(context.Context, string, time.Duration) (UnlockFunc, error)
	}

	// UnlockFunc func type.
	UnlockFunc func() error
)

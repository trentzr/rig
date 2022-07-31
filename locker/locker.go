package locker

import (
	"context"
	"time"
)

type (
	// Locker
	Locker interface {
		Lock(context.Context, string, time.Duration) (UnlockFunc, error)
	}

	// UnlockFunc
	UnlockFunc func() error
)

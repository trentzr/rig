package redis

import (
	"context"
	"time"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"github.com/trentzr/rig/locker"
)

type (
	// redisLocker struct.
	redisLocker struct {
		locker       *redislock.Client
		retryCount   int
		retryTimeout time.Duration
	}
)

// New create new redis locker instance.
func New(client redis.UniversalClient, opts ...option) *redisLocker {

	rl := redislock.New(client)

	l := &redisLocker{
		locker:       rl,
		retryCount:   3,
		retryTimeout: 5 * time.Second,
	}
	for _, opt := range opts {
		opt(l)
	}

	return l
}

// Lock method obtain lock for target key.
func (l *redisLocker) Lock(ctx context.Context, key string, ttl time.Duration) (locker.UnlockFunc, error) {

	lock, err := l.locker.Obtain(ctx, key, ttl, &redislock.Options{
		RetryStrategy: redislock.LimitRetry(
			redislock.LinearBackoff(l.retryTimeout),
			l.retryCount,
		),
	})
	if err != nil {
		if err == redislock.ErrNotObtained {
			return nil, locker.ErrNotObtained
		}
	}

	return func() error {
		return lock.Release(ctx)
	}, nil
}

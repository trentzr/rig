package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/trentzr/rig/locker"
)

type (
	// redisLocker struct.
	redisLocker struct {
		client       redis.UniversalClient
		retryCount   int
		retryTimeout time.Duration
	}

	// lock is helper type for lock.
	lock struct {
		ctx context.Context
	}
)

func New(client redis.UniversalClient, opts ...option) *redisLocker {

	l := &redisLocker{
		client:       client,
		retryCount:   3,
		retryTimeout: 5 * time.Second,
	}
	for _, opt := range opts {
		opt(l)
	}

	return l
}

func (l *redisLocker) Lock(ctx context.Context, key string, ttl time.Duration) (locker.UnlockFunc, error) {
	return nil, nil
}

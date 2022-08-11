package redis

import "time"

// option is locker modifier.
type option func(*redisLocker)

// WithRetryCount setup retry count.
// Default is 3.
func WithRetryCount(n int) option {
	return func(rl *redisLocker) {
		rl.retryCount = n
	}
}

// WithRetryTimeout setup timeout between retries.
// Default 5 sec.
func WithRetryTimeout(timeout time.Duration) option {
	return func(rl *redisLocker) {
		rl.retryTimeout = timeout
	}
}

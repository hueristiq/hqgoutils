// Package ratelimiter provides an implementation of client rate limiting in Go.
package ratelimiter

import (
	"sync"
	"time"
)

// RateLimiter implements rate limiting to limit the number of requests made within a certain time period.
type RateLimiter struct {
	requestsPerMinute     int
	minimumDelayInSeconds int
	timeOfLastRequest     time.Time
	lock                  sync.Mutex
}

// Options implements the structure of RateLimiter creation options.
type Options struct {
	RequestsPerMinute     int
	MinimumDelayInSeconds int
}

// New creates a new *RateLimiter with the specified *Options.
func New(options *Options) (limiter *RateLimiter) {
	limiter = &RateLimiter{
		requestsPerMinute:     options.RequestsPerMinute,
		minimumDelayInSeconds: options.MinimumDelayInSeconds,
		timeOfLastRequest:     time.Now(),
	}

	return
}

// Wait waits until the next request can be made within the rate limit.
func (limiter *RateLimiter) Wait() {
	limiter.lock.Lock()
	defer limiter.lock.Unlock()

	// Calculate the minimum duration to wait before allowing the next request.
	interval := time.Duration(time.Minute.Nanoseconds() / int64(limiter.requestsPerMinute))
	elapsed := time.Since(limiter.timeOfLastRequest)
	remaining := interval - elapsed

	if remaining > 0 {
		// Sleep for the remaining or minimum delay duration to ensure we wait until the rate limit allows the next request.
		timeToSleep := remaining
		if limiter.minimumDelayInSeconds > int(timeToSleep) {
			timeToSleep = time.Duration(limiter.minimumDelayInSeconds) * time.Second
		}

		time.Sleep(timeToSleep)
	}

	// Update the last access time to the current time.
	limiter.timeOfLastRequest = time.Now()
}

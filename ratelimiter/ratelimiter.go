package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	RPM  int        // Requests per minute allowed
	ToLR time.Time  // Time of last request
	lock sync.Mutex // Mutex to ensure thread-safe access to state
}

// New creates a new RateLimiter with the specified requests per minute limit.
func New(RPM int) *RateLimiter {
	return &RateLimiter{
		RPM:  RPM,
		ToLR: time.Now(),
	}
}

// Wait waits until the next request can be made within the rate limit.
func (limiter *RateLimiter) Wait() {
	limiter.lock.Lock()
	defer limiter.lock.Unlock()

	// Calculate the minimum duration to wait before allowing the next request.
	interval := time.Duration(time.Minute.Nanoseconds() / int64(limiter.RPM))
	elapsed := time.Since(limiter.ToLR)
	remaining := interval - elapsed

	if remaining > 0 {
		// Sleep for the remaining duration to ensure we wait until the rate limit allows the next request.
		time.Sleep(remaining)
	}

	// Update the last access time to the current time.
	limiter.ToLR = time.Now()
}

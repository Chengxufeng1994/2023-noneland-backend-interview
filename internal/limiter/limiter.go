package limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity  int
	rate      float64
	tokens    float64
	lastToken time.Time
	mtx       sync.Mutex
}

func (tb *TokenBucket) Allow() bool {
	tb.mtx.Lock()
	defer tb.mtx.Unlock()
	now := time.Now()
	tb.tokens = tb.tokens + tb.rate*now.Sub(tb.lastToken).Seconds()
	if tb.tokens > float64(tb.capacity) {
		tb.tokens = float64(tb.capacity)
	}

	if tb.tokens >= 1 {
		tb.tokens--
		tb.lastToken = now
		return true
	} else {
		return false
	}
}

type RateLimiter interface {
	Allow() bool
}

func New(capacity int, rate float64) RateLimiter {
	tb := &TokenBucket{
		capacity:  capacity,
		rate:      rate,
		tokens:    0,
		lastToken: time.Now(),
	}

	return tb
}

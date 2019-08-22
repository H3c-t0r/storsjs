// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package defence

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"storj.io/storj/internal/sync2"
)

// limited stores information about attacker entity.
type limited struct {
	limiter *rate.Limiter
	expire  time.Time
}

// Limiter is used to store and manage a list of banned entities.
type Limiter struct {
	attackers map[string]*limited

	// Attempts defines how many times attacker could perform an operation.
	attempts int
	// AttemptsPeriod defines period in which attempts will count. For example, 5 attempts per minute.
	attemptsPeriod time.Duration
	lockInterval   time.Duration

	mu   sync.Mutex
	loop *sync2.Cycle
}

// NewLimiter is a constructor for Limiter.
func NewLimiter(attempts int, lockInterval, clearPeriod time.Duration) *Limiter {
	return &Limiter{
		attackers:    map[string]*limited{},
		attempts:     attempts,
		lockInterval: lockInterval,
		loop:         sync2.NewCycle(clearPeriod),
	}
}

// Limit is use to add new fail attack.
func (limiter *Limiter) Limit(key string) bool {
	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	now := time.Now()

	attacker, found := limiter.attackers[key]
	if !found {
		attacker = &limited{
			limiter: rate.NewLimiter(rate.Every(limiter.lockInterval), limiter.attempts),
			expire:  now.Add(limiter.lockInterval),
		}
		limiter.attackers[key] = attacker
	}

	return attacker.limiter.AllowN(now, 1)
}

// Run is used to clean all attackers whose ban is expired.
func (limiter *Limiter) Run(ctx context.Context) error {
	return limiter.loop.Run(ctx, func(ctx context.Context) error {
		return limiter.cleanUp(ctx, time.Now())
	})
}

func (limiter *Limiter) cleanUp(ctx context.Context, cleanUpTime time.Time) error {
	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	for key, limit := range limiter.attackers {
		select {
		case <-ctx.Done():
			limiter.loop.Close()
			return ctx.Err()
		default:
		}

		if cleanUpTime.After(limit.expire) {
			delete(limiter.attackers, key)
		}
	}

	return nil
}

// Close should be used when limiter is no longer needed.
func (limiter *Limiter) Close() {
	limiter.loop.Close()
}

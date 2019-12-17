// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package date_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"storj.io/storj/private/date"
)

func TestMonthBoundary(t *testing.T) {
	now := time.Now().UTC()

	start, end := date.MonthBoundary(now)
	assert.Equal(t, start, time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, end, time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, -1, time.UTC))
}

func TestDayBoundary(t *testing.T) {
	now := time.Now().UTC()

	start, end := date.DayBoundary(now)
	assert.Equal(t, start, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC))
	assert.Equal(t, end, time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, -1, time.UTC))
}

func TestTimeTruncateDown(t *testing.T) {
	now := time.Now().UTC()

	truncated := date.TimeTruncateDown(now)

	assert.Equal(t, now.Year(), truncated.Year())
	assert.Equal(t, now.Month(), truncated.Month())
	assert.Equal(t, now.Day(), truncated.Day())
	assert.Equal(t, now.Hour(), truncated.Hour())
	assert.Equal(t, truncated.Minute(), 0)
	assert.Equal(t, truncated.Second(), 0)
	assert.Equal(t, truncated.Nanosecond(), 0)
}

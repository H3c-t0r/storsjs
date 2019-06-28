// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information

package rewards

import (
	"fmt"
)

// USD describes USD currency in cents.
type USD struct {
	cents int64
}

// USDFromDollars converts dollars to USD amount.
func USDFromDollars(dollars int) USD {
	return USD{int64(dollars) * 100}
}

// USDFromCents converts cents to USD amount.
func USDFromCents(cents int) USD {
	return USD{int64(cents)}
}

// Cents returns amount in cents.
func (usd USD) Cents() int { return int(usd.cents) }

// String returns the value in dollars.
func (usd USD) String() string {
	if usd.cents < 0 {
		return fmt.Sprintf("-%d.%02d", -usd.cents/100, -usd.cents%100)
	}
	return fmt.Sprintf("%d.%02d", usd.cents/100, usd.cents%100)
}

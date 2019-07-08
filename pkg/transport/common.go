// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package transport

import (
	"time"

	"github.com/zeebo/errs"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"
)

var (
	mon = monkit.Package()
	//Error is the errs class of standard Transport Client errors
	Error = errs.Class("transport error")
)

const (
	// DefaultDialTimeout is the default time to wait for a connection to be established
	DefaultDialTimeout = 20 * time.Second
	// DefaultRequestTimeout is the default time to wait for a response
	DefaultRequestTimeout = 20 * time.Second
)

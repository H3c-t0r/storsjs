// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package storagenodedb

import (
	"github.com/zeebo/errs"
)

// ErrSatellitesDB represents errors from the satellites database.
var ErrSatellitesDB = errs.Class("satellitesdb error")

const (
	// SatellitesDBName represents the database name.
	SatellitesDBName = "satellites"
)

// reputation works with node reputation DB
type satellitesDB struct {
	migratableDB
}

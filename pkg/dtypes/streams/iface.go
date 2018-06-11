// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package streams

import (
	"context"
	"io"
	"time"

	"storj.io/storj/pkg/dtypes"
	"storj.io/storj/pkg/ranger"
)

type StreamStore interface {
	Put(ctx context.Context, path dtypes.Path, data io.Reader, metadata []byte,
		expiration time.Time) error
	Get(ctx context.Context, path dtypes.Path) (ranger.Ranger, dtypes.Meta, error)
	Delete(ctx context.Context, path dtypes.Path) error
	List(ctx context.Context, startingPath, endingPath dtypes.Path) (
		paths []dtypes.Path, truncated bool, err error)
}

// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package encryption

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeriveRootKey(t *testing.T) {
	// ensure that we can derive with no errors
	_, err := DeriveRootKey([]byte("password"), []byte("salt"), "")
	assert.NoError(t, err)
	_, err = DeriveRootKey([]byte("password"), []byte("salt"), "any/path")
	assert.NoError(t, err)
}

func TestDeriveDefaultPassword(t *testing.T) {
	// ensure that we can derive with no errors
	_, err := DeriveDefaultPassword([]byte("password"), []byte("salt"))
	assert.NoError(t, err)
}

// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"crypto/rand"
	"testing"

	"github.com/lib/pq"
	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"storj.io/storj/pkg/storj"
)

func TestBytesToUUID(t *testing.T) {
	t.Run("Invalid input", func(t *testing.T) {
		str := "not UUID string"
		bytes := []byte(str)

		_, err := bytesToUUID(bytes)

		assert.NotNil(t, err)
		assert.Error(t, err)
	})

	t.Run("Valid input", func(t *testing.T) {
		id, err := uuid.New()
		assert.NoError(t, err)

		result, err := bytesToUUID(id[:])
		assert.NoError(t, err)
		assert.Equal(t, result, *id)
	})
}

func TestNodeIDsArray(t *testing.T) {
	ids := make(storj.NodeIDList, 10)
	for i := range ids {
		_, _ = rand.Read(ids[i][:])
	}

	got, err := nodeIDsArray(ids).Value() // returns a []byte
	require.NoError(t, err)

	expected, err := pq.ByteaArray(ids.Bytes()).Value() // returns a string
	require.NoError(t, err)

	assert.Equal(t, expected.(string), string(got.([]byte)))
}

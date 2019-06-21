// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package encryption

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"storj.io/storj/pkg/paths"
	"storj.io/storj/pkg/storj"
)

func newStore(key storj.Key) *Store {
	store := NewStore()
	store.Add("bucket", paths.Unencrypted{}, paths.Encrypted{}, key)
	return store
}

func TestStoreEncryption(t *testing.T) {
	forAllCiphers(func(cipher storj.Cipher) {
		for i, rawPath := range []string{
			"",
			"/",
			"//",
			"file.txt",
			"file.txt/",
			"fold1/file.txt",
			"fold1/fold2/file.txt",
			"/fold1/fold2/fold3/file.txt",
		} {
			errTag := fmt.Sprintf("test:%d path:%q cipher:%v", i, rawPath, cipher)

			var key storj.Key
			copy(key[:], randData(storj.KeySize))
			store := newStore(key)
			path := paths.NewUnencrypted(rawPath)

			encPath, err := StoreEncryptPath("bucket", path, cipher, store)
			if !assert.NoError(t, err, errTag) {
				continue
			}

			decPath, err := StoreDecryptPath("bucket", encPath, cipher, store)
			if !assert.NoError(t, err, errTag) {
				continue
			}

			assert.Equal(t, rawPath, decPath.Raw(), errTag)
		}
	})
}

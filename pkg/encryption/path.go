// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package encryption

import (
	"crypto/hmac"
	"crypto/sha512"

	"storj.io/storj/pkg/storj"
)

var (
	emptyComponent          = []byte{'\x01'}
	notEmptyComponentPrefix = byte('\x02')

	escape1 = byte('\x2e')
	escape2 = byte('\xfe')
	escape3 = byte('\x01')
)

// EncryptPath encrypts path with the given key
func EncryptPath(path storj.Path, cipher storj.CipherSuite, key *storj.Key) (encrypted storj.Path, err error) {
	// do not encrypt empty paths
	if len(path) == 0 {
		return path, nil
	}

	if cipher == storj.EncNull {
		return path, nil
	}

	comps := storj.SplitPath(path)
	for i, comp := range comps {
		comps[i], err = encryptPathComponent(comp, cipher, key)
		if err != nil {
			return "", err
		}
		key, err = DeriveKey(key, "path:"+comp)
		if err != nil {
			return "", err
		}
	}
	return storj.JoinPaths(comps...), nil
}

// DecryptPath decrypts path with the given key
func DecryptPath(path storj.Path, cipher storj.CipherSuite, key *storj.Key) (decrypted storj.Path, err error) {
	if cipher == storj.EncNull {
		return path, nil
	}

	comps := storj.SplitPath(path)
	for i, comp := range comps {
		comps[i], err = decryptPathComponent(comp, cipher, key)
		if err != nil {
			return "", err
		}
		key, err = DeriveKey(key, "path:"+comps[i])
		if err != nil {
			return "", err
		}
	}
	return storj.JoinPaths(comps...), nil
}

// DerivePathKey derives the key for the given depth from the given root key.
// This method must be called on an unencrypted path.
func DerivePathKey(path storj.Path, key *storj.Key, depth int) (derivedKey *storj.Key, err error) {
	if depth < 0 {
		return nil, Error.New("negative depth")
	}

	// do not derive key from empty path
	if len(path) == 0 {
		return key, nil
	}

	comps := storj.SplitPath(path)
	if depth > len(comps) {
		return nil, Error.New("depth greater than path length")
	}

	derivedKey = key
	for i := 0; i < depth; i++ {
		derivedKey, err = DeriveKey(derivedKey, "path:"+comps[i])
		if err != nil {
			return nil, err
		}
	}
	return derivedKey, nil
}

// DeriveContentKey derives the key for the encrypted object data using the root key.
// This method must be called on an unencrypted path.
func DeriveContentKey(path storj.Path, key *storj.Key) (derivedKey *storj.Key, err error) {
	comps := storj.SplitPath(path)
	if len(comps) == 0 {
		return nil, Error.New("path is empty")
	}
	derivedKey, err = DerivePathKey(path, key, len(comps))
	if err != nil {
		return nil, err
	}
	derivedKey, err = DeriveKey(derivedKey, "content")
	if err != nil {
		return nil, err
	}
	return derivedKey, nil
}

func encryptPathComponent(comp string, cipher storj.CipherSuite, key *storj.Key) (string, error) {
	// derive the key for the current path component
	derivedKey, err := DeriveKey(key, "path:"+comp)
	if err != nil {
		return "", err
	}

	// use the derived key to derive the nonce
	mac := hmac.New(sha512.New, derivedKey[:])
	_, err = mac.Write([]byte("nonce"))
	if err != nil {
		return "", Error.Wrap(err)
	}

	nonce := new(storj.Nonce)
	copy(nonce[:], mac.Sum(nil))

	// encrypt the path components with the parent's key and the derived nonce
	cipherText, err := Encrypt([]byte(comp), cipher, key, nonce)
	if err != nil {
		return "", Error.Wrap(err)
	}

	nonceSize := storj.NonceSize
	if cipher == storj.EncAESGCM {
		nonceSize = AESGCMNonceSize
	}

	// keep the nonce together with the cipher text
	return string(encodeSegment(append(nonce[:nonceSize], cipherText...))), nil
}

func decryptPathComponent(comp string, cipher storj.CipherSuite, key *storj.Key) (string, error) {
	if comp == "" {
		return "", nil
	}

	data := decodeSegment([]byte(comp))

	nonceSize := storj.NonceSize
	if cipher == storj.EncAESGCM {
		nonceSize = AESGCMNonceSize
	}

	// extract the nonce from the cipher text
	nonce := new(storj.Nonce)
	copy(nonce[:], data[:nonceSize])

	decrypted, err := Decrypt(data[nonceSize:], cipher, key, nonce)
	if err != nil {
		return "", Error.Wrap(err)
	}

	return string(decrypted), nil
}

func encodeSegment(segment []byte) []byte {
	if len(segment) == 0 {
		return emptyComponent
	}

	result := make([]byte, 0, len(segment)*2+1)
	result = append(result, notEmptyComponentPrefix)
	for i := 0; i < len(segment); i++ {
		switch {
		case segment[i] == escape1:
			result = append(result, []byte{escape1, 1}...)
		case segment[i] == escape1+1:
			result = append(result, []byte{escape1, 2}...)
		case segment[i] == escape2:
			result = append(result, []byte{escape2, 1}...)
		case segment[i] == escape2+1:
			result = append(result, []byte{escape2, 2}...)
		case segment[i] == escape3 - 1:
			result = append(result, []byte{escape3, 1}...)
		case segment[i] == escape3:
			result = append(result, []byte{escape3, 2}...)
		default:
			result = append(result, segment[i])
		}
	}
	return result
}

func decodeSegment(segment []byte) []byte {
	if len(segment) == 0 {
		return segment
	}
	if segment[0] == emptyComponent[0] {
		return []byte{}
	}

	// TODO should first byte different than x02 should be invalid?
	result := make([]byte, 0, len(segment))
	for i := 1; i < len(segment); i++ {
		switch {
		case i == len(segment)-1:
			result = append(result, segment[i])
		case segment[i] == escape1 || segment[i] == escape2:
			result = append(result, segment[i]+segment[i+1]-1)
			i++
		case segment[i] == escape3:
			result = append(result, segment[i+1]-1)
			i++
		default:
			result = append(result, segment[i])
		}
	}
	return result
}

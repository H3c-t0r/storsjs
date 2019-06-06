// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package streams

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/eestream"
	"storj.io/storj/pkg/encryption"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/ranger"
	"storj.io/storj/pkg/storage/meta"
	"storj.io/storj/pkg/storage/segments"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/pkg/vpath"
	"storj.io/storj/storage"
)

var (
	mon   = monkit.Package()
	Error = errs.Class("streams")
)

// Meta info about a stream
type Meta struct {
	Modified   time.Time
	Expiration time.Time
	Size       int64
	Data       []byte
}

// convertMeta converts segment metadata to stream metadata
func convertMeta(lastSegmentMeta segments.Meta, stream pb.StreamInfo, streamMeta pb.StreamMeta) Meta {
	return Meta{
		Modified:   lastSegmentMeta.Modified,
		Expiration: lastSegmentMeta.Expiration,
		Size:       ((stream.NumberOfSegments - 1) * stream.SegmentsSize) + stream.LastSegmentSize,
		Data:       stream.Metadata,
	}
}

// Store interface methods for streams to satisfy to be a store
type Store interface {
	Meta(ctx context.Context, path storj.Path, pathCipher storj.Cipher) (Meta, error)
	Get(ctx context.Context, path storj.Path, pathCipher storj.Cipher) (ranger.Ranger, Meta, error)
	Put(ctx context.Context, path storj.Path, pathCipher storj.Cipher, data io.Reader, metadata []byte, expiration time.Time) (Meta, error)
	Delete(ctx context.Context, path storj.Path, pathCipher storj.Cipher) error
	List(ctx context.Context, prefix, startAfter, endBefore storj.Path, pathCipher storj.Cipher, recursive bool, limit int, metaFlags uint32) (items []ListItem, more bool, err error)
}

// streamStore is a store for streams
type streamStore struct {
	segments     segments.Store
	segmentSize  int64
	searcher     *vpath.Searcher
	encBlockSize int
	cipher       storj.Cipher
}

// NewStreamStore stuff
func NewStreamStore(segments segments.Store, segmentSize int64, searcher *vpath.Searcher, encBlockSize int, cipher storj.Cipher) (Store, error) {
	if segmentSize <= 0 {
		return nil, errs.New("segment size must be larger than 0")
	}
	if encBlockSize <= 0 {
		return nil, errs.New("encryption block size must be larger than 0")
	}

	return &streamStore{
		segments:     segments,
		segmentSize:  segmentSize,
		searcher:     searcher,
		encBlockSize: encBlockSize,
		cipher:       cipher,
	}, nil
}

// encryptPath finds the appropriate key for the unencrypted path and encrypts the remainder.
func (s *streamStore) encryptPath(ctx context.Context, path storj.Path, pathCipher storj.Cipher) (encPath storj.Path, pathKey *storj.Key, err error) {
	defer mon.Task()(&ctx)(&err)

	// TODO(jeff): we can do stuff with the returned map if base is nil because it allows us
	// to decrypt specific components. This is important during list.
	_, consumed, base := s.searcher.Lookup(path)
	if base == nil {
		return "", nil, Error.New("cannot find encryption information for: %q", path)
	}

	remainingPath := path[len(consumed):]
	pathKey, err = encryption.DeriveContentKey(remainingPath, &base.Key)
	if err != nil {
		return "", nil, Error.Wrap(err)
	}

	encryptedRemaining, err := encryption.EncryptPath(remainingPath, pathCipher, &base.Key)
	if err != nil {
		return "", nil, Error.Wrap(err)
	}

	return storj.JoinPaths(consumed, encryptedRemaining), pathKey, nil
}

// Put breaks up data as it comes in into s.segmentSize length pieces, then
// store the first piece at s0/<path>, second piece at s1/<path>, and the
// *last* piece at l/<path>. Store the given metadata, along with the number
// of segments, in a new protobuf, in the metadata of l/<path>.
func (s *streamStore) Put(ctx context.Context, path storj.Path, pathCipher storj.Cipher, data io.Reader, metadata []byte, expiration time.Time) (m Meta, err error) {
	defer mon.Task()(&ctx)(&err)

	// previously file uploaded?
	err = s.Delete(ctx, path, pathCipher)
	if err != nil && !storage.ErrKeyNotFound.Has(err) {
		// something wrong happened checking for an existing file with the same name
		return Meta{}, err
	}

	m, lastSegment, err := s.upload(ctx, path, pathCipher, data, metadata, expiration)
	if err != nil {
		s.cancelHandler(context.Background(), lastSegment, path, pathCipher)
	}

	return m, err
}

func (s *streamStore) upload(ctx context.Context, path storj.Path, pathCipher storj.Cipher, data io.Reader, metadata []byte, expiration time.Time) (m Meta, lastSegment int64, err error) {
	defer mon.Task()(&ctx)(&err)

	var currentSegment int64
	var streamSize int64
	var putMeta segments.Meta

	defer func() {
		select {
		case <-ctx.Done():
			s.cancelHandler(context.Background(), currentSegment, path, pathCipher)
		default:
		}
	}()

	encPath, pathKey, err := s.encryptPath(ctx, path, pathCipher)
	if err != nil {
		return Meta{}, currentSegment, Error.Wrap(err)
	}

	eofReader := NewEOFReader(data)

	for !eofReader.isEOF() && !eofReader.hasError() {
		// generate random key for encrypting the segment's content
		var contentKey storj.Key
		_, err = rand.Read(contentKey[:])
		if err != nil {
			return Meta{}, currentSegment, err
		}

		// Initialize the content nonce with the segment's index incremented by 1.
		// The increment by 1 is to avoid nonce reuse with the metadata encryption,
		// which is encrypted with the zero nonce.
		var contentNonce storj.Nonce
		_, err := encryption.Increment(&contentNonce, currentSegment+1)
		if err != nil {
			return Meta{}, currentSegment, err
		}

		encrypter, err := encryption.NewEncrypter(s.cipher, &contentKey, &contentNonce, s.encBlockSize)
		if err != nil {
			return Meta{}, currentSegment, err
		}

		// generate random nonce for encrypting the content key
		var keyNonce storj.Nonce
		_, err = rand.Read(keyNonce[:])
		if err != nil {
			return Meta{}, currentSegment, err
		}

		encryptedKey, err := encryption.EncryptKey(&contentKey, s.cipher, pathKey, &keyNonce)
		if err != nil {
			return Meta{}, currentSegment, err
		}

		sizeReader := NewSizeReader(eofReader)
		segmentReader := io.LimitReader(sizeReader, s.segmentSize)
		peekReader := segments.NewPeekThresholdReader(segmentReader)
		largeData, err := peekReader.IsLargerThan(encrypter.InBlockSize())
		if err != nil {
			return Meta{}, currentSegment, err
		}
		var transformedReader io.Reader
		if largeData {
			paddedReader := eestream.PadReader(ioutil.NopCloser(peekReader), encrypter.InBlockSize())
			transformedReader = encryption.TransformReader(paddedReader, encrypter, 0)
		} else {
			data, err := ioutil.ReadAll(peekReader)
			if err != nil {
				return Meta{}, currentSegment, err
			}
			cipherData, err := encryption.Encrypt(data, s.cipher, &contentKey, &contentNonce)
			if err != nil {
				return Meta{}, currentSegment, err
			}
			transformedReader = bytes.NewReader(cipherData)
		}

		putMeta, err = s.segments.Put(ctx, transformedReader, expiration, func() (storj.Path, []byte, error) {
			if !eofReader.isEOF() {
				segmentPath := getSegmentPath(encPath, currentSegment)

				if s.cipher == storj.Unencrypted {
					return segmentPath, nil, nil
				}

				segmentMeta, err := proto.Marshal(&pb.SegmentMeta{
					EncryptedKey: encryptedKey,
					KeyNonce:     keyNonce[:],
				})
				if err != nil {
					return "", nil, err
				}

				return segmentPath, segmentMeta, nil
			}

			lastSegmentPath := storj.JoinPaths("l", encPath)

			streamInfo, err := proto.Marshal(&pb.StreamInfo{
				NumberOfSegments: currentSegment + 1,
				SegmentsSize:     s.segmentSize,
				LastSegmentSize:  sizeReader.Size(),
				Metadata:         metadata,
			})
			if err != nil {
				return "", nil, err
			}

			// encrypt metadata with the content encryption key and zero nonce
			encryptedStreamInfo, err := encryption.Encrypt(streamInfo, s.cipher, &contentKey, &storj.Nonce{})
			if err != nil {
				return "", nil, err
			}

			streamMeta := pb.StreamMeta{
				EncryptedStreamInfo: encryptedStreamInfo,
				EncryptionType:      int32(s.cipher),
				EncryptionBlockSize: int32(s.encBlockSize),
			}

			if s.cipher != storj.Unencrypted {
				streamMeta.LastSegmentMeta = &pb.SegmentMeta{
					EncryptedKey: encryptedKey,
					KeyNonce:     keyNonce[:],
				}
			}

			lastSegmentMeta, err := proto.Marshal(&streamMeta)
			if err != nil {
				return "", nil, err
			}

			return lastSegmentPath, lastSegmentMeta, nil
		})
		if err != nil {
			return Meta{}, currentSegment, err
		}

		currentSegment++
		streamSize += sizeReader.Size()
	}

	if eofReader.hasError() {
		return Meta{}, currentSegment, eofReader.err
	}

	resultMeta := Meta{
		Modified:   putMeta.Modified,
		Expiration: expiration,
		Size:       streamSize,
		Data:       metadata,
	}

	return resultMeta, currentSegment, nil
}

// getSegmentPath returns the unique path for a particular segment
func getSegmentPath(path storj.Path, segNum int64) storj.Path {
	return storj.JoinPaths(fmt.Sprintf("s%d", segNum), path)
}

// Get returns a ranger that knows what the overall size is (from l/<path>)
// and then returns the appropriate data from segments s0/<path>, s1/<path>,
// ..., l/<path>.
func (s *streamStore) Get(ctx context.Context, path storj.Path, pathCipher storj.Cipher) (rr ranger.Ranger, meta Meta, err error) {
	defer mon.Task()(&ctx)(&err)

	encPath, pathKey, err := s.encryptPath(ctx, path, pathCipher)
	if err != nil {
		return nil, Meta{}, err
	}

	lastSegmentRanger, lastSegmentMeta, err := s.segments.Get(ctx, storj.JoinPaths("l", encPath))
	if err != nil {
		return nil, Meta{}, err
	}

	streamInfo, streamMeta, err := DecryptStreamInfo(ctx, lastSegmentMeta.Data, pathKey)
	if err != nil {
		return nil, Meta{}, err
	}

	stream := pb.StreamInfo{}
	err = proto.Unmarshal(streamInfo, &stream)
	if err != nil {
		return nil, Meta{}, err
	}

	// TODO(jeff): this should be exposed nicely in the encryption package.
	derivedKey, err := encryption.DeriveKey(pathKey, "content")
	if err != nil {
		return nil, Meta{}, err
	}

	var rangers []ranger.Ranger
	for i := int64(0); i < stream.NumberOfSegments-1; i++ {
		currentPath := getSegmentPath(encPath, i)
		size := stream.SegmentsSize
		var contentNonce storj.Nonce
		_, err := encryption.Increment(&contentNonce, i+1)
		if err != nil {
			return nil, Meta{}, err
		}
		rr := &lazySegmentRanger{
			segments:      s.segments,
			path:          currentPath,
			size:          size,
			derivedKey:    derivedKey,
			startingNonce: &contentNonce,
			encBlockSize:  int(streamMeta.EncryptionBlockSize),
			cipher:        storj.Cipher(streamMeta.EncryptionType),
		}
		rangers = append(rangers, rr)
	}

	var contentNonce storj.Nonce
	_, err = encryption.Increment(&contentNonce, stream.NumberOfSegments)
	if err != nil {
		return nil, Meta{}, err
	}
	encryptedKey, keyNonce := getEncryptedKeyAndNonce(streamMeta.LastSegmentMeta)
	decryptedLastSegmentRanger, err := decryptRanger(
		ctx,
		lastSegmentRanger,
		stream.LastSegmentSize,
		storj.Cipher(streamMeta.EncryptionType),
		derivedKey,
		encryptedKey,
		keyNonce,
		&contentNonce,
		int(streamMeta.EncryptionBlockSize),
	)
	if err != nil {
		return nil, Meta{}, err
	}

	rangers = append(rangers, decryptedLastSegmentRanger)
	catRangers := ranger.Concat(rangers...)
	meta = convertMeta(lastSegmentMeta, stream, streamMeta)
	return catRangers, meta, nil
}

// Meta implements Store.Meta
func (s *streamStore) Meta(ctx context.Context, path storj.Path, pathCipher storj.Cipher) (meta Meta, err error) {
	defer mon.Task()(&ctx)(&err)

	encPath, pathKey, err := s.encryptPath(ctx, path, pathCipher)
	if err != nil {
		return Meta{}, err
	}

	lastSegmentMeta, err := s.segments.Meta(ctx, storj.JoinPaths("l", encPath))
	if err != nil {
		return Meta{}, err
	}

	streamInfo, streamMeta, err := DecryptStreamInfo(ctx, lastSegmentMeta.Data, pathKey)
	if err != nil {
		return Meta{}, err
	}
	var stream pb.StreamInfo
	if err := proto.Unmarshal(streamInfo, &stream); err != nil {
		return Meta{}, err
	}

	return convertMeta(lastSegmentMeta, stream, streamMeta), nil
}

// Delete all the segments, with the last one last
func (s *streamStore) Delete(ctx context.Context, path storj.Path, pathCipher storj.Cipher) (err error) {
	defer mon.Task()(&ctx)(&err)

	encPath, pathKey, err := s.encryptPath(ctx, path, pathCipher)
	if err != nil {
		return err
	}

	lastSegmentMeta, err := s.segments.Meta(ctx, storj.JoinPaths("l", encPath))
	if err != nil {
		return err
	}

	streamInfo, _, err := DecryptStreamInfo(ctx, lastSegmentMeta.Data, pathKey)
	if err != nil {
		return err
	}
	var stream pb.StreamInfo
	if err := proto.Unmarshal(streamInfo, &stream); err != nil {
		return err
	}

	for i := 0; i < int(stream.NumberOfSegments-1); i++ {
		currentPath := getSegmentPath(encPath, int64(i))
		err := s.segments.Delete(ctx, currentPath)
		if err != nil {
			return err
		}
	}

	return s.segments.Delete(ctx, storj.JoinPaths("l", encPath))
}

// ListItem is a single item in a listing
type ListItem struct {
	Path     storj.Path
	Meta     Meta
	IsPrefix bool
}

// List all the paths inside l/, stripping off the l/ prefix
func (s *streamStore) List(ctx context.Context, prefix, startAfter, endBefore storj.Path, pathCipher storj.Cipher, recursive bool, limit int, metaFlags uint32) (items []ListItem, more bool, err error) {
	defer mon.Task()(&ctx)(&err)

	if metaFlags&meta.Size != 0 {
		// Calculating the stream's size require also the user-defined metadata,
		// where stream store keeps info about the number of segments and their size.
		metaFlags |= meta.UserDefined
	}

	prefix = strings.TrimSuffix(prefix, "/")

	encPrefix, prefixKey, err := s.encryptPath(ctx, prefix, pathCipher)
	if err != nil {
		return nil, false, err
	}

	encStartAfter, err := s.encryptMarker(ctx, startAfter, pathCipher, prefixKey)
	if err != nil {
		return nil, false, err
	}

	encEndBefore, err := s.encryptMarker(ctx, endBefore, pathCipher, prefixKey)
	if err != nil {
		return nil, false, err
	}

	segments, more, err := s.segments.List(ctx, storj.JoinPaths("l", encPrefix), encStartAfter, encEndBefore, recursive, limit, metaFlags)
	if err != nil {
		return nil, false, err
	}

	items = make([]ListItem, len(segments))
	for i, item := range segments {
		path, err := s.decryptMarker(ctx, item.Path, pathCipher, prefixKey)
		if err != nil {
			return nil, false, err
		}

		// TODO(jeff): this should be exposed nicely in the encryption package.
		pathKey, err := encryption.DeriveKey(prefixKey, "path:"+path)
		if err != nil {
			return nil, false, err
		}

		streamInfo, streamMeta, err := DecryptStreamInfo(ctx, item.Meta.Data, pathKey)
		if err != nil {
			return nil, false, err
		}
		var stream pb.StreamInfo
		if err := proto.Unmarshal(streamInfo, &stream); err != nil {
			return nil, false, err
		}

		newMeta := convertMeta(item.Meta, stream, streamMeta)
		items[i] = ListItem{Path: path, Meta: newMeta, IsPrefix: item.IsPrefix}
	}

	return items, more, nil
}

// encryptMarker is a helper method for encrypting startAfter and endBefore markers
func (s *streamStore) encryptMarker(ctx context.Context, marker storj.Path, pathCipher storj.Cipher, prefixKey *storj.Key) (_ storj.Path, err error) {
	defer mon.Task()(&ctx)(&err)

	// TODO(jeff): Figure out what's going on here. empty prefix was being checked
	// by comparing it to the root key. That's a bit weird.
	return encryption.EncryptPath(marker, pathCipher, prefixKey)
}

// decryptMarker is a helper method for decrypting listed path markers
func (s *streamStore) decryptMarker(ctx context.Context, marker storj.Path, pathCipher storj.Cipher, prefixKey *storj.Key) (_ storj.Path, err error) {
	defer mon.Task()(&ctx)(&err)

	// TODO(jeff): Figure out what's going on here. empty prefix was being checked
	// by comparing it to the root key. That's a bit weird.
	return encryption.DecryptPath(marker, pathCipher, prefixKey)
}

type lazySegmentRanger struct {
	ranger        ranger.Ranger
	segments      segments.Store
	path          storj.Path
	size          int64
	derivedKey    *storj.Key
	startingNonce *storj.Nonce
	encBlockSize  int
	cipher        storj.Cipher
}

// Size implements Ranger.Size
func (lr *lazySegmentRanger) Size() int64 {
	return lr.size
}

// Range implements Ranger.Range to be lazily connected
func (lr *lazySegmentRanger) Range(ctx context.Context, offset, length int64) (_ io.ReadCloser, err error) {
	defer mon.Task()(&ctx)(&err)
	if lr.ranger == nil {
		rr, m, err := lr.segments.Get(ctx, lr.path)
		if err != nil {
			return nil, err
		}
		segmentMeta := pb.SegmentMeta{}
		err = proto.Unmarshal(m.Data, &segmentMeta)
		if err != nil {
			return nil, err
		}
		encryptedKey, keyNonce := getEncryptedKeyAndNonce(&segmentMeta)
		lr.ranger, err = decryptRanger(ctx, rr, lr.size, lr.cipher, lr.derivedKey, encryptedKey, keyNonce, lr.startingNonce, lr.encBlockSize)
		if err != nil {
			return nil, err
		}
	}
	return lr.ranger.Range(ctx, offset, length)
}

// decryptRanger returns a decrypted ranger of the given rr ranger
func decryptRanger(ctx context.Context, rr ranger.Ranger, decryptedSize int64, cipher storj.Cipher, derivedKey *storj.Key, encryptedKey storj.EncryptedPrivateKey, encryptedKeyNonce, startingNonce *storj.Nonce, encBlockSize int) (decrypted ranger.Ranger, err error) {
	defer mon.Task()(&ctx)(&err)
	contentKey, err := encryption.DecryptKey(encryptedKey, cipher, derivedKey, encryptedKeyNonce)
	if err != nil {
		return nil, err
	}

	decrypter, err := encryption.NewDecrypter(cipher, contentKey, startingNonce, encBlockSize)
	if err != nil {
		return nil, err
	}

	var rd ranger.Ranger
	if rr.Size()%int64(decrypter.InBlockSize()) != 0 {
		reader, err := rr.Range(ctx, 0, rr.Size())
		if err != nil {
			return nil, err
		}
		defer func() { err = errs.Combine(err, reader.Close()) }()
		cipherData, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}
		data, err := encryption.Decrypt(cipherData, cipher, contentKey, startingNonce)
		if err != nil {
			return nil, err
		}
		return ranger.ByteRanger(data), nil
	}

	rd, err = encryption.Transform(rr, decrypter)
	if err != nil {
		return nil, err
	}
	return eestream.Unpad(rd, int(rd.Size()-decryptedSize))
}

// CancelHandler handles clean up of segments on receiving CTRL+C
func (s *streamStore) cancelHandler(ctx context.Context, totalSegments int64, path storj.Path, pathCipher storj.Cipher) {
	defer mon.Task()(&ctx)(nil)

	encPath, _, err := s.encryptPath(ctx, path, pathCipher)
	if err != nil {
		zap.S().Warnf("Failed deleting a segment due to encryption: %v", err)
		return
	}

	for i := int64(0); i < totalSegments; i++ {
		currentPath := getSegmentPath(encPath, i)
		err = s.segments.Delete(ctx, currentPath)
		if err != nil {
			zap.S().Warnf("Failed deleting a segment %v %v", currentPath, err)
		}
	}
}

func getEncryptedKeyAndNonce(m *pb.SegmentMeta) (storj.EncryptedPrivateKey, *storj.Nonce) {
	if m == nil {
		return nil, nil
	}

	var nonce storj.Nonce
	copy(nonce[:], m.KeyNonce)

	return m.EncryptedKey, &nonce
}

// DecryptStreamInfo decrypts stream info
func DecryptStreamInfo(ctx context.Context, streamMetaBytes []byte, key *storj.Key) (
	streamInfo []byte, streamMeta pb.StreamMeta, err error) {
	defer mon.Task()(&ctx)(&err)

	err = proto.Unmarshal(streamMetaBytes, &streamMeta)
	if err != nil {
		return nil, pb.StreamMeta{}, err
	}

	cipher := storj.Cipher(streamMeta.EncryptionType)
	encryptedKey, keyNonce := getEncryptedKeyAndNonce(streamMeta.LastSegmentMeta)
	contentKey, err := encryption.DecryptKey(encryptedKey, cipher, key, keyNonce)
	if err != nil {
		return nil, pb.StreamMeta{}, err
	}

	// decrypt metadata with the content encryption key and zero nonce
	streamInfo, err = encryption.Decrypt(streamMeta.EncryptedStreamInfo, cipher, contentKey, &storj.Nonce{})
	return streamInfo, streamMeta, err
}

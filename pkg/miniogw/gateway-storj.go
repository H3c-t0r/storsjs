// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package miniogw

import (
	"context"
	"io"
	"time"

	minio "github.com/minio/minio/cmd"
	"github.com/minio/minio/pkg/auth"
	"github.com/minio/minio/pkg/hash"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/paths"
	"storj.io/storj/pkg/storage/meta"
	"storj.io/storj/pkg/storage/objects"
)

var (
	mon = monkit.Package()
	//Error is the errs class of standard End User Client errors
	Error = errs.Class("Storj Gateway error")
)

func closeHelper(c io.Closer, errptr *error) {
	err := c.Close()
	if err != nil {
		if *errptr == nil {
			*errptr = err
			return
		}
		zap.S().Errorf("error closing: %s", err)
	}
}

// NewStorjGateway creates a *Storj object from an existing ObjectStore
func NewStorjGateway(os objects.Store) *Storj {
	return &Storj{os: os}
}

//Storj is the implementation of a minio cmd.Gateway
type Storj struct {
	os objects.Store
}

// Name implements cmd.Gateway
func (s *Storj) Name() string {
	return "storj"
}

// NewGatewayLayer implements cmd.Gateway
func (s *Storj) NewGatewayLayer(creds auth.Credentials) (
	minio.ObjectLayer, error) {
	return &storjObjects{storj: s}, nil
}

// Production implements cmd.Gateway
func (s *Storj) Production() bool {
	return false
}

type storjObjects struct {
	minio.GatewayUnsupported
	storj *Storj
}

func (s *storjObjects) DeleteBucket(ctx context.Context, bucket string) (err error) {
	defer mon.Task()(&ctx)(&err)
	panic("TODO")
}

func (s *storjObjects) DeleteObject(ctx context.Context, bucket, object string) (err error) {
	defer mon.Task()(&ctx)(&err)
	objpath := paths.New(bucket, object)
	return s.storj.os.Delete(ctx, objpath)
}

func (s *storjObjects) GetBucketInfo(ctx context.Context, bucket string) (
	bucketInfo minio.BucketInfo, err error) {
	defer mon.Task()(&ctx)(&err)
	panic("TODO")
}

func (s *storjObjects) GetObject(ctx context.Context, bucket, object string,
	startOffset int64, length int64, writer io.Writer, etag string) (err error) {
	defer mon.Task()(&ctx)(&err)

	// handle invalid parameters
	if writer == nil || bucket == "" || object == "" {
		return Error.New("Invalid argument(s)")
	}

	objpath := paths.New(bucket, object)
	rr, _, err := s.storj.os.Get(ctx, objpath)
	if err != nil {
		return err
	}
	defer closeHelper(rr, &err)

	r, err := rr.Range(ctx, startOffset, length)
	if err != nil {
		return err
	}
	defer closeHelper(r, &err)

	_, err = io.Copy(writer, r)
	return err
}

func (s *storjObjects) GetObjectInfo(ctx context.Context, bucket,
	object string) (objInfo minio.ObjectInfo, err error) {
	defer mon.Task()(&ctx)(&err)

	// handle invalid parameters
	if (bucket == "") || (object == "") {
		return objInfo, Error.New("Invalid argument(s)")
	}

	objPath := paths.New(bucket, object)
	m, err := s.storj.os.Meta(ctx, objPath)
	if err != nil {
		return objInfo, err
	}

	return minio.ObjectInfo{
		Name:        object,
		Bucket:      bucket,
		ModTime:     m.Modified,
		Size:        m.Size,
		ETag:        m.Checksum,
		ContentType: m.ContentType,
		UserDefined: m.UserDefined,
	}, err
}

func (s *storjObjects) ListBuckets(ctx context.Context) (
	buckets []minio.BucketInfo, err error) {
	defer mon.Task()(&ctx)(&err)
	buckets = nil
	err = nil
	return buckets, err
}

func (s *storjObjects) ListObjects(ctx context.Context, bucket, prefix, marker,
	delimiter string, maxKeys int) (result minio.ListObjectsInfo, err error) {
	defer mon.Task()(&ctx)(&err)
	startAfter := paths.New(marker)
	var fl []minio.ObjectInfo
	items, more, err := s.storj.os.List(ctx, paths.New(bucket, prefix), startAfter, nil, true, maxKeys, meta.All)
	if err != nil {
		return result, err
	}
	if len(items) > 0 {
		//Populate the objectlist (aka filelist)
		f := make([]minio.ObjectInfo, len(items))
		for i, fi := range items {
			f[i] = minio.ObjectInfo{
				Bucket:      fi.Path[0],
				Name:        fi.Path[1:].String(),
				ModTime:     fi.Meta.Modified,
				Size:        fi.Meta.Size,
				ContentType: fi.Meta.ContentType,
				UserDefined: fi.Meta.UserDefined,
				ETag:        fi.Meta.Checksum,
			}
		}
		startAfter = items[len(items)-1].Path[len(paths.New(bucket, prefix)):]
		fl = f
	}

	result = minio.ListObjectsInfo{
		IsTruncated: more,
		Objects:     fl,
	}
	if more {
		result.NextMarker = startAfter.String()
	}

	return result, err
}

func (s *storjObjects) MakeBucketWithLocation(ctx context.Context,
	bucket string, location string) (err error) {
	defer mon.Task()(&ctx)(&err)
	return nil
}

func (s *storjObjects) PutObject(ctx context.Context, bucket, object string,
	data *hash.Reader, metadata map[string]string) (objInfo minio.ObjectInfo,
	err error) {
	defer mon.Task()(&ctx)(&err)

	// handle invalid parameters
	if data == nil || bucket == "" || object == "" {
		return objInfo, Error.New("Invalid argument(s)")
	}

	objPath := paths.New(bucket, object)
	tempContType := metadata["content-type"]
	delete(metadata, "content-type")

	//metadata serialized
	serMetaInfo := objects.SerializableMeta{
		ContentType: tempContType,
		UserDefined: metadata,
	}

	// [TODO @ASK] setting zero value means the object never expires
	expTime := time.Time{}

	m, err := s.storj.os.Put(ctx, objPath, data, serMetaInfo, expTime)
	if err != nil {
		return objInfo, err
	}
	return minio.ObjectInfo{
		Name:        object,
		Bucket:      bucket,
		ModTime:     m.Modified,
		Size:        m.Size,
		ETag:        m.Checksum,
		ContentType: m.ContentType,
		UserDefined: m.UserDefined,
	}, err
}

func (s *storjObjects) Shutdown(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)
	return nil
}

func (s *storjObjects) StorageInfo(context.Context) minio.StorageInfo {
	return minio.StorageInfo{}
}

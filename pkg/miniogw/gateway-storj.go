// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package storj

import (
	"context"
	"io"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/minio/cli"
	minio "github.com/minio/minio/cmd"
	"github.com/minio/minio/pkg/auth"
	"github.com/minio/minio/pkg/hash"
	"github.com/zeebo/errs"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/objects"
	"storj.io/storj/pkg/paths"
	mpb "storj.io/storj/protos/objects"
)

var (
	//Error is the errs class of standard Object Store errors
	Error = errs.Class("objectstore error")
	mon   = monkit.Package()
)

func init() {
	minio.RegisterGatewayCommand(cli.Command{
		Name:            "storj",
		Usage:           "Storj",
		Action:          storjGatewayMain,
		HideHelpCommand: true,
	})
}

func storjGatewayMain(ctx *cli.Context) {
	s := &Storj{os: mockObjectStore()}
	minio.StartGateway(ctx, s)
}

func mockObjectStore() objects.ObjectStore {
	return &objects.Objects{}
}

//S3Bucket structure
type S3Bucket struct {
	bucket   minio.BucketInfo
	filelist S3FileList
}

//S3FileList structure
type S3FileList struct {
	file minio.ListObjectsInfo
}

// Storj is the implementation of a minio cmd.Gateway
type Storj struct {
	bucketlist []S3Bucket
	os         objects.ObjectStore
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
	TempDir string // Temporary storage location for file transfers.
	storj   *Storj
}

func (s *storjObjects) DeleteBucket(ctx context.Context, bucket string) (err error) {
	defer mon.Task()(&ctx)(&err)
	panic("TODO")
}

func (s *storjObjects) DeleteObject(ctx context.Context, bucket, object string) (err error) {
	defer mon.Task()(&ctx)(&err)
	objpath := paths.New(bucket, object)
	return s.storj.os.DeleteObject(ctx, objpath)
}

func (s *storjObjects) GetBucketInfo(ctx context.Context, bucket string) (
	bucketInfo minio.BucketInfo, err error) {
	defer mon.Task()(&ctx)(&err)
	panic("TODO")
}

func (s *storjObjects) GetObject(ctx context.Context, bucket, object string,
	startOffset int64, length int64, writer io.Writer, etag string) (err error) {
	defer mon.Task()(&ctx)(&err)
	objpath := paths.New(bucket, object)
	rr, _, err := s.storj.os.GetObject(ctx, objpath)
	defer rr.Close()
	r, err := rr.Range(ctx, startOffset, length)
	if err != nil {
		return err
	}
	defer r.Close()
	_, err = io.Copy(writer, r)
	return err
}

func (s *storjObjects) GetObjectInfo(ctx context.Context, bucket,
	object string) (objInfo minio.ObjectInfo, err error) {
	defer mon.Task()(&ctx)(&err)
	objPath := paths.New(bucket, object)
	_, m, err := s.storj.os.GetObject(ctx, objPath)
	newmetainfo := &mpb.StorjMetaInfo{}
	err = proto.Unmarshal(m.Data, newmetainfo)
	if err != nil {
		return objInfo, Error.New("ObjectStore GetObject() error")
	}
	objInfo = minio.ObjectInfo{
		ModTime: m.Modified,
	}
	return objInfo, err
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
	result = minio.ListObjectsInfo{}
	err = nil
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
	//metadata serialized
	serMetaInfo := &mpb.StorjMetaInfo{
		Metadata: metadata,
		Bucket:   bucket,
		Name:     object,
	}
	metainfo, err := proto.Marshal(serMetaInfo)
	objPath := paths.New(bucket, object)
	t := time.Now()
	expAfterTenMin := t.Add(time.Minute * 10)
	err = s.storj.os.PutObject(ctx, objPath, data, metainfo, expAfterTenMin)
	return minio.ObjectInfo{
		Name:    object,
		Bucket:  bucket,
		ModTime: time.Now(),
		Size:    data.Size(),
		ETag:    minio.GenETag(),
	}, err
}

func (s *storjObjects) Shutdown(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)
	panic("TODO")
}

func (s *storjObjects) StorageInfo(context.Context) minio.StorageInfo {
	return minio.StorageInfo{}
}

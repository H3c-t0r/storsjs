// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package tally

import (
	"context"
	"time"

	"github.com/zeebo/errs"
	"go.uber.org/zap"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/internal/sync2"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite/accounting"
	"storj.io/storj/satellite/accounting/live"
	"storj.io/storj/satellite/metainfo"
)

// Error is a standard error class for this package.
var (
	Error = errs.Class("tally error")
	mon   = monkit.Package()
)

// Error is a standard error class for this package.
var (
	Error = errs.Class("tally error")
	mon   = monkit.Package()
)

// Config contains configurable values for the tally service
type Config struct {
	Interval time.Duration `help:"how frequently the tally service should run" releaseDefault:"1h" devDefault:"30s"`
}

// Service is the tally service for data stored on each storage node
//
// architecture: Chore
type Service struct {
	log  *zap.Logger
	Loop sync2.Cycle

	metainfoLoop            *metainfo.Loop
	liveAccounting          live.Service
	storagenodeAccountingDB accounting.StoragenodeAccounting
	projectAccountingDB     accounting.ProjectAccounting
}

// New creates a new tally Service
func New(log *zap.Logger, sdb accounting.StoragenodeAccounting, pdb accounting.ProjectAccounting, liveAccounting live.Service, metainfoLoop *metainfo.Loop, interval time.Duration) *Service {
	return &Service{
		log:  log,
		Loop: *sync2.NewCycle(interval),

		metainfoLoop:            metainfoLoop,
		liveAccounting:          liveAccounting,
		storagenodeAccountingDB: sdb,
		projectAccountingDB:     pdb,
	}
}

// Run the tally service loop
func (service *Service) Run(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)
	service.log.Info("Tally service starting up")

	return service.Loop.Run(ctx, func(ctx context.Context) error {
		err := service.Tally(ctx)
		if err != nil {
			service.log.Error("tally failed", zap.Error(err))
		}
		return nil
	})
}

// Close stops the service and releases any resources.
func (service *Service) Close() error {
	service.Loop.Close()
	return nil
}

// Tally calculates data-at-rest usage once
func (service *Service) Tally(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	// The live accounting store will only keep a delta to space used relative
	// to the latest tally. Since a new tally is beginning, we will zero it out
	// now. There is a window between this call and the point where the tally DB
	// transaction starts, during which some changes in space usage may be
	// double-counted (counted in the tally and also counted as a delta to
	// the tally). If that happens, it will be fixed at the time of the next
	// tally run.
	service.liveAccounting.ResetTotals()

	// Fetch when the last tally happened so we can roughly calculate the byte-hours.
	lastTime, err := service.storagenodeAccountingDB.LastTimestamp(ctx, accounting.LastAtRestTally)
	if err != nil {
		return Error.Wrap(err)
	}
	if lastTime.IsZero() {
		lastTime = time.Now()
	}

	// add up all nodes and buckets
	tally := &Tally{
		Node:   make(map[storj.NodeID]float64),
		Bucket: make(map[string]*accounting.BucketTally),
	}
	err = service.metainfoLoop.Join(ctx, tally)
	if err != nil {
		return Error.Wrap(err)
	}
	finishTime := time.Now()

	// calculate byte hours, not just bytes
	hours := time.Since(lastTime).Hours()
	for id := range tally.Node {
		tally.Node[id] *= hours
	}

	// save the new results
	var errAtRest, errBucketInfo error
	if len(tally.Node) > 0 {
		err = service.storagenodeAccountingDB.SaveTallies(ctx, finishTime, tally.Node)
		if err != nil {
			errAtRest = errs.New("StorageNodeAccounting.SaveTallies failed: %v", err)
		}
	}

	if len(tally.Bucket) > 0 {
		err = service.projectAccountingDB.SaveTallies(ctx, finishTime, tally.Bucket)
		if err != nil {
			errAtRest = errs.New("ProjectAccounting.SaveTallies failed: %v", err)
		}
	}

	// report bucket metrics
	if len(tally.Bucket) > 0 {
		var total accounting.BucketTally
		for _, bucket := range tally.Bucket {
			bucketReport(bucket, "bucket")
			total.Combine(bucket)
		}
		bucketReport(&total, "total")
	}

	// return errors if something went wrong.
	return errs.Combine(errAtRest, errBucketInfo)
}

var _ metainfo.Observer = (*Tally)(nil)

// Tally observes metainfo and adds up tallies for nodes and buckets
type Tally struct {
	Node   map[storj.NodeID]float64
	Bucket map[string]*accounting.BucketTally
}

// ensureBucket returns bucket corresponding to the passed in path
func (tally *Tally) ensureBucket(ctx context.Context, path metainfo.ScopedPath) (*accounting.BucketTally, error) {
	bucketID := storj.JoinPaths(path.ProjectIDString, path.BucketName)

	bucket, exists := tally.Bucket[bucketID]
	if !exists {
		bucket = &accounting.BucketTally{}
		bucket.ProjectID = path.ProjectID
		bucket.BucketName = []byte(path.BucketName)
		tally.Bucket[bucketID] = bucket
	}

	return bucket, nil
}

// Object is called for each object once.
func (tally *Tally) Object(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) (err error) {
	bucket, err := tally.ensureBucket(ctx, path)
	if err != nil {
		return err
	}

	bucket.ObjectCount++
	return nil
}

// InlineSegment is called for each inline segment.
func (tally *Tally) InlineSegment(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) (err error) {
	bucket, err := tally.ensureBucket(ctx, path)
	if err != nil {
		return err
	}

	bucket.InlineSegments++
	bucket.InlineBytes += int64(len(pointer.InlineSegment))
	bucket.MetadataSize += int64(len(pointer.Metadata))

	return nil
}

// RemoteSegment is called for each remote segment.
func (tally *Tally) RemoteSegment(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) (err error) {
	bucket, err := tally.ensureBucket(ctx, path)
	if err != nil {
		return err
	}

	bucket.RemoteSegments++
	bucket.RemoteBytes += pointer.GetSegmentSize()
	bucket.MetadataSize += int64(len(pointer.Metadata))

	// add node info
	remote := pointer.GetRemote()
	redundancy := remote.GetRedundancy()
	segmentSize := pointer.GetSegmentSize()
	minimumRequired := redundancy.GetMinReq()

	if remote == nil || redundancy == nil || minimumRequired <= 0 {
		// TODO: tally.log.Error("failed sanity check")
		return nil
	}

	pieceSize := float64(segmentSize / int64(minimumRequired))
	for _, piece := range pointer.GetRemote().GetRemotePieces() {
		tally.Node[piece.NodeId] += pieceSize
	}
	return nil
}

// bucketTallyAdd groups all the data based the passed pointer
func bucketTallyAdd(s *accounting.BucketTally, pointer *pb.Pointer, last bool) {
	switch pointer.GetType() {
	case pb.Pointer_INLINE:
		s.InlineSegments++
		s.InlineBytes += int64(len(pointer.InlineSegment))
		s.MetadataSize += int64(len(pointer.Metadata))

	case pb.Pointer_REMOTE:
		s.RemoteSegments++
		s.RemoteBytes += pointer.GetSegmentSize()
		s.MetadataSize += int64(len(pointer.Metadata))
	}

	if last {
		s.ObjectCount++
	}
}

// using custom name to avoid breaking monitoring
var monAccounting = monkit.ScopeNamed("storj.io/storj/satellite/accounting")

// bucketReport reports the stats thru monkit
func bucketReport(tally *accounting.BucketTally, prefix string) {
	monAccounting.IntVal(prefix + ".objects").Observe(tally.ObjectCount)

	monAccounting.IntVal(prefix + ".segments").Observe(tally.Segments())
	monAccounting.IntVal(prefix + ".inline_segments").Observe(tally.InlineSegments)
	monAccounting.IntVal(prefix + ".remote_segments").Observe(tally.RemoteSegments)

	monAccounting.IntVal(prefix + ".bytes").Observe(tally.Bytes())
	monAccounting.IntVal(prefix + ".inline_bytes").Observe(tally.InlineBytes)
	monAccounting.IntVal(prefix + ".remote_bytes").Observe(tally.RemoteBytes)
}

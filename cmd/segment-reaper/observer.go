// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"encoding/base64"
	"encoding/csv"
	"strconv"

	"github.com/gogo/protobuf/proto"
	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite/metainfo"
)

// object represents object with segments.
type object struct {
	// TODO verify if we have more than 64 segments for object in network
	segments bitmask

	expectedNumberOfSegments byte

	hasLastSegment bool
	// if skip is true then segments from this object shouldn't be treated as zombie segments
	// and printed out, e.g. when one of segments is out of specified date rage
	skip bool
}

// bucketsObjects keeps a list of objects associated with their path per bucket
// name.
type bucketsObjects map[string]map[storj.Path]*object

func newObserver(db metainfo.PointerDB, w *csv.Writer) *observer {
	return &observer{
		db:     db,
		writer: w,

		objects: make(bucketsObjects),
	}
}

// observer metainfo.Loop observer for zombie reaper.
type observer struct {
	db     metainfo.PointerDB
	writer *csv.Writer

	lastProjectID string

	objects            bucketsObjects
	inlineSegments     int
	lastInlineSegments int
	remoteSegments     int
}

// RemoteSegment processes a segment to collect data needed to detect zombie segment.
func (obsvr *observer) RemoteSegment(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) (err error) {
	return obsvr.processSegment(ctx, path, pointer)
}

// InlineSegment processes a segment to collect data needed to detect zombie segment.
func (obsvr *observer) InlineSegment(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) (err error) {
	return obsvr.processSegment(ctx, path, pointer)
}

// Object not used in this implementation.
func (obsvr *observer) Object(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) (err error) {
	return nil
}

func (obsvr *observer) processSegment(ctx context.Context, path metainfo.ScopedPath, pointer *pb.Pointer) error {
	if obsvr.lastProjectID != "" && obsvr.lastProjectID != path.ProjectIDString {
		err := obsvr.analyzeProject(ctx)
		if err != nil {
			return err
		}

		// cleanup map to free memory
		obsvr.clearBucketsObjects()
	}

	obsvr.lastProjectID = path.ProjectIDString

	isLastSegment := path.Segment == "l"
	object := findOrCreate(path.BucketName, path.EncryptedObjectPath, obsvr.objects)
	if isLastSegment {
		object.hasLastSegment = true

		streamMeta := pb.StreamMeta{}
		err := proto.Unmarshal(pointer.Metadata, &streamMeta)
		if err != nil {
			return errs.New("unexpected error unmarshalling pointer metadata %s", err)
		}

		if streamMeta.NumberOfSegments > 0 {
			if streamMeta.NumberOfSegments > int64(maxNumOfSegments) {
				object.skip = true
				zap.S().Warn("unsupported number of segments", zap.Int64("index", streamMeta.NumberOfSegments))
			}
			object.expectedNumberOfSegments = byte(streamMeta.NumberOfSegments)
		}
	} else {
		segmentIndex, err := strconv.Atoi(path.Segment[1:])
		if err != nil {
			return err
		}
		if segmentIndex >= int(maxNumOfSegments) {
			object.skip = true
			zap.S().Warn("unsupported segment index", zap.Int("index", segmentIndex))
		}

		ok, err := object.segments.Has(segmentIndex)
		if err != nil {
			return err
		}
		if ok {
			// TODO make path displayable
			return errs.New("fatal error this segment is duplicated: %s", path.Raw)
		}

		err = object.segments.Set(segmentIndex)
		if err != nil {
			return err
		}
	}

	// collect number of pointers for report
	if pointer.Type == pb.Pointer_INLINE {
		obsvr.inlineSegments++
		if isLastSegment {
			obsvr.lastInlineSegments++
		}
	} else {
		obsvr.remoteSegments++
	}

	return nil
}

// analyzeProject analyzes the objects in obsv.objects field for detecting bad
// segments and writing them to objs.writer.
func (obsvr *observer) analyzeProject(ctx context.Context) error {
	for bucket, objects := range obsvr.objects {
		for path, object := range objects {
			if object.skip {
				continue
			}

			brokenObject := false
			if object.hasLastSegment {
				if object.expectedNumberOfSegments == 0 {
					if !object.segments.IsSequence() {
						brokenObject = true
					}
				} else if object.segments.Count() != int(object.expectedNumberOfSegments)-1 {
					// TODO should we also check if its valid sequence
					// expectedNumberOfSegments-1 because 'segments' doesn't contain last segment
					brokenObject = true
				}

				if brokenObject {
					err := printSegment(ctx, obsvr.db, obsvr.lastProjectID, "l", bucket, path, obsvr.writer)
					if err != nil {
						return err
					}
				}
			} else {
				brokenObject = true
			}

			if brokenObject {
				for index := 0; index < maxNumOfSegments; index++ {
					has, err := object.segments.Has(index)
					if err != nil {
						return nil
					}
					if has {
						err := printSegment(ctx, obsvr.db, obsvr.lastProjectID, "s"+strconv.Itoa(index), bucket, path, obsvr.writer)
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

func printSegment(ctx context.Context, db metainfo.PointerDB, projectID string, bucket string, segmentIndex string, path string, csvWriter *csv.Writer) error {
	creationDate, err := pointerCreationDate(ctx, db, projectID, segmentIndex, bucket, path)
	if err != nil {
		return err
	}
	encodedPath := base64.StdEncoding.EncodeToString([]byte(path))
	csvWriter.Write([]string{
		projectID,
		segmentIndex,
		bucket,
		encodedPath,
		creationDate,
	})
	return nil
}

func pointerCreationDate(ctx context.Context, db metainfo.PointerDB, projectID string, bucket string, segmentIndex string, path string) (string, error) {
	key := []byte(storj.JoinPaths(projectID, segmentIndex, bucket, path))
	pointerBytes, err := db.Get(ctx, key)
	if err != nil {
		return "", err
	}

	pointer := &pb.Pointer{}
	err = proto.Unmarshal(pointerBytes, pointer)
	if err != nil {
		return "", err
	}
	return pointer.CreationDate.String(), nil
}

// clearBucketsObjects clears up the buckets objects map for reusing it.
func (obsvr *observer) clearBucketsObjects() {
	// This is an idiomatic way of not having to destroy and recreate a new map
	// each time that a empty map is required.
	// See https://github.com/golang/go/issues/20138
	for b := range obsvr.objects {
		delete(obsvr.objects, b)
	}
}

func findOrCreate(bucketName string, path string, buckets bucketsObjects) *object {
	objects, ok := buckets[bucketName]
	if !ok {
		objects = make(map[storj.Path]*object)
		buckets[bucketName] = objects
	}

	obj, ok := objects[path]
	if !ok {
		obj = &object{}
		objects[path] = obj
	}

	return obj
}

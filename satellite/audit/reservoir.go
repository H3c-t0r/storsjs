// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package audit

import (
	"math/rand"
	"time"

	"storj.io/common/uuid"
	"storj.io/storj/satellite/metabase"
	"storj.io/storj/satellite/metabase/segmentloop"
)

const maxReservoirSize = 3

// Reservoir holds a certain number of segments to reflect a random sample.
type Reservoir struct {
	Segments [maxReservoirSize]Segment
	size     int8
	index    int64
}

// NewReservoir instantiates a Reservoir.
func NewReservoir(size int) *Reservoir {
	if size < 1 {
		size = 1
	} else if size > maxReservoirSize {
		size = maxReservoirSize
	}
	return &Reservoir{
		size:  int8(size),
		index: 0,
	}
}

// Sample makes sure that for every segment in metainfo from index i=size..n-1,
// pick a random number r = rand(0..i), and if r < size, replace reservoir.Segments[r] with segment.
func (reservoir *Reservoir) Sample(r *rand.Rand, segment Segment) {
	if reservoir.index < int64(reservoir.size) {
		reservoir.Segments[reservoir.index] = segment
	} else {
		random := r.Int63n(reservoir.index + 1)
		if random < int64(reservoir.size) {
			reservoir.Segments[random] = segment
		}
	}
	reservoir.index++
}

// Segment is a segment to audit.
type Segment struct {
	StreamID  uuid.UUID
	Position  metabase.SegmentPosition
	ExpiresAt *time.Time
}

// NewSegment creates a new segment to audit from a metainfo loop segment.
func NewSegment(loopSegment *segmentloop.Segment) Segment {
	return Segment{
		StreamID:  loopSegment.StreamID,
		Position:  loopSegment.Position,
		ExpiresAt: loopSegment.ExpiresAt,
	}
}

// Expired checks if segment is expired relative to now.
func (segment *Segment) Expired(now time.Time) bool {
	return segment.ExpiresAt != nil && segment.ExpiresAt.Before(now)
}

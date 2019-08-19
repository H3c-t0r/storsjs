// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package audit_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"storj.io/storj/internal/memory"
	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/internal/testrand"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite/audit"
	"storj.io/storj/storage"
	"storj.io/storj/uplink"
)

// TestAuditPathCollector does the following:
// - start testplanet with 5 nodes and a reservoir size of 3
// - upload 5 files
// - iterate over all the segments in satellite.Metainfo and store them in allPieces map
// - create a audit observer and call metaloop.Join(auditObs)
//
// Then for every node in testplanet:
//    - expect that there is a reservoir for that node on the audit observer
//    - that the reservoir size is <= 3
//    - that every item in the reservoir is unique
//    - that looking up each pieceID in allPieces results in the same node ID
func TestAuditPathCollector(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 5, UplinkCount: 1,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		audits := planet.Satellites[0].Audit
		satellite := planet.Satellites[0]
		err := audits.Service.Close()
		require.NoError(t, err)

		ul := planet.Uplinks[0]

		// upload 5 remote files with 1 segment
		for i := 0; i < 5; i++ {
			testData := testrand.Bytes(8 * memory.KiB)
			path := "/some/remote/path/" + string(i)
			err := ul.UploadWithConfig(ctx, satellite, &uplink.RSConfig{
				MinThreshold:     3,
				RepairThreshold:  4,
				SuccessThreshold: 5,
				MaxThreshold:     5,
			}, "testbucket", path, testData)
			require.NoError(t, err)
		}

		allPieces := make(map[storj.PieceID]storj.NodeID)

		err = satellite.Metainfo.Service.Iterate(ctx, "", "", true, false,
			func(ctx context.Context, it storage.Iterator) error {
				var item storage.ListItem

				// iterate over every segment in metainfo
				for it.Next(ctx, &item) {
					pointer := &pb.Pointer{}

					err = proto.Unmarshal(item.Value, pointer)
					require.NoError(t, err)

					remote := pointer.GetRemote()
					for _, piece := range remote.GetRemotePieces() {
						pieceID := remote.RootPieceId.Derive(piece.NodeId, piece.PieceNum)
						allPieces[pieceID] = piece.NodeId
					}

					// if context has been canceled exit. Otherwise, continue
					select {
					case <-ctx.Done():
						return ctx.Err()
					default:
					}
				}
				return nil
			})
		require.NoError(t, err)

		r := rand.New(rand.NewSource(time.Now().Unix()))
		observer := audit.NewPathCollector(3, r)
		err = audits.ReservoirService.MetainfoLoop.Join(ctx, observer)
		require.NoError(t, err)

		for _, node := range planet.StorageNodes {
			// expect a reservoir for every node
			require.NotNil(t, observer.Reservoirs[node.ID()])
			require.True(t, len(observer.Reservoirs[node.ID()].Paths) > 1)
			require.True(t, len(observer.Reservoirs[node.ID()].Paths) <= 2)
			repeats := make(map[storj.Path]bool)
			for _, path := range observer.Reservoirs[node.ID()].Paths {
				assert.False(t, repeats[path], "expected every item in reservoir to be unique")
				repeats[path] = true
			}
		}
	})
}

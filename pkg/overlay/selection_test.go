// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package overlay_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeebo/errs"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/pkg/overlay"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite"
	"storj.io/storj/satellite/satellitedb/satellitedbtest"
)

func TestOffline(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 4, UplinkCount: 1,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]
		service := satellite.Overlay.Service
		// TODO: handle cleanup

		result, err := service.KnownUnreliableOrOffline(ctx, []storj.NodeID{
			planet.StorageNodes[0].ID(),
		})
		require.NoError(t, err)
		require.Empty(t, result)

		result, err = service.KnownUnreliableOrOffline(ctx, []storj.NodeID{
			planet.StorageNodes[0].ID(),
			planet.StorageNodes[1].ID(),
			planet.StorageNodes[2].ID(),
		})
		require.NoError(t, err)
		require.Empty(t, result)

		result, err = service.KnownUnreliableOrOffline(ctx, []storj.NodeID{
			planet.StorageNodes[0].ID(),
			storj.NodeID{1, 2, 3, 4}, //note that this succeeds by design
			planet.StorageNodes[2].ID(),
		})
		require.NoError(t, err)
		require.Len(t, result, 1)
		require.Equal(t, result[0], storj.NodeID{1, 2, 3, 4})
	})
}

func BenchmarkOffline(b *testing.B) {
	satellitedbtest.Bench(b, func(b *testing.B, db satellite.DB) {
		const (
			TotalNodeCount = 10000
			OnlineCount    = 90
			OfflineCount   = 10
		)

		overlaydb := db.OverlayCache()
		ctx := context.Background()

		var online []storj.NodeID
		var offline []storj.NodeID

		nodes := make(map[storj.NodeID]bool, TotalNodeCount)
		for i := 0; i < TotalNodeCount; i++ {
			var id storj.NodeID
			_, _ = rand.Read(id[:]) // math/rand never returns error

			overlaydb.UpdateAddress(ctx, &pb.Node{
				Id: id,
			})
			nodes[id] = true
		}

		// pick random node ids to check
		for id := range nodes {
			online = append(online, id)
			if len(online) >= OnlineCount {
				break
			}
		}

		// create random offline node ids to check
		for i := 0; i < OfflineCount; i++ {
			var id storj.NodeID
			_, _ = rand.Read(id[:]) // math/rand never returns error
			offline = append(offline, id)
		}

		var check []storj.NodeID
		check = append(check, offline...)
		check = append(check, online...)

		criteria := &overlay.NodeCriteria{
			AuditCount:         0,
			AuditSuccessRatio:  0,
			OnlineWindow:       1000 * time.Hour,
			UptimeCount:        0,
			UptimeSuccessRatio: 0,
		}

		b.ResetTimer()
		defer b.StopTimer()
		for i := 0; i < b.N; i++ {
			badNodes, err := overlaydb.KnownUnreliableOrOffline(ctx, criteria, check)
			if err != nil {
				b.Fatal(err)
			}
			if len(badNodes) != len(offline) {
				require.Len(b, badNodes, len(offline))
			}
		}
	})
}

func TestNodeSelection(t *testing.T) {
	t.Skip("flaky")
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 10, UplinkCount: 1,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		var err error
		satellite := planet.Satellites[0]

		// This sets a reputable audit count for a certain number of nodes.
		for i, node := range planet.StorageNodes {
			for k := 0; k < i; k++ {
				_, err := satellite.DB.OverlayCache().UpdateStats(ctx, &overlay.UpdateRequest{
					NodeID:       node.ID(),
					IsUp:         true,
					AuditSuccess: true,
				})
				assert.NoError(t, err)
			}
		}

		// ensure all storagenodes are in overlay service
		for _, storageNode := range planet.StorageNodes {
			err = satellite.Overlay.Service.Put(ctx, storageNode.ID(), storageNode.Local().Node)
			assert.NoError(t, err)
		}

		type test struct {
			Preferences    overlay.NodeSelectionConfig
			ExcludeCount   int
			RequestCount   int
			ExpectedCount  int
			ShouldFailWith *errs.Class
		}

		for i, tt := range []test{
			{ // all reputable nodes, only reputable nodes requested
				Preferences: overlay.NodeSelectionConfig{
					AuditCount:        0,
					NewNodePercentage: 0,
					OnlineWindow:      time.Hour,
				},
				RequestCount:  5,
				ExpectedCount: 5,
			},
			{ // all reputable nodes, reputable and new nodes requested
				Preferences: overlay.NodeSelectionConfig{
					AuditCount:        0,
					NewNodePercentage: 1,
					OnlineWindow:      time.Hour,
				},
				RequestCount:  5,
				ExpectedCount: 5,
			},
			{ // all reputable nodes except one, reputable and new nodes requested
				Preferences: overlay.NodeSelectionConfig{
					AuditCount:        1,
					NewNodePercentage: 1,
					OnlineWindow:      time.Hour,
				},
				RequestCount:  5,
				ExpectedCount: 6,
			},
			{ // 50-50 reputable and new nodes, reputable and new nodes requested (new node ratio 1.0)
				Preferences: overlay.NodeSelectionConfig{
					AuditCount:        5,
					NewNodePercentage: 1,
					OnlineWindow:      time.Hour,
				},
				RequestCount:  2,
				ExpectedCount: 4,
			},
			{ // 50-50 reputable and new nodes, reputable and new nodes requested (new node ratio 0.5)
				Preferences: overlay.NodeSelectionConfig{
					AuditCount:        5,
					NewNodePercentage: 0.5,
					OnlineWindow:      time.Hour,
				},
				RequestCount:  4,
				ExpectedCount: 6,
			},
			{ // all new nodes except one, reputable and new nodes requested (happy path)
				Preferences: overlay.NodeSelectionConfig{
					AuditCount:        8,
					NewNodePercentage: 1,
					OnlineWindow:      time.Hour,
				},
				RequestCount:  1,
				ExpectedCount: 2,
			},
			{ // all new nodes except one, reputable and new nodes requested (not happy path)
				Preferences: overlay.NodeSelectionConfig{
					AuditCount:        9,
					NewNodePercentage: 1,
					OnlineWindow:      time.Hour,
				},
				RequestCount:   2,
				ExpectedCount:  3,
				ShouldFailWith: &overlay.ErrNotEnoughNodes,
			},
			{ // all new nodes, reputable and new nodes requested
				Preferences: overlay.NodeSelectionConfig{
					AuditCount:        50,
					NewNodePercentage: 1,
					OnlineWindow:      time.Hour,
				},
				RequestCount:   2,
				ExpectedCount:  2,
				ShouldFailWith: &overlay.ErrNotEnoughNodes,
			},
			{ // audit threshold edge case (1)
				Preferences: overlay.NodeSelectionConfig{
					AuditCount:        9,
					NewNodePercentage: 0,
					OnlineWindow:      time.Hour,
				},
				RequestCount:  1,
				ExpectedCount: 1,
			},
			{ // audit threshold edge case (2)
				Preferences: overlay.NodeSelectionConfig{
					AuditCount:        0,
					NewNodePercentage: 1,
					OnlineWindow:      time.Hour,
				},
				RequestCount:  1,
				ExpectedCount: 1,
			},
			{ // excluded node ids being excluded
				Preferences: overlay.NodeSelectionConfig{
					AuditCount:        5,
					NewNodePercentage: 0,
					OnlineWindow:      time.Hour,
				},
				ExcludeCount:   7,
				RequestCount:   5,
				ExpectedCount:  3,
				ShouldFailWith: &overlay.ErrNotEnoughNodes,
			},
		} {
			t.Logf("#%2d. %+v", i, tt)
			service := planet.Satellites[0].Overlay.Service

			var excludedNodes []storj.NodeID
			for _, storageNode := range planet.StorageNodes[:tt.ExcludeCount] {
				excludedNodes = append(excludedNodes, storageNode.ID())
			}

			response, err := service.FindStorageNodesWithPreferences(ctx, overlay.FindStorageNodesRequest{
				FreeBandwidth:  0,
				FreeDisk:       0,
				RequestedCount: tt.RequestCount,
				ExcludedNodes:  excludedNodes,
			}, &tt.Preferences)

			t.Log(len(response), err)
			if tt.ShouldFailWith != nil {
				assert.Error(t, err)
				assert.True(t, tt.ShouldFailWith.Has(err))
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.ExpectedCount, len(response))
		}
	})
}

// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package bandwidth_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/storagenode"
	"storj.io/storj/storagenode/bandwidth"
	"storj.io/storj/storagenode/storagenodedb/storagenodedbtest"
)

func TestDB(t *testing.T) {
	storagenodedbtest.Run(t, func(t *testing.T, db storagenode.DB) {
		ctx := testcontext.New(t)
		defer ctx.Cleanup()

		bandwidthdb := db.Bandwidth()

		satellite0 := testplanet.MustPregeneratedSignedIdentity(0).ID
		satellite1 := testplanet.MustPregeneratedSignedIdentity(1).ID

		now := time.Now()

		// ensure zero queries work
		usage, err := bandwidthdb.Summary(ctx, now, now)
		require.NoError(t, err)
		require.Equal(t, &bandwidth.Usage{}, usage)

		usageBySatellite, err := bandwidthdb.SummaryBySatellite(ctx, now, now)
		require.NoError(t, err)
		require.Equal(t, map[storj.NodeID]*bandwidth.Usage{}, usageBySatellite)

		actions := []pb.Action{
			pb.Action_INVALID,

			pb.Action_PUT,
			pb.Action_GET,
			pb.Action_GET_AUDIT,
			pb.Action_GET_REPAIR,
			pb.Action_PUT_REPAIR,
			pb.Action_DELETE,

			pb.Action_PUT,
			pb.Action_GET,
			pb.Action_GET_AUDIT,
			pb.Action_GET_REPAIR,
			pb.Action_PUT_REPAIR,
			pb.Action_DELETE,
		}

		expectedUsage := &bandwidth.Usage{}
		expectedUsageTotal := &bandwidth.Usage{}

		// add bandwidth usages
		for _, action := range actions {
			expectedUsage.Include(action, int64(action))
			expectedUsageTotal.Include(action, int64(2*action))

			err := bandwidthdb.Add(ctx, satellite0, action, int64(action), now)
			require.NoError(t, err)

			err = bandwidthdb.Add(ctx, satellite1, action, int64(action), now.Add(2*time.Hour))
			require.NoError(t, err)
		}

		// test summarizing
		usage, err = bandwidthdb.Summary(ctx, now.Add(-10*time.Hour), now.Add(10*time.Hour))
		require.NoError(t, err)
		require.Equal(t, expectedUsageTotal, usage)

		expectedUsageBySatellite := map[storj.NodeID]*bandwidth.Usage{
			satellite0: expectedUsage,
			satellite1: expectedUsage,
		}
		usageBySatellite, err = bandwidthdb.SummaryBySatellite(ctx, now.Add(-10*time.Hour), now.Add(10*time.Hour))
		require.NoError(t, err)
		require.Equal(t, expectedUsageBySatellite, usageBySatellite)

		// only range capturing second satellite
		usage, err = bandwidthdb.Summary(ctx, now.Add(time.Hour), now.Add(10*time.Hour))
		require.NoError(t, err)
		require.Equal(t, expectedUsage, usage)

		// only range capturing second satellite
		expectedUsageBySatellite = map[storj.NodeID]*bandwidth.Usage{
			satellite1: expectedUsage,
		}
		usageBySatellite, err = bandwidthdb.SummaryBySatellite(ctx, now.Add(time.Hour), now.Add(10*time.Hour))
		require.NoError(t, err)
		require.Equal(t, expectedUsage, usage)
	})
}

// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package accounting_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/stretchr/testify/require"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/pkg/accounting"
	"storj.io/storj/pkg/paths"
	"storj.io/storj/satellite"
	"storj.io/storj/satellite/satellitedb/satellitedbtest"
)

func TestSaveBucketTallies(t *testing.T) {
	satellitedbtest.Run(t, func(t *testing.T, db satellite.DB) {
		ctx := testcontext.New(t)
		defer ctx.Cleanup()

		// Setup: create bucket storage tallies
		projectID, err := uuid.New()
		require.NoError(t, err)
		bucketTallies, expectedTallies, err := createBucketStorageTallies(*projectID)
		require.NoError(t, err)

		// Execute test:  retrieve the save tallies and confirm they contains the expected data
		intervalStart := time.Now()
		pdb := db.ProjectAccounting()
		actualTallies, err := pdb.SaveTallies(ctx, intervalStart, bucketTallies)
		require.NoError(t, err)
		for _, tally := range actualTallies {
			require.Contains(t, expectedTallies, tally)
		}
	})
}

func createBucketStorageTallies(projectID uuid.UUID) (map[paths.BucketID]*accounting.BucketTally, []accounting.BucketTally, error) {
	bucketTallies := make(map[paths.BucketID]*accounting.BucketTally)
	var expectedTallies []accounting.BucketTally

	for i := 0; i < 4; i++ {

		bucketName := fmt.Sprintf("%s%d", "testbucket", i)
		bucketID := paths.NewBucketID(projectID, bucketName)
		// bucketIDComponents, err := metainfo.ParsePath(bucketID)
		// if err != nil {
		// 	return nil, nil, err
		// }
		// Setup: The data in this tally should match the pointer that the uplink.upload created
		pid, err := projectID.MarshalJSON()
		if err != nil {
			return nil, nil, err
		}
		tally := accounting.BucketTally{
			BucketName:     []byte(bucketName),
			ProjectID:      pid,
			InlineSegments: int64(1),
			RemoteSegments: int64(1),
			Files:          int64(1),
			InlineBytes:    int64(1),
			RemoteBytes:    int64(1),
			MetadataSize:   int64(1),
		}
		bucketTallies[bucketID] = &tally
		expectedTallies = append(expectedTallies, tally)

	}
	return bucketTallies, expectedTallies, nil
}

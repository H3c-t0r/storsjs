// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package gc_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"storj.io/storj/internal/memory"
	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/internal/testrand"
	"storj.io/storj/pkg/encryption"
	"storj.io/storj/pkg/paths"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/satellite"
)

// TestGarbageCollection does the following:
// * Set up a network with one storagenode
// * Upload two objects
// * Delete one object from the metainfo service on the satellite
// * Trigger a bloom filter generation
// * Check that pieces of the deleted object are deleted on the storagenode
// * Check that pieces of the kept object are not deleted on the storagenode
func TestGarbageCollection(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 1, UplinkCount: 1,
		Reconfigure: testplanet.Reconfigure{
			Satellite: func(log *zap.Logger, index int, config *satellite.Config) {
				config.GarbageCollection.FalsePositiveRate = 0.0001
				config.GarbageCollection.Interval = 0
			},
		},
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		satellite := planet.Satellites[0]
		upl := planet.Uplinks[0]
		targetNode := planet.StorageNodes[0]

		checker := satellite.Repair.Checker
		checker.Loop.Stop()

		// Upload two objects
		testData1 := testrand.Bytes(8 * memory.KiB)
		testData2 := testrand.Bytes(8 * memory.KiB)

		err := upl.Upload(ctx, satellite, "testbucket", "test/path/1", testData1)
		require.NoError(t, err)
		deletedEncPath, pointerToDelete := getPointer(ctx, t, satellite, upl, "testbucket", "test/path/1")
		var deletedPieceID storj.PieceID
		for _, p := range pointerToDelete.GetRemote().GetRemotePieces() {
			if p.NodeId == targetNode.ID() {
				deletedPieceID = pointerToDelete.GetRemote().RootPieceId.Derive(p.NodeId, p.PieceNum)
				break
			}
		}
		require.NotNil(t, deletedPieceID)

		err = upl.Upload(ctx, satellite, "testbucket", "test/path/2", testData2)
		require.NoError(t, err)
		_, pointerToKeep := getPointer(ctx, t, satellite, upl, "testbucket", "test/path/2")
		var keptPieceID storj.PieceID
		for _, p := range pointerToKeep.GetRemote().GetRemotePieces() {
			if p.NodeId == targetNode.ID() {
				keptPieceID = pointerToKeep.GetRemote().RootPieceId.Derive(p.NodeId, p.PieceNum)
				break
			}
		}
		require.NotNil(t, keptPieceID)

		// Delete object from metainfo service on satellite
		err = satellite.Metainfo.Service.Delete(ctx, deletedEncPath)
		require.NoError(t, err)

		// Check that piece of the deleted object is on the storagenode
		pieceInfo, err := targetNode.DB.PieceInfo().Get(ctx, satellite.ID(), deletedPieceID)
		require.NoError(t, err)
		require.NotNil(t, pieceInfo)

		// Sleep is here because we need at least one full second between the next piece tracker
		// creation datetime and the piece upload datetimes
		time.Sleep(1 * time.Second)

		// Trigger bloom filter generation by running checker
		// We trigger it twice because the first piece tracker was created before the upload
		for i := 0; i < 2; i++ {
			err = checker.IdentifyInjuredSegments(ctx)
			require.NoError(t, err)
			checker.WaitForGCSend()
		}

		// Check that piece of the deleted object is not on the storagenode
		pieceInfo, err = targetNode.DB.PieceInfo().Get(ctx, satellite.ID(), deletedPieceID)
		require.Error(t, err)
		require.Nil(t, pieceInfo)

		// Check that piece of the kept object is on the storagenode
		pieceInfo, err = targetNode.DB.PieceInfo().Get(ctx, satellite.ID(), keptPieceID)
		require.NoError(t, err)
		require.NotNil(t, pieceInfo)
	})
}

func getPointer(ctx *testcontext.Context, t *testing.T, satellite *satellite.Peer, upl *testplanet.Uplink, bucket, path string) (lastSegPath string, pointer *pb.Pointer) {
	projects, err := satellite.DB.Console().Projects().GetAll(ctx)
	require.NoError(t, err)
	require.Len(t, projects, 1)

	encParameters := upl.GetConfig(satellite).GetEncryptionParameters()
	cipherSuite := encParameters.CipherSuite
	store := encryption.NewStore()
	store.SetDefaultKey(new(storj.Key))
	encryptedPath, err := encryption.EncryptPath(bucket, paths.NewUnencrypted(path), cipherSuite, store)
	require.NoError(t, err)

	lastSegPath = storj.JoinPaths(projects[0].ID.String(), "l", bucket, encryptedPath.Raw())
	pointer, err = satellite.Metainfo.Service.Get(ctx, lastSegPath)
	require.NoError(t, err)

	return lastSegPath, pointer
}

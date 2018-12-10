// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package tally

import (
	"context"
	"crypto/ecdsa"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	
	"storj.io/storj/internal/identity"
	"storj.io/storj/internal/teststorj"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/pkg/accounting"
	dbManager "storj.io/storj/pkg/bwagreement/database-manager"
	"storj.io/storj/pkg/bwagreement/test"
	"storj.io/storj/pkg/overlay"
	"storj.io/storj/pkg/overlay/mocks"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/pointerdb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/storage/teststore"
	"storj.io/storj/internal/testcontext"
)

var ctx = context.Background()

func TestIdentifyActiveNodes(t *testing.T) {
	//TODO
}
func TestCategorizeNodes(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	planet, err := testplanet.New(t, 1, 30, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Check(planet.Shutdown)

	planet.Start(ctx)

	kad := planet.Satellites[0].Kademlia
	logger := zap.NewNop()
	pointerdb := pointerdb.NewServer(teststore.New(), &overlay.Cache{}, logger, pointerdb.Config{}, nil)

	const N = 50
	nodes := []*pb.Node{}
	nodeIDs := storj.NodeIDList{}
	expectedOnline := []*pb.Node{}
	for i := 0; i < N; i++ {
		nodeID := teststorj.NodeIDFromString(strconv.Itoa(i))
		n := &pb.Node{Id: nodeID, Type: pb.NodeType_STORAGE, Address: &pb.NodeAddress{Address: ""}}
		nodes = append(nodes, n)
		if i%(rand.Intn(5)+2) == 0 {
			id := teststorj.NodeIDFromString("id" + nodeID.String())
			nodeIDs = append(nodeIDs, id)
		} else {
			nodeIDs = append(nodeIDs, nodeID)
			expectedOnline = append(expectedOnline, n)
		}
	}
	overlayServer := mocks.NewOverlay(nodes)
	limit := 0
	interval := time.Second

	accountingDb, err := accounting.NewDB("sqlite3://file::memory:?mode=memory&cache=shared")
	assert.NoError(t, err)
	defer func() { _ = accountingDb.Close() }()

	bwDb, err := dbManager.NewDBManager("sqlite3", "file::memory:?mode=memory&cache=shared")
	assert.NoError(t, err)
	defer func() { _ = accountingDb.Close() }()
	tally, err := newTally(ctx, logger, accountingDb, bwDb, pointerdb, overlayServer, kad, limit, interval)
	assert.NoError(t, err)
	var nodeData = make(map[string]int64)
	online, err := tally.categorize(ctx, nodeIDs, nodeData)
	assert.NoError(t, err)
	assert.Equal(t, expectedOnline, online)
}

func TestTallyAtRestStorage(t *testing.T) {
	//TODO -- 
}

func TestUpdateRawTable(t *testing.T) {
	//TODO
}

func TestQueryNoAgreements(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	planet, err := testplanet.New(t, 1, 30, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Check(planet.Shutdown)

	planet.Start(ctx)

	kad := planet.Satellites[0].Kademlia
	//get stuff we need
	pointerdb := pointerdb.NewServer(teststore.New(), &overlay.Cache{}, zap.NewNop(), pointerdb.Config{}, nil)
	overlayServer := mocks.NewOverlay([]*pb.Node{})
	accountingDb, err := accounting.NewDB("sqlite3://file::memory:?mode=memory&cache=shared")
	assert.NoError(t, err)
	defer func() { _ = accountingDb.Close() }()
	bwDb, err := dbManager.NewDBManager("sqlite3", "file::memory:?mode=memory&cache=shared")
	assert.NoError(t, err)
	defer func() { _ = accountingDb.Close() }()
	tally, err := newTally(ctx, zap.NewNop(), accountingDb, bwDb, pointerdb, overlayServer, kad, 0, time.Second)
	assert.NoError(t, err)
	//check the db
	err = tally.Query(ctx)
	assert.NoError(t, err)
}

func TestQueryWithBw(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	planet, err := testplanet.New(t, 1, 30, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer ctx.Check(planet.Shutdown)

	planet.Start(ctx)

	kad := planet.Satellites[0].Kademlia
	//get stuff we need
	pointerdb := pointerdb.NewServer(teststore.New(), &overlay.Cache{}, zap.NewNop(), pointerdb.Config{}, nil)
	overlayServer := mocks.NewOverlay([]*pb.Node{})
	accountingDb, err := accounting.NewDB("sqlite3://file::memory:?mode=memory&cache=shared")
	assert.NoError(t, err)
	defer func() { _ = accountingDb.Close() }()
	bwDb, err := dbManager.NewDBManager("sqlite3", "file::memory:?mode=memory&cache=shared")
	assert.NoError(t, err)
	defer func() { _ = accountingDb.Close() }()
	tally, err := newTally(ctx, zap.NewNop(), accountingDb, bwDb, pointerdb, overlayServer, kad, 0, time.Second)
	assert.NoError(t, err)
	//get a private key
	fiC, err := testidentity.NewTestIdentity()
	assert.NoError(t, err)
	k, ok := fiC.Key.(*ecdsa.PrivateKey)
	assert.True(t, ok)
	//generate an agreement with the key
	pba, err := test.GeneratePayerBandwidthAllocation(pb.PayerBandwidthAllocation_GET, k)
	assert.NoError(t, err)
	rba, err := test.GenerateRenterBandwidthAllocation(pba, k)
	assert.NoError(t, err)
	//save to db
	_, err = bwDb.Create(ctx, rba)
	assert.NoError(t, err)

	//check the db
	err = tally.Query(ctx)
	assert.NoError(t, err)
}
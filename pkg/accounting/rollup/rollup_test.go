// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package rollup_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/pkg/storj"
)

func TestQuery(t *testing.T) {
	tests := []struct {
		days  int
		rest  float64
		bw    []int64
		nodes int
	}{
		{
			days:  1,
			rest:  float64(5000),
			bw:    []int64{1000, 2000, 3000, 4000},
			nodes: 4,
		},
		{
			days:  2,
			rest:  float64(10000),
			bw:    []int64{2000, 4000, 6000, 8000},
			nodes: 10,
		},
		// TODO: sum data totals per node in QueryPaymentInfo, then this should work
		// {
		// 	days:  3,
		// 	rest:  float64(5000),
		// 	bw:    []int64{1000, 2000, 3000, 4000},
		// 	nodes: 10,
		// },
	}

	for _, tt := range tests {
		ctx := testcontext.New(t)
		defer ctx.Cleanup()

		planet, err := testplanet.New(t, 1, tt.nodes, 0)
		assert.NoError(t, err)
		defer ctx.Check(planet.Shutdown)

		planet.Start(ctx)
		time.Sleep(2 * time.Second)

		nodeData, bwTotals := createData(planet, tt.rest, tt.bw)

		// Set timestamp back by the number of days we want to save
		timestamp := time.Now().UTC().AddDate(0, 0, -1*tt.days)
		start := timestamp

		// Save data for n days
		for i := 0; i < tt.days; i++ {
			err = planet.Satellites[0].DB.Accounting().SaveAtRestRaw(ctx, timestamp, timestamp, nodeData)
			assert.NoError(t, err)

			err = planet.Satellites[0].DB.Accounting().SaveBWRaw(ctx, timestamp, timestamp, bwTotals)
			assert.NoError(t, err)

			// Advance time by 24 hours
			timestamp = timestamp.Add(time.Hour * 24)
		}

		end := timestamp

		err = planet.Satellites[0].Accounting.Rollup.Query(ctx)
		assert.NoError(t, err)

		// rollup.Query cuts off the hr/min/sec before saving, we need to do the same when querying
		start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
		end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())

		rows, err := planet.Satellites[0].DB.Accounting().QueryPaymentInfo(ctx, start, end)
		assert.NoError(t, err)
		if tt.days <= 1 {
			assert.Equal(t, 0, len(rows))
			continue
		}
		assert.Equal(t, (tt.days-1)*tt.nodes, len(rows))

		// verify data is correct
		var nodeIDs []*storj.NodeID
		for _, n := range planet.StorageNodes {
			ptr := n.Identity.ID
			nodeIDs = append(nodeIDs, &ptr)
		}
		for _, r := range rows {
			assert.Equal(t, tt.bw[0], r.PutTotal)
			assert.Equal(t, tt.bw[1], r.GetTotal)
			assert.Equal(t, tt.bw[2], r.GetAuditTotal)
			assert.Equal(t, tt.bw[3], r.GetRepairTotal)
			assert.Equal(t, tt.rest, r.AtRestTotal)
			assert.True(t, contains(r.NodeID, nodeIDs))
		}
	}
}

// contains
func contains(entry storj.NodeID, list []*storj.NodeID) bool {
	for i, n := range list {
		if n == nil {
			continue
		}
		if entry == *n {
			list[i] = nil
			return true
		}
	}
	return false
}

func createData(planet *testplanet.Planet, atRest float64, bw []int64) (nodeData map[storj.NodeID]float64, bwTotals map[storj.NodeID][]int64) {
	nodeData = make(map[storj.NodeID]float64)
	bwTotals = make(map[storj.NodeID][]int64)
	for _, n := range planet.StorageNodes {
		id := n.Identity.ID
		nodeData[id] = atRest
		bwTotals[id] = bw
	}
	return nodeData, bwTotals
}

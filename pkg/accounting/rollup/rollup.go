// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package rollup

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"storj.io/storj/pkg/accounting"
	"storj.io/storj/pkg/storj"
	dbx "storj.io/storj/satellite/satellitedb/dbx"
)

// Rollup is the service for totalling data on storage nodes for 1, 7, 30 day intervals
type Rollup interface {
	Run(ctx context.Context) error
}

type rollup struct {
	logger *zap.Logger
	ticker *time.Ticker
	db     accounting.DB
}

func newRollup(logger *zap.Logger, db accounting.DB, interval time.Duration) *rollup {
	return &rollup{
		logger: logger,
		ticker: time.NewTicker(interval),
		db:     db,
	}
}

// Run the rollup loop
func (r *rollup) Run(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)
	for {
		err = r.Query(ctx)
		if err != nil {
			r.logger.Error("Query failed", zap.Error(err))
		}
		select {
		case <-r.ticker.C: // wait for the next interval to happen
		case <-ctx.Done(): // or the rollup is canceled via context
			return ctx.Err()
		}
	}
}

func (r *rollup) Query(ctx context.Context) error {
	//only rollup new things - get LastRollup
	var latestTally time.Time
	lastRollup, isNil, err := r.db.LastRawTime(ctx, accounting.LastRollup)
	if err != nil {
		return Error.Wrap(err)
	}
	var tallies []*dbx.AccountingRaw
	if isNil {
		r.logger.Info("Rollup found no existing raw tally data")
		tallies, err = r.db.GetRaw(ctx)
	} else {
		tallies, err = r.db.GetRawSince(ctx, lastRollup)
	}
	if err != nil {
		return Error.Wrap(err)
	}
	if len(tallies) == 0 {
		r.logger.Info("Rollup found no new tallies")
		return nil
	}
	//loop through tallies and build rollup
	rollupStats := make(accounting.RollupStats)
	for _, tallyRow := range tallies {
		node, err := storj.NodeIDFromBytes(tallyRow.NodeId)
		if err != nil {
			return Error.Wrap(err)
		}
		if tallyRow.CreatedAt.After(latestTally) {
			latestTally = tallyRow.CreatedAt
		}
		//create or get AccoutingRollup
		iDay := tallyRow.IntervalEndTime
		iDay = time.Date(iDay.Year(), iDay.Month(), iDay.Day(), 0, 0, 0, 0, iDay.Location())
		if rollupStats[iDay] == nil {
			rollupStats[iDay] = make(map[storj.NodeID]*dbx.AccountingRollup, 0)
		}
		if rollupStats[iDay][node] == nil {
			rollupStats[iDay][node] = &dbx.AccountingRollup{NodeId: node.Bytes(), StartTime: iDay}
		}
		//increment Rollups
		switch tallyRow.DataType {
		case accounting.BandwidthPut:
			rollupStats[iDay][node].PutTotal += int64(tallyRow.DataTotal)
		case accounting.BandwidthGet:
			rollupStats[iDay][node].GetTotal += int64(tallyRow.DataTotal)
		case accounting.BandwidthGetAudit:
			rollupStats[iDay][node].GetAuditTotal += int64(tallyRow.DataTotal)
		case accounting.BandwidthGetRepair:
			rollupStats[iDay][node].GetRepairTotal += int64(tallyRow.DataTotal)
		case accounting.BandwidthPutRepair:
			rollupStats[iDay][node].PutRepairTotal += int64(tallyRow.DataTotal)
		case accounting.AtRest:
			rollupStats[iDay][node].AtRestTotal += tallyRow.DataTotal
		default:
			return Error.Wrap(fmt.Errorf("Bad tally datatype in rollup : %d", tallyRow.DataType))
		}
	}
	//push to database
	latestTally = time.Date(latestTally.Year(), latestTally.Month(), latestTally.Day(), 0, 0, 0, 0, latestTally.Location())
	return Error.Wrap(r.db.SaveRollup(ctx, latestTally, rollupStats))
}

// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"

	"github.com/zeebo/errs"

	"github.com/golang/protobuf/ptypes"
	"storj.io/storj/pkg/pb"
	dbx "storj.io/storj/satellite/satellitedb/dbx"
)

type ordersDB struct {
	db *dbx.DB
}

// SaveInlineOrder saves inline order
func (db *ordersDB) SaveInlineOrder(ctx context.Context, bucketID []byte) error {
	return nil
}

// SaveRemoteOrder saves remote order
func (db *ordersDB) SaveRemoteOrder(ctx context.Context, bucketID []byte, orderLimits []*pb.OrderLimit2) error {
	if len(orderLimits) == 0 {
		return nil
	}

	expires, err := ptypes.Timestamp(orderLimits[0].OrderExpiration)
	if err != nil {
		return err
	}

	tx, err := db.db.Open(ctx)
	if err != nil {
		return err
	}

	serialNumber := orderLimits[0].SerialNumber

	_, err = tx.Create_SerialNumber(
		ctx,
		dbx.SerialNumber_SerialNumber(serialNumber.Bytes()),
		dbx.SerialNumber_BucketId(bucketID),
		dbx.SerialNumber_ExpiresAt(expires),
	)
	if err != nil {
		return errs.Combine(err, tx.Rollback())
	}

	// TODO store allocated bandwidth in rollup tables

	return tx.Commit()
}

// SettleOrder settle remote order
func (db *ordersDB) SettleRemoteOrder(ctx context.Context, orderLimit *pb.OrderLimit2, order *pb.Order2) error {
	tx, err := db.db.Open(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			err = tx.Commit()
		} else {
			err = errs.Combine(err, tx.Rollback())
		}
	}()

	serialNumber := dbx.SerialNumber_SerialNumber(order.SerialNumber.Bytes())
	dbxSerialNumber, err := tx.Find_SerialNumber_By_SerialNumber(ctx, serialNumber)
	if err != nil {
		return err
	}

	if dbxSerialNumber == nil {
		return errs.New("serial number not found")
	}

	serialNumberID := dbx.UsedSerial_SerialNumberId(dbxSerialNumber.Id)
	storageNodeID := dbx.UsedSerial_StorageNodeId(orderLimit.StorageNodeId.Bytes())
	_, err = tx.Create_UsedSerial(ctx, serialNumberID, storageNodeID)
	if err != nil {
		return err
	}

	// TODO store settle bandwidth in rollup tables

	return nil
}

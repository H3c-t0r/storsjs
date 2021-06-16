// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package storage

import (
	"context"
	"time"

	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/common/rpc"
	"storj.io/common/storj"
	"storj.io/storj/multinode/nodes"
	"storj.io/storj/private/multinodepb"
)

var (
	mon = monkit.Package()
	// Error is an error class for storage service error.
	Error = errs.Class("storage")
)

// Service exposes all storage related logic.
//
// architecture: Service
type Service struct {
	log    *zap.Logger
	dialer rpc.Dialer
	nodes  nodes.DB
}

// NewService creates new instance of Service.
func NewService(log *zap.Logger, dialer rpc.Dialer, nodes nodes.DB) *Service {
	return &Service{
		log:    log,
		dialer: dialer,
		nodes:  nodes,
	}
}

// Usage retrieves node's daily storage usage for provided interval.
func (service *Service) Usage(ctx context.Context, nodeID storj.NodeID, from, to time.Time) (_ []UsageStamp, err error) {
	defer mon.Task()(&ctx)(&err)

	node, err := service.nodes.Get(ctx, nodeID)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	stamps, err := service.dialUsage(ctx, node, from, to)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	return stamps, nil
}

// UsageSatellite retrieves node's daily storage usage for provided interval and satellite.
func (service *Service) UsageSatellite(ctx context.Context, nodeID, satelliteID storj.NodeID, from, to time.Time) (_ []UsageStamp, err error) {
	defer mon.Task()(&ctx)(&err)

	node, err := service.nodes.Get(ctx, nodeID)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	stamps, err := service.dialUsageSatellite(ctx, node, satelliteID, from, to)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	return stamps, nil
}

// TotalUsage retrieves aggregated daily storage usage for provided interval.
func (service *Service) TotalUsage(ctx context.Context, from, to time.Time) (_ []UsageStamp, err error) {
	defer mon.Task()(&ctx)(&err)

	nodesList, err := service.nodes.List(ctx)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	cache := make(UsageStampDailyCache)

	for _, node := range nodesList {
		stamps, err := service.dialUsage(ctx, node, from, to)
		if err != nil {
			return nil, Error.Wrap(err)
		}

		for _, stamp := range stamps {
			cache.Add(stamp)
		}
	}

	return cache.Sorted(), nil
}

// TotalUsageSatellite retrieves aggregated daily storage usage for provided interval and satellite.
func (service *Service) TotalUsageSatellite(ctx context.Context, satelliteID storj.NodeID, from, to time.Time) (_ []UsageStamp, err error) {
	defer mon.Task()(&ctx)(&err)

	nodesList, err := service.nodes.List(ctx)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	cache := make(UsageStampDailyCache)

	for _, node := range nodesList {
		stamps, err := service.dialUsageSatellite(ctx, node, satelliteID, from, to)
		if err != nil {
			return nil, Error.Wrap(err)
		}

		for _, stamp := range stamps {
			cache.Add(stamp)
		}
	}

	return cache.Sorted(), nil
}

// TotalDiskSpace returns all info about all storagenodes disk space usage.
func (service *Service) TotalDiskSpace(ctx context.Context) (totalDiskSpace DiskSpace, err error) {
	defer mon.Task()(&ctx)(&err)

	nodes, err := service.nodes.List(ctx)
	if err != nil {
		return DiskSpace{}, Error.Wrap(err)
	}

	for _, node := range nodes {
		diskSpace, err := service.dialDiskSpace(ctx, node)
		if err != nil {
			// TODO: how should we handle offline node?
			continue
		}
		totalDiskSpace.Add(diskSpace)
	}

	return totalDiskSpace, nil
}

// DiskSpace returns all info about concrete storagenode disk space usage.
func (service *Service) DiskSpace(ctx context.Context, nodeID storj.NodeID) (_ DiskSpace, err error) {
	defer mon.Task()(&ctx)(&err)

	node, err := service.nodes.Get(ctx, nodeID)
	if err != nil {
		return DiskSpace{}, Error.Wrap(err)
	}

	return service.dialDiskSpace(ctx, node)
}

// dialDiskSpace dials node and retrieves all info about concrete storagenode disk space usage.
func (service *Service) dialDiskSpace(ctx context.Context, node nodes.Node) (diskSpace DiskSpace, err error) {
	conn, err := service.dialer.DialNodeURL(ctx, storj.NodeURL{
		ID:      node.ID,
		Address: node.PublicAddress,
	})
	if err != nil {
		return DiskSpace{}, Error.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, conn.Close())
	}()

	storageClient := multinodepb.NewDRPCStorageClient(conn)

	diskSpaceResponse, err := storageClient.DiskSpace(ctx, &multinodepb.DiskSpaceRequest{
		Header: &multinodepb.RequestHeader{
			ApiKey: node.APISecret,
		},
	})
	if err != nil {
		return DiskSpace{}, Error.Wrap(err)
	}

	return DiskSpace{
		Allocated: diskSpaceResponse.Allocated,
		Used:      diskSpaceResponse.UsedPieces,
		Trash:     diskSpaceResponse.UsedTrash,
		Free:      diskSpaceResponse.Free,
		Available: diskSpaceResponse.Available,
		Overused:  diskSpaceResponse.Overused,
	}, nil
}

// dialUsage dials node and retrieves it's storage usage for provided interval.
func (service *Service) dialUsage(ctx context.Context, node nodes.Node, from, to time.Time) (_ []UsageStamp, err error) {
	defer mon.Task()(&ctx)(&err)

	conn, err := service.dialer.DialNodeURL(ctx, storj.NodeURL{
		ID:      node.ID,
		Address: node.PublicAddress,
	})
	if err != nil {
		return nil, Error.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, conn.Close())
	}()

	storageClient := multinodepb.NewDRPCStorageClient(conn)

	req := &multinodepb.StorageUsageRequest{
		Header: &multinodepb.RequestHeader{
			ApiKey: node.APISecret,
		},
		From: from,
		To:   to,
	}
	resp, err := storageClient.Usage(ctx, req)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	var stamps []UsageStamp
	for _, usage := range resp.GetStorageUsage() {
		stamps = append(stamps, UsageStamp{
			AtRestTotal:   usage.GetAtRestTotal(),
			IntervalStart: usage.GetIntervalStart(),
		})
	}

	return stamps, nil
}

// dialUsageSatellite dials node and retrieves it's storage usage for provided interval and satellite.
func (service *Service) dialUsageSatellite(ctx context.Context, node nodes.Node, satelliteID storj.NodeID, from, to time.Time) (_ []UsageStamp, err error) {
	defer mon.Task()(&ctx)(&err)

	conn, err := service.dialer.DialNodeURL(ctx, storj.NodeURL{
		ID:      node.ID,
		Address: node.PublicAddress,
	})
	if err != nil {
		return nil, Error.Wrap(err)
	}
	defer func() {
		err = errs.Combine(err, conn.Close())
	}()

	storageClient := multinodepb.NewDRPCStorageClient(conn)

	req := &multinodepb.StorageUsageSatelliteRequest{
		Header: &multinodepb.RequestHeader{
			ApiKey: node.APISecret,
		},
		SatelliteId: satelliteID,
		From:        from,
		To:          to,
	}
	resp, err := storageClient.UsageSatellite(ctx, req)
	if err != nil {
		return nil, Error.Wrap(err)
	}

	var stamps []UsageStamp
	for _, usage := range resp.GetStorageUsage() {
		stamps = append(stamps, UsageStamp{
			AtRestTotal:   usage.GetAtRestTotal(),
			IntervalStart: usage.GetIntervalStart(),
		})
	}

	return stamps, nil
}

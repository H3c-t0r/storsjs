// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package heldamount

import (
	"context"
	"time"

	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/common/pb"
	"storj.io/common/rpc"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/storagenode/trust"
)

var (
	// ErrHeldAmountService defines held amount service error
	ErrHeldAmountService = errs.Class("heldamount service error")

	mon = monkit.Package()
)

// Client encapsulates HeldAmountClient with underlying connection
//
// architecture: Client
type Client struct {
	conn *rpc.Conn
	pb.DRPCHeldAmountClient
}

// Close closes underlying client connection
func (c *Client) Close() error {
	return c.conn.Close()
}

// TODO: separate service on service and endpoint.

// Service retrieves info from satellites using an rpc client
//
// architecture: Service
type Service struct {
	log *zap.Logger

	db DB

	dialer rpc.Dialer
	trust  *trust.Pool
}

// NewService creates new instance of service
func NewService(log *zap.Logger, db DB, dialer rpc.Dialer, trust *trust.Pool) *Service {
	return &Service{
		log:    log,
		db:     db,
		dialer: dialer,
		trust:  trust,
	}
}

// GetPaystubStats retrieves held amount for particular satellite from satellite using grpc.
func (service *Service) GetPaystubStats(ctx context.Context, satelliteID storj.NodeID, period string) (_ *PayStub, err error) {
	defer mon.Task()(&ctx)(&err)

	client, err := service.dial(ctx, satelliteID)
	if err != nil {
		return nil, ErrHeldAmountService.Wrap(err)
	}
	defer func() { err = errs.Combine(err, client.Close()) }()

	requestedPeriod, err := stringToTime(period)
	if err != nil {
		service.log.Error("stringToTime", zap.Error(err))
		return nil, ErrHeldAmountService.Wrap(err)
	}

	resp, err := client.GetPayStub(ctx, &pb.GetHeldAmountRequest{Period: requestedPeriod})
	if err != nil {
		service.log.Error("GetPayStub", zap.Error(err))
		return nil, ErrHeldAmountService.Wrap(err)
	}
	service.log.Error("paystub = = = =", zap.Any("", resp))
	return &PayStub{
		Period:         period,
		SatelliteID:    satelliteID,
		Created:        resp.CreatedAt,
		Codes:          resp.Codes,
		UsageAtRest:    float64(resp.UsageAtRest),
		UsageGet:       resp.UsageGet,
		UsagePut:       resp.UsagePut,
		UsageGetRepair: resp.CompGetRepair,
		UsagePutRepair: resp.CompPutRepair,
		UsageGetAudit:  resp.UsageGetAudit,
		CompAtRest:     resp.CompAtRest,
		CompGet:        resp.CompGet,
		CompPut:        resp.CompPut,
		CompGetRepair:  resp.CompGetRepair,
		CompPutRepair:  resp.CompPutRepair,
		CompGetAudit:   resp.CompGetAudit,
		SurgePercent:   resp.SurgePercent,
		Held:           resp.Held,
		Owed:           resp.Owed,
		Disposed:       resp.Disposed,
		Paid:           resp.Paid,
	}, nil
}

// GetPayment retrieves payment data from particular satellite using grpc.
func (service *Service) GetPayment(ctx context.Context, satelliteID storj.NodeID, period string) (_ *Payment, err error) {
	defer mon.Task()(&ctx)(&err)

	client, err := service.dial(ctx, satelliteID)
	if err != nil {
		return nil, ErrHeldAmountService.Wrap(err)
	}
	defer func() { err = errs.Combine(err, client.Close()) }()

	requestedPeriod, err := stringToTime(period)
	if err != nil {
		return nil, ErrHeldAmountService.Wrap(err)
	}

	resp, err := client.GetPayment(ctx, &pb.GetPaymentRequest{Period: requestedPeriod})
	if err != nil {
		return nil, ErrHeldAmountService.Wrap(err)
	}

	return &Payment{
		ID:          resp.Id,
		Created:     resp.CreatedAt,
		SatelliteID: satelliteID,
		Period:      period,
		Amount:      resp.Amount,
		Receipt:     resp.Receipt,
		Notes:       resp.Notes,
	}, nil
}

// SatellitePayStubMonthlyCached retrieves held amount for particular satellite for selected month from storagenode database.
func (service *Service) SatellitePayStubMonthlyCached(ctx context.Context, satelliteID storj.NodeID, period string) (payStub *PayStub, err error) {
	defer mon.Task()(&ctx, &satelliteID, &period)(&err)

	payStub, err = service.db.GetPayStub(ctx, satelliteID, period)
	if err != nil {
		return nil, ErrHeldAmountService.Wrap(err)
	}

	return payStub, nil
}

// AllPayStubsMonthlyCached retrieves held amount for particular satellite from storagenode database.
func (service *Service) AllPayStubsMonthlyCached(ctx context.Context, period string) (payStubs []PayStub, err error) {
	defer mon.Task()(&ctx, &period)(&err)

	payStubs, err = service.db.AllPayStubs(ctx, period)
	if err != nil {
		return nil, ErrHeldAmountService.Wrap(err)
	}

	return payStubs, nil
}

// GetPaymentCached retrieves payment data from particular satellite from storagenode database.
func (service *Service) GetPaymentCached(ctx context.Context, satelliteID storj.NodeID, period string) (_ *Payment, err error) {
	defer mon.Task()(&ctx, &satelliteID, &period)(&err)

	return service.db.GetPayment(ctx, satelliteID, period)
}

// dial dials the HeldAmount client for the satellite by id
func (service *Service) dial(ctx context.Context, satelliteID storj.NodeID) (_ *Client, err error) {
	defer mon.Task()(&ctx)(&err)

	address, err := service.trust.GetAddress(ctx, satelliteID)
	if err != nil {
		return nil, errs.New("unable to find satellite %s: %w", satelliteID, err)
	}

	conn, err := service.dialer.DialAddressID(ctx, address, satelliteID)
	if err != nil {
		return nil, errs.New("unable to connect to the satellite %s: %w", satelliteID, err)
	}

	return &Client{
		conn:                 conn,
		DRPCHeldAmountClient: pb.NewDRPCHeldAmountClient(conn.Raw()),
	}, nil
}

func stringToTime(period string) (_ time.Time, err error) {
	layout := "2006-01"
	per := period[0:7]
	result, err := time.Parse(layout, per)
	if err != nil {
		return time.Time{}, err
	}

	return result, nil
}

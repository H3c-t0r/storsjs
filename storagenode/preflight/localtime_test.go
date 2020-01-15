// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package preflight_test

import (
	"context"
	"net"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"

	"storj.io/common/pb"
	"storj.io/common/peertls/tlsopts"
	"storj.io/common/rpc"
	"storj.io/storj/pkg/server"

	"storj.io/common/identity/testidentity"
	"storj.io/common/peertls/extensions"
	"storj.io/common/testcontext"
	"storj.io/storj/private/testplanet"
	"storj.io/storj/storagenode/preflight"
	"storj.io/storj/storagenode/trust"
)

type mockServer struct {
	localTime time.Time
	pb.NodeServer
}

func TestLocalTime_InSync(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount: 1, StorageNodeCount: 1, UplinkCount: 0,
	}, func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
		storagenode := planet.StorageNodes[0]
		err := storagenode.Preflight.LocalTime.Check(ctx)
		require.NoError(t, err)
	})
}

func TestLocalTime_OutOfSync(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	log := zaptest.NewLogger(t)

	// set up mock satellite server configuration
	mockSatID, err := testidentity.NewTestIdentity(ctx)
	require.NoError(t, err)
	config := server.Config{
		Address:        "127.0.0.1:0",
		PrivateAddress: "127.0.0.1:0",

		Config: tlsopts.Config{
			PeerIDVersions: "*",
			Extensions: extensions.Config{
				Revocation:          false,
				WhitelistSignedLeaf: false,
			},
		},
	}
	mockSatTLSOptions, err := tlsopts.NewOptions(mockSatID, config.Config, nil)
	require.NoError(t, err)

	t.Run("Less than 24h", func(t *testing.T) {
		// register mock GetTime endpoint to mock server
		contactServer, err := server.New(log, mockSatTLSOptions, config.Address, config.PrivateAddress, nil)
		require.NoError(t, err)
		defer func() {
			err := contactServer.Close()
			require.NoError(t, err)
		}()
		pb.DRPCRegisterNode(contactServer.DRPC(), &mockServer{
			localTime: time.Now().UTC().Add(-2 * time.Hour),
		})

		go func() {
			err := contactServer.Run(ctx)
			require.NoError(t, err)
		}()

		// get mock server address
		_, portStr, err := net.SplitHostPort(contactServer.Addr().String())
		require.NoError(t, err)
		port, err := strconv.Atoi(portStr)
		require.NoError(t, err)
		url := trust.SatelliteURL{
			ID:   mockSatID.ID,
			Host: "127.0.0.1",
			Port: port,
		}
		require.NoError(t, err)

		// set up storagenode client
		source, err := trust.NewStaticURLSource(url.String())
		require.NoError(t, err)

		identity, err := testidentity.NewTestIdentity(ctx)
		require.NoError(t, err)
		tlsOptions, err := tlsopts.NewOptions(identity, config.Config, nil)
		require.NoError(t, err)
		dialer := rpc.NewDefaultDialer(tlsOptions)
		pool, err := trust.NewPool(log, trust.Dialer(dialer), trust.Config{
			Sources:   []trust.Source{source},
			CachePath: ctx.File("trust-cache.json"),
		})
		require.NoError(t, err)
		err = pool.Refresh(ctx)
		require.NoError(t, err)

		// should not return any error when node's clock is off no more than 24
		localtime := preflight.NewLocalTime(log, preflight.Config{
			EnabledLocalTime: true,
		}, pool, dialer)
		err = localtime.Check(ctx)
		require.NoError(t, err)
	})

	t.Run("More than 24h", func(t *testing.T) {
		// register mock GetTime endpoint to mock server
		contactServer, err := server.New(log, mockSatTLSOptions, config.Address, config.PrivateAddress, nil)
		require.NoError(t, err)
		defer func() {
			err := contactServer.Close()
			require.NoError(t, err)
		}()

		pb.DRPCRegisterNode(contactServer.DRPC(), &mockServer{
			localTime: time.Now().UTC().Add(-25 * time.Hour),
		})

		go func() {
			err := contactServer.Run(ctx)
			require.NoError(t, err)
		}()

		// get mock server address
		_, portStr, err := net.SplitHostPort(contactServer.Addr().String())
		require.NoError(t, err)
		port, err := strconv.Atoi(portStr)
		require.NoError(t, err)
		url := trust.SatelliteURL{
			ID:   mockSatID.ID,
			Host: "127.0.0.1",
			Port: port,
		}
		require.NoError(t, err)

		// set up storagenode client
		source, err := trust.NewStaticURLSource(url.String())
		require.NoError(t, err)

		identity, err := testidentity.NewTestIdentity(ctx)
		require.NoError(t, err)
		tlsOptions, err := tlsopts.NewOptions(identity, config.Config, nil)
		require.NoError(t, err)
		dialer := rpc.NewDefaultDialer(tlsOptions)
		pool, err := trust.NewPool(log, trust.Dialer(dialer), trust.Config{
			Sources:   []trust.Source{source},
			CachePath: ctx.File("trust-cache.json"),
		})
		require.NoError(t, err)
		err = pool.Refresh(ctx)
		require.NoError(t, err)

		// should return an error when node's clock is off by more than 24h with all trusted satellites
		localtime := preflight.NewLocalTime(log, preflight.Config{
			EnabledLocalTime: true,
		}, pool, dialer)
		err = localtime.Check(ctx)
		require.Error(t, err)
	})
}

func (mock *mockServer) GetTime(ctx context.Context, req *pb.GetTimeRequest) (*pb.GetTimeResponse, error) {
	return &pb.GetTimeResponse{
		Timestamp: mock.localTime,
	}, nil
}

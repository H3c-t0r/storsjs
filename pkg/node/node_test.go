// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package node

import (
	"context"
	"fmt"
	"net"
	"testing"

	"storj.io/storj/pkg/dht/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"storj.io/storj/pkg/provider"
	proto "storj.io/storj/protos/overlay"
)

var ctx = context.Background()

func TestLookup(t *testing.T) {
	cases := []struct {
		self        proto.Node
		to          proto.Node
		find        proto.Node
		expectedErr error
	}{
		{
			self:        proto.Node{Id: "hello", Address: &proto.NodeAddress{Address: ":7070"}},
			to:          proto.Node{Id: "hello", Address: &proto.NodeAddress{Address: ":8080"}},
			find:        proto.Node{Id: "hello", Address: &proto.NodeAddress{Address: ":9090"}},
			expectedErr: nil,
		},
	}

	for _, v := range cases {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
		assert.NoError(t, err)

		srv, mock, err := newTestServer(ctx)
		assert.NoError(t, err)
		go srv.Serve(lis)
		defer srv.Stop()
		ctrl := gomock.NewController(t)

		mdht := mock_dht.NewMockDHT(ctrl)
		mrt := mock_dht.NewMockRoutingTable(ctrl)

		mdht.EXPECT().GetRoutingTable(gomock.Any()).Return(mrt, nil)
		mrt.EXPECT().ConnectionSuccess(gomock.Any()).Return(nil)

		ca, err := provider.NewCA(ctx, 12, 4)
		assert.NoError(t, err)
		identity, err := ca.NewIdentity()
		assert.NoError(t, err)

		nc, err := NewNodeClient(identity, v.self, mdht)
		assert.NoError(t, err)

		_, err = nc.Lookup(ctx, v.to, v.find)
		assert.Equal(t, v.expectedErr, err)
		assert.Equal(t, 1, mock.queryCalled)
	}
}

func newTestServer(ctx context.Context) (*grpc.Server, *mockNodeServer, error) {
	ca, err := provider.NewCA(ctx, 12, 4)
	if err != nil {
		return nil, nil, err
	}
	identity, err := ca.NewIdentity()
	if err != nil {
		return nil, nil, err
	}
	identOpt, err := identity.ServerOption()
	if err != nil {
		return nil, nil, err
	}

	grpcServer := grpc.NewServer(identOpt)
	mn := &mockNodeServer{queryCalled: 0}

	proto.RegisterNodesServer(grpcServer, mn)

	return grpcServer, mn, nil

}

type mockNodeServer struct {
	queryCalled int
}

func (mn *mockNodeServer) Query(ctx context.Context, req *proto.QueryRequest) (*proto.QueryResponse, error) {
	mn.queryCalled++
	return &proto.QueryResponse{}, nil
}

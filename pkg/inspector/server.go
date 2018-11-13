package inspector

import (
	"context"
	"fmt"

	"github.com/zeebo/errs"
	"go.uber.org/zap"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/dht"
	"storj.io/storj/pkg/overlay"
	"storj.io/storj/pkg/pb"
)

var (
	// ServerError is a gRPC server error for Inspector
	ServerError = errs.Class("inspector server error:")
)

// Server holds references to cache and kad
type Server struct {
	dht     dht.DHT
	cache   *overlay.Cache
	logger  *zap.Logger
	metrics *monkit.Registry
}

// CountNodes returns the number of nodes in the cache and in kademlia
func (srv *Server) CountNodes(ctx context.Context, req *pb.CountNodesRequest) (*pb.CountNodesResponse, error) {
	return &pb.CountNodesResponse{
		Kademlia: 0,
		Overlay:  0,
	}, nil
}

// GetBuckets returns all kademlia buckets for current kademlia instance
func (srv *Server) GetBuckets(ctx context.Context, req *pb.GetBucketsRequest) (*pb.GetBucketsResponse, error) {
	results := []*pb.BucketList{}
	rt, err := srv.dht.GetRoutingTable(ctx)
	if err != nil {
		return &pb.GetBucketsResponse{}, ServerError.New("")
	}

	buckets, err := rt.GetBuckets()
	number := len(buckets)

	if err != nil {
		return nil, err
	}

	for _, b := range buckets {
		fmt.Printf("\nRanging over bucket: %+v\n", b)
	}

	return &pb.GetBucketsResponse{
		Total:   int64(number),
		Buckets: results,
	}, nil
}

// func createBucketList(bucket dht.Bucket) *pb.BucketList {
// 	var bl []*pb.BucketList
// 	for _, n := range bucket {
// 		node := &pb.Node{
// 			Id: n.Id,
// 			Address: &pb.NodeAddress{
// 				Address: n.Address.Address,
// 			},
// 		}
// 		// push to list
// 		bl = append(bl, n)
// 	}
// 	return bl
// }

// GetBucket retrieves all of a given K buckets contents
func (srv *Server) GetBucket(ctx context.Context, req *pb.GetBucketRequest) (*pb.GetBucketResponse, error) {
	rt, err := srv.dht.GetRoutingTable(ctx)
	if err != nil {
		return nil, err
	}
	bucket, ok := rt.GetBucket(req.Id)
	if !ok {
		return &pb.GetBucketResponse{}, nil
	}

	return &pb.GetBucketResponse{
		Id:    req.Id,
		Nodes: bucket.Nodes(),
	}, nil
}

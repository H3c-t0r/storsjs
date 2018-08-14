// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package kademlia

import (
	"context"
	"encoding/binary"
	"encoding/hex"

	"sync"
	"time"

	"github.com/zeebo/errs"
	"go.uber.org/zap"
	"storj.io/storj/pkg/dht"
	"storj.io/storj/pkg/node"
	proto "storj.io/storj/protos/overlay"
	"storj.io/storj/storage"
	"storj.io/storj/storage/boltdb"
)

// RoutingErr is the class for all errors pertaining to routing table operations
var RoutingErr = errs.Class("routing table error")

// RoutingTable implements the RoutingTable interface
type RoutingTable struct {
	self         *proto.Node
	kadBucketDB  storage.KeyValueStore
	nodeBucketDB storage.KeyValueStore
	transport    *proto.NodeTransport
	mutex        *sync.Mutex
	nodeClient   node.Client
	idLength     int // kbucket and node id bit length (SHA256) = 256
	bucketSize   int // max number of nodes stored in a kbucket = 20 (k)
}

//RoutingOptions for configuring RoutingTable
type RoutingOptions struct {
	kpath      string
	npath      string
	idLength   int
	bucketSize int
}

// NewRoutingTable returns a newly configured instance of a RoutingTable
func NewRoutingTable(localNode *proto.Node, options *RoutingOptions) (*RoutingTable, error) {
	logger := zap.L()
	kdb, err := boltdb.NewClient(logger, options.kpath, boltdb.KBucket)
	if err != nil {
		return nil, RoutingErr.New("could not create kadBucketDB: %s", err)
	}
	ndb, err := boltdb.NewClient(logger, options.npath, boltdb.NodeBucket)
	if err != nil {
		return nil, RoutingErr.New("could not create nodeBucketDB: %s", err)
	}
	rt := &RoutingTable{
		self:         localNode,
		kadBucketDB:  kdb,
		nodeBucketDB: ndb,
		transport:    &defaultTransport,
		mutex:        &sync.Mutex{},
		idLength:     options.idLength,
		bucketSize:   options.bucketSize,
	}
	err = rt.addNode(localNode)
	if err != nil {
		return nil, RoutingErr.New("could not add localNode to routing table: %s", err)
	}
	return rt, nil
}

// Local returns the local nodes ID
func (rt *RoutingTable) Local() proto.Node {
	return *rt.self
}

// K returns the currently configured maximum of nodes to store in a bucket
func (rt *RoutingTable) K() int {
	return rt.bucketSize
}

// CacheSize returns the total current size of the cache (number of nodes)
func (rt *RoutingTable) CacheSize() int {
	nodeKeys, err := rt.kadBucketDB.List(nil, 0)
	if err != nil {

	}
	return len(nodeKeys)
}

// GetBucket retrieves the corresponding kbucket from node id
// Note: id doesn't need to be stored at time of search
func (rt *RoutingTable) GetBucket(id string) (bucket dht.Bucket, ok bool) {
	i, err := hex.DecodeString(id)
	if err != nil {
		return &KBucket{}, false
	}
	bucketID, err := rt.getKBucketID(i)
	if err != nil {
		return &KBucket{}, false
	}
	if bucketID == nil {
		return &KBucket{}, false
	}
	unmarshaledNodes, err := rt.getUnmarshaledNodesFromBucket(bucketID)
	if err != nil {
		return &KBucket{}, false
	}
	return &KBucket{nodes: unmarshaledNodes}, true
}

// GetBuckets retrieves all buckets from the local node
func (rt *RoutingTable) GetBuckets() (k []dht.Bucket, err error) {
	bs := []dht.Bucket{}
	kbuckets, err := rt.kadBucketDB.List(nil, 0)
	if err != nil {
		return bs, RoutingErr.New("could not get bucket ids %s", err)
	}
	for _, v := range kbuckets {
		unmarshaledNodes, err := rt.getUnmarshaledNodesFromBucket(v)
		if err != nil {
			return bs, err
		}
		bs = append(bs, &KBucket{nodes: unmarshaledNodes})
	}
	return bs, nil
}

// FindNear finds all Nodes near the provided nodeID up to the provided limit
// 1. check if target is in routing table, if not add target to rt
// 1a. If target was successfully added, return target
// 2. If target was not successfully added, return nearest k nodes to sender
// 2a. If target is in routing table already, return target
func (rt *RoutingTable) FindNear(sender, target *proto.Node, limit int) ([]*proto.Node, error) {
	nodeInRT, err := rt.nodeAlreadyExists(storage.Key(target.Id))
	if err != nil {
		//TODO
	}
	if !nodeInRT {
		err := rt.addNode(target)
	}
	nodes, err := rt.findNear(StringToNodeID(target.Id), rt.K())
	if err != nil {
		//TODO
	}
	return nodes, nil
}

// ConnectionSuccess handles the details of what kademlia should do when
// a successful connection is made to node on the network
func (rt *RoutingTable) ConnectionSuccess(id string, address proto.NodeAddress) {
	// TODO: What should we do ?
	
	return
}

// ConnectionFailed handles the details of what kademlia should do when
// a connection fails for a node on the network
func (rt *RoutingTable) ConnectionFailed(id string, address proto.NodeAddress) {
	// TODO: What should we do ?
	return
}

// SetBucketTimestamp updates the last updated time for a bucket
func (rt *RoutingTable) SetBucketTimestamp(id string, now time.Time) error {
	rt.mutex.Lock()
	defer rt.mutex.Unlock()
	err := rt.createOrUpdateKBucket([]byte(id), now)
	if err != nil {
		return NodeErr.New("could not update bucket timestamp %s", err)
	}
	return nil
}

// GetBucketTimestamp retrieves the last updated time for a bucket
func (rt *RoutingTable) GetBucketTimestamp(id string, bucket dht.Bucket) (time.Time, error) {
	t, err := rt.kadBucketDB.Get([]byte(id))
	if err != nil {
		return time.Now(), RoutingErr.New("could not get bucket timestamp %s", err)
	}

	timestamp, _ := binary.Varint(t)

	return time.Unix(0, timestamp).UTC(), nil
}

// GetNodeRoutingTable gets a routing table for a given node rather than the local node's routing table
func GetNodeRoutingTable(ctx context.Context, ID NodeID) (RoutingTable, error) {
	//TODO: What should we do?
	return RoutingTable{}, nil
}

// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information

package kademlia

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/zeebo/errs"

	"storj.io/storj/pkg/pb"
	"storj.io/storj/storage"
)

func tempdir(t testing.TB) (dir string, cleanup func()) {
	dir, err := ioutil.TempDir("", "storj-kademlia")
	if err != nil {
		t.Fatal(err)
	}
	return dir, func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatal(err)
		}
	}
}

func createRoutingTable(t *testing.T, localNodeID []byte) (*RoutingTable, func()) {
	tempdir, cleanup := tempdir(t)

	if localNodeID == nil {
		localNodeID = []byte("AA")
	}
	localNode := &pb.Node{Id: string(localNodeID)}
	options := &RoutingOptions{
		kpath:        filepath.Join(tempdir, "Kadbucket"),
		npath:        filepath.Join(tempdir, "Nodebucket"),
		idLength:     16,
		bucketSize:   6,
		rcBucketSize: 2,
	}
	rt, err := NewRoutingTable(localNode, options)
	if err != nil {
		t.Fatal(err)
	}
	return rt, func() {
		err := rt.Close()
		cleanup()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func mockNode(id string) *pb.Node {
	var node pb.Node
	node.Id = id
	return &node
}

func TestAddNode(t *testing.T) {
	rt, cleanup := createRoutingTable(t, []byte("OO")) //localNode [79, 79] or [01001111, 01001111]
	defer cleanup()
	bucket, err := rt.kadBucketDB.Get(storage.Key([]byte{255, 255}))
	assert.NoError(t, err)
	assert.NotNil(t, bucket)
	var ok bool
	//add node to unfilled kbucket
	node1 := mockNode("PO") //[80, 79] or [01010000, 01001111]
	ok, err = rt.addNode(node1)
	assert.True(t, ok)
	assert.NoError(t, err)
	kadKeys, err := rt.kadBucketDB.List(nil, 0)
	assert.NoError(t, err)
	nodeKeys, err := rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(kadKeys))
	assert.Equal(t, 2, len(nodeKeys))

	//add node to full kbucket and split
	node2 := mockNode("NO") //[78, 79] or [01001110, 01001111]
	ok, err = rt.addNode(node2)
	assert.True(t, ok)
	assert.NoError(t, err)

	node3 := mockNode("MO") //[77, 79] or [01001101, 01001111]
	ok, err = rt.addNode(node3)
	assert.True(t, ok)
	assert.NoError(t, err)

	node4 := mockNode("LO") //[76, 79] or [01001100, 01001111]
	ok, err = rt.addNode(node4)
	assert.True(t, ok)
	assert.NoError(t, err)

	node5 := mockNode("QO") //[81, 79] or [01010001, 01001111]
	ok, err = rt.addNode(node5)
	assert.True(t, ok)
	assert.NoError(t, err)

	kadKeys, err = rt.kadBucketDB.List(nil, 0)
	assert.NoError(t, err)
	nodeKeys, err = rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(kadKeys))
	assert.Equal(t, 6, len(nodeKeys))

	//splitting here
	node6 := mockNode("SO")
	ok, err = rt.addNode(node6)
	assert.True(t, ok)
	assert.NoError(t, err)

	kadKeys, err = rt.kadBucketDB.List(nil, 0)
	assert.NoError(t, err)
	nodeKeys, err = rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(kadKeys))
	assert.Equal(t, 7, len(nodeKeys))

	// //check how many keys in each bucket
	a, err := rt.getNodeIDsWithinKBucket(kadKeys[0])
	assert.NoError(t, err)
	assert.Equal(t, 0, len(a))
	b, err := rt.getNodeIDsWithinKBucket(kadKeys[1])
	assert.NoError(t, err)
	assert.Equal(t, 4, len(b))
	c, err := rt.getNodeIDsWithinKBucket(kadKeys[2])
	assert.NoError(t, err)
	assert.Equal(t, 3, len(c))
	d, err := rt.getNodeIDsWithinKBucket(kadKeys[3])
	assert.NoError(t, err)
	assert.Equal(t, 0, len(d))
	e, err := rt.getNodeIDsWithinKBucket(kadKeys[4])
	assert.NoError(t, err)
	assert.Equal(t, 0, len(e))

	//add node to full kbucket and drop
	node7 := mockNode("?O")
	ok, err = rt.addNode(node7)
	assert.True(t, ok)
	assert.NoError(t, err)

	node8 := mockNode(">O")
	ok, err = rt.addNode(node8)
	assert.True(t, ok)
	assert.NoError(t, err)

	node9 := mockNode("=O")
	ok, err = rt.addNode(node9)
	assert.True(t, ok)
	assert.NoError(t, err)

	node10 := mockNode(";O")
	ok, err = rt.addNode(node10)
	assert.True(t, ok)
	assert.NoError(t, err)

	node11 := mockNode(":O")
	ok, err = rt.addNode(node11)
	assert.True(t, ok)
	assert.NoError(t, err)

	node12 := mockNode("9O")
	ok, err = rt.addNode(node12)
	assert.True(t, ok)
	assert.NoError(t, err)

	kadKeys, err = rt.kadBucketDB.List(nil, 0)
	assert.NoError(t, err)
	nodeKeys, err = rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(kadKeys))
	assert.Equal(t, 13, len(nodeKeys))

	a, err = rt.getNodeIDsWithinKBucket(kadKeys[0])
	assert.NoError(t, err)
	assert.Equal(t, 6, len(a))

	//should drop
	node13 := mockNode("8O")
	ok, err = rt.addNode(node13)
	assert.False(t, ok)
	assert.NoError(t, err)
	//check if node13 it into the replacement cache
	ns := rt.replacementCache[string([]byte{63, 255})]
	assert.Equal(t, node13.Id, ns[0].Id)

	kadKeys, err = rt.kadBucketDB.List(nil, 0)
	assert.NoError(t, err)
	nodeKeys, err = rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(kadKeys))
	assert.Equal(t, 13, len(nodeKeys))
	a, err = rt.getNodeIDsWithinKBucket(kadKeys[0])
	assert.NoError(t, err)
	assert.Equal(t, 6, len(a))

	//add node to highly unbalanced tree
	//adding to bucket 1
	node14 := mockNode("KO") //75
	ok, err = rt.addNode(node14)
	assert.True(t, ok)
	assert.NoError(t, err)
	node15 := mockNode("JO") //74
	ok, err = rt.addNode(node15)
	assert.True(t, ok)
	assert.NoError(t, err)

	//adding to bucket 2
	node16 := mockNode("]O") //93
	ok, err = rt.addNode(node16)
	assert.True(t, ok)
	assert.NoError(t, err)
	node17 := mockNode("^O") //94
	ok, err = rt.addNode(node17)
	assert.True(t, ok)
	assert.NoError(t, err)
	node18 := mockNode("_O") //95
	ok, err = rt.addNode(node18)
	assert.True(t, ok)
	assert.NoError(t, err)

	b, err = rt.getNodeIDsWithinKBucket(kadKeys[1])
	assert.NoError(t, err)
	assert.Equal(t, 6, len(b))
	c, err = rt.getNodeIDsWithinKBucket(kadKeys[2])
	assert.NoError(t, err)
	assert.Equal(t, 6, len(c))
	kadKeys, err = rt.kadBucketDB.List(nil, 0)
	assert.NoError(t, err)
	nodeKeys, err = rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(kadKeys))
	assert.Equal(t, 18, len(nodeKeys))

	//split bucket 2
	node19 := mockNode("@O")
	ok, err = rt.addNode(node19)
	assert.True(t, ok)
	assert.NoError(t, err)
	kadKeys, err = rt.kadBucketDB.List(nil, 0)
	assert.NoError(t, err)
	nodeKeys, err = rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)

	assert.Equal(t, 6, len(kadKeys))
	assert.Equal(t, 19, len(nodeKeys))
}

func TestUpdateNode(t *testing.T) {
	rt, cleanup := createRoutingTable(t, []byte("AA"))
	defer cleanup()
	node := mockNode("BB")
	ok, err := rt.addNode(node)
	assert.True(t, ok)
	assert.NoError(t, err)
	val, err := rt.nodeBucketDB.Get(storage.Key(node.Id))
	assert.NoError(t, err)
	unmarshaled, err := unmarshalNodes(storage.Keys{storage.Key(node.Id)}, []storage.Value{val})
	assert.NoError(t, err)
	x := unmarshaled[0].Address
	assert.Nil(t, x)

	node.Address = &pb.NodeAddress{Address: "BB"}
	err = rt.updateNode(node)
	assert.NoError(t, err)
	val, err = rt.nodeBucketDB.Get(storage.Key(node.Id))
	assert.NoError(t, err)
	unmarshaled, err = unmarshalNodes(storage.Keys{storage.Key(node.Id)}, []storage.Value{val})
	assert.NoError(t, err)
	y := unmarshaled[0].Address.Address
	assert.Equal(t, "BB", y)
}

func TestRemoveNode(t *testing.T) {
	rt, cleanup := createRoutingTable(t, []byte("AA"))
	defer cleanup()
	kadBucketID := []byte{255, 255}
	node := mockNode("BB")
	ok, err := rt.addNode(node)
	assert.True(t, ok)
	assert.NoError(t, err)
	val, err := rt.nodeBucketDB.Get(storage.Key(node.Id))
	assert.NoError(t, err)
	assert.NotNil(t, val)
	node2 := mockNode("CC")
	rt.addToReplacementCache(kadBucketID, node2)
	err = rt.removeNode(kadBucketID, storage.Key(node.Id))
	assert.NoError(t, err)
	val, err = rt.nodeBucketDB.Get(storage.Key(node.Id))
	assert.Nil(t, val)
	assert.Error(t, err)
	val2, err := rt.nodeBucketDB.Get(storage.Key(node2.Id))
	assert.NoError(t, err)
	assert.NotNil(t, val2)
	assert.Equal(t, 0, len(rt.replacementCache[string(kadBucketID)]))

	//try to remove node not in rt
	err = rt.removeNode(kadBucketID, storage.Key("DD"))
	assert.NoError(t, err)
}

func TestCreateOrUpdateKBucket(t *testing.T) {
	id := []byte{255, 255}
	rt, cleanup := createRoutingTable(t, nil)
	defer cleanup()
	err := rt.createOrUpdateKBucket(storage.Key(id), time.Now())
	assert.NoError(t, err)
	val, e := rt.kadBucketDB.Get(storage.Key(id))
	assert.NotNil(t, val)
	assert.NoError(t, e)

}

func TestGetKBucketID(t *testing.T) {
	kadIDA := storage.Key([]byte{255, 255})
	nodeIDA := []byte("AA")
	rt, cleanup := createRoutingTable(t, nodeIDA)
	defer cleanup()
	keyA, err := rt.getKBucketID(nodeIDA)
	assert.NoError(t, err)
	assert.Equal(t, kadIDA, keyA)
}

func TestXorTwoIds(t *testing.T) {
	x := xorTwoIds([]byte{191}, []byte{159})
	assert.Equal(t, []byte{32}, x) //00100000
}

func TestSortByXOR(t *testing.T) {
	node1 := []byte{127, 255} //xor 0
	rt, cleanup := createRoutingTable(t, node1)
	defer cleanup()
	node2 := []byte{143, 255} //xor 240
	assert.NoError(t, rt.nodeBucketDB.Put(node2, []byte("")))
	node3 := []byte{255, 255} //xor 128
	assert.NoError(t, rt.nodeBucketDB.Put(node3, []byte("")))
	node4 := []byte{191, 255} //xor 192
	assert.NoError(t, rt.nodeBucketDB.Put(node4, []byte("")))
	node5 := []byte{133, 255} //xor 250
	assert.NoError(t, rt.nodeBucketDB.Put(node5, []byte("")))
	nodes, err := rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	expectedNodes := storage.Keys{node1, node5, node2, node4, node3}
	assert.Equal(t, expectedNodes, nodes)
	sortedNodes := sortByXOR(nodes, node1)
	expectedSorted := storage.Keys{node1, node3, node4, node2, node5}
	assert.Equal(t, expectedSorted, sortedNodes)
	nodes, err = rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	assert.Equal(t, expectedNodes, nodes)
}

func TestDetermineFurthestIDWithinK(t *testing.T) {
	node1 := []byte{127, 255} //xor 0
	rt, cleanup := createRoutingTable(t, node1)
	defer cleanup()
	rt.self.Id = string(node1)
	assert.NoError(t, rt.nodeBucketDB.Put(node1, []byte("")))
	expectedFurthest := node1
	nodes, err := rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	furthest, err := rt.determineFurthestIDWithinK(nodes)
	assert.NoError(t, err)
	assert.Equal(t, expectedFurthest, furthest)

	node2 := []byte{143, 255} //xor 240
	assert.NoError(t, rt.nodeBucketDB.Put(node2, []byte("")))
	expectedFurthest = node2
	nodes, err = rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	furthest, err = rt.determineFurthestIDWithinK(nodes)
	assert.NoError(t, err)
	assert.Equal(t, expectedFurthest, furthest)

	node3 := []byte{255, 255} //xor 128
	assert.NoError(t, rt.nodeBucketDB.Put(node3, []byte("")))
	expectedFurthest = node2
	nodes, err = rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	furthest, err = rt.determineFurthestIDWithinK(nodes)
	assert.NoError(t, err)
	assert.Equal(t, expectedFurthest, furthest)

	node4 := []byte{191, 255} //xor 192
	assert.NoError(t, rt.nodeBucketDB.Put(node4, []byte("")))
	expectedFurthest = node2
	nodes, err = rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	furthest, err = rt.determineFurthestIDWithinK(nodes)
	assert.NoError(t, err)
	assert.Equal(t, expectedFurthest, furthest)

	node5 := []byte{133, 255} //xor 250
	assert.NoError(t, rt.nodeBucketDB.Put(node5, []byte("")))
	expectedFurthest = node5
	nodes, err = rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	furthest, err = rt.determineFurthestIDWithinK(nodes)
	assert.NoError(t, err)
	assert.Equal(t, expectedFurthest, furthest)
}

func TestNodeIsWithinNearestK(t *testing.T) {
	selfNode := []byte{127, 255}
	rt, cleanup := createRoutingTable(t, selfNode)
	defer cleanup()
	rt.bucketSize = 2
	expectTrue, err := rt.nodeIsWithinNearestK(selfNode)
	assert.NoError(t, err)
	assert.True(t, expectTrue)

	furthestNode := []byte{143, 255}
	expectTrue, err = rt.nodeIsWithinNearestK(furthestNode)
	assert.NoError(t, err)
	assert.True(t, expectTrue)
	assert.NoError(t, rt.nodeBucketDB.Put(furthestNode, []byte("")))

	node1 := []byte{255, 255}
	expectTrue, err = rt.nodeIsWithinNearestK(node1)
	assert.NoError(t, err)
	assert.True(t, expectTrue)
	assert.NoError(t, rt.nodeBucketDB.Put(node1, []byte("")))

	node2 := []byte{191, 255}
	expectTrue, err = rt.nodeIsWithinNearestK(node2)
	assert.NoError(t, err)
	assert.True(t, expectTrue)
	assert.NoError(t, rt.nodeBucketDB.Put(node1, []byte("")))

	node3 := []byte{133, 255}
	expectFalse, err := rt.nodeIsWithinNearestK(node3)
	assert.NoError(t, err)
	assert.False(t, expectFalse)
}

func TestKadBucketContainsLocalNode(t *testing.T) {
	nodeIDA := []byte{183, 255} //[10110111, 1111111]
	rt, cleanup := createRoutingTable(t, nodeIDA)
	defer cleanup()
	kadIDA := storage.Key([]byte{255, 255})
	kadIDB := storage.Key([]byte{127, 255})
	now := time.Now()
	err := rt.createOrUpdateKBucket(kadIDB, now)
	assert.NoError(t, err)
	resultTrue, err := rt.kadBucketContainsLocalNode(kadIDA)
	assert.NoError(t, err)
	resultFalse, err := rt.kadBucketContainsLocalNode(kadIDB)
	assert.NoError(t, err)
	assert.True(t, resultTrue)
	assert.False(t, resultFalse)
}

func TestKadBucketHasRoom(t *testing.T) {
	node1 := []byte{255, 255}
	kadIDA := storage.Key([]byte{255, 255})
	rt, cleanup := createRoutingTable(t, node1)
	defer cleanup()
	node2 := []byte{191, 255}
	node3 := []byte{127, 255}
	node4 := []byte{63, 255}
	node5 := []byte{159, 255}
	node6 := []byte{0, 127}
	resultA, err := rt.kadBucketHasRoom(kadIDA)
	assert.NoError(t, err)
	assert.True(t, resultA)
	assert.NoError(t, rt.nodeBucketDB.Put(node2, []byte("")))
	assert.NoError(t, rt.nodeBucketDB.Put(node3, []byte("")))
	assert.NoError(t, rt.nodeBucketDB.Put(node4, []byte("")))
	assert.NoError(t, rt.nodeBucketDB.Put(node5, []byte("")))
	assert.NoError(t, rt.nodeBucketDB.Put(node6, []byte("")))
	resultB, err := rt.kadBucketHasRoom(kadIDA)
	assert.NoError(t, err)
	assert.False(t, resultB)
}

func TestGetNodeIDsWithinKBucket(t *testing.T) {
	nodeIDA := []byte{183, 255} //[10110111, 1111111]
	rt, cleanup := createRoutingTable(t, nodeIDA)
	defer cleanup()
	kadIDA := storage.Key([]byte{255, 255})
	kadIDB := storage.Key([]byte{127, 255})
	now := time.Now()
	assert.NoError(t, rt.createOrUpdateKBucket(kadIDB, now))

	nodeIDB := []byte{111, 255} //[01101111, 1111111]
	nodeIDC := []byte{47, 255}  //[00101111, 1111111]

	assert.NoError(t, rt.nodeBucketDB.Put(nodeIDB, []byte("")))
	assert.NoError(t, rt.nodeBucketDB.Put(nodeIDC, []byte("")))

	expectedA := storage.Keys{nodeIDA}
	expectedB := storage.Keys{nodeIDC, nodeIDB}

	A, err := rt.getNodeIDsWithinKBucket(kadIDA)
	assert.NoError(t, err)
	B, err := rt.getNodeIDsWithinKBucket(kadIDB)
	assert.NoError(t, err)

	assert.Equal(t, expectedA, A)
	assert.Equal(t, expectedB, B)
}

func TestGetNodesFromIDs(t *testing.T) {
	nodeA := mockNode("AA")
	nodeB := mockNode("BB")
	nodeC := mockNode("CC")
	nodeIDA := []byte(nodeA.Id)
	nodeIDB := []byte(nodeB.Id)
	nodeIDC := []byte(nodeC.Id)
	a, err := proto.Marshal(nodeA)
	assert.NoError(t, err)
	b, err := proto.Marshal(nodeB)
	assert.NoError(t, err)
	c, err := proto.Marshal(nodeC)
	assert.NoError(t, err)
	rt, cleanup := createRoutingTable(t, nodeIDA)
	defer cleanup()

	assert.NoError(t, rt.nodeBucketDB.Put(nodeIDA, a))
	assert.NoError(t, rt.nodeBucketDB.Put(nodeIDB, b))
	assert.NoError(t, rt.nodeBucketDB.Put(nodeIDC, c))
	expected := []storage.Value{a, b, c}

	nodeIDs, err := rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	_, values, err := rt.getNodesFromIDs(nodeIDs)
	assert.NoError(t, err)
	assert.Equal(t, expected, values)
}

func TestUnmarshalNodes(t *testing.T) {
	nodeA := mockNode("AA")
	nodeB := mockNode("BB")
	nodeC := mockNode("CC")

	nodeIDA := []byte(nodeA.Id)
	nodeIDB := []byte(nodeB.Id)
	nodeIDC := []byte(nodeC.Id)
	a, err := proto.Marshal(nodeA)
	assert.NoError(t, err)
	b, err := proto.Marshal(nodeB)
	assert.NoError(t, err)
	c, err := proto.Marshal(nodeC)
	assert.NoError(t, err)
	rt, cleanup := createRoutingTable(t, nodeIDA)
	defer cleanup()
	assert.NoError(t, rt.nodeBucketDB.Put(nodeIDA, a))
	assert.NoError(t, rt.nodeBucketDB.Put(nodeIDB, b))
	assert.NoError(t, rt.nodeBucketDB.Put(nodeIDC, c))
	nodeIDs, err := rt.nodeBucketDB.List(nil, 0)
	assert.NoError(t, err)
	ids, values, err := rt.getNodesFromIDs(nodeIDs)
	assert.NoError(t, err)
	nodes, err := unmarshalNodes(ids, values)
	assert.NoError(t, err)
	expected := []*pb.Node{nodeA, nodeB, nodeC}
	for i, v := range expected {
		assert.True(t, proto.Equal(v, nodes[i]))
	}
}

func TestGetUnmarshaledNodesFromBucket(t *testing.T) {
	bucketID := []byte{255, 255}
	nodeA := mockNode("AA")
	rt, cleanup := createRoutingTable(t, []byte(nodeA.Id))
	defer cleanup()
	nodeB := mockNode("BB")
	nodeC := mockNode("CC")
	var err error
	_, err = rt.addNode(nodeB)
	assert.NoError(t, err)
	_, err = rt.addNode(nodeC)
	assert.NoError(t, err)
	nodes, err := rt.getUnmarshaledNodesFromBucket(bucketID)
	expected := []*pb.Node{nodeA, nodeB, nodeC}
	assert.NoError(t, err)
	for i, v := range expected {
		assert.True(t, proto.Equal(v, nodes[i]))
	}
}

func TestGetKBucketRange(t *testing.T) {
	rt, cleanup := createRoutingTable(t, nil)
	defer cleanup()
	idA := []byte{255, 255}
	idB := []byte{127, 255}
	idC := []byte{63, 255}
	assert.NoError(t, rt.kadBucketDB.Put(idA, []byte("")))
	assert.NoError(t, rt.kadBucketDB.Put(idB, []byte("")))
	assert.NoError(t, rt.kadBucketDB.Put(idC, []byte("")))
	cases := []struct {
		testID   string
		id       []byte
		expected storage.Keys
	}{
		{testID: "A",
			id:       idA,
			expected: storage.Keys{idB, idA},
		},
		{testID: "B",
			id:       idB,
			expected: storage.Keys{idC, idB}},
		{testID: "C",
			id:       idC,
			expected: storage.Keys{rt.createZeroAsStorageKey(), idC},
		},
	}
	for _, c := range cases {
		t.Run(c.testID, func(t *testing.T) {
			ep, err := rt.getKBucketRange(c.id)
			assert.NoError(t, err)
			assert.Equal(t, c.expected, ep)
		})
	}
}

func TestCreateFirstBucketID(t *testing.T) {
	rt, cleanup := createRoutingTable(t, nil)
	defer cleanup()
	x := rt.createFirstBucketID()
	expected := []byte{255, 255}
	assert.Equal(t, x, expected)
}

func TestCreateZeroAsStorageKey(t *testing.T) {
	rt, cleanup := createRoutingTable(t, nil)
	defer cleanup()
	zero := rt.createZeroAsStorageKey()
	expected := []byte{0, 0}
	assert.Equal(t, zero, storage.Key(expected))
}

func TestDetermineLeafDepth(t *testing.T) {
	rt, cleanup := createRoutingTable(t, nil)
	defer cleanup()
	idA := []byte{255, 255}
	idB := []byte{127, 255}
	idC := []byte{63, 255}

	cases := []struct {
		testID  string
		id      []byte
		depth   int
		addNode func()
	}{
		{testID: "A",
			id:    idA,
			depth: 0,
			addNode: func() {
				e := rt.kadBucketDB.Put(idA, []byte(""))
				assert.NoError(t, e)
			},
		},
		{testID: "B",
			id:    idB,
			depth: 1,
			addNode: func() {
				e := rt.kadBucketDB.Put(idB, []byte(""))
				assert.NoError(t, e)
			},
		},
		{testID: "C",
			id:    idA,
			depth: 1,
			addNode: func() {
				e := rt.kadBucketDB.Put(idC, []byte(""))
				assert.NoError(t, e)
			},
		},
		{testID: "D",
			id:      idB,
			depth:   2,
			addNode: func() {},
		},
		{testID: "E",
			id:      idC,
			depth:   2,
			addNode: func() {},
		},
	}
	for _, c := range cases {
		t.Run(c.testID, func(t *testing.T) {
			c.addNode()
			d, err := rt.determineLeafDepth(c.id)
			assert.NoError(t, err)
			assert.Equal(t, c.depth, d)
		})
	}
}

func TestDetermineDifferingBitIndex(t *testing.T) {
	rt, cleanup := createRoutingTable(t, nil)
	defer cleanup()
	cases := []struct {
		testID   string
		bucketID []byte
		key      []byte
		expected int
		err      *errs.Class
	}{
		{testID: "A",
			bucketID: []byte{191, 255},
			key:      []byte{255, 255},
			expected: 1,
			err:      nil,
		},
		{testID: "B",
			bucketID: []byte{255, 255},
			key:      []byte{191, 255},
			expected: 1,
			err:      nil,
		},
		{testID: "C",
			bucketID: []byte{95, 255},
			key:      []byte{127, 255},
			expected: 2,
			err:      nil,
		},
		{testID: "D",
			bucketID: []byte{95, 255},
			key:      []byte{79, 255},
			expected: 3,
			err:      nil,
		},
		{testID: "E",
			bucketID: []byte{95, 255},
			key:      []byte{63, 255},
			expected: 2,
			err:      nil,
		},
		{testID: "F",
			bucketID: []byte{95, 255},
			key:      []byte{79, 255},
			expected: 3,
			err:      nil,
		},
		{testID: "G",
			bucketID: []byte{255, 255},
			key:      []byte{255, 255},
			expected: -2,
			err:      &RoutingErr,
		},
		{testID: "H",
			bucketID: []byte{255, 255},
			key:      []byte{0, 0},
			expected: -1,
			err:      nil,
		},
		{testID: "I",
			bucketID: []byte{127, 255},
			key:      []byte{0, 0},
			expected: 0,
			err:      nil,
		},
		{testID: "J",
			bucketID: []byte{63, 255},
			key:      []byte{0, 0},
			expected: 1,
			err:      nil,
		},
		{testID: "K",
			bucketID: []byte{31, 255},
			key:      []byte{0, 0},
			expected: 2,
			err:      nil,
		},
		{testID: "L",
			bucketID: []byte{95, 255},
			key:      []byte{63, 255},
			expected: 2,
			err:      nil,
		},
	}

	for _, c := range cases {
		t.Run(c.testID, func(t *testing.T) {
			diff, err := rt.determineDifferingBitIndex(c.bucketID, c.key)
			assertErrClass(t, c.err, err)
			assert.Equal(t, c.expected, diff)
		})
	}
}

func TestSplitBucket(t *testing.T) {
	rt, cleanup := createRoutingTable(t, nil)
	defer cleanup()
	cases := []struct {
		testID string
		idA    []byte
		idB    []byte
		depth  int
	}{
		{testID: "A: [11111111, 11111111] -> [10111111, 11111111]",
			idA:   []byte{255, 255},
			idB:   []byte{191, 255},
			depth: 1,
		},
		{testID: "B: [10111111, 11111111] -> [10011111, 11111111]",
			idA:   []byte{191, 255},
			idB:   []byte{159, 255},
			depth: 2,
		},
		{testID: "C: [01111111, 11111111] -> [00111111, 11111111]",
			idA:   []byte{127, 255},
			idB:   []byte{63, 255},
			depth: 1,
		},
		{testID: "D: [00000000, 11111111] -> [00000000, 01111111]",
			idA:   []byte{0, 255},
			idB:   []byte{0, 127},
			depth: 8,
		},
		{testID: "E: [01011111, 11111111] -> [01010111, 11111111]",
			idA:   []byte{95, 255},
			idB:   []byte{87, 255},
			depth: 4,
		},
		{testID: "F: [01011111, 11111111] -> [01001111, 11111111]",
			idA:   []byte{95, 255},
			idB:   []byte{79, 255},
			depth: 3,
		},
	}
	for _, c := range cases {
		t.Run(c.testID, func(t *testing.T) {
			newID := rt.splitBucket(c.idA, c.depth)
			assert.Equal(t, c.idB, newID)
		})
	}
}

func assertErrClass(t *testing.T, class *errs.Class, err error) {
	t.Helper()
	if class != nil {
		assert.True(t, class.Has(err))
	} else {
		assert.NoError(t, err)
	}
}

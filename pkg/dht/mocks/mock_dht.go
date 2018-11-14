// Code generated by MockGen. DO NOT EDIT.
// Source: storj.io/storj/pkg/dht (interfaces: DHT,RoutingTable)

// Package mock_dht is a generated GoMock package.
package mock_dht

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"

	dht "storj.io/storj/pkg/dht"
	pb "storj.io/storj/pkg/pb"
	"storj.io/storj/storage"
)

// MockDHT is a mock of DHT interface
type MockDHT struct {
	ctrl     *gomock.Controller
	recorder *MockDHTMockRecorder
}

// MockDHTMockRecorder is the mock recorder for MockDHT
type MockDHTMockRecorder struct {
	mock *MockDHT
}

// NewMockDHT creates a new mock instance
func NewMockDHT(ctrl *gomock.Controller) *MockDHT {
	mock := &MockDHT{ctrl: ctrl}
	mock.recorder = &MockDHTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDHT) EXPECT() *MockDHTMockRecorder {
	return m.recorder
}

// Bootstrap mocks base method
func (m *MockDHT) Bootstrap(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Bootstrap", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bootstrap indicates an expected call of Bootstrap
func (mr *MockDHTMockRecorder) Bootstrap(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrap", reflect.TypeOf((*MockDHT)(nil).Bootstrap), arg0)
}

// Disconnect mocks base method
func (m *MockDHT) Disconnect() error {
	ret := m.ctrl.Call(m, "Disconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect
func (mr *MockDHTMockRecorder) Disconnect() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockDHT)(nil).Disconnect))
}

// FindNode mocks base method
func (m *MockDHT) FindNode(arg0 context.Context, arg1 dht.NodeID) (pb.Node, error) {
	ret := m.ctrl.Call(m, "FindNode", arg0, arg1)
	ret0, _ := ret[0].(pb.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindNode indicates an expected call of FindNode
func (mr *MockDHTMockRecorder) FindNode(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindNode", reflect.TypeOf((*MockDHT)(nil).FindNode), arg0, arg1)
}

// GetNodes mocks base method
func (m *MockDHT) GetNodes(arg0 context.Context, arg1 string, arg2 int, arg3 ...pb.Restriction) ([]*pb.Node, error) {
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNodes", varargs...)
	ret0, _ := ret[0].([]*pb.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodes indicates an expected call of GetNodes
func (mr *MockDHTMockRecorder) GetNodes(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodes", reflect.TypeOf((*MockDHT)(nil).GetNodes), varargs...)
}

// GetRoutingTable mocks base method
func (m *MockDHT) GetRoutingTable(arg0 context.Context) (dht.RoutingTable, error) {
	ret := m.ctrl.Call(m, "GetRoutingTable", arg0)
	ret0, _ := ret[0].(dht.RoutingTable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoutingTable indicates an expected call of GetRoutingTable
func (mr *MockDHTMockRecorder) GetRoutingTable(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoutingTable", reflect.TypeOf((*MockDHT)(nil).GetRoutingTable), arg0)
}

// Ping mocks base method
func (m *MockDHT) Ping(arg0 context.Context, arg1 pb.Node) (pb.Node, error) {
	ret := m.ctrl.Call(m, "Ping", arg0, arg1)
	ret0, _ := ret[0].(pb.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping
func (mr *MockDHTMockRecorder) Ping(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDHT)(nil).Ping), arg0, arg1)
}

// MockRoutingTable is a mock of RoutingTable interface
type MockRoutingTable struct {
	ctrl     *gomock.Controller
	recorder *MockRoutingTableMockRecorder
}

// MockRoutingTableMockRecorder is the mock recorder for MockRoutingTable
type MockRoutingTableMockRecorder struct {
	mock *MockRoutingTable
}

// NewMockRoutingTable creates a new mock instance
func NewMockRoutingTable(ctrl *gomock.Controller) *MockRoutingTable {
	mock := &MockRoutingTable{ctrl: ctrl}
	mock.recorder = &MockRoutingTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoutingTable) EXPECT() *MockRoutingTableMockRecorder {
	return m.recorder
}

// CacheSize mocks base method
func (m *MockRoutingTable) CacheSize() int {
	ret := m.ctrl.Call(m, "CacheSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// CacheSize indicates an expected call of CacheSize
func (mr *MockRoutingTableMockRecorder) CacheSize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CacheSize", reflect.TypeOf((*MockRoutingTable)(nil).CacheSize))
}

// ConnectionFailed mocks base method
func (m *MockRoutingTable) ConnectionFailed(arg0 *pb.Node) error {
	ret := m.ctrl.Call(m, "ConnectionFailed", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConnectionFailed indicates an expected call of ConnectionFailed
func (mr *MockRoutingTableMockRecorder) ConnectionFailed(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionFailed", reflect.TypeOf((*MockRoutingTable)(nil).ConnectionFailed), arg0)
}

// ConnectionSuccess mocks base method
func (m *MockRoutingTable) ConnectionSuccess(arg0 *pb.Node) error {
	ret := m.ctrl.Call(m, "ConnectionSuccess", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConnectionSuccess indicates an expected call of ConnectionSuccess
func (mr *MockRoutingTableMockRecorder) ConnectionSuccess(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionSuccess", reflect.TypeOf((*MockRoutingTable)(nil).ConnectionSuccess), arg0)
}

// FindNear mocks base method
func (m *MockRoutingTable) FindNear(arg0 dht.NodeID, arg1 int) ([]*pb.Node, error) {
	ret := m.ctrl.Call(m, "FindNear", arg0, arg1)
	ret0, _ := ret[0].([]*pb.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindNear indicates an expected call of FindNear
func (mr *MockRoutingTableMockRecorder) FindNear(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindNear", reflect.TypeOf((*MockRoutingTable)(nil).FindNear), arg0, arg1)
}

// GetBucket mocks base method
func (m *MockRoutingTable) GetBucket(arg0 string) (dht.Bucket, bool) {
	ret := m.ctrl.Call(m, "GetBucket", arg0)
	ret0, _ := ret[0].(dht.Bucket)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetBucket indicates an expected call of GetBucket
func (mr *MockRoutingTableMockRecorder) GetBucket(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucket", reflect.TypeOf((*MockRoutingTable)(nil).GetBucket), arg0)
}

// GetBucketTimestamp mocks base method
func (m *MockRoutingTable) GetBucketTimestamp(arg0 string, arg1 dht.Bucket) (time.Time, error) {
	ret := m.ctrl.Call(m, "GetBucketTimestamp", arg0, arg1)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketTimestamp indicates an expected call of GetBucketTimestamp
func (mr *MockRoutingTableMockRecorder) GetBucketTimestamp(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketTimestamp", reflect.TypeOf((*MockRoutingTable)(nil).GetBucketTimestamp), arg0, arg1)
}

// GetBuckets mocks base method
func (m *MockRoutingTable) GetBuckets() ([]dht.Bucket, error) {
	ret := m.ctrl.Call(m, "GetBuckets")
	ret0, _ := ret[0].([]dht.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockRoutingTable) GetBucketIds() (storage.Keys, error) {
	return nil, nil
}

// GetBuckets indicates an expected call of GetBuckets
func (mr *MockRoutingTableMockRecorder) GetBuckets() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuckets", reflect.TypeOf((*MockRoutingTable)(nil).GetBuckets))
}

// K mocks base method
func (m *MockRoutingTable) K() int {
	ret := m.ctrl.Call(m, "K")
	ret0, _ := ret[0].(int)
	return ret0
}

// K indicates an expected call of K
func (mr *MockRoutingTableMockRecorder) K() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "K", reflect.TypeOf((*MockRoutingTable)(nil).K))
}

// Local mocks base method
func (m *MockRoutingTable) Local() pb.Node {
	ret := m.ctrl.Call(m, "Local")
	ret0, _ := ret[0].(pb.Node)
	return ret0
}

// Local indicates an expected call of Local
func (mr *MockRoutingTableMockRecorder) Local() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Local", reflect.TypeOf((*MockRoutingTable)(nil).Local))
}

// SetBucketTimestamp mocks base method
func (m *MockRoutingTable) SetBucketTimestamp(arg0 string, arg1 time.Time) error {
	ret := m.ctrl.Call(m, "SetBucketTimestamp", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBucketTimestamp indicates an expected call of SetBucketTimestamp
func (mr *MockRoutingTableMockRecorder) SetBucketTimestamp(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBucketTimestamp", reflect.TypeOf((*MockRoutingTable)(nil).SetBucketTimestamp), arg0, arg1)
}

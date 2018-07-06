// Code generated by MockGen. DO NOT EDIT.
// Source: storj.io/storj/protos/netstate (interfaces: NetStateClient)

// Package netstate is a generated GoMock package.
package netstate

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
	netstate "storj.io/storj/protos/netstate"
)

// MockNetStateClient is a mock of NetStateClient interface
type MockNetStateClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetStateClientMockRecorder
}

// MockNetStateClientMockRecorder is the mock recorder for MockNetStateClient
type MockNetStateClientMockRecorder struct {
	mock *MockNetStateClient
}

// NewMockNetStateClient creates a new mock instance
func NewMockNetStateClient(ctrl *gomock.Controller) *MockNetStateClient {
	mock := &MockNetStateClient{ctrl: ctrl}
	mock.recorder = &MockNetStateClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetStateClient) EXPECT() *MockNetStateClientMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockNetStateClient) Delete(arg0 context.Context, arg1 *netstate.DeleteRequest, arg2 ...grpc.CallOption) (*netstate.DeleteResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*netstate.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockNetStateClientMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNetStateClient)(nil).Delete), varargs...)
}

// Get mocks base method
func (m *MockNetStateClient) Get(arg0 context.Context, arg1 *netstate.GetRequest, arg2 ...grpc.CallOption) (*netstate.GetResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*netstate.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockNetStateClientMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNetStateClient)(nil).Get), varargs...)
}

// List mocks base method
func (m *MockNetStateClient) List(arg0 context.Context, arg1 *netstate.ListRequest, arg2 ...grpc.CallOption) (*netstate.ListResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*netstate.ListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockNetStateClientMockRecorder) List(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNetStateClient)(nil).List), varargs...)
}

// Put mocks base method
func (m *MockNetStateClient) Put(arg0 context.Context, arg1 *netstate.PutRequest, arg2 ...grpc.CallOption) (*netstate.PutResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Put", varargs...)
	ret0, _ := ret[0].(*netstate.PutResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put
func (mr *MockNetStateClientMockRecorder) Put(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockNetStateClient)(nil).Put), varargs...)
}
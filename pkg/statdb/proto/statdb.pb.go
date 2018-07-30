// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statdb.proto

package statdb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Node is info for a updating a single farmer, used in the Update rpc calls
type Node struct {
	NodeId               []byte   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	LatencyList          []int64  `protobuf:"varint,2,rep,packed,name=latency_list,json=latencyList,proto3" json:"latency_list,omitempty"`
	AuditSuccess         bool     `protobuf:"varint,3,opt,name=audit_success,json=auditSuccess,proto3" json:"audit_success,omitempty"`
	IsUp                 bool     `protobuf:"varint,4,opt,name=is_up,json=isUp,proto3" json:"is_up,omitempty"`
	UpdateLatency        bool     `protobuf:"varint,5,opt,name=update_latency,json=updateLatency,proto3" json:"update_latency,omitempty"`
	UpdateAuditSuccess   bool     `protobuf:"varint,6,opt,name=update_audit_success,json=updateAuditSuccess,proto3" json:"update_audit_success,omitempty"`
	UpdateUptime         bool     `protobuf:"varint,7,opt,name=update_uptime,json=updateUptime,proto3" json:"update_uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_c54091604b1717e8, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Node) GetLatencyList() []int64 {
	if m != nil {
		return m.LatencyList
	}
	return nil
}

func (m *Node) GetAuditSuccess() bool {
	if m != nil {
		return m.AuditSuccess
	}
	return false
}

func (m *Node) GetIsUp() bool {
	if m != nil {
		return m.IsUp
	}
	return false
}

func (m *Node) GetUpdateLatency() bool {
	if m != nil {
		return m.UpdateLatency
	}
	return false
}

func (m *Node) GetUpdateAuditSuccess() bool {
	if m != nil {
		return m.UpdateAuditSuccess
	}
	return false
}

func (m *Node) GetUpdateUptime() bool {
	if m != nil {
		return m.UpdateUptime
	}
	return false
}

// NodeStats is info about a single farmer stored in the stats db
type NodeStats struct {
	NodeId               []byte   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Latency_90           int64    `protobuf:"varint,2,opt,name=latency_90,json=latency90,proto3" json:"latency_90,omitempty"`
	AuditSuccessRatio    float64  `protobuf:"fixed64,3,opt,name=audit_success_ratio,json=auditSuccessRatio,proto3" json:"audit_success_ratio,omitempty"`
	UptimeRatio          float64  `protobuf:"fixed64,4,opt,name=uptime_ratio,json=uptimeRatio,proto3" json:"uptime_ratio,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeStats) Reset()         { *m = NodeStats{} }
func (m *NodeStats) String() string { return proto.CompactTextString(m) }
func (*NodeStats) ProtoMessage()    {}
func (*NodeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_c54091604b1717e8, []int{1}
}
func (m *NodeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeStats.Unmarshal(m, b)
}
func (m *NodeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeStats.Marshal(b, m, deterministic)
}
func (dst *NodeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeStats.Merge(dst, src)
}
func (m *NodeStats) XXX_Size() int {
	return xxx_messageInfo_NodeStats.Size(m)
}
func (m *NodeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeStats.DiscardUnknown(m)
}

var xxx_messageInfo_NodeStats proto.InternalMessageInfo

func (m *NodeStats) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *NodeStats) GetLatency_90() int64 {
	if m != nil {
		return m.Latency_90
	}
	return 0
}

func (m *NodeStats) GetAuditSuccessRatio() float64 {
	if m != nil {
		return m.AuditSuccessRatio
	}
	return 0
}

func (m *NodeStats) GetUptimeRatio() float64 {
	if m != nil {
		return m.UptimeRatio
	}
	return 0
}

// CreateRequest is a request message for the Create rpc call
type CreateRequest struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	APIKey               []byte   `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_c54091604b1717e8, []int{2}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *CreateRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// CreateResponse is a response message for the Create rpc call
type CreateResponse struct {
	Stats                *NodeStats `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_c54091604b1717e8, []int{3}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetStats() *NodeStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// GetRequest is a request message for the Get rpc call
type GetRequest struct {
	NodeId               []byte   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	APIKey               []byte   `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_c54091604b1717e8, []int{4}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *GetRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// GetResponse is a response message for the Get rpc call
type GetResponse struct {
	Stats                *NodeStats `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_c54091604b1717e8, []int{5}
}
func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (dst *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(dst, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetStats() *NodeStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// UpdateRequest is a request message for the Update rpc call
type UpdateRequest struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	APIKey               []byte   `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_c54091604b1717e8, []int{6}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *UpdateRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// UpdateRequest is a response message for the Update rpc call
type UpdateResponse struct {
	Stats                *NodeStats `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_c54091604b1717e8, []int{7}
}
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(dst, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetStats() *NodeStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// UpdateBatchRequest is a request message for the UpdateBatch rpc call
type UpdateBatchRequest struct {
	NodeList             []*Node  `protobuf:"bytes,1,rep,name=node_list,json=nodeList,proto3" json:"node_list,omitempty"`
	APIKey               []byte   `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBatchRequest) Reset()         { *m = UpdateBatchRequest{} }
func (m *UpdateBatchRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBatchRequest) ProtoMessage()    {}
func (*UpdateBatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_c54091604b1717e8, []int{8}
}
func (m *UpdateBatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBatchRequest.Unmarshal(m, b)
}
func (m *UpdateBatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBatchRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateBatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBatchRequest.Merge(dst, src)
}
func (m *UpdateBatchRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBatchRequest.Size(m)
}
func (m *UpdateBatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBatchRequest proto.InternalMessageInfo

func (m *UpdateBatchRequest) GetNodeList() []*Node {
	if m != nil {
		return m.NodeList
	}
	return nil
}

func (m *UpdateBatchRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// UpdateBatchResponse is a response message for the UpdateBatch rpc call
type UpdateBatchResponse struct {
	StatsList            []*NodeStats `protobuf:"bytes,1,rep,name=stats_list,json=statsList,proto3" json:"stats_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateBatchResponse) Reset()         { *m = UpdateBatchResponse{} }
func (m *UpdateBatchResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateBatchResponse) ProtoMessage()    {}
func (*UpdateBatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_statdb_c54091604b1717e8, []int{9}
}
func (m *UpdateBatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBatchResponse.Unmarshal(m, b)
}
func (m *UpdateBatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBatchResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateBatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBatchResponse.Merge(dst, src)
}
func (m *UpdateBatchResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateBatchResponse.Size(m)
}
func (m *UpdateBatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBatchResponse proto.InternalMessageInfo

func (m *UpdateBatchResponse) GetStatsList() []*NodeStats {
	if m != nil {
		return m.StatsList
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "statdb.Node")
	proto.RegisterType((*NodeStats)(nil), "statdb.NodeStats")
	proto.RegisterType((*CreateRequest)(nil), "statdb.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "statdb.CreateResponse")
	proto.RegisterType((*GetRequest)(nil), "statdb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "statdb.GetResponse")
	proto.RegisterType((*UpdateRequest)(nil), "statdb.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "statdb.UpdateResponse")
	proto.RegisterType((*UpdateBatchRequest)(nil), "statdb.UpdateBatchRequest")
	proto.RegisterType((*UpdateBatchResponse)(nil), "statdb.UpdateBatchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatDBClient is the client API for StatDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatDBClient interface {
	// Create a db entry for the provided farmer ID
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get uses a farmer ID to get that farmer's stats
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Update updates farmer stats for a single farmer
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// UpdateBatch updates farmer stats for multiple farmers at a time
	UpdateBatch(ctx context.Context, in *UpdateBatchRequest, opts ...grpc.CallOption) (*UpdateBatchResponse, error)
}

type statDBClient struct {
	cc *grpc.ClientConn
}

func NewStatDBClient(cc *grpc.ClientConn) StatDBClient {
	return &statDBClient{cc}
}

func (c *statDBClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/statdb.StatDB/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statDBClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/statdb.StatDB/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statDBClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/statdb.StatDB/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statDBClient) UpdateBatch(ctx context.Context, in *UpdateBatchRequest, opts ...grpc.CallOption) (*UpdateBatchResponse, error) {
	out := new(UpdateBatchResponse)
	err := c.cc.Invoke(ctx, "/statdb.StatDB/UpdateBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatDBServer is the server API for StatDB service.
type StatDBServer interface {
	// Create a db entry for the provided farmer ID
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get uses a farmer ID to get that farmer's stats
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Update updates farmer stats for a single farmer
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// UpdateBatch updates farmer stats for multiple farmers at a time
	UpdateBatch(context.Context, *UpdateBatchRequest) (*UpdateBatchResponse, error)
}

func RegisterStatDBServer(s *grpc.Server, srv StatDBServer) {
	s.RegisterService(&_StatDB_serviceDesc, srv)
}

func _StatDB_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatDBServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statdb.StatDB/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatDBServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatDB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatDBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statdb.StatDB/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatDBServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatDB_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatDBServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statdb.StatDB/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatDBServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatDB_UpdateBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatDBServer).UpdateBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statdb.StatDB/UpdateBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatDBServer).UpdateBatch(ctx, req.(*UpdateBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "statdb.StatDB",
	HandlerType: (*StatDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StatDB_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StatDB_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StatDB_Update_Handler,
		},
		{
			MethodName: "UpdateBatch",
			Handler:    _StatDB_UpdateBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "statdb.proto",
}

func init() { proto.RegisterFile("statdb.proto", fileDescriptor_statdb_c54091604b1717e8) }

var fileDescriptor_statdb_c54091604b1717e8 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x45, 0x91, 0xac, 0xc4, 0x23, 0x39, 0x90, 0x71, 0x9b, 0x0a, 0x97, 0x82, 0xaa, 0x50, 0xea,
	0x5e, 0x8c, 0x49, 0xa1, 0xc1, 0x87, 0x1e, 0x92, 0x96, 0x1a, 0xd3, 0x50, 0xca, 0x06, 0xd3, 0xa3,
	0x50, 0xac, 0x85, 0x0a, 0x52, 0x4b, 0xf5, 0x8e, 0x0e, 0xf9, 0x91, 0x7e, 0x68, 0x8f, 0x3d, 0x95,
	0x9d, 0x5d, 0x61, 0x29, 0x8d, 0x0f, 0x86, 0x1e, 0xf7, 0xbd, 0xd9, 0xf7, 0xde, 0xbe, 0xb1, 0x0c,
	0xa1, 0xa2, 0x8c, 0xf2, 0xdb, 0x49, 0xb5, 0x29, 0xa9, 0x44, 0xdf, 0x9c, 0x92, 0x3f, 0x0e, 0x78,
	0x5f, 0xca, 0x5c, 0xe2, 0x33, 0x38, 0x5c, 0x97, 0xb9, 0x4c, 0x8b, 0x3c, 0x72, 0x62, 0x67, 0x1c,
	0x0a, 0x5f, 0x1f, 0x17, 0x39, 0xbe, 0x84, 0xf0, 0x2e, 0x23, 0xb9, 0x5e, 0xdd, 0xa7, 0x77, 0x85,
	0xa2, 0xe8, 0x20, 0x76, 0xc7, 0xae, 0x08, 0x2c, 0x76, 0x5d, 0x28, 0xc2, 0x33, 0x18, 0x64, 0x75,
	0x5e, 0x50, 0xaa, 0xea, 0xd5, 0x4a, 0x2a, 0x15, 0xb9, 0xb1, 0x33, 0x3e, 0x12, 0x21, 0x83, 0x37,
	0x06, 0xc3, 0x21, 0xf4, 0x0a, 0x95, 0xd6, 0x55, 0xe4, 0x31, 0xe9, 0x15, 0x6a, 0x59, 0xe1, 0x2b,
	0x38, 0xae, 0xab, 0x3c, 0x23, 0x99, 0x5a, 0xbd, 0xa8, 0xc7, 0xec, 0xc0, 0xa0, 0xd7, 0x06, 0xc4,
	0x29, 0x3c, 0xb1, 0x63, 0x5d, 0x1f, 0x9f, 0x87, 0xd1, 0x70, 0x97, 0x6d, 0xb7, 0x33, 0xb0, 0x12,
	0x69, 0x5d, 0x51, 0xf1, 0x43, 0x46, 0x87, 0x26, 0x92, 0x01, 0x97, 0x8c, 0x25, 0xbf, 0x1c, 0xe8,
	0xeb, 0xc7, 0xdf, 0x50, 0x46, 0x6a, 0x77, 0x03, 0x2f, 0x00, 0x9a, 0x06, 0x66, 0xd3, 0xe8, 0x20,
	0x76, 0xc6, 0xae, 0xe8, 0x5b, 0x64, 0x36, 0xc5, 0x09, 0x0c, 0x3b, 0xa9, 0xd2, 0x4d, 0x46, 0x45,
	0xc9, 0x1d, 0x38, 0xe2, 0xa4, 0xdd, 0x81, 0xd0, 0x84, 0x2e, 0xd4, 0x64, 0xb2, 0x83, 0x1e, 0x0f,
	0x06, 0x06, 0xe3, 0x91, 0x64, 0x01, 0x83, 0x0f, 0x1b, 0x99, 0x91, 0x14, 0xf2, 0x67, 0x2d, 0x15,
	0x61, 0x0c, 0x9e, 0x0e, 0xc3, 0xc1, 0x82, 0xf3, 0x70, 0x62, 0x77, 0xa9, 0xc3, 0x0b, 0x66, 0xf0,
	0x14, 0xfc, 0xcb, 0xaf, 0x8b, 0xcf, 0xf2, 0x9e, 0x03, 0x86, 0xc2, 0x9e, 0x92, 0x19, 0x1c, 0x37,
	0x52, 0xaa, 0x2a, 0xd7, 0x4a, 0xe2, 0x6b, 0xe8, 0xe9, 0xeb, 0xca, 0x8a, 0x9d, 0xb4, 0xc5, 0xb8,
	0x09, 0x61, 0xf8, 0xe4, 0x3d, 0xc0, 0x5c, 0x52, 0x13, 0x61, 0x67, 0x3d, 0xbb, 0x9c, 0xdf, 0x41,
	0xc0, 0xd7, 0xf7, 0xb5, 0x5d, 0xc0, 0x60, 0xc9, 0x5b, 0xfa, 0x2f, 0x8f, 0x6f, 0xa4, 0xf6, 0x4d,
	0xf1, 0x0d, 0xd0, 0x5c, 0xbd, 0xca, 0x68, 0xf5, 0xbd, 0x89, 0xf2, 0x06, 0xfa, 0x5c, 0x02, 0x7f,
	0x09, 0x4e, 0xec, 0xfe, 0x93, 0xe7, 0x48, 0xd3, 0xfc, 0x51, 0xec, 0xca, 0x34, 0x87, 0x61, 0x47,
	0xd8, 0x06, 0x9b, 0x02, 0xb0, 0x71, 0x5b, 0xfa, 0x91, 0x74, 0x7d, 0x1e, 0xd2, 0x06, 0xe7, 0xbf,
	0x1d, 0xf0, 0x35, 0xf8, 0xf1, 0x0a, 0x2f, 0xc0, 0x37, 0x4b, 0xc6, 0xa7, 0xcd, 0x95, 0xce, 0xef,
	0x67, 0x74, 0xfa, 0x10, 0xb6, 0xae, 0x13, 0x70, 0xe7, 0x92, 0x10, 0x1b, 0x7a, 0xbb, 0xef, 0xd1,
	0xb0, 0x83, 0xd9, 0xf9, 0x0b, 0xf0, 0x4d, 0xf8, 0xad, 0x51, 0x67, 0x57, 0x5b, 0xa3, 0x07, 0xbd,
	0x7f, 0x82, 0xa0, 0xf5, 0x6a, 0x1c, 0x75, 0xc7, 0xda, 0x1d, 0x8f, 0x9e, 0x3f, 0xca, 0x19, 0x9d,
	0x5b, 0x9f, 0xff, 0xbe, 0xde, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x96, 0xbe, 0xeb, 0xce,
	0x04, 0x00, 0x00,
}

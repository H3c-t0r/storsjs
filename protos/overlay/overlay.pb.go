// Code generated by protoc-gen-go. DO NOT EDIT.
// source: overlay.proto

package overlay

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

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

// NodeTransport is an enum of possible transports for the overlay network
type NodeTransport int32

const (
	NodeTransport_TCP NodeTransport = 0
)

var NodeTransport_name = map[int32]string{
	0: "TCP",
}
var NodeTransport_value = map[string]int32{
	"TCP": 0,
}

func (x NodeTransport) String() string {
	return proto.EnumName(NodeTransport_name, int32(x))
}
func (NodeTransport) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8e97a5f3853567d6, []int{0}
}

// NodeTyp is an enum of possible node types
type NodeType int32

const (
	NodeType_ADMIN   NodeType = 0
	NodeType_STORAGE NodeType = 1
)

var NodeType_name = map[int32]string{
	0: "ADMIN",
	1: "STORAGE",
}
var NodeType_value = map[string]int32{
	"ADMIN":   0,
	"STORAGE": 1,
}

func (x NodeType) String() string {
	return proto.EnumName(NodeType_name, int32(x))
}
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8e97a5f3853567d6, []int{1}
}

// LookupRequest is is request message for the lookup rpc call
type LookupRequest struct {
	NodeID               string   `protobuf:"bytes,1,opt,name=nodeID" json:"nodeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupRequest) Reset()         { *m = LookupRequest{} }
func (m *LookupRequest) String() string { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()    {}
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8e97a5f3853567d6, []int{0}
}
func (m *LookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupRequest.Unmarshal(m, b)
}
func (m *LookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupRequest.Marshal(b, m, deterministic)
}
func (dst *LookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupRequest.Merge(dst, src)
}
func (m *LookupRequest) XXX_Size() int {
	return xxx_messageInfo_LookupRequest.Size(m)
}
func (m *LookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookupRequest proto.InternalMessageInfo

func (m *LookupRequest) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

// LookupResponse is is response message for the lookup rpc call
type LookupResponse struct {
	Node                 *Node    `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupResponse) Reset()         { *m = LookupResponse{} }
func (m *LookupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()    {}
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8e97a5f3853567d6, []int{1}
}
func (m *LookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupResponse.Unmarshal(m, b)
}
func (m *LookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupResponse.Marshal(b, m, deterministic)
}
func (dst *LookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupResponse.Merge(dst, src)
}
func (m *LookupResponse) XXX_Size() int {
	return xxx_messageInfo_LookupResponse.Size(m)
}
func (m *LookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LookupResponse proto.InternalMessageInfo

func (m *LookupResponse) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// FindStorageNodesResponse is is response message for the FindStorageNodes rpc call
type FindStorageNodesResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindStorageNodesResponse) Reset()         { *m = FindStorageNodesResponse{} }
func (m *FindStorageNodesResponse) String() string { return proto.CompactTextString(m) }
func (*FindStorageNodesResponse) ProtoMessage()    {}
func (*FindStorageNodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8e97a5f3853567d6, []int{2}
}
func (m *FindStorageNodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindStorageNodesResponse.Unmarshal(m, b)
}
func (m *FindStorageNodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindStorageNodesResponse.Marshal(b, m, deterministic)
}
func (dst *FindStorageNodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindStorageNodesResponse.Merge(dst, src)
}
func (m *FindStorageNodesResponse) XXX_Size() int {
	return xxx_messageInfo_FindStorageNodesResponse.Size(m)
}
func (m *FindStorageNodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindStorageNodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindStorageNodesResponse proto.InternalMessageInfo

func (m *FindStorageNodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// FindStorageNodesRequest is is request message for the FindStorageNodes rpc call
type FindStorageNodesRequest struct {
	ObjectSize           int64              `protobuf:"varint,1,opt,name=objectSize" json:"objectSize,omitempty"`
	ContractLength       *duration.Duration `protobuf:"bytes,2,opt,name=contractLength" json:"contractLength,omitempty"`
	Opts                 *OverlayOptions    `protobuf:"bytes,3,opt,name=opts" json:"opts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FindStorageNodesRequest) Reset()         { *m = FindStorageNodesRequest{} }
func (m *FindStorageNodesRequest) String() string { return proto.CompactTextString(m) }
func (*FindStorageNodesRequest) ProtoMessage()    {}
func (*FindStorageNodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8e97a5f3853567d6, []int{3}
}
func (m *FindStorageNodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindStorageNodesRequest.Unmarshal(m, b)
}
func (m *FindStorageNodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindStorageNodesRequest.Marshal(b, m, deterministic)
}
func (dst *FindStorageNodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindStorageNodesRequest.Merge(dst, src)
}
func (m *FindStorageNodesRequest) XXX_Size() int {
	return xxx_messageInfo_FindStorageNodesRequest.Size(m)
}
func (m *FindStorageNodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindStorageNodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindStorageNodesRequest proto.InternalMessageInfo

func (m *FindStorageNodesRequest) GetObjectSize() int64 {
	if m != nil {
		return m.ObjectSize
	}
	return 0
}

func (m *FindStorageNodesRequest) GetContractLength() *duration.Duration {
	if m != nil {
		return m.ContractLength
	}
	return nil
}

func (m *FindStorageNodesRequest) GetOpts() *OverlayOptions {
	if m != nil {
		return m.Opts
	}
	return nil
}

// NodeAddress contains the information needed to communicate with a node on the network
type NodeAddress struct {
	Transport            NodeTransport `protobuf:"varint,1,opt,name=transport,enum=NodeTransport" json:"transport,omitempty"`
	Address              string        `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NodeAddress) Reset()         { *m = NodeAddress{} }
func (m *NodeAddress) String() string { return proto.CompactTextString(m) }
func (*NodeAddress) ProtoMessage()    {}
func (*NodeAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8e97a5f3853567d6, []int{4}
}
func (m *NodeAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAddress.Unmarshal(m, b)
}
func (m *NodeAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAddress.Marshal(b, m, deterministic)
}
func (dst *NodeAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAddress.Merge(dst, src)
}
func (m *NodeAddress) XXX_Size() int {
	return xxx_messageInfo_NodeAddress.Size(m)
}
func (m *NodeAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAddress proto.InternalMessageInfo

func (m *NodeAddress) GetTransport() NodeTransport {
	if m != nil {
		return m.Transport
	}
	return NodeTransport_TCP
}

func (m *NodeAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// OverlayOptions is a set of criteria that a node must meet to be considered for a storage opportunity
type OverlayOptions struct {
	MaxLatency           *duration.Duration `protobuf:"bytes,1,opt,name=maxLatency" json:"maxLatency,omitempty"`
	MinReputation        *NodeRep           `protobuf:"bytes,2,opt,name=minReputation" json:"minReputation,omitempty"`
	MinSpeedKbps         int64              `protobuf:"varint,3,opt,name=minSpeedKbps" json:"minSpeedKbps,omitempty"`
	Limit                int64              `protobuf:"varint,4,opt,name=limit" json:"limit,omitempty"`
	Restictions          *NodeRestrictions  `protobuf:"bytes,5,opt,name=restictions" json:"restictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OverlayOptions) Reset()         { *m = OverlayOptions{} }
func (m *OverlayOptions) String() string { return proto.CompactTextString(m) }
func (*OverlayOptions) ProtoMessage()    {}
func (*OverlayOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8e97a5f3853567d6, []int{5}
}
func (m *OverlayOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OverlayOptions.Unmarshal(m, b)
}
func (m *OverlayOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OverlayOptions.Marshal(b, m, deterministic)
}
func (dst *OverlayOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OverlayOptions.Merge(dst, src)
}
func (m *OverlayOptions) XXX_Size() int {
	return xxx_messageInfo_OverlayOptions.Size(m)
}
func (m *OverlayOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_OverlayOptions.DiscardUnknown(m)
}

var xxx_messageInfo_OverlayOptions proto.InternalMessageInfo

func (m *OverlayOptions) GetMaxLatency() *duration.Duration {
	if m != nil {
		return m.MaxLatency
	}
	return nil
}

func (m *OverlayOptions) GetMinReputation() *NodeRep {
	if m != nil {
		return m.MinReputation
	}
	return nil
}

func (m *OverlayOptions) GetMinSpeedKbps() int64 {
	if m != nil {
		return m.MinSpeedKbps
	}
	return 0
}

func (m *OverlayOptions) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *OverlayOptions) GetRestictions() *NodeRestrictions {
	if m != nil {
		return m.Restictions
	}
	return nil
}

// NodeRep is the reputation characteristics of a node
type NodeRep struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRep) Reset()         { *m = NodeRep{} }
func (m *NodeRep) String() string { return proto.CompactTextString(m) }
func (*NodeRep) ProtoMessage()    {}
func (*NodeRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8e97a5f3853567d6, []int{6}
}
func (m *NodeRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRep.Unmarshal(m, b)
}
func (m *NodeRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRep.Marshal(b, m, deterministic)
}
func (dst *NodeRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRep.Merge(dst, src)
}
func (m *NodeRep) XXX_Size() int {
	return xxx_messageInfo_NodeRep.Size(m)
}
func (m *NodeRep) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRep.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRep proto.InternalMessageInfo

//  NodeRestrictions contains all relevant data about a nodes ability to store data
type NodeRestrictions struct {
	FreeBandwidth        int64    `protobuf:"varint,1,opt,name=freeBandwidth" json:"freeBandwidth,omitempty"`
	FreeDisk             int64    `protobuf:"varint,2,opt,name=freeDisk" json:"freeDisk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeRestrictions) Reset()         { *m = NodeRestrictions{} }
func (m *NodeRestrictions) String() string { return proto.CompactTextString(m) }
func (*NodeRestrictions) ProtoMessage()    {}
func (*NodeRestrictions) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8e97a5f3853567d6, []int{7}
}
func (m *NodeRestrictions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeRestrictions.Unmarshal(m, b)
}
func (m *NodeRestrictions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeRestrictions.Marshal(b, m, deterministic)
}
func (dst *NodeRestrictions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeRestrictions.Merge(dst, src)
}
func (m *NodeRestrictions) XXX_Size() int {
	return xxx_messageInfo_NodeRestrictions.Size(m)
}
func (m *NodeRestrictions) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeRestrictions.DiscardUnknown(m)
}

var xxx_messageInfo_NodeRestrictions proto.InternalMessageInfo

func (m *NodeRestrictions) GetFreeBandwidth() int64 {
	if m != nil {
		return m.FreeBandwidth
	}
	return 0
}

func (m *NodeRestrictions) GetFreeDisk() int64 {
	if m != nil {
		return m.FreeDisk
	}
	return 0
}

// Node represents a node in the overlay network
type Node struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address              *NodeAddress      `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	Type                 NodeType          `protobuf:"varint,3,opt,name=type,enum=NodeType" json:"type,omitempty"`
	Restrictions         *NodeRestrictions `protobuf:"bytes,4,opt,name=restrictions" json:"restrictions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_overlay_8e97a5f3853567d6, []int{8}
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

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() *NodeAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Node) GetType() NodeType {
	if m != nil {
		return m.Type
	}
	return NodeType_ADMIN
}

func (m *Node) GetRestrictions() *NodeRestrictions {
	if m != nil {
		return m.Restrictions
	}
	return nil
}

func init() {
	proto.RegisterType((*LookupRequest)(nil), "LookupRequest")
	proto.RegisterType((*LookupResponse)(nil), "LookupResponse")
	proto.RegisterType((*FindStorageNodesResponse)(nil), "FindStorageNodesResponse")
	proto.RegisterType((*FindStorageNodesRequest)(nil), "FindStorageNodesRequest")
	proto.RegisterType((*NodeAddress)(nil), "NodeAddress")
	proto.RegisterType((*OverlayOptions)(nil), "OverlayOptions")
	proto.RegisterType((*NodeRep)(nil), "NodeRep")
	proto.RegisterType((*NodeRestrictions)(nil), "NodeRestrictions")
	proto.RegisterType((*Node)(nil), "Node")
	proto.RegisterEnum("NodeTransport", NodeTransport_name, NodeTransport_value)
	proto.RegisterEnum("NodeType", NodeType_name, NodeType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OverlayClient is the client API for Overlay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OverlayClient interface {
	// Lookup finds a nodes address from the network
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error)
	// FindStorageNodes finds a list of nodes in the network that meet the specified request parameters
	FindStorageNodes(ctx context.Context, in *FindStorageNodesRequest, opts ...grpc.CallOption) (*FindStorageNodesResponse, error)
}

type overlayClient struct {
	cc *grpc.ClientConn
}

func NewOverlayClient(cc *grpc.ClientConn) OverlayClient {
	return &overlayClient{cc}
}

func (c *overlayClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := c.cc.Invoke(ctx, "/Overlay/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *overlayClient) FindStorageNodes(ctx context.Context, in *FindStorageNodesRequest, opts ...grpc.CallOption) (*FindStorageNodesResponse, error) {
	out := new(FindStorageNodesResponse)
	err := c.cc.Invoke(ctx, "/Overlay/FindStorageNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OverlayServer is the server API for Overlay service.
type OverlayServer interface {
	// Lookup finds a nodes address from the network
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)
	// FindStorageNodes finds a list of nodes in the network that meet the specified request parameters
	FindStorageNodes(context.Context, *FindStorageNodesRequest) (*FindStorageNodesResponse, error)
}

func RegisterOverlayServer(s *grpc.Server, srv OverlayServer) {
	s.RegisterService(&_Overlay_serviceDesc, srv)
}

func _Overlay_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Overlay/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Overlay_FindStorageNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindStorageNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverlayServer).FindStorageNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Overlay/FindStorageNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverlayServer).FindStorageNodes(ctx, req.(*FindStorageNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Overlay_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Overlay",
	HandlerType: (*OverlayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lookup",
			Handler:    _Overlay_Lookup_Handler,
		},
		{
			MethodName: "FindStorageNodes",
			Handler:    _Overlay_FindStorageNodes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "overlay.proto",
}

func init() { proto.RegisterFile("overlay.proto", fileDescriptor_overlay_8e97a5f3853567d6) }

var fileDescriptor_overlay_8e97a5f3853567d6 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x6d, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0xeb, 0x7c, 0x79, 0x92, 0xb8, 0x61, 0x84, 0xc0, 0x09, 0x02, 0x55, 0x06, 0xf1, 0x51,
	0x90, 0x7f, 0xa4, 0x42, 0x88, 0x9f, 0x81, 0x40, 0x55, 0x11, 0x1a, 0xb4, 0x09, 0x07, 0x70, 0xe2,
	0x6d, 0xba, 0x34, 0xf1, 0x2e, 0xbb, 0x1b, 0x20, 0x48, 0xdc, 0x82, 0x0b, 0x70, 0x3d, 0x4e, 0x81,
	0xbc, 0xeb, 0xa4, 0x71, 0xa0, 0xfc, 0x9c, 0x37, 0x6f, 0x76, 0xde, 0xbc, 0xd9, 0x81, 0x26, 0xff,
	0x42, 0xe5, 0x3c, 0x5e, 0x45, 0x42, 0x72, 0xcd, 0x3b, 0x7e, 0xb2, 0x94, 0xb1, 0x66, 0x3c, 0xb5,
	0x71, 0xf8, 0x08, 0x9a, 0x03, 0xce, 0x2f, 0x97, 0x82, 0xd0, 0xcf, 0x4b, 0xaa, 0x34, 0xde, 0x82,
	0x4a, 0xca, 0x13, 0x7a, 0xda, 0x0f, 0x9c, 0x43, 0xe7, 0xb1, 0x47, 0xf2, 0x28, 0x7c, 0x0a, 0xfe,
	0x9a, 0xa8, 0x04, 0x4f, 0x15, 0xc5, 0x36, 0x94, 0xb2, 0x9c, 0xe1, 0xd5, 0xbb, 0xe5, 0xe8, 0x8c,
	0x27, 0x94, 0x18, 0x28, 0x7c, 0x01, 0xc1, 0x5b, 0x96, 0x26, 0x23, 0xcd, 0x65, 0x3c, 0xa3, 0x59,
	0x42, 0x6d, 0xca, 0xee, 0x40, 0x39, 0xe3, 0xa8, 0xc0, 0x39, 0x74, 0xaf, 0xea, 0x2c, 0x16, 0xfe,
	0x72, 0xe0, 0xf6, 0xdf, 0x95, 0x56, 0xd9, 0x3d, 0x00, 0x3e, 0xf9, 0x44, 0xa7, 0x7a, 0xc4, 0xbe,
	0xdb, 0xae, 0x2e, 0xd9, 0x42, 0xb0, 0x07, 0xfe, 0x94, 0xa7, 0x5a, 0xc6, 0x53, 0x3d, 0xa0, 0xe9,
	0x4c, 0x5f, 0x04, 0xfb, 0x46, 0x59, 0x3b, 0x9a, 0x71, 0x3e, 0x9b, 0x53, 0x3b, 0xf1, 0x64, 0x79,
	0x1e, 0xf5, 0x73, 0x0f, 0xc8, 0x4e, 0x01, 0xde, 0x87, 0x12, 0x17, 0x5a, 0x05, 0xae, 0x29, 0x3c,
	0x88, 0x86, 0xd6, 0xbb, 0xa1, 0xc8, 0xd8, 0x8a, 0x98, 0x64, 0xf8, 0x11, 0xea, 0x99, 0xae, 0x5e,
	0x92, 0x48, 0xaa, 0x14, 0x3e, 0x03, 0x4f, 0xcb, 0x38, 0x55, 0x82, 0x4b, 0x6d, 0x54, 0xf9, 0x5d,
	0xdf, 0xcc, 0x34, 0x5e, 0xa3, 0xe4, 0x8a, 0x80, 0x01, 0x54, 0x63, 0x5b, 0x68, 0xd4, 0x79, 0x64,
	0x1d, 0x86, 0xbf, 0x1d, 0xf0, 0x8b, 0xfd, 0xf0, 0x25, 0xc0, 0x22, 0xfe, 0x36, 0x88, 0x35, 0x4d,
	0xa7, 0xab, 0xdc, 0xe7, 0xff, 0x4c, 0xb3, 0x45, 0xc6, 0x08, 0x9a, 0x0b, 0x96, 0x12, 0x2a, 0x96,
	0xda, 0x24, 0x73, 0x2f, 0x6a, 0xd6, 0x6d, 0x2a, 0x48, 0x31, 0x8d, 0x21, 0x34, 0x16, 0x2c, 0x1d,
	0x09, 0x4a, 0x93, 0x77, 0x13, 0x61, 0x1d, 0x70, 0x49, 0x01, 0xc3, 0x9b, 0x50, 0x9e, 0xb3, 0x05,
	0xd3, 0x41, 0xc9, 0x24, 0x6d, 0x80, 0xc7, 0x50, 0x97, 0x54, 0x69, 0x36, 0x35, 0x9a, 0x83, 0xb2,
	0xe9, 0x73, 0x23, 0xef, 0xa3, 0xb4, 0xcc, 0x13, 0x64, 0x9b, 0x15, 0x7a, 0x50, 0xcd, 0x85, 0x84,
	0x63, 0x68, 0xed, 0x72, 0xf1, 0x01, 0x34, 0xcf, 0x25, 0xa5, 0xaf, 0xe2, 0x34, 0xf9, 0xca, 0x12,
	0x7d, 0x91, 0x6f, 0xbb, 0x08, 0x62, 0x07, 0x6a, 0x19, 0xd0, 0x67, 0xea, 0xd2, 0x8c, 0xe7, 0x92,
	0x4d, 0x1c, 0xfe, 0x74, 0xa0, 0x94, 0x3d, 0x8b, 0x3e, 0xec, 0xb3, 0x24, 0xff, 0xcb, 0xfb, 0x2c,
	0xc1, 0x87, 0xc5, 0x05, 0xd4, 0xbb, 0x8d, 0x68, 0x6b, 0x9b, 0x9b, 0x75, 0xe0, 0x5d, 0x28, 0xe9,
	0x95, 0xa0, 0xc6, 0x08, 0xbf, 0xeb, 0xd9, 0x8d, 0xae, 0x04, 0x25, 0x06, 0xc6, 0xe7, 0xd0, 0x90,
	0x5b, 0x8a, 0x8d, 0x25, 0xff, 0x1c, 0xbb, 0x40, 0x3b, 0x0a, 0xa0, 0x59, 0xf8, 0x1a, 0x58, 0x05,
	0x77, 0xfc, 0xfa, 0x43, 0x6b, 0xef, 0x28, 0x84, 0xda, 0xba, 0x05, 0x7a, 0x50, 0xee, 0xf5, 0xdf,
	0x9f, 0x9e, 0xb5, 0xf6, 0xb0, 0x0e, 0xd5, 0xd1, 0x78, 0x48, 0x7a, 0x27, 0x6f, 0x5a, 0x4e, 0xf7,
	0x07, 0x54, 0xf3, 0x1f, 0x82, 0x4f, 0xa0, 0x62, 0xcf, 0x11, 0xfd, 0xa8, 0x70, 0xc0, 0x9d, 0x83,
	0x68, 0xe7, 0x4e, 0x4f, 0xa0, 0xb5, 0x7b, 0x52, 0x18, 0x44, 0xd7, 0x5c, 0x59, 0xa7, 0x1d, 0x5d,
	0x77, 0xb9, 0x93, 0x8a, 0xf9, 0x72, 0xc7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xd2, 0x9b,
	0x54, 0x53, 0x04, 0x00, 0x00,
}

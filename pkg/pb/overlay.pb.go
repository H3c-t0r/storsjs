// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: overlay.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
	drpc "storj.io/drpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Restriction_Operator int32

const (
	Restriction_LT  Restriction_Operator = 0
	Restriction_EQ  Restriction_Operator = 1
	Restriction_GT  Restriction_Operator = 2
	Restriction_LTE Restriction_Operator = 3
	Restriction_GTE Restriction_Operator = 4
)

var Restriction_Operator_name = map[int32]string{
	0: "LT",
	1: "EQ",
	2: "GT",
	3: "LTE",
	4: "GTE",
}

var Restriction_Operator_value = map[string]int32{
	"LT":  0,
	"EQ":  1,
	"GT":  2,
	"LTE": 3,
	"GTE": 4,
}

func (x Restriction_Operator) String() string {
	return proto.EnumName(Restriction_Operator_name, int32(x))
}

func (Restriction_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{6, 0}
}

type Restriction_Operand int32

const (
	Restriction_FREE_BANDWIDTH Restriction_Operand = 0
	Restriction_FREE_DISK      Restriction_Operand = 1
)

var Restriction_Operand_name = map[int32]string{
	0: "FREE_BANDWIDTH",
	1: "FREE_DISK",
}

var Restriction_Operand_value = map[string]int32{
	"FREE_BANDWIDTH": 0,
	"FREE_DISK":      1,
}

func (x Restriction_Operand) String() string {
	return proto.EnumName(Restriction_Operand_name, int32(x))
}

func (Restriction_Operand) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{6, 1}
}

type QueryRequest struct {
	Sender               *Node      `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Target               *Node      `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Limit                int64      `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Pingback             bool       `protobuf:"varint,4,opt,name=pingback,proto3" json:"pingback,omitempty"`
	Vouchers             []*Voucher `protobuf:"bytes,5,rep,name=vouchers,proto3" json:"vouchers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{0}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetSender() *Node {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *QueryRequest) GetTarget() *Node {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *QueryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryRequest) GetPingback() bool {
	if m != nil {
		return m.Pingback
	}
	return false
}

func (m *QueryRequest) GetVouchers() []*Voucher {
	if m != nil {
		return m.Vouchers
	}
	return nil
}

type QueryResponse struct {
	Sender               *Node    `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Response             []*Node  `protobuf:"bytes,2,rep,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{1}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetSender() *Node {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *QueryResponse) GetResponse() []*Node {
	if m != nil {
		return m.Response
	}
	return nil
}

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{2}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{3}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

// TODO: add fields that validate who is requesting the info
type InfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoRequest) Reset()         { *m = InfoRequest{} }
func (m *InfoRequest) String() string { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()    {}
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{4}
}
func (m *InfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoRequest.Unmarshal(m, b)
}
func (m *InfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoRequest.Marshal(b, m, deterministic)
}
func (m *InfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoRequest.Merge(m, src)
}
func (m *InfoRequest) XXX_Size() int {
	return xxx_messageInfo_InfoRequest.Size(m)
}
func (m *InfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InfoRequest proto.InternalMessageInfo

type InfoResponse struct {
	Type                 NodeType      `protobuf:"varint,2,opt,name=type,proto3,enum=node.NodeType" json:"type,omitempty"`
	Operator             *NodeOperator `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Capacity             *NodeCapacity `protobuf:"bytes,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Version              *NodeVersion  `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{5}
}
func (m *InfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoResponse.Unmarshal(m, b)
}
func (m *InfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoResponse.Marshal(b, m, deterministic)
}
func (m *InfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse.Merge(m, src)
}
func (m *InfoResponse) XXX_Size() int {
	return xxx_messageInfo_InfoResponse.Size(m)
}
func (m *InfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse proto.InternalMessageInfo

func (m *InfoResponse) GetType() NodeType {
	if m != nil {
		return m.Type
	}
	return NodeType_INVALID
}

func (m *InfoResponse) GetOperator() *NodeOperator {
	if m != nil {
		return m.Operator
	}
	return nil
}

func (m *InfoResponse) GetCapacity() *NodeCapacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}

func (m *InfoResponse) GetVersion() *NodeVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

type Restriction struct {
	Operator             Restriction_Operator `protobuf:"varint,1,opt,name=operator,proto3,enum=overlay.Restriction_Operator" json:"operator,omitempty"`
	Operand              Restriction_Operand  `protobuf:"varint,2,opt,name=operand,proto3,enum=overlay.Restriction_Operand" json:"operand,omitempty"`
	Value                int64                `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Restriction) Reset()         { *m = Restriction{} }
func (m *Restriction) String() string { return proto.CompactTextString(m) }
func (*Restriction) ProtoMessage()    {}
func (*Restriction) Descriptor() ([]byte, []int) {
	return fileDescriptor_61fc82527fbe24ad, []int{6}
}
func (m *Restriction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Restriction.Unmarshal(m, b)
}
func (m *Restriction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Restriction.Marshal(b, m, deterministic)
}
func (m *Restriction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Restriction.Merge(m, src)
}
func (m *Restriction) XXX_Size() int {
	return xxx_messageInfo_Restriction.Size(m)
}
func (m *Restriction) XXX_DiscardUnknown() {
	xxx_messageInfo_Restriction.DiscardUnknown(m)
}

var xxx_messageInfo_Restriction proto.InternalMessageInfo

func (m *Restriction) GetOperator() Restriction_Operator {
	if m != nil {
		return m.Operator
	}
	return Restriction_LT
}

func (m *Restriction) GetOperand() Restriction_Operand {
	if m != nil {
		return m.Operand
	}
	return Restriction_FREE_BANDWIDTH
}

func (m *Restriction) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("overlay.Restriction_Operator", Restriction_Operator_name, Restriction_Operator_value)
	proto.RegisterEnum("overlay.Restriction_Operand", Restriction_Operand_name, Restriction_Operand_value)
	proto.RegisterType((*QueryRequest)(nil), "overlay.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "overlay.QueryResponse")
	proto.RegisterType((*PingRequest)(nil), "overlay.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "overlay.PingResponse")
	proto.RegisterType((*InfoRequest)(nil), "overlay.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "overlay.InfoResponse")
	proto.RegisterType((*Restriction)(nil), "overlay.Restriction")
}

func init() { proto.RegisterFile("overlay.proto", fileDescriptor_61fc82527fbe24ad) }

var fileDescriptor_61fc82527fbe24ad = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x77, 0xfa, 0x33, 0xbe, 0xfe, 0x20, 0x3b, 0xec, 0x4a, 0x28, 0x0a, 0x25, 0x07, 0x29,
	0xa8, 0x39, 0x64, 0x65, 0x41, 0x6f, 0xae, 0x8d, 0x6b, 0x71, 0x59, 0xdd, 0x31, 0xac, 0xa0, 0x07,
	0x49, 0x93, 0x31, 0x06, 0xeb, 0x4c, 0x9c, 0x4c, 0x0b, 0xf9, 0xaf, 0xbc, 0x79, 0xf4, 0xff, 0xf2,
	0x24, 0x33, 0x99, 0xa4, 0x71, 0x55, 0xf0, 0x94, 0x79, 0xef, 0xfb, 0x79, 0x9d, 0x6f, 0x5e, 0xbf,
	0x81, 0x09, 0xdf, 0x51, 0xb1, 0x89, 0x4a, 0x2f, 0x17, 0x5c, 0x72, 0x3c, 0x34, 0xe5, 0x0c, 0x52,
	0x9e, 0xf2, 0xaa, 0x39, 0x03, 0xc6, 0x13, 0x6a, 0xce, 0xd3, 0x1d, 0xdf, 0xc6, 0x9f, 0xa8, 0x28,
	0xaa, 0xda, 0xfd, 0x8e, 0x60, 0x7c, 0xb5, 0xa5, 0xa2, 0x24, 0xf4, 0xeb, 0x96, 0x16, 0x12, 0xbb,
	0x30, 0x28, 0x28, 0x4b, 0xa8, 0x70, 0xd0, 0x1c, 0x2d, 0x46, 0x3e, 0x78, 0x7a, 0xfa, 0x92, 0x27,
	0x94, 0x18, 0x45, 0x31, 0x32, 0x12, 0x29, 0x95, 0x4e, 0xe7, 0x4f, 0xa6, 0x52, 0xf0, 0x11, 0xf4,
	0x37, 0xd9, 0x97, 0x4c, 0x3a, 0xdd, 0x39, 0x5a, 0x74, 0x49, 0x55, 0xe0, 0x19, 0x58, 0x79, 0xc6,
	0xd2, 0x75, 0x14, 0x7f, 0x76, 0x7a, 0x73, 0xb4, 0xb0, 0x48, 0x53, 0xe3, 0x87, 0x60, 0xd5, 0xe6,
	0x9c, 0xfe, 0xbc, 0xbb, 0x18, 0xf9, 0x87, 0x5e, 0xe3, 0xf6, 0xba, 0x3a, 0x90, 0x06, 0x71, 0xdf,
	0xc3, 0xc4, 0x18, 0x2f, 0x72, 0xce, 0x0a, 0xfa, 0x5f, 0xce, 0xef, 0x81, 0x25, 0x0c, 0xef, 0x74,
	0xf4, 0x1d, 0x6d, 0xaa, 0xd1, 0xdc, 0x09, 0x8c, 0x5e, 0x67, 0x2c, 0x35, 0x4b, 0x71, 0xa7, 0x30,
	0xae, 0xca, 0xbd, 0xbc, 0x62, 0x1f, 0x79, 0x2d, 0xff, 0x40, 0x30, 0xae, 0xea, 0xc6, 0x4a, 0x4f,
	0x96, 0x39, 0xd5, 0xeb, 0x99, 0xfa, 0xd3, 0xfd, 0x15, 0x61, 0x99, 0x53, 0xa2, 0x35, 0xec, 0x81,
	0xc5, 0x73, 0x2a, 0x22, 0xc9, 0x85, 0xde, 0xd1, 0xc8, 0xc7, 0x7b, 0xee, 0x95, 0x51, 0x48, 0xc3,
	0x28, 0x3e, 0x8e, 0xf2, 0x28, 0xce, 0x64, 0xa9, 0x57, 0xf7, 0x1b, 0xff, 0xcc, 0x28, 0xa4, 0x61,
	0xf0, 0x7d, 0x18, 0xee, 0xa8, 0x28, 0x32, 0xce, 0x9c, 0xbe, 0xc6, 0x0f, 0xf7, 0xf8, 0x75, 0x25,
	0x90, 0x9a, 0x70, 0x7f, 0x22, 0x18, 0x11, 0x5a, 0x48, 0x91, 0xc5, 0x32, 0xe3, 0x0c, 0x3f, 0x6e,
	0x99, 0x43, 0xfa, 0x25, 0xee, 0x7a, 0x75, 0xd2, 0x5a, 0x9c, 0xf7, 0x17, 0x9f, 0xa7, 0x30, 0xd4,
	0x67, 0x96, 0x98, 0xd7, 0xbf, 0xf3, 0xef, 0x49, 0x96, 0x90, 0x1a, 0x56, 0x81, 0xd9, 0x45, 0x9b,
	0x2d, 0xad, 0x03, 0xa3, 0x0b, 0xf7, 0x11, 0x58, 0xf5, 0x1d, 0x78, 0x00, 0x9d, 0x8b, 0xd0, 0x3e,
	0x50, 0xcf, 0xe0, 0xca, 0x46, 0xea, 0x79, 0x1e, 0xda, 0x1d, 0x3c, 0x84, 0xee, 0x45, 0x18, 0xd8,
	0x5d, 0x75, 0x38, 0x0f, 0x03, 0xbb, 0xe7, 0x3e, 0x80, 0xa1, 0xf9, 0x7d, 0x8c, 0x61, 0xfa, 0x9c,
	0x04, 0xc1, 0x87, 0xb3, 0xa7, 0x97, 0xcb, 0xb7, 0xab, 0x65, 0xf8, 0xc2, 0x3e, 0xc0, 0x13, 0xb8,
	0xa5, 0x7b, 0xcb, 0xd5, 0x9b, 0x97, 0x36, 0xf2, 0xbf, 0x21, 0xe8, 0xab, 0xad, 0x14, 0xf8, 0x14,
	0xfa, 0x3a, 0x53, 0xf8, 0xb8, 0xf1, 0xdc, 0xfe, 0x38, 0x66, 0xb7, 0x6f, 0xb6, 0xcd, 0xff, 0x7d,
	0x02, 0x3d, 0x95, 0x0f, 0x7c, 0xd4, 0xe8, 0xad, 0xf4, 0xcc, 0x8e, 0x6f, 0x74, 0xcd, 0xd0, 0x13,
	0xb5, 0x72, 0x4d, 0xa8, 0xec, 0xb4, 0x66, 0x5b, 0xd1, 0x6a, 0xcd, 0xb6, 0x03, 0x76, 0xd6, 0x7b,
	0xd7, 0xc9, 0xd7, 0xeb, 0x81, 0xfe, 0x86, 0x4f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xae, 0xb0,
	0xad, 0x04, 0x05, 0x04, 0x00, 0x00,
}

type DRPCNodesClient interface {
	DRPCConn() drpc.Conn

	Query(ctx context.Context, in *QueryRequest) (*QueryResponse, error)
	Ping(ctx context.Context, in *PingRequest) (*PingResponse, error)
	RequestInfo(ctx context.Context, in *InfoRequest) (*InfoResponse, error)
}

type drpcNodesClient struct {
	cc drpc.Conn
}

func NewDRPCNodesClient(cc drpc.Conn) DRPCNodesClient {
	return &drpcNodesClient{cc}
}

func (c *drpcNodesClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcNodesClient) Query(ctx context.Context, in *QueryRequest) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/overlay.Nodes/Query", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcNodesClient) Ping(ctx context.Context, in *PingRequest) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/overlay.Nodes/Ping", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcNodesClient) RequestInfo(ctx context.Context, in *InfoRequest) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/overlay.Nodes/RequestInfo", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCNodesServer interface {
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	RequestInfo(context.Context, *InfoRequest) (*InfoResponse, error)
}

type DRPCNodesDescription struct{}

func (DRPCNodesDescription) NumMethods() int { return 3 }

func (DRPCNodesDescription) Method(n int) (string, drpc.Handler, interface{}, bool) {
	switch n {
	case 0:
		return "/overlay.Nodes/Query",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodesServer).
					Query(
						ctx,
						in1.(*QueryRequest),
					)
			}, DRPCNodesServer.Query, true
	case 1:
		return "/overlay.Nodes/Ping",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodesServer).
					Ping(
						ctx,
						in1.(*PingRequest),
					)
			}, DRPCNodesServer.Ping, true
	case 2:
		return "/overlay.Nodes/RequestInfo",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCNodesServer).
					RequestInfo(
						ctx,
						in1.(*InfoRequest),
					)
			}, DRPCNodesServer.RequestInfo, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterNodes(srv drpc.Server, impl DRPCNodesServer) {
	srv.Register(impl, DRPCNodesDescription{})
}

type DRPCNodes_QueryStream interface {
	drpc.Stream
	SendAndClose(*QueryResponse) error
}

type drpcNodesQueryStream struct {
	drpc.Stream
}

func (x *drpcNodesQueryStream) SendAndClose(m *QueryResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNodes_PingStream interface {
	drpc.Stream
	SendAndClose(*PingResponse) error
}

type drpcNodesPingStream struct {
	drpc.Stream
}

func (x *drpcNodesPingStream) SendAndClose(m *PingResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCNodes_RequestInfoStream interface {
	drpc.Stream
	SendAndClose(*InfoResponse) error
}

type drpcNodesRequestInfoStream struct {
	drpc.Stream
}

func (x *drpcNodesRequestInfoStream) SendAndClose(m *InfoResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodesClient is the client API for Nodes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodesClient interface {
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	RequestInfo(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}

type nodesClient struct {
	cc *grpc.ClientConn
}

func NewNodesClient(cc *grpc.ClientConn) NodesClient {
	return &nodesClient{cc}
}

func (c *nodesClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/overlay.Nodes/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/overlay.Nodes/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodesClient) RequestInfo(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/overlay.Nodes/RequestInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodesServer is the server API for Nodes service.
type NodesServer interface {
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	RequestInfo(context.Context, *InfoRequest) (*InfoResponse, error)
}

func RegisterNodesServer(s *grpc.Server, srv NodesServer) {
	s.RegisterService(&_Nodes_serviceDesc, srv)
}

func _Nodes_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Nodes/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nodes_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Nodes/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nodes_RequestInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodesServer).RequestInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/overlay.Nodes/RequestInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodesServer).RequestInfo(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nodes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "overlay.Nodes",
	HandlerType: (*NodesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Nodes_Query_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Nodes_Ping_Handler,
		},
		{
			MethodName: "RequestInfo",
			Handler:    _Nodes_RequestInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "overlay.proto",
}

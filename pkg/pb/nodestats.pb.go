// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nodestats.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

type UptimeCheckRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UptimeCheckRequest) Reset()         { *m = UptimeCheckRequest{} }
func (m *UptimeCheckRequest) String() string { return proto.CompactTextString(m) }
func (*UptimeCheckRequest) ProtoMessage()    {}
func (*UptimeCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{0}
}
func (m *UptimeCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UptimeCheckRequest.Unmarshal(m, b)
}
func (m *UptimeCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UptimeCheckRequest.Marshal(b, m, deterministic)
}
func (m *UptimeCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeCheckRequest.Merge(m, src)
}
func (m *UptimeCheckRequest) XXX_Size() int {
	return xxx_messageInfo_UptimeCheckRequest.Size(m)
}
func (m *UptimeCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeCheckRequest proto.InternalMessageInfo

type UptimeCheckResponse struct {
	TotalCount           int64    `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	SuccessCount         int64    `protobuf:"varint,2,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
	ReputationAlpha      float64  `protobuf:"fixed64,3,opt,name=reputation_alpha,json=reputationAlpha,proto3" json:"reputation_alpha,omitempty"`
	ReputationBeta       float64  `protobuf:"fixed64,4,opt,name=reputation_beta,json=reputationBeta,proto3" json:"reputation_beta,omitempty"`
	ReputationScore      float64  `protobuf:"fixed64,5,opt,name=reputation_score,json=reputationScore,proto3" json:"reputation_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UptimeCheckResponse) Reset()         { *m = UptimeCheckResponse{} }
func (m *UptimeCheckResponse) String() string { return proto.CompactTextString(m) }
func (*UptimeCheckResponse) ProtoMessage()    {}
func (*UptimeCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0b184ee117142aa, []int{1}
}
func (m *UptimeCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UptimeCheckResponse.Unmarshal(m, b)
}
func (m *UptimeCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UptimeCheckResponse.Marshal(b, m, deterministic)
}
func (m *UptimeCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeCheckResponse.Merge(m, src)
}
func (m *UptimeCheckResponse) XXX_Size() int {
	return xxx_messageInfo_UptimeCheckResponse.Size(m)
}
func (m *UptimeCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeCheckResponse proto.InternalMessageInfo

func (m *UptimeCheckResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *UptimeCheckResponse) GetSuccessCount() int64 {
	if m != nil {
		return m.SuccessCount
	}
	return 0
}

func (m *UptimeCheckResponse) GetReputationAlpha() float64 {
	if m != nil {
		return m.ReputationAlpha
	}
	return 0
}

func (m *UptimeCheckResponse) GetReputationBeta() float64 {
	if m != nil {
		return m.ReputationBeta
	}
	return 0
}

func (m *UptimeCheckResponse) GetReputationScore() float64 {
	if m != nil {
		return m.ReputationScore
	}
	return 0
}

func init() {
	proto.RegisterType((*UptimeCheckRequest)(nil), "nodestats.UptimeCheckRequest")
	proto.RegisterType((*UptimeCheckResponse)(nil), "nodestats.UptimeCheckResponse")
}

func init() { proto.RegisterFile("nodestats.proto", fileDescriptor_e0b184ee117142aa) }

var fileDescriptor_e0b184ee117142aa = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0x71, 0x17, 0xd0, 0x84, 0x41, 0xc5, 0xac, 0x1e, 0x1a, 0x12, 0x95, 0xd4, 0x83, 0x78,
	0xe1, 0xa0, 0x4f, 0x20, 0xdc, 0x39, 0x94, 0x78, 0xd1, 0x43, 0xb3, 0x5d, 0x26, 0x40, 0xc4, 0xce,
	0xda, 0x99, 0x7d, 0x59, 0x9f, 0xc6, 0xec, 0x96, 0x58, 0xac, 0xf1, 0xfa, 0xdb, 0x2f, 0x93, 0xec,
	0x1f, 0x86, 0x25, 0xad, 0x90, 0xc5, 0x08, 0x4f, 0x5d, 0x45, 0x42, 0xba, 0xff, 0x03, 0x23, 0x58,
	0xd3, 0x9a, 0x6a, 0x4e, 0xaf, 0x40, 0xbf, 0x38, 0xd9, 0x7e, 0xe0, 0x7c, 0x83, 0xf6, 0x3d, 0xc3,
	0x4f, 0x8f, 0x2c, 0xe9, 0x97, 0x82, 0xcb, 0x5f, 0xcc, 0x8e, 0x4a, 0x46, 0x7d, 0x0b, 0x03, 0x21,
	0x31, 0xbb, 0xdc, 0x92, 0x2f, 0x25, 0x51, 0x63, 0x35, 0xe9, 0x66, 0x10, 0x69, 0x1e, 0x44, 0xdf,
	0xc1, 0x19, 0x7b, 0x6b, 0x91, 0x79, 0x3f, 0xe9, 0xc4, 0xc9, 0xe9, 0x1e, 0xeb, 0xd1, 0x03, 0x5c,
	0x54, 0xe8, 0xbc, 0x18, 0xd9, 0x52, 0x99, 0x9b, 0x9d, 0xdb, 0x98, 0xa4, 0x3b, 0x56, 0x13, 0x95,
	0x0d, 0x1b, 0x7f, 0x0e, 0xac, 0xef, 0xe1, 0x80, 0xf2, 0x02, 0xc5, 0x24, 0xbd, 0xb8, 0x3c, 0x6f,
	0x78, 0x86, 0x62, 0x5a, 0x37, 0xd9, 0x52, 0x85, 0xc9, 0x71, 0xfb, 0xe6, 0x32, 0xf0, 0xe3, 0x1b,
	0xf4, 0x17, 0xb4, 0xc2, 0x65, 0x68, 0xa1, 0x17, 0x30, 0x38, 0xf8, 0xa8, 0xbe, 0x9e, 0x36, 0xdd,
	0xfe, 0x76, 0x19, 0xdd, 0xfc, 0xf7, 0x5c, 0xf7, 0x49, 0x8f, 0x66, 0xbd, 0xd7, 0x8e, 0x2b, 0x8a,
	0x93, 0x18, 0xf7, 0xe9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x45, 0x2f, 0x78, 0x86, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeStatsClient is the client API for NodeStats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeStatsClient interface {
	UptimeCheck(ctx context.Context, in *UptimeCheckRequest, opts ...grpc.CallOption) (*UptimeCheckResponse, error)
}

type nodeStatsClient struct {
	cc *grpc.ClientConn
}

func NewNodeStatsClient(cc *grpc.ClientConn) NodeStatsClient {
	return &nodeStatsClient{cc}
}

func (c *nodeStatsClient) UptimeCheck(ctx context.Context, in *UptimeCheckRequest, opts ...grpc.CallOption) (*UptimeCheckResponse, error) {
	out := new(UptimeCheckResponse)
	err := c.cc.Invoke(ctx, "/nodestats.NodeStats/UptimeCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeStatsServer is the server API for NodeStats service.
type NodeStatsServer interface {
	UptimeCheck(context.Context, *UptimeCheckRequest) (*UptimeCheckResponse, error)
}

func RegisterNodeStatsServer(s *grpc.Server, srv NodeStatsServer) {
	s.RegisterService(&_NodeStats_serviceDesc, srv)
}

func _NodeStats_UptimeCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UptimeCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeStatsServer).UptimeCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nodestats.NodeStats/UptimeCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeStatsServer).UptimeCheck(ctx, req.(*UptimeCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeStats_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nodestats.NodeStats",
	HandlerType: (*NodeStatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UptimeCheck",
			Handler:    _NodeStats_UptimeCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nodestats.proto",
}

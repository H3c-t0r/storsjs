// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vouchers.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Voucher is a signed message verifying that a node has been vetted by a particular satellite
type Voucher struct {
	SatelliteId          NodeID               `protobuf:"bytes,1,opt,name=satellite_id,json=satelliteId,proto3,customtype=NodeID" json:"satellite_id"`
	StorageNodeId        NodeID               `protobuf:"bytes,2,opt,name=storage_node_id,json=storageNodeId,proto3,customtype=NodeID" json:"storage_node_id"`
	Expiration           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	SatelliteSignature   []byte               `protobuf:"bytes,4,opt,name=satellite_signature,json=satelliteSignature,proto3" json:"satellite_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Voucher) Reset()         { *m = Voucher{} }
func (m *Voucher) String() string { return proto.CompactTextString(m) }
func (*Voucher) ProtoMessage()    {}
func (*Voucher) Descriptor() ([]byte, []int) {
	return fileDescriptor_3659b9a115b8060d, []int{0}
}
func (m *Voucher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Voucher.Unmarshal(m, b)
}
func (m *Voucher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Voucher.Marshal(b, m, deterministic)
}
func (m *Voucher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Voucher.Merge(m, src)
}
func (m *Voucher) XXX_Size() int {
	return xxx_messageInfo_Voucher.Size(m)
}
func (m *Voucher) XXX_DiscardUnknown() {
	xxx_messageInfo_Voucher.DiscardUnknown(m)
}

var xxx_messageInfo_Voucher proto.InternalMessageInfo

func (m *Voucher) GetExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *Voucher) GetSatelliteSignature() []byte {
	if m != nil {
		return m.SatelliteSignature
	}
	return nil
}

type VoucherRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoucherRequest) Reset()         { *m = VoucherRequest{} }
func (m *VoucherRequest) String() string { return proto.CompactTextString(m) }
func (*VoucherRequest) ProtoMessage()    {}
func (*VoucherRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3659b9a115b8060d, []int{1}
}
func (m *VoucherRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoucherRequest.Unmarshal(m, b)
}
func (m *VoucherRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoucherRequest.Marshal(b, m, deterministic)
}
func (m *VoucherRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoucherRequest.Merge(m, src)
}
func (m *VoucherRequest) XXX_Size() int {
	return xxx_messageInfo_VoucherRequest.Size(m)
}
func (m *VoucherRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VoucherRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VoucherRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Voucher)(nil), "vouchers.Voucher")
	proto.RegisterType((*VoucherRequest)(nil), "vouchers.VoucherRequest")
}

func init() { proto.RegisterFile("vouchers.proto", fileDescriptor_3659b9a115b8060d) }

var fileDescriptor_3659b9a115b8060d = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xc7, 0x9b, 0x52, 0xb5, 0xd5, 0xb5, 0x14, 0x30, 0x4b, 0x94, 0x25, 0x55, 0xa6, 0x4e, 0x8e,
	0x28, 0x12, 0x42, 0x8c, 0x55, 0x97, 0x2e, 0x0c, 0x01, 0x31, 0xb0, 0x54, 0x0e, 0x3e, 0x8c, 0xa5,
	0x34, 0x17, 0x6c, 0x07, 0xf1, 0x88, 0x3c, 0x03, 0x43, 0x25, 0xde, 0x04, 0x91, 0x2f, 0x40, 0x30,
	0xfe, 0xbf, 0xe4, 0x9f, 0x0f, 0x66, 0x2f, 0x54, 0x3e, 0x3c, 0xa1, 0xb1, 0xbc, 0x30, 0xe4, 0x88,
	0x8d, 0x5b, 0x1d, 0x80, 0x22, 0x45, 0xb5, 0x1b, 0x84, 0x8a, 0x48, 0x65, 0x18, 0x57, 0x2a, 0x2d,
	0x1f, 0x63, 0xa7, 0x77, 0x68, 0x9d, 0xd8, 0x15, 0x75, 0x21, 0xfa, 0xf0, 0x60, 0x74, 0x57, 0x2f,
	0xd9, 0x19, 0x4c, 0xad, 0x70, 0x98, 0x65, 0xda, 0xe1, 0x56, 0x4b, 0xdf, 0x9b, 0x7b, 0x8b, 0xe9,
	0x6a, 0xf6, 0xb6, 0x0f, 0x7b, 0xef, 0xfb, 0x70, 0x78, 0x4d, 0x12, 0x37, 0xeb, 0x64, 0xd2, 0x75,
	0x36, 0x92, 0x5d, 0xc0, 0x91, 0x75, 0x64, 0x84, 0xc2, 0x6d, 0x4e, 0xb2, 0x5a, 0xf5, 0xff, 0x5d,
	0x1d, 0x36, 0xb5, 0x4a, 0x4a, 0x76, 0x05, 0x80, 0xaf, 0x85, 0x36, 0xc2, 0x69, 0xca, 0xfd, 0x83,
	0xb9, 0xb7, 0x98, 0x2c, 0x03, 0x5e, 0xc3, 0xf2, 0x16, 0x96, 0xdf, 0xb6, 0xb0, 0xc9, 0x8f, 0x36,
	0x8b, 0xe1, 0xf4, 0x1b, 0xd3, 0x6a, 0x95, 0x0b, 0x57, 0x1a, 0xf4, 0x07, 0x5f, 0xef, 0x26, 0xac,
	0x8b, 0x6e, 0xda, 0x24, 0x3a, 0x86, 0x59, 0xf3, 0xc5, 0x04, 0x9f, 0x4b, 0xb4, 0x6e, 0xb9, 0x86,
	0x71, 0xe3, 0x58, 0x76, 0x09, 0xa3, 0xc6, 0x66, 0x3e, 0xef, 0x8e, 0xfa, 0x7b, 0x10, 0x9c, 0xfc,
	0x49, 0xa2, 0xde, 0x6a, 0x70, 0xdf, 0x2f, 0xd2, 0x74, 0x58, 0xe1, 0x9e, 0x7f, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x0a, 0x25, 0x8e, 0x62, 0x91, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VouchersClient is the client API for Vouchers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VouchersClient interface {
	Request(ctx context.Context, in *VoucherRequest, opts ...grpc.CallOption) (*Voucher, error)
}

type vouchersClient struct {
	cc *grpc.ClientConn
}

func NewVouchersClient(cc *grpc.ClientConn) VouchersClient {
	return &vouchersClient{cc}
}

func (c *vouchersClient) Request(ctx context.Context, in *VoucherRequest, opts ...grpc.CallOption) (*Voucher, error) {
	out := new(Voucher)
	err := c.cc.Invoke(ctx, "/vouchers.Vouchers/Request", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VouchersServer is the server API for Vouchers service.
type VouchersServer interface {
	Request(context.Context, *VoucherRequest) (*Voucher, error)
}

func RegisterVouchersServer(s *grpc.Server, srv VouchersServer) {
	s.RegisterService(&_Vouchers_serviceDesc, srv)
}

func _Vouchers_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoucherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VouchersServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vouchers.Vouchers/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VouchersServer).Request(ctx, req.(*VoucherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Vouchers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vouchers.Vouchers",
	HandlerType: (*VouchersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _Vouchers_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vouchers.proto",
}

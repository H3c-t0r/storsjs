// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: referralmanager.proto

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

type GetTokensRequest struct {
	UserId               []byte   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokensRequest) Reset()         { *m = GetTokensRequest{} }
func (m *GetTokensRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokensRequest) ProtoMessage()    {}
func (*GetTokensRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45d96ad24f1e021c, []int{0}
}
func (m *GetTokensRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokensRequest.Unmarshal(m, b)
}
func (m *GetTokensRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokensRequest.Marshal(b, m, deterministic)
}
func (m *GetTokensRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokensRequest.Merge(m, src)
}
func (m *GetTokensRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokensRequest.Size(m)
}
func (m *GetTokensRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokensRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokensRequest proto.InternalMessageInfo

func (m *GetTokensRequest) GetUserId() []byte {
	if m != nil {
		return m.UserId
	}
	return nil
}

type GetTokensResponse struct {
	Token                [][]byte `protobuf:"bytes,1,rep,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokensResponse) Reset()         { *m = GetTokensResponse{} }
func (m *GetTokensResponse) String() string { return proto.CompactTextString(m) }
func (*GetTokensResponse) ProtoMessage()    {}
func (*GetTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45d96ad24f1e021c, []int{1}
}
func (m *GetTokensResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokensResponse.Unmarshal(m, b)
}
func (m *GetTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokensResponse.Marshal(b, m, deterministic)
}
func (m *GetTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokensResponse.Merge(m, src)
}
func (m *GetTokensResponse) XXX_Size() int {
	return xxx_messageInfo_GetTokensResponse.Size(m)
}
func (m *GetTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokensResponse proto.InternalMessageInfo

func (m *GetTokensResponse) GetToken() [][]byte {
	if m != nil {
		return m.Token
	}
	return nil
}

type RedeemTokenRequest struct {
	Token                []byte   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId               []byte   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SatelliteId          []byte   `protobuf:"bytes,3,opt,name=satellite_id,json=satelliteId,proto3" json:"satellite_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedeemTokenRequest) Reset()         { *m = RedeemTokenRequest{} }
func (m *RedeemTokenRequest) String() string { return proto.CompactTextString(m) }
func (*RedeemTokenRequest) ProtoMessage()    {}
func (*RedeemTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45d96ad24f1e021c, []int{2}
}
func (m *RedeemTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedeemTokenRequest.Unmarshal(m, b)
}
func (m *RedeemTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedeemTokenRequest.Marshal(b, m, deterministic)
}
func (m *RedeemTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedeemTokenRequest.Merge(m, src)
}
func (m *RedeemTokenRequest) XXX_Size() int {
	return xxx_messageInfo_RedeemTokenRequest.Size(m)
}
func (m *RedeemTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RedeemTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RedeemTokenRequest proto.InternalMessageInfo

func (m *RedeemTokenRequest) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *RedeemTokenRequest) GetUserId() []byte {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *RedeemTokenRequest) GetSatelliteId() []byte {
	if m != nil {
		return m.SatelliteId
	}
	return nil
}

type RedeemTokenResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedeemTokenResponse) Reset()         { *m = RedeemTokenResponse{} }
func (m *RedeemTokenResponse) String() string { return proto.CompactTextString(m) }
func (*RedeemTokenResponse) ProtoMessage()    {}
func (*RedeemTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45d96ad24f1e021c, []int{3}
}
func (m *RedeemTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedeemTokenResponse.Unmarshal(m, b)
}
func (m *RedeemTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedeemTokenResponse.Marshal(b, m, deterministic)
}
func (m *RedeemTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedeemTokenResponse.Merge(m, src)
}
func (m *RedeemTokenResponse) XXX_Size() int {
	return xxx_messageInfo_RedeemTokenResponse.Size(m)
}
func (m *RedeemTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RedeemTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RedeemTokenResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetTokensRequest)(nil), "referralmanager.GetTokensRequest")
	proto.RegisterType((*GetTokensResponse)(nil), "referralmanager.GetTokensResponse")
	proto.RegisterType((*RedeemTokenRequest)(nil), "referralmanager.RedeemTokenRequest")
	proto.RegisterType((*RedeemTokenResponse)(nil), "referralmanager.RedeemTokenResponse")
}

func init() { proto.RegisterFile("referralmanager.proto", fileDescriptor_45d96ad24f1e021c) }

var fileDescriptor_45d96ad24f1e021c = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4a, 0x4d, 0x4b,
	0x2d, 0x2a, 0x4a, 0xcc, 0xc9, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x47, 0x13, 0x96, 0xe2, 0x4a, 0xcf, 0x4f, 0xcf, 0x87, 0x48, 0x2a, 0x69, 0x73,
	0x09, 0xb8, 0xa7, 0x96, 0x84, 0xe4, 0x67, 0xa7, 0xe6, 0x15, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16,
	0x97, 0x08, 0x89, 0x73, 0xb1, 0x97, 0x16, 0xa7, 0x16, 0xc5, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0xf0, 0x04, 0xb1, 0x81, 0xb8, 0x9e, 0x29, 0x4a, 0x9a, 0x5c, 0x82, 0x48, 0x8a, 0x8b, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x44, 0xb8, 0x58, 0x4b, 0x40, 0x22, 0x12, 0x8c, 0x0a, 0xcc, 0x1a,
	0x3c, 0x41, 0x10, 0x8e, 0x52, 0x1a, 0x97, 0x50, 0x50, 0x6a, 0x4a, 0x6a, 0x6a, 0x2e, 0x58, 0x35,
	0xcc, 0x64, 0x24, 0xb5, 0x8c, 0x70, 0xb5, 0xc8, 0xf6, 0x31, 0x21, 0xdb, 0x27, 0xa4, 0xc8, 0xc5,
	0x53, 0x9c, 0x58, 0x92, 0x9a, 0x93, 0x93, 0x59, 0x92, 0x0a, 0x92, 0x65, 0x06, 0xcb, 0x72, 0xc3,
	0xc5, 0x3c, 0x53, 0x94, 0x44, 0xb9, 0x84, 0x51, 0xec, 0x81, 0x38, 0xca, 0x68, 0x3f, 0x23, 0x17,
	0x7f, 0x10, 0xd4, 0xdb, 0xbe, 0x10, 0x6f, 0x0b, 0x05, 0x71, 0x71, 0xc2, 0x5d, 0x2f, 0xa4, 0xa8,
	0x87, 0x1e, 0x58, 0xe8, 0xc1, 0x20, 0xa5, 0x84, 0x4f, 0x09, 0xd4, 0xf3, 0x11, 0x5c, 0xdc, 0x48,
	0xd6, 0x0b, 0x29, 0x63, 0x68, 0xc1, 0x0c, 0x04, 0x29, 0x15, 0xfc, 0x8a, 0x20, 0x26, 0x3b, 0xb1,
	0x44, 0x31, 0x15, 0x24, 0x25, 0xb1, 0x81, 0x63, 0xc9, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x46,
	0x89, 0x48, 0x05, 0xdb, 0x01, 0x00, 0x00,
}

type DRPCReferralManagerClient interface {
	DRPCConn() drpc.Conn

	// GetTokens retrieves a list of unredeemed tokens for a user
	GetTokens(ctx context.Context, in *GetTokensRequest) (*GetTokensResponse, error)
	// RedeemToken saves newly created user info in referral manager
	RedeemToken(ctx context.Context, in *RedeemTokenRequest) (*RedeemTokenResponse, error)
}

type drpcReferralManagerClient struct {
	cc drpc.Conn
}

func NewDRPCReferralManagerClient(cc drpc.Conn) DRPCReferralManagerClient {
	return &drpcReferralManagerClient{cc}
}

func (c *drpcReferralManagerClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcReferralManagerClient) GetTokens(ctx context.Context, in *GetTokensRequest) (*GetTokensResponse, error) {
	out := new(GetTokensResponse)
	err := c.cc.Invoke(ctx, "/referralmanager.ReferralManager/GetTokens", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcReferralManagerClient) RedeemToken(ctx context.Context, in *RedeemTokenRequest) (*RedeemTokenResponse, error) {
	out := new(RedeemTokenResponse)
	err := c.cc.Invoke(ctx, "/referralmanager.ReferralManager/RedeemToken", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCReferralManagerServer interface {
	// GetTokens retrieves a list of unredeemed tokens for a user
	GetTokens(context.Context, *GetTokensRequest) (*GetTokensResponse, error)
	// RedeemToken saves newly created user info in referral manager
	RedeemToken(context.Context, *RedeemTokenRequest) (*RedeemTokenResponse, error)
}

type DRPCReferralManagerDescription struct{}

func (DRPCReferralManagerDescription) NumMethods() int { return 2 }

func (DRPCReferralManagerDescription) Method(n int) (string, drpc.Handler, interface{}, bool) {
	switch n {
	case 0:
		return "/referralmanager.ReferralManager/GetTokens",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCReferralManagerServer).
					GetTokens(
						ctx,
						in1.(*GetTokensRequest),
					)
			}, DRPCReferralManagerServer.GetTokens, true
	case 1:
		return "/referralmanager.ReferralManager/RedeemToken",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCReferralManagerServer).
					RedeemToken(
						ctx,
						in1.(*RedeemTokenRequest),
					)
			}, DRPCReferralManagerServer.RedeemToken, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterReferralManager(srv drpc.Server, impl DRPCReferralManagerServer) {
	srv.Register(impl, DRPCReferralManagerDescription{})
}

type DRPCReferralManager_GetTokensStream interface {
	drpc.Stream
	SendAndClose(*GetTokensResponse) error
}

type drpcReferralManagerGetTokensStream struct {
	drpc.Stream
}

func (x *drpcReferralManagerGetTokensStream) SendAndClose(m *GetTokensResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCReferralManager_RedeemTokenStream interface {
	drpc.Stream
	SendAndClose(*RedeemTokenResponse) error
}

type drpcReferralManagerRedeemTokenStream struct {
	drpc.Stream
}

func (x *drpcReferralManagerRedeemTokenStream) SendAndClose(m *RedeemTokenResponse) error {
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

// ReferralManagerClient is the client API for ReferralManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReferralManagerClient interface {
	// GetTokens retrieves a list of unredeemed tokens for a user
	GetTokens(ctx context.Context, in *GetTokensRequest, opts ...grpc.CallOption) (*GetTokensResponse, error)
	// RedeemToken saves newly created user info in referral manager
	RedeemToken(ctx context.Context, in *RedeemTokenRequest, opts ...grpc.CallOption) (*RedeemTokenResponse, error)
}

type referralManagerClient struct {
	cc *grpc.ClientConn
}

func NewReferralManagerClient(cc *grpc.ClientConn) ReferralManagerClient {
	return &referralManagerClient{cc}
}

func (c *referralManagerClient) GetTokens(ctx context.Context, in *GetTokensRequest, opts ...grpc.CallOption) (*GetTokensResponse, error) {
	out := new(GetTokensResponse)
	err := c.cc.Invoke(ctx, "/referralmanager.ReferralManager/GetTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *referralManagerClient) RedeemToken(ctx context.Context, in *RedeemTokenRequest, opts ...grpc.CallOption) (*RedeemTokenResponse, error) {
	out := new(RedeemTokenResponse)
	err := c.cc.Invoke(ctx, "/referralmanager.ReferralManager/RedeemToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReferralManagerServer is the server API for ReferralManager service.
type ReferralManagerServer interface {
	// GetTokens retrieves a list of unredeemed tokens for a user
	GetTokens(context.Context, *GetTokensRequest) (*GetTokensResponse, error)
	// RedeemToken saves newly created user info in referral manager
	RedeemToken(context.Context, *RedeemTokenRequest) (*RedeemTokenResponse, error)
}

func RegisterReferralManagerServer(s *grpc.Server, srv ReferralManagerServer) {
	s.RegisterService(&_ReferralManager_serviceDesc, srv)
}

func _ReferralManager_GetTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferralManagerServer).GetTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/referralmanager.ReferralManager/GetTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferralManagerServer).GetTokens(ctx, req.(*GetTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReferralManager_RedeemToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedeemTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReferralManagerServer).RedeemToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/referralmanager.ReferralManager/RedeemToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReferralManagerServer).RedeemToken(ctx, req.(*RedeemTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReferralManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "referralmanager.ReferralManager",
	HandlerType: (*ReferralManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTokens",
			Handler:    _ReferralManager_GetTokens_Handler,
		},
		{
			MethodName: "RedeemToken",
			Handler:    _ReferralManager_RedeemToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "referralmanager.proto",
}

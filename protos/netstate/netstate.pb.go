// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netstate.proto

/*
Package netstate is a generated protocol buffer package.

It is generated from these files:
	netstate.proto

It has these top-level messages:
	RedundancyScheme
	EncryptionScheme
	RemotePiece
	RemoteSegment
	Pointer
	PutRequest
	GetRequest
	ListRequest
	PutResponse
	GetResponse
	ListResponse
	DeleteRequest
	DeleteResponse
*/
package netstate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type RedundancyScheme_SchemeType int32

const (
	RedundancyScheme_RS RedundancyScheme_SchemeType = 0
)

var RedundancyScheme_SchemeType_name = map[int32]string{
	0: "RS",
}
var RedundancyScheme_SchemeType_value = map[string]int32{
	"RS": 0,
}

func (x RedundancyScheme_SchemeType) String() string {
	return proto.EnumName(RedundancyScheme_SchemeType_name, int32(x))
}
func (RedundancyScheme_SchemeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type EncryptionScheme_EncryptionType int32

const (
	EncryptionScheme_AESGCM    EncryptionScheme_EncryptionType = 0
	EncryptionScheme_SECRETBOX EncryptionScheme_EncryptionType = 1
)

var EncryptionScheme_EncryptionType_name = map[int32]string{
	0: "AESGCM",
	1: "SECRETBOX",
}
var EncryptionScheme_EncryptionType_value = map[string]int32{
	"AESGCM":    0,
	"SECRETBOX": 1,
}

func (x EncryptionScheme_EncryptionType) String() string {
	return proto.EnumName(EncryptionScheme_EncryptionType_name, int32(x))
}
func (EncryptionScheme_EncryptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type Pointer_DataType int32

const (
	Pointer_INLINE Pointer_DataType = 0
	Pointer_REMOTE Pointer_DataType = 1
)

var Pointer_DataType_name = map[int32]string{
	0: "INLINE",
	1: "REMOTE",
}
var Pointer_DataType_value = map[string]int32{
	"INLINE": 0,
	"REMOTE": 1,
}

func (x Pointer_DataType) String() string {
	return proto.EnumName(Pointer_DataType_name, int32(x))
}
func (Pointer_DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type RedundancyScheme struct {
	Type RedundancyScheme_SchemeType `protobuf:"varint,1,opt,name=type,enum=netstate.RedundancyScheme_SchemeType" json:"type,omitempty"`
	// these values apply to RS encoding
	MinReq           int64 `protobuf:"varint,2,opt,name=min_req,json=minReq" json:"min_req,omitempty"`
	Total            int64 `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	RepairThreshold  int64 `protobuf:"varint,4,opt,name=repair_threshold,json=repairThreshold" json:"repair_threshold,omitempty"`
	SuccessThreshold int64 `protobuf:"varint,5,opt,name=success_threshold,json=successThreshold" json:"success_threshold,omitempty"`
}

func (m *RedundancyScheme) Reset()                    { *m = RedundancyScheme{} }
func (m *RedundancyScheme) String() string            { return proto.CompactTextString(m) }
func (*RedundancyScheme) ProtoMessage()               {}
func (*RedundancyScheme) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RedundancyScheme) GetType() RedundancyScheme_SchemeType {
	if m != nil {
		return m.Type
	}
	return RedundancyScheme_RS
}

func (m *RedundancyScheme) GetMinReq() int64 {
	if m != nil {
		return m.MinReq
	}
	return 0
}

func (m *RedundancyScheme) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RedundancyScheme) GetRepairThreshold() int64 {
	if m != nil {
		return m.RepairThreshold
	}
	return 0
}

func (m *RedundancyScheme) GetSuccessThreshold() int64 {
	if m != nil {
		return m.SuccessThreshold
	}
	return 0
}

type EncryptionScheme struct {
	Type                   EncryptionScheme_EncryptionType `protobuf:"varint,1,opt,name=type,enum=netstate.EncryptionScheme_EncryptionType" json:"type,omitempty"`
	EncryptedEncryptionKey []byte                          `protobuf:"bytes,2,opt,name=encrypted_encryption_key,json=encryptedEncryptionKey,proto3" json:"encrypted_encryption_key,omitempty"`
	EncryptedStartingNonce []byte                          `protobuf:"bytes,3,opt,name=encrypted_starting_nonce,json=encryptedStartingNonce,proto3" json:"encrypted_starting_nonce,omitempty"`
}

func (m *EncryptionScheme) Reset()                    { *m = EncryptionScheme{} }
func (m *EncryptionScheme) String() string            { return proto.CompactTextString(m) }
func (*EncryptionScheme) ProtoMessage()               {}
func (*EncryptionScheme) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EncryptionScheme) GetType() EncryptionScheme_EncryptionType {
	if m != nil {
		return m.Type
	}
	return EncryptionScheme_AESGCM
}

func (m *EncryptionScheme) GetEncryptedEncryptionKey() []byte {
	if m != nil {
		return m.EncryptedEncryptionKey
	}
	return nil
}

func (m *EncryptionScheme) GetEncryptedStartingNonce() []byte {
	if m != nil {
		return m.EncryptedStartingNonce
	}
	return nil
}

type RemotePiece struct {
	PieceNum int64  `protobuf:"varint,1,opt,name=piece_num,json=pieceNum" json:"piece_num,omitempty"`
	NodeId   string `protobuf:"bytes,2,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Size     int64  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
}

func (m *RemotePiece) Reset()                    { *m = RemotePiece{} }
func (m *RemotePiece) String() string            { return proto.CompactTextString(m) }
func (*RemotePiece) ProtoMessage()               {}
func (*RemotePiece) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RemotePiece) GetPieceNum() int64 {
	if m != nil {
		return m.PieceNum
	}
	return 0
}

func (m *RemotePiece) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *RemotePiece) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type RemoteSegment struct {
	Redundancy   *RedundancyScheme `protobuf:"bytes,1,opt,name=redundancy" json:"redundancy,omitempty"`
	PieceId      string            `protobuf:"bytes,2,opt,name=piece_id,json=pieceId" json:"piece_id,omitempty"`
	RemotePieces []*RemotePiece    `protobuf:"bytes,3,rep,name=remote_pieces,json=remotePieces" json:"remote_pieces,omitempty"`
	MerkleRoot   []byte            `protobuf:"bytes,4,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	MerkleSize   int64             `protobuf:"varint,5,opt,name=merkle_size,json=merkleSize" json:"merkle_size,omitempty"`
}

func (m *RemoteSegment) Reset()                    { *m = RemoteSegment{} }
func (m *RemoteSegment) String() string            { return proto.CompactTextString(m) }
func (*RemoteSegment) ProtoMessage()               {}
func (*RemoteSegment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RemoteSegment) GetRedundancy() *RedundancyScheme {
	if m != nil {
		return m.Redundancy
	}
	return nil
}

func (m *RemoteSegment) GetPieceId() string {
	if m != nil {
		return m.PieceId
	}
	return ""
}

func (m *RemoteSegment) GetRemotePieces() []*RemotePiece {
	if m != nil {
		return m.RemotePieces
	}
	return nil
}

func (m *RemoteSegment) GetMerkleRoot() []byte {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

func (m *RemoteSegment) GetMerkleSize() int64 {
	if m != nil {
		return m.MerkleSize
	}
	return 0
}

type Pointer struct {
	Type                     Pointer_DataType           `protobuf:"varint,1,opt,name=type,enum=netstate.Pointer_DataType" json:"type,omitempty"`
	Encryption               *EncryptionScheme          `protobuf:"bytes,2,opt,name=encryption" json:"encryption,omitempty"`
	InlineSegment            []byte                     `protobuf:"bytes,3,opt,name=inline_segment,json=inlineSegment,proto3" json:"inline_segment,omitempty"`
	Remote                   *RemoteSegment             `protobuf:"bytes,4,opt,name=remote" json:"remote,omitempty"`
	EncryptedUnencryptedSize []byte                     `protobuf:"bytes,5,opt,name=encrypted_unencrypted_size,json=encryptedUnencryptedSize,proto3" json:"encrypted_unencrypted_size,omitempty"`
	CreationDate             *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=creation_date,json=creationDate" json:"creation_date,omitempty"`
	ExpirationDate           *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=expiration_date,json=expirationDate" json:"expiration_date,omitempty"`
}

func (m *Pointer) Reset()                    { *m = Pointer{} }
func (m *Pointer) String() string            { return proto.CompactTextString(m) }
func (*Pointer) ProtoMessage()               {}
func (*Pointer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Pointer) GetType() Pointer_DataType {
	if m != nil {
		return m.Type
	}
	return Pointer_INLINE
}

func (m *Pointer) GetEncryption() *EncryptionScheme {
	if m != nil {
		return m.Encryption
	}
	return nil
}

func (m *Pointer) GetInlineSegment() []byte {
	if m != nil {
		return m.InlineSegment
	}
	return nil
}

func (m *Pointer) GetRemote() *RemoteSegment {
	if m != nil {
		return m.Remote
	}
	return nil
}

func (m *Pointer) GetEncryptedUnencryptedSize() []byte {
	if m != nil {
		return m.EncryptedUnencryptedSize
	}
	return nil
}

func (m *Pointer) GetCreationDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationDate
	}
	return nil
}

func (m *Pointer) GetExpirationDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpirationDate
	}
	return nil
}

// PutRequest is a request message for the Put rpc call
type PutRequest struct {
	Path    []byte   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Pointer *Pointer `protobuf:"bytes,2,opt,name=pointer" json:"pointer,omitempty"`
	APIKey  []byte   `protobuf:"bytes,3,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PutRequest) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *PutRequest) GetPointer() *Pointer {
	if m != nil {
		return m.Pointer
	}
	return nil
}

func (m *PutRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// GetRequest is a request message for the Get rpc call
type GetRequest struct {
	Path   []byte `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	APIKey []byte `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetRequest) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *GetRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// ListRequest is a request message for the List rpc call
type ListRequest struct {
	StartingPathKey []byte `protobuf:"bytes,1,opt,name=starting_path_key,json=startingPathKey,proto3" json:"starting_path_key,omitempty"`
	Limit           int64  `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	APIKey          []byte `protobuf:"bytes,3,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListRequest) GetStartingPathKey() []byte {
	if m != nil {
		return m.StartingPathKey
	}
	return nil
}

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// PutResponse is a response message for the Put rpc call
type PutResponse struct {
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// GetResponse is a response message for the Get rpc call
type GetResponse struct {
	Pointer []byte `protobuf:"bytes,1,opt,name=pointer,proto3" json:"pointer,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetResponse) GetPointer() []byte {
	if m != nil {
		return m.Pointer
	}
	return nil
}

// ListResponse is a response message for the List rpc call
type ListResponse struct {
	Paths     [][]byte `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	Truncated bool     `protobuf:"varint,2,opt,name=truncated" json:"truncated,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListResponse) GetPaths() [][]byte {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *ListResponse) GetTruncated() bool {
	if m != nil {
		return m.Truncated
	}
	return false
}

type DeleteRequest struct {
	Path   []byte `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	APIKey []byte `protobuf:"bytes,2,opt,name=APIKey,proto3" json:"APIKey,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeleteRequest) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *DeleteRequest) GetAPIKey() []byte {
	if m != nil {
		return m.APIKey
	}
	return nil
}

// DeleteResponse is a response message for the Delete rpc call
type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*RedundancyScheme)(nil), "netstate.RedundancyScheme")
	proto.RegisterType((*EncryptionScheme)(nil), "netstate.EncryptionScheme")
	proto.RegisterType((*RemotePiece)(nil), "netstate.RemotePiece")
	proto.RegisterType((*RemoteSegment)(nil), "netstate.RemoteSegment")
	proto.RegisterType((*Pointer)(nil), "netstate.Pointer")
	proto.RegisterType((*PutRequest)(nil), "netstate.PutRequest")
	proto.RegisterType((*GetRequest)(nil), "netstate.GetRequest")
	proto.RegisterType((*ListRequest)(nil), "netstate.ListRequest")
	proto.RegisterType((*PutResponse)(nil), "netstate.PutResponse")
	proto.RegisterType((*GetResponse)(nil), "netstate.GetResponse")
	proto.RegisterType((*ListResponse)(nil), "netstate.ListResponse")
	proto.RegisterType((*DeleteRequest)(nil), "netstate.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "netstate.DeleteResponse")
	proto.RegisterEnum("netstate.RedundancyScheme_SchemeType", RedundancyScheme_SchemeType_name, RedundancyScheme_SchemeType_value)
	proto.RegisterEnum("netstate.EncryptionScheme_EncryptionType", EncryptionScheme_EncryptionType_name, EncryptionScheme_EncryptionType_value)
	proto.RegisterEnum("netstate.Pointer_DataType", Pointer_DataType_name, Pointer_DataType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NetState service

type NetStateClient interface {
	// Put formats and hands off a file path to be saved to boltdb
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// Get formats and hands off a file path to get a small value from boltdb
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// List calls the bolt client's List function and returns all file paths
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Delete formats and hands off a file path to delete from boltdb
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type netStateClient struct {
	cc *grpc.ClientConn
}

func NewNetStateClient(cc *grpc.ClientConn) NetStateClient {
	return &netStateClient{cc}
}

func (c *netStateClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/netstate.NetState/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netStateClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/netstate.NetState/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netStateClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/netstate.NetState/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netStateClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/netstate.NetState/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetState service

type NetStateServer interface {
	// Put formats and hands off a file path to be saved to boltdb
	Put(context.Context, *PutRequest) (*PutResponse, error)
	// Get formats and hands off a file path to get a small value from boltdb
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// List calls the bolt client's List function and returns all file paths
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Delete formats and hands off a file path to delete from boltdb
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterNetStateServer(s *grpc.Server, srv NetStateServer) {
	s.RegisterService(&_NetState_serviceDesc, srv)
}

func _NetState_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetStateServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netstate.NetState/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetStateServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetState_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetStateServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netstate.NetState/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetStateServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetState_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetStateServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netstate.NetState/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetStateServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetState_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetStateServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/netstate.NetState/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetStateServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetState_serviceDesc = grpc.ServiceDesc{
	ServiceName: "netstate.NetState",
	HandlerType: (*NetStateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _NetState_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _NetState_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NetState_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NetState_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "netstate.proto",
}

func init() { proto.RegisterFile("netstate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 880 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x8b, 0xdb, 0x46,
	0x10, 0x3f, 0x9d, 0xee, 0x64, 0xdf, 0xf8, 0xcf, 0xe9, 0x16, 0xe7, 0xa2, 0xba, 0x85, 0x1e, 0x82,
	0x50, 0xa7, 0x07, 0x2e, 0xb8, 0x14, 0xd2, 0x4b, 0x4b, 0x49, 0xee, 0xcc, 0x61, 0x92, 0x38, 0x66,
	0xed, 0xd2, 0xbe, 0x09, 0x45, 0x9a, 0xda, 0x22, 0xd6, 0x4a, 0x27, 0xad, 0xa0, 0xee, 0xf7, 0xea,
	0x37, 0x2a, 0x7d, 0xe8, 0x63, 0x3f, 0x41, 0xd1, 0xee, 0x4a, 0x5a, 0x3b, 0xa4, 0x85, 0x3e, 0x69,
	0x67, 0xe6, 0x37, 0xb3, 0xf3, 0x9b, 0xf9, 0x69, 0xa1, 0xcf, 0x90, 0xe7, 0xdc, 0xe7, 0x38, 0x4e,
	0xb3, 0x84, 0x27, 0xa4, 0x5d, 0xd9, 0xc3, 0x73, 0x1e, 0xc5, 0x98, 0x73, 0x3f, 0x4e, 0x65, 0xc8,
	0xfd, 0xcb, 0x00, 0x9b, 0x62, 0x58, 0xb0, 0xd0, 0x67, 0xc1, 0x6e, 0x19, 0x6c, 0x30, 0x46, 0xf2,
	0x2d, 0x9c, 0xf0, 0x5d, 0x8a, 0x8e, 0x71, 0x65, 0x8c, 0xfa, 0x93, 0x27, 0xe3, 0xba, 0xdc, 0x21,
	0x72, 0x2c, 0x3f, 0xab, 0x5d, 0x8a, 0x54, 0xa4, 0x90, 0xc7, 0xd0, 0x8a, 0x23, 0xe6, 0x65, 0xf8,
	0xe0, 0x1c, 0x5f, 0x19, 0x23, 0x93, 0x5a, 0x71, 0xc4, 0x28, 0x3e, 0x90, 0x01, 0x9c, 0xf2, 0x84,
	0xfb, 0x5b, 0xc7, 0x14, 0x6e, 0x69, 0x90, 0xa7, 0x60, 0x67, 0x98, 0xfa, 0x51, 0xe6, 0xf1, 0x4d,
	0x86, 0xf9, 0x26, 0xd9, 0x86, 0xce, 0x89, 0x00, 0x9c, 0x4b, 0xff, 0xaa, 0x72, 0x93, 0x6b, 0xb8,
	0xc8, 0x8b, 0x20, 0xc0, 0x3c, 0xd7, 0xb0, 0xa7, 0x02, 0x6b, 0xab, 0x40, 0x0d, 0x76, 0x07, 0x00,
	0x4d, 0x6b, 0xc4, 0x82, 0x63, 0xba, 0xb4, 0x8f, 0xdc, 0xbf, 0x0d, 0xb0, 0xa7, 0x2c, 0xc8, 0x76,
	0x29, 0x8f, 0x12, 0xa6, 0xc8, 0x7e, 0xbf, 0x47, 0xf6, 0x69, 0x43, 0xf6, 0x10, 0xa9, 0x39, 0x34,
	0xc2, 0xcf, 0xc0, 0x41, 0xe9, 0xc7, 0xd0, 0xc3, 0x1a, 0xe1, 0xbd, 0xc7, 0x9d, 0x98, 0x40, 0x97,
	0x5e, 0xd6, 0xf1, 0xa6, 0xc0, 0x2b, 0xdc, 0xed, 0x67, 0xe6, 0xdc, 0xcf, 0x78, 0xc4, 0xd6, 0x1e,
	0x4b, 0x58, 0x80, 0x62, 0x48, 0x7a, 0xe6, 0x52, 0x85, 0xe7, 0x65, 0xd4, 0xbd, 0x86, 0xfe, 0x7e,
	0x2f, 0x04, 0xc0, 0x7a, 0x31, 0x5d, 0xde, 0xdf, 0xbe, 0xb1, 0x8f, 0x48, 0x0f, 0xce, 0x96, 0xd3,
	0x5b, 0x3a, 0x5d, 0xbd, 0x7c, 0xfb, 0xb3, 0x6d, 0xb8, 0x3f, 0x41, 0x87, 0x62, 0x9c, 0x70, 0x5c,
	0x44, 0x18, 0x20, 0xf9, 0x14, 0xce, 0xd2, 0xf2, 0xe0, 0xb1, 0x22, 0x16, 0x9c, 0x4d, 0xda, 0x16,
	0x8e, 0x79, 0x11, 0x97, 0xdb, 0x63, 0x49, 0x88, 0x5e, 0x14, 0x8a, 0xde, 0xcf, 0xa8, 0x55, 0x9a,
	0xb3, 0x90, 0x10, 0x38, 0xc9, 0xa3, 0xdf, 0x50, 0x2d, 0x4f, 0x9c, 0xdd, 0x3f, 0x0c, 0xe8, 0xc9,
	0xca, 0x4b, 0x5c, 0xc7, 0xc8, 0x38, 0xb9, 0x01, 0xc8, 0x6a, 0x85, 0x88, 0xe2, 0x9d, 0xc9, 0xf0,
	0xe3, 0xea, 0xa1, 0x1a, 0x9a, 0x7c, 0x02, 0xb2, 0x8d, 0xe6, 0xee, 0x96, 0xb0, 0x67, 0x21, 0xb9,
	0x81, 0x5e, 0x26, 0xee, 0xf1, 0x84, 0x27, 0x77, 0xcc, 0x2b, 0x73, 0xd4, 0x99, 0x3c, 0xd2, 0x2b,
	0xd7, 0x04, 0x69, 0x37, 0x6b, 0x8c, 0x9c, 0x7c, 0x0e, 0x9d, 0x18, 0xb3, 0xf7, 0x5b, 0xf4, 0xb2,
	0x24, 0xe1, 0x42, 0x5b, 0x5d, 0x0a, 0xd2, 0x45, 0x93, 0x84, 0x6b, 0x00, 0x41, 0x50, 0x0a, 0x4a,
	0x01, 0x96, 0x25, 0xcd, 0xdf, 0x4d, 0x68, 0x2d, 0x92, 0x88, 0x71, 0xcc, 0xc8, 0x78, 0x4f, 0x2b,
	0x1a, 0x35, 0x05, 0x18, 0xdf, 0xf9, 0xdc, 0xd7, 0xc4, 0x71, 0x03, 0xd0, 0x48, 0x42, 0xd0, 0xda,
	0x1b, 0xc8, 0xa1, 0xc2, 0xa8, 0x86, 0x26, 0x4f, 0xa0, 0x1f, 0xb1, 0x6d, 0xc4, 0xd0, 0xcb, 0xe5,
	0x78, 0x95, 0x28, 0x7a, 0xd2, 0x5b, 0xcd, 0xfc, 0x2b, 0xb0, 0x24, 0x61, 0xc1, 0xad, 0x33, 0x79,
	0x7c, 0x38, 0x15, 0x05, 0xa4, 0x0a, 0x46, 0xbe, 0x83, 0x61, 0x23, 0xbb, 0x82, 0x69, 0x12, 0xac,
	0xf8, 0x77, 0x69, 0x23, 0xcc, 0x1f, 0x1b, 0x40, 0x39, 0x0d, 0xf2, 0x03, 0xf4, 0x82, 0x0c, 0x7d,
	0x21, 0xf1, 0xd0, 0xe7, 0xe8, 0x58, 0x8a, 0xd4, 0x3a, 0x49, 0xd6, 0x5b, 0xf5, 0xe0, 0xbc, 0x2b,
	0x7e, 0x19, 0xaf, 0xaa, 0x87, 0x86, 0x76, 0xab, 0x84, 0x3b, 0x9f, 0x23, 0xb9, 0x85, 0x73, 0xfc,
	0x35, 0x8d, 0x32, 0xad, 0x44, 0xeb, 0x3f, 0x4b, 0xf4, 0x9b, 0x94, 0xb2, 0x88, 0xeb, 0x42, 0xbb,
	0x9a, 0x74, 0x29, 0xfd, 0xd9, 0xfc, 0xf5, 0x6c, 0x3e, 0xb5, 0x8f, 0xca, 0x33, 0x9d, 0xbe, 0x79,
	0xbb, 0x9a, 0xda, 0x86, 0x8b, 0x00, 0x8b, 0x82, 0x53, 0x7c, 0x28, 0x30, 0xe7, 0xa5, 0x80, 0x53,
	0x9f, 0x6f, 0xc4, 0xe6, 0xba, 0x54, 0x9c, 0xc9, 0x35, 0xb4, 0x52, 0xb9, 0x37, 0xb5, 0x9a, 0x8b,
	0x0f, 0x16, 0x4a, 0x2b, 0x04, 0xb9, 0x04, 0xeb, 0xc5, 0x62, 0xf6, 0x0a, 0x77, 0x6a, 0x0d, 0xca,
	0x72, 0x9f, 0x01, 0xdc, 0xe3, 0xbf, 0x5e, 0xd3, 0x64, 0x1e, 0xef, 0x65, 0xae, 0xa1, 0xf3, 0x3a,
	0xca, 0xeb, 0xd4, 0x2f, 0xe1, 0xa2, 0x7e, 0x04, 0xca, 0x3c, 0xf1, 0x82, 0xc8, 0x3a, 0xe7, 0x55,
	0x60, 0xe1, 0xf3, 0x4d, 0xf9, 0x74, 0x0c, 0xe0, 0x74, 0x1b, 0xc5, 0x11, 0x57, 0x6f, 0xac, 0x34,
	0x3e, 0xda, 0x62, 0x0f, 0x3a, 0x62, 0x12, 0x79, 0x9a, 0xb0, 0x1c, 0xdd, 0x2f, 0xa0, 0x23, 0x3a,
	0x96, 0x26, 0x71, 0x9a, 0x29, 0xc8, 0xdb, 0x2a, 0xd3, 0x7d, 0x09, 0x5d, 0xd9, 0xa0, 0x42, 0x0e,
	0xe0, 0xb4, 0x6c, 0x2c, 0x77, 0x8c, 0x2b, 0x73, 0xd4, 0xa5, 0xd2, 0x20, 0x9f, 0xc1, 0x19, 0xcf,
	0x0a, 0x16, 0xf8, 0x1c, 0xe5, 0x9f, 0xdb, 0xa6, 0x8d, 0xc3, 0x7d, 0x0e, 0xbd, 0x3b, 0xdc, 0x22,
	0xc7, 0xff, 0x33, 0x21, 0x1b, 0xfa, 0x55, 0xb2, 0x6c, 0x61, 0xf2, 0xa7, 0x01, 0xed, 0x39, 0xf2,
	0x65, 0xb9, 0x23, 0x32, 0x01, 0x73, 0x51, 0x70, 0x32, 0xd0, 0xb6, 0x56, 0x2f, 0x7c, 0xf8, 0xe8,
	0xc0, 0xab, 0x38, 0x4c, 0xc0, 0xbc, 0xc7, 0xbd, 0x9c, 0x66, 0x7b, 0x7a, 0x8e, 0x3e, 0xa1, 0x6f,
	0xe0, 0xa4, 0x9c, 0x03, 0xd1, 0xc2, 0xda, 0xe2, 0x86, 0x97, 0x87, 0x6e, 0x95, 0xf6, 0x1c, 0x2c,
	0xd9, 0x3d, 0xd1, 0xfe, 0xc9, 0xbd, 0x61, 0x0c, 0x9d, 0x0f, 0x03, 0x32, 0xf9, 0x9d, 0x25, 0xfe,
	0x82, 0xaf, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x7b, 0x10, 0x5b, 0xcb, 0x07, 0x00, 0x00,
}

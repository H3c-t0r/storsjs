// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: inspector.proto

package pb

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// GetStats
type GetStatsRequest struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsRequest) Reset()         { *m = GetStatsRequest{} }
func (m *GetStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatsRequest) ProtoMessage()    {}
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{0}
}
func (m *GetStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsRequest.Unmarshal(m, b)
}
func (m *GetStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsRequest.Marshal(b, m, deterministic)
}
func (dst *GetStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsRequest.Merge(dst, src)
}
func (m *GetStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatsRequest.Size(m)
}
func (m *GetStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsRequest proto.InternalMessageInfo

func (m *GetStatsRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type GetStatsResponse struct {
	AuditCount           int64    `protobuf:"varint,1,opt,name=audit_count,json=auditCount,proto3" json:"audit_count,omitempty"`
	UptimeRatio          float64  `protobuf:"fixed64,2,opt,name=uptime_ratio,json=uptimeRatio,proto3" json:"uptime_ratio,omitempty"`
	AuditRatio           float64  `protobuf:"fixed64,3,opt,name=audit_ratio,json=auditRatio,proto3" json:"audit_ratio,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsResponse) Reset()         { *m = GetStatsResponse{} }
func (m *GetStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatsResponse) ProtoMessage()    {}
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{1}
}
func (m *GetStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsResponse.Unmarshal(m, b)
}
func (m *GetStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsResponse.Marshal(b, m, deterministic)
}
func (dst *GetStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsResponse.Merge(dst, src)
}
func (m *GetStatsResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatsResponse.Size(m)
}
func (m *GetStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsResponse proto.InternalMessageInfo

func (m *GetStatsResponse) GetAuditCount() int64 {
	if m != nil {
		return m.AuditCount
	}
	return 0
}

func (m *GetStatsResponse) GetUptimeRatio() float64 {
	if m != nil {
		return m.UptimeRatio
	}
	return 0
}

func (m *GetStatsResponse) GetAuditRatio() float64 {
	if m != nil {
		return m.AuditRatio
	}
	return 0
}

// CreateStats
type CreateStatsRequest struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	AuditCount           int64    `protobuf:"varint,2,opt,name=audit_count,json=auditCount,proto3" json:"audit_count,omitempty"`
	AuditSuccessCount    int64    `protobuf:"varint,3,opt,name=audit_success_count,json=auditSuccessCount,proto3" json:"audit_success_count,omitempty"`
	UptimeCount          int64    `protobuf:"varint,4,opt,name=uptime_count,json=uptimeCount,proto3" json:"uptime_count,omitempty"`
	UptimeSuccessCount   int64    `protobuf:"varint,5,opt,name=uptime_success_count,json=uptimeSuccessCount,proto3" json:"uptime_success_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateStatsRequest) Reset()         { *m = CreateStatsRequest{} }
func (m *CreateStatsRequest) String() string { return proto.CompactTextString(m) }
func (*CreateStatsRequest) ProtoMessage()    {}
func (*CreateStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{2}
}
func (m *CreateStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStatsRequest.Unmarshal(m, b)
}
func (m *CreateStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStatsRequest.Marshal(b, m, deterministic)
}
func (dst *CreateStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStatsRequest.Merge(dst, src)
}
func (m *CreateStatsRequest) XXX_Size() int {
	return xxx_messageInfo_CreateStatsRequest.Size(m)
}
func (m *CreateStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStatsRequest proto.InternalMessageInfo

func (m *CreateStatsRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *CreateStatsRequest) GetAuditCount() int64 {
	if m != nil {
		return m.AuditCount
	}
	return 0
}

func (m *CreateStatsRequest) GetAuditSuccessCount() int64 {
	if m != nil {
		return m.AuditSuccessCount
	}
	return 0
}

func (m *CreateStatsRequest) GetUptimeCount() int64 {
	if m != nil {
		return m.UptimeCount
	}
	return 0
}

func (m *CreateStatsRequest) GetUptimeSuccessCount() int64 {
	if m != nil {
		return m.UptimeSuccessCount
	}
	return 0
}

type CreateStatsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateStatsResponse) Reset()         { *m = CreateStatsResponse{} }
func (m *CreateStatsResponse) String() string { return proto.CompactTextString(m) }
func (*CreateStatsResponse) ProtoMessage()    {}
func (*CreateStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{3}
}
func (m *CreateStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStatsResponse.Unmarshal(m, b)
}
func (m *CreateStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStatsResponse.Marshal(b, m, deterministic)
}
func (dst *CreateStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStatsResponse.Merge(dst, src)
}
func (m *CreateStatsResponse) XXX_Size() int {
	return xxx_messageInfo_CreateStatsResponse.Size(m)
}
func (m *CreateStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStatsResponse proto.InternalMessageInfo

// CountNodes
type CountNodesResponse struct {
	Kademlia             int64    `protobuf:"varint,1,opt,name=kademlia,proto3" json:"kademlia,omitempty"`
	Overlay              int64    `protobuf:"varint,2,opt,name=overlay,proto3" json:"overlay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountNodesResponse) Reset()         { *m = CountNodesResponse{} }
func (m *CountNodesResponse) String() string { return proto.CompactTextString(m) }
func (*CountNodesResponse) ProtoMessage()    {}
func (*CountNodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{4}
}
func (m *CountNodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountNodesResponse.Unmarshal(m, b)
}
func (m *CountNodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountNodesResponse.Marshal(b, m, deterministic)
}
func (dst *CountNodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountNodesResponse.Merge(dst, src)
}
func (m *CountNodesResponse) XXX_Size() int {
	return xxx_messageInfo_CountNodesResponse.Size(m)
}
func (m *CountNodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountNodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountNodesResponse proto.InternalMessageInfo

func (m *CountNodesResponse) GetKademlia() int64 {
	if m != nil {
		return m.Kademlia
	}
	return 0
}

func (m *CountNodesResponse) GetOverlay() int64 {
	if m != nil {
		return m.Overlay
	}
	return 0
}

type CountNodesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountNodesRequest) Reset()         { *m = CountNodesRequest{} }
func (m *CountNodesRequest) String() string { return proto.CompactTextString(m) }
func (*CountNodesRequest) ProtoMessage()    {}
func (*CountNodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{5}
}
func (m *CountNodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountNodesRequest.Unmarshal(m, b)
}
func (m *CountNodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountNodesRequest.Marshal(b, m, deterministic)
}
func (dst *CountNodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountNodesRequest.Merge(dst, src)
}
func (m *CountNodesRequest) XXX_Size() int {
	return xxx_messageInfo_CountNodesRequest.Size(m)
}
func (m *CountNodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CountNodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CountNodesRequest proto.InternalMessageInfo

// GetBuckets
type GetBucketsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBucketsRequest) Reset()         { *m = GetBucketsRequest{} }
func (m *GetBucketsRequest) String() string { return proto.CompactTextString(m) }
func (*GetBucketsRequest) ProtoMessage()    {}
func (*GetBucketsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{6}
}
func (m *GetBucketsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBucketsRequest.Unmarshal(m, b)
}
func (m *GetBucketsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBucketsRequest.Marshal(b, m, deterministic)
}
func (dst *GetBucketsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBucketsRequest.Merge(dst, src)
}
func (m *GetBucketsRequest) XXX_Size() int {
	return xxx_messageInfo_GetBucketsRequest.Size(m)
}
func (m *GetBucketsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBucketsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBucketsRequest proto.InternalMessageInfo

type GetBucketsResponse struct {
	Total                int64    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Ids                  [][]byte `protobuf:"bytes,2,rep,name=ids" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBucketsResponse) Reset()         { *m = GetBucketsResponse{} }
func (m *GetBucketsResponse) String() string { return proto.CompactTextString(m) }
func (*GetBucketsResponse) ProtoMessage()    {}
func (*GetBucketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{7}
}
func (m *GetBucketsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBucketsResponse.Unmarshal(m, b)
}
func (m *GetBucketsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBucketsResponse.Marshal(b, m, deterministic)
}
func (dst *GetBucketsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBucketsResponse.Merge(dst, src)
}
func (m *GetBucketsResponse) XXX_Size() int {
	return xxx_messageInfo_GetBucketsResponse.Size(m)
}
func (m *GetBucketsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBucketsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBucketsResponse proto.InternalMessageInfo

func (m *GetBucketsResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *GetBucketsResponse) GetIds() [][]byte {
	if m != nil {
		return m.Ids
	}
	return nil
}

// GetBucket
type GetBucketRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBucketRequest) Reset()         { *m = GetBucketRequest{} }
func (m *GetBucketRequest) String() string { return proto.CompactTextString(m) }
func (*GetBucketRequest) ProtoMessage()    {}
func (*GetBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{8}
}
func (m *GetBucketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBucketRequest.Unmarshal(m, b)
}
func (m *GetBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBucketRequest.Marshal(b, m, deterministic)
}
func (dst *GetBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBucketRequest.Merge(dst, src)
}
func (m *GetBucketRequest) XXX_Size() int {
	return xxx_messageInfo_GetBucketRequest.Size(m)
}
func (m *GetBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBucketRequest proto.InternalMessageInfo

func (m *GetBucketRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetBucketResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nodes                []*Node  `protobuf:"bytes,2,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBucketResponse) Reset()         { *m = GetBucketResponse{} }
func (m *GetBucketResponse) String() string { return proto.CompactTextString(m) }
func (*GetBucketResponse) ProtoMessage()    {}
func (*GetBucketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{9}
}
func (m *GetBucketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBucketResponse.Unmarshal(m, b)
}
func (m *GetBucketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBucketResponse.Marshal(b, m, deterministic)
}
func (dst *GetBucketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBucketResponse.Merge(dst, src)
}
func (m *GetBucketResponse) XXX_Size() int {
	return xxx_messageInfo_GetBucketResponse.Size(m)
}
func (m *GetBucketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBucketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBucketResponse proto.InternalMessageInfo

func (m *GetBucketResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetBucketResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Bucket struct {
	Nodes                []*Node  `protobuf:"bytes,2,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{10}
}
func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (dst *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(dst, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type BucketList struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BucketList) Reset()         { *m = BucketList{} }
func (m *BucketList) String() string { return proto.CompactTextString(m) }
func (*BucketList) ProtoMessage()    {}
func (*BucketList) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{11}
}
func (m *BucketList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketList.Unmarshal(m, b)
}
func (m *BucketList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketList.Marshal(b, m, deterministic)
}
func (dst *BucketList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketList.Merge(dst, src)
}
func (m *BucketList) XXX_Size() int {
	return xxx_messageInfo_BucketList.Size(m)
}
func (m *BucketList) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketList.DiscardUnknown(m)
}

var xxx_messageInfo_BucketList proto.InternalMessageInfo

func (m *BucketList) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// PingNode
type PingNodeRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingNodeRequest) Reset()         { *m = PingNodeRequest{} }
func (m *PingNodeRequest) String() string { return proto.CompactTextString(m) }
func (*PingNodeRequest) ProtoMessage()    {}
func (*PingNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{12}
}
func (m *PingNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingNodeRequest.Unmarshal(m, b)
}
func (m *PingNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingNodeRequest.Marshal(b, m, deterministic)
}
func (dst *PingNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingNodeRequest.Merge(dst, src)
}
func (m *PingNodeRequest) XXX_Size() int {
	return xxx_messageInfo_PingNodeRequest.Size(m)
}
func (m *PingNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingNodeRequest proto.InternalMessageInfo

func (m *PingNodeRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PingNodeRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type PingNodeResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingNodeResponse) Reset()         { *m = PingNodeResponse{} }
func (m *PingNodeResponse) String() string { return proto.CompactTextString(m) }
func (*PingNodeResponse) ProtoMessage()    {}
func (*PingNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_inspector_bc69ade8473655f5, []int{13}
}
func (m *PingNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingNodeResponse.Unmarshal(m, b)
}
func (m *PingNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingNodeResponse.Marshal(b, m, deterministic)
}
func (dst *PingNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingNodeResponse.Merge(dst, src)
}
func (m *PingNodeResponse) XXX_Size() int {
	return xxx_messageInfo_PingNodeResponse.Size(m)
}
func (m *PingNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingNodeResponse proto.InternalMessageInfo

func (m *PingNodeResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*GetStatsRequest)(nil), "inspector.GetStatsRequest")
	proto.RegisterType((*GetStatsResponse)(nil), "inspector.GetStatsResponse")
	proto.RegisterType((*CreateStatsRequest)(nil), "inspector.CreateStatsRequest")
	proto.RegisterType((*CreateStatsResponse)(nil), "inspector.CreateStatsResponse")
	proto.RegisterType((*CountNodesResponse)(nil), "inspector.CountNodesResponse")
	proto.RegisterType((*CountNodesRequest)(nil), "inspector.CountNodesRequest")
	proto.RegisterType((*GetBucketsRequest)(nil), "inspector.GetBucketsRequest")
	proto.RegisterType((*GetBucketsResponse)(nil), "inspector.GetBucketsResponse")
	proto.RegisterType((*GetBucketRequest)(nil), "inspector.GetBucketRequest")
	proto.RegisterType((*GetBucketResponse)(nil), "inspector.GetBucketResponse")
	proto.RegisterType((*Bucket)(nil), "inspector.Bucket")
	proto.RegisterType((*BucketList)(nil), "inspector.BucketList")
	proto.RegisterType((*PingNodeRequest)(nil), "inspector.PingNodeRequest")
	proto.RegisterType((*PingNodeResponse)(nil), "inspector.PingNodeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InspectorClient is the client API for Inspector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InspectorClient interface {
	// Kad/Overlay commands:
	// CountNodes returns the number of nodes in the cache and in the routing table
	CountNodes(ctx context.Context, in *CountNodesRequest, opts ...grpc.CallOption) (*CountNodesResponse, error)
	// GetBuckets returns the k buckets from a Kademlia instance
	GetBuckets(ctx context.Context, in *GetBucketsRequest, opts ...grpc.CallOption) (*GetBucketsResponse, error)
	// GetBucket returns the details of a single k bucket from the kademlia instance
	GetBucket(ctx context.Context, in *GetBucketRequest, opts ...grpc.CallOption) (*GetBucketResponse, error)
	// PingNodes sends a PING RPC to a node and returns it's availability
	PingNode(ctx context.Context, in *PingNodeRequest, opts ...grpc.CallOption) (*PingNodeResponse, error)
	// StatDB commands:
	// GetStats returns the stats for a particular node ID
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
	// CreateStats creates a node with specified stats
	CreateStats(ctx context.Context, in *CreateStatsRequest, opts ...grpc.CallOption) (*CreateStatsResponse, error)
}

type inspectorClient struct {
	cc *grpc.ClientConn
}

func NewInspectorClient(cc *grpc.ClientConn) InspectorClient {
	return &inspectorClient{cc}
}

func (c *inspectorClient) CountNodes(ctx context.Context, in *CountNodesRequest, opts ...grpc.CallOption) (*CountNodesResponse, error) {
	out := new(CountNodesResponse)
	err := c.cc.Invoke(ctx, "/inspector.Inspector/CountNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inspectorClient) GetBuckets(ctx context.Context, in *GetBucketsRequest, opts ...grpc.CallOption) (*GetBucketsResponse, error) {
	out := new(GetBucketsResponse)
	err := c.cc.Invoke(ctx, "/inspector.Inspector/GetBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inspectorClient) GetBucket(ctx context.Context, in *GetBucketRequest, opts ...grpc.CallOption) (*GetBucketResponse, error) {
	out := new(GetBucketResponse)
	err := c.cc.Invoke(ctx, "/inspector.Inspector/GetBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inspectorClient) PingNode(ctx context.Context, in *PingNodeRequest, opts ...grpc.CallOption) (*PingNodeResponse, error) {
	out := new(PingNodeResponse)
	err := c.cc.Invoke(ctx, "/inspector.Inspector/PingNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inspectorClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, "/inspector.Inspector/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inspectorClient) CreateStats(ctx context.Context, in *CreateStatsRequest, opts ...grpc.CallOption) (*CreateStatsResponse, error) {
	out := new(CreateStatsResponse)
	err := c.cc.Invoke(ctx, "/inspector.Inspector/CreateStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InspectorServer is the server API for Inspector service.
type InspectorServer interface {
	// Kad/Overlay commands:
	// CountNodes returns the number of nodes in the cache and in the routing table
	CountNodes(context.Context, *CountNodesRequest) (*CountNodesResponse, error)
	// GetBuckets returns the k buckets from a Kademlia instance
	GetBuckets(context.Context, *GetBucketsRequest) (*GetBucketsResponse, error)
	// GetBucket returns the details of a single k bucket from the kademlia instance
	GetBucket(context.Context, *GetBucketRequest) (*GetBucketResponse, error)
	// PingNodes sends a PING RPC to a node and returns it's availability
	PingNode(context.Context, *PingNodeRequest) (*PingNodeResponse, error)
	// StatDB commands:
	// GetStats returns the stats for a particular node ID
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	// CreateStats creates a node with specified stats
	CreateStats(context.Context, *CreateStatsRequest) (*CreateStatsResponse, error)
}

func RegisterInspectorServer(s *grpc.Server, srv InspectorServer) {
	s.RegisterService(&_Inspector_serviceDesc, srv)
}

func _Inspector_CountNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InspectorServer).CountNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inspector.Inspector/CountNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InspectorServer).CountNodes(ctx, req.(*CountNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inspector_GetBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBucketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InspectorServer).GetBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inspector.Inspector/GetBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InspectorServer).GetBuckets(ctx, req.(*GetBucketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inspector_GetBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InspectorServer).GetBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inspector.Inspector/GetBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InspectorServer).GetBucket(ctx, req.(*GetBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inspector_PingNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InspectorServer).PingNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inspector.Inspector/PingNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InspectorServer).PingNode(ctx, req.(*PingNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inspector_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InspectorServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inspector.Inspector/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InspectorServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inspector_CreateStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InspectorServer).CreateStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inspector.Inspector/CreateStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InspectorServer).CreateStats(ctx, req.(*CreateStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Inspector_serviceDesc = grpc.ServiceDesc{
	ServiceName: "inspector.Inspector",
	HandlerType: (*InspectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountNodes",
			Handler:    _Inspector_CountNodes_Handler,
		},
		{
			MethodName: "GetBuckets",
			Handler:    _Inspector_GetBuckets_Handler,
		},
		{
			MethodName: "GetBucket",
			Handler:    _Inspector_GetBucket_Handler,
		},
		{
			MethodName: "PingNode",
			Handler:    _Inspector_PingNode_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _Inspector_GetStats_Handler,
		},
		{
			MethodName: "CreateStats",
			Handler:    _Inspector_CreateStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inspector.proto",
}

func init() { proto.RegisterFile("inspector.proto", fileDescriptor_inspector_bc69ade8473655f5) }

var fileDescriptor_inspector_bc69ade8473655f5 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0x92, 0xb5, 0x6b, 0x6f, 0x37, 0xba, 0xb9, 0x43, 0x54, 0xe9, 0x80, 0x62, 0x5e, 0x2a,
	0x24, 0x2a, 0x18, 0x8f, 0xf0, 0xb4, 0x4a, 0x8c, 0xa2, 0x09, 0xa1, 0xec, 0x8d, 0x97, 0x2a, 0x8b,
	0x2d, 0x64, 0xb5, 0x8b, 0x43, 0xec, 0x80, 0xf8, 0x79, 0xfc, 0x06, 0xfe, 0x10, 0xf2, 0x47, 0xe2,
	0x24, 0xcd, 0xd0, 0xde, 0x72, 0xef, 0x39, 0xf7, 0x5c, 0xdf, 0x73, 0x63, 0xc3, 0x98, 0xa5, 0x22,
	0xa3, 0x89, 0xe4, 0xf9, 0x32, 0xcb, 0xb9, 0xe4, 0x68, 0x58, 0x25, 0xc2, 0x63, 0xfe, 0x93, 0xe6,
	0xbb, 0xf8, 0xb7, 0x41, 0xf0, 0x2b, 0x18, 0x5f, 0x51, 0x79, 0x23, 0x63, 0x29, 0x22, 0xfa, 0xa3,
	0xa0, 0x42, 0xa2, 0x27, 0x70, 0x98, 0x72, 0x42, 0x37, 0x8c, 0x4c, 0xbd, 0xb9, 0xb7, 0x18, 0x46,
	0x7d, 0x15, 0xae, 0x09, 0xfe, 0x05, 0x27, 0x8e, 0x2b, 0x32, 0x9e, 0x0a, 0x8a, 0x9e, 0xc3, 0x28,
	0x2e, 0x08, 0x93, 0x9b, 0x84, 0x17, 0xa9, 0xd4, 0x05, 0x41, 0x04, 0x3a, 0xb5, 0x52, 0x19, 0xf4,
	0x02, 0x8e, 0x8a, 0x4c, 0xb2, 0x3b, 0xba, 0xc9, 0x63, 0xc9, 0xf8, 0xd4, 0x9f, 0x7b, 0x0b, 0x2f,
	0x1a, 0x99, 0x5c, 0xa4, 0x52, 0x4e, 0xc3, 0x30, 0x02, 0xcd, 0x30, 0x1a, 0x9a, 0x80, 0xff, 0x7a,
	0x80, 0x56, 0x39, 0x8d, 0x25, 0x7d, 0xd0, 0x41, 0xdb, 0x87, 0xf2, 0xf7, 0x0e, 0xb5, 0x84, 0x89,
	0x21, 0x88, 0x22, 0x49, 0xa8, 0x10, 0x96, 0x18, 0x68, 0xe2, 0xa9, 0x86, 0x6e, 0x0c, 0xd2, 0x1e,
	0xc2, 0x10, 0x0f, 0x34, 0xd1, 0x0e, 0x61, 0x28, 0x6f, 0xe0, 0xcc, 0x52, 0x9a, 0x9a, 0x3d, 0x4d,
	0x45, 0x06, 0xab, 0x8b, 0xe2, 0xc7, 0x30, 0x69, 0x0c, 0x65, 0x1c, 0xc5, 0x9f, 0x01, 0x69, 0xfc,
	0x0b, 0x27, 0xd4, 0xf9, 0x1c, 0xc2, 0x60, 0x1b, 0x13, 0x7a, 0xb7, 0x63, 0xb1, 0x35, 0xb9, 0x8a,
	0xd1, 0x14, 0x0e, 0xed, 0x52, 0xed, 0xa8, 0x65, 0x88, 0x27, 0x70, 0x5a, 0xd7, 0xd2, 0xb6, 0xa9,
	0xe4, 0x15, 0x95, 0x97, 0x45, 0xb2, 0xa5, 0x95, 0x97, 0xf8, 0x03, 0xa0, 0x7a, 0xd2, 0x76, 0x3d,
	0x83, 0x9e, 0xe4, 0x32, 0xde, 0xd9, 0x96, 0x26, 0x40, 0x27, 0x10, 0x30, 0x22, 0xa6, 0xfe, 0x3c,
	0x58, 0x1c, 0x45, 0xea, 0x13, 0x63, 0xfd, 0x67, 0x98, 0xea, 0x72, 0x3b, 0x8f, 0xc0, 0xaf, 0x16,
	0xe3, 0x33, 0x82, 0x3f, 0xd5, 0xda, 0x56, 0x0d, 0x5a, 0x24, 0xf4, 0x12, 0x7a, 0x6a, 0x87, 0x46,
	0x7c, 0x74, 0x71, 0xbc, 0x2c, 0xff, 0x56, 0x35, 0x41, 0x64, 0x30, 0xfc, 0x1a, 0xfa, 0x46, 0xe6,
	0x61, 0xf4, 0xb7, 0x00, 0x86, 0x7e, 0xcd, 0x44, 0xad, 0xc4, 0xfb, 0x4f, 0xc9, 0x7b, 0x18, 0x7f,
	0x65, 0xe9, 0x77, 0x9d, 0xea, 0x1e, 0x47, 0x99, 0x1e, 0x13, 0x92, 0x53, 0x21, 0xb4, 0xe9, 0xc3,
	0xa8, 0x0c, 0x95, 0x19, 0xae, 0xd8, 0xcd, 0xc9, 0xb7, 0xba, 0x7a, 0x10, 0xf9, 0x7c, 0x7b, 0xf1,
	0x27, 0x80, 0xe1, 0xba, 0xbc, 0x93, 0x68, 0x0d, 0xe0, 0xd6, 0x84, 0xce, 0x97, 0xee, 0xfa, 0xee,
	0x6d, 0x2f, 0x7c, 0x7a, 0x0f, 0x6a, 0x1b, 0xad, 0x01, 0xdc, 0x1e, 0x1b, 0x52, 0x7b, 0x3b, 0x6f,
	0x48, 0x75, 0x2c, 0xff, 0x23, 0x0c, 0xab, 0x2c, 0x9a, 0x75, 0x71, 0x4b, 0xa1, 0xf3, 0x6e, 0xd0,
	0xea, 0xac, 0x60, 0x50, 0xfa, 0x81, 0xc2, 0x1a, 0xb3, 0xe5, 0x70, 0x38, 0xeb, 0xc4, 0x9c, 0x48,
	0xf9, 0xf6, 0x34, 0x44, 0x5a, 0x8f, 0x57, 0x38, 0xeb, 0xc4, 0xac, 0xc8, 0x35, 0x8c, 0x6a, 0x37,
	0x0e, 0x35, 0xac, 0xdc, 0x7b, 0x5e, 0xc2, 0x67, 0xf7, 0xc1, 0x46, 0xed, 0xf2, 0xe0, 0x9b, 0x9f,
	0xdd, 0xde, 0xf6, 0xf5, 0x3b, 0xfa, 0xee, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x4d, 0xa6,
	0x45, 0x74, 0x05, 0x00, 0x00,
}

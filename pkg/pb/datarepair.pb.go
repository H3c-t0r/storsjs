// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datarepair.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// InjuredSegment is the queue item used for the data repair queue
type InjuredSegment struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	LostPieces           []int32  `protobuf:"varint,2,rep,packed,name=lost_pieces,json=lostPieces,proto3" json:"lost_pieces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InjuredSegment) Reset()         { *m = InjuredSegment{} }
func (m *InjuredSegment) String() string { return proto.CompactTextString(m) }
func (*InjuredSegment) ProtoMessage()    {}
func (*InjuredSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_datarepair_ff8a1b3aa647676c, []int{0}
}
func (m *InjuredSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InjuredSegment.Unmarshal(m, b)
}
func (m *InjuredSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InjuredSegment.Marshal(b, m, deterministic)
}
func (dst *InjuredSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InjuredSegment.Merge(dst, src)
}
func (m *InjuredSegment) XXX_Size() int {
	return xxx_messageInfo_InjuredSegment.Size(m)
}
func (m *InjuredSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_InjuredSegment.DiscardUnknown(m)
}

var xxx_messageInfo_InjuredSegment proto.InternalMessageInfo

func (m *InjuredSegment) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *InjuredSegment) GetLostPieces() []int32 {
	if m != nil {
		return m.LostPieces
	}
	return nil
}

// IdentifyRequest is a request message for the identifyInjuredSegments private method
type IdentifyRequest struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	First                string   `protobuf:"bytes,2,opt,name=first,proto3" json:"first,omitempty"`
	Recurse              bool     `protobuf:"varint,3,opt,name=recurse,proto3" json:"recurse,omitempty"`
	Reverse              bool     `protobuf:"varint,4,opt,name=reverse,proto3" json:"reverse,omitempty"`
	Limit                int32    `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentifyRequest) Reset()         { *m = IdentifyRequest{} }
func (m *IdentifyRequest) String() string { return proto.CompactTextString(m) }
func (*IdentifyRequest) ProtoMessage()    {}
func (*IdentifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_datarepair_ff8a1b3aa647676c, []int{1}
}
func (m *IdentifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdentifyRequest.Unmarshal(m, b)
}
func (m *IdentifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdentifyRequest.Marshal(b, m, deterministic)
}
func (dst *IdentifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentifyRequest.Merge(dst, src)
}
func (m *IdentifyRequest) XXX_Size() int {
	return xxx_messageInfo_IdentifyRequest.Size(m)
}
func (m *IdentifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdentifyRequest proto.InternalMessageInfo

func (m *IdentifyRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *IdentifyRequest) GetFirst() string {
	if m != nil {
		return m.First
	}
	return ""
}

func (m *IdentifyRequest) GetRecurse() bool {
	if m != nil {
		return m.Recurse
	}
	return false
}

func (m *IdentifyRequest) GetReverse() bool {
	if m != nil {
		return m.Reverse
	}
	return false
}

func (m *IdentifyRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*InjuredSegment)(nil), "repair.InjuredSegment")
	proto.RegisterType((*IdentifyRequest)(nil), "repair.IdentifyRequest")
}

func init() { proto.RegisterFile("datarepair.proto", fileDescriptor_datarepair_ff8a1b3aa647676c) }

var fileDescriptor_datarepair_ff8a1b3aa647676c = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xcf, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x06, 0x60, 0xd2, 0x6d, 0xaa, 0x8e, 0xa0, 0x12, 0x44, 0x72, 0x33, 0xec, 0x29, 0x27, 0x2f,
	0xbe, 0x81, 0xe0, 0x61, 0x6f, 0x12, 0x6f, 0x5e, 0x24, 0xbb, 0x9d, 0x6a, 0x64, 0xdb, 0xc6, 0xc9,
	0x54, 0xf4, 0x11, 0x7c, 0xeb, 0xa5, 0x49, 0x7b, 0x9b, 0xef, 0xff, 0xe1, 0x87, 0x81, 0x9b, 0xd6,
	0xb3, 0x27, 0x8c, 0x3e, 0xd0, 0x43, 0xa4, 0x91, 0x47, 0xd5, 0x14, 0x6d, 0x9f, 0xe1, 0x6a, 0x37,
	0x7c, 0x4d, 0x84, 0xed, 0x2b, 0x7e, 0xf4, 0x38, 0xb0, 0x52, 0x50, 0x47, 0xcf, 0x9f, 0x5a, 0x18,
	0x61, 0x2f, 0x5c, 0xbe, 0xd5, 0x3d, 0x5c, 0x1e, 0xc7, 0xc4, 0xef, 0x31, 0xe0, 0x01, 0x93, 0xae,
	0xcc, 0xc6, 0x4a, 0x07, 0x73, 0xf4, 0x92, 0x93, 0xed, 0xbf, 0x80, 0xeb, 0x5d, 0x8b, 0x03, 0x87,
	0xee, 0xcf, 0xe1, 0xf7, 0x84, 0x89, 0xd5, 0x1d, 0x34, 0x91, 0xb0, 0x0b, 0xbf, 0xcb, 0xd4, 0x22,
	0x75, 0x0b, 0xb2, 0x0b, 0x94, 0x58, 0x57, 0x39, 0x2e, 0x50, 0x1a, 0xce, 0x08, 0x0f, 0x13, 0x25,
	0xd4, 0x1b, 0x23, 0xec, 0xb9, 0x5b, 0x59, 0x9a, 0x1f, 0x9c, 0x9b, 0x7a, 0x6d, 0x32, 0xe7, 0xa5,
	0x63, 0xe8, 0x03, 0x6b, 0x69, 0x84, 0x95, 0xae, 0xe0, 0xa9, 0x7e, 0xab, 0xe2, 0x7e, 0xdf, 0xe4,
	0x3f, 0x1f, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x01, 0xe7, 0x72, 0x19, 0xfb, 0x00, 0x00, 0x00,
}

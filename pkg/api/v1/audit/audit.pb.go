// Code generated by protoc-gen-go. DO NOT EDIT.
// source: audit.proto

package audit

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Audit struct {
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ModifiedAt           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
	Vers                 int64                `protobuf:"varint,3,opt,name=vers,proto3" json:"vers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Audit) Reset()         { *m = Audit{} }
func (m *Audit) String() string { return proto.CompactTextString(m) }
func (*Audit) ProtoMessage()    {}
func (*Audit) Descriptor() ([]byte, []int) {
	return fileDescriptor_5594839dd8e38a1b, []int{0}
}

func (m *Audit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Audit.Unmarshal(m, b)
}
func (m *Audit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Audit.Marshal(b, m, deterministic)
}
func (m *Audit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Audit.Merge(m, src)
}
func (m *Audit) XXX_Size() int {
	return xxx_messageInfo_Audit.Size(m)
}
func (m *Audit) XXX_DiscardUnknown() {
	xxx_messageInfo_Audit.DiscardUnknown(m)
}

var xxx_messageInfo_Audit proto.InternalMessageInfo

func (m *Audit) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Audit) GetModifiedAt() *timestamp.Timestamp {
	if m != nil {
		return m.ModifiedAt
	}
	return nil
}

func (m *Audit) GetVers() int64 {
	if m != nil {
		return m.Vers
	}
	return 0
}

func init() {
	proto.RegisterType((*Audit)(nil), "v1.Audit")
}

func init() { proto.RegisterFile("audit.proto", fileDescriptor_5594839dd8e38a1b) }

var fileDescriptor_5594839dd8e38a1b = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xcf, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xc0, 0x71, 0xb9, 0x05, 0x24, 0x9c, 0xcd, 0x53, 0xd5, 0x85, 0x8a, 0xa9, 0x0b, 0x3e, 0x15,
	0xa6, 0xaa, 0x53, 0x78, 0x84, 0x8a, 0x89, 0x05, 0x9d, 0x9d, 0xab, 0x63, 0x81, 0x3f, 0x94, 0x9c,
	0x23, 0xf1, 0x1c, 0xbc, 0x30, 0xc2, 0x21, 0x73, 0xb7, 0xd3, 0xe9, 0xff, 0xd3, 0xe9, 0x64, 0x83,
	0xa5, 0xf3, 0xac, 0xf3, 0x90, 0x38, 0xa9, 0xd5, 0x74, 0xd8, 0x3e, 0xb8, 0x94, 0xdc, 0x17, 0x41,
	0xdd, 0x98, 0x72, 0x01, 0xf6, 0x81, 0x46, 0xc6, 0x90, 0xe7, 0xe8, 0xf1, 0x47, 0xc8, 0xdb, 0xf6,
	0x0f, 0xa9, 0xa3, 0x94, 0x76, 0x20, 0x64, 0xea, 0x3e, 0x90, 0x37, 0x62, 0x27, 0xf6, 0xcd, 0xf3,
	0x56, 0xcf, 0x5e, 0x2f, 0x5e, 0xbf, 0x2d, 0xfe, 0x7c, 0xff, 0x5f, 0xb7, 0xac, 0x4e, 0xb2, 0x09,
	0xa9, 0xf3, 0x17, 0x3f, 0xdb, 0xd5, 0x55, 0x2b, 0x97, 0xbc, 0x65, 0xa5, 0xe4, 0xcd, 0x44, 0xc3,
	0xb8, 0x59, 0xef, 0xc4, 0x7e, 0x7d, 0xae, 0xf3, 0xeb, 0xe9, 0xfd, 0xe8, 0x3c, 0xf7, 0xc5, 0x68,
	0x9b, 0x02, 0x98, 0x12, 0xdd, 0xf7, 0xd8, 0x13, 0x65, 0xb0, 0x29, 0x32, 0x5a, 0x7e, 0x0a, 0x18,
	0xd1, 0x51, 0xa0, 0xc8, 0x90, 0x3f, 0x1d, 0x60, 0xf6, 0x30, 0x1d, 0xa0, 0x7e, 0x6f, 0xee, 0xea,
	0xc1, 0x97, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0x85, 0x6a, 0x6d, 0x0d, 0x01, 0x00, 0x00,
}

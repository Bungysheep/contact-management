// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contactsystem.proto

package contactsystem

import (
	context "context"
	fmt "fmt"
	audit "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	message "github.com/bungysheep/contact-management/pkg/api/v1/message"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ContactSystem struct {
	ContactSystemCode    string       `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Details              string       `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	Status               string       `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Audit                *audit.Audit `protobuf:"bytes,5,opt,name=audit,proto3" json:"audit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContactSystem) Reset()         { *m = ContactSystem{} }
func (m *ContactSystem) String() string { return proto.CompactTextString(m) }
func (*ContactSystem) ProtoMessage()    {}
func (*ContactSystem) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{0}
}

func (m *ContactSystem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactSystem.Unmarshal(m, b)
}
func (m *ContactSystem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactSystem.Marshal(b, m, deterministic)
}
func (m *ContactSystem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactSystem.Merge(m, src)
}
func (m *ContactSystem) XXX_Size() int {
	return xxx_messageInfo_ContactSystem.Size(m)
}
func (m *ContactSystem) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactSystem.DiscardUnknown(m)
}

var xxx_messageInfo_ContactSystem proto.InternalMessageInfo

func (m *ContactSystem) GetContactSystemCode() string {
	if m != nil {
		return m.ContactSystemCode
	}
	return ""
}

func (m *ContactSystem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ContactSystem) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *ContactSystem) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ContactSystem) GetAudit() *audit.Audit {
	if m != nil {
		return m.Audit
	}
	return nil
}

type DoReadContactSystemRequest struct {
	ContactSystemCode    string   `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoReadContactSystemRequest) Reset()         { *m = DoReadContactSystemRequest{} }
func (m *DoReadContactSystemRequest) String() string { return proto.CompactTextString(m) }
func (*DoReadContactSystemRequest) ProtoMessage()    {}
func (*DoReadContactSystemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{1}
}

func (m *DoReadContactSystemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReadContactSystemRequest.Unmarshal(m, b)
}
func (m *DoReadContactSystemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReadContactSystemRequest.Marshal(b, m, deterministic)
}
func (m *DoReadContactSystemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReadContactSystemRequest.Merge(m, src)
}
func (m *DoReadContactSystemRequest) XXX_Size() int {
	return xxx_messageInfo_DoReadContactSystemRequest.Size(m)
}
func (m *DoReadContactSystemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReadContactSystemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoReadContactSystemRequest proto.InternalMessageInfo

func (m *DoReadContactSystemRequest) GetContactSystemCode() string {
	if m != nil {
		return m.ContactSystemCode
	}
	return ""
}

type DoReadContactSystemResponse struct {
	ContactSystem        *ContactSystem   `protobuf:"bytes,1,opt,name=contact_system,json=contactSystem,proto3" json:"contact_system,omitempty"`
	Message              *message.Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DoReadContactSystemResponse) Reset()         { *m = DoReadContactSystemResponse{} }
func (m *DoReadContactSystemResponse) String() string { return proto.CompactTextString(m) }
func (*DoReadContactSystemResponse) ProtoMessage()    {}
func (*DoReadContactSystemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{2}
}

func (m *DoReadContactSystemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReadContactSystemResponse.Unmarshal(m, b)
}
func (m *DoReadContactSystemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReadContactSystemResponse.Marshal(b, m, deterministic)
}
func (m *DoReadContactSystemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReadContactSystemResponse.Merge(m, src)
}
func (m *DoReadContactSystemResponse) XXX_Size() int {
	return xxx_messageInfo_DoReadContactSystemResponse.Size(m)
}
func (m *DoReadContactSystemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReadContactSystemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoReadContactSystemResponse proto.InternalMessageInfo

func (m *DoReadContactSystemResponse) GetContactSystem() *ContactSystem {
	if m != nil {
		return m.ContactSystem
	}
	return nil
}

func (m *DoReadContactSystemResponse) GetMessage() *message.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type DoReadAllContactSystemRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoReadAllContactSystemRequest) Reset()         { *m = DoReadAllContactSystemRequest{} }
func (m *DoReadAllContactSystemRequest) String() string { return proto.CompactTextString(m) }
func (*DoReadAllContactSystemRequest) ProtoMessage()    {}
func (*DoReadAllContactSystemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{3}
}

func (m *DoReadAllContactSystemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReadAllContactSystemRequest.Unmarshal(m, b)
}
func (m *DoReadAllContactSystemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReadAllContactSystemRequest.Marshal(b, m, deterministic)
}
func (m *DoReadAllContactSystemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReadAllContactSystemRequest.Merge(m, src)
}
func (m *DoReadAllContactSystemRequest) XXX_Size() int {
	return xxx_messageInfo_DoReadAllContactSystemRequest.Size(m)
}
func (m *DoReadAllContactSystemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReadAllContactSystemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoReadAllContactSystemRequest proto.InternalMessageInfo

type DoReadAllContactSystemResponse struct {
	ContactSystems       []*ContactSystem `protobuf:"bytes,1,rep,name=contact_systems,json=contactSystems,proto3" json:"contact_systems,omitempty"`
	Message              *message.Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DoReadAllContactSystemResponse) Reset()         { *m = DoReadAllContactSystemResponse{} }
func (m *DoReadAllContactSystemResponse) String() string { return proto.CompactTextString(m) }
func (*DoReadAllContactSystemResponse) ProtoMessage()    {}
func (*DoReadAllContactSystemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{4}
}

func (m *DoReadAllContactSystemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReadAllContactSystemResponse.Unmarshal(m, b)
}
func (m *DoReadAllContactSystemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReadAllContactSystemResponse.Marshal(b, m, deterministic)
}
func (m *DoReadAllContactSystemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReadAllContactSystemResponse.Merge(m, src)
}
func (m *DoReadAllContactSystemResponse) XXX_Size() int {
	return xxx_messageInfo_DoReadAllContactSystemResponse.Size(m)
}
func (m *DoReadAllContactSystemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReadAllContactSystemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoReadAllContactSystemResponse proto.InternalMessageInfo

func (m *DoReadAllContactSystemResponse) GetContactSystems() []*ContactSystem {
	if m != nil {
		return m.ContactSystems
	}
	return nil
}

func (m *DoReadAllContactSystemResponse) GetMessage() *message.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type DoSaveContactSystemRequest struct {
	ContactSystem        *ContactSystem `protobuf:"bytes,1,opt,name=contact_system,json=contactSystem,proto3" json:"contact_system,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DoSaveContactSystemRequest) Reset()         { *m = DoSaveContactSystemRequest{} }
func (m *DoSaveContactSystemRequest) String() string { return proto.CompactTextString(m) }
func (*DoSaveContactSystemRequest) ProtoMessage()    {}
func (*DoSaveContactSystemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{5}
}

func (m *DoSaveContactSystemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoSaveContactSystemRequest.Unmarshal(m, b)
}
func (m *DoSaveContactSystemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoSaveContactSystemRequest.Marshal(b, m, deterministic)
}
func (m *DoSaveContactSystemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoSaveContactSystemRequest.Merge(m, src)
}
func (m *DoSaveContactSystemRequest) XXX_Size() int {
	return xxx_messageInfo_DoSaveContactSystemRequest.Size(m)
}
func (m *DoSaveContactSystemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoSaveContactSystemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoSaveContactSystemRequest proto.InternalMessageInfo

func (m *DoSaveContactSystemRequest) GetContactSystem() *ContactSystem {
	if m != nil {
		return m.ContactSystem
	}
	return nil
}

type DoSaveContactSystemResponse struct {
	Result               bool             `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Message              *message.Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DoSaveContactSystemResponse) Reset()         { *m = DoSaveContactSystemResponse{} }
func (m *DoSaveContactSystemResponse) String() string { return proto.CompactTextString(m) }
func (*DoSaveContactSystemResponse) ProtoMessage()    {}
func (*DoSaveContactSystemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{6}
}

func (m *DoSaveContactSystemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoSaveContactSystemResponse.Unmarshal(m, b)
}
func (m *DoSaveContactSystemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoSaveContactSystemResponse.Marshal(b, m, deterministic)
}
func (m *DoSaveContactSystemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoSaveContactSystemResponse.Merge(m, src)
}
func (m *DoSaveContactSystemResponse) XXX_Size() int {
	return xxx_messageInfo_DoSaveContactSystemResponse.Size(m)
}
func (m *DoSaveContactSystemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoSaveContactSystemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoSaveContactSystemResponse proto.InternalMessageInfo

func (m *DoSaveContactSystemResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *DoSaveContactSystemResponse) GetMessage() *message.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type DoDeleteContactSystemRequest struct {
	ContactSystemCode    string   `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoDeleteContactSystemRequest) Reset()         { *m = DoDeleteContactSystemRequest{} }
func (m *DoDeleteContactSystemRequest) String() string { return proto.CompactTextString(m) }
func (*DoDeleteContactSystemRequest) ProtoMessage()    {}
func (*DoDeleteContactSystemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{7}
}

func (m *DoDeleteContactSystemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoDeleteContactSystemRequest.Unmarshal(m, b)
}
func (m *DoDeleteContactSystemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoDeleteContactSystemRequest.Marshal(b, m, deterministic)
}
func (m *DoDeleteContactSystemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoDeleteContactSystemRequest.Merge(m, src)
}
func (m *DoDeleteContactSystemRequest) XXX_Size() int {
	return xxx_messageInfo_DoDeleteContactSystemRequest.Size(m)
}
func (m *DoDeleteContactSystemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoDeleteContactSystemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoDeleteContactSystemRequest proto.InternalMessageInfo

func (m *DoDeleteContactSystemRequest) GetContactSystemCode() string {
	if m != nil {
		return m.ContactSystemCode
	}
	return ""
}

type DoDeleteContactSystemResponse struct {
	Result               bool             `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Message              *message.Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DoDeleteContactSystemResponse) Reset()         { *m = DoDeleteContactSystemResponse{} }
func (m *DoDeleteContactSystemResponse) String() string { return proto.CompactTextString(m) }
func (*DoDeleteContactSystemResponse) ProtoMessage()    {}
func (*DoDeleteContactSystemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{8}
}

func (m *DoDeleteContactSystemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoDeleteContactSystemResponse.Unmarshal(m, b)
}
func (m *DoDeleteContactSystemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoDeleteContactSystemResponse.Marshal(b, m, deterministic)
}
func (m *DoDeleteContactSystemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoDeleteContactSystemResponse.Merge(m, src)
}
func (m *DoDeleteContactSystemResponse) XXX_Size() int {
	return xxx_messageInfo_DoDeleteContactSystemResponse.Size(m)
}
func (m *DoDeleteContactSystemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoDeleteContactSystemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoDeleteContactSystemResponse proto.InternalMessageInfo

func (m *DoDeleteContactSystemResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *DoDeleteContactSystemResponse) GetMessage() *message.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterType((*ContactSystem)(nil), "v1.ContactSystem")
	proto.RegisterType((*DoReadContactSystemRequest)(nil), "v1.DoReadContactSystemRequest")
	proto.RegisterType((*DoReadContactSystemResponse)(nil), "v1.DoReadContactSystemResponse")
	proto.RegisterType((*DoReadAllContactSystemRequest)(nil), "v1.DoReadAllContactSystemRequest")
	proto.RegisterType((*DoReadAllContactSystemResponse)(nil), "v1.DoReadAllContactSystemResponse")
	proto.RegisterType((*DoSaveContactSystemRequest)(nil), "v1.DoSaveContactSystemRequest")
	proto.RegisterType((*DoSaveContactSystemResponse)(nil), "v1.DoSaveContactSystemResponse")
	proto.RegisterType((*DoDeleteContactSystemRequest)(nil), "v1.DoDeleteContactSystemRequest")
	proto.RegisterType((*DoDeleteContactSystemResponse)(nil), "v1.DoDeleteContactSystemResponse")
}

func init() { proto.RegisterFile("contactsystem.proto", fileDescriptor_dc89358642882005) }

var fileDescriptor_dc89358642882005 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0x56, 0xd6, 0xad, 0x8e, 0x5a, 0x98, 0x87, 0xaa, 0x10, 0xc6, 0xda, 0x59, 0x80, 0xa6,
	0x09, 0x1a, 0xb5, 0xdc, 0xa0, 0xdd, 0x95, 0xf6, 0x12, 0xb8, 0x48, 0x25, 0x2e, 0x10, 0x62, 0x72,
	0x93, 0xa3, 0x2c, 0x90, 0xc6, 0xa1, 0x76, 0x8a, 0x26, 0x34, 0x2e, 0xe0, 0x11, 0x78, 0x15, 0xde,
	0x84, 0x17, 0xe0, 0x82, 0x07, 0x41, 0xb6, 0x93, 0x31, 0xa3, 0x64, 0x2a, 0x68, 0x97, 0xe7, 0xef,
	0x3b, 0xdf, 0xf9, 0xfc, 0xc9, 0x68, 0x37, 0x60, 0xa9, 0xa0, 0x81, 0xe0, 0x67, 0x5c, 0xc0, 0x62,
	0x90, 0x2d, 0x99, 0x60, 0x78, 0x63, 0x35, 0x74, 0xf7, 0x22, 0xc6, 0xa2, 0x04, 0x3c, 0x9a, 0xc5,
	0x1e, 0x4d, 0x53, 0x26, 0xa8, 0x88, 0x59, 0xca, 0x75, 0x87, 0x6b, 0xd3, 0x3c, 0x8c, 0x45, 0x11,
	0xb4, 0x17, 0xc0, 0x39, 0x8d, 0x40, 0x87, 0xe4, 0xbb, 0x85, 0xda, 0x13, 0x8d, 0x3a, 0x53, 0xa8,
	0x78, 0x70, 0xb1, 0xe6, 0x44, 0xef, 0x39, 0x09, 0x58, 0x08, 0x8e, 0xd5, 0xb7, 0x0e, 0x5b, 0xfe,
	0x4e, 0x70, 0xb9, 0x77, 0xc2, 0x42, 0xc0, 0x7d, 0x64, 0x87, 0xc0, 0x83, 0x65, 0x9c, 0xc9, 0x9d,
	0xce, 0x86, 0xea, 0xbb, 0x9c, 0xc2, 0x0e, 0xda, 0x0a, 0x41, 0xd0, 0x38, 0xe1, 0x4e, 0x43, 0x55,
	0xcb, 0x10, 0x77, 0x51, 0x93, 0x0b, 0x2a, 0x72, 0xee, 0xdc, 0x50, 0x85, 0x22, 0xc2, 0x3d, 0xb4,
	0xa9, 0x38, 0x3b, 0x9b, 0x7d, 0xeb, 0xd0, 0x1e, 0xb5, 0x06, 0xab, 0xe1, 0x60, 0x2c, 0x13, 0xbe,
	0xce, 0x93, 0xe7, 0xc8, 0x9d, 0x32, 0x1f, 0x68, 0x68, 0x70, 0xf7, 0xe1, 0x43, 0x0e, 0x5c, 0xfc,
	0xeb, 0x09, 0xe4, 0x33, 0xba, 0x5b, 0x89, 0xc6, 0x33, 0x96, 0x72, 0xc0, 0x4f, 0x51, 0xc7, 0x84,
	0x53, 0x48, 0xf6, 0x68, 0x47, 0xd2, 0x32, 0x47, 0xda, 0x06, 0x38, 0x7e, 0x80, 0xb6, 0x0a, 0xb9,
	0x95, 0x2e, 0xf6, 0xc8, 0x96, 0x23, 0x2f, 0x74, 0xca, 0x2f, 0x6b, 0xa4, 0x87, 0xee, 0xe9, 0xfd,
	0xe3, 0x24, 0xa9, 0x3a, 0x88, 0x7c, 0xb5, 0xd0, 0x7e, 0x5d, 0x47, 0x41, 0xf2, 0x18, 0xdd, 0x34,
	0x49, 0x72, 0xc7, 0xea, 0x37, 0xaa, 0x59, 0x76, 0x0c, 0x96, 0x7c, 0x5d, 0x9a, 0xaf, 0xa4, 0xe8,
	0x33, 0xba, 0x82, 0x4a, 0xd1, 0xff, 0x5b, 0x25, 0xf2, 0x46, 0xca, 0x5f, 0x81, 0x5b, 0x5c, 0xd6,
	0x45, 0xcd, 0x25, 0xf0, 0x3c, 0x11, 0x0a, 0x70, 0xdb, 0x2f, 0xa2, 0x75, 0x59, 0xbf, 0x44, 0x7b,
	0x53, 0x36, 0x85, 0x04, 0x04, 0x5c, 0x8b, 0x59, 0xde, 0xca, 0xc7, 0xaa, 0xc4, 0xbb, 0x16, 0xbe,
	0xa3, 0x9f, 0x0d, 0x74, 0xdb, 0x00, 0x9e, 0xc1, 0x72, 0x15, 0x07, 0x80, 0x3f, 0xa2, 0xa6, 0xf6,
	0x00, 0xde, 0x97, 0x83, 0xf5, 0xfe, 0x77, 0x7b, 0xb5, 0x75, 0x4d, 0x91, 0x3c, 0xfa, 0xf2, 0xe3,
	0xd7, 0xb7, 0x8d, 0x87, 0xf8, 0xbe, 0x67, 0xfc, 0x28, 0xde, 0xa7, 0x0a, 0x25, 0xce, 0xf1, 0x3b,
	0xd4, 0xba, 0x30, 0x1f, 0x3e, 0xf8, 0x83, 0x5d, 0xe3, 0x56, 0x97, 0x5c, 0xd5, 0x52, 0x30, 0xe8,
	0x2a, 0x06, 0xb7, 0x70, 0xc7, 0x64, 0x80, 0x43, 0x79, 0xa4, 0xf4, 0x42, 0x79, 0x64, 0x9d, 0xdf,
	0xca, 0x23, 0x6b, 0x7d, 0x43, 0xee, 0xa8, 0x15, 0xbb, 0xe4, 0xaf, 0x15, 0xc7, 0xd6, 0x11, 0x3e,
	0x47, 0xdb, 0xe5, 0x1b, 0xe2, 0xbe, 0xc6, 0xa9, 0x77, 0x88, 0x7b, 0x70, 0x45, 0x87, 0x29, 0xe8,
	0xd1, 0x5a, 0x82, 0x3e, 0x9b, 0xbc, 0x1e, 0x47, 0xb1, 0x38, 0xcd, 0xe7, 0x83, 0x80, 0x2d, 0xbc,
	0x79, 0x9e, 0x46, 0x67, 0xfc, 0x14, 0x20, 0x2b, 0x87, 0x1f, 0x2f, 0x68, 0x4a, 0x23, 0x58, 0x40,
	0x2a, 0xbc, 0xec, 0x7d, 0xa4, 0xbe, 0xf5, 0xd5, 0xd0, 0x84, 0x9e, 0x37, 0xd5, 0x07, 0xfe, 0xe4,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x80, 0xf6, 0xe3, 0x15, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContactSystemServiceClient is the client API for ContactSystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContactSystemServiceClient interface {
	DoRead(ctx context.Context, in *DoReadContactSystemRequest, opts ...grpc.CallOption) (*DoReadContactSystemResponse, error)
	DoReadAll(ctx context.Context, in *DoReadAllContactSystemRequest, opts ...grpc.CallOption) (*DoReadAllContactSystemResponse, error)
	DoSave(ctx context.Context, in *DoSaveContactSystemRequest, opts ...grpc.CallOption) (*DoSaveContactSystemResponse, error)
	DoDelete(ctx context.Context, in *DoDeleteContactSystemRequest, opts ...grpc.CallOption) (*DoDeleteContactSystemResponse, error)
}

type contactSystemServiceClient struct {
	cc *grpc.ClientConn
}

func NewContactSystemServiceClient(cc *grpc.ClientConn) ContactSystemServiceClient {
	return &contactSystemServiceClient{cc}
}

func (c *contactSystemServiceClient) DoRead(ctx context.Context, in *DoReadContactSystemRequest, opts ...grpc.CallOption) (*DoReadContactSystemResponse, error) {
	out := new(DoReadContactSystemResponse)
	err := c.cc.Invoke(ctx, "/v1.ContactSystemService/DoRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactSystemServiceClient) DoReadAll(ctx context.Context, in *DoReadAllContactSystemRequest, opts ...grpc.CallOption) (*DoReadAllContactSystemResponse, error) {
	out := new(DoReadAllContactSystemResponse)
	err := c.cc.Invoke(ctx, "/v1.ContactSystemService/DoReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactSystemServiceClient) DoSave(ctx context.Context, in *DoSaveContactSystemRequest, opts ...grpc.CallOption) (*DoSaveContactSystemResponse, error) {
	out := new(DoSaveContactSystemResponse)
	err := c.cc.Invoke(ctx, "/v1.ContactSystemService/DoSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactSystemServiceClient) DoDelete(ctx context.Context, in *DoDeleteContactSystemRequest, opts ...grpc.CallOption) (*DoDeleteContactSystemResponse, error) {
	out := new(DoDeleteContactSystemResponse)
	err := c.cc.Invoke(ctx, "/v1.ContactSystemService/DoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactSystemServiceServer is the server API for ContactSystemService service.
type ContactSystemServiceServer interface {
	DoRead(context.Context, *DoReadContactSystemRequest) (*DoReadContactSystemResponse, error)
	DoReadAll(context.Context, *DoReadAllContactSystemRequest) (*DoReadAllContactSystemResponse, error)
	DoSave(context.Context, *DoSaveContactSystemRequest) (*DoSaveContactSystemResponse, error)
	DoDelete(context.Context, *DoDeleteContactSystemRequest) (*DoDeleteContactSystemResponse, error)
}

// UnimplementedContactSystemServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContactSystemServiceServer struct {
}

func (*UnimplementedContactSystemServiceServer) DoRead(ctx context.Context, req *DoReadContactSystemRequest) (*DoReadContactSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoRead not implemented")
}
func (*UnimplementedContactSystemServiceServer) DoReadAll(ctx context.Context, req *DoReadAllContactSystemRequest) (*DoReadAllContactSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoReadAll not implemented")
}
func (*UnimplementedContactSystemServiceServer) DoSave(ctx context.Context, req *DoSaveContactSystemRequest) (*DoSaveContactSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSave not implemented")
}
func (*UnimplementedContactSystemServiceServer) DoDelete(ctx context.Context, req *DoDeleteContactSystemRequest) (*DoDeleteContactSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoDelete not implemented")
}

func RegisterContactSystemServiceServer(s *grpc.Server, srv ContactSystemServiceServer) {
	s.RegisterService(&_ContactSystemService_serviceDesc, srv)
}

func _ContactSystemService_DoRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoReadContactSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactSystemServiceServer).DoRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ContactSystemService/DoRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactSystemServiceServer).DoRead(ctx, req.(*DoReadContactSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactSystemService_DoReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoReadAllContactSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactSystemServiceServer).DoReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ContactSystemService/DoReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactSystemServiceServer).DoReadAll(ctx, req.(*DoReadAllContactSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactSystemService_DoSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoSaveContactSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactSystemServiceServer).DoSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ContactSystemService/DoSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactSystemServiceServer).DoSave(ctx, req.(*DoSaveContactSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactSystemService_DoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoDeleteContactSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactSystemServiceServer).DoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ContactSystemService/DoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactSystemServiceServer).DoDelete(ctx, req.(*DoDeleteContactSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContactSystemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ContactSystemService",
	HandlerType: (*ContactSystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoRead",
			Handler:    _ContactSystemService_DoRead_Handler,
		},
		{
			MethodName: "DoReadAll",
			Handler:    _ContactSystemService_DoReadAll_Handler,
		},
		{
			MethodName: "DoSave",
			Handler:    _ContactSystemService_DoSave_Handler,
		},
		{
			MethodName: "DoDelete",
			Handler:    _ContactSystemService_DoDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contactsystem.proto",
}

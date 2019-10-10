// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contactsystem.proto

package contactsystem

import (
	context "context"
	fmt "fmt"
	audit "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	proto "github.com/golang/protobuf/proto"
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

type DoReadRequest struct {
	ContactSystemCode    string   `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoReadRequest) Reset()         { *m = DoReadRequest{} }
func (m *DoReadRequest) String() string { return proto.CompactTextString(m) }
func (*DoReadRequest) ProtoMessage()    {}
func (*DoReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{1}
}

func (m *DoReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReadRequest.Unmarshal(m, b)
}
func (m *DoReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReadRequest.Marshal(b, m, deterministic)
}
func (m *DoReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReadRequest.Merge(m, src)
}
func (m *DoReadRequest) XXX_Size() int {
	return xxx_messageInfo_DoReadRequest.Size(m)
}
func (m *DoReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoReadRequest proto.InternalMessageInfo

func (m *DoReadRequest) GetContactSystemCode() string {
	if m != nil {
		return m.ContactSystemCode
	}
	return ""
}

type DoReadResponse struct {
	ContactSystem        *ContactSystem `protobuf:"bytes,1,opt,name=contact_system,json=contactSystem,proto3" json:"contact_system,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DoReadResponse) Reset()         { *m = DoReadResponse{} }
func (m *DoReadResponse) String() string { return proto.CompactTextString(m) }
func (*DoReadResponse) ProtoMessage()    {}
func (*DoReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{2}
}

func (m *DoReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReadResponse.Unmarshal(m, b)
}
func (m *DoReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReadResponse.Marshal(b, m, deterministic)
}
func (m *DoReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReadResponse.Merge(m, src)
}
func (m *DoReadResponse) XXX_Size() int {
	return xxx_messageInfo_DoReadResponse.Size(m)
}
func (m *DoReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoReadResponse proto.InternalMessageInfo

func (m *DoReadResponse) GetContactSystem() *ContactSystem {
	if m != nil {
		return m.ContactSystem
	}
	return nil
}

type DoReadAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoReadAllRequest) Reset()         { *m = DoReadAllRequest{} }
func (m *DoReadAllRequest) String() string { return proto.CompactTextString(m) }
func (*DoReadAllRequest) ProtoMessage()    {}
func (*DoReadAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{3}
}

func (m *DoReadAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReadAllRequest.Unmarshal(m, b)
}
func (m *DoReadAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReadAllRequest.Marshal(b, m, deterministic)
}
func (m *DoReadAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReadAllRequest.Merge(m, src)
}
func (m *DoReadAllRequest) XXX_Size() int {
	return xxx_messageInfo_DoReadAllRequest.Size(m)
}
func (m *DoReadAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReadAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoReadAllRequest proto.InternalMessageInfo

type DoReadAllResponse struct {
	ContactSystems       []*ContactSystem `protobuf:"bytes,1,rep,name=contact_systems,json=contactSystems,proto3" json:"contact_systems,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DoReadAllResponse) Reset()         { *m = DoReadAllResponse{} }
func (m *DoReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*DoReadAllResponse) ProtoMessage()    {}
func (*DoReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{4}
}

func (m *DoReadAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReadAllResponse.Unmarshal(m, b)
}
func (m *DoReadAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReadAllResponse.Marshal(b, m, deterministic)
}
func (m *DoReadAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReadAllResponse.Merge(m, src)
}
func (m *DoReadAllResponse) XXX_Size() int {
	return xxx_messageInfo_DoReadAllResponse.Size(m)
}
func (m *DoReadAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReadAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoReadAllResponse proto.InternalMessageInfo

func (m *DoReadAllResponse) GetContactSystems() []*ContactSystem {
	if m != nil {
		return m.ContactSystems
	}
	return nil
}

type DoSaveRequest struct {
	ContactSystem        *ContactSystem `protobuf:"bytes,1,opt,name=contact_system,json=contactSystem,proto3" json:"contact_system,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DoSaveRequest) Reset()         { *m = DoSaveRequest{} }
func (m *DoSaveRequest) String() string { return proto.CompactTextString(m) }
func (*DoSaveRequest) ProtoMessage()    {}
func (*DoSaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{5}
}

func (m *DoSaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoSaveRequest.Unmarshal(m, b)
}
func (m *DoSaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoSaveRequest.Marshal(b, m, deterministic)
}
func (m *DoSaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoSaveRequest.Merge(m, src)
}
func (m *DoSaveRequest) XXX_Size() int {
	return xxx_messageInfo_DoSaveRequest.Size(m)
}
func (m *DoSaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoSaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoSaveRequest proto.InternalMessageInfo

func (m *DoSaveRequest) GetContactSystem() *ContactSystem {
	if m != nil {
		return m.ContactSystem
	}
	return nil
}

type DoSaveResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoSaveResponse) Reset()         { *m = DoSaveResponse{} }
func (m *DoSaveResponse) String() string { return proto.CompactTextString(m) }
func (*DoSaveResponse) ProtoMessage()    {}
func (*DoSaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{6}
}

func (m *DoSaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoSaveResponse.Unmarshal(m, b)
}
func (m *DoSaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoSaveResponse.Marshal(b, m, deterministic)
}
func (m *DoSaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoSaveResponse.Merge(m, src)
}
func (m *DoSaveResponse) XXX_Size() int {
	return xxx_messageInfo_DoSaveResponse.Size(m)
}
func (m *DoSaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoSaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoSaveResponse proto.InternalMessageInfo

func (m *DoSaveResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type DoDeleteRequest struct {
	ContactSystemCode    string   `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoDeleteRequest) Reset()         { *m = DoDeleteRequest{} }
func (m *DoDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DoDeleteRequest) ProtoMessage()    {}
func (*DoDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{7}
}

func (m *DoDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoDeleteRequest.Unmarshal(m, b)
}
func (m *DoDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoDeleteRequest.Marshal(b, m, deterministic)
}
func (m *DoDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoDeleteRequest.Merge(m, src)
}
func (m *DoDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DoDeleteRequest.Size(m)
}
func (m *DoDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoDeleteRequest proto.InternalMessageInfo

func (m *DoDeleteRequest) GetContactSystemCode() string {
	if m != nil {
		return m.ContactSystemCode
	}
	return ""
}

type DoDeleteResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoDeleteResponse) Reset()         { *m = DoDeleteResponse{} }
func (m *DoDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DoDeleteResponse) ProtoMessage()    {}
func (*DoDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc89358642882005, []int{8}
}

func (m *DoDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoDeleteResponse.Unmarshal(m, b)
}
func (m *DoDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoDeleteResponse.Marshal(b, m, deterministic)
}
func (m *DoDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoDeleteResponse.Merge(m, src)
}
func (m *DoDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DoDeleteResponse.Size(m)
}
func (m *DoDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoDeleteResponse proto.InternalMessageInfo

func (m *DoDeleteResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*ContactSystem)(nil), "v1.ContactSystem")
	proto.RegisterType((*DoReadRequest)(nil), "v1.DoReadRequest")
	proto.RegisterType((*DoReadResponse)(nil), "v1.DoReadResponse")
	proto.RegisterType((*DoReadAllRequest)(nil), "v1.DoReadAllRequest")
	proto.RegisterType((*DoReadAllResponse)(nil), "v1.DoReadAllResponse")
	proto.RegisterType((*DoSaveRequest)(nil), "v1.DoSaveRequest")
	proto.RegisterType((*DoSaveResponse)(nil), "v1.DoSaveResponse")
	proto.RegisterType((*DoDeleteRequest)(nil), "v1.DoDeleteRequest")
	proto.RegisterType((*DoDeleteResponse)(nil), "v1.DoDeleteResponse")
}

func init() { proto.RegisterFile("contactsystem.proto", fileDescriptor_dc89358642882005) }

var fileDescriptor_dc89358642882005 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x25, 0xef, 0xf9, 0xe2, 0xeb, 0x0d, 0xed, 0xb3, 0xf3, 0xea, 0x23, 0x64, 0x63, 0xc8, 0xaa,
	0x08, 0x26, 0xb4, 0x22, 0x88, 0x1b, 0x89, 0xed, 0x46, 0x37, 0x42, 0xba, 0x73, 0x53, 0xa6, 0xc9,
	0x25, 0x0d, 0x26, 0x99, 0x98, 0x99, 0x04, 0xfa, 0x5f, 0xfe, 0x94, 0x7f, 0x21, 0x99, 0x49, 0xda,
	0x8c, 0x88, 0x60, 0x97, 0x73, 0xee, 0xbd, 0xe7, 0x9c, 0x7b, 0x2e, 0x03, 0x8f, 0x31, 0x2b, 0x05,
	0x8d, 0x05, 0x3f, 0x71, 0x81, 0x85, 0x5f, 0xd5, 0x4c, 0x30, 0x72, 0xd3, 0xae, 0x1c, 0x8b, 0x36,
	0x49, 0x26, 0x14, 0xe0, 0xfd, 0x34, 0x60, 0xba, 0x51, 0x8d, 0x3b, 0xd9, 0x48, 0xfc, 0xf3, 0xe4,
	0x5e, 0x8d, 0xee, 0x63, 0x96, 0xa0, 0x6d, 0xb8, 0xc6, 0x72, 0x12, 0xcd, 0xe3, 0x71, 0xef, 0x86,
	0x25, 0x48, 0x5c, 0xb0, 0x12, 0xe4, 0x71, 0x9d, 0x55, 0x22, 0x63, 0xa5, 0x7d, 0x23, 0xfb, 0xc6,
	0x10, 0xb1, 0xe1, 0x79, 0x82, 0x82, 0x66, 0x39, 0xb7, 0x6f, 0x65, 0x75, 0x78, 0x92, 0x27, 0x30,
	0xb9, 0xa0, 0xa2, 0xe1, 0xf6, 0x33, 0x59, 0xe8, 0x5f, 0xe4, 0x15, 0xdc, 0x49, 0x93, 0xf6, 0x9d,
	0x6b, 0x2c, 0xad, 0xf5, 0xc4, 0x6f, 0x57, 0x7e, 0xd8, 0x01, 0x91, 0xc2, 0xbd, 0x8f, 0x30, 0xdd,
	0xb2, 0x08, 0x69, 0x12, 0xe1, 0x8f, 0x06, 0xb9, 0xf8, 0x5f, 0xd7, 0xde, 0x17, 0x98, 0x0d, 0x04,
	0xbc, 0x62, 0x25, 0x47, 0xf2, 0x1e, 0x66, 0x3a, 0x83, 0x1c, 0xb6, 0xd6, 0xf3, 0x4e, 0x5c, 0x8b,
	0x28, 0x9a, 0x6a, 0x7c, 0x1e, 0x81, 0x17, 0x8a, 0x2b, 0xcc, 0xf3, 0xde, 0x8f, 0xf7, 0x15, 0xe6,
	0x23, 0xac, 0x97, 0xf8, 0x00, 0x0f, 0xba, 0x04, 0xb7, 0x0d, 0xf7, 0xf6, 0xef, 0x1a, 0x33, 0x4d,
	0x83, 0x7b, 0x9f, 0xbb, 0x8d, 0x77, 0xb4, 0xc5, 0x61, 0xe3, 0xeb, 0xfd, 0x2e, 0xbb, 0xdd, 0x15,
	0x55, 0x6f, 0xec, 0x09, 0xcc, 0x1a, 0x79, 0x93, 0x0b, 0xc9, 0x71, 0x1f, 0xf5, 0x2f, 0x2f, 0x84,
	0x87, 0x2d, 0xdb, 0x62, 0x8e, 0x02, 0xaf, 0x0d, 0xfa, 0x75, 0x17, 0xce, 0x40, 0xf1, 0x6f, 0xb9,
	0xf5, 0x2f, 0x03, 0x16, 0x9a, 0xf3, 0x1d, 0xd6, 0x6d, 0x16, 0x23, 0x09, 0xc0, 0x54, 0x69, 0x12,
	0xb9, 0x9d, 0x76, 0x7a, 0x87, 0x8c, 0xa1, 0xf3, 0x31, 0x27, 0xe7, 0xf8, 0xc9, 0xe2, 0xd2, 0x70,
	0xb9, 0x90, 0xf3, 0xf2, 0x0f, 0xb4, 0x9f, 0x94, 0x52, 0x5d, 0x38, 0x83, 0xd4, 0x28, 0xf3, 0x41,
	0x4a, 0xcb, 0xee, 0x1d, 0xdc, 0x0f, 0x0b, 0x92, 0x47, 0x55, 0xd7, 0x12, 0x73, 0x16, 0x3a, 0xa8,
	0xc6, 0x3e, 0x6d, 0xbe, 0x85, 0x69, 0x26, 0x8e, 0xcd, 0xc1, 0x8f, 0x59, 0x11, 0x1c, 0x9a, 0x32,
	0x3d, 0xf1, 0x23, 0x62, 0x15, 0xf4, 0x11, 0xbe, 0x29, 0x68, 0x49, 0x53, 0x2c, 0xb0, 0x14, 0x41,
	0xf5, 0x3d, 0x0d, 0x68, 0x95, 0x05, 0xed, 0x2a, 0xd0, 0x3e, 0xf5, 0xc1, 0x94, 0x9f, 0xf8, 0xed,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0x61, 0xfd, 0x77, 0xec, 0x03, 0x00, 0x00,
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
	DoRead(ctx context.Context, in *DoReadRequest, opts ...grpc.CallOption) (*DoReadResponse, error)
	DoReadAll(ctx context.Context, in *DoReadAllRequest, opts ...grpc.CallOption) (*DoReadAllResponse, error)
	DoSave(ctx context.Context, in *DoSaveRequest, opts ...grpc.CallOption) (*DoSaveResponse, error)
	DoDelete(ctx context.Context, in *DoDeleteRequest, opts ...grpc.CallOption) (*DoDeleteResponse, error)
}

type contactSystemServiceClient struct {
	cc *grpc.ClientConn
}

func NewContactSystemServiceClient(cc *grpc.ClientConn) ContactSystemServiceClient {
	return &contactSystemServiceClient{cc}
}

func (c *contactSystemServiceClient) DoRead(ctx context.Context, in *DoReadRequest, opts ...grpc.CallOption) (*DoReadResponse, error) {
	out := new(DoReadResponse)
	err := c.cc.Invoke(ctx, "/v1.ContactSystemService/DoRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactSystemServiceClient) DoReadAll(ctx context.Context, in *DoReadAllRequest, opts ...grpc.CallOption) (*DoReadAllResponse, error) {
	out := new(DoReadAllResponse)
	err := c.cc.Invoke(ctx, "/v1.ContactSystemService/DoReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactSystemServiceClient) DoSave(ctx context.Context, in *DoSaveRequest, opts ...grpc.CallOption) (*DoSaveResponse, error) {
	out := new(DoSaveResponse)
	err := c.cc.Invoke(ctx, "/v1.ContactSystemService/DoSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactSystemServiceClient) DoDelete(ctx context.Context, in *DoDeleteRequest, opts ...grpc.CallOption) (*DoDeleteResponse, error) {
	out := new(DoDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.ContactSystemService/DoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactSystemServiceServer is the server API for ContactSystemService service.
type ContactSystemServiceServer interface {
	DoRead(context.Context, *DoReadRequest) (*DoReadResponse, error)
	DoReadAll(context.Context, *DoReadAllRequest) (*DoReadAllResponse, error)
	DoSave(context.Context, *DoSaveRequest) (*DoSaveResponse, error)
	DoDelete(context.Context, *DoDeleteRequest) (*DoDeleteResponse, error)
}

// UnimplementedContactSystemServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContactSystemServiceServer struct {
}

func (*UnimplementedContactSystemServiceServer) DoRead(ctx context.Context, req *DoReadRequest) (*DoReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoRead not implemented")
}
func (*UnimplementedContactSystemServiceServer) DoReadAll(ctx context.Context, req *DoReadAllRequest) (*DoReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoReadAll not implemented")
}
func (*UnimplementedContactSystemServiceServer) DoSave(ctx context.Context, req *DoSaveRequest) (*DoSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSave not implemented")
}
func (*UnimplementedContactSystemServiceServer) DoDelete(ctx context.Context, req *DoDeleteRequest) (*DoDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoDelete not implemented")
}

func RegisterContactSystemServiceServer(s *grpc.Server, srv ContactSystemServiceServer) {
	s.RegisterService(&_ContactSystemService_serviceDesc, srv)
}

func _ContactSystemService_DoRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoReadRequest)
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
		return srv.(ContactSystemServiceServer).DoRead(ctx, req.(*DoReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactSystemService_DoReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoReadAllRequest)
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
		return srv.(ContactSystemServiceServer).DoReadAll(ctx, req.(*DoReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactSystemService_DoSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoSaveRequest)
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
		return srv.(ContactSystemServiceServer).DoSave(ctx, req.(*DoSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactSystemService_DoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoDeleteRequest)
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
		return srv.(ContactSystemServiceServer).DoDelete(ctx, req.(*DoDeleteRequest))
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

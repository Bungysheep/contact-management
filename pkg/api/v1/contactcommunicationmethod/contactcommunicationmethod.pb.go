// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contactcommunicationmethod.proto

package contactcommunicationmethod

import (
	context "context"
	fmt "fmt"
	audit "github.com/bungysheep/contact-management/pkg/api/v1/audit"
	contactcommunicationmethodfield "github.com/bungysheep/contact-management/pkg/api/v1/contactcommunicationmethodfield"
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

type ContactCommunicationMethod struct {
	ContactSystemCode               string                                                             `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	ContactId                       int64                                                              `protobuf:"varint,2,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	ContactCommunicationMethodId    int64                                                              `protobuf:"varint,3,opt,name=contact_communication_method_id,json=contactCommunicationMethodId,proto3" json:"contact_communication_method_id,omitempty"`
	CommunicationMethodCode         string                                                             `protobuf:"bytes,4,opt,name=communication_method_code,json=communicationMethodCode,proto3" json:"communication_method_code,omitempty"`
	CommunicationMethodLabelCode    string                                                             `protobuf:"bytes,5,opt,name=communication_method_label_code,json=communicationMethodLabelCode,proto3" json:"communication_method_label_code,omitempty"`
	CommunicationMethodLabelCaption string                                                             `protobuf:"bytes,6,opt,name=communication_method_label_caption,json=communicationMethodLabelCaption,proto3" json:"communication_method_label_caption,omitempty"`
	FormatValue                     string                                                             `protobuf:"bytes,7,opt,name=format_value,json=formatValue,proto3" json:"format_value,omitempty"`
	IsDefault                       bool                                                               `protobuf:"varint,8,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty"`
	ContactCommunicationMethodField []*contactcommunicationmethodfield.ContactCommunicationMethodField `protobuf:"bytes,9,rep,name=contact_communication_method_field,json=contactCommunicationMethodField,proto3" json:"contact_communication_method_field,omitempty"`
	Audit                           *audit.Audit                                                       `protobuf:"bytes,10,opt,name=audit,proto3" json:"audit,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                                                           `json:"-"`
	XXX_unrecognized                []byte                                                             `json:"-"`
	XXX_sizecache                   int32                                                              `json:"-"`
}

func (m *ContactCommunicationMethod) Reset()         { *m = ContactCommunicationMethod{} }
func (m *ContactCommunicationMethod) String() string { return proto.CompactTextString(m) }
func (*ContactCommunicationMethod) ProtoMessage()    {}
func (*ContactCommunicationMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e73ba28293f485, []int{0}
}

func (m *ContactCommunicationMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactCommunicationMethod.Unmarshal(m, b)
}
func (m *ContactCommunicationMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactCommunicationMethod.Marshal(b, m, deterministic)
}
func (m *ContactCommunicationMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactCommunicationMethod.Merge(m, src)
}
func (m *ContactCommunicationMethod) XXX_Size() int {
	return xxx_messageInfo_ContactCommunicationMethod.Size(m)
}
func (m *ContactCommunicationMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactCommunicationMethod.DiscardUnknown(m)
}

var xxx_messageInfo_ContactCommunicationMethod proto.InternalMessageInfo

func (m *ContactCommunicationMethod) GetContactSystemCode() string {
	if m != nil {
		return m.ContactSystemCode
	}
	return ""
}

func (m *ContactCommunicationMethod) GetContactId() int64 {
	if m != nil {
		return m.ContactId
	}
	return 0
}

func (m *ContactCommunicationMethod) GetContactCommunicationMethodId() int64 {
	if m != nil {
		return m.ContactCommunicationMethodId
	}
	return 0
}

func (m *ContactCommunicationMethod) GetCommunicationMethodCode() string {
	if m != nil {
		return m.CommunicationMethodCode
	}
	return ""
}

func (m *ContactCommunicationMethod) GetCommunicationMethodLabelCode() string {
	if m != nil {
		return m.CommunicationMethodLabelCode
	}
	return ""
}

func (m *ContactCommunicationMethod) GetCommunicationMethodLabelCaption() string {
	if m != nil {
		return m.CommunicationMethodLabelCaption
	}
	return ""
}

func (m *ContactCommunicationMethod) GetFormatValue() string {
	if m != nil {
		return m.FormatValue
	}
	return ""
}

func (m *ContactCommunicationMethod) GetIsDefault() bool {
	if m != nil {
		return m.IsDefault
	}
	return false
}

func (m *ContactCommunicationMethod) GetContactCommunicationMethodField() []*contactcommunicationmethodfield.ContactCommunicationMethodField {
	if m != nil {
		return m.ContactCommunicationMethodField
	}
	return nil
}

func (m *ContactCommunicationMethod) GetAudit() *audit.Audit {
	if m != nil {
		return m.Audit
	}
	return nil
}

type DoReadContactCommunicationMethodRequest struct {
	ContactSystemCode            string   `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	ContactId                    int64    `protobuf:"varint,2,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	ContactCommunicationMethodId int64    `protobuf:"varint,3,opt,name=contact_communication_method_id,json=contactCommunicationMethodId,proto3" json:"contact_communication_method_id,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *DoReadContactCommunicationMethodRequest) Reset() {
	*m = DoReadContactCommunicationMethodRequest{}
}
func (m *DoReadContactCommunicationMethodRequest) String() string { return proto.CompactTextString(m) }
func (*DoReadContactCommunicationMethodRequest) ProtoMessage()    {}
func (*DoReadContactCommunicationMethodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e73ba28293f485, []int{1}
}

func (m *DoReadContactCommunicationMethodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReadContactCommunicationMethodRequest.Unmarshal(m, b)
}
func (m *DoReadContactCommunicationMethodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReadContactCommunicationMethodRequest.Marshal(b, m, deterministic)
}
func (m *DoReadContactCommunicationMethodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReadContactCommunicationMethodRequest.Merge(m, src)
}
func (m *DoReadContactCommunicationMethodRequest) XXX_Size() int {
	return xxx_messageInfo_DoReadContactCommunicationMethodRequest.Size(m)
}
func (m *DoReadContactCommunicationMethodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReadContactCommunicationMethodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoReadContactCommunicationMethodRequest proto.InternalMessageInfo

func (m *DoReadContactCommunicationMethodRequest) GetContactSystemCode() string {
	if m != nil {
		return m.ContactSystemCode
	}
	return ""
}

func (m *DoReadContactCommunicationMethodRequest) GetContactId() int64 {
	if m != nil {
		return m.ContactId
	}
	return 0
}

func (m *DoReadContactCommunicationMethodRequest) GetContactCommunicationMethodId() int64 {
	if m != nil {
		return m.ContactCommunicationMethodId
	}
	return 0
}

type DoReadContactCommunicationMethodResponse struct {
	ContactCommunicationMethod *ContactCommunicationMethod `protobuf:"bytes,1,opt,name=contact_communication_method,json=contactCommunicationMethod,proto3" json:"contact_communication_method,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                    `json:"-"`
	XXX_unrecognized           []byte                      `json:"-"`
	XXX_sizecache              int32                       `json:"-"`
}

func (m *DoReadContactCommunicationMethodResponse) Reset() {
	*m = DoReadContactCommunicationMethodResponse{}
}
func (m *DoReadContactCommunicationMethodResponse) String() string { return proto.CompactTextString(m) }
func (*DoReadContactCommunicationMethodResponse) ProtoMessage()    {}
func (*DoReadContactCommunicationMethodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e73ba28293f485, []int{2}
}

func (m *DoReadContactCommunicationMethodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReadContactCommunicationMethodResponse.Unmarshal(m, b)
}
func (m *DoReadContactCommunicationMethodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReadContactCommunicationMethodResponse.Marshal(b, m, deterministic)
}
func (m *DoReadContactCommunicationMethodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReadContactCommunicationMethodResponse.Merge(m, src)
}
func (m *DoReadContactCommunicationMethodResponse) XXX_Size() int {
	return xxx_messageInfo_DoReadContactCommunicationMethodResponse.Size(m)
}
func (m *DoReadContactCommunicationMethodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReadContactCommunicationMethodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoReadContactCommunicationMethodResponse proto.InternalMessageInfo

func (m *DoReadContactCommunicationMethodResponse) GetContactCommunicationMethod() *ContactCommunicationMethod {
	if m != nil {
		return m.ContactCommunicationMethod
	}
	return nil
}

type DoReadAllContactCommunicationMethodRequest struct {
	ContactSystemCode    string   `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	ContactId            int64    `protobuf:"varint,2,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoReadAllContactCommunicationMethodRequest) Reset() {
	*m = DoReadAllContactCommunicationMethodRequest{}
}
func (m *DoReadAllContactCommunicationMethodRequest) String() string {
	return proto.CompactTextString(m)
}
func (*DoReadAllContactCommunicationMethodRequest) ProtoMessage() {}
func (*DoReadAllContactCommunicationMethodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e73ba28293f485, []int{3}
}

func (m *DoReadAllContactCommunicationMethodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReadAllContactCommunicationMethodRequest.Unmarshal(m, b)
}
func (m *DoReadAllContactCommunicationMethodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReadAllContactCommunicationMethodRequest.Marshal(b, m, deterministic)
}
func (m *DoReadAllContactCommunicationMethodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReadAllContactCommunicationMethodRequest.Merge(m, src)
}
func (m *DoReadAllContactCommunicationMethodRequest) XXX_Size() int {
	return xxx_messageInfo_DoReadAllContactCommunicationMethodRequest.Size(m)
}
func (m *DoReadAllContactCommunicationMethodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReadAllContactCommunicationMethodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoReadAllContactCommunicationMethodRequest proto.InternalMessageInfo

func (m *DoReadAllContactCommunicationMethodRequest) GetContactSystemCode() string {
	if m != nil {
		return m.ContactSystemCode
	}
	return ""
}

func (m *DoReadAllContactCommunicationMethodRequest) GetContactId() int64 {
	if m != nil {
		return m.ContactId
	}
	return 0
}

type DoReadAllContactCommunicationMethodResponse struct {
	ContactCommunicationMethod []*ContactCommunicationMethod `protobuf:"bytes,1,rep,name=contact_communication_method,json=contactCommunicationMethod,proto3" json:"contact_communication_method,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                      `json:"-"`
	XXX_unrecognized           []byte                        `json:"-"`
	XXX_sizecache              int32                         `json:"-"`
}

func (m *DoReadAllContactCommunicationMethodResponse) Reset() {
	*m = DoReadAllContactCommunicationMethodResponse{}
}
func (m *DoReadAllContactCommunicationMethodResponse) String() string {
	return proto.CompactTextString(m)
}
func (*DoReadAllContactCommunicationMethodResponse) ProtoMessage() {}
func (*DoReadAllContactCommunicationMethodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e73ba28293f485, []int{4}
}

func (m *DoReadAllContactCommunicationMethodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoReadAllContactCommunicationMethodResponse.Unmarshal(m, b)
}
func (m *DoReadAllContactCommunicationMethodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoReadAllContactCommunicationMethodResponse.Marshal(b, m, deterministic)
}
func (m *DoReadAllContactCommunicationMethodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoReadAllContactCommunicationMethodResponse.Merge(m, src)
}
func (m *DoReadAllContactCommunicationMethodResponse) XXX_Size() int {
	return xxx_messageInfo_DoReadAllContactCommunicationMethodResponse.Size(m)
}
func (m *DoReadAllContactCommunicationMethodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoReadAllContactCommunicationMethodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoReadAllContactCommunicationMethodResponse proto.InternalMessageInfo

func (m *DoReadAllContactCommunicationMethodResponse) GetContactCommunicationMethod() []*ContactCommunicationMethod {
	if m != nil {
		return m.ContactCommunicationMethod
	}
	return nil
}

type DoSaveContactCommunicationMethodRequest struct {
	ContactCommunicationMethod *ContactCommunicationMethod `protobuf:"bytes,1,opt,name=contact_communication_method,json=contactCommunicationMethod,proto3" json:"contact_communication_method,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                    `json:"-"`
	XXX_unrecognized           []byte                      `json:"-"`
	XXX_sizecache              int32                       `json:"-"`
}

func (m *DoSaveContactCommunicationMethodRequest) Reset() {
	*m = DoSaveContactCommunicationMethodRequest{}
}
func (m *DoSaveContactCommunicationMethodRequest) String() string { return proto.CompactTextString(m) }
func (*DoSaveContactCommunicationMethodRequest) ProtoMessage()    {}
func (*DoSaveContactCommunicationMethodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e73ba28293f485, []int{5}
}

func (m *DoSaveContactCommunicationMethodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoSaveContactCommunicationMethodRequest.Unmarshal(m, b)
}
func (m *DoSaveContactCommunicationMethodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoSaveContactCommunicationMethodRequest.Marshal(b, m, deterministic)
}
func (m *DoSaveContactCommunicationMethodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoSaveContactCommunicationMethodRequest.Merge(m, src)
}
func (m *DoSaveContactCommunicationMethodRequest) XXX_Size() int {
	return xxx_messageInfo_DoSaveContactCommunicationMethodRequest.Size(m)
}
func (m *DoSaveContactCommunicationMethodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoSaveContactCommunicationMethodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoSaveContactCommunicationMethodRequest proto.InternalMessageInfo

func (m *DoSaveContactCommunicationMethodRequest) GetContactCommunicationMethod() *ContactCommunicationMethod {
	if m != nil {
		return m.ContactCommunicationMethod
	}
	return nil
}

type DoSaveContactCommunicationMethodResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoSaveContactCommunicationMethodResponse) Reset() {
	*m = DoSaveContactCommunicationMethodResponse{}
}
func (m *DoSaveContactCommunicationMethodResponse) String() string { return proto.CompactTextString(m) }
func (*DoSaveContactCommunicationMethodResponse) ProtoMessage()    {}
func (*DoSaveContactCommunicationMethodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e73ba28293f485, []int{6}
}

func (m *DoSaveContactCommunicationMethodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoSaveContactCommunicationMethodResponse.Unmarshal(m, b)
}
func (m *DoSaveContactCommunicationMethodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoSaveContactCommunicationMethodResponse.Marshal(b, m, deterministic)
}
func (m *DoSaveContactCommunicationMethodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoSaveContactCommunicationMethodResponse.Merge(m, src)
}
func (m *DoSaveContactCommunicationMethodResponse) XXX_Size() int {
	return xxx_messageInfo_DoSaveContactCommunicationMethodResponse.Size(m)
}
func (m *DoSaveContactCommunicationMethodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoSaveContactCommunicationMethodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoSaveContactCommunicationMethodResponse proto.InternalMessageInfo

func (m *DoSaveContactCommunicationMethodResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type DoDeleteContactCommunicationMethodRequest struct {
	ContactSystemCode            string   `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	ContactId                    int64    `protobuf:"varint,2,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	ContactCommunicationMethodId int64    `protobuf:"varint,3,opt,name=contact_communication_method_id,json=contactCommunicationMethodId,proto3" json:"contact_communication_method_id,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *DoDeleteContactCommunicationMethodRequest) Reset() {
	*m = DoDeleteContactCommunicationMethodRequest{}
}
func (m *DoDeleteContactCommunicationMethodRequest) String() string { return proto.CompactTextString(m) }
func (*DoDeleteContactCommunicationMethodRequest) ProtoMessage()    {}
func (*DoDeleteContactCommunicationMethodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e73ba28293f485, []int{7}
}

func (m *DoDeleteContactCommunicationMethodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoDeleteContactCommunicationMethodRequest.Unmarshal(m, b)
}
func (m *DoDeleteContactCommunicationMethodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoDeleteContactCommunicationMethodRequest.Marshal(b, m, deterministic)
}
func (m *DoDeleteContactCommunicationMethodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoDeleteContactCommunicationMethodRequest.Merge(m, src)
}
func (m *DoDeleteContactCommunicationMethodRequest) XXX_Size() int {
	return xxx_messageInfo_DoDeleteContactCommunicationMethodRequest.Size(m)
}
func (m *DoDeleteContactCommunicationMethodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoDeleteContactCommunicationMethodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoDeleteContactCommunicationMethodRequest proto.InternalMessageInfo

func (m *DoDeleteContactCommunicationMethodRequest) GetContactSystemCode() string {
	if m != nil {
		return m.ContactSystemCode
	}
	return ""
}

func (m *DoDeleteContactCommunicationMethodRequest) GetContactId() int64 {
	if m != nil {
		return m.ContactId
	}
	return 0
}

func (m *DoDeleteContactCommunicationMethodRequest) GetContactCommunicationMethodId() int64 {
	if m != nil {
		return m.ContactCommunicationMethodId
	}
	return 0
}

type DoDeleteContactCommunicationMethodResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoDeleteContactCommunicationMethodResponse) Reset() {
	*m = DoDeleteContactCommunicationMethodResponse{}
}
func (m *DoDeleteContactCommunicationMethodResponse) String() string {
	return proto.CompactTextString(m)
}
func (*DoDeleteContactCommunicationMethodResponse) ProtoMessage() {}
func (*DoDeleteContactCommunicationMethodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e73ba28293f485, []int{8}
}

func (m *DoDeleteContactCommunicationMethodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoDeleteContactCommunicationMethodResponse.Unmarshal(m, b)
}
func (m *DoDeleteContactCommunicationMethodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoDeleteContactCommunicationMethodResponse.Marshal(b, m, deterministic)
}
func (m *DoDeleteContactCommunicationMethodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoDeleteContactCommunicationMethodResponse.Merge(m, src)
}
func (m *DoDeleteContactCommunicationMethodResponse) XXX_Size() int {
	return xxx_messageInfo_DoDeleteContactCommunicationMethodResponse.Size(m)
}
func (m *DoDeleteContactCommunicationMethodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoDeleteContactCommunicationMethodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoDeleteContactCommunicationMethodResponse proto.InternalMessageInfo

func (m *DoDeleteContactCommunicationMethodResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*ContactCommunicationMethod)(nil), "v1.ContactCommunicationMethod")
	proto.RegisterType((*DoReadContactCommunicationMethodRequest)(nil), "v1.DoReadContactCommunicationMethodRequest")
	proto.RegisterType((*DoReadContactCommunicationMethodResponse)(nil), "v1.DoReadContactCommunicationMethodResponse")
	proto.RegisterType((*DoReadAllContactCommunicationMethodRequest)(nil), "v1.DoReadAllContactCommunicationMethodRequest")
	proto.RegisterType((*DoReadAllContactCommunicationMethodResponse)(nil), "v1.DoReadAllContactCommunicationMethodResponse")
	proto.RegisterType((*DoSaveContactCommunicationMethodRequest)(nil), "v1.DoSaveContactCommunicationMethodRequest")
	proto.RegisterType((*DoSaveContactCommunicationMethodResponse)(nil), "v1.DoSaveContactCommunicationMethodResponse")
	proto.RegisterType((*DoDeleteContactCommunicationMethodRequest)(nil), "v1.DoDeleteContactCommunicationMethodRequest")
	proto.RegisterType((*DoDeleteContactCommunicationMethodResponse)(nil), "v1.DoDeleteContactCommunicationMethodResponse")
}

func init() { proto.RegisterFile("contactcommunicationmethod.proto", fileDescriptor_64e73ba28293f485) }

var fileDescriptor_64e73ba28293f485 = []byte{
	// 590 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x25, 0xdb, 0xdd, 0xba, 0xb9, 0xf5, 0xc5, 0x11, 0x34, 0x86, 0xd5, 0x66, 0x23, 0x62, 0x74,
	0xdd, 0x84, 0xd6, 0x37, 0xdf, 0x76, 0x5b, 0x85, 0xc5, 0x0f, 0x24, 0x05, 0x1f, 0x7c, 0xa9, 0xd3,
	0xe4, 0xb6, 0x0d, 0x26, 0x99, 0xd8, 0x4c, 0x02, 0x8b, 0x3f, 0x41, 0xc1, 0x9f, 0x24, 0xbe, 0xf8,
	0x17, 0xfc, 0x3b, 0x32, 0x33, 0xa9, 0x1f, 0xd8, 0xa4, 0x01, 0x51, 0xf4, 0xb1, 0x77, 0xce, 0xb9,
	0x73, 0xee, 0xb9, 0x67, 0x4a, 0xc0, 0x0a, 0x58, 0xca, 0x69, 0xc0, 0x03, 0x96, 0x24, 0x45, 0x1a,
	0x05, 0x94, 0x47, 0x2c, 0x4d, 0x90, 0x2f, 0x59, 0xe8, 0x66, 0x2b, 0xc6, 0x19, 0xd9, 0x29, 0x07,
	0x66, 0x8f, 0x16, 0x61, 0xc4, 0x55, 0xc1, 0xbc, 0x55, 0x4f, 0x99, 0x47, 0x18, 0x57, 0x3c, 0xfb,
	0xcb, 0x2e, 0x98, 0x23, 0x85, 0x1c, 0xfd, 0x88, 0x7c, 0x2a, 0x91, 0xc4, 0x85, 0xcb, 0x55, 0x9f,
	0x69, 0x7e, 0x9e, 0x73, 0x4c, 0xa6, 0x01, 0x0b, 0xd1, 0xd0, 0x2c, 0xcd, 0xd1, 0xfd, 0x4b, 0xd5,
	0xd1, 0x44, 0x9e, 0x8c, 0x58, 0x88, 0xe4, 0x3a, 0xc0, 0x1a, 0x1f, 0x85, 0xc6, 0x8e, 0xa5, 0x39,
	0x1d, 0x5f, 0xaf, 0x2a, 0x67, 0x21, 0x79, 0x08, 0xfd, 0xf5, 0xf1, 0x4f, 0xba, 0xa6, 0x4a, 0x98,
	0xe0, 0x74, 0x24, 0xe7, 0x20, 0xa8, 0xd5, 0x74, 0x16, 0x92, 0x07, 0x70, 0x6d, 0x23, 0x5d, 0x6a,
	0xdb, 0x95, 0xda, 0xae, 0x06, 0xbf, 0x32, 0xa5, 0x42, 0x29, 0x61, 0x03, 0x37, 0xa6, 0x33, 0x8c,
	0x55, 0x87, 0x3d, 0xd9, 0xe1, 0x60, 0x43, 0x87, 0x27, 0x02, 0x24, 0xdb, 0x3c, 0x06, 0xbb, 0xa9,
	0x0d, 0xcd, 0x44, 0xd1, 0xe8, 0xca, 0x4e, 0xfd, 0xda, 0x4e, 0x0a, 0x46, 0x0e, 0xe1, 0xe2, 0x9c,
	0xad, 0x12, 0xca, 0xa7, 0x25, 0x8d, 0x0b, 0x34, 0x2e, 0x48, 0x5a, 0x4f, 0xd5, 0x5e, 0x88, 0x92,
	0x30, 0x36, 0xca, 0xa7, 0x21, 0xce, 0x69, 0x11, 0x73, 0x63, 0xdf, 0xd2, 0x9c, 0x7d, 0x5f, 0x8f,
	0xf2, 0xb1, 0x2a, 0x90, 0x4c, 0xc8, 0x69, 0x30, 0x56, 0xae, 0xdc, 0xd0, 0xad, 0x8e, 0xd3, 0x1b,
	0xde, 0x74, 0xcb, 0x81, 0x5b, 0xbf, 0xf3, 0x47, 0x02, 0xea, 0xf7, 0x83, 0x66, 0x00, 0xe9, 0xc3,
	0x9e, 0x8c, 0x9b, 0x01, 0x96, 0xe6, 0xf4, 0x86, 0xba, 0x68, 0x7a, 0x22, 0x0a, 0xbe, 0xaa, 0xdb,
	0x1f, 0x35, 0xb8, 0x3d, 0x66, 0x3e, 0xd2, 0xb0, 0xfe, 0x2e, 0x1f, 0xdf, 0x14, 0x98, 0xf3, 0x7f,
	0x33, 0x66, 0xf6, 0x7b, 0x0d, 0x9c, 0xed, 0x13, 0xe4, 0x19, 0x4b, 0x73, 0x24, 0xaf, 0xe0, 0xa0,
	0xe9, 0x4e, 0x39, 0x4b, 0x6f, 0x78, 0xa3, 0xd9, 0x7b, 0xdf, 0xac, 0x17, 0x64, 0xbf, 0x85, 0xbb,
	0x4a, 0xcd, 0x49, 0x1c, 0xff, 0x6d, 0x4b, 0xed, 0x0f, 0x1a, 0x1c, 0xb5, 0xba, 0xbd, 0xb5, 0x1d,
	0x9d, 0xdf, 0xb4, 0xe3, 0x9d, 0xcc, 0xd7, 0x84, 0x96, 0xb8, 0xdd, 0x8c, 0x3f, 0xbf, 0x9c, 0x53,
	0x11, 0x95, 0x6d, 0x62, 0x2a, 0x6f, 0xae, 0x40, 0x77, 0x85, 0xb9, 0x78, 0xc7, 0x9a, 0x7c, 0xc7,
	0xd5, 0x2f, 0xfb, 0x93, 0x06, 0x77, 0xc6, 0x6c, 0x8c, 0x31, 0x72, 0xfc, 0x5f, 0xdf, 0xcc, 0x58,
	0x84, 0x74, 0xfb, 0x08, 0xcd, 0x4e, 0x0c, 0x3f, 0x77, 0xe0, 0xb0, 0x9e, 0x3e, 0xc1, 0x55, 0x19,
	0x05, 0x48, 0x02, 0xe8, 0xaa, 0x48, 0x92, 0x23, 0xb1, 0xb9, 0x96, 0x7f, 0x36, 0xe6, 0xbd, 0x76,
	0xe0, 0x4a, 0x62, 0x0c, 0xfa, 0xb7, 0xdc, 0x13, 0xf7, 0x3b, 0xb5, 0xcd, 0x23, 0x34, 0xbd, 0xd6,
	0xf8, 0xea, 0x36, 0x39, 0x92, 0x88, 0xd1, 0x7a, 0xa4, 0x56, 0xf9, 0x5e, 0x8f, 0xd4, 0x32, 0x7f,
	0x11, 0xec, 0xaf, 0x77, 0x44, 0x8e, 0x15, 0xb3, 0x65, 0xe8, 0x4c, 0xb7, 0x2d, 0x5c, 0x5d, 0x75,
	0xfa, 0xfc, 0xe5, 0xb3, 0x45, 0xc4, 0x97, 0xc5, 0xcc, 0x0d, 0x58, 0xe2, 0xcd, 0x8a, 0x74, 0x71,
	0x9e, 0x2f, 0x11, 0x33, 0xaf, 0x0a, 0xd1, 0x71, 0x42, 0x53, 0xba, 0xc0, 0x04, 0x53, 0xee, 0x65,
	0xaf, 0x17, 0x1e, 0xcd, 0x22, 0xaf, 0x1c, 0x78, 0xf5, 0xdf, 0x2e, 0xb3, 0xae, 0xfc, 0x6e, 0xb9,
	0xff, 0x35, 0x00, 0x00, 0xff, 0xff, 0x13, 0x0e, 0xcc, 0xbb, 0x13, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContactCommunicationMethodServiceClient is the client API for ContactCommunicationMethodService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContactCommunicationMethodServiceClient interface {
	DoRead(ctx context.Context, in *DoReadContactCommunicationMethodRequest, opts ...grpc.CallOption) (*DoReadContactCommunicationMethodResponse, error)
	DoReadAll(ctx context.Context, in *DoReadAllContactCommunicationMethodRequest, opts ...grpc.CallOption) (*DoReadAllContactCommunicationMethodResponse, error)
	DoSave(ctx context.Context, in *DoSaveContactCommunicationMethodRequest, opts ...grpc.CallOption) (*DoSaveContactCommunicationMethodResponse, error)
	DoDelete(ctx context.Context, in *DoDeleteContactCommunicationMethodRequest, opts ...grpc.CallOption) (*DoDeleteContactCommunicationMethodResponse, error)
}

type contactCommunicationMethodServiceClient struct {
	cc *grpc.ClientConn
}

func NewContactCommunicationMethodServiceClient(cc *grpc.ClientConn) ContactCommunicationMethodServiceClient {
	return &contactCommunicationMethodServiceClient{cc}
}

func (c *contactCommunicationMethodServiceClient) DoRead(ctx context.Context, in *DoReadContactCommunicationMethodRequest, opts ...grpc.CallOption) (*DoReadContactCommunicationMethodResponse, error) {
	out := new(DoReadContactCommunicationMethodResponse)
	err := c.cc.Invoke(ctx, "/v1.ContactCommunicationMethodService/DoRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactCommunicationMethodServiceClient) DoReadAll(ctx context.Context, in *DoReadAllContactCommunicationMethodRequest, opts ...grpc.CallOption) (*DoReadAllContactCommunicationMethodResponse, error) {
	out := new(DoReadAllContactCommunicationMethodResponse)
	err := c.cc.Invoke(ctx, "/v1.ContactCommunicationMethodService/DoReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactCommunicationMethodServiceClient) DoSave(ctx context.Context, in *DoSaveContactCommunicationMethodRequest, opts ...grpc.CallOption) (*DoSaveContactCommunicationMethodResponse, error) {
	out := new(DoSaveContactCommunicationMethodResponse)
	err := c.cc.Invoke(ctx, "/v1.ContactCommunicationMethodService/DoSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactCommunicationMethodServiceClient) DoDelete(ctx context.Context, in *DoDeleteContactCommunicationMethodRequest, opts ...grpc.CallOption) (*DoDeleteContactCommunicationMethodResponse, error) {
	out := new(DoDeleteContactCommunicationMethodResponse)
	err := c.cc.Invoke(ctx, "/v1.ContactCommunicationMethodService/DoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactCommunicationMethodServiceServer is the server API for ContactCommunicationMethodService service.
type ContactCommunicationMethodServiceServer interface {
	DoRead(context.Context, *DoReadContactCommunicationMethodRequest) (*DoReadContactCommunicationMethodResponse, error)
	DoReadAll(context.Context, *DoReadAllContactCommunicationMethodRequest) (*DoReadAllContactCommunicationMethodResponse, error)
	DoSave(context.Context, *DoSaveContactCommunicationMethodRequest) (*DoSaveContactCommunicationMethodResponse, error)
	DoDelete(context.Context, *DoDeleteContactCommunicationMethodRequest) (*DoDeleteContactCommunicationMethodResponse, error)
}

// UnimplementedContactCommunicationMethodServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContactCommunicationMethodServiceServer struct {
}

func (*UnimplementedContactCommunicationMethodServiceServer) DoRead(ctx context.Context, req *DoReadContactCommunicationMethodRequest) (*DoReadContactCommunicationMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoRead not implemented")
}
func (*UnimplementedContactCommunicationMethodServiceServer) DoReadAll(ctx context.Context, req *DoReadAllContactCommunicationMethodRequest) (*DoReadAllContactCommunicationMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoReadAll not implemented")
}
func (*UnimplementedContactCommunicationMethodServiceServer) DoSave(ctx context.Context, req *DoSaveContactCommunicationMethodRequest) (*DoSaveContactCommunicationMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSave not implemented")
}
func (*UnimplementedContactCommunicationMethodServiceServer) DoDelete(ctx context.Context, req *DoDeleteContactCommunicationMethodRequest) (*DoDeleteContactCommunicationMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoDelete not implemented")
}

func RegisterContactCommunicationMethodServiceServer(s *grpc.Server, srv ContactCommunicationMethodServiceServer) {
	s.RegisterService(&_ContactCommunicationMethodService_serviceDesc, srv)
}

func _ContactCommunicationMethodService_DoRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoReadContactCommunicationMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactCommunicationMethodServiceServer).DoRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ContactCommunicationMethodService/DoRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactCommunicationMethodServiceServer).DoRead(ctx, req.(*DoReadContactCommunicationMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactCommunicationMethodService_DoReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoReadAllContactCommunicationMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactCommunicationMethodServiceServer).DoReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ContactCommunicationMethodService/DoReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactCommunicationMethodServiceServer).DoReadAll(ctx, req.(*DoReadAllContactCommunicationMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactCommunicationMethodService_DoSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoSaveContactCommunicationMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactCommunicationMethodServiceServer).DoSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ContactCommunicationMethodService/DoSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactCommunicationMethodServiceServer).DoSave(ctx, req.(*DoSaveContactCommunicationMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactCommunicationMethodService_DoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoDeleteContactCommunicationMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactCommunicationMethodServiceServer).DoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ContactCommunicationMethodService/DoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactCommunicationMethodServiceServer).DoDelete(ctx, req.(*DoDeleteContactCommunicationMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContactCommunicationMethodService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ContactCommunicationMethodService",
	HandlerType: (*ContactCommunicationMethodServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoRead",
			Handler:    _ContactCommunicationMethodService_DoRead_Handler,
		},
		{
			MethodName: "DoReadAll",
			Handler:    _ContactCommunicationMethodService_DoReadAll_Handler,
		},
		{
			MethodName: "DoSave",
			Handler:    _ContactCommunicationMethodService_DoSave_Handler,
		},
		{
			MethodName: "DoDelete",
			Handler:    _ContactCommunicationMethodService_DoDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contactcommunicationmethod.proto",
}

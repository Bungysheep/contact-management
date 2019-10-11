// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contactcommunicationmethodfield.proto

package contactcommunicationmethodfield

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ContactCommunicationMethodField struct {
	ContactSystemCode            string   `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	ContactId                    int64    `protobuf:"varint,2,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	ContactCommunicationMethodId int64    `protobuf:"varint,3,opt,name=contact_communication_method_id,json=contactCommunicationMethodId,proto3" json:"contact_communication_method_id,omitempty"`
	CommunicationMethodCode      string   `protobuf:"bytes,4,opt,name=communication_method_code,json=communicationMethodCode,proto3" json:"communication_method_code,omitempty"`
	FieldCode                    string   `protobuf:"bytes,5,opt,name=field_code,json=fieldCode,proto3" json:"field_code,omitempty"`
	FieldValue                   string   `protobuf:"bytes,6,opt,name=field_value,json=fieldValue,proto3" json:"field_value,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *ContactCommunicationMethodField) Reset()         { *m = ContactCommunicationMethodField{} }
func (m *ContactCommunicationMethodField) String() string { return proto.CompactTextString(m) }
func (*ContactCommunicationMethodField) ProtoMessage()    {}
func (*ContactCommunicationMethodField) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcf9983a0d68aa46, []int{0}
}

func (m *ContactCommunicationMethodField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactCommunicationMethodField.Unmarshal(m, b)
}
func (m *ContactCommunicationMethodField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactCommunicationMethodField.Marshal(b, m, deterministic)
}
func (m *ContactCommunicationMethodField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactCommunicationMethodField.Merge(m, src)
}
func (m *ContactCommunicationMethodField) XXX_Size() int {
	return xxx_messageInfo_ContactCommunicationMethodField.Size(m)
}
func (m *ContactCommunicationMethodField) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactCommunicationMethodField.DiscardUnknown(m)
}

var xxx_messageInfo_ContactCommunicationMethodField proto.InternalMessageInfo

func (m *ContactCommunicationMethodField) GetContactSystemCode() string {
	if m != nil {
		return m.ContactSystemCode
	}
	return ""
}

func (m *ContactCommunicationMethodField) GetContactId() int64 {
	if m != nil {
		return m.ContactId
	}
	return 0
}

func (m *ContactCommunicationMethodField) GetContactCommunicationMethodId() int64 {
	if m != nil {
		return m.ContactCommunicationMethodId
	}
	return 0
}

func (m *ContactCommunicationMethodField) GetCommunicationMethodCode() string {
	if m != nil {
		return m.CommunicationMethodCode
	}
	return ""
}

func (m *ContactCommunicationMethodField) GetFieldCode() string {
	if m != nil {
		return m.FieldCode
	}
	return ""
}

func (m *ContactCommunicationMethodField) GetFieldValue() string {
	if m != nil {
		return m.FieldValue
	}
	return ""
}

func init() {
	proto.RegisterType((*ContactCommunicationMethodField)(nil), "v1.ContactCommunicationMethodField")
}

func init() {
	proto.RegisterFile("contactcommunicationmethodfield.proto", fileDescriptor_bcf9983a0d68aa46)
}

var fileDescriptor_bcf9983a0d68aa46 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x69, 0x57, 0x17, 0x1a, 0x4f, 0xd6, 0x83, 0x15, 0x94, 0x2e, 0x82, 0xb0, 0x17, 0x1b,
	0x16, 0x6f, 0x1e, 0x2d, 0x0a, 0x7b, 0xf0, 0xb2, 0x8b, 0x1e, 0xbc, 0x94, 0x34, 0x19, 0xdb, 0xe0,
	0x26, 0x53, 0x6c, 0x5a, 0xd8, 0xdf, 0xe5, 0x1f, 0x94, 0x4e, 0x5a, 0x50, 0x54, 0xbc, 0xbe, 0xf7,
	0xbd, 0xf0, 0x91, 0x61, 0x57, 0x12, 0xad, 0x13, 0xd2, 0x49, 0x34, 0xa6, 0xb3, 0x5a, 0x0a, 0xa7,
	0xd1, 0x1a, 0x70, 0x35, 0xaa, 0x57, 0x0d, 0x3b, 0x95, 0x35, 0xef, 0xe8, 0x30, 0x0e, 0xfb, 0xd5,
	0xe5, 0x47, 0xc8, 0xd2, 0xdc, 0xd3, 0xf9, 0x57, 0xfa, 0x91, 0xe8, 0x87, 0x81, 0x8e, 0x33, 0x76,
	0x32, 0x3e, 0x58, 0xb4, 0xfb, 0xd6, 0x81, 0x29, 0x24, 0x2a, 0x48, 0x82, 0x45, 0xb0, 0x8c, 0x36,
	0xc7, 0x63, 0xb5, 0xa5, 0x26, 0x47, 0x05, 0xf1, 0x05, 0x63, 0x13, 0xaf, 0x55, 0x12, 0x2e, 0x82,
	0xe5, 0x6c, 0x13, 0x8d, 0xc9, 0x5a, 0xc5, 0xf7, 0x2c, 0x9d, 0xea, 0x6f, 0x82, 0x85, 0x37, 0x1c,
	0x36, 0x33, 0xda, 0x9c, 0xcb, 0x3f, 0xc5, 0xd6, 0x2a, 0xbe, 0x65, 0x67, 0xbf, 0xce, 0xc9, 0xed,
	0x80, 0xdc, 0x4e, 0xe5, 0xcf, 0xe5, 0x64, 0x48, 0x1f, 0xe1, 0xe1, 0x43, 0x82, 0x23, 0x4a, 0xa8,
	0x4e, 0xd9, 0x91, 0xaf, 0x7b, 0xb1, 0xeb, 0x20, 0x99, 0x53, 0xef, 0x17, 0xcf, 0x43, 0x72, 0xf7,
	0xf4, 0xb2, 0xad, 0xb4, 0xab, 0xbb, 0x32, 0x93, 0x68, 0x78, 0xd9, 0xd9, 0x6a, 0xdf, 0xd6, 0x00,
	0x0d, 0x1f, 0x8d, 0xaf, 0x8d, 0xb0, 0xa2, 0x02, 0x03, 0xd6, 0xf1, 0xe6, 0xad, 0xe2, 0xa2, 0xd1,
	0xbc, 0x5f, 0xf1, 0x7f, 0xce, 0x52, 0xce, 0xe9, 0x2e, 0x37, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x41, 0xce, 0x2e, 0x8e, 0xc0, 0x01, 0x00, 0x00,
}

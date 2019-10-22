// Code generated by protoc-gen-go. DO NOT EDIT.
// source: communicationmethodlabel.proto

package communicationmethodlabel

import (
	context "context"
	fmt "fmt"
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

type CommunicationMethodLabel struct {
	ContactSystemCode            string   `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	CommunicationMethodCode      string   `protobuf:"bytes,2,opt,name=communication_method_code,json=communicationMethodCode,proto3" json:"communication_method_code,omitempty"`
	CommunicationMethodLabelCode string   `protobuf:"bytes,3,opt,name=communication_method_label_code,json=communicationMethodLabelCode,proto3" json:"communication_method_label_code,omitempty"`
	Caption                      string   `protobuf:"bytes,4,opt,name=caption,proto3" json:"caption,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *CommunicationMethodLabel) Reset()         { *m = CommunicationMethodLabel{} }
func (m *CommunicationMethodLabel) String() string { return proto.CompactTextString(m) }
func (*CommunicationMethodLabel) ProtoMessage()    {}
func (*CommunicationMethodLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_4102a5fd7ad80fc5, []int{0}
}

func (m *CommunicationMethodLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommunicationMethodLabel.Unmarshal(m, b)
}
func (m *CommunicationMethodLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommunicationMethodLabel.Marshal(b, m, deterministic)
}
func (m *CommunicationMethodLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunicationMethodLabel.Merge(m, src)
}
func (m *CommunicationMethodLabel) XXX_Size() int {
	return xxx_messageInfo_CommunicationMethodLabel.Size(m)
}
func (m *CommunicationMethodLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunicationMethodLabel.DiscardUnknown(m)
}

var xxx_messageInfo_CommunicationMethodLabel proto.InternalMessageInfo

func (m *CommunicationMethodLabel) GetContactSystemCode() string {
	if m != nil {
		return m.ContactSystemCode
	}
	return ""
}

func (m *CommunicationMethodLabel) GetCommunicationMethodCode() string {
	if m != nil {
		return m.CommunicationMethodCode
	}
	return ""
}

func (m *CommunicationMethodLabel) GetCommunicationMethodLabelCode() string {
	if m != nil {
		return m.CommunicationMethodLabelCode
	}
	return ""
}

func (m *CommunicationMethodLabel) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

type DoReadRequest struct {
	ContactSystemCode            string   `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	CommunicationMethodCode      string   `protobuf:"bytes,2,opt,name=communication_method_code,json=communicationMethodCode,proto3" json:"communication_method_code,omitempty"`
	CommunicationMethodLabelCode string   `protobuf:"bytes,3,opt,name=communication_method_label_code,json=communicationMethodLabelCode,proto3" json:"communication_method_label_code,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *DoReadRequest) Reset()         { *m = DoReadRequest{} }
func (m *DoReadRequest) String() string { return proto.CompactTextString(m) }
func (*DoReadRequest) ProtoMessage()    {}
func (*DoReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4102a5fd7ad80fc5, []int{1}
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

func (m *DoReadRequest) GetCommunicationMethodCode() string {
	if m != nil {
		return m.CommunicationMethodCode
	}
	return ""
}

func (m *DoReadRequest) GetCommunicationMethodLabelCode() string {
	if m != nil {
		return m.CommunicationMethodLabelCode
	}
	return ""
}

type DoReadResponse struct {
	CommunicationMethodLabel *CommunicationMethodLabel `protobuf:"bytes,1,opt,name=communication_method_label,json=communicationMethodLabel,proto3" json:"communication_method_label,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *DoReadResponse) Reset()         { *m = DoReadResponse{} }
func (m *DoReadResponse) String() string { return proto.CompactTextString(m) }
func (*DoReadResponse) ProtoMessage()    {}
func (*DoReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4102a5fd7ad80fc5, []int{2}
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

func (m *DoReadResponse) GetCommunicationMethodLabel() *CommunicationMethodLabel {
	if m != nil {
		return m.CommunicationMethodLabel
	}
	return nil
}

type DoReadAllRequest struct {
	ContactSystemCode       string   `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	CommunicationMethodCode string   `protobuf:"bytes,2,opt,name=communication_method_code,json=communicationMethodCode,proto3" json:"communication_method_code,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *DoReadAllRequest) Reset()         { *m = DoReadAllRequest{} }
func (m *DoReadAllRequest) String() string { return proto.CompactTextString(m) }
func (*DoReadAllRequest) ProtoMessage()    {}
func (*DoReadAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4102a5fd7ad80fc5, []int{3}
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

func (m *DoReadAllRequest) GetContactSystemCode() string {
	if m != nil {
		return m.ContactSystemCode
	}
	return ""
}

func (m *DoReadAllRequest) GetCommunicationMethodCode() string {
	if m != nil {
		return m.CommunicationMethodCode
	}
	return ""
}

type DoReadAllResponse struct {
	CommunicationMethodLabel []*CommunicationMethodLabel `protobuf:"bytes,1,rep,name=communication_method_label,json=communicationMethodLabel,proto3" json:"communication_method_label,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                    `json:"-"`
	XXX_unrecognized         []byte                      `json:"-"`
	XXX_sizecache            int32                       `json:"-"`
}

func (m *DoReadAllResponse) Reset()         { *m = DoReadAllResponse{} }
func (m *DoReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*DoReadAllResponse) ProtoMessage()    {}
func (*DoReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4102a5fd7ad80fc5, []int{4}
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

func (m *DoReadAllResponse) GetCommunicationMethodLabel() []*CommunicationMethodLabel {
	if m != nil {
		return m.CommunicationMethodLabel
	}
	return nil
}

type DoSaveRequest struct {
	CommunicationMethodLabel *CommunicationMethodLabel `protobuf:"bytes,1,opt,name=communication_method_label,json=communicationMethodLabel,proto3" json:"communication_method_label,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *DoSaveRequest) Reset()         { *m = DoSaveRequest{} }
func (m *DoSaveRequest) String() string { return proto.CompactTextString(m) }
func (*DoSaveRequest) ProtoMessage()    {}
func (*DoSaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4102a5fd7ad80fc5, []int{5}
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

func (m *DoSaveRequest) GetCommunicationMethodLabel() *CommunicationMethodLabel {
	if m != nil {
		return m.CommunicationMethodLabel
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
	return fileDescriptor_4102a5fd7ad80fc5, []int{6}
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
	ContactSystemCode            string   `protobuf:"bytes,1,opt,name=contact_system_code,json=contactSystemCode,proto3" json:"contact_system_code,omitempty"`
	CommunicationMethodCode      string   `protobuf:"bytes,2,opt,name=communication_method_code,json=communicationMethodCode,proto3" json:"communication_method_code,omitempty"`
	CommunicationMethodLabelCode string   `protobuf:"bytes,3,opt,name=communication_method_label_code,json=communicationMethodLabelCode,proto3" json:"communication_method_label_code,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *DoDeleteRequest) Reset()         { *m = DoDeleteRequest{} }
func (m *DoDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DoDeleteRequest) ProtoMessage()    {}
func (*DoDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4102a5fd7ad80fc5, []int{7}
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

func (m *DoDeleteRequest) GetCommunicationMethodCode() string {
	if m != nil {
		return m.CommunicationMethodCode
	}
	return ""
}

func (m *DoDeleteRequest) GetCommunicationMethodLabelCode() string {
	if m != nil {
		return m.CommunicationMethodLabelCode
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
	return fileDescriptor_4102a5fd7ad80fc5, []int{8}
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
	proto.RegisterType((*CommunicationMethodLabel)(nil), "v1.CommunicationMethodLabel")
	proto.RegisterType((*DoReadRequest)(nil), "v1.DoReadRequest")
	proto.RegisterType((*DoReadResponse)(nil), "v1.DoReadResponse")
	proto.RegisterType((*DoReadAllRequest)(nil), "v1.DoReadAllRequest")
	proto.RegisterType((*DoReadAllResponse)(nil), "v1.DoReadAllResponse")
	proto.RegisterType((*DoSaveRequest)(nil), "v1.DoSaveRequest")
	proto.RegisterType((*DoSaveResponse)(nil), "v1.DoSaveResponse")
	proto.RegisterType((*DoDeleteRequest)(nil), "v1.DoDeleteRequest")
	proto.RegisterType((*DoDeleteResponse)(nil), "v1.DoDeleteResponse")
}

func init() { proto.RegisterFile("communicationmethodlabel.proto", fileDescriptor_4102a5fd7ad80fc5) }

var fileDescriptor_4102a5fd7ad80fc5 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0x25, 0xad, 0xd4, 0xf6, 0x8a, 0x3f, 0x9d, 0x56, 0x8d, 0xa1, 0x58, 0xc9, 0xaa, 0x08, 0x26,
	0xb4, 0x22, 0x88, 0x3b, 0x6d, 0xdd, 0x55, 0x17, 0xe9, 0xae, 0x9b, 0x32, 0x99, 0x5c, 0xd2, 0xd0,
	0x24, 0x13, 0x9b, 0x49, 0xa0, 0x1b, 0x1f, 0x4e, 0x7c, 0x17, 0x5f, 0xc0, 0x07, 0x90, 0xcc, 0x24,
	0x36, 0x51, 0x23, 0x1f, 0x7c, 0x7c, 0x7c, 0x74, 0x99, 0xb9, 0xf7, 0xdc, 0x33, 0xe7, 0x9e, 0x93,
	0x81, 0xe7, 0x8c, 0x47, 0x51, 0x16, 0x07, 0x8c, 0x8a, 0x80, 0xc7, 0x11, 0x8a, 0x3d, 0xf7, 0x42,
	0xea, 0x62, 0x68, 0x25, 0x47, 0x2e, 0x38, 0xe9, 0xe4, 0x73, 0xf3, 0x87, 0x06, 0xfa, 0xb2, 0xde,
	0xf6, 0x49, 0xb6, 0xad, 0x8b, 0x36, 0x62, 0xc1, 0x88, 0xf1, 0x58, 0x50, 0x26, 0x76, 0xe9, 0x29,
	0x15, 0x18, 0xed, 0x18, 0xf7, 0x50, 0xd7, 0x5e, 0x68, 0xb3, 0x81, 0x33, 0x2c, 0x4b, 0x1b, 0x59,
	0x59, 0x72, 0x0f, 0xc9, 0x3b, 0x78, 0xd6, 0xa0, 0xdc, 0x29, 0x4e, 0x85, 0xea, 0x48, 0xd4, 0x53,
	0xf6, 0x37, 0x99, 0xc4, 0x7e, 0x84, 0xe9, 0x3f, 0xb1, 0xf2, 0xc2, 0x6a, 0x42, 0x57, 0x4e, 0x98,
	0xb0, 0x96, 0xeb, 0xca, 0x31, 0x3a, 0xdc, 0x65, 0x34, 0x29, 0x2a, 0xfa, 0x1d, 0xd9, 0x5e, 0x7d,
	0x9a, 0xdf, 0x34, 0xb8, 0xbf, 0xe2, 0x0e, 0x52, 0xcf, 0xc1, 0x2f, 0x19, 0xa6, 0xe2, 0x02, 0xe5,
	0x99, 0x21, 0x3c, 0xa8, 0x34, 0xa4, 0x09, 0x8f, 0x53, 0x24, 0x5b, 0x30, 0xda, 0x07, 0x4b, 0x2d,
	0xf7, 0x16, 0x13, 0x2b, 0x9f, 0x5b, 0x6d, 0x2e, 0x3b, 0x7a, 0x1b, 0xa3, 0xf9, 0x15, 0x1e, 0x29,
	0xb6, 0xf7, 0x61, 0x78, 0x0b, 0x4b, 0x33, 0x39, 0x0c, 0x6b, 0xfc, 0x57, 0x14, 0xdc, 0xbd, 0x86,
	0xe0, 0x43, 0x11, 0x91, 0x0d, 0xcd, 0xb1, 0x52, 0x7b, 0x93, 0xdb, 0x9d, 0x15, 0x5e, 0x2a, 0xb2,
	0x52, 0xda, 0x13, 0xe8, 0x1d, 0x31, 0xcd, 0x42, 0x21, 0x27, 0xf7, 0x9d, 0xf2, 0xcb, 0xfc, 0xae,
	0xc1, 0xc3, 0x15, 0x5f, 0x61, 0x88, 0x02, 0x2f, 0x38, 0xbc, 0x2f, 0x8b, 0x38, 0x55, 0x2a, 0xfe,
	0x2f, 0x79, 0xf1, 0x53, 0x83, 0x69, 0xdb, 0x4e, 0x37, 0x78, 0xcc, 0x03, 0x86, 0xc4, 0x86, 0x9e,
	0x8a, 0x07, 0x19, 0x16, 0x16, 0x34, 0x7e, 0x6e, 0x83, 0xd4, 0x8f, 0x4a, 0xb2, 0xb7, 0x30, 0xf8,
	0x9d, 0x27, 0x32, 0x3e, 0x37, 0x9c, 0xe3, 0x6d, 0x3c, 0xfe, 0xe3, 0xb4, 0x44, 0x4a, 0xaa, 0xc2,
	0xab, 0x8a, 0xaa, 0x16, 0x92, 0x8a, 0xaa, 0x61, 0xe5, 0x1b, 0xe8, 0x57, 0x5a, 0xc9, 0x48, 0xd5,
	0x1b, 0xfe, 0x19, 0xe3, 0xe6, 0xa1, 0x82, 0x7d, 0xf8, 0xbc, 0x5d, 0xfb, 0x81, 0xd8, 0x67, 0xae,
	0xc5, 0x78, 0x64, 0xbb, 0x59, 0xec, 0x9f, 0xd2, 0x3d, 0x62, 0x62, 0x97, 0x86, 0xbe, 0x8a, 0x68,
	0x4c, 0x7d, 0x8c, 0x30, 0x16, 0x76, 0x72, 0xf0, 0x6d, 0x9a, 0x04, 0x76, 0x3e, 0xb7, 0xdb, 0x1e,
	0x7a, 0xb7, 0x27, 0x5f, 0xfa, 0xd7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x36, 0x26, 0xda, 0xaa,
	0x0b, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommunicationMethodLabelServiceClient is the client API for CommunicationMethodLabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommunicationMethodLabelServiceClient interface {
	DoRead(ctx context.Context, in *DoReadRequest, opts ...grpc.CallOption) (*DoReadResponse, error)
	DoReadAll(ctx context.Context, in *DoReadAllRequest, opts ...grpc.CallOption) (*DoReadAllResponse, error)
	DoSave(ctx context.Context, in *DoSaveRequest, opts ...grpc.CallOption) (*DoSaveResponse, error)
	DoDelete(ctx context.Context, in *DoDeleteRequest, opts ...grpc.CallOption) (*DoDeleteResponse, error)
}

type communicationMethodLabelServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommunicationMethodLabelServiceClient(cc *grpc.ClientConn) CommunicationMethodLabelServiceClient {
	return &communicationMethodLabelServiceClient{cc}
}

func (c *communicationMethodLabelServiceClient) DoRead(ctx context.Context, in *DoReadRequest, opts ...grpc.CallOption) (*DoReadResponse, error) {
	out := new(DoReadResponse)
	err := c.cc.Invoke(ctx, "/v1.CommunicationMethodLabelService/DoRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationMethodLabelServiceClient) DoReadAll(ctx context.Context, in *DoReadAllRequest, opts ...grpc.CallOption) (*DoReadAllResponse, error) {
	out := new(DoReadAllResponse)
	err := c.cc.Invoke(ctx, "/v1.CommunicationMethodLabelService/DoReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationMethodLabelServiceClient) DoSave(ctx context.Context, in *DoSaveRequest, opts ...grpc.CallOption) (*DoSaveResponse, error) {
	out := new(DoSaveResponse)
	err := c.cc.Invoke(ctx, "/v1.CommunicationMethodLabelService/DoSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationMethodLabelServiceClient) DoDelete(ctx context.Context, in *DoDeleteRequest, opts ...grpc.CallOption) (*DoDeleteResponse, error) {
	out := new(DoDeleteResponse)
	err := c.cc.Invoke(ctx, "/v1.CommunicationMethodLabelService/DoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationMethodLabelServiceServer is the server API for CommunicationMethodLabelService service.
type CommunicationMethodLabelServiceServer interface {
	DoRead(context.Context, *DoReadRequest) (*DoReadResponse, error)
	DoReadAll(context.Context, *DoReadAllRequest) (*DoReadAllResponse, error)
	DoSave(context.Context, *DoSaveRequest) (*DoSaveResponse, error)
	DoDelete(context.Context, *DoDeleteRequest) (*DoDeleteResponse, error)
}

// UnimplementedCommunicationMethodLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommunicationMethodLabelServiceServer struct {
}

func (*UnimplementedCommunicationMethodLabelServiceServer) DoRead(ctx context.Context, req *DoReadRequest) (*DoReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoRead not implemented")
}
func (*UnimplementedCommunicationMethodLabelServiceServer) DoReadAll(ctx context.Context, req *DoReadAllRequest) (*DoReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoReadAll not implemented")
}
func (*UnimplementedCommunicationMethodLabelServiceServer) DoSave(ctx context.Context, req *DoSaveRequest) (*DoSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSave not implemented")
}
func (*UnimplementedCommunicationMethodLabelServiceServer) DoDelete(ctx context.Context, req *DoDeleteRequest) (*DoDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoDelete not implemented")
}

func RegisterCommunicationMethodLabelServiceServer(s *grpc.Server, srv CommunicationMethodLabelServiceServer) {
	s.RegisterService(&_CommunicationMethodLabelService_serviceDesc, srv)
}

func _CommunicationMethodLabelService_DoRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationMethodLabelServiceServer).DoRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CommunicationMethodLabelService/DoRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationMethodLabelServiceServer).DoRead(ctx, req.(*DoReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationMethodLabelService_DoReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationMethodLabelServiceServer).DoReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CommunicationMethodLabelService/DoReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationMethodLabelServiceServer).DoReadAll(ctx, req.(*DoReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationMethodLabelService_DoSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationMethodLabelServiceServer).DoSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CommunicationMethodLabelService/DoSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationMethodLabelServiceServer).DoSave(ctx, req.(*DoSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationMethodLabelService_DoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationMethodLabelServiceServer).DoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.CommunicationMethodLabelService/DoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationMethodLabelServiceServer).DoDelete(ctx, req.(*DoDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommunicationMethodLabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.CommunicationMethodLabelService",
	HandlerType: (*CommunicationMethodLabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoRead",
			Handler:    _CommunicationMethodLabelService_DoRead_Handler,
		},
		{
			MethodName: "DoReadAll",
			Handler:    _CommunicationMethodLabelService_DoReadAll_Handler,
		},
		{
			MethodName: "DoSave",
			Handler:    _CommunicationMethodLabelService_DoSave_Handler,
		},
		{
			MethodName: "DoDelete",
			Handler:    _CommunicationMethodLabelService_DoDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communicationmethodlabel.proto",
}

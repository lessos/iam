// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package iamapi

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

type MsgItem struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Action               uint32   `protobuf:"varint,2,opt,name=action" json:"action,omitempty"`
	ToUser               string   `protobuf:"bytes,3,opt,name=to_user,json=toUser" json:"to_user,omitempty"`
	ToEmail              string   `protobuf:"bytes,4,opt,name=to_email,json=toEmail" json:"to_email,omitempty"`
	FromUser             string   `protobuf:"bytes,5,opt,name=from_user,json=fromUser" json:"from_user,omitempty"`
	FromEmail            string   `protobuf:"bytes,6,opt,name=from_email,json=fromEmail" json:"from_email,omitempty"`
	Priority             uint32   `protobuf:"varint,7,opt,name=priority" json:"priority,omitempty"`
	Title                string   `protobuf:"bytes,8,opt,name=title" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,9,opt,name=body" json:"body,omitempty"`
	Retry                uint32   `protobuf:"varint,10,opt,name=retry" json:"retry,omitempty"`
	Log                  string   `protobuf:"bytes,12,opt,name=log" json:"log,omitempty"`
	Posted               uint32   `protobuf:"varint,13,opt,name=posted" json:"posted,omitempty"`
	Created              uint32   `protobuf:"varint,14,opt,name=created" json:"created,omitempty"`
	Updated              uint32   `protobuf:"varint,15,opt,name=updated" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgItem) Reset()         { *m = MsgItem{} }
func (m *MsgItem) String() string { return proto.CompactTextString(m) }
func (*MsgItem) ProtoMessage()    {}
func (*MsgItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_types_b4f1893ae3f38817, []int{0}
}
func (m *MsgItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgItem.Unmarshal(m, b)
}
func (m *MsgItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgItem.Marshal(b, m, deterministic)
}
func (dst *MsgItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgItem.Merge(dst, src)
}
func (m *MsgItem) XXX_Size() int {
	return xxx_messageInfo_MsgItem.Size(m)
}
func (m *MsgItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgItem.DiscardUnknown(m)
}

var xxx_messageInfo_MsgItem proto.InternalMessageInfo

func (m *MsgItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgItem) GetAction() uint32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *MsgItem) GetToUser() string {
	if m != nil {
		return m.ToUser
	}
	return ""
}

func (m *MsgItem) GetToEmail() string {
	if m != nil {
		return m.ToEmail
	}
	return ""
}

func (m *MsgItem) GetFromUser() string {
	if m != nil {
		return m.FromUser
	}
	return ""
}

func (m *MsgItem) GetFromEmail() string {
	if m != nil {
		return m.FromEmail
	}
	return ""
}

func (m *MsgItem) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *MsgItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgItem) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MsgItem) GetRetry() uint32 {
	if m != nil {
		return m.Retry
	}
	return 0
}

func (m *MsgItem) GetLog() string {
	if m != nil {
		return m.Log
	}
	return ""
}

func (m *MsgItem) GetPosted() uint32 {
	if m != nil {
		return m.Posted
	}
	return 0
}

func (m *MsgItem) GetCreated() uint32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *MsgItem) GetUpdated() uint32 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgItem)(nil), "iamapi.MsgItem")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_types_b4f1893ae3f38817) }

var fileDescriptor_types_b4f1893ae3f38817 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xbf, 0x4e, 0x03, 0x31,
	0x0c, 0xc6, 0x75, 0xd7, 0xf6, 0xfe, 0x18, 0x5a, 0x90, 0x85, 0xc0, 0x80, 0x90, 0x2a, 0xa6, 0x4e,
	0x2c, 0x3c, 0x03, 0x03, 0x03, 0x4b, 0x25, 0xe6, 0x2a, 0x6d, 0x42, 0x15, 0xe9, 0x0e, 0x47, 0x89,
	0x3b, 0xdc, 0xb3, 0xf2, 0x32, 0x28, 0x4e, 0x61, 0xf3, 0xcf, 0xbf, 0xef, 0x3b, 0x9d, 0x03, 0x17,
	0x32, 0x05, 0x97, 0x5e, 0x42, 0x64, 0x61, 0x6c, 0xbc, 0x19, 0x4d, 0xf0, 0xcf, 0x3f, 0x35, 0xb4,
	0x1f, 0xe9, 0xf8, 0x2e, 0x6e, 0xc4, 0x15, 0xd4, 0xde, 0x52, 0xb5, 0xae, 0x36, 0xfd, 0xb6, 0xf6,
	0x16, 0x6f, 0xa1, 0x31, 0x07, 0xf1, 0xfc, 0x4d, 0xf5, 0xba, 0xda, 0x2c, 0xb7, 0x67, 0xc2, 0x3b,
	0x68, 0x85, 0x77, 0xa7, 0xe4, 0x22, 0xcd, 0x34, 0xdc, 0x08, 0x7f, 0x26, 0x17, 0xf1, 0x1e, 0x3a,
	0xe1, 0x9d, 0x1b, 0x8d, 0x1f, 0x68, 0xae, 0xa6, 0x15, 0x7e, 0xcb, 0x88, 0x8f, 0xd0, 0x7f, 0x45,
	0x1e, 0x4b, 0x6b, 0xa1, 0xae, 0xcb, 0x0b, 0xed, 0x3d, 0x01, 0xa8, 0x2c, 0xcd, 0x46, 0xad, 0xc6,
	0x4b, 0xf7, 0x01, 0xba, 0x10, 0x3d, 0x47, 0x2f, 0x13, 0xb5, 0xfa, 0x27, 0xff, 0x8c, 0x37, 0xb0,
	0x10, 0x2f, 0x83, 0xa3, 0x4e, 0x5b, 0x05, 0x10, 0x61, 0xbe, 0x67, 0x3b, 0x51, 0xaf, 0x4b, 0x9d,
	0x73, 0x32, 0x3a, 0x89, 0x13, 0x81, 0x7e, 0xa2, 0x00, 0x5e, 0xc3, 0x6c, 0xe0, 0x23, 0x5d, 0x6a,
	0x30, 0x8f, 0xf9, 0xea, 0xc0, 0x49, 0x9c, 0xa5, 0x65, 0xb9, 0xba, 0x10, 0x12, 0xb4, 0x87, 0xe8,
	0x4c, 0x16, 0x2b, 0x15, 0x7f, 0x98, 0xcd, 0x29, 0x58, 0x35, 0x57, 0xc5, 0x9c, 0x71, 0xdf, 0xe8,
	0x63, 0xbf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x91, 0xa0, 0x65, 0x51, 0x7b, 0x01, 0x00, 0x00,
}

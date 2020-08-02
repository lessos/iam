// Copyright 2019 Eryx <evorui аt gmail dοt com>, All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: iamauth.proto

package iamauth

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AuthKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKey string `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty" toml:"access_key,omitempty"` // struct:object_slice_key
	SecretKey string `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty" toml:"secret_key,omitempty"`
	User      string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty" toml:"user,omitempty"`
}

func (x *AuthKey) Reset() {
	*x = AuthKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthKey) ProtoMessage() {}

func (x *AuthKey) ProtoReflect() protoreflect.Message {
	mi := &file_iamauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthKey.ProtoReflect.Descriptor instead.
func (*AuthKey) Descriptor() ([]byte, []int) {
	return file_iamauth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthKey) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *AuthKey) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *AuthKey) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type UserPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty"` // struct:object_slice_key
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty"`
	Roles   []uint32 `protobuf:"varint,4,rep,packed,name=roles,proto3" json:"roles,omitempty" toml:"roles,omitempty"`
	Groups  []string `protobuf:"bytes,5,rep,name=groups,proto3" json:"groups,omitempty" toml:"groups,omitempty"`
	Expired int64    `protobuf:"varint,9,opt,name=expired,proto3" json:"expired,omitempty" toml:"expired,omitempty"` // unix time in seconds
}

func (x *UserPayload) Reset() {
	*x = UserPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPayload) ProtoMessage() {}

func (x *UserPayload) ProtoReflect() protoreflect.Message {
	mi := &file_iamauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPayload.ProtoReflect.Descriptor instead.
func (*UserPayload) Descriptor() ([]byte, []int) {
	return file_iamauth_proto_rawDescGZIP(), []int{1}
}

func (x *UserPayload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserPayload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserPayload) GetRoles() []uint32 {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *UserPayload) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *UserPayload) GetExpired() int64 {
	if x != nil {
		return x.Expired
	}
	return 0
}

type AppPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty"`
	User      string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty" toml:"user,omitempty"`
	AccessKey string `protobuf:"bytes,3,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty" toml:"access_key,omitempty"`
	Created   int64  `protobuf:"varint,9,opt,name=created,proto3" json:"created,omitempty" toml:"created,omitempty"` // unix time in milliseconds
}

func (x *AppPayload) Reset() {
	*x = AppPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamauth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppPayload) ProtoMessage() {}

func (x *AppPayload) ProtoReflect() protoreflect.Message {
	mi := &file_iamauth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppPayload.ProtoReflect.Descriptor instead.
func (*AppPayload) Descriptor() ([]byte, []int) {
	return file_iamauth_proto_rawDescGZIP(), []int{2}
}

func (x *AppPayload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppPayload) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *AppPayload) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *AppPayload) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

var File_iamauth_proto protoreflect.FileDescriptor

var file_iamauth_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x61, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x69, 0x61, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x22, 0x5b, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68,
	0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b,
	0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x79, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x22, 0x69, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x0d, 0x48, 0x03, 0x5a,
	0x09, 0x2e, 0x3b, 0x69, 0x61, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_iamauth_proto_rawDescOnce sync.Once
	file_iamauth_proto_rawDescData = file_iamauth_proto_rawDesc
)

func file_iamauth_proto_rawDescGZIP() []byte {
	file_iamauth_proto_rawDescOnce.Do(func() {
		file_iamauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_iamauth_proto_rawDescData)
	})
	return file_iamauth_proto_rawDescData
}

var file_iamauth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_iamauth_proto_goTypes = []interface{}{
	(*AuthKey)(nil),     // 0: iamauth.AuthKey
	(*UserPayload)(nil), // 1: iamauth.UserPayload
	(*AppPayload)(nil),  // 2: iamauth.AppPayload
}
var file_iamauth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_iamauth_proto_init() }
func file_iamauth_proto_init() {
	if File_iamauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iamauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iamauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iamauth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_iamauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_iamauth_proto_goTypes,
		DependencyIndexes: file_iamauth_proto_depIdxs,
		MessageInfos:      file_iamauth_proto_msgTypes,
	}.Build()
	File_iamauth_proto = out.File
	file_iamauth_proto_rawDesc = nil
	file_iamauth_proto_goTypes = nil
	file_iamauth_proto_depIdxs = nil
}

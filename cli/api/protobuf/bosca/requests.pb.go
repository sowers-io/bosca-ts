//
// Copyright 2024 Sowers, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: bosca/requests.proto

package bosca

import (
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

type IntIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IntIdRequest) Reset() {
	*x = IntIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_requests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntIdRequest) ProtoMessage() {}

func (x *IntIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_requests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntIdRequest.ProtoReflect.Descriptor instead.
func (*IntIdRequest) Descriptor() ([]byte, []int) {
	return file_bosca_requests_proto_rawDescGZIP(), []int{0}
}

func (x *IntIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_requests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_requests_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_bosca_requests_proto_rawDescGZIP(), []int{1}
}

func (x *IdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SupplementaryIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *SupplementaryIdRequest) Reset() {
	*x = SupplementaryIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_requests_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplementaryIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplementaryIdRequest) ProtoMessage() {}

func (x *SupplementaryIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_requests_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupplementaryIdRequest.ProtoReflect.Descriptor instead.
func (*SupplementaryIdRequest) Descriptor() ([]byte, []int) {
	return file_bosca_requests_proto_rawDescGZIP(), []int{2}
}

func (x *SupplementaryIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SupplementaryIdRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type IdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdResponse) Reset() {
	*x = IdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_requests_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdResponse) ProtoMessage() {}

func (x *IdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_requests_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdResponse.ProtoReflect.Descriptor instead.
func (*IdResponse) Descriptor() ([]byte, []int) {
	return file_bosca_requests_proto_rawDescGZIP(), []int{3}
}

func (x *IdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type IdsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IdsResponse) Reset() {
	*x = IdsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_requests_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsResponse) ProtoMessage() {}

func (x *IdsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_requests_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdsResponse.ProtoReflect.Descriptor instead.
func (*IdsResponse) Descriptor() ([]byte, []int) {
	return file_bosca_requests_proto_rawDescGZIP(), []int{4}
}

func (x *IdsResponse) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type IdResponsesId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Error *string `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *IdResponsesId) Reset() {
	*x = IdResponsesId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_requests_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdResponsesId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdResponsesId) ProtoMessage() {}

func (x *IdResponsesId) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_requests_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdResponsesId.ProtoReflect.Descriptor instead.
func (*IdResponsesId) Descriptor() ([]byte, []int) {
	return file_bosca_requests_proto_rawDescGZIP(), []int{5}
}

func (x *IdResponsesId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IdResponsesId) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

type IdResponses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []*IdResponsesId `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *IdResponses) Reset() {
	*x = IdResponses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_requests_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdResponses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdResponses) ProtoMessage() {}

func (x *IdResponses) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_requests_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdResponses.ProtoReflect.Descriptor instead.
func (*IdResponses) Descriptor() ([]byte, []int) {
	return file_bosca_requests_proto_rawDescGZIP(), []int{6}
}

func (x *IdResponses) GetId() []*IdResponsesId {
	if x != nil {
		return x.Id
	}
	return nil
}

type Url struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Url) Reset() {
	*x = Url{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_requests_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url) ProtoMessage() {}

func (x *Url) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_requests_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url.ProtoReflect.Descriptor instead.
func (*Url) Descriptor() ([]byte, []int) {
	return file_bosca_requests_proto_rawDescGZIP(), []int{7}
}

func (x *Url) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type IdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *IdsRequest) Reset() {
	*x = IdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_requests_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsRequest) ProtoMessage() {}

func (x *IdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_requests_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdsRequest.ProtoReflect.Descriptor instead.
func (*IdsRequest) Descriptor() ([]byte, []int) {
	return file_bosca_requests_proto_rawDescGZIP(), []int{8}
}

func (x *IdsRequest) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_bosca_requests_proto protoreflect.FileDescriptor

var file_bosca_requests_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x22, 0x1e, 0x0a,
	0x0c, 0x49, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a,
	0x09, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x16, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x1c, 0x0a, 0x0a, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0b, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x44, 0x0a, 0x0d, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x33, 0x0a, 0x0b, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x17, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x1c, 0x0a, 0x0a, 0x49, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x1d, 0x5a, 0x1b, 0x62, 0x6f, 0x73, 0x63, 0x61,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bosca_requests_proto_rawDescOnce sync.Once
	file_bosca_requests_proto_rawDescData = file_bosca_requests_proto_rawDesc
)

func file_bosca_requests_proto_rawDescGZIP() []byte {
	file_bosca_requests_proto_rawDescOnce.Do(func() {
		file_bosca_requests_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosca_requests_proto_rawDescData)
	})
	return file_bosca_requests_proto_rawDescData
}

var file_bosca_requests_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_bosca_requests_proto_goTypes = []any{
	(*IntIdRequest)(nil),           // 0: bosca.IntIdRequest
	(*IdRequest)(nil),              // 1: bosca.IdRequest
	(*SupplementaryIdRequest)(nil), // 2: bosca.SupplementaryIdRequest
	(*IdResponse)(nil),             // 3: bosca.IdResponse
	(*IdsResponse)(nil),            // 4: bosca.IdsResponse
	(*IdResponsesId)(nil),          // 5: bosca.IdResponsesId
	(*IdResponses)(nil),            // 6: bosca.IdResponses
	(*Url)(nil),                    // 7: bosca.Url
	(*IdsRequest)(nil),             // 8: bosca.IdsRequest
}
var file_bosca_requests_proto_depIdxs = []int32{
	5, // 0: bosca.IdResponses.id:type_name -> bosca.IdResponsesId
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bosca_requests_proto_init() }
func file_bosca_requests_proto_init() {
	if File_bosca_requests_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bosca_requests_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IntIdRequest); i {
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
		file_bosca_requests_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*IdRequest); i {
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
		file_bosca_requests_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SupplementaryIdRequest); i {
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
		file_bosca_requests_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*IdResponse); i {
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
		file_bosca_requests_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*IdsResponse); i {
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
		file_bosca_requests_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*IdResponsesId); i {
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
		file_bosca_requests_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*IdResponses); i {
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
		file_bosca_requests_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Url); i {
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
		file_bosca_requests_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*IdsRequest); i {
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
	file_bosca_requests_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bosca_requests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosca_requests_proto_goTypes,
		DependencyIndexes: file_bosca_requests_proto_depIdxs,
		MessageInfos:      file_bosca_requests_proto_msgTypes,
	}.Build()
	File_bosca_requests_proto = out.File
	file_bosca_requests_proto_rawDesc = nil
	file_bosca_requests_proto_goTypes = nil
	file_bosca_requests_proto_depIdxs = nil
}

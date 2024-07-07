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
// source: bosca/content/traits.proto

package content

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

type Traits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Traits []*Trait `protobuf:"bytes,1,rep,name=traits,proto3" json:"traits,omitempty"`
}

func (x *Traits) Reset() {
	*x = Traits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_content_traits_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Traits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Traits) ProtoMessage() {}

func (x *Traits) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_content_traits_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Traits.ProtoReflect.Descriptor instead.
func (*Traits) Descriptor() ([]byte, []int) {
	return file_bosca_content_traits_proto_rawDescGZIP(), []int{0}
}

func (x *Traits) GetTraits() []*Trait {
	if x != nil {
		return x.Traits
	}
	return nil
}

type Trait struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	WorkflowIds []string `protobuf:"bytes,4,rep,name=workflow_ids,json=workflowIds,proto3" json:"workflow_ids,omitempty"`
}

func (x *Trait) Reset() {
	*x = Trait{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_content_traits_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trait) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trait) ProtoMessage() {}

func (x *Trait) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_content_traits_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trait.ProtoReflect.Descriptor instead.
func (*Trait) Descriptor() ([]byte, []int) {
	return file_bosca_content_traits_proto_rawDescGZIP(), []int{1}
}

func (x *Trait) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Trait) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Trait) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Trait) GetWorkflowIds() []string {
	if x != nil {
		return x.WorkflowIds
	}
	return nil
}

type TraitWorkflowIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraitId    string `protobuf:"bytes,1,opt,name=trait_id,json=traitId,proto3" json:"trait_id,omitempty"`
	WorkflowId string `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
}

func (x *TraitWorkflowIdRequest) Reset() {
	*x = TraitWorkflowIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_content_traits_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraitWorkflowIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraitWorkflowIdRequest) ProtoMessage() {}

func (x *TraitWorkflowIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_content_traits_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraitWorkflowIdRequest.ProtoReflect.Descriptor instead.
func (*TraitWorkflowIdRequest) Descriptor() ([]byte, []int) {
	return file_bosca_content_traits_proto_rawDescGZIP(), []int{2}
}

func (x *TraitWorkflowIdRequest) GetTraitId() string {
	if x != nil {
		return x.TraitId
	}
	return ""
}

func (x *TraitWorkflowIdRequest) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

type TraitWorkflowActivityIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraitId    string `protobuf:"bytes,1,opt,name=trait_id,json=traitId,proto3" json:"trait_id,omitempty"`
	WorkflowId string `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	ActivityId string `protobuf:"bytes,3,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
}

func (x *TraitWorkflowActivityIdRequest) Reset() {
	*x = TraitWorkflowActivityIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_content_traits_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraitWorkflowActivityIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraitWorkflowActivityIdRequest) ProtoMessage() {}

func (x *TraitWorkflowActivityIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_content_traits_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraitWorkflowActivityIdRequest.ProtoReflect.Descriptor instead.
func (*TraitWorkflowActivityIdRequest) Descriptor() ([]byte, []int) {
	return file_bosca_content_traits_proto_rawDescGZIP(), []int{3}
}

func (x *TraitWorkflowActivityIdRequest) GetTraitId() string {
	if x != nil {
		return x.TraitId
	}
	return ""
}

func (x *TraitWorkflowActivityIdRequest) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *TraitWorkflowActivityIdRequest) GetActivityId() string {
	if x != nil {
		return x.ActivityId
	}
	return ""
}

var File_bosca_content_traits_proto protoreflect.FileDescriptor

var file_bosca_content_traits_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x62, 0x6f,
	0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x36, 0x0a, 0x06, 0x54,
	0x72, 0x61, 0x69, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x69, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x74, 0x52, 0x06, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x73, 0x22, 0x70, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x49, 0x64, 0x73, 0x22, 0x54, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x69, 0x74, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x7d, 0x0a, 0x1e, 0x54,
	0x72, 0x61, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x72, 0x61, 0x69, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x25, 0x5a, 0x23, 0x62, 0x6f,
	0x73, 0x63, 0x61, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bosca_content_traits_proto_rawDescOnce sync.Once
	file_bosca_content_traits_proto_rawDescData = file_bosca_content_traits_proto_rawDesc
)

func file_bosca_content_traits_proto_rawDescGZIP() []byte {
	file_bosca_content_traits_proto_rawDescOnce.Do(func() {
		file_bosca_content_traits_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosca_content_traits_proto_rawDescData)
	})
	return file_bosca_content_traits_proto_rawDescData
}

var file_bosca_content_traits_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bosca_content_traits_proto_goTypes = []any{
	(*Traits)(nil),                         // 0: bosca.content.Traits
	(*Trait)(nil),                          // 1: bosca.content.Trait
	(*TraitWorkflowIdRequest)(nil),         // 2: bosca.content.TraitWorkflowIdRequest
	(*TraitWorkflowActivityIdRequest)(nil), // 3: bosca.content.TraitWorkflowActivityIdRequest
}
var file_bosca_content_traits_proto_depIdxs = []int32{
	1, // 0: bosca.content.Traits.traits:type_name -> bosca.content.Trait
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bosca_content_traits_proto_init() }
func file_bosca_content_traits_proto_init() {
	if File_bosca_content_traits_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bosca_content_traits_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Traits); i {
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
		file_bosca_content_traits_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Trait); i {
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
		file_bosca_content_traits_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TraitWorkflowIdRequest); i {
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
		file_bosca_content_traits_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TraitWorkflowActivityIdRequest); i {
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
			RawDescriptor: file_bosca_content_traits_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosca_content_traits_proto_goTypes,
		DependencyIndexes: file_bosca_content_traits_proto_depIdxs,
		MessageInfos:      file_bosca_content_traits_proto_msgTypes,
	}.Build()
	File_bosca_content_traits_proto = out.File
	file_bosca_content_traits_proto_rawDesc = nil
	file_bosca_content_traits_proto_goTypes = nil
	file_bosca_content_traits_proto_depIdxs = nil
}

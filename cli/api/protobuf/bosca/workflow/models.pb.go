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
// source: bosca/workflow/models.proto

package workflow

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

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          string            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Name          string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description   string            `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Configuration map[string]string `protobuf:"bytes,5,rep,name=configuration,proto3" json:"configuration,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_workflow_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_workflow_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_bosca_workflow_models_proto_rawDescGZIP(), []int{0}
}

func (x *Model) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Model) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Model) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Model) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Model) GetConfiguration() map[string]string {
	if x != nil {
		return x.Configuration
	}
	return nil
}

type Models struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Models []*Model `protobuf:"bytes,1,rep,name=models,proto3" json:"models,omitempty"`
}

func (x *Models) Reset() {
	*x = Models{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_workflow_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Models) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Models) ProtoMessage() {}

func (x *Models) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_workflow_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Models.ProtoReflect.Descriptor instead.
func (*Models) Descriptor() ([]byte, []int) {
	return file_bosca_workflow_models_proto_rawDescGZIP(), []int{1}
}

func (x *Models) GetModels() []*Model {
	if x != nil {
		return x.Models
	}
	return nil
}

var File_bosca_workflow_models_proto protoreflect.FileDescriptor

var file_bosca_workflow_models_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62,
	0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xf3, 0x01,
	0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x4e, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x40, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x37, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x2d, 0x0a,
	0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x42, 0x26, 0x5a, 0x24,
	0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bosca_workflow_models_proto_rawDescOnce sync.Once
	file_bosca_workflow_models_proto_rawDescData = file_bosca_workflow_models_proto_rawDesc
)

func file_bosca_workflow_models_proto_rawDescGZIP() []byte {
	file_bosca_workflow_models_proto_rawDescOnce.Do(func() {
		file_bosca_workflow_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosca_workflow_models_proto_rawDescData)
	})
	return file_bosca_workflow_models_proto_rawDescData
}

var file_bosca_workflow_models_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bosca_workflow_models_proto_goTypes = []any{
	(*Model)(nil),  // 0: bosca.workflow.Model
	(*Models)(nil), // 1: bosca.workflow.Models
	nil,            // 2: bosca.workflow.Model.ConfigurationEntry
}
var file_bosca_workflow_models_proto_depIdxs = []int32{
	2, // 0: bosca.workflow.Model.configuration:type_name -> bosca.workflow.Model.ConfigurationEntry
	0, // 1: bosca.workflow.Models.models:type_name -> bosca.workflow.Model
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bosca_workflow_models_proto_init() }
func file_bosca_workflow_models_proto_init() {
	if File_bosca_workflow_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bosca_workflow_models_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Model); i {
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
		file_bosca_workflow_models_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Models); i {
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
			RawDescriptor: file_bosca_workflow_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosca_workflow_models_proto_goTypes,
		DependencyIndexes: file_bosca_workflow_models_proto_depIdxs,
		MessageInfos:      file_bosca_workflow_models_proto_msgTypes,
	}.Build()
	File_bosca_workflow_models_proto = out.File
	file_bosca_workflow_models_proto_rawDesc = nil
	file_bosca_workflow_models_proto_goTypes = nil
	file_bosca_workflow_models_proto_depIdxs = nil
}

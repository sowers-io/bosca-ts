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
// source: bosca/ai/table.proto

package search

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

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header []string `protobuf:"bytes,1,rep,name=header,proto3" json:"header,omitempty"`
	Rows   []*Row   `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_ai_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_ai_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_bosca_ai_table_proto_rawDescGZIP(), []int{0}
}

func (x *Table) GetHeader() []string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Table) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Columns []*Column `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_ai_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_ai_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_bosca_ai_table_proto_rawDescGZIP(), []int{1}
}

func (x *Row) GetColumns() []*Column {
	if x != nil {
		return x.Columns
	}
	return nil
}

type Column struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Column) Reset() {
	*x = Column{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_ai_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Column) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Column) ProtoMessage() {}

func (x *Column) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_ai_table_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Column.ProtoReflect.Descriptor instead.
func (*Column) Descriptor() ([]byte, []int) {
	return file_bosca_ai_table_proto_rawDescGZIP(), []int{2}
}

func (x *Column) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_bosca_ai_table_proto protoreflect.FileDescriptor

var file_bosca_ai_table_proto_rawDesc = []byte{
	0x0a, 0x14, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f, 0x61, 0x69, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x61, 0x69,
	0x22, 0x42, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x61, 0x69, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x22, 0x31, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x2a, 0x0a, 0x07, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62,
	0x6f, 0x73, 0x63, 0x61, 0x2e, 0x61, 0x69, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x07,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0x1e, 0x0a, 0x06, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x62, 0x6f, 0x73, 0x63, 0x61,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bosca_ai_table_proto_rawDescOnce sync.Once
	file_bosca_ai_table_proto_rawDescData = file_bosca_ai_table_proto_rawDesc
)

func file_bosca_ai_table_proto_rawDescGZIP() []byte {
	file_bosca_ai_table_proto_rawDescOnce.Do(func() {
		file_bosca_ai_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosca_ai_table_proto_rawDescData)
	})
	return file_bosca_ai_table_proto_rawDescData
}

var file_bosca_ai_table_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_bosca_ai_table_proto_goTypes = []any{
	(*Table)(nil),  // 0: bosca.ai.Table
	(*Row)(nil),    // 1: bosca.ai.Row
	(*Column)(nil), // 2: bosca.ai.Column
}
var file_bosca_ai_table_proto_depIdxs = []int32{
	1, // 0: bosca.ai.Table.rows:type_name -> bosca.ai.Row
	2, // 1: bosca.ai.Row.columns:type_name -> bosca.ai.Column
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bosca_ai_table_proto_init() }
func file_bosca_ai_table_proto_init() {
	if File_bosca_ai_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bosca_ai_table_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Table); i {
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
		file_bosca_ai_table_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Row); i {
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
		file_bosca_ai_table_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Column); i {
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
			RawDescriptor: file_bosca_ai_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosca_ai_table_proto_goTypes,
		DependencyIndexes: file_bosca_ai_table_proto_depIdxs,
		MessageInfos:      file_bosca_ai_table_proto_msgTypes,
	}.Build()
	File_bosca_ai_table_proto = out.File
	file_bosca_ai_table_proto_rawDesc = nil
	file_bosca_ai_table_proto_goTypes = nil
	file_bosca_ai_table_proto_depIdxs = nil
}

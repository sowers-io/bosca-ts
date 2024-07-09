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
// source: bosca/workflow/service.proto

package workflow

import (
	bosca "bosca.io/api/protobuf/bosca"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_bosca_workflow_service_proto protoreflect.FileDescriptor

var file_bosca_workflow_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x62, 0x6f,
	0x73, 0x63, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x6f,
	0x73, 0x63, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x62, 0x6f, 0x73, 0x63, 0x61,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x62,
	0x6f, 0x73, 0x63, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x62, 0x6f, 0x73,
	0x63, 0x61, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x8f, 0x13, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x12, 0x0c, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x16, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x55, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x10, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x12, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x51, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x73, 0x12, 0x0c, 0x2e, 0x62, 0x6f, 0x73,
	0x63, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x73, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x73, 0x12,
	0x58, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x10, 0x2e, 0x62,
	0x6f, 0x73, 0x63, 0x61, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19,
	0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x0c,
	0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x62,
	0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x6d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20,
	0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x80, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x10, 0x2e, 0x62, 0x6f,
	0x73, 0x63, 0x61, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x76, 0x31, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x12, 0x57, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x73, 0x12, 0x0c, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x19, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x5e, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x10, 0x2e, 0x62, 0x6f,
	0x73, 0x63, 0x61, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12,
	0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6e, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x10, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5e, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x12, 0x0c, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1e, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x7d, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x2e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0xd4, 0x01, 0x0a, 0x21,
	0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x29, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x62,
	0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x54, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x4e, 0x12, 0x4c, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x7b, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0xbf, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x73, 0x12, 0x29, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62,
	0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x73, 0x22, 0x4d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x47, 0x12, 0x45, 0x2f,
	0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x7b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x69, 0x64, 0x7d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x7b,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x73, 0x12, 0x9b, 0x01, 0x0a, 0x17, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x12, 0x2e, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x42,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x3a, 0x01, 0x2a, 0x22, 0x37, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x7b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0xaa, 0x01, 0x0a, 0x1a, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x12, 0x31, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x45, 0x3a, 0x01, 0x2a, 0x22, 0x40, 0x2f,
	0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x7b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x8a, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x12, 0x28, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e,
	0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x39, 0x3a, 0x01, 0x2a, 0x22, 0x34, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x7b, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x9e, 0x01, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x2a, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61,
	0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4a, 0x6f, 0x62, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x7d,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x30, 0x01, 0x12, 0xb3, 0x01,
	0x0a, 0x1c, 0x53, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29,
	0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x0c, 0x2e, 0x62, 0x6f, 0x73, 0x63,
	0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x54, 0x3a,
	0x01, 0x2a, 0x22, 0x4f, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x62, 0x6f, 0x73,
	0x63, 0x61, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_bosca_workflow_service_proto_goTypes = []any{
	(*bosca.Empty)(nil),                       // 0: bosca.Empty
	(*bosca.IdRequest)(nil),                   // 1: bosca.IdRequest
	(*WorkflowActivityIdRequest)(nil),         // 2: bosca.workflow.WorkflowActivityIdRequest
	(*BeginTransitionWorkflowRequest)(nil),    // 3: bosca.workflow.BeginTransitionWorkflowRequest
	(*CompleteTransitionWorkflowRequest)(nil), // 4: bosca.workflow.CompleteTransitionWorkflowRequest
	(*WorkflowExecutionRequest)(nil),          // 5: bosca.workflow.WorkflowExecutionRequest
	(*WorkflowActivityJobRequest)(nil),        // 6: bosca.workflow.WorkflowActivityJobRequest
	(*WorkflowActivityJobStatus)(nil),         // 7: bosca.workflow.WorkflowActivityJobStatus
	(*Models)(nil),                            // 8: bosca.workflow.Models
	(*Model)(nil),                             // 9: bosca.workflow.Model
	(*Prompts)(nil),                           // 10: bosca.workflow.Prompts
	(*Prompt)(nil),                            // 11: bosca.workflow.Prompt
	(*StorageSystems)(nil),                    // 12: bosca.workflow.StorageSystems
	(*StorageSystem)(nil),                     // 13: bosca.workflow.StorageSystem
	(*StorageSystemModels)(nil),               // 14: bosca.workflow.StorageSystemModels
	(*Workflows)(nil),                         // 15: bosca.workflow.Workflows
	(*Workflow)(nil),                          // 16: bosca.workflow.Workflow
	(*WorkflowState)(nil),                     // 17: bosca.workflow.WorkflowState
	(*WorkflowStates)(nil),                    // 18: bosca.workflow.WorkflowStates
	(*WorkflowActivities)(nil),                // 19: bosca.workflow.WorkflowActivities
	(*WorkflowActivityStorageSystems)(nil),    // 20: bosca.workflow.WorkflowActivityStorageSystems
	(*WorkflowActivityPrompts)(nil),           // 21: bosca.workflow.WorkflowActivityPrompts
	(*WorkflowActivityJob)(nil),               // 22: bosca.workflow.WorkflowActivityJob
}
var file_bosca_workflow_service_proto_depIdxs = []int32{
	0,  // 0: bosca.workflow.WorkflowService.GetModels:input_type -> bosca.Empty
	1,  // 1: bosca.workflow.WorkflowService.GetModel:input_type -> bosca.IdRequest
	0,  // 2: bosca.workflow.WorkflowService.GetPrompts:input_type -> bosca.Empty
	1,  // 3: bosca.workflow.WorkflowService.GetPrompt:input_type -> bosca.IdRequest
	0,  // 4: bosca.workflow.WorkflowService.GetStorageSystems:input_type -> bosca.Empty
	1,  // 5: bosca.workflow.WorkflowService.GetStorageSystem:input_type -> bosca.IdRequest
	1,  // 6: bosca.workflow.WorkflowService.GetStorageSystemModels:input_type -> bosca.IdRequest
	0,  // 7: bosca.workflow.WorkflowService.GetWorkflows:input_type -> bosca.Empty
	1,  // 8: bosca.workflow.WorkflowService.GetWorkflow:input_type -> bosca.IdRequest
	1,  // 9: bosca.workflow.WorkflowService.GetWorkflowState:input_type -> bosca.IdRequest
	0,  // 10: bosca.workflow.WorkflowService.GetWorkflowStates:input_type -> bosca.Empty
	1,  // 11: bosca.workflow.WorkflowService.GetWorkflowActivities:input_type -> bosca.IdRequest
	2,  // 12: bosca.workflow.WorkflowService.GetWorkflowActivityStorageSystems:input_type -> bosca.workflow.WorkflowActivityIdRequest
	2,  // 13: bosca.workflow.WorkflowService.GetWorkflowActivityPrompts:input_type -> bosca.workflow.WorkflowActivityIdRequest
	3,  // 14: bosca.workflow.WorkflowService.BeginTransitionWorkflow:input_type -> bosca.workflow.BeginTransitionWorkflowRequest
	4,  // 15: bosca.workflow.WorkflowService.CompleteTransitionWorkflow:input_type -> bosca.workflow.CompleteTransitionWorkflowRequest
	5,  // 16: bosca.workflow.WorkflowService.ExecuteWorkflow:input_type -> bosca.workflow.WorkflowExecutionRequest
	6,  // 17: bosca.workflow.WorkflowService.GetWorkflowActivityJobs:input_type -> bosca.workflow.WorkflowActivityJobRequest
	7,  // 18: bosca.workflow.WorkflowService.SetWorkflowActivityJobStatus:input_type -> bosca.workflow.WorkflowActivityJobStatus
	8,  // 19: bosca.workflow.WorkflowService.GetModels:output_type -> bosca.workflow.Models
	9,  // 20: bosca.workflow.WorkflowService.GetModel:output_type -> bosca.workflow.Model
	10, // 21: bosca.workflow.WorkflowService.GetPrompts:output_type -> bosca.workflow.Prompts
	11, // 22: bosca.workflow.WorkflowService.GetPrompt:output_type -> bosca.workflow.Prompt
	12, // 23: bosca.workflow.WorkflowService.GetStorageSystems:output_type -> bosca.workflow.StorageSystems
	13, // 24: bosca.workflow.WorkflowService.GetStorageSystem:output_type -> bosca.workflow.StorageSystem
	14, // 25: bosca.workflow.WorkflowService.GetStorageSystemModels:output_type -> bosca.workflow.StorageSystemModels
	15, // 26: bosca.workflow.WorkflowService.GetWorkflows:output_type -> bosca.workflow.Workflows
	16, // 27: bosca.workflow.WorkflowService.GetWorkflow:output_type -> bosca.workflow.Workflow
	17, // 28: bosca.workflow.WorkflowService.GetWorkflowState:output_type -> bosca.workflow.WorkflowState
	18, // 29: bosca.workflow.WorkflowService.GetWorkflowStates:output_type -> bosca.workflow.WorkflowStates
	19, // 30: bosca.workflow.WorkflowService.GetWorkflowActivities:output_type -> bosca.workflow.WorkflowActivities
	20, // 31: bosca.workflow.WorkflowService.GetWorkflowActivityStorageSystems:output_type -> bosca.workflow.WorkflowActivityStorageSystems
	21, // 32: bosca.workflow.WorkflowService.GetWorkflowActivityPrompts:output_type -> bosca.workflow.WorkflowActivityPrompts
	0,  // 33: bosca.workflow.WorkflowService.BeginTransitionWorkflow:output_type -> bosca.Empty
	0,  // 34: bosca.workflow.WorkflowService.CompleteTransitionWorkflow:output_type -> bosca.Empty
	0,  // 35: bosca.workflow.WorkflowService.ExecuteWorkflow:output_type -> bosca.Empty
	22, // 36: bosca.workflow.WorkflowService.GetWorkflowActivityJobs:output_type -> bosca.workflow.WorkflowActivityJob
	0,  // 37: bosca.workflow.WorkflowService.SetWorkflowActivityJobStatus:output_type -> bosca.Empty
	19, // [19:38] is the sub-list for method output_type
	0,  // [0:19] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_bosca_workflow_service_proto_init() }
func file_bosca_workflow_service_proto_init() {
	if File_bosca_workflow_service_proto != nil {
		return
	}
	file_bosca_workflow_workflows_proto_init()
	file_bosca_workflow_models_proto_init()
	file_bosca_workflow_prompts_proto_init()
	file_bosca_workflow_storage_systems_proto_init()
	file_bosca_workflow_transitions_proto_init()
	file_bosca_workflow_activities_proto_init()
	file_bosca_workflow_execution_context_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bosca_workflow_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bosca_workflow_service_proto_goTypes,
		DependencyIndexes: file_bosca_workflow_service_proto_depIdxs,
	}.Build()
	File_bosca_workflow_service_proto = out.File
	file_bosca_workflow_service_proto_rawDesc = nil
	file_bosca_workflow_service_proto_goTypes = nil
	file_bosca_workflow_service_proto_depIdxs = nil
}

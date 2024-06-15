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
// 	protoc        v5.27.0
// source: bosca/content/metadata.proto

package content

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DefaultId              string                 `protobuf:"bytes,1,opt,name=default_id,json=defaultId,proto3" json:"default_id,omitempty"`
	Id                     string                 `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                   string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ContentType            string                 `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	SourceId               *string                `protobuf:"bytes,5,opt,name=source_id,json=sourceId,proto3,oneof" json:"source_id,omitempty"`
	SourceIdentifier       *string                `protobuf:"bytes,6,opt,name=source_identifier,json=sourceIdentifier,proto3,oneof" json:"source_identifier,omitempty"`
	LanguageTag            string                 `protobuf:"bytes,7,opt,name=language_tag,json=languageTag,proto3" json:"language_tag,omitempty"`
	ContentLength          int64                  `protobuf:"varint,8,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	TraitIds               []string               `protobuf:"bytes,11,rep,name=trait_ids,json=traitIds,proto3" json:"trait_ids,omitempty"`
	CategoryIds            []string               `protobuf:"bytes,12,rep,name=category_ids,json=categoryIds,proto3" json:"category_ids,omitempty"`
	Tags                   []string               `protobuf:"bytes,13,rep,name=tags,proto3" json:"tags,omitempty"`
	Attributes             map[string]string      `protobuf:"bytes,14,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Created                *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=created,proto3" json:"created,omitempty"`
	Modified               *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=modified,proto3" json:"modified,omitempty"`
	WorkflowStateId        string                 `protobuf:"bytes,31,opt,name=workflow_state_id,json=workflowStateId,proto3" json:"workflow_state_id,omitempty"`
	WorkflowStatePendingId *string                `protobuf:"bytes,32,opt,name=workflow_state_pending_id,json=workflowStatePendingId,proto3,oneof" json:"workflow_state_pending_id,omitempty"`
	Metadata               *structpb.Struct       `protobuf:"bytes,33,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_content_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_content_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_bosca_content_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetDefaultId() string {
	if x != nil {
		return x.DefaultId
	}
	return ""
}

func (x *Metadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Metadata) GetSourceId() string {
	if x != nil && x.SourceId != nil {
		return *x.SourceId
	}
	return ""
}

func (x *Metadata) GetSourceIdentifier() string {
	if x != nil && x.SourceIdentifier != nil {
		return *x.SourceIdentifier
	}
	return ""
}

func (x *Metadata) GetLanguageTag() string {
	if x != nil {
		return x.LanguageTag
	}
	return ""
}

func (x *Metadata) GetContentLength() int64 {
	if x != nil {
		return x.ContentLength
	}
	return 0
}

func (x *Metadata) GetTraitIds() []string {
	if x != nil {
		return x.TraitIds
	}
	return nil
}

func (x *Metadata) GetCategoryIds() []string {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

func (x *Metadata) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Metadata) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Metadata) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Metadata) GetModified() *timestamppb.Timestamp {
	if x != nil {
		return x.Modified
	}
	return nil
}

func (x *Metadata) GetWorkflowStateId() string {
	if x != nil {
		return x.WorkflowStateId
	}
	return ""
}

func (x *Metadata) GetWorkflowStatePendingId() string {
	if x != nil && x.WorkflowStatePendingId != nil {
		return *x.WorkflowStatePendingId
	}
	return ""
}

func (x *Metadata) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type AddMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection string    `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Metadata   *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *AddMetadataRequest) Reset() {
	*x = AddMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_content_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMetadataRequest) ProtoMessage() {}

func (x *AddMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_content_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMetadataRequest.ProtoReflect.Descriptor instead.
func (*AddMetadataRequest) Descriptor() ([]byte, []int) {
	return file_bosca_content_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *AddMetadataRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *AddMetadataRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Metadatas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata []*Metadata `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Metadatas) Reset() {
	*x = Metadatas{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_content_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadatas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadatas) ProtoMessage() {}

func (x *Metadatas) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_content_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadatas.ProtoReflect.Descriptor instead.
func (*Metadatas) Descriptor() ([]byte, []int) {
	return file_bosca_content_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *Metadatas) GetMetadata() []*Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type AddMetadataRelationshipRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetadataId1  string `protobuf:"bytes,1,opt,name=metadata_id1,json=metadataId1,proto3" json:"metadata_id1,omitempty"`
	MetadataId2  string `protobuf:"bytes,2,opt,name=metadata_id2,json=metadataId2,proto3" json:"metadata_id2,omitempty"`
	Relationship string `protobuf:"bytes,3,opt,name=relationship,proto3" json:"relationship,omitempty"`
}

func (x *AddMetadataRelationshipRequest) Reset() {
	*x = AddMetadataRelationshipRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_content_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMetadataRelationshipRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMetadataRelationshipRequest) ProtoMessage() {}

func (x *AddMetadataRelationshipRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_content_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMetadataRelationshipRequest.ProtoReflect.Descriptor instead.
func (*AddMetadataRelationshipRequest) Descriptor() ([]byte, []int) {
	return file_bosca_content_metadata_proto_rawDescGZIP(), []int{3}
}

func (x *AddMetadataRelationshipRequest) GetMetadataId1() string {
	if x != nil {
		return x.MetadataId1
	}
	return ""
}

func (x *AddMetadataRelationshipRequest) GetMetadataId2() string {
	if x != nil {
		return x.MetadataId2
	}
	return ""
}

func (x *AddMetadataRelationshipRequest) GetRelationship() string {
	if x != nil {
		return x.Relationship
	}
	return ""
}

type AddMetadataTraitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetadataId string `protobuf:"bytes,1,opt,name=metadata_id,json=metadataId,proto3" json:"metadata_id,omitempty"`
	TraitId    string `protobuf:"bytes,2,opt,name=trait_id,json=traitId,proto3" json:"trait_id,omitempty"`
}

func (x *AddMetadataTraitRequest) Reset() {
	*x = AddMetadataTraitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_content_metadata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMetadataTraitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMetadataTraitRequest) ProtoMessage() {}

func (x *AddMetadataTraitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_content_metadata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMetadataTraitRequest.ProtoReflect.Descriptor instead.
func (*AddMetadataTraitRequest) Descriptor() ([]byte, []int) {
	return file_bosca_content_metadata_proto_rawDescGZIP(), []int{4}
}

func (x *AddMetadataTraitRequest) GetMetadataId() string {
	if x != nil {
		return x.MetadataId
	}
	return ""
}

func (x *AddMetadataTraitRequest) GetTraitId() string {
	if x != nil {
		return x.TraitId
	}
	return ""
}

type AddSupplementaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ContentType   string `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	ContentLength int64  `protobuf:"varint,5,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
}

func (x *AddSupplementaryRequest) Reset() {
	*x = AddSupplementaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_content_metadata_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSupplementaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSupplementaryRequest) ProtoMessage() {}

func (x *AddSupplementaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_content_metadata_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSupplementaryRequest.ProtoReflect.Descriptor instead.
func (*AddSupplementaryRequest) Descriptor() ([]byte, []int) {
	return file_bosca_content_metadata_proto_rawDescGZIP(), []int{5}
}

func (x *AddSupplementaryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AddSupplementaryRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AddSupplementaryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddSupplementaryRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *AddSupplementaryRequest) GetContentLength() int64 {
	if x != nil {
		return x.ContentLength
	}
	return 0
}

type SupplementaryIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *SupplementaryIdRequest) Reset() {
	*x = SupplementaryIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bosca_content_metadata_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SupplementaryIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupplementaryIdRequest) ProtoMessage() {}

func (x *SupplementaryIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bosca_content_metadata_proto_msgTypes[6]
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
	return file_bosca_content_metadata_proto_rawDescGZIP(), []int{6}
}

func (x *SupplementaryIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SupplementaryIdRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_bosca_content_metadata_proto protoreflect.FileDescriptor

var file_bosca_content_metadata_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x06, 0x0a,
	0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x30, 0x0a, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x10,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x61, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x54, 0x61, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x72, 0x61, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x72, 0x61, 0x69, 0x74, 0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x12, 0x47, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x36, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x1f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x19, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x16, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x21, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x1c, 0x0a, 0x1a,
	0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x22, 0x69, 0x0a, 0x12, 0x41, 0x64,
	0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x40, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8a, 0x01, 0x0a, 0x1e, 0x41, 0x64, 0x64, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64, 0x31, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x32, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64, 0x32,
	0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x22, 0x55, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x54, 0x72, 0x61, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x74, 0x49, 0x64, 0x22, 0x9b, 0x01, 0x0a, 0x17,
	0x41, 0x64, 0x64, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x3c, 0x0a, 0x16, 0x53, 0x75, 0x70,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x62, 0x6f, 0x73, 0x63, 0x61,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bosca_content_metadata_proto_rawDescOnce sync.Once
	file_bosca_content_metadata_proto_rawDescData = file_bosca_content_metadata_proto_rawDesc
)

func file_bosca_content_metadata_proto_rawDescGZIP() []byte {
	file_bosca_content_metadata_proto_rawDescOnce.Do(func() {
		file_bosca_content_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_bosca_content_metadata_proto_rawDescData)
	})
	return file_bosca_content_metadata_proto_rawDescData
}

var file_bosca_content_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_bosca_content_metadata_proto_goTypes = []any{
	(*Metadata)(nil),                       // 0: bosca.content.Metadata
	(*AddMetadataRequest)(nil),             // 1: bosca.content.AddMetadataRequest
	(*Metadatas)(nil),                      // 2: bosca.content.Metadatas
	(*AddMetadataRelationshipRequest)(nil), // 3: bosca.content.AddMetadataRelationshipRequest
	(*AddMetadataTraitRequest)(nil),        // 4: bosca.content.AddMetadataTraitRequest
	(*AddSupplementaryRequest)(nil),        // 5: bosca.content.AddSupplementaryRequest
	(*SupplementaryIdRequest)(nil),         // 6: bosca.content.SupplementaryIdRequest
	nil,                                    // 7: bosca.content.Metadata.AttributesEntry
	(*timestamppb.Timestamp)(nil),          // 8: google.protobuf.Timestamp
	(*structpb.Struct)(nil),                // 9: google.protobuf.Struct
}
var file_bosca_content_metadata_proto_depIdxs = []int32{
	7, // 0: bosca.content.Metadata.attributes:type_name -> bosca.content.Metadata.AttributesEntry
	8, // 1: bosca.content.Metadata.created:type_name -> google.protobuf.Timestamp
	8, // 2: bosca.content.Metadata.modified:type_name -> google.protobuf.Timestamp
	9, // 3: bosca.content.Metadata.metadata:type_name -> google.protobuf.Struct
	0, // 4: bosca.content.AddMetadataRequest.metadata:type_name -> bosca.content.Metadata
	0, // 5: bosca.content.Metadatas.metadata:type_name -> bosca.content.Metadata
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_bosca_content_metadata_proto_init() }
func file_bosca_content_metadata_proto_init() {
	if File_bosca_content_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bosca_content_metadata_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Metadata); i {
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
		file_bosca_content_metadata_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddMetadataRequest); i {
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
		file_bosca_content_metadata_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Metadatas); i {
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
		file_bosca_content_metadata_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AddMetadataRelationshipRequest); i {
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
		file_bosca_content_metadata_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AddMetadataTraitRequest); i {
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
		file_bosca_content_metadata_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AddSupplementaryRequest); i {
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
		file_bosca_content_metadata_proto_msgTypes[6].Exporter = func(v any, i int) any {
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
	}
	file_bosca_content_metadata_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bosca_content_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bosca_content_metadata_proto_goTypes,
		DependencyIndexes: file_bosca_content_metadata_proto_depIdxs,
		MessageInfos:      file_bosca_content_metadata_proto_msgTypes,
	}.Build()
	File_bosca_content_metadata_proto = out.File
	file_bosca_content_metadata_proto_rawDesc = nil
	file_bosca_content_metadata_proto_goTypes = nil
	file_bosca_content_metadata_proto_depIdxs = nil
}

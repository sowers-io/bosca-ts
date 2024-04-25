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
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: profiles/profiles.proto

package profiles

import (
	protobuf "bosca.io/api/protobuf"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type ProfileVisibility int32

const (
	ProfileVisibility_system             ProfileVisibility = 0
	ProfileVisibility_user               ProfileVisibility = 1
	ProfileVisibility_friends            ProfileVisibility = 2
	ProfileVisibility_friends_of_friends ProfileVisibility = 3
	ProfileVisibility_public             ProfileVisibility = 4
)

// Enum value maps for ProfileVisibility.
var (
	ProfileVisibility_name = map[int32]string{
		0: "system",
		1: "user",
		2: "friends",
		3: "friends_of_friends",
		4: "public",
	}
	ProfileVisibility_value = map[string]int32{
		"system":             0,
		"user":               1,
		"friends":            2,
		"friends_of_friends": 3,
		"public":             4,
	}
)

func (x ProfileVisibility) Enum() *ProfileVisibility {
	p := new(ProfileVisibility)
	*p = x
	return p
}

func (x ProfileVisibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProfileVisibility) Descriptor() protoreflect.EnumDescriptor {
	return file_profiles_profiles_proto_enumTypes[0].Descriptor()
}

func (ProfileVisibility) Type() protoreflect.EnumType {
	return &file_profiles_profiles_proto_enumTypes[0]
}

func (x ProfileVisibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProfileVisibility.Descriptor instead.
func (ProfileVisibility) EnumDescriptor() ([]byte, []int) {
	return file_profiles_profiles_proto_rawDescGZIP(), []int{0}
}

type ProfileConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarTemplateUrl string `protobuf:"bytes,1,opt,name=avatar_template_url,json=avatarTemplateUrl,proto3" json:"avatar_template_url,omitempty"`
}

func (x *ProfileConfiguration) Reset() {
	*x = ProfileConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_profiles_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileConfiguration) ProtoMessage() {}

func (x *ProfileConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_profiles_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileConfiguration.ProtoReflect.Descriptor instead.
func (*ProfileConfiguration) Descriptor() ([]byte, []int) {
	return file_profiles_profiles_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileConfiguration) GetAvatarTemplateUrl() string {
	if x != nil {
		return x.AvatarTemplateUrl
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Principal  string                 `protobuf:"bytes,2,opt,name=principal,proto3" json:"principal,omitempty"`
	Name       string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Attributes []*ProfileAttribute    `protobuf:"bytes,4,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Visibility ProfileVisibility      `protobuf:"varint,5,opt,name=visibility,proto3,enum=bosca.profiles.ProfileVisibility" json:"visibility,omitempty"`
	Created    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_profiles_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_profiles_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_profiles_profiles_proto_rawDescGZIP(), []int{1}
}

func (x *Profile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Profile) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetAttributes() []*ProfileAttribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Profile) GetVisibility() ProfileVisibility {
	if x != nil {
		return x.Visibility
	}
	return ProfileVisibility_system
}

func (x *Profile) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

type ProfileAttributeType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ProfileAttributeType) Reset() {
	*x = ProfileAttributeType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_profiles_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileAttributeType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileAttributeType) ProtoMessage() {}

func (x *ProfileAttributeType) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_profiles_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileAttributeType.ProtoReflect.Descriptor instead.
func (*ProfileAttributeType) Descriptor() ([]byte, []int) {
	return file_profiles_profiles_proto_rawDescGZIP(), []int{2}
}

func (x *ProfileAttributeType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProfileAttributeType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProfileAttributeType) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ProfileAttributeTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Types []*ProfileAttributeType `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
}

func (x *ProfileAttributeTypes) Reset() {
	*x = ProfileAttributeTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_profiles_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileAttributeTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileAttributeTypes) ProtoMessage() {}

func (x *ProfileAttributeTypes) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_profiles_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileAttributeTypes.ProtoReflect.Descriptor instead.
func (*ProfileAttributeTypes) Descriptor() ([]byte, []int) {
	return file_profiles_profiles_proto_rawDescGZIP(), []int{3}
}

func (x *ProfileAttributeTypes) GetTypes() []*ProfileAttributeType {
	if x != nil {
		return x.Types
	}
	return nil
}

type ProfileAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TypeId     string                 `protobuf:"bytes,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	Visibility ProfileVisibility      `protobuf:"varint,4,opt,name=visibility,proto3,enum=bosca.profiles.ProfileVisibility" json:"visibility,omitempty"`
	Value      *anypb.Any             `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Confidence float32                `protobuf:"fixed32,6,opt,name=confidence,proto3" json:"confidence,omitempty"`
	Priority   float32                `protobuf:"fixed32,7,opt,name=priority,proto3" json:"priority,omitempty"`
	Source     string                 `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
	Created    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created,proto3" json:"created,omitempty"`
	Expiration *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *ProfileAttribute) Reset() {
	*x = ProfileAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profiles_profiles_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileAttribute) ProtoMessage() {}

func (x *ProfileAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_profiles_profiles_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileAttribute.ProtoReflect.Descriptor instead.
func (*ProfileAttribute) Descriptor() ([]byte, []int) {
	return file_profiles_profiles_proto_rawDescGZIP(), []int{4}
}

func (x *ProfileAttribute) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProfileAttribute) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *ProfileAttribute) GetVisibility() ProfileVisibility {
	if x != nil {
		return x.Visibility
	}
	return ProfileVisibility_system
}

func (x *ProfileAttribute) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ProfileAttribute) GetConfidence() float32 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *ProfileAttribute) GetPriority() float32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *ProfileAttribute) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ProfileAttribute) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *ProfileAttribute) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

var File_profiles_profiles_proto protoreflect.FileDescriptor

var file_profiles_profiles_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x6f, 0x73, 0x63, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x46, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x86, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70,
	0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x6f, 0x73,
	0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62,
	0x6f, 0x73, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x22, 0x5c, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x53, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x22, 0xf0, 0x02, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x12, 0x41, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x5a, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x0a, 0x0a, 0x06,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x10, 0x02, 0x12,
	0x16, 0x0a, 0x12, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x73, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x10, 0x04, 0x32, 0xc9, 0x02, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x62, 0x6f,
	0x73, 0x63, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x62, 0x6f, 0x73, 0x63,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x0c, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x2f, 0x6d, 0x79, 0x12, 0x7a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12,
	0x11, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42,
	0x20, 0x5a, 0x1e, 0x62, 0x6f, 0x73, 0x63, 0x61, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profiles_profiles_proto_rawDescOnce sync.Once
	file_profiles_profiles_proto_rawDescData = file_profiles_profiles_proto_rawDesc
)

func file_profiles_profiles_proto_rawDescGZIP() []byte {
	file_profiles_profiles_proto_rawDescOnce.Do(func() {
		file_profiles_profiles_proto_rawDescData = protoimpl.X.CompressGZIP(file_profiles_profiles_proto_rawDescData)
	})
	return file_profiles_profiles_proto_rawDescData
}

var file_profiles_profiles_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_profiles_profiles_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_profiles_profiles_proto_goTypes = []interface{}{
	(ProfileVisibility)(0),        // 0: bosca.profiles.ProfileVisibility
	(*ProfileConfiguration)(nil),  // 1: bosca.profiles.ProfileConfiguration
	(*Profile)(nil),               // 2: bosca.profiles.Profile
	(*ProfileAttributeType)(nil),  // 3: bosca.profiles.ProfileAttributeType
	(*ProfileAttributeTypes)(nil), // 4: bosca.profiles.ProfileAttributeTypes
	(*ProfileAttribute)(nil),      // 5: bosca.profiles.ProfileAttribute
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 7: google.protobuf.Any
	(*protobuf.Empty)(nil),        // 8: bosca.Empty
	(*protobuf.IdsRequest)(nil),   // 9: bosca.IdsRequest
}
var file_profiles_profiles_proto_depIdxs = []int32{
	5,  // 0: bosca.profiles.Profile.attributes:type_name -> bosca.profiles.ProfileAttribute
	0,  // 1: bosca.profiles.Profile.visibility:type_name -> bosca.profiles.ProfileVisibility
	6,  // 2: bosca.profiles.Profile.created:type_name -> google.protobuf.Timestamp
	3,  // 3: bosca.profiles.ProfileAttributeTypes.types:type_name -> bosca.profiles.ProfileAttributeType
	0,  // 4: bosca.profiles.ProfileAttribute.visibility:type_name -> bosca.profiles.ProfileVisibility
	7,  // 5: bosca.profiles.ProfileAttribute.value:type_name -> google.protobuf.Any
	6,  // 6: bosca.profiles.ProfileAttribute.created:type_name -> google.protobuf.Timestamp
	6,  // 7: bosca.profiles.ProfileAttribute.expiration:type_name -> google.protobuf.Timestamp
	8,  // 8: bosca.profiles.ProfilesService.GetConfiguration:input_type -> bosca.Empty
	8,  // 9: bosca.profiles.ProfilesService.GetMyProfile:input_type -> bosca.Empty
	9,  // 10: bosca.profiles.ProfilesService.GetProfileAttributeTypes:input_type -> bosca.IdsRequest
	1,  // 11: bosca.profiles.ProfilesService.GetConfiguration:output_type -> bosca.profiles.ProfileConfiguration
	2,  // 12: bosca.profiles.ProfilesService.GetMyProfile:output_type -> bosca.profiles.Profile
	4,  // 13: bosca.profiles.ProfilesService.GetProfileAttributeTypes:output_type -> bosca.profiles.ProfileAttributeTypes
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_profiles_profiles_proto_init() }
func file_profiles_profiles_proto_init() {
	if File_profiles_profiles_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profiles_profiles_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileConfiguration); i {
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
		file_profiles_profiles_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_profiles_profiles_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileAttributeType); i {
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
		file_profiles_profiles_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileAttributeTypes); i {
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
		file_profiles_profiles_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileAttribute); i {
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
			RawDescriptor: file_profiles_profiles_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profiles_profiles_proto_goTypes,
		DependencyIndexes: file_profiles_profiles_proto_depIdxs,
		EnumInfos:         file_profiles_profiles_proto_enumTypes,
		MessageInfos:      file_profiles_profiles_proto_msgTypes,
	}.Build()
	File_profiles_profiles_proto = out.File
	file_profiles_profiles_proto_rawDesc = nil
	file_profiles_profiles_proto_goTypes = nil
	file_profiles_profiles_proto_depIdxs = nil
}

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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: content/content.proto

package content

import (
	protobuf "bosca.io/api/protobuf"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ContentService_GetRootCollectionItems_FullMethodName   = "/bosca.content.ContentService/GetRootCollectionItems"
	ContentService_AddCollection_FullMethodName            = "/bosca.content.ContentService/AddCollection"
	ContentService_GetCollectionPermissions_FullMethodName = "/bosca.content.ContentService/GetCollectionPermissions"
	ContentService_AddCollectionPermission_FullMethodName  = "/bosca.content.ContentService/AddCollectionPermission"
	ContentService_GetMetadata_FullMethodName              = "/bosca.content.ContentService/GetMetadata"
	ContentService_AddMetadata_FullMethodName              = "/bosca.content.ContentService/AddMetadata"
	ContentService_SetMetadataUploaded_FullMethodName      = "/bosca.content.ContentService/SetMetadataUploaded"
	ContentService_GetMetadataPermissions_FullMethodName   = "/bosca.content.ContentService/GetMetadataPermissions"
	ContentService_AddMetadataPermissions_FullMethodName   = "/bosca.content.ContentService/AddMetadataPermissions"
	ContentService_AddMetadataPermission_FullMethodName    = "/bosca.content.ContentService/AddMetadataPermission"
	ContentService_SetMetadataStatus_FullMethodName        = "/bosca.content.ContentService/SetMetadataStatus"
)

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentServiceClient interface {
	GetRootCollectionItems(ctx context.Context, in *protobuf.Empty, opts ...grpc.CallOption) (*CollectionItems, error)
	AddCollection(ctx context.Context, in *AddCollectionRequest, opts ...grpc.CallOption) (*SignedUrl, error)
	GetCollectionPermissions(ctx context.Context, in *protobuf.IdRequest, opts ...grpc.CallOption) (*Permissions, error)
	AddCollectionPermission(ctx context.Context, in *Permission, opts ...grpc.CallOption) (*protobuf.Empty, error)
	GetMetadata(ctx context.Context, in *protobuf.IdRequest, opts ...grpc.CallOption) (*Metadata, error)
	AddMetadata(ctx context.Context, in *AddMetadataRequest, opts ...grpc.CallOption) (*SignedUrl, error)
	SetMetadataUploaded(ctx context.Context, in *protobuf.IdRequest, opts ...grpc.CallOption) (*protobuf.Empty, error)
	GetMetadataPermissions(ctx context.Context, in *protobuf.IdRequest, opts ...grpc.CallOption) (*Permissions, error)
	AddMetadataPermissions(ctx context.Context, in *Permissions, opts ...grpc.CallOption) (*protobuf.Empty, error)
	AddMetadataPermission(ctx context.Context, in *Permission, opts ...grpc.CallOption) (*protobuf.Empty, error)
	SetMetadataStatus(ctx context.Context, in *SetMetadataStatusRequest, opts ...grpc.CallOption) (*protobuf.Empty, error)
}

type contentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentServiceClient(cc grpc.ClientConnInterface) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) GetRootCollectionItems(ctx context.Context, in *protobuf.Empty, opts ...grpc.CallOption) (*CollectionItems, error) {
	out := new(CollectionItems)
	err := c.cc.Invoke(ctx, ContentService_GetRootCollectionItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) AddCollection(ctx context.Context, in *AddCollectionRequest, opts ...grpc.CallOption) (*SignedUrl, error) {
	out := new(SignedUrl)
	err := c.cc.Invoke(ctx, ContentService_AddCollection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetCollectionPermissions(ctx context.Context, in *protobuf.IdRequest, opts ...grpc.CallOption) (*Permissions, error) {
	out := new(Permissions)
	err := c.cc.Invoke(ctx, ContentService_GetCollectionPermissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) AddCollectionPermission(ctx context.Context, in *Permission, opts ...grpc.CallOption) (*protobuf.Empty, error) {
	out := new(protobuf.Empty)
	err := c.cc.Invoke(ctx, ContentService_AddCollectionPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetMetadata(ctx context.Context, in *protobuf.IdRequest, opts ...grpc.CallOption) (*Metadata, error) {
	out := new(Metadata)
	err := c.cc.Invoke(ctx, ContentService_GetMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) AddMetadata(ctx context.Context, in *AddMetadataRequest, opts ...grpc.CallOption) (*SignedUrl, error) {
	out := new(SignedUrl)
	err := c.cc.Invoke(ctx, ContentService_AddMetadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) SetMetadataUploaded(ctx context.Context, in *protobuf.IdRequest, opts ...grpc.CallOption) (*protobuf.Empty, error) {
	out := new(protobuf.Empty)
	err := c.cc.Invoke(ctx, ContentService_SetMetadataUploaded_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetMetadataPermissions(ctx context.Context, in *protobuf.IdRequest, opts ...grpc.CallOption) (*Permissions, error) {
	out := new(Permissions)
	err := c.cc.Invoke(ctx, ContentService_GetMetadataPermissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) AddMetadataPermissions(ctx context.Context, in *Permissions, opts ...grpc.CallOption) (*protobuf.Empty, error) {
	out := new(protobuf.Empty)
	err := c.cc.Invoke(ctx, ContentService_AddMetadataPermissions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) AddMetadataPermission(ctx context.Context, in *Permission, opts ...grpc.CallOption) (*protobuf.Empty, error) {
	out := new(protobuf.Empty)
	err := c.cc.Invoke(ctx, ContentService_AddMetadataPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) SetMetadataStatus(ctx context.Context, in *SetMetadataStatusRequest, opts ...grpc.CallOption) (*protobuf.Empty, error) {
	out := new(protobuf.Empty)
	err := c.cc.Invoke(ctx, ContentService_SetMetadataStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
// All implementations must embed UnimplementedContentServiceServer
// for forward compatibility
type ContentServiceServer interface {
	GetRootCollectionItems(context.Context, *protobuf.Empty) (*CollectionItems, error)
	AddCollection(context.Context, *AddCollectionRequest) (*SignedUrl, error)
	GetCollectionPermissions(context.Context, *protobuf.IdRequest) (*Permissions, error)
	AddCollectionPermission(context.Context, *Permission) (*protobuf.Empty, error)
	GetMetadata(context.Context, *protobuf.IdRequest) (*Metadata, error)
	AddMetadata(context.Context, *AddMetadataRequest) (*SignedUrl, error)
	SetMetadataUploaded(context.Context, *protobuf.IdRequest) (*protobuf.Empty, error)
	GetMetadataPermissions(context.Context, *protobuf.IdRequest) (*Permissions, error)
	AddMetadataPermissions(context.Context, *Permissions) (*protobuf.Empty, error)
	AddMetadataPermission(context.Context, *Permission) (*protobuf.Empty, error)
	SetMetadataStatus(context.Context, *SetMetadataStatusRequest) (*protobuf.Empty, error)
	mustEmbedUnimplementedContentServiceServer()
}

// UnimplementedContentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContentServiceServer struct {
}

func (UnimplementedContentServiceServer) GetRootCollectionItems(context.Context, *protobuf.Empty) (*CollectionItems, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRootCollectionItems not implemented")
}
func (UnimplementedContentServiceServer) AddCollection(context.Context, *AddCollectionRequest) (*SignedUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCollection not implemented")
}
func (UnimplementedContentServiceServer) GetCollectionPermissions(context.Context, *protobuf.IdRequest) (*Permissions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollectionPermissions not implemented")
}
func (UnimplementedContentServiceServer) AddCollectionPermission(context.Context, *Permission) (*protobuf.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCollectionPermission not implemented")
}
func (UnimplementedContentServiceServer) GetMetadata(context.Context, *protobuf.IdRequest) (*Metadata, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedContentServiceServer) AddMetadata(context.Context, *AddMetadataRequest) (*SignedUrl, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMetadata not implemented")
}
func (UnimplementedContentServiceServer) SetMetadataUploaded(context.Context, *protobuf.IdRequest) (*protobuf.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMetadataUploaded not implemented")
}
func (UnimplementedContentServiceServer) GetMetadataPermissions(context.Context, *protobuf.IdRequest) (*Permissions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadataPermissions not implemented")
}
func (UnimplementedContentServiceServer) AddMetadataPermissions(context.Context, *Permissions) (*protobuf.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMetadataPermissions not implemented")
}
func (UnimplementedContentServiceServer) AddMetadataPermission(context.Context, *Permission) (*protobuf.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMetadataPermission not implemented")
}
func (UnimplementedContentServiceServer) SetMetadataStatus(context.Context, *SetMetadataStatusRequest) (*protobuf.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMetadataStatus not implemented")
}
func (UnimplementedContentServiceServer) mustEmbedUnimplementedContentServiceServer() {}

// UnsafeContentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServiceServer will
// result in compilation errors.
type UnsafeContentServiceServer interface {
	mustEmbedUnimplementedContentServiceServer()
}

func RegisterContentServiceServer(s grpc.ServiceRegistrar, srv ContentServiceServer) {
	s.RegisterService(&ContentService_ServiceDesc, srv)
}

func _ContentService_GetRootCollectionItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetRootCollectionItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_GetRootCollectionItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetRootCollectionItems(ctx, req.(*protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_AddCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AddCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_AddCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AddCollection(ctx, req.(*AddCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetCollectionPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetCollectionPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_GetCollectionPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetCollectionPermissions(ctx, req.(*protobuf.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_AddCollectionPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Permission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AddCollectionPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_AddCollectionPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AddCollectionPermission(ctx, req.(*Permission))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_GetMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetMetadata(ctx, req.(*protobuf.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_AddMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AddMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_AddMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AddMetadata(ctx, req.(*AddMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_SetMetadataUploaded_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).SetMetadataUploaded(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_SetMetadataUploaded_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).SetMetadataUploaded(ctx, req.(*protobuf.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetMetadataPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetMetadataPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_GetMetadataPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetMetadataPermissions(ctx, req.(*protobuf.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_AddMetadataPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Permissions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AddMetadataPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_AddMetadataPermissions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AddMetadataPermissions(ctx, req.(*Permissions))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_AddMetadataPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Permission)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AddMetadataPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_AddMetadataPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AddMetadataPermission(ctx, req.(*Permission))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_SetMetadataStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMetadataStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).SetMetadataStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_SetMetadataStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).SetMetadataStatus(ctx, req.(*SetMetadataStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentService_ServiceDesc is the grpc.ServiceDesc for ContentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosca.content.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRootCollectionItems",
			Handler:    _ContentService_GetRootCollectionItems_Handler,
		},
		{
			MethodName: "AddCollection",
			Handler:    _ContentService_AddCollection_Handler,
		},
		{
			MethodName: "GetCollectionPermissions",
			Handler:    _ContentService_GetCollectionPermissions_Handler,
		},
		{
			MethodName: "AddCollectionPermission",
			Handler:    _ContentService_AddCollectionPermission_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _ContentService_GetMetadata_Handler,
		},
		{
			MethodName: "AddMetadata",
			Handler:    _ContentService_AddMetadata_Handler,
		},
		{
			MethodName: "SetMetadataUploaded",
			Handler:    _ContentService_SetMetadataUploaded_Handler,
		},
		{
			MethodName: "GetMetadataPermissions",
			Handler:    _ContentService_GetMetadataPermissions_Handler,
		},
		{
			MethodName: "AddMetadataPermissions",
			Handler:    _ContentService_AddMetadataPermissions_Handler,
		},
		{
			MethodName: "AddMetadataPermission",
			Handler:    _ContentService_AddMetadataPermission_Handler,
		},
		{
			MethodName: "SetMetadataStatus",
			Handler:    _ContentService_SetMetadataStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content/content.proto",
}

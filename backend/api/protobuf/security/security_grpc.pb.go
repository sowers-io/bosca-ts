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
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: security/security.proto

package profiles

import (
	protobuf "bosca.io/api/protobuf"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SecurityService_GetGroups_FullMethodName = "/bosca.security.SecurityService/GetGroups"
)

// SecurityServiceClient is the client API for SecurityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityServiceClient interface {
	GetGroups(ctx context.Context, in *protobuf.Empty, opts ...grpc.CallOption) (*GetGroupsResponse, error)
}

type securityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityServiceClient(cc grpc.ClientConnInterface) SecurityServiceClient {
	return &securityServiceClient{cc}
}

func (c *securityServiceClient) GetGroups(ctx context.Context, in *protobuf.Empty, opts ...grpc.CallOption) (*GetGroupsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGroupsResponse)
	err := c.cc.Invoke(ctx, SecurityService_GetGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityServiceServer is the server API for SecurityService service.
// All implementations must embed UnimplementedSecurityServiceServer
// for forward compatibility
type SecurityServiceServer interface {
	GetGroups(context.Context, *protobuf.Empty) (*GetGroupsResponse, error)
	mustEmbedUnimplementedSecurityServiceServer()
}

// UnimplementedSecurityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecurityServiceServer struct {
}

func (UnimplementedSecurityServiceServer) GetGroups(context.Context, *protobuf.Empty) (*GetGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedSecurityServiceServer) mustEmbedUnimplementedSecurityServiceServer() {}

// UnsafeSecurityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityServiceServer will
// result in compilation errors.
type UnsafeSecurityServiceServer interface {
	mustEmbedUnimplementedSecurityServiceServer()
}

func RegisterSecurityServiceServer(s grpc.ServiceRegistrar, srv SecurityServiceServer) {
	s.RegisterService(&SecurityService_ServiceDesc, srv)
}

func _SecurityService_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityService_GetGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).GetGroups(ctx, req.(*protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SecurityService_ServiceDesc is the grpc.ServiceDesc for SecurityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecurityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bosca.security.SecurityService",
	HandlerType: (*SecurityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroups",
			Handler:    _SecurityService_GetGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "security/security.proto",
}

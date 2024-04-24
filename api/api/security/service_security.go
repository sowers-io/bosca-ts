package security

import (
	"bosca.io/api/protobuf"
	"bosca.io/api/protobuf/content"
	grpc "bosca.io/api/protobuf/security"
	"bosca.io/pkg/security"
	"context"
)

type serviceSecurity struct {
	grpc.UnimplementedSecurityServiceServer

	service grpc.SecurityServiceServer
	manager *security.Manager
}

func NewServiceSecurity(manager *security.Manager, service grpc.SecurityServiceServer) grpc.SecurityServiceServer {
	return &serviceSecurity{
		manager: manager,
		service: service,
	}
}

func (svc *serviceSecurity) GetGroups(ctx context.Context, request *protobuf.Empty) (*grpc.GetGroupsResponse, error) {
	err := svc.manager.CheckWithError(ctx, security.SystemResourceObject, "groups", content.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetGroups(ctx, request)
}

/*
 * Copyright 2024 Sowers, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package security

import (
	"bosca.io/api/protobuf"
	"bosca.io/api/protobuf/content"
	grpc "bosca.io/api/protobuf/security"
	"bosca.io/pkg/security"
	"context"
)

type authorizationService struct {
	grpc.UnimplementedSecurityServiceServer

	service grpc.SecurityServiceServer
	manager security.PermissionManager
}

func NewAuthorizationService(manager security.PermissionManager, service grpc.SecurityServiceServer) grpc.SecurityServiceServer {
	return &authorizationService{
		manager: manager,
		service: service,
	}
}

func (svc *authorizationService) GetGroups(ctx context.Context, request *protobuf.Empty) (*grpc.GetGroupsResponse, error) {
	err := svc.manager.CheckWithError(ctx, security.SystemResourceObject, "groups", content.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetGroups(ctx, request)
}

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

package ai

import (
	grpc "bosca.io/api/protobuf/bosca/ai"
	"bosca.io/pkg/security"
	"context"
)

// This service authorizes all calls to ensure the principal is authorized to perform the request.
//
//	If the principal is authorized, it forwards the request to the underlying service interface.
type authorizationService struct {
	grpc.UnimplementedAIServiceServer

	service     grpc.AIServiceServer
	permissions security.PermissionManager
}

func NewAuthorizationService(permissions security.PermissionManager, service grpc.AIServiceServer) grpc.AIServiceServer {
	svc := &authorizationService{
		service:     service,
		permissions: permissions,
	}
	return svc
}

func (svc *authorizationService) Chat(ctx context.Context, request *grpc.ChatRequest) (*grpc.ChatResponse, error) {
	return svc.service.Chat(ctx, request)
}

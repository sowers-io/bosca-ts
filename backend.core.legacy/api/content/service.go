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

package content

import (
	grpc "bosca.io/api/protobuf/bosca/content"
	grpcWorkflow "bosca.io/api/protobuf/bosca/workflow"
	"bosca.io/pkg/objectstore"
	"bosca.io/pkg/security"
)

type service struct {
	grpc.UnimplementedContentServiceServer

	ds          *DataStore
	objectStore objectstore.ObjectStore

	serviceAccountId    string
	serviceAccountToken string
	permissions         security.PermissionManager

	workflowClient grpcWorkflow.WorkflowServiceClient
}

const RootCollectionId = "00000000-0000-0000-0000-000000000000"

func NewService(dataStore *DataStore, serviceAccountId, serviceAccountToken string, objectStore objectstore.ObjectStore, permissions security.PermissionManager, workflowClient grpcWorkflow.WorkflowServiceClient) grpc.ContentServiceServer {
	go initializeService(dataStore)
	return &service{
		ds:                  dataStore,
		objectStore:         objectStore,
		serviceAccountId:    serviceAccountId,
		serviceAccountToken: serviceAccountToken,
		permissions:         permissions,
		workflowClient:      workflowClient,
	}
}

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
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/objectstore"
	"bosca.io/pkg/security"
	"go.temporal.io/sdk/client"
)

type service struct {
	grpc.UnimplementedContentServiceServer

	ds          *DataStore
	objectStore objectstore.ObjectStore

	serviceAccountId string
	permissions      security.PermissionManager
	temporalClient   client.Client
}

const RootCollectionId = "00000000-0000-0000-0000-000000000000"

func NewService(cfg *configuration.ServerConfiguration, dataStore *DataStore, serviceAccountId string, objectStore objectstore.ObjectStore, permissions security.PermissionManager, temporalClient client.Client) grpc.ContentServiceServer {
	go initializeService(cfg, dataStore)
	return &service{
		ds:               dataStore,
		objectStore:      objectStore,
		serviceAccountId: serviceAccountId,
		permissions:      permissions,
		temporalClient:   temporalClient,
	}
}

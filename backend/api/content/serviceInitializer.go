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
	grpc "bosca.io/api/protobuf/content"
	"bosca.io/pkg/security"
	"context"
	"log"
)

func initializeService(permissions security.PermissionManager, dataSource *DataStore) {
	ctx := context.Background()
	if added, err := dataSource.AddRootCollection(ctx); added {
		if err != nil {
			log.Fatalf("error initializing root collection: %v", err)
		}
		err = permissions.CreateRelationship(ctx, grpc.PermissionObjectType_collection_type, &grpc.Permission{
			Id:          RootCollectionId,
			Subject:     security.AdministratorGroup,
			SubjectType: grpc.PermissionSubjectType_group,
			Relation:    grpc.PermissionRelation_owners,
		})
		if err != nil {
			_ = dataSource.DeleteCollection(ctx, RootCollectionId)
			log.Fatalf("error initializing root collection permission: %v", err)
		}
	}
}

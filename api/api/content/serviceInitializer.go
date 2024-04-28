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
		err = permissions.CreateRelationship(ctx, security.CollectionObject, &grpc.Permission{
			Id:       RootCollectionId,
			Subject:  security.AdministratorGroup,
			Group:    true,
			Relation: grpc.PermissionRelation_owners,
		})
		if err != nil {
			_ = dataSource.DeleteCollection(ctx, RootCollectionId)
			log.Fatalf("error initializing root collection permission: %v", err)
		}
	}
}

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

package workflow

import (
	"bosca.io/api/graphql/common"
	"bosca.io/api/protobuf/bosca"
	"bosca.io/api/protobuf/bosca/content"
	grpc "bosca.io/api/protobuf/bosca/workflow"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/objectstore"
	"bosca.io/pkg/security"
	"context"
	opts "google.golang.org/grpc"
	"log/slog"
	"time"
)

type service struct {
	grpc.UnimplementedWorkflowServiceServer

	ds          *DataStore
	objectStore objectstore.ObjectStore

	serviceAccountId    string
	serviceAccountToken string
	permissions         security.PermissionManager

	contentClient content.ContentServiceClient
}

func NewService(cfg *configuration.ServerConfiguration, dataStore *DataStore, serviceAccountId, serviceAccountToken string, objectStore objectstore.ObjectStore, permissions security.PermissionManager, contentClient content.ContentServiceClient) grpc.WorkflowServiceServer {
	go initializeService(cfg, dataStore)
	svc := &service{
		ds:                  dataStore,
		objectStore:         objectStore,
		serviceAccountId:    serviceAccountId,
		serviceAccountToken: serviceAccountToken,
		permissions:         permissions,
		contentClient:       contentClient,
	}
	go func() {
		ctx := context.Background()
		for {
			time.Sleep(5 * time.Second)
			queues, err := svc.ds.GetAllQueues(ctx)
			if err != nil {
				slog.ErrorContext(ctx, "failed to get queues", slog.Any("error", err))
			}
			// TODO: check if there are actually jobs in the queues
			err = svc.ds.NotifyJobAvailable(ctx, queues)
			if err != nil {
				slog.ErrorContext(ctx, "notifying there are jobs... even if there are not", slog.Any("error", err))
			}
		}
	}()
	return svc
}

func (svc *service) getMetadata(ctx context.Context, metadataId string) (*content.Metadata, error) {
	var md *content.Metadata
	var err error
	// permissions can be laggy when first created
	for tries := 0; tries < 10; tries++ {
		md, err = svc.contentClient.GetMetadata(ctx, &bosca.IdRequest{Id: metadataId}, opts.PerRPCCredsCallOption{Creds: &common.Authorization{
			HeaderValue: "Token " + svc.serviceAccountToken,
		}})
		if err != nil {
			if err.Error() == "rpc error: code = Unauthenticated desc = permission check failed" {
				time.Sleep(3 * time.Second)
			} else {
				slog.ErrorContext(ctx, "failed to get metadata", slog.String("id", metadataId), slog.Any("error", err))
				return nil, err
			}
		} else if md != nil {
			break
		}
	}
	return md, err
}

func (svc *service) findMetadata(ctx context.Context, attributes map[string]string) ([]*content.Metadata, error) {
	md, err := svc.contentClient.FindMetadata(ctx, &content.FindMetadataRequest{Attributes: attributes}, opts.PerRPCCredsCallOption{Creds: &common.Authorization{
		HeaderValue: "Token " + svc.serviceAccountToken,
	}})
	if err != nil {
		return nil, err
	}
	return md.Metadata, nil
}

func (svc *service) findCollection(ctx context.Context, attributes map[string]string) ([]*content.Collection, error) {
	md, err := svc.contentClient.FindCollection(ctx, &content.FindCollectionRequest{Attributes: attributes}, opts.PerRPCCredsCallOption{Creds: &common.Authorization{
		HeaderValue: "Token " + svc.serviceAccountToken,
	}})
	if err != nil {
		return nil, err
	}
	return md.Collections, nil
}

func (svc *service) getCollection(ctx context.Context, collectionId string) (*content.Collection, error) {
	var col *content.Collection
	var err error
	// permissions can be laggy when first created
	for tries := 0; tries < 10; tries++ {
		col, err = svc.contentClient.GetCollection(ctx, &bosca.IdRequest{Id: collectionId}, opts.PerRPCCredsCallOption{Creds: &common.Authorization{
			HeaderValue: "Token " + svc.serviceAccountToken,
		}})
		if err != nil {
			if err.Error() == "rpc error: code = Unauthenticated desc = permission check failed" {
				time.Sleep(3 * time.Second)
			} else {
				slog.ErrorContext(ctx, "failed to get collection", slog.String("id", collectionId), slog.Any("error", err))
				return nil, err
			}
		} else if col != nil {
			break
		}
	}
	return col, err
}

func (svc *service) setBeginMetadataWorkflowState(ctx context.Context, metadata *content.Metadata, toState *grpc.WorkflowState, status string) error {
	_, err := svc.contentClient.SetWorkflowState(ctx, &content.SetWorkflowStateRequest{
		MetadataId: metadata.Id,
		StateId:    toState.Id,
		Status:     status,
	}, opts.PerRPCCredsCallOption{Creds: &common.Authorization{
		HeaderValue: "Token " + svc.serviceAccountToken,
	}})
	return err
}

func (svc *service) setCompleteMetadataWorkflowState(ctx context.Context, metadata *content.Metadata, status string) error {
	_, err := svc.contentClient.SetWorkflowStateComplete(ctx, &content.SetWorkflowStateCompleteRequest{
		MetadataId: metadata.Id,
		Status:     status,
	}, opts.PerRPCCredsCallOption{Creds: &common.Authorization{
		HeaderValue: "Token " + svc.serviceAccountToken,
	}})
	return err
}

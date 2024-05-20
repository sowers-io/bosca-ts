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

package search

import (
	grpcRequests "bosca.io/api/protobuf"
	grpcContent "bosca.io/api/protobuf/content"
	grpc "bosca.io/api/protobuf/search"
	"bosca.io/pkg/search"
	"context"
)

type service struct {
	grpc.UnimplementedSearchServiceServer

	contentClient grpcContent.ContentServiceClient

	semanticClient search.SemanticClient
	standardClient search.StandardClient
}

func NewService(contentClient grpcContent.ContentServiceClient, semanticClient search.SemanticClient, standardClient search.StandardClient) grpc.SearchServiceServer {
	return &service{
		contentClient:  contentClient,
		semanticClient: semanticClient,
		standardClient: standardClient,
	}
}

func (s *service) Search(ctx context.Context, request *grpc.SearchRequest) (*grpc.SearchResponse, error) {
	if request.Limit == 0 {
		request.Limit = 25
	}
	results, err := s.semanticClient.Search(ctx, s.semanticClient.GetMetadataIndex(), &search.SemanticQuery{
		Query:  request.Query,
		Offset: request.Offset,
		Limit:  request.Limit,
	})
	if err != nil {
		return nil, err
	}
	metadata, err := s.contentClient.GetMetadatas(ctx, &grpcRequests.IdsRequest{Id: results.MetadataIds})
	if err != nil {
		return nil, err
	}
	return &grpc.SearchResponse{
		Metadata: metadata.Metadata,
	}, nil
}

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
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/vectorstores"
	"log/slog"
)

type service struct {
	grpc.UnimplementedAIServiceServer

	serviceAccountId string
	permissions      security.PermissionManager

	store vectorstores.VectorStore
	model llms.Model
}

func NewService(serviceAccountId string, permissions security.PermissionManager, store vectorstores.VectorStore, model llms.Model) grpc.AIServiceServer {
	return &service{
		serviceAccountId: serviceAccountId,
		permissions:      permissions,
		store:            store,
		model:            model,
	}
}

func (s *service) Chat(ctx context.Context, req *grpc.ChatRequest) (*grpc.ChatResponse, error) {
	slog.DebugContext(ctx, "chat request", slog.String("request", req.Query))
	response, err := chains.Run(
		ctx,
		chains.NewRetrievalQAFromLLM(
			s.model,
			vectorstores.ToRetriever(s.store, 10000, vectorstores.WithScoreThreshold(0.6)),
		),
		req.Query,
	)
	if err != nil {
		slog.ErrorContext(ctx, "failed chat request", slog.String("request", req.Query), slog.Any("error", err))
		return nil, err
	}
	slog.InfoContext(ctx, "chat request", slog.String("request", req.Query), slog.String("response", response), slog.Any("error", err))
	return &grpc.ChatResponse{
		Response: response,
	}, nil
}

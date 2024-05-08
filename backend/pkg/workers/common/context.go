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

package common

import (
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/search"
	"context"
	"go.temporal.io/sdk/workflow"
	"google.golang.org/grpc/metadata"
	"net/http"
)

const configurationKey = "configuration"
const httpKey = "http"
const contentServiceKey = "contentService"
const searchClientKey = "searchClient"

type contextPropagator struct {
	cfg            *configuration.WorkerConfiguration
	contentService content.ContentServiceClient
	httpClient     *http.Client
	searchClient   search.Client
}

func NewContextPropagator(cfg *configuration.WorkerConfiguration, httpClient *http.Client, contentService content.ContentServiceClient, searchClient search.Client) workflow.ContextPropagator {
	return &contextPropagator{
		cfg:            cfg,
		contentService: contentService,
		httpClient:     httpClient,
		searchClient:   searchClient,
	}
}

func (c *contextPropagator) Inject(context.Context, workflow.HeaderWriter) error {
	return nil
}

func (c *contextPropagator) Extract(ctx context.Context, _ workflow.HeaderReader) (context.Context, error) {
	ctx = context.WithValue(ctx, configurationKey, c.cfg)
	ctx = context.WithValue(ctx, httpKey, c.httpClient)
	ctx = context.WithValue(ctx, contentServiceKey, c.contentService)
	ctx = context.WithValue(ctx, searchClientKey, c.searchClient)
	return ctx, nil
}

func (c *contextPropagator) InjectFromWorkflow(workflow.Context, workflow.HeaderWriter) error {
	return nil
}

func (c *contextPropagator) ExtractToWorkflow(ctx workflow.Context, writer workflow.HeaderReader) (workflow.Context, error) {
	ctx = workflow.WithValue(ctx, configurationKey, c.cfg)
	ctx = workflow.WithValue(ctx, httpKey, c.httpClient)
	ctx = workflow.WithValue(ctx, contentServiceKey, c.contentService)
	ctx = workflow.WithValue(ctx, searchClientKey, c.searchClient)
	return ctx, nil
}

func GetServiceAuthorizedContext(context context.Context) context.Context {
	cfg := GetConfiguration(context)
	md, exists := metadata.FromOutgoingContext(context)
	if !exists {
		md = metadata.New(make(map[string]string))
	}
	md["x-service-authorization"] = []string{"Token " + cfg.Security.ServiceAccountToken}
	return metadata.NewOutgoingContext(context, md)
}

func GetConfiguration(context context.Context) *configuration.WorkerConfiguration {
	return context.Value(configurationKey).(*configuration.WorkerConfiguration)
}

func GetHttpClient(context context.Context) *http.Client {
	return context.Value(httpKey).(*http.Client)
}

func GetContentService(context context.Context) content.ContentServiceClient {
	return context.Value(contentServiceKey).(content.ContentServiceClient)
}

func GetSearchClient(context context.Context) search.Client {
	return context.Value(searchClientKey).(search.Client)
}

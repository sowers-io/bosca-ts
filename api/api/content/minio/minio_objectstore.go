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

package minio

import (
	"bosca.io/api/content"
	model "bosca.io/api/protobuf/content"
	"bosca.io/pkg/configuration"
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type store struct {
	client *minio.Client
	bucket string
}

func NewMinioObjectStore(cfg *configuration.ServerConfiguration) content.ObjectStore {
	client, err := minio.New(cfg.Storage.S3.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Storage.S3.AccessKeyID, cfg.Storage.S3.SecretAccessKey, ""),
	})
	if err != nil {
		log.Fatalf("failed to create minio store: %v", err)
	}
	return &store{
		client: client,
		bucket: cfg.Storage.S3.Bucket,
	}
}

func (m *store) CreateUploadUrl(ctx context.Context, id string, _ string, contentType string, _ map[string]string) (*model.SignedUrl, error) {
	urlParams := url.Values{}
	headers := http.Header{
		"Content-Type": []string{contentType},
	}
	u, err := m.client.PresignHeader(ctx, "PUT", m.bucket, id, 5*time.Minute, urlParams, headers)
	if err != nil {
		return nil, err
	}
	return &model.SignedUrl{
		Url: u.String(),
	}, nil
}

func (m *store) CreateDownloadUrl(ctx context.Context, id string) (*model.SignedUrl, error) {
	urlParams := url.Values{}
	u, err := m.client.PresignedGetObject(ctx, m.bucket, id, 5*time.Minute, urlParams)
	if err != nil {
		return nil, err
	}
	return &model.SignedUrl{
		Url: u.String(),
	}, nil
}

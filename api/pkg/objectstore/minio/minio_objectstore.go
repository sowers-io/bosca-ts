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
	model "bosca.io/api/protobuf/content"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/objectstore"
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type objectStore struct {
	client *minio.Client
	bucket string
}

func NewMinioObjectStore(cfg *configuration.StorageConfiguration) objectstore.ObjectStore {
	client, err := minio.New(cfg.S3.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.S3.AccessKeyID, cfg.S3.SecretAccessKey, ""),
	})
	if err != nil {
		log.Fatalf("failed to create minio store: %v", err)
	}
	return &objectStore{
		client: client,
		bucket: cfg.S3.Bucket,
	}
}

func (m *objectStore) CreateUploadUrl(ctx context.Context, id string, _ string, contentType string, _ map[string]string) (*model.SignedUrl, error) {
	urlParams := url.Values{}
	headers := http.Header{
		"Content-Type": []string{contentType},
	}
	u, err := m.client.PresignHeader(ctx, "PUT", m.bucket, id, 5*time.Minute, urlParams, headers)
	if err != nil {
		return nil, err
	}
	return &model.SignedUrl{
		Id:  id,
		Url: u.String(),
	}, nil
}

func (m *objectStore) CreateDownloadUrl(ctx context.Context, id string) (*model.SignedUrl, error) {
	urlParams := url.Values{}
	u, err := m.client.PresignedGetObject(ctx, m.bucket, id, 5*time.Minute, urlParams)
	if err != nil {
		return nil, err
	}
	return &model.SignedUrl{
		Id:  id,
		Url: u.String(),
	}, nil
}

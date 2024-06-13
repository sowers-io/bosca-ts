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
	model "bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/objectstore"
	"context"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
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

func (m *objectStore) CreateUploadUrl(ctx context.Context, id string, _ string, contentType string, contentLength int64, _ map[string]string) (*model.SignedUrl, error) {
	urlParams := url.Values{}
	contentLengthStr := strconv.FormatInt(contentLength, 10)
	headers := http.Header{
		"Content-Type":   []string{contentType},
		"Content-Length": []string{contentLengthStr},
	}
	u, err := m.client.PresignHeader(ctx, "PUT", m.bucket, id, 5*time.Minute, urlParams, headers)
	if err != nil {
		return nil, err
	}
	return &model.SignedUrl{
		Id:     id,
		Method: "PUT",
		Headers: []*model.SignedUrlHeader{
			{
				Name:  "Content-Type",
				Value: contentType,
			},
			{
				Name:  "Content-Length",
				Value: contentLengthStr,
			},
		},
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
		Id:     id,
		Method: "GET",
		Url:    u.String(),
	}, nil
}

func (m *objectStore) Delete(ctx context.Context, id string) error {
	for object := range m.client.ListObjects(ctx, m.bucket, minio.ListObjectsOptions{
		Prefix: id,
	}) {
		err := m.client.RemoveObject(ctx, m.bucket, object.Key, minio.RemoveObjectOptions{})
		if err != nil {
			slog.ErrorContext(ctx, "failed to delete file", slog.String("file", object.Key), slog.Any("error", err))
			return err
		}
	}
	return nil
}

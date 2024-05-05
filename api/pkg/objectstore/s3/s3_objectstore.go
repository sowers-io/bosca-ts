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

package s3

import (
	model "bosca.io/api/protobuf/content"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/objectstore"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"time"
)

type objectStore struct {
	client  *s3.Client
	presign *s3.PresignClient
	bucket  string
}

func NewS3ObjectStore(cfg *configuration.StorageConfiguration) objectstore.ObjectStore {
	s3config, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(cfg.S3.Region),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: cfg.S3.Endpoint,
			}, nil
		})),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.S3.AccessKeyID, cfg.S3.SecretAccessKey, "")),
	)
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}
	s3client := s3.NewFromConfig(s3config)
	return &objectStore{
		client:  s3client,
		presign: s3.NewPresignClient(s3client),
		bucket:  cfg.S3.Bucket,
	}
}

func (m *objectStore) toSignedUrl(id string, request *v4.PresignedHTTPRequest) *model.SignedUrl {
	headers := make([]*model.SignedUrlHeader, 0)
	for key, values := range request.SignedHeader {
		for _, value := range values {
			headers = append(headers, &model.SignedUrlHeader{
				Name:  key,
				Value: value,
			})
		}
	}
	return &model.SignedUrl{
		Id:      id,
		Url:     request.URL,
		Method:  request.Method,
		Headers: headers,
	}
}

func (m *objectStore) CreateUploadUrl(ctx context.Context, id string, _ string, contentType string, _ map[string]string) (*model.SignedUrl, error) {
	request, err := m.presign.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(m.bucket),
		Key:         aws.String(id),
		ContentType: &contentType,
	}, func(opts *s3.PresignOptions) {
		opts.Expires = 5 * time.Minute
	})
	if err != nil {
		return nil, err
	}
	return m.toSignedUrl(id, request), nil
}

func (m *objectStore) CreateDownloadUrl(ctx context.Context, id string) (*model.SignedUrl, error) {
	request, err := m.presign.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(m.bucket),
		Key:    aws.String(id),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = 5 * time.Minute
	})
	if err != nil {
		return nil, err
	}
	return m.toSignedUrl(id, request), nil
}

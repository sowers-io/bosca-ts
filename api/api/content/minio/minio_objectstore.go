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
	client, err := minio.New(cfg.Storage.Minio.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.Storage.Minio.AccessKeyID, cfg.Storage.Minio.SecretAccessKey, ""),
	})
	if err != nil {
		log.Fatalf("failed to create minio store: %v", err)
	}
	return &store{
		client: client,
		bucket: cfg.Storage.Minio.Bucket,
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

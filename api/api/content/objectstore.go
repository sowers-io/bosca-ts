package content

import (
	"bosca.io/api/protobuf/content"
	"context"
)

type ObjectStore interface {
	CreateUploadUrl(ctx context.Context, id string, name string, contentType string, attributes map[string]string) (*content.SignedUrl, error)

	CreateDownloadUrl(ctx context.Context, id string) (*content.SignedUrl, error)
}

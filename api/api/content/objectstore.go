package content

import "bosca.io/api/protobuf/content"

type ObjectStore interface {
	CreateUploadUrl(id string, name string, contentType string, attributes map[string]string) (content.SignedUrl, error)

	CreateDownloadUrl(id string) (content.SignedUrl, error)
}

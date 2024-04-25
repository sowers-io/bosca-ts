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

package content

import (
	"bosca.io/api/protobuf/content"
	"context"
)

type ObjectStore interface {
	CreateUploadUrl(ctx context.Context, id string, name string, contentType string, attributes map[string]string) (*content.SignedUrl, error)

	CreateDownloadUrl(ctx context.Context, id string) (*content.SignedUrl, error)
}

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

package factory

import (
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/objectstore"
	"bosca.io/pkg/objectstore/minio"
	"bosca.io/pkg/objectstore/s3"
	"fmt"
)

func NewObjectStore(cfg *configuration.StorageConfiguration) (objectstore.ObjectStore, error) {
	var os objectstore.ObjectStore
	switch cfg.Type {
	case configuration.StorageTypeMinio:
		os = minio.NewMinioObjectStore(cfg)
		break
	case configuration.StorageTypeS3:
		os = s3.NewS3ObjectStore(cfg)
		break
	default:
		return nil, fmt.Errorf("unknown storage type: %s", cfg.Type)
	}
	return os, nil
}

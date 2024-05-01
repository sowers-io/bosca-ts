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

package configuration

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type ServerConfiguration struct {
	RestPort    int                    `envconfig:"REST_PORT"`
	GrpcPort    int                    `envconfig:"GRPC_PORT"`
	StorageType string                 `envconfig:"STORAGE_TYPE"`
	Database    *DatabaseConfiguration `ignored:"true"`
	Security    *SecurityConfiguration `ignored:"true"`
	Storage     *StorageConfiguration  `ignored:"true"`
}

type SecurityConfiguration struct {
	SessionEndpoint        string `envconfig:"SESSION_ENDPOINT" required:"true"`
	SessionEndpointType    string `envconfig:"SESSION_ENDPOINT_TYPE" default:"ory"`
	PermissionsEndPoint    string `envconfig:"PERMISSIONS_ENDPOINT" default:"localhost:50051"`
	PermissionsSharedToken string `envconfig:"PERMISSIONS_SHARED_TOKEN" required:"true"`
}

type DatabaseConfiguration struct {
	ConnectionString string `envconfig:"CONNECTION_STRING" required:"true"`
}

const StorageTypeMinio = "minio"
const StorageTypeS3 = "s3"

type StorageConfiguration struct {
	S3 *S3Configuration
}

type S3Configuration struct {
	Endpoint        string `envconfig:"ENDPOINT"`
	Region          string `envconfig:"REGION"`
	Bucket          string `envconfig:"BUCKET" default:"bosca"`
	AccessKeyID     string `envconfig:"ACCESS_KEY_ID"`
	SecretAccessKey string `envconfig:"SECRET_ACCESS_KEY"`
}

func getBaseConfiguration(defaultRestPort, defaultGrpcPort int) *ServerConfiguration {
	configuration := &ServerConfiguration{}
	err := envconfig.Process("bosca", configuration)
	if err != nil {
		log.Fatalf("failed to process base configuration: %v", err)
	}
	if configuration.RestPort == 0 {
		configuration.RestPort = defaultRestPort
	}
	if configuration.GrpcPort == 0 {
		configuration.GrpcPort = defaultGrpcPort
	}
	return configuration
}

func getDatabaseConfiguration(databasePrefix string) *DatabaseConfiguration {
	database := &DatabaseConfiguration{}
	err := envconfig.Process("bosca_"+databasePrefix, database)
	if err != nil {
		log.Fatalf("failed to process database configuration: %v", err)
	}
	return database
}

func getSecurityConfiguration() *SecurityConfiguration {
	permissions := &SecurityConfiguration{}
	err := envconfig.Process("bosca", permissions)
	if err != nil {
		log.Fatalf("failed to process security configuration: %v", err)
	}
	return permissions
}

func getStorageConfiguration(databasePrefix string, storageType string) *StorageConfiguration {
	if databasePrefix == "content" {
		switch storageType {
		case StorageTypeMinio:
			fallthrough
		case StorageTypeS3:
			cfg := &StorageConfiguration{
				S3: &S3Configuration{},
			}
			err := envconfig.Process("bosca_s3", cfg.S3)
			if err != nil {
				log.Fatalf("failed to process storage configuration: %v", err)
			}
			return cfg
		default:
			panic(errors.New("unknown storage type: " + storageType))
		}
	}
	return nil
}

func NewServerConfiguration(databasePrefix string, defaultRestPort, defaultGrpcPort int) *ServerConfiguration {
	configuration := getBaseConfiguration(defaultRestPort, defaultGrpcPort)
	configuration.Database = getDatabaseConfiguration(databasePrefix)
	configuration.Storage = getStorageConfiguration(databasePrefix, configuration.StorageType)
	configuration.Security = getSecurityConfiguration()
	return configuration
}

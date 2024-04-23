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
	RestPort                int                       `envconfig:"REST_PORT"`
	GrpcPort                int                       `envconfig:"GRPC_PORT"`
	OathKeeperConfiguration string                    `envconfig:"OAUTH_KEEPER_CONFIGURATION" default:"conf/oathkeeper.yaml"`
	StorageType             string                    `envconfig:"STORAGE_TYPE"`
	Database                *DatabaseConfiguration    `ignored:"true"`
	Permissions             *PermissionsConfiguration `ignored:"true"`
	Storage                 *StorageConfiguration     `ignored:"true"`
}

type PermissionsConfiguration struct {
	ReadEndPoint  string `envconfig:"KETO_READ_ENDPOINT" default:"http://localhost:4466"`
	WriteEndPoint string `envconfig:"KETO_WRITE_ENDPOINT" default:"http://localhost:4465"`
}

type DatabaseConfiguration struct {
	ConnectionString string `envconfig:"CONNECTION_STRING" required:"true"`
}

const StorageTypeMinio = "minio"

type StorageConfiguration struct {
	Minio *MinioConfiguration
}

type MinioConfiguration struct {
	Endpoint        string `envconfig:"ENDPOINT"`
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

func getPermissionsConfiguration() *PermissionsConfiguration {
	permissions := &PermissionsConfiguration{}
	err := envconfig.Process("bosca", permissions)
	if err != nil {
		log.Fatalf("failed to process permissions configuration: %v", err)
	}
	return permissions
}

func getStorageConfiguration(databasePrefix string, storageType string) *StorageConfiguration {
	if databasePrefix == "content" {
		switch storageType {
		case StorageTypeMinio:
			cfg := &StorageConfiguration{
				Minio: &MinioConfiguration{},
			}
			err := envconfig.Process("bosca_minio", cfg.Minio)
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
	configuration.Permissions = getPermissionsConfiguration()
	return configuration
}

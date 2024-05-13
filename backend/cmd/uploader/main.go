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

package main

import (
	"bosca.io/api/graphql/common"
	"bosca.io/api/protobuf"
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/clients"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/security"
	securityFactory "bosca.io/pkg/security/factory"
	"bosca.io/pkg/security/identity"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	opts "google.golang.org/grpc"
	"log"
	"net/http"
	"regexp"

	tusd "github.com/tus/tusd/v2/pkg/handler"
	"github.com/tus/tusd/v2/pkg/s3store"
)

func verify(cfg *configuration.ServerConfiguration, contentClient content.ContentServiceClient, subjectFinder security.SubjectFinder, hook tusd.HookEvent) error {
	var subjectId string
	var subjectType string
	var err error
	cookies := hook.HTTPRequest.Header["Cookie"]
	if cookies != nil && len(cookies) > 0 {
		subjectId, subjectType, err = subjectFinder.FindSubjectId(context.Background(), true, cookies[0])
		if err != nil {
			authorization := hook.HTTPRequest.Header["Authorization"]
			if authorization != nil && len(authorization) > 0 {
				subjectId, subjectType, err = subjectFinder.FindSubjectId(context.Background(), false, authorization[0])
				if err != nil {
					return err
				}
			} else {
				return errors.New("missing authorization header")
			}
		}
	} else {
		return errors.New("missing authorization header")
	}

	var subjectPermissionType content.PermissionSubjectType
	if subjectType == identity.SubjectTypeServiceAccount {
		subjectPermissionType = content.PermissionSubjectType_service_account
	} else {
		subjectPermissionType = content.PermissionSubjectType_user
	}

	collection := hook.Upload.MetaData["collection"]
	if collection == "" {
		collection = "00000000-0000-0000-0000-000000000000"
	}

	permission := &content.PermissionCheckRequest{
		Object:      collection,
		ObjectType:  content.PermissionObjectType_collection_type,
		Subject:     subjectId,
		SubjectType: subjectPermissionType,
		Action:      content.PermissionAction_edit,
	}

	result, err := contentClient.CheckPermission(context.Background(), permission, opts.PerRPCCredsCallOption{Creds: &common.Authorization{
		HeaderValue: "Token " + cfg.Security.ServiceAccountToken,
	}})

	if err != nil || result == nil || !result.Allowed {
		return errors.New("unauthorized")
	}
	return nil
}

func main() {
	cfg := configuration.NewServerConfiguration("", 8099, 0)

	contentConnection, err := clients.NewClientConnection(cfg.ClientEndPoints.ContentApiAddress)
	if err != nil {
		log.Fatalf("failed to create content client: %v", err)
	}
	defer contentConnection.Close()

	contentClient := content.NewContentServiceClient(contentConnection)

	s3config, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(cfg.Storage.S3.Region),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:               cfg.Storage.S3.Endpoint,
				HostnameImmutable: cfg.Storage.Type == configuration.StorageTypeMinio,
			}, nil
		})),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.Storage.S3.AccessKeyID, cfg.Storage.S3.SecretAccessKey, "")),
	)

	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}

	s3client := s3.NewFromConfig(s3config)

	store := s3store.New(cfg.Storage.S3.Bucket, s3client)

	composer := tusd.NewStoreComposer()
	store.UseIn(composer)

	cors := &tusd.CorsConfig{
		Disable:          false,
		AllowOrigin:      regexp.MustCompile(".*"),
		AllowCredentials: true,
		AllowMethods:     "POST, HEAD, PATCH, OPTIONS, GET, DELETE",
		AllowHeaders:     "Cookie, Origin, X-Requested-With, X-Request-ID, X-HTTP-Method-Override, Content-Type, Upload-Length, Upload-Offset, Tus-Resumable, Upload-Metadata, Upload-Defer-Length, Upload-Concat, Upload-Incomplete, Upload-Complete, Upload-Draft-Interop-Version",
		MaxAge:           "86400",
		ExposeHeaders:    "Cookie, Upload-Offset, Location, Upload-Length, Tus-Version, Tus-Resumable, Tus-Max-Size, Tus-Extension, Upload-Metadata, Upload-Defer-Length, Upload-Concat, Upload-Incomplete, Upload-Complete, Upload-Draft-Interop-Version",
	}

	sessionInterceptor := securityFactory.NewSessionInterceptor(cfg.Security.SessionEndpointType)
	subjectFinder := security.NewSubjectFinder(cfg.Security.SessionEndpoint, cfg.Security.ServiceAccountId, cfg.Security.ServiceAccountToken, sessionInterceptor)

	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:              "/uploads/",
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
		DisableDownload:       true,
		Cors:                  cors,
		PreUploadCreateCallback: func(hook tusd.HookEvent) (response tusd.HTTPResponse, changes tusd.FileInfoChanges, error error) {
			if err := verify(cfg, contentClient, subjectFinder, hook); err != nil {
				response.StatusCode = 401
				error = err
				return
			}
			response.StatusCode = 200
			error = nil
			return
		},
	})

	if err != nil {
		log.Fatalf("unable to create handler: %s", err)
	}

	go func() {
		for {
			event := <-handler.CompleteUploads

			if err := verify(cfg, contentClient, subjectFinder, event); err != nil {
				log.Printf("cannot complete upload: %v", err)
				return
			}

			collection := event.Upload.MetaData["collection"]
			if collection == "" {
				collection = "00000000-0000-0000-0000-000000000000"
			}

			metadata := &content.Metadata{
				Name:          event.Upload.MetaData["name"],
				ContentType:   event.Upload.MetaData["filetype"],
				ContentLength: event.Upload.Size,
				Source:        &event.Upload.ID,
			}

			id, err := contentClient.AddMetadata(context.Background(), &content.AddMetadataRequest{
				Collection: collection,
				Metadata:   metadata,
			}, opts.PerRPCCredsCallOption{Creds: &common.Authorization{
				HeaderValue: "Token " + cfg.Security.ServiceAccountToken,
			}})

			if err != nil {
				log.Printf("ERROR: unable to set metadata uploaded: %v", err)
				return
			}

			_, err = contentClient.SetMetadataUploaded(context.Background(), &protobuf.IdRequest{
				Id: id.Id,
			}, opts.PerRPCCredsCallOption{Creds: &common.Authorization{HeaderValue: "Token " + cfg.Security.ServiceAccountToken}})

			if err != nil {
				log.Printf("ERROR: unable to set metadata uploaded: %v", err)
			}
		}
	}()

	http.Handle("/uploads/", http.StripPrefix("/uploads/", handler))
	http.Handle("/uploads", http.StripPrefix("/uploads", handler))
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.RestPort), nil)
	if err != nil {
		log.Fatalf("unable to listen: %s", err)
	}
}

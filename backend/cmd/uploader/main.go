package main

import (
	"bosca.io/api/graphql/common"
	"bosca.io/api/protobuf"
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/clients"
	"bosca.io/pkg/configuration"
	"context"
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

	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:              "/uploads/",
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
		DisableDownload:       true,
		Cors:                  cors,
		PreUploadCreateCallback: func(hook tusd.HookEvent) (response tusd.HTTPResponse, changes tusd.FileInfoChanges, error error) {
			if hook.Upload.ID == "" {
				var id *protobuf.IdResponse
				collection := hook.Upload.MetaData["collection"]
				if collection == "" {
					collection = "00000000-0000-0000-0000-000000000000"
				}
				metadata := &content.Metadata{
					Name:          hook.Upload.MetaData["name"],
					ContentType:   hook.Upload.MetaData["filetype"],
					ContentLength: hook.Upload.Size,
				}
				id, error = contentClient.AddMetadata(context.Background(), &content.AddMetadataRequest{
					Collection: collection,
					Metadata:   metadata,
				}, opts.PerRPCCredsCallOption{Creds: &common.Authorization{HeaderValue: "Cookie " + hook.HTTPRequest.Header.Get("Cookie")}})
				if error != nil {
					response.StatusCode = 500
					error = fmt.Errorf("unable to add metadata: %s", error)
					return
				}
				changes.ID = id.Id
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
			_, err := contentClient.SetMetadataUploaded(context.Background(), &protobuf.IdRequest{
				Id: event.Upload.Storage["Key"],
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

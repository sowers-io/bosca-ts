package security

import (
	"bosca.io/pkg/configuration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"github.com/authzed/authzed-go/v1"
	"github.com/authzed/grpcutil"
)

func NewSpiceDBClient(cfg *configuration.ServerConfiguration) *authzed.Client {
	client, err := authzed.NewClient(
		cfg.Permissions.EndPoint,
		grpcutil.WithInsecureBearerToken(cfg.Permissions.SharedToken),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("unable to initialize client: %s", err)
	}
	return client
}

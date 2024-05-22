package cli

import (
	"bosca.io/cmd/cli/commands/login"
	"context"
	"google.golang.org/grpc/metadata"
)

func GetAuthenticatedContext(ctx context.Context) (context.Context, error) {
	session, err := login.GetSession()
	if err != nil {
		return nil, err
	}
	return metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", "Bearer "+session)), nil
}

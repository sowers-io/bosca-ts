package identity

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func GetUserId(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "failed to get grpc metadata")
	}
	userID := md.Get("X-User")
	if len(userID) == 0 {
		return "", status.Error(codes.Unauthenticated, "user is missing in metadata")
	}
	return userID[0], nil
}

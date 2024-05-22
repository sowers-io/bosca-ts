package cli

import (
	grpc "bosca.io/api/protobuf/content"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/clients"
	"github.com/spf13/cobra"
)

func NewContentClient(cmd *cobra.Command) (grpc.ContentServiceClient, error) {
	connection, err := clients.NewClientConnection(cmd.Flag(flags.EndpointFlag).Value.String())
	if err != nil {
		return nil, err
	}
	return grpc.NewContentServiceClient(connection), nil
}

package cli

import (
	grpc "bosca.io/api/protobuf/search"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/clients"
	"github.com/spf13/cobra"
)

func NewSearchClient(cmd *cobra.Command) (grpc.SearchServiceClient, error) {
	connection, err := clients.NewClientConnection(cmd.Flag(flags.EndpointFlag).Value.String())
	if err != nil {
		return nil, err
	}
	return grpc.NewSearchServiceClient(connection), nil
}

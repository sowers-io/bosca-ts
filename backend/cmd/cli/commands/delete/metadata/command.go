package collection

import (
	grpcRequests "bosca.io/api/protobuf"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"context"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "metadata [id]",
	Short: "Delete metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := cli.NewContentClient(cmd)
		if err != nil {
			return err
		}

		ctx, err := cli.GetAuthenticatedContext(context.Background())
		if err != nil {
			return err
		}

		_, err = client.DeleteMetadata(ctx, &grpcRequests.IdRequest{
			Id: args[0],
		})

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	Command.Flags().String(flags.EndpointFlag, "localhost:5013", "The endpoint to use.")
}

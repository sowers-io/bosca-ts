package collection

import (
	grpc "bosca.io/api/protobuf/content"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"context"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "collection [name]",
	Short: "Create a collection",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := cli.NewContentClient(cmd)
		if err != nil {
			return err
		}

		ctx, err := cli.GetAuthenticatedContext(context.Background())
		if err != nil {
			return err
		}

		parent := cmd.Flag(flags.ParentFlag)
		response, err := client.AddCollection(ctx, &grpc.AddCollectionRequest{
			Parent: parent.Value.String(),
			Collection: &grpc.Collection{
				Name: args[0],
			},
		})

		if err != nil {
			return err
		}

		cmd.Println(response.Id)

		return nil
	},
}

func init() {
	Command.Flags().String(flags.EndpointFlag, "localhost:5013", "The endpoint to use.")
}

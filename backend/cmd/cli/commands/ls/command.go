package ls

import (
	grpcRequests "bosca.io/api/protobuf"
	"bosca.io/api/protobuf/content"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"context"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "ls",
	Short: "List all the metadata in the current collection.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		client, err := cli.NewContentClient(cmd)
		if err != nil {
			return err
		}

		ctx, err := cli.GetAuthenticatedContext(context.Background())
		if err != nil {
			return err
		}

		var items *content.CollectionItems

		if len(args) == 0 {
			items, err = client.GetRootCollectionItems(ctx, &grpcRequests.Empty{})
		} else {
			items, err = client.GetCollectionItems(ctx, &grpcRequests.IdRequest{
				Id: args[0],
			})
		}

		if err != nil {
			return err
		}

		tbl := table.NewWriter()
		tbl.AppendHeader(table.Row{"ID", "Name", "Type", "Status", "State", "Language"})

		for _, item := range items.Items {
			collection := item.GetCollection()
			if collection != nil {
				tbl.AppendRow(table.Row{
					collection.Id,
					collection.Name,
					"collection: " + collection.Type.String(),
					"--",
					"--",
					"--",
				})
			}
			metadata := item.GetMetadata()
			if metadata != nil {
				tbl.AppendRow(table.Row{
					metadata.Id,
					metadata.Name,
					metadata.ContentType,
					metadata.Status.String(),
					metadata.WorkflowStateId,
					metadata.LanguageTag,
				})
			}
		}

		fmt.Printf("%s", tbl.Render())

		return nil
	},
}

func init() {
	Command.Flags().String(flags.EndpointFlag, "localhost:5013", "The endpoint to use.")
}

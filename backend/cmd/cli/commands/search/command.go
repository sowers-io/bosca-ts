package search

import (
	grpc "bosca.io/api/protobuf/search"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"context"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "search [query]",
	Short: "Search across all metadata",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		client, err := cli.NewSearchClient(cmd)
		if err != nil {
			return err
		}

		ctx, err := cli.GetAuthenticatedContext(context.Background())
		if err != nil {
			return err
		}

		cmd.Printf("Searching for %s\n", args[0])

		response, err := client.Search(ctx, &grpc.SearchRequest{
			Query: args[0],
		})

		if err != nil {
			return err
		}

		tbl := table.NewWriter()
		tbl.AppendHeader(table.Row{"ID", "Name", "Type", "Status", "State", "Language"})

		for _, metadata := range response.Metadata {
			tbl.AppendRow(table.Row{
				metadata.Id,
				metadata.Name,
				metadata.ContentType,
				metadata.Status.String(),
				metadata.WorkflowStateId,
				metadata.LanguageTag,
			})
		}

		fmt.Printf("%s", tbl.Render())

		return nil
	},
}

func init() {
	Command.Flags().String(flags.EndpointFlag, "localhost:5015", "The endpoint to use.")
}

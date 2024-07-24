/*
 * Copyright 2024 Sowers, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ls

import (
	grpcRequests "bosca.io/api/protobuf/bosca"
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"context"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "ls",
	Short: "List all the workflow in the current collection.",
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
		tbl.AppendHeader(table.Row{"ID", "Name", "Type", "State", "Language"})

		for _, item := range items.Items {
			collection := item.GetCollection()
			if collection != nil {
				tbl.AppendRow(table.Row{
					collection.Id,
					collection.Name,
					"collection: " + collection.Type.String(),
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
					metadata.WorkflowStateId,
					metadata.LanguageTag,
				})
			}
		}

		cmd.Printf("%s", tbl.Render())

		return nil
	},
}

func init() {
	Command.Flags().String(flags.EndpointFlag, "localhost:5013", "The endpoint to use.")
}

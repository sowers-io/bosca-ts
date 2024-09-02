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

package collection

import (
	"context"

	grpcRequests "bosca.io/api/protobuf/bosca"
	grpc "bosca.io/api/protobuf/bosca/content"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"github.com/spf13/cobra"
)

func deleteAll(ctx context.Context, id string, client grpc.ContentServiceClient, recursive bool) error {
	if recursive {
		items, err := client.GetCollectionItems(ctx, &grpcRequests.IdRequest{
			Id: id,
		})
		if err != nil {
			return err
		}
		for _, item := range items.Items {
			if item.GetCollection() != nil {
				err = deleteAll(ctx, item.GetCollection().Id, client, true)
				if err != nil {
					return err
				}
			}
			if item.GetMetadata() != nil {
				_, err = client.DeleteMetadata(ctx, &grpcRequests.IdRequest{
					Id: item.GetMetadata().Id,
				})
				if err != nil {
					return err
				}
			}
		}
	}
	_, err := client.DeleteCollection(ctx, &grpcRequests.IdRequest{
		Id: id,
	})
	if err != nil {
		return err
	}
	return nil
}

var Command = &cobra.Command{
	Use:   "collection [name]",
	Short: "Delete a collection",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := cli.NewContentClient(cmd)
		if err != nil {
			return err
		}
		ctx, err := cli.GetAuthenticatedContext(context.Background())
		if err != nil {
			return err
		}
		err = deleteAll(ctx, args[0], client, cmd.Flag(flags.RecursiveFlag).Value.String() == "true")
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	Command.Flags().String(flags.EndpointFlag, "localhost:7000", "The endpoint to use.")
	Command.Flags().Bool(flags.RecursiveFlag, false, "Recursive deletion")
}

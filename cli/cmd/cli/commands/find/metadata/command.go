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

package metadata

import (
	grpc "bosca.io/api/protobuf/bosca/content"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"context"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"strings"
)

var Command = &cobra.Command{
	Use:   "metadata",
	Short: "Find a metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := cli.NewContentClient(cmd)
		if err != nil {
			return err
		}

		ctx, err := cli.GetAuthenticatedContext(context.Background())
		if err != nil {
			return err
		}

		attributes := make(map[string]string)

		findArgs := cmd.Flag(flags.ArgsFlag).Value.String()

		for _, arg := range strings.Split(findArgs, ",") {
			kv := strings.Split(arg, "=")
			if len(kv) != 2 {
				continue
			}
			attributes[kv[0]] = kv[1]
		}

		cmd.Printf("Args: %v\n", attributes)

		response, err := client.FindMetadata(ctx, &grpc.FindMetadataRequest{
			Attributes: attributes,
		})

		if err != nil {
			return err
		}

		tbl := table.NewWriter()
		tbl.AppendHeader(table.Row{"ID", "Name", "Type", "Attributes"})

		for _, collection := range response.Metadata {
			tbl.AppendRow(table.Row{
				collection.Id,
				collection.Name,
				collection.ContentType,
				collection.Attributes,
			})
		}

		cmd.Printf("%s", tbl.Render())

		return nil
	},
}

func init() {
	Command.Flags().String(flags.EndpointFlag, "localhost:7000", "The endpoint to use.")
	Command.Flags().String(flags.ArgsFlag, "", "The args")
}

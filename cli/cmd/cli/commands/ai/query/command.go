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

package query

import (
	grpc "bosca.io/api/protobuf/bosca/ai"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"context"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "query [query]",
	Short: "Query the Specified Storage Engine using the model associated with the engine",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		client, err := cli.NewAIClient(cmd)
		if err != nil {
			return err
		}

		ctx, err := cli.GetAuthenticatedContext(context.Background())
		if err != nil {
			return err
		}

		cmd.Printf("Sending message: %s\n", args[0])

		response, err := client.QueryStorage(ctx, &grpc.QueryStorageRequest{
			Query:         args[0],
			StorageSystem: cmd.Flag(flags.StorageSystemFlag).Value.String(),
		})

		if err != nil {
			return err
		}

		cmd.Println(response)

		return nil
	},
}

func init() {
	Command.Flags().String(flags.StorageSystemFlag, "", "The storage system to use.")
	Command.Flags().String(flags.EndpointFlag, "localhost:5007", "The endpoint to use.")
}

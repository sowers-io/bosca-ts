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

package transition

import (
	grpc "bosca.io/api/protobuf/bosca/content"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"context"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "transition [metadata id] [state]",
	Short: "Transition workflow to a new state",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := cli.NewContentClient(cmd)
		if err != nil {
			return err
		}

		ctx, err := cli.GetAuthenticatedContext(context.Background())
		if err != nil {
			return err
		}

		_, err = client.BeginTransitionWorkflow(ctx, &grpc.TransitionWorkflowRequest{
			MetadataId: args[0],
			StateId:    args[1],
			Retry:      cmd.Flag(flags.RetryFlag).Value.String() == "true",
		})

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	Command.Flags().Bool(flags.RetryFlag, false, "Retry transition")
	Command.Flags().String(flags.EndpointFlag, "localhost:5013", "The endpoint to use.")
}

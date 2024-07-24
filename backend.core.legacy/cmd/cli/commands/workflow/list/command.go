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

package list

import (
	"bosca.io/api/protobuf/bosca"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"context"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "list",
	Short: "List workflows",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := cli.NewWorkflowClient(cmd)
		if err != nil {
			return err
		}

		ctx, err := cli.GetAuthenticatedContext(context.Background())
		if err != nil {
			return err
		}

		workflows, err := client.GetWorkflows(ctx, &bosca.Empty{})

		if err != nil {
			return err
		}

		tbl := table.NewWriter()
		tbl.AppendHeader(table.Row{"ID", "Name", "Description"})

		for _, workflow := range workflows.Workflows {
			tbl.AppendRow(table.Row{workflow.Id, workflow.Name, workflow.Description})
		}

		cmd.Printf("%s", tbl.Render())
		return nil
	},
}

func init() {
	Command.Flags().Bool(flags.RetryFlag, false, "Retry transition")
	Command.Flags().String(flags.EndpointFlag, "localhost:5011", "The endpoint to use.")
}

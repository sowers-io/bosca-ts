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
	"bosca.io/api/protobuf/bosca/workflow"
	"bosca.io/cmd/cli/commands/flags"
	"bosca.io/pkg/cli"
	"context"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"strings"
)

var Command = &cobra.Command{
	Use:   "metadata [workflow]",
	Short: "Execute metadata workflows",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := cli.NewWorkflowClient(cmd)
		if err != nil {
			return err
		}

		ctx, err := cli.GetAuthenticatedContext(context.Background())
		if err != nil {
			return err
		}

		findArgs := cmd.Flag(flags.ArgsFlag).Value.String()

		attributes := make(map[string]string)
		for _, arg := range strings.Split(findArgs, ",") {
			kv := strings.Split(arg, "=")
			if len(kv) != 2 {
				continue
			}
			attributes[kv[0]] = kv[1]
		}

		cmd.Printf("Args: %v\n", attributes)

		executions, err := client.FindAndExecuteWorkflow(ctx, &workflow.FindAndWorkflowExecutionRequest{
			WorkflowId:         args[0],
			MetadataAttributes: attributes,
			WaitForCompletion:  cmd.Flag(flags.WaitFlag).Value.String() == "true",
		})

		tbl := table.NewWriter()
		tbl.AppendHeader(table.Row{"ID", "Success", "Complete", "Error"})

		for _, execution := range executions.Responses {
			tbl.AppendRow(table.Row{execution.ExecutionId, execution.Success, execution.Complete, execution.Error})
		}

		cmd.Printf("%s", tbl.Render())

		return err
	},
}

func init() {
	Command.Flags().String(flags.ArgsFlag, "", "The args to use to find items")
	Command.Flags().Bool(flags.WaitFlag, false, "Wait for completion")
	Command.Flags().String(flags.EndpointFlag, "localhost:5011", "The endpoint to use.")
}

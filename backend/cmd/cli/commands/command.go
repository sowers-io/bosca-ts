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

package commands

import (
	"bosca.io/cmd/cli/commands/create"
	del "bosca.io/cmd/cli/commands/delete"
	"bosca.io/cmd/cli/commands/login"
	"bosca.io/cmd/cli/commands/ls"
	"bosca.io/cmd/cli/commands/search"
	"bosca.io/cmd/cli/commands/upload"
	"bosca.io/cmd/cli/commands/workflow"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var command = &cobra.Command{
	Use:   "bosca",
	Short: "Bosca is an AI powered platform for creating and managing content.",
	Long:  "Bosca is an AI powered platform for creating and managing content. See more at bosca.io",
}

func init() {
	command.AddCommand(ls.Command, login.Command, workflow.Command, create.Command, del.Command, upload.Command, search.Command)
}

func Execute() {
	if err := command.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package commands

import (
	"bosca.io/cmd/cli/commands/create"
	del "bosca.io/cmd/cli/commands/delete"
	"bosca.io/cmd/cli/commands/login"
	"bosca.io/cmd/cli/commands/ls"
	"bosca.io/cmd/cli/commands/search"
	"bosca.io/cmd/cli/commands/upload"
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
	command.AddCommand(ls.Command, login.Command, create.Command, del.Command, upload.Command, search.Command)
}

func Execute() {
	if err := command.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

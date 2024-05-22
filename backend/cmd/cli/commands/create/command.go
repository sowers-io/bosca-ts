package create

import (
	"bosca.io/cmd/cli/commands/create/collection"
	"bosca.io/cmd/cli/commands/flags"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "create",
	Short: "Create a resource",
}

func init() {
	Command.PersistentFlags().String(flags.ParentFlag, "", "The parent collection id")
	Command.AddCommand(collection.Command)
}

package delete

import (
	"bosca.io/cmd/cli/commands/delete/collection"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource",
}

func init() {
	Command.AddCommand(collection.Command)
}

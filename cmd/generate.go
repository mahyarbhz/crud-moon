package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [entity-name]",
	Short: "Generate CRUD structure",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entityName := args[0]
		fmt.Printf("Generating CRUD for entity: %s\n", entityName)
	},
}

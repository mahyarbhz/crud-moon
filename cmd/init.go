package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [module-name]", // Ensure there's no space or prefix here
	Short: "Initialize your base project",
	Long:  `Initialize your base project folder ([module-name]), use flag "--not-complex" for a simpler directory design.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]
		fmt.Printf("Creating the great %s project!\n", moduleName)
		err := os.Mkdir(moduleName, 0755)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)
		}
		err = os.WriteFile(moduleName+"/.env", []byte("Hello, Gophers!"), 0666)
		if err != nil {
			log.Fatal(err)
		}
	},
}

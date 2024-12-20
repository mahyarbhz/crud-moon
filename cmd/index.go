package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "crud-moon",
	Short: "A CLI tool for generating CRUD code",
	Long: `Welcome to CRUD Moon CLI tool!
CRUD Moon is a CLI tool for generating CRUD code based on the DDD architecture for Echo framework using GORM.
****
Start with "init own" for creating a new project with DDD design, or use "custom -h" command inside your app folder for defining your own folder structure.
****`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(generateCmd)
}

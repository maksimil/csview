package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open [flags] [file]",
	Short: "opens a file",
	Run:   runOpen,
}

func runOpen(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}

func init() {
	rootCmd.AddCommand(openCmd)

	openCmd.Args = cobra.ExactArgs(1)
}

package cmd

import (
	"csview/src"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open [flags] [file]",
	Short: "opens a file",
	Run:   src.RunOpen,
}

func init() {
	rootCmd.AddCommand(openCmd)

	openCmd.Args = cobra.ExactArgs(1)
}

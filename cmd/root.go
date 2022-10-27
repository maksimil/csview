package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "csview",
	Short: "Open and edit csv files vim-style",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
}

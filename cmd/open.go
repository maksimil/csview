package cmd

import (
	"csview/tui"
	"time"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open [flags] [file]",
	Short: "opens a file",
	Run:   runOpen,
}

func runOpen(cmd *cobra.Command, args []string) {
	canvas := tui.CreateCanvas()

	canvas.Batch(func(b *tui.Batch) {
		b.PutString(0, 0, "hi")
	})

	defer canvas.Batch(func(b *tui.Batch) { b.Clear() })

	time.Sleep(time.Second)

	canvas.Batch(func(b *tui.Batch) {
		b.Clear()
		b.PutString(1, 1, "hi1")
	})

	time.Sleep(time.Second)
}

func init() {
	rootCmd.AddCommand(openCmd)

	openCmd.Args = cobra.ExactArgs(1)
}

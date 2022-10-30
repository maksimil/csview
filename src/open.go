package src

import (
	"time"

	"github.com/spf13/cobra"
)

func RunOpen(cmd *cobra.Command, args []string) {
	canvas := CreateCanvas()

	canvas.Batch(func(b *Batch) {
		b.PutString(0, 0, "hi")
		b.Mv(0, 0)
	})

	defer canvas.Batch(func(b *Batch) { b.Clear() })

	time.Sleep(time.Second)

	canvas.Batch(func(b *Batch) {
		b.Clear()
		b.PutString(1, 1, "hi1")
		b.Mv(0, 0)
	})

	time.Sleep(time.Second)
}

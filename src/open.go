package src

import (
	"time"

	"github.com/spf13/cobra"
)

func RunOpen(cmd *cobra.Command, args []string) {
	draw_communication := RunDrawRoutine()
	draw_communication.StateUpdate <- StateUpdate

	{
		csv_source := "hi;hi;123\nhii;i;34;"
		draw_communication.DrawState.Mutex.Lock()
		draw_communication.DrawState.Document = CsvDocument{
			Path: "path",
			Data: ParseCsv(csv_source),
		}
		draw_communication.DrawState.Mutex.Unlock()
		draw_communication.StateUpdate <- StateUpdate
	}

	time.Sleep(time.Second)

	{
		csv_source := "hi;hi;1234\nhii;i;34;"
		draw_communication.DrawState.Mutex.Lock()
		draw_communication.DrawState.Document = CsvDocument{
			Path: "path",
			Data: ParseCsv(csv_source),
		}
		draw_communication.DrawState.Mutex.Unlock()
		draw_communication.StateUpdate <- StateUpdate
	}

	time.Sleep(time.Second)

	draw_communication.StateUpdate <- Exit
	<-draw_communication.TuiClosed
}

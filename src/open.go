package src

import (
	"time"

	"github.com/spf13/cobra"
)

func RunOpen(cmd *cobra.Command, args []string) {
	draw_communication := RunDrawRoutine()
	draw_communication.StateUpdate <- StateUpdate

	{
		csv_source := "Name;Age\nAnderson;1\nVV;10"
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
		csv_source := "Name;Age\nAnderson;1\nVV;12"
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

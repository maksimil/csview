package src

import (
	"time"

	"github.com/spf13/cobra"
)

func RunOpen(cmd *cobra.Command, args []string) {
	draw_communication := RunDrawRoutine()

	{
		csv_source := "Name;Age\nAnderson;1\nVV;10"
		draw_communication.UpdateState(func(draw_state *DrawState) {
			draw_state.Document = CsvDocument{
				Path: "path",
				Data: ParseCsv(csv_source),
			}

		})
	}

	time.Sleep(time.Second)

	{
		csv_source := "Name;Age\nAnderson;1\nVV;12"
		draw_communication.UpdateState(func(draw_state *DrawState) {
			draw_state.Document = CsvDocument{
				Path: "path",
				Data: ParseCsv(csv_source),
			}

		})
	}

	time.Sleep(time.Second)

	draw_communication.Close()
}

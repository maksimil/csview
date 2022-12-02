package src

import (
	"sync"

	"github.com/rs/zerolog/log"
)

type DrawState struct {
	Mutex    sync.Mutex
	Document CsvDocument
}

type UpdateMessage int8

const (
	StateUpdate UpdateMessage = iota
	Exit
)

type DrawRoutineCommunication struct {
	StateUpdate chan UpdateMessage
	TuiClosed   chan bool
	DrawState   *DrawState
}

func RunDrawRoutine() DrawRoutineCommunication {
	draw_state := DrawState{}
	communication := DrawRoutineCommunication{}

	communication.StateUpdate = make(chan UpdateMessage)
	communication.TuiClosed = make(chan bool)
	communication.DrawState = &draw_state

	go func() {
		canvas := CreateCanvas()
		log.Info().Msg("Created canvas")

		for message := range communication.StateUpdate {
			log.Info().Interface("message", message).Msg("Processing message")

			if message == Exit {
				log.Info().Msg("Exiting drawer")
				canvas.Batch(func(b *Batch) { b.Clear() })
				communication.TuiClosed <- true
				return
			}

			if message == StateUpdate {
				draw_state.Mutex.Lock()
				canvas.Batch(func(b *Batch) {
					DrawFunction(b, &draw_state)
				})
				draw_state.Mutex.Unlock()
			}
		}

	}()

	return communication
}

func DrawFunction(b *Batch, state_ptr *DrawState) {
	b.Clear()

	column_count := 0

	for _, csv_line := range state_ptr.Document.Data {
		if len(csv_line) > column_count {
			column_count = len(csv_line)
		}
	}

	accumulated_x := 0
	for column_index := 0; column_index < column_count; column_index++ {
		column_width := make(chan int)
		go func() {
			column_width_calculated := 1

			for _, csv_line := range state_ptr.Document.Data {
				if column_index < len(csv_line) && len(csv_line[column_index]) > column_width_calculated {
					column_width_calculated = len(csv_line[column_index])
				}
			}

			column_width <- column_width_calculated
		}()

		for line_index, csv_line := range state_ptr.Document.Data {
			if column_index < len(csv_line) {
				b.PutStringf(accumulated_x, line_index, csv_line[column_index])
			}
		}

		accumulated_x += <-column_width
	}
}

package src

import (
	"strings"
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

const (
	BOX_DRAWING_HORIZONTAL = "─"
	BOX_DRAWING_VERTICAL   = "│"
	BOX_DRAWING_CROSS      = "┼"
)

func DrawFunction(b *Batch, state_ptr *DrawState) {
	b.Clear()

	// drawing the table
	column_count := 0
	for _, csv_line := range state_ptr.Document.Data {
		if len(csv_line) > column_count {
			column_count = len(csv_line)
		}
	}

	accumulated_x := 0
	for column_index := 0; column_index < column_count+1; column_index++ {
		// calculating column width
		column_width := 0
		for _, csv_line := range state_ptr.Document.Data {
			if column_index < len(csv_line) && len(csv_line[column_index]) > column_width {
				column_width = len(csv_line[column_index])
			}
		}

		// drawing the table
		b.PutStringf(accumulated_x, 0,
			"%s%s", BOX_DRAWING_CROSS, strings.Repeat(BOX_DRAWING_HORIZONTAL, column_width))
		for line_index, csv_line := range state_ptr.Document.Data {
			s := ""
			if column_index < len(csv_line) {
				s = csv_line[column_index]
			}
			b.PutStringf(accumulated_x, 2*line_index+1, "%s%s", BOX_DRAWING_VERTICAL, s)

			b.PutStringf(accumulated_x, 2*line_index+2,
				"%s%s", BOX_DRAWING_CROSS, strings.Repeat(BOX_DRAWING_HORIZONTAL, column_width))
		}

		accumulated_x += column_width + 1
	}
}

package src

import (
	"sync"

	"github.com/rs/zerolog/log"
)

type DrawState struct {
	Document CsvDocument
}

type UpdateMessage int8

const (
	StateUpdate UpdateMessage = iota
	Exit
)

type DrawRoutineCommunication struct {
	StateUpdate    chan UpdateMessage
	TuiClosed      chan bool
	DrawState      *DrawState
	DrawStateMutex *sync.Mutex
}

func CreateCommunication() DrawRoutineCommunication {
	state := DrawState{}
	var mutex sync.Mutex

	return DrawRoutineCommunication{
		StateUpdate:    make(chan UpdateMessage),
		TuiClosed:      make(chan bool),
		DrawState:      &state,
		DrawStateMutex: &mutex,
	}
}

func (communication *DrawRoutineCommunication) UpdateState(updater func(*DrawState)) {
	communication.DrawStateMutex.Lock()
	updater(communication.DrawState)
	communication.DrawStateMutex.Unlock()
	communication.StateUpdate <- StateUpdate
}

func (communication *DrawRoutineCommunication) Close() {
	communication.StateUpdate <- Exit
	<-communication.TuiClosed
}

func RunDrawRoutine() DrawRoutineCommunication {
	communication := CreateCommunication()

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
				communication.DrawStateMutex.Lock()
				canvas.Batch(func(b *Batch) {
					DrawFunction(b, communication.DrawState)
				})
				communication.DrawStateMutex.Unlock()
			}
		}
	}()

	communication.StateUpdate <- StateUpdate

	return communication
}

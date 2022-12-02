package src

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/term"
)

type Canvas struct {
	Width  int
	Height int
	state  *term.State
}

func CreateCanvas() Canvas {
	width, height, err := term.GetSize(0)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to init terminal")
	}

	log.Info().Int("width", width).Int("height", height).Msg("Terminal size")

	state, err := term.MakeRaw(0)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to set terminal to raw")
	}

	canvas := Canvas{
		Width:  width,
		Height: height,
		state:  state,
	}

	canvas.Batch(func(b *Batch) {
		b.print(strings.Repeat("\n", b.Canvas.Height-1))
		b.Mv(0, 0)
	})

	return Canvas{
		Width:  width,
		Height: height,
	}
}

func (canvas *Canvas) Batch(drawer func(*Batch)) {
	batch := Batch{
		input:  "",
		Canvas: canvas,
	}

	drawer(&batch)

	fmt.Print(batch.input)
}

func (canvas *Canvas) Close() {
	canvas.Batch(func(b *Batch) { b.Clear() })
	term.Restore(0, canvas.state)
}

type Batch struct {
	input  string
	Canvas *Canvas
}

func (batch *Batch) print(s string) {
	batch.input += s
}

func (batch *Batch) printf(s string, a ...any) {
	batch.input += fmt.Sprintf(s, a...)
}

func (batch *Batch) Mv(x, y int) {
	xop := ""
	yop := ""

	if x > 0 {
		xop = fmt.Sprintf("\x1b[%vC", x)
	}

	if y > 0 {
		yop = fmt.Sprintf("\x1b[%vB", y)
	}

	batch.printf("\x1b[H%v%v", xop, yop)
}

func (batch *Batch) PutString(x, y int, s string) {
	batch.Mv(x, y)
	batch.print(s)
}

func (batch *Batch) PutStringf(x, y int, s string, args ...any) {
	batch.PutString(x, y, fmt.Sprintf(s, args...))
}

func (batch *Batch) Clear() {
	batch.Mv(0, 0)
	batch.print("\x1b[J")
}

package tui

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"golang.org/x/term"
)

type Canvas struct {
	Width  int
	Height int
}

func CreateCanvas() Canvas {
	width, height, err := term.GetSize(0)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to init terminal")
	}

	log.Info().Int("width", width).Int("height", height).Msg("Terminal size")

	canvas := Canvas{
		Width:  width,
		Height: height,
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

func (batch *Batch) Mv(x, y int32) {
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

func (batch *Batch) PutString(x, y int32, s string) {
	batch.Mv(x, y)
	batch.print(s)
}

func (batch *Batch) PutStringf(x, y int32, s string, args ...any) {
	batch.PutString(x, y, fmt.Sprintf(s, args...))
}

func (batch *Batch) Clear() {
	batch.Mv(0, 0)
	batch.print("\x1b[J")
}

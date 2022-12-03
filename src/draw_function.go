package src

import (
	"regexp"
	"strings"

	"github.com/gookit/color"
)

// const (
//   BOX_DRAWING_HORIZONTAL = "─"
//   BOX_DRAWING_VERTICAL   = "│"
//   BOX_DRAWING_CROSS      = "┼"
// )

const (
	BOX_DRAWING_HORIZONTAL = "━"
	BOX_DRAWING_VERTICAL   = "┃"
	BOX_DRAWING_CROSS      = "╋"
)

var (
	BORDER_COLOR = color.New(color.FgGray, color.BgBlack)
	NUMBER_COLOR = color.New(color.FgBlue, color.BgBlack)
)

func SHorizontalDivider(width int) string {
	return BORDER_COLOR.Sprintf(
		"%s%s", BOX_DRAWING_CROSS, strings.Repeat(BOX_DRAWING_HORIZONTAL, width))
}

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
		b.PutString(accumulated_x, 0, SHorizontalDivider(column_width))
		for line_index, csv_line := range state_ptr.Document.Data {
			s := ""
			if column_index < len(csv_line) {
				s = csv_line[column_index]
			}
			s = FormatCellContents(s, column_width)

			b.PutStringf(accumulated_x, 2*line_index+1, "%s%s", BORDER_COLOR.Sprint(BOX_DRAWING_VERTICAL), s)
			b.PutStringf(accumulated_x, 2*line_index+2, SHorizontalDivider(column_width))
		}

		accumulated_x += column_width + 1
	}
}

var (
	NUMBER_MATCH = regexp.MustCompile(`^[\d ]*[.,]*[\d ]*$`)
)

func FormatCellContents(cell_contents string, cell_width int) string {
	if NUMBER_MATCH.Match([]byte(cell_contents)) {
		return NUMBER_COLOR.Sprintf("%*s", cell_width, cell_contents)
	}

	return cell_contents
}

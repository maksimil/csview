package src

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestParseCsv(t *testing.T) {
	g := Goblin(t)
	g.Describe("ParseCsv", func() {
		// Passing Test
		g.It("Parse one-line", func() {
			source := "hi;it;me"

			res := ParseCsv(source)

			expect := [][]string{
				{"hi", "it", "me"},
			}

			g.Assert(res).Eql(expect)
		})

		g.It("Parse multi-line", func() {
			source := "hi;it;me\nhah;h\ni\ng;1;23"

			res := ParseCsv(source)

			expect := [][]string{
				{"hi", "it", "me"},
				{"hah", "h"},
				{"i"},
				{"g", "1", "23"},
			}

			g.Assert(res).Eql(expect)
		})

		g.It("Parse with \\r\\n", func() {
			source := "hi;it;me\r\nhah;h\r\ni\r\ng;1;23"

			res := ParseCsv(source)

			expect := [][]string{
				{"hi", "it", "me"},
				{"hah", "h"},
				{"i"},
				{"g", "1", "23"},
			}

			g.Assert(res).Eql(expect)
		})

		g.It("Parse with quotes", func() {
			source := "\"hi;it\";me\r\n\"hah;\"h\ni\ng;\"1;2;3\";23"

			res := ParseCsv(source)

			expect := [][]string{
				{"hi;it", "me"},
				{"hah;h"},
				{"i"},
				{"g", "1;2;3", "23"},
			}

			g.Assert(res).Eql(expect)
		})

		g.It("Parse with wrong quotes", func() {
			source := "hi;it;\"me\r\nhah;\"h\ni\ng;\"1;2;3\";23"

			res := ParseCsv(source)

			expect := [][]string{
				{"hi", "it", "me"},
				{"hah", "h"},
				{"i"},
				{"g", "1;2;3", "23"},
			}

			g.Assert(res).Eql(expect)
		})

		g.It("Parse lines weird ;", func() {

			source := "hi;;it;\nme;1"

			res := ParseCsv(source)

			expect := [][]string{
				{"hi", "", "it", ""},
				{"me", "1"},
			}

			g.Assert(res).Eql(expect)
		})
	})
}

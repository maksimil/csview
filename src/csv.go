package src

import (
	"os"
	"strings"
)

type CsvDocument struct {
	Path string
	Data [][]string
}

func ParseCsv(source string) [][]string {
	lines := strings.Split(strings.ReplaceAll(source, "\r\n", "\n"), "\n")

	data := make([][]string, len(lines))

	for i, line := range lines {
		quoted := false

		data_row := []string{""}

		for _, c := range line {
			if c == '"' {
				quoted = !quoted
			} else if c == ';' && !quoted {
				data_row = append(data_row, "")
			} else {
				data_row[len(data_row)-1] += string(c)
			}
		}

		data[i] = data_row
	}

	return data
}

func ParseFile(path string) (CsvDocument, error) {
	f, err := os.ReadFile(path)

	if err != nil {
		return CsvDocument{}, err
	}

	data := ParseCsv(string(f))
	return CsvDocument{
		Data: data,
		Path: path,
	}, nil
}

package asciit

import (
	"embed"
	"strconv"
	"strings"
)

//go:embed ascii.txt
var f embed.FS

func Table() map[int]string {
	data, _ := f.ReadFile("ascii.txt")
	rows := strings.Split(string(data), "\n")

	ranges := NewRanges()
	ranges.load(rows)

	table := map[int]string{}

	for i := range ranges.data {
		row := ranges.data[i]
		cols := strings.Split(row, " ")
		codeTrim := strings.TrimSpace(cols[0])
		code, err := strconv.Atoi(codeTrim)
		if err != nil {
			panic(err)
		}

		table[code] = cols[4]
	}

	return table
}

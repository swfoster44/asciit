package asciit

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

type entry struct {
	code int
	char string
	tag  int
}

//go:embed ascii.txt
var f embed.FS

func Create() {
	data, _ := f.ReadFile("ascii.txt")
	rows := strings.Split(string(data), "\n")

	table := make(map[int]entry, 128)
	tagger := NewTagger()

	// -1 is for a hidden space or something
	// i don't want to track down at this time
	for i := 0; i < len(rows)-1; i++ {
		cols := strings.Split(rows[i], " ")
		e := entry{}

		codeTrim := strings.TrimSpace(cols[0])
		code, err := strconv.Atoi(codeTrim)
		if err != nil {
			panic(err)
		}

		e.code = code
		e.char = cols[4]
		if ok := tagger.tag(&e); ok {
			table[e.code] = e
		}
		// entries = append(entries, e)
	}
	fmt.Println(table)
}

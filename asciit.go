package asciit

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed ascii.txt
var f embed.FS

const (
	NUMERIC = iota + 100
	ALPHAL
	ALPHAU
	SPACE
	SYMBOL
)

type Code = uint64

type Entry struct {
	Code Code
	Hex  string
	Html string
	Char string
	Desc string
	Tag  int
}

func tagger(c Code) int {
	switch {
	case c >= 48 && c <= 57:
		return NUMERIC
	case c >= 97 && c <= 122:
		return ALPHAL
	case c >= 65 && c <= 90:
		return ALPHAU
	case (c >= 8 && c <= 13) || c == 32:
		return SPACE
	case c >= 33 && c <= 47:
		return SYMBOL
	case c >= 58 && c <= 64:
		return SYMBOL
	case c >= 91 && c <= 96:
		return SYMBOL
	case c >= 123 && c <= 126:
		return SYMBOL
	default:
		return 0
	}
}

type Table struct {
	codeMap map[Code]*Entry
	strMap  map[string]*Entry
}

func (T *Table) ByCode(c Code) *Entry {
	return T.codeMap[c]
}
func (T *Table) ByStr(s string) *Entry {
	return T.strMap[s]
}

func newTable() *Table {
	cm := map[Code]*Entry{}
	sm := map[string]*Entry{}
	return &Table{codeMap: cm, strMap: sm}
}

func New() *Table {
	data, _ := f.ReadFile("ascii.txt")
	rows := strings.Split(string(data), "\n")

	table := newTable()

	// - 1 because spliting on new line
	for i := 0; i < len(rows)-1; i++ {
		row := rows[i]
		cols := strings.Split(row, " ")

		code, err := strconv.ParseUint(strings.TrimSpace(cols[0]), 10, 64)
		if err != nil {
			panic(err)
		}

		tag := tagger(code)

		if tag != 0 {
			hex := cols[1]
			html := cols[3]
			char := cols[4]
			var desc string = ""

			if len(cols) > 5 {
				desc = strings.Join(cols[5:], " ")
			}

			e := &Entry{code, hex, html, char, desc, tag}
			table.codeMap[code] = e
			table.strMap[char] = e
			// fmt.Printf("%v, %v, %v, %v, %v, %v\n", code, hex, html, char, desc, tag)
		}
	}
	return table
}

package asciit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRanges(t *testing.T) {
	assert := assert.New(t)

	cr := codeRange{48, 57}
	assert.True(cr.InRange(57))
	assert.True(cr.InRange(48))

	symbols := symbols_ranges()
	assert.Len(symbols.crange, 4)

	ranges := NewRanges()
	assert.Len(ranges.allRanges, 5)

	table := Table()
	assert.Len(table, 101)

	char := table[97]
	assert.Equal("a", char)
}

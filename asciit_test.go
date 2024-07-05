package asciit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTagger(t *testing.T) {
	assert := assert.New(t)

	cr := codeRange{49, 9}
	assert.True(cr.InRange(9))
}

package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractCodeData(t *testing.T) {
	// given/when
	file, line, name := ExtractCodeData()

	// then
	assert.NotEmpty(t, file)
	assert.NotEmpty(t, line)
	assert.NotEmpty(t, name)
}

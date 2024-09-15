package fileio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileReader_Read(t *testing.T) {

	// get file path from the home directory

	reader := NewFileReader("../../examples/testdata/test.txt")
	lines, err := reader.ReadLines()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(lines))
	assert.Equal(t, "hello", lines[0])
	assert.Equal(t, "world", lines[1])
	assert.Equal(t, "foo", lines[2])
}

package fileio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileWriter_Write(t *testing.T) {

	reader := NewFileReader("../../examples/testdata/testwriterinit.txt")
	linesInit, err := reader.ReadLines()
	assert.Nil(t, err)
	writer := NewFileWriter("../../examples/testdata/testwriterfinal.txt")
	err = writer.WriteLines(linesInit)
	assert.Nil(t, err)
	readerOne := NewFileReader("../../examples/testdata/testwriterfinal.txt")
	linesFinal, err := readerOne.ReadLines()
	assert.Nil(t, err)
	assert.Equal(t, linesInit, linesFinal)

}

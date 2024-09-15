package fileio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileWriter struct {
	filePath string
}

func NewFileWriter(filepath string) *FileWriter {
	return &FileWriter{
		filePath: filepath,
	}
}

func (f *FileWriter) WriteLines(lines []string) error {
	file, err := os.OpenFile(f.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	fmt.Printf("File %s written successfully\n", f.filePath)
	return writer.Flush()
}

func (f *FileWriter) WriteLinesWithSlice(lines []string) {
	err := os.WriteFile(f.filePath, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		fmt.Println("Error in file:", err)
		return
	}

}

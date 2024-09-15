package fileio

import (
	"bufio"
	"fmt"
	"os"
)

type FileReader struct {
	filePath string
}

func NewFileReader(filepath string) *FileReader {
	return &FileReader{
		filePath: filepath,
	}
}

// open up the file with the folder provided on the command line
func (f *FileReader) ReadLines() ([]string, error) {
	file, err := os.Open(f.filePath)
	if err != nil {
		return nil, err
	}
	var lines []string
	defer file.Close()
	buf := make([]byte, 0, 64*1024)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func (f *FileReader) ReadAll() (string, error) {
	file, err := os.Open(f.filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf := make([]byte, 0, 64*1024)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(buf, 1024*1024)
	var content string
	for scanner.Scan() {
		content += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content, nil
}

// reads files buffer, more memory effecient
func (f *FileReader) ReadStream() (*bufio.Scanner, func(), error) {
	file, err := os.Open(f.filePath)
	if err != nil {
		fmt.Println("Error reading files")
		return nil, nil, err
	}

	buf := make([]byte, 64*1024)
	scanner := bufio.NewScanner(file)
	scanner.Buffer(buf, 1024*1024)

	cleanup := func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}
	return scanner, cleanup, nil
}

// // reads files line by line
// func (f *FileReader) ReadEachLine() ([]string, error) {
// 	file, err := os.Open(f.filePath)
// 	if err != nil {
// 		fmt.Println("Error reading files")
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	var allLine []string

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		exist := findConsole.ReadConsoleLogs(line)
// 		if !exist {
// 			allLine = append(allLine, line)
// 		} else {
// 			allLine = append(allLine, "")

// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return nil, err
// 	}

// 	return allLine, nil
// }

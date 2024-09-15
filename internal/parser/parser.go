package parser

import (
	"fmt"
	"log"
	"sync"

	"github.com/dlclark/regexp2"
	"github.com/wadmit/era/internal/fileio"
	"github.com/wadmit/era/internal/parser/rules"
	"github.com/wadmit/era/internal/transform"
)

type Parser interface {
	Parse() error
}

type Parse struct {
	fileExtension []string
	outputReg     []string
}

func processLine(file string, regexList []*regexp2.Regexp, resultChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	matched := false
	for _, reg := range regexList {
		match, _ := reg.MatchString(file)
		if match {
			matched = true
			break
		}
	}
	if !matched {
		resultChan <- file
	}
}

func ParseAndWrite(transformed transform.Transform, configMap *rules.ConfigMap) {
	// Generate rules for the given extension
	rule := rules.GenerateRulesForExtensions(transformed.Extension, configMap)
	if rule == nil {
		log.Printf("No rules found for extension: %s", transformed.Extension)
		return
	}

	// Create a slice to hold the filtered lines
	filteredLines := make([]string, len(transformed.ContentLines))

	// Create a wait group for synchronization
	var wg sync.WaitGroup

	// Start processing lines concurrently
	for i, file := range transformed.ContentLines {
		wg.Add(1)
		go func(i int, file string) {
			defer wg.Done()
			matched := false
			for _, reg := range rule.Regex {
				match, _ := reg.MatchString(file)
				if match {
					matched = true
					fmt.Printf("Matched line: %s with regex: %s\n", file, reg.String()) // Debugging line
					break
				}
			}
			if !matched {
				filteredLines[i] = file // Preserve the order by storing in the correct index
			}
		}(i, file)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Remove empty strings from the filteredLines slice
	var finalLines []string
	for _, line := range filteredLines {
		if line != "" {
			finalLines = append(finalLines, line)
		}
	}

	// Write the filtered lines to the file
	fileWriter := fileio.NewFileWriter(transformed.FilePath)
	if err := fileWriter.WriteLines(finalLines); err != nil {
		log.Printf("Error writing to file: %v", err)
	}
}

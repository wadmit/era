package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/wadmit/era/internal/fileio"
	"github.com/wadmit/era/internal/parser/rules"
	"github.com/wadmit/era/internal/transform"
)

// Report structure to store the matched information
type Report struct {
	FileName     string      `json:"filename"`
	MatchedLines []MatchInfo `json:"matched_lines"`
}

// MatchInfo holds the information about each matched line
type MatchInfo struct {
	LineIndex int    `json:"line_index"`
	Content   string `json:"content"`
}

func ParseAndWrite(transformed transform.Transform, configMap *rules.ConfigMap, reportPath string) {
	// Generate rules for the given extension
	rule := rules.GenerateRulesForExtensions(transformed.Extension, configMap)
	if rule == nil {
		log.Printf("No rules found for extension: %s", transformed.Extension)
		return
	}

	// Pre-allocate filtered lines with the size of the content lines
	filteredLines := make([]string, len(transformed.ContentLines))

	// Buffered channel to collect results for matched lines
	reportChan := make(chan Report, len(transformed.ContentLines))

	// Wait group for synchronizing goroutines
	var wg sync.WaitGroup

	// Process each line concurrently
	for i, line := range transformed.ContentLines {
		wg.Add(1)
		go func(i int, line string) {
			defer wg.Done()

			// Check if the line matches any of the rules
			var matchedLines []MatchInfo
			for _, reg := range rule.Regex {
				match, _ := reg.MatchString(line)
				if match {
					matchedLines = append(matchedLines, MatchInfo{
						LineIndex: i,
						Content:   line,
					})
					break // Exit early if a match is found
				}
			}

			// If matched, send the result to the report channel
			if len(matchedLines) > 0 {
				reportChan <- Report{
					FileName:     transformed.FilePath,
					MatchedLines: matchedLines,
				}
			} else {
				// Preserve the order of filtered lines
				filteredLines[i] = line
			}
		}(i, line)
	}

	// Close the report channel once all processing is done
	go func() {
		wg.Wait()
		close(reportChan)
	}()

	// Collect results from the report channel
	var reports []Report
	for report := range reportChan {
		reports = append(reports, report)
	}

	// Remove empty strings from the filtered lines
	finalLines := make([]string, 0, len(filteredLines))
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

	// Ensure the directory for the report path exists, create it if necessary
	if err := os.MkdirAll(filepath.Dir(reportPath), os.ModePerm); err != nil {
		log.Printf("Error creating report directory: %v", err)
		return
	}

	// Handle JSON appending in a valid way
	var existingReports []Report

	// Check if the file exists
	if _, err := os.Stat(reportPath); err == nil {
		// File exists, read the current content
		data, err := ioutil.ReadFile(reportPath)
		if err != nil {
			log.Printf("Error reading report file: %v", err)
			return
		}

		// Unmarshal the existing content into the report slice
		if err := json.Unmarshal(data, &existingReports); err != nil {
			log.Printf("Error unmarshaling existing report file: %v", err)
			return
		}
	}

	// Merge new reports with existing ones
	for _, newReport := range reports {
		found := false
		for i, existingReport := range existingReports {
			if existingReport.FileName == newReport.FileName {
				// Append matched lines to the existing report
				existingReports[i].MatchedLines = append(existingReports[i].MatchedLines, newReport.MatchedLines...)
				found = true
				break
			}
		}
		if !found {
			// If the file is not already in the report, add it as a new entry
			existingReports = append(existingReports, newReport)
		}
	}

	// Convert the entire report (existing + new) to JSON format
	reportJSON, err := json.MarshalIndent(existingReports, "", "  ")
	if err != nil {
		log.Printf("Error marshaling report to JSON: %v", err)
		return
	}

	// Write the updated JSON report to the file (overwrite)
	if err := os.WriteFile(reportPath, reportJSON, 0644); err != nil {
		log.Printf("Error writing JSON report to file: %v", err)
	}
}

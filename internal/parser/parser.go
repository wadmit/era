package parser

import (
	"encoding/json"
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
	LineIndex  int    `json:"line_index"`
	Content    string `json:"content"`
	StartIndex int    `json:"start_index"`
	EndIndex   int    `json:"end_index"`
}

func ParseAndWrite(transformed transform.Transform, configMap *rules.ConfigMap, reportPath string) {
	rule := rules.GenerateRulesForExtensions(transformed.Extension, configMap)
	if rule == nil {
		log.Printf("No rules found for extension: %s", transformed.Extension)
		return
	}

	// Use separate filtering and matching logic
	reportChan := make(chan Report, len(transformed.ContentLines))
	var wg sync.WaitGroup

	for i, line := range transformed.ContentLines {
		wg.Add(1)
		go func(i int, line string) {
			defer wg.Done()

			// First, check if the line should be ignored
			if rules.ShouldIgnoreLine(line, rule.IgnoreRegex) {
				return
			}

			var matchInfos []MatchInfo
			// Now apply each regex to the line
			for _, reg := range rule.Regex {
				matches := reg.FindAllStringIndex(line, -1)
				for _, match := range matches {
					matchInfos = append(matchInfos, MatchInfo{
						LineIndex:  i,
						Content:    line[match[0]:match[1]],
						StartIndex: match[0],
						EndIndex:   match[1],
					})
				}
			}

			if len(matchInfos) > 0 {
				reportChan <- Report{
					FileName:     transformed.FilePath,
					MatchedLines: matchInfos,
				}
			}
		}(i, line)
	}

	go func() {
		wg.Wait()
		close(reportChan)
	}()

	var reports []Report
	for report := range reportChan {
		reports = append(reports, report)
	}

	// Remove matched substrings from the content
	finalLines := make([]string, len(transformed.ContentLines))
	copy(finalLines, transformed.ContentLines)
	for _, report := range reports {
		for _, matchInfo := range report.MatchedLines {
			finalLines[matchInfo.LineIndex] = finalLines[matchInfo.LineIndex][:matchInfo.StartIndex] + finalLines[matchInfo.LineIndex][matchInfo.EndIndex:]
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
		data, err := os.ReadFile(reportPath)
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

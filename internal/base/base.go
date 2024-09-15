package base

import (
	"fmt"
	"os"

	"github.com/wadmit/era/internal/parser"
	"github.com/wadmit/era/internal/parser/rules"
	"github.com/wadmit/era/internal/transform"
	"github.com/wadmit/era/internal/types"
	"github.com/wadmit/era/internal/utils"
)

// DetectAndChangeFile processes files, modifies their content, and then writes it back or handles it as needed.
func DetectAndChangeFile(root string, cfg *types.Config, configMap *rules.ConfigMap) {
	// Create channels for transformed files and file results
	fileTransformChan := make(chan transform.Transform) // Buffered channel for batch processing
	resultChan := make(chan transform.Transform, 10)    // Buffered channel for file results
	// // Create a transformer and start processing files in a separate goroutine
	fileTransformer := &transform.FileTransformer{}
	go fileTransformer.Transform(root, cfg, fileTransformChan)
	// // Start worker pool for processing transformed files
	go func() {
		for transformed := range fileTransformChan {
			resultChan <- transformed
		}
		close(resultChan)
	}()

	jsonFilePath, err := utils.CreateReportPath(cfg.ReportPath)
	if err != nil {
		fmt.Print("Error: Unable to create report path")
		os.Exit(1)
		return

	}
	for transformed := range resultChan {
		parser.ParseAndWrite(transformed, configMap, jsonFilePath)
	}
}

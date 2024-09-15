package base

import (
	"github.com/wadmit/eradicate/internal/parser"
	"github.com/wadmit/eradicate/internal/parser/rules"
	"github.com/wadmit/eradicate/internal/transform"
	"github.com/wadmit/eradicate/internal/types"
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

	// // Process each transformed file
	for transformed := range resultChan {
		// fmt.Print(transformed.ContentLines)
		parser.ParseAndWrite(transformed, configMap)
	}
}
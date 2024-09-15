package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wadmit/eradicate/internal/types"
)

func TestFileWalker(t *testing.T) {
	// Setup: Create a temporary directory structure
	tempDir, err := os.MkdirTemp("", "filewalker_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir) // Clean up after test

	// Create directories and files
	err = os.Mkdir(filepath.Join(tempDir, "dir1"), 0755)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Mkdir(filepath.Join(tempDir, "dir2"), 0755)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(tempDir, "dir1", "file1.txt"), []byte("content"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(tempDir, "dir2", "file2.log"), []byte("content"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(tempDir, "dir2", "file3.tmp"), []byte("content"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Create a symlink to dir1
	err = os.Symlink(filepath.Join(tempDir, "dir1"), filepath.Join(tempDir, "symlink_to_dir1"))
	if err != nil {
		t.Fatal(err)
	}

	// Create a symlink to a file
	err = os.Symlink(filepath.Join(tempDir, "dir1", "file1.txt"), filepath.Join(tempDir, "symlink_to_file1.txt"))
	if err != nil {
		t.Fatal(err)
	}

	// Define ignore rules, including file extensions
	config := &types.Config{
		IgnoreDirs:           []string{"dir2", "/symlink_to_dir1"},
		IgnoreFiles:          []string{"file2.log", "/symlink_to_file1.txt"},
		IgnoreFileExtensions: []string{".tmp"},
	}

	// Run the FileWalker
	fileChan := make(chan string)
	go FileWalker(tempDir, fileChan, config)

	// Collect results
	var foundFiles []string
	for file := range fileChan {
		relPath, err := filepath.Rel(tempDir, file)
		if err != nil {
			t.Fatal(err)
		}
		foundFiles = append(foundFiles, relPath)
	}
	fmt.Print(foundFiles)

	// Validate the output
	expectedFiles := []string{
		"dir1/file1.txt", // This should be included
	}

	assert.ElementsMatch(t, expectedFiles, foundFiles)
}

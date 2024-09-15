package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/wadmit/era/internal/types"
)

// FileWalker walks through every folder and file in the given directory, sending file paths
// to the provided channel. It skips directories and files listed in the config's IgnoreDirs and IgnoreFiles,
// and also files with specific extensions listed in IgnoreFileExtensions. Symlinks are resolved and handled accordingly.
func FileWalker(root string, fileChan chan<- string, config *types.Config) {
	ignoreDirsSet := make(map[string]struct{})
	for _, dir := range config.IgnoreDirs {
		ignoreDirsSet[dir] = struct{}{}
	}

	ignoreFilesSet := make(map[string]struct{})
	for _, file := range config.IgnoreFiles {
		ignoreFilesSet[file] = struct{}{}
	}

	ignoreExtSet := make(map[string]struct{})
	for _, ext := range config.IgnoreFileExtensions {
		ignoreExtSet[ext] = struct{}{}
	}

	// Set to track seen file paths
	seenFiles := make(map[string]struct{})

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Resolve symlink
		if info.Mode()&os.ModeSymlink != 0 {
			realPath, err := filepath.EvalSymlinks(path)
			if err != nil {
				return err
			}
			info, err = os.Stat(realPath)
			if err != nil {
				return err
			}
			path = realPath
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		// Early exit for ignored directories
		if info.IsDir() {
			if _, found := ignoreDirsSet[info.Name()]; found || strings.HasPrefix(relPath, "/") && (strings.HasPrefix(relPath, "/"+info.Name())) {
				return filepath.SkipDir
			}
		} else {
			if _, found := ignoreFilesSet[info.Name()]; found || strings.HasPrefix(relPath, "/") && (strings.HasPrefix(relPath, "/"+info.Name())) {
				return nil
			}

			ext := filepath.Ext(info.Name())
			_, found := ignoreExtSet[ext]
			if found {
				return nil
			}
		}

		// Check and send file paths if not already processed
		if !info.IsDir() {
			if _, seen := seenFiles[path]; !seen {
				seenFiles[path] = struct{}{}
				fileChan <- path
			}
		}

		return nil
	})

	if err != nil {
		fmt.Print("err", err)
		panic(err)
	}
}

func GetFileExt(filePath string) string {
	return filepath.Ext(filePath)
}

func CreateReportPath(reportDir string) (string, error) {
	// Ensure the directory exists, or create it
	if _, err := os.Stat(reportDir); os.IsNotExist(err) {
		if err := os.MkdirAll(reportDir, 0755); err != nil {
			return "", fmt.Errorf("error creating directory: %v", err)
		}
	}

	// Generate 16 random bytes for uniqueness
	randBytes := make([]byte, 16)
	if _, err := rand.Read(randBytes); err != nil {
		return "", fmt.Errorf("error generating random bytes: %v", err)
	}

	// Combine timestamp and random bytes for uniqueness
	timestamp := time.Now().Format("20060102-150405")
	uniqueID := hex.EncodeToString(randBytes)
	jsonFileName := fmt.Sprintf("report-%s-%s.json", timestamp, uniqueID)

	// Create the full path for the JSON report
	jsonFilePath := filepath.Join(reportDir, jsonFileName)
	return jsonFilePath, nil
}

package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

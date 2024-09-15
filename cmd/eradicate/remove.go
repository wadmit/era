package eradicate

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wadmit/era/internal/base"
	"github.com/wadmit/era/internal/config"
	"github.com/wadmit/era/internal/parser"
	"github.com/wadmit/era/internal/parser/rules"
	"github.com/wadmit/era/internal/transform"
	"github.com/wadmit/era/internal/types"
)

var RemoveCommand = &cobra.Command{
	Use:   "remove",
	Short: "Remove the files based on the rules",
	Run: func(cmd *cobra.Command, args []string) {

		fileFlag, _ := cmd.Flags().GetStringArray("file")
		dirFlag, _ := cmd.Flags().GetString("dir")

		if len(fileFlag) == 0 && dirFlag == "" {
			fmt.Println("Error: Provide either --file or --dir, but not both.")
			cmd.Help()
			os.Exit(1)
		}

		// Load the config
		cfg, err := config.LoadDefaultConfig()
		if err != nil {
			fmt.Println("Error: Unable to load config file:", err)
			os.Exit(1)
		}

		configMap := rules.LoadRules(cfg)

		// handle multiple files flag
		for _, file := range fileFlag {
			handleFileRemoval(file, cfg, configMap)
		}
		// Handle directory removal
		if dirFlag != "" {
			handleDirectoryRemoval(dirFlag, cfg, configMap)
		}
	},
}

// Check if the file exists
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Handle file removal logic
func handleFileRemoval(file string, cfg *types.Config, configMap *rules.ConfigMap) {
	// join the root path with the file
	filePath := JoinRootPath(file, cfg.Root)
	if !fileExists(filePath) {
		fmt.Println("Error: File does not exist")
		return
	}

	// Check if the file is in the ignore list
	if rules.Contains(cfg.IgnoreFiles, file) {
		fmt.Println("Error: File is in ignore list")
		return
	}

	transformedFile, err := transform.TransformFile(filePath)
	if err != nil {
		fmt.Println("Error reading file as if file path is invalid or path is directory", err)
		return
	}

	parser.ParseAndWrite(transformedFile, configMap)
}

// Handle directory removal logic
func handleDirectoryRemoval(dir string, cfg *types.Config, configMap *rules.ConfigMap) {
	if !fileExists(dir) {
		fmt.Println("Error: Directory does not exist")
		return
	}
	// Check if the directory is in the ignore list
	if rules.Contains(cfg.IgnoreDirs, dir) {
		fmt.Println("Error: Directory is in ignore list")
		return
	}
	base.DetectAndChangeFile(dir, cfg, configMap)
}

func JoinRootPath(path, root string) string {
	if root == "" {
		return path
	}
	return root + "/" + path
}

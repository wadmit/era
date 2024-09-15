package utils

import (
	"os"
	"path/filepath"
)

func FindConfigFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		configPath := filepath.Join(dir, "etc.yaml")
		if _, err := os.Stat(configPath); err == nil {
			return configPath, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", os.ErrNotExist
}

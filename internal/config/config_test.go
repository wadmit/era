package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/wadmit/era/internal/utils"
)

func TestLoadConfig(t *testing.T) {
	// Test loading default config
	cfg, err := LoadConfig()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, ".", cfg.Root)

	// Create a temporary config file
	tempDir, err := os.MkdirTemp("", "config_test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	tempConfigPath := filepath.Join(tempDir, "etc.yaml")
	v := viper.New()
	v.SetConfigFile(tempConfigPath)
	v.SetConfigType("yaml")
	v.Set("root", "/custom/root")
	v.Set("reportPath", "/custom/report")
	v.Set("IgnoreKeyword", []string{"custom_ignore"})
	v.Set("IgnoreFileExtensions", []string{".custom"})
	v.Set("IgnoreDir", []string{"custom_dir"})
	v.Set("IgnoreFiles", []string{"custom_file"})
	v.Set("ListenType", "GitBeforePush")

	err = v.WriteConfigAs(tempConfigPath)
	assert.NoError(t, err)

	// Change working directory to temp dir
	oldWd, _ := os.Getwd()
	err = os.Chdir(tempDir)
	assert.NoError(t, err)
	defer os.Chdir(oldWd)

	// Test loading custom config
	cfg, err = LoadConfig()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "/custom/root", cfg.Root)
	assert.Equal(t, "/custom/report", cfg.ReportPath)
	assert.Equal(t, []string{"custom_ignore"}, cfg.IgnoreKeyword)
	assert.Equal(t, []string{".custom"}, cfg.IgnoreFileExtensions)
	assert.Equal(t, []string{"custom_dir"}, cfg.IgnoreDirs)
	assert.Equal(t, []string{"custom_file"}, cfg.IgnoreFiles)
	assert.Equal(t, "GitBeforePush", cfg.ListenType)
}

func TestInitConfig(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "init_config_test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Change working directory to temp dir
	oldWd, _ := os.Getwd()
	err = os.Chdir(tempDir)
	assert.NoError(t, err)
	defer os.Chdir(oldWd)

	// Create the default config file in the temporary directory
	defaultConfigPath := filepath.Join(tempDir, "generate")
	err = os.Mkdir(defaultConfigPath, 0755)
	assert.NoError(t, err)

	defaultConfigContent := `
			root: .
			reportPath: .
			IgnoreKeyword: ["erd:ignore", "erd:ignoreAll"]
			IgnoreFileExtensions: [".exe", ".dll", ".so", ".dylib", ".zip", ".tar", ".gz", ".rar"]
			IgnoreDir: ["node_modules", "vendor", ".git", ".idea", ".vscode", ".vs", ".hg", ".svn", ".bzr", ".fslckout", "_darcs", "_sgbak"]
			IgnoreFiles: []
			ListenType: "command"
	`
	err = os.WriteFile(filepath.Join(defaultConfigPath, "etc.yaml"), []byte(defaultConfigContent), 0644)
	assert.NoError(t, err)

	// Run InitConfig
	err = InitConfig()
	assert.NoError(t, err)

	// Check if config file was created
	_, err = os.Stat("etc.yaml")
	assert.NoError(t, err)

	// Load the created config
	cfg, err := LoadConfig()
	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	// Check if values match default config
	defaultCfg, err := LoadDefaultConfig()
	assert.NoError(t, err)
	assert.Equal(t, defaultCfg.Root, cfg.Root)
	assert.Equal(t, defaultCfg.ReportPath, cfg.ReportPath)
	assert.Equal(t, defaultCfg.IgnoreKeyword, cfg.IgnoreKeyword)
	assert.Equal(t, defaultCfg.IgnoreFileExtensions, cfg.IgnoreFileExtensions)
	assert.Equal(t, defaultCfg.IgnoreDirs, cfg.IgnoreDirs)
	assert.Equal(t, defaultCfg.IgnoreFiles, cfg.IgnoreFiles)
	assert.Equal(t, defaultCfg.ListenType, cfg.ListenType)

	// Test initializing when config already exists
	err = InitConfig()
	assert.NoError(t, err) // Should not return an error, just skip initialization
}

func TestFindConfigFile(t *testing.T) {
	// Create a temporary directory structure
	tempDir, err := os.MkdirTemp("", "find_config_test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	subDir := filepath.Join(tempDir, "subdir")
	err = os.Mkdir(subDir, 0755)
	assert.NoError(t, err)

	configPath := filepath.Join(tempDir, "etc.yaml")
	_, err = os.Create(configPath)
	assert.NoError(t, err)

	// Test finding config in parent directory
	oldWd, _ := os.Getwd()
	err = os.Chdir(subDir)
	assert.NoError(t, err)
	defer os.Chdir(oldWd)

	foundPath, err := utils.FindConfigFile()
	assert.NoError(t, err)
	assert.Equal(t, configPath, foundPath)

	// Test when no config file exists
	err = os.RemoveAll(tempDir)
	assert.NoError(t, err)

	foundPath, err = utils.FindConfigFile()
	assert.Error(t, err)
	assert.Equal(t, "", foundPath)
}

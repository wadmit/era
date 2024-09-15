package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func InitConfig() error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	v := viper.New()

	configPath, err := os.Getwd()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get current directory")
		return err
	}
	configPath = filepath.Join(configPath, "etc.yaml")

	if _, err := os.Stat(configPath); err == nil {
		log.Warn().Msg("Config file already exists. Skipping initialization.")
		return nil
	}

	defaultConfig, err := LoadDefaultConfig()
	if err != nil {
		log.Error().Err(err).Msg("Failed to load default config")
		return err
	}

	v.SetConfigFile(configPath)
	v.SetConfigType("yml")

	v.Set("root", defaultConfig.Root)
	v.Set("reportPath", defaultConfig.ReportPath)
	v.Set("IgnoreKeyword", defaultConfig.IgnoreKeyword)
	v.Set("IgnoreFileExtensions", defaultConfig.IgnoreFileExtensions)
	v.Set("IgnoreDir", defaultConfig.IgnoreDirs)
	v.Set("IgnoreFiles", defaultConfig.IgnoreFiles)
	v.Set("ListenType", defaultConfig.ListenType)

	if err := v.WriteConfig(); err != nil {
		log.Error().Err(err).Msg("Failed to write config file")
		return err
	}

	fmt.Println("Configuration file created successfully at:", configPath)
	return nil
}

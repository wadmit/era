package config

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/wadmit/eradicate/internal/types"
	"github.com/wadmit/eradicate/internal/utils"
)

const (
	ListenTypeCommand        = "command"
	ListenTypeGitBeforePush  = "git-before-push"
	ListenTypeGitAfterCommit = "git-after-commit"
)

func LoadConfig() (*types.Config, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	v := viper.New()
	v.SetConfigName("etc")
	v.SetConfigType("yaml")

	configPath, err := utils.FindConfigFile()
	if err == nil {
		v.SetConfigFile(configPath)
	} else {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Error().Err(err).Msg("Failed to get current directory")
			return nil, err
		}
		v.AddConfigPath(currentDir)
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msg("Config file not found, using default config")
			return LoadDefaultConfig()
		}
		log.Error().Err(err).Msg("Failed to read config file")
		return nil, err
	}

	var config types.Config
	if err := v.Unmarshal(&config); err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal config")
		return nil, err
	}

	return &config, nil
}

func LoadDefaultConfig() (*types.Config, error) {
	config := *GenerateErdConfig()
	return &config, nil
}

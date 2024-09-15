package eradicate

import (
	"github.com/spf13/cobra"
	"github.com/wadmit/eradicate/internal/config"
)

var InitCommand = &cobra.Command{
	Use:   "init",
	Short: "It will generate a default configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig()
	},
}

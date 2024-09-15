package eradicate

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wadmit/era/internal/base"
	"github.com/wadmit/era/internal/config"
	"github.com/wadmit/era/internal/parser/rules"
)

var CleanCommand = &cobra.Command{
	Use:   "clean",
	Short: "It will go through the files and clean them",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println("Error: Loading config file")
			panic(err)
		}
		fmt.Print("Cleaning files\n", cfg)
		configMap := rules.LoadRules(cfg)
		base.DetectAndChangeFile(cfg.Root, cfg, configMap)
	},
}

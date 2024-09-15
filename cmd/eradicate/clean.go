package eradicate

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wadmit/eradicate/internal/base"
	"github.com/wadmit/eradicate/internal/config"
	"github.com/wadmit/eradicate/internal/parser/rules"
)

var CleanCommand = &cobra.Command{
	Use:   "clean",
	Short: "It will go through the files and clean them",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Println("Error loading config file")
			panic(err)
		}
		fmt.Print("Cleaning files\n", cfg)
		// root := cmd.Flag("root").Value.String()
		// if root == "" {
		// 	root = cfg.Root
		// }
		configMap := rules.LoadRules(cfg)
		base.DetectAndChangeFile(cfg.Root, cfg, configMap)
	},
}

package eradicate

import (
	"fmt"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure" // For ASCII art
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// ASCII art for the tool
var asciiArt = figure.NewFigure("Eradicate", "", true).String()

var rootCmd = &cobra.Command{
	Use:   "era",
	Short: "Eradicate(era) is a magic trick that does magic",
	Long: asciiArt + `
Eradicate (era) is a powerful tool for cleaning and transforming your files with ease.
It comes with built-in commands to initialize configurations and clean files based on custom rules.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Print ASCII art and help message when no command is provided
		cmd.Help()
	},
}

func init() {

	rootCmd.AddCommand(InitCommand)
	rootCmd.AddCommand(CleanCommand)
	rootCmd.AddCommand(RemoveCommand)

	RemoveCommand.Flags().StringArrayP("file", "f", []string{}, "Files to remove")
	RemoveCommand.Flags().StringP("dir", "d", "", "Directory to remove")

	// Adding a global flag for verbosity
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		// Handle global flags like verbosity
		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			log.Logger = log.With().Caller().Logger()
			log.Info().Msg("Verbose mode enabled")
		}
	}
}

// Execute is the entry point for the CLI application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if strings.Contains(err.Error(), "unknown flag") {
			// Exit code 126: Command invoked cannot execute
			fmt.Println("Error: Unknown flag or command")
			os.Exit(126)
		} else {
			log.Fatal().Msg(err.Error())
		}
	}
}

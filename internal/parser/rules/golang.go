package rules

import (
	"github.com/wadmit/eradicate/internal/types"
)

func GoLangConfig(cfg *types.Config) *Rule {

	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`fmt\.Print\((.*?)\)`,
		`fmt\.Printf\((.*?)\)`,
		`fmt\.Println\((.*?)\)`,
		`log\.Print\((.*?)\)`,
		`log\.Printf\((.*?)\)`,
		`log\.Println\((.*?)\)`,
		`os\.Stdout\.Write\((.*?)\)`,
		`os\.Stdout\.WriteString\((.*?)\)`,
		`os\.Stderr\.Write\((.*?)\)`,
		`os\.Stderr\.WriteString\((.*?)\)`,
		`os\.Stdout\.WriteAt\((.*?)\)`,
		`os\.Stdout\.WriteStringAt\((.*?)\)`,
		`os\.Stderr\.WriteAt\((.*?)\)`,
		`os\.Stderr\.WriteStringAt\((.*?)\)`,
		`os\.Stdout\.WriteFile\((.*?)\)`,
		`os\.Stderr\.WriteFile\((.*?)\)`,
		`os\.Stdout\.WriteStringFile\((.*?)\)`,
		`os\.Stderr\.WriteStringFile\((.*?)\)`,
		`os\.Stdout\.WriteStringFileAt\((.*?)\)`,
		`os\.Stderr\.WriteStringFileAt\((.*?)\)`,
		`os\.Stdout\.WriteStringFileAt\((.*?)\)`,
	}
	r := Rule{
		Description:    "GoLang Config",
		RuleID:         "GoLangConfig",
		FileExtensions: []string{".go"},
		Keywords:       []string{"golang"},
		Regex:          GenerateCombinedRegex(ignoreKeywords, outputRegexPatterns),
	}
	return &r
}

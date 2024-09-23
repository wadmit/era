package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/wadmit/era/internal/types"
)

func GoLangConfig(cfg *types.Config) *Rule {
	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`fmt\.Print\((.*?)\)\s*;?`,
		`fmt\.Printf\((.*?)\)\s*;?`,
		`fmt\.Println\((.*?)\)\s*;?`,
		`log\.Print\((.*?)\)\s*;?`,
		`log\.Printf\((.*?)\)\s*;?`,
		`log\.Println\((.*?)\)\s*;?`,
		`os\.Stdout\.Write\((.*?)\)\s*;?`,
		`os\.Stdout\.WriteString\((.*?)\)\s*;?`,
		`os\.Stderr\.Write\((.*?)\)\s*;?`,
		`os\.Stderr\.WriteString\((.*?)\)\s*;?`,
		`os\.Stdout\.WriteAt\((.*?)\)\s*;?`,
		`os\.Stderr\.WriteAt\((.*?)\)\s*;?`,
		`os\.Stdout\.WriteFile\((.*?)\)\s*;?`,
		`os\.Stderr\.WriteFile\((.*?)\)\s*;?`,
		`os\.Stdout\.WriteStringFile\((.*?)\)\s*;?`,
		`os\.Stderr\.WriteStringFile\((.*?)\)\s*;?`,
		`os\.Stdout\.WriteStringFileAt\((.*?)\)\s*;?`,
		`os\.Stderr\.WriteStringFileAt\((.*?)\)\s*;?`,
	}

	ignoreRegex := regexp.MustCompile(fmt.Sprintf(`(%s)`, strings.Join(ignoreKeywords, "|")))

	// Compile each output regex pattern
	compiledPatterns := make([]*regexp.Regexp, len(outputRegexPatterns))
	for i, pattern := range outputRegexPatterns {
		compiledPatterns[i] = regexp.MustCompile(pattern)
	}

	return &Rule{
		Description:    "GoLang Config",
		RuleID:         "GoLangConfig",
		FileExtensions: []string{".go"},
		Keywords:       []string{"golang"},
		Regex:          compiledPatterns,
		IgnoreRegex:    ignoreRegex,
	}
}

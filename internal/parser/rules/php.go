package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/wadmit/era/internal/types"
)

func PhpConfig(cfg *types.Config) *Rule {
	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`echo\s*(.*?)\s*;?`,          // Matches echo statements with optional spaces and semicolon
		`print\s*\((.*?)\)\s*;?`,     // Matches print() with optional spaces and semicolon
		`printf\s*\((.*?)\)\s*;?`,    // Matches printf() with optional spaces and semicolon
		`error_log\s*\((.*?)\)\s*;?`, // Matches error_log() with optional spaces and semicolon
	}

	ignoreRegex := regexp.MustCompile(fmt.Sprintf(`(%s)`, strings.Join(ignoreKeywords, "|")))
	// Compile each output regex pattern
	compiledPatterns := make([]*regexp.Regexp, len(outputRegexPatterns))
	for i, pattern := range outputRegexPatterns {
		compiledPatterns[i] = regexp.MustCompile(pattern)
	}

	return &Rule{
		Description:    "PHP Config",
		RuleID:         "PhpConfig",
		FileExtensions: []string{".php"},
		Keywords:       []string{"php"},
		Regex:          compiledPatterns,
		IgnoreRegex:    ignoreRegex,
	}
}

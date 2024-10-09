package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/wadmit/era/internal/types"
)

func DartConfig(cfg *types.Config) *Rule {
	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`print\s*\((.*?)\)\s*;?`,       // Matches print() with  spaces and  semicolon
		`debugPrint\s*\((.*?)\)\s*;?`,  // Matches debugPrint() with  spaces and  semicolon
		`Logger\(\)\.log\s*\((.*?)\)\s*;?`, // Matches Logger.log() with spaces and semicolon
		`Logger\(\)\.(info|warning|severe|shout)\s*\((.*?)\)\s*;?`, // Matches different Logger methods: info, warning, severe, shout
		`log\s*\((.*?)\)\s*;?`,         // Matches the log() function from the dart:developer package
	}

	ignoreRegex := regexp.MustCompile(fmt.Sprintf(`(%s)`, strings.Join(ignoreKeywords, "|")))

	// Compile each output regex pattern
	compiledPatterns := make([]*regexp.Regexp, len(outputRegexPatterns))
	for i, pattern := range outputRegexPatterns {
		compiledPatterns[i] = regexp.MustCompile(pattern)
	}

	return &Rule{
		Description:    "Dart Config",
		RuleID:         "DartConfig",
		FileExtensions: []string{".dart"},
		Keywords:       []string{"dart"},
		Regex:          compiledPatterns,
		IgnoreRegex:    ignoreRegex,
	}
}
package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/wadmit/era/internal/types"
)

func RubyConfig(cfg *types.Config) *Rule {
	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`puts\s*(.*?)\s*$`,              // Matches puts with optional spaces and end of line
		`print\s*(.*?)\s*$`,             // Matches print with optional spaces and end of line
		`warn\s*(.*?)\s*$`,              // Matches warn with optional spaces and end of line
		`logger\.debug\s*\((.*?)\)\s*$`, // Matches logger.debug() with optional spaces and end of line
		`logger\.info\s*\((.*?)\)\s*$`,  // Matches logger.info() with optional spaces and end of line
		`logger\.warn\s*\((.*?)\)\s*$`,  // Matches logger.warn() with optional spaces and end of line
		`logger\.error\s*\((.*?)\)\s*$`, // Matches logger.error() with optional spaces and end of line
		`logger\.fatal\s*\((.*?)\)\s*$`, // Matches logger.fatal() with optional spaces and end of line
	}

	ignoreRegex := regexp.MustCompile(fmt.Sprintf(`(%s)`, strings.Join(ignoreKeywords, "|")))
	// Compile each output regex pattern
	compiledPatterns := make([]*regexp.Regexp, len(outputRegexPatterns))
	for i, pattern := range outputRegexPatterns {
		compiledPatterns[i] = regexp.MustCompile(pattern)
	}

	return &Rule{
		Description:    "Ruby Config",
		RuleID:         "RubyConfig",
		FileExtensions: []string{".rb"},
		Keywords:       []string{"ruby"},
		Regex:          compiledPatterns,
		IgnoreRegex:    ignoreRegex,
	}
}

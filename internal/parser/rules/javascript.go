package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/wadmit/era/internal/types"
)

func JavaScriptConfig(cfg *types.Config) *Rule {
	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`console\.log\((.*?)\)\s*;?`,
		`console\.warn\((.*?)\)\s*;?`,
		`console\.error\((.*?)\)\s*;?`,
		`console\.info\((.*?)\)\s*;?`,
		`console\.debug\((.*?)\)\s*;?`,
	}

	ignoreRegex := regexp.MustCompile(fmt.Sprintf(`(%s)`, strings.Join(ignoreKeywords, "|")))
	// Compile each output regex pattern
	compiledPatterns := make([]*regexp.Regexp, len(outputRegexPatterns))
	for i, pattern := range outputRegexPatterns {
		compiledPatterns[i] = regexp.MustCompile(pattern)
	}

	return &Rule{
		Description:    "JavaScript Config",
		RuleID:         "JavaScriptConfig",
		FileExtensions: []string{".js", ".ts", ".tsx", ".jsx", ".vue", ".svelte"},
		Keywords:       []string{"javascript"},
		Regex:          compiledPatterns,
		IgnoreRegex:    ignoreRegex,
	}
}

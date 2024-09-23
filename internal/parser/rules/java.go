package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/wadmit/era/internal/types"
)

func JavaConfig(cfg *types.Config) *Rule {
	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`System\.out\.print\((.*?)\)\s*;?`,
		`System\.out\.println\((.*?)\)\s*;?`,
		`System\.out\.printf\((.*?)\)\s*;?`,
		`System\.err\.print\((.*?)\)\s*;?`,
		`System\.err\.println\((.*?)\)\s*;?`,
		`System\.err\.printf\((.*?)\)\s*;?`,
		`Logger\.log\((.*?)\)\s*;?`,
		`Logger\.info\((.*?)\)\s*;?`,
		`Logger\.warning\((.*?)\)\s*;?`,
		`Logger\.severe\((.*?)\)\s*;?`,
	}

	ignoreRegex := regexp.MustCompile(fmt.Sprintf(`(%s)`, strings.Join(ignoreKeywords, "|")))
	// Compile each output regex pattern
	compiledPatterns := make([]*regexp.Regexp, len(outputRegexPatterns))
	for i, pattern := range outputRegexPatterns {
		compiledPatterns[i] = regexp.MustCompile(pattern)
	}

	return &Rule{
		Description:    "Java Config",
		RuleID:         "JavaConfig",
		FileExtensions: []string{".java"},
		Keywords:       []string{"java"},
		Regex:          compiledPatterns,
		IgnoreRegex:    ignoreRegex,
	}
}

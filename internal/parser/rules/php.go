package rules

import "github.com/wadmit/era/internal/types"

func PhpConfig(cfg *types.Config) *Rule {

	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`echo\s*(.*?)\s*;`,
		`print\((.*?)\)\s*;`,
		`printf\((.*?)\)\s*;`,
		`error_log\((.*?)\)\s*;`,
	}
	r := Rule{
		Description:    "PHP Config",
		RuleID:         "PhpConfig",
		FileExtensions: []string{".php"},
		Keywords:       []string{"php"},
		Regex:          GenerateCombinedRegex(ignoreKeywords, outputRegexPatterns),
	}
	return &r
}

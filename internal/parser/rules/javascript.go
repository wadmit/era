package rules

import (
	"github.com/wadmit/eradicate/internal/types"
)

func JavaScriptConfig(cfg *types.Config) *Rule {
	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`console\.log\((.*?)\)`,
		`console\.warn\((.*?)\)`,
		`console\.error\((.*?)\)`,
		`console\.info\((.*?)\)`,
		`console\.debug\((.*?)\)`,
	}
	r := Rule{
		Description:    "Javascript Config",
		RuleID:         "JavascriptConfig",
		FileExtensions: []string{".js", ".ts", ".tsx", ".jsx", ".vue", ".svelte"},
		Keywords:       []string{"golang"},
		Regex:          GenerateCombinedRegex(ignoreKeywords, outputRegexPatterns),
	}
	return &r
}

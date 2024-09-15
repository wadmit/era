package rules

import "github.com/wadmit/era/internal/types"

func JavaConfig(cfg *types.Config) *Rule {

	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`System\.out\.print\((.*?)\)`,
		`System\.out\.println\((.*?)\)`,
		`System\.out\.printf\((.*?)\)`,
		`System\.err\.print\((.*?)\)`,
		`System\.err\.println\((.*?)\)`,
		`System\.err\.printf\((.*?)\)`,
		`Logger\.log\((.*?)\)`,
		`Logger\.info\((.*?)\)`,
		`Logger\.warning\((.*?)\)`,
		`Logger\.severe\((.*?)\)`,
	}
	r := Rule{
		Description:    "Java Config",
		RuleID:         "JavaConfig",
		FileExtensions: []string{".java"},
		Keywords:       []string{"java"},
		Regex:          GenerateCombinedRegex(ignoreKeywords, outputRegexPatterns),
	}
	return &r
}

package rules

import (
	"github.com/wadmit/era/internal/types"
)

func PythonConfig(cfg *types.Config) *Rule {

	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`print\((.*?)\)`,
		`sys\.stdout\.write\((.*?)\)`,
		`sys\.stderr\.write\((.*?)\)`,
		`sys\.stdout\.writelines\((.*?)\)`,
		`sys\.stderr\.writelines\((.*?)\)`,
		`logging\.debug\((.*?)\)`,
		`logging\.info\((.*?)\)`,
		`logging\.warning\((.*?)\)`,
		`logging\.error\((.*?)\)`,
		`logging\.critical\((.*?)\)`,
	}
	r := Rule{
		Description:    "Python Config",
		RuleID:         "PythonConfig",
		FileExtensions: []string{".py"},
		Keywords:       []string{"python"},
		Regex:          GenerateCombinedRegex(ignoreKeywords, outputRegexPatterns),
	}
	return &r
}

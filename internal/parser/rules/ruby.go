package rules

import "github.com/wadmit/era/internal/types"

func RubyConfig(cfg *types.Config) *Rule {

	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`puts\s*(.*?)$`,
		`print\s*(.*?)$`,
		`warn\s*(.*?)$`,
		`logger\.debug\((.*?)\)`,
		`logger\.info\((.*?)\)`,
		`logger\.warn\((.*?)\)`,
		`logger\.error\((.*?)\)`,
		`logger\.fatal\((.*?)\)`,
	}
	r := Rule{
		Description:    "Ruby Config",
		RuleID:         "RubyConfig",
		FileExtensions: []string{".rb"},
		Keywords:       []string{"ruby"},
		Regex:          GenerateCombinedRegex(ignoreKeywords, outputRegexPatterns),
	}
	return &r
}

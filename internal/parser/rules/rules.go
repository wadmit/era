package rules

import (
	"regexp"

	"github.com/wadmit/era/internal/types"
)

type Rule struct {
	Description    string
	RuleID         string
	Regex          []*regexp.Regexp
	IgnoreRegex    *regexp.Regexp
	Keywords       []string
	FileExtensions []string
}
type ConfigMap map[string]*Rule

func LoadRules(cfg *types.Config) *ConfigMap {
	configMap := ConfigMap{
		"javascript": JavaScriptConfig(cfg),
		"golang":     GoLangConfig(cfg),
		"python":     PythonConfig(cfg),
		"java":       JavaConfig(cfg),
		"ruby":       RubyConfig(cfg),
		"php":        PhpConfig(cfg),
	}
	return &configMap
}

func GenerateRulesForExtensions(fileExtension string, configMap *ConfigMap) *Rule {
	for _, rule := range *configMap {
		if Contains(rule.FileExtensions, fileExtension) {
			return rule
		}
	}
	return nil
}

// This function will be used to check if a line should be ignored
func ShouldIgnoreLine(line string, ignoreRegex *regexp.Regexp) bool {
	return ignoreRegex.MatchString(line)
}

// checks if a string is in a slice
func Contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

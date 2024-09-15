package rules

import (
	"fmt"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/wadmit/era/internal/types"
)

type Rule struct {
	Description    string
	RuleID         string
	Regex          []*regexp2.Regexp
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

func GenerateCombinedRegex(ignoreKeywords []string, outputRegexPatterns []string) []*regexp2.Regexp {
	// Escape the ignore keywords to be safe in regex
	escapedKeywords := make([]string, len(ignoreKeywords))
	for i, keyword := range ignoreKeywords {
		escapedKeywords[i] = quoteMetaRegex(keyword)
	}

	// Create the ignore regex pattern
	ignorePattern := fmt.Sprintf(`(?!.*//.*(%s))`, strings.Join(escapedKeywords, "|"))

	// Combine the ignore pattern with each output regex pattern
	combinedPatterns := make([]*regexp2.Regexp, len(outputRegexPatterns))
	for i, outputPattern := range outputRegexPatterns {
		pattern := ignorePattern + ".*" + outputPattern
		combinedPatterns[i] = regexp2.MustCompile(pattern, regexp2.RE2)
	}
	return combinedPatterns
}

// quoteMetaRegex escapes regex metacharacters in a string
func quoteMetaRegex(s string) string {
	specialChars := `\.+*?()|[]{}^$`
	result := strings.Builder{}
	for _, char := range s {
		if strings.ContainsRune(specialChars, char) {
			result.WriteRune('\\')
		}
		result.WriteRune(char)
	}
	return result.String()
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

package rules

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/wadmit/era/internal/types"
)

func PythonConfig(cfg *types.Config) *Rule {
	ignoreKeywords := cfg.IgnoreKeyword
	outputRegexPatterns := []string{
		`print\s*\((.*?)\)\s*;?`,
		`sys\.stdout\.write\s*\((.*?)\)\s*;?`,      // Matches sys.stdout.write() with optional spaces and semicolon
		`sys\.stderr\.write\s*\((.*?)\)\s*;?`,      // Matches sys.stderr.write() with optional spaces and semicolon
		`sys\.stdout\.writelines\s*\((.*?)\)\s*;?`, // Matches sys.stdout.writelines() with optional spaces and semicolon
		`sys\.stderr\.writelines\s*\((.*?)\)\s*;?`, // Matches sys.stderr.writelines() with optional spaces and semicolon
		`logging\.debug\s*\((.*?)\)\s*;?`,          // Matches logging.debug() with optional spaces and semicolon
		`logging\.info\s*\((.*?)\)\s*;?`,           // Matches logging.info() with optional spaces and semicolon
		`logging\.warning\s*\((.*?)\)\s*;?`,        // Matches logging.warning() with optional spaces and semicolon
		`logging\.error\s*\((.*?)\)\s*;?`,          // Matches logging.error() with optional spaces and semicolon
		`logging\.critical\s*\((.*?)\)\s*;?`,       // Matches logging.critical() with optional spaces and semicolon
	}

	ignoreRegex := regexp.MustCompile(fmt.Sprintf(`(%s)`, strings.Join(ignoreKeywords, "|")))

	// Compile each output regex pattern
	compiledPatterns := make([]*regexp.Regexp, len(outputRegexPatterns))
	for i, pattern := range outputRegexPatterns {
		compiledPatterns[i] = regexp.MustCompile(pattern)
	}

	return &Rule{
		Description:    "Python Config",
		RuleID:         "PythonConfig",
		FileExtensions: []string{".py"},
		Keywords:       []string{"python"},
		Regex:          compiledPatterns,
		IgnoreRegex:    ignoreRegex,
	}
}

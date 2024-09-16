package config

import "github.com/wadmit/era/internal/types"

func GenerateErdConfig() *types.Config {
	return &types.Config{
		Root:                 ".",
		ReportPath:           "era-reports",
		IgnoreKeyword:        []string{"erd:ignore", "erd:ignoreAll"},
		IgnoreFileExtensions: []string{".exe", ".dll", ".so", ".dylib", ".zip", ".tar", ".gz", ".rar"},
		IgnoreDirs:           []string{"node_modules", "vendor", ".git", ".idea", ".vscode", ".vs", ".hg", ".svn", ".bzr", ".fslckout", "_darcs", "_sgbak", ".bzr", ".bzrignore", ".bzrtags", ".bzrcheckout", ".bzrcommit", ".bzrpush"},
		IgnoreFiles:          []string{},
		ListenType:           "command",
	}
}

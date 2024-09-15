package types

type Config struct {
	Root                 string   `mapstructure:"root"`
	ReportPath           string   `mapstructure:"reportPath"`
	IgnoreKeyword        []string `mapstructure:"IgnoreKeyword"`
	IgnoreFileExtensions []string `mapstructure:"IgnoreFileExtensions"`
	IgnoreDirs           []string `mapstructure:"IgnoreDir"`
	IgnoreFiles          []string `mapstructure:"IgnoreFiles"`
	ListenType           string   `mapstructure:"ListenType"`
}

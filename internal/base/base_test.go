package base

import (
	"testing"

	"github.com/wadmit/eradicate/internal/config"
	"github.com/wadmit/eradicate/internal/parser/rules"
)

func TestBaseTest(t *testing.T) {
	root := "../../examples/languages/javascript"
	cfg, err := config.LoadConfig()
	configMap := rules.LoadRules(cfg)
	if err != nil {
		t.Fatal(err)
	}

	DetectAndChangeFile(root, cfg, configMap)
}

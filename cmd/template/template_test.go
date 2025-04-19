package template

import (
	"testing"
)

func TestTemplateCmd_HasSubcommands(t *testing.T) {
	if len(TemplateCmd.Commands()) == 0 {
		t.Error("expected TemplateCmd to have subcommands")
	}
}

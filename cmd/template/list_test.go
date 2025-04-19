package template

import (
	"bytes"
	"strings"
	"testing"
)

func TestTemplateListCommand_PrintsTemplates(t *testing.T) {
	var buf bytes.Buffer

	listCmd.SetOut(&buf)
	listCmd.SetErr(&buf)
	listCmd.SetArgs([]string{})

	listCmd.Run(listCmd, []string{})

	output := buf.String()

	if !strings.Contains(output, "Available Templates:") {
		t.Errorf("expected 'Available Templates:', got: %s", output)
	}

	if !strings.Contains(output, "default") {
		t.Errorf("expected 'default' template in list, got: %s", output)
	}
}

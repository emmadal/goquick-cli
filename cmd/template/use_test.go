package template

import (
	"bytes"
	"strings"
	"testing"
)

func TestUseCommand_TemplateExists(t *testing.T) {
	var buf bytes.Buffer
	useCmd.SetOut(&buf)
	useCmd.SetErr(&buf)

	// `template use default`
	useCmd.Flags().Parse([]string{})
	useCmd.Run(useCmd, []string{"default"})

	output := buf.String()
	if !strings.Contains(output, "Template 'default' selected successfully") {
		t.Errorf("expected success message, got: %s", output)
	}
}

func TestUseCommand_TemplateDoesNotExist(t *testing.T) {
	var buf bytes.Buffer
	useCmd.SetOut(&buf)
	useCmd.SetErr(&buf)

	useCmd.Flags().Parse([]string{})
	useCmd.Run(useCmd, []string{"not-a-template"})

	output := buf.String()
	if !strings.Contains(output, "Template 'not-a-template' not found") {
		t.Errorf("expected error message for non‑existent template, got: %s", output)
	}
}

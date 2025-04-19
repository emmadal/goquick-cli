package template

import (
	"testing"
)

func TestListTemplates_ReturnsDefaultTemplates(t *testing.T) {
	manager := NewManager()

	templates, err := manager.ListTemplates()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(templates) < 3 {
		t.Errorf("expected at least 3 templates, got %d", len(templates))
	}

	expectedNames := map[string]bool{
		"default": true,
		"api":     true,
		"cli":     true,
	}

	for _, tmpl := range templates {
		if !expectedNames[tmpl.Name] {
			t.Errorf("unexpected template name: %s", tmpl.Name)
		}
	}
}

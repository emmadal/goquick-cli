package ui

import (
	"testing"

	"github.com/goquick-run/cli/internal/template"

	tea "github.com/charmbracelet/bubbletea"
)

func TestUpdate_SelectTemplate_EnterMovesToProjectName(t *testing.T) {
	m := Model{
		templates:     []template.Template{{Name: "default"}},
		selectedIndex: 0,
		step:          stepSelectTemplate,
	}

	msg := tea.KeyMsg{Type: tea.KeyEnter}
	updated, _ := m.Update(msg)

	if updated.(Model).step != stepEnterProjectName {
		t.Errorf("expected stepEnterProjectName, got %v", updated.(Model).step)
	}
}

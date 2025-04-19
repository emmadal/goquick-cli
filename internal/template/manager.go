package template

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Template represents a reusable project/component template
type Template struct {
	Name        string
	Description string
	Category    string
	Source      string
}

// Manager handles available templates (built-in and custom)
type Manager struct {
	templatesDir string
}

// NewManager creates a new template manager and ensures the local template directory exists
func NewManager() *Manager {
	// Get the user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}

	// Define the template directory: ~/.quick/templates
	templatesDir := filepath.Join(home, ".quick", "templates")

	// Create the directory if it doesn't exist
	if _, err := os.Stat(templatesDir); os.IsNotExist(err) {
		_ = os.MkdirAll(templatesDir, 0755)
	}

	return &Manager{
		templatesDir: templatesDir,
	}
}

// ListTemplates returns all available templates (currently only built-in)
func (m *Manager) ListTemplates() ([]Template, error) {
	// Default templates (built-in)
	defaultTemplates := []Template{
		{
			Name:        "default",
			Description: "Default template for Quick projects",
			Category:    "web",
			Source:      "built-in",
		},
		{
			Name:        "api",
			Description: "Template for RESTful APIs",
			Category:    "api",
			Source:      "built-in",
		},
		{
			Name:        "cli",
			Description: "Template for CLI applications",
			Category:    "cli",
			Source:      "built-in",
		},
	}

	// TODO: Load custom templates from the ~/.quick/templates directory

	return defaultTemplates, nil
}

// AddTemplate registers a new template by name, description, category, and source
func (m *Manager) AddTemplate(name, description, category, source string) error {
	// Prevent duplicate names
	exists, _ := m.TemplateExists(name)
	if exists {
		return errors.New("a template with this name already exists")
	}

	// TODO: Implement logic to add a new template:
	// - If source is a Git URL: clone the repo into templatesDir
	// - If source is a local directory: copy the contents to templatesDir

	fmt.Printf("Template '%s' added successfully!\n", name)
	return nil
}

// TemplateExists checks whether a template with the given name exists
func (m *Manager) TemplateExists(name string) (bool, error) {
	templates, err := m.ListTemplates()
	if err != nil {
		return false, err
	}

	for _, tmpl := range templates {
		if tmpl.Name == name {
			return true, nil
		}
	}

	return false, nil
}

// GetTemplate retrieves a template by name
func (m *Manager) GetTemplate(name string) (*Template, error) {
	templates, err := m.ListTemplates()
	if err != nil {
		return nil, err
	}

	for _, tmpl := range templates {
		if tmpl.Name == name {
			return &tmpl, nil
		}
	}

	return nil, fmt.Errorf("template '%s' not found", name)
}

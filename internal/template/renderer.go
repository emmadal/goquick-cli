package template

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// TemplateData holds dynamic values to be used when rendering templates
type TemplateData struct {
	ProjectName    string
	PackageName    string
	ControllerName string
	// TODO: Add additional fields as needed (e.g., module path, license, author)
}

// Renderer handles rendering of templates into real code/files
type Renderer struct {
	manager *Manager
}

// NewRenderer returns a new instance of Renderer
func NewRenderer(manager *Manager) *Renderer {
	return &Renderer{
		manager: manager,
	}
}

// RenderTemplate renders all files in a given template into the target path
func (r *Renderer) RenderTemplate(templateName, targetPath string, data TemplateData) error {
	// Retrieve the template by name
	tmpl, err := r.manager.GetTemplate(templateName)
	if err != nil {
		return err
	}

	fmt.Sprintf("%v", tmpl)
	// TODO: Implement full rendering logic:
	// - Load all template files from the template directory
	// - Recursively render each file (if needed)
	// - Replace variables using text/template
	// - Preserve folder structure when copying

	// Example (future):
	// templateDir := filepath.Join(r.manager.templatesDir, tmpl.Name)
	// filepath.WalkDir(templateDir, ...)

	return nil
}

// RenderFile renders a single template file to a target destination
func (r *Renderer) RenderFile(templatePath, targetPath string, data TemplateData) error {
	// Read the contents of the template file
	content, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	// Parse the template content
	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(content))
	if err != nil {
		return err
	}

	// Execute the template with provided data
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	// Ensure the target directory exists
	targetDir := filepath.Dir(targetPath)
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}

	// Write the rendered output to the target path
	return os.WriteFile(targetPath, buf.Bytes(), 0644)
}

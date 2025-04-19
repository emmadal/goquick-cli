package project

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/goquick-run/cli/internal/template"
)

// ProjectResult holds the result of a project creation
type ProjectResult struct {
	Name string
	Path string
}

// Creator is responsible for creating projects and components
type Creator struct {
	templateManager  *template.Manager
	templateRenderer *template.Renderer
}

// NewCreator returns a new instance of Creator
func NewCreator() *Creator {
	manager := template.NewManager()
	renderer := template.NewRenderer(manager)

	return &Creator{
		templateManager:  manager,
		templateRenderer: renderer,
	}
}

// CreateProject creates a new Quick project
func (c *Creator) CreateProject(name, templateName string) (*ProjectResult, error) {
	// If no name is provided, use the current directory name
	if name == "" {
		currentDir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		name = filepath.Base(currentDir)
	}

	// Retrieve the selected template
	tmpl, err := c.templateManager.GetTemplate(templateName)
	if err != nil {
		return nil, err
	}

	// Set the target project path
	//projectPath := name
	projectPath := filepath.Join(".", name)
	absPath, err := filepath.Abs(projectPath)
	if err != nil {
		return nil, err
	}

	if name != "." {
		// Create the target directory
		if err := os.MkdirAll(projectPath, 0755); err != nil {
			return nil, err
		}
	}

	// Prepare the template data
	data := template.TemplateData{
		ProjectName: name,
		PackageName: name,
	}

	// Render the template
	if err := c.templateRenderer.RenderTemplate(tmpl.Name, projectPath, data); err != nil {
		return nil, err
	}

	return &ProjectResult{
		Name: name,
		Path: absPath,
	}, nil
}

// AddController adds a new controller to the current Quick project
func (c *Creator) AddController(name string) error {
	// Ensure we are inside a Quick project
	if !c.isQuickProject() {
		return errors.New("not inside a Quick project")
	}

	// TODO: Implement controller addition logic
	// - Check if controller already exists
	// - Render the controller template
	// - Add routes or handlers if needed

	return nil
}

// isQuickProject checks whether the current directory is a valid Quick project
func (c *Creator) isQuickProject() bool {
	// Basic check: look for go.mod
	if _, err := os.Stat("go.mod"); err != nil {
		return false
	}

	// TODO: Add more robust validation:
	// - Check for required folder structure (e.g., cmd/, internal/)
	// - Look for .quick.yaml or other Quick-specific files

	return true
}

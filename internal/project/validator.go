package project

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Validator validates project names, controller names, and project structure
type Validator struct{}

// NewValidator returns a new Validator instance
func NewValidator() *Validator {
	return &Validator{}
}

// ValidateProjectName checks if the project name is valid
func (v *Validator) ValidateProjectName(name string) error {
	// Name cannot be empty
	if name == "" {
		return errors.New("project name cannot be empty")
	}

	// Name must contain only allowed characters
	validName := regexp.MustCompile(`^[a-zA-Z0-9_\-]+$`)
	if !validName.MatchString(name) {
		return errors.New("project name contains invalid characters (only letters, numbers, hyphens, and underscores are allowed)")
	}

	// Directory must not already exist (unless current directory is used)
	if name != "." {
		if _, err := os.Stat(name); err == nil {
			return errors.New("a directory with this project name already exists")
		}
	}

	return nil
}

// ValidateControllerName checks if the controller name is valid
func (v *Validator) ValidateControllerName(name string) error {
	if name == "" {
		return errors.New("controller name cannot be empty")
	}

	validName := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !validName.MatchString(name) {
		return errors.New("controller name contains invalid characters (only letters, numbers, and underscores are allowed)")
	}

	// Enforce lowercase naming for consistency
	if strings.ToLower(name) != name {
		return errors.New("controller name must be in lowercase")
	}

	return nil
}

// ValidateProjectStructure checks for basic elements of a valid Quick project
func (v *Validator) ValidateProjectStructure(projectPath string) error {
	// Check if project directory exists
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		return errors.New("project directory does not exist")
	}

	// Check for presence of go.mod file
	if _, err := os.Stat(filepath.Join(projectPath, "go.mod")); os.IsNotExist(err) {
		return errors.New("go.mod file not found in project directory")
	}

	// TODO: Add deeper validation (e.g., check for expected folders like cmd/, internal/, etc.)
	// TODO: Check for .quick.yaml or other metadata in the future

	return nil
}

package project

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateProject_Success(t *testing.T) {
	tmpDir := t.TempDir()
	projectName := "testapp"
	templateName := "default"

	// Change working dir to tmpDir
	oldWD, _ := os.Getwd()
	defer os.Chdir(oldWD)
	_ = os.Chdir(tmpDir)

	creator := NewCreator()

	result, err := creator.CreateProject(projectName, templateName)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expectedPath := filepath.Join(tmpDir, projectName)
	expectedPath, _ = filepath.EvalSymlinks(expectedPath)
	actualPath, _ := filepath.EvalSymlinks(result.Path)
	if actualPath != expectedPath {
		t.Errorf("expected path %s, got %s", expectedPath, actualPath)
	}

	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Errorf("expected directory %s to be created", expectedPath)
	}
}

func TestCreateProject_InvalidTemplate(t *testing.T) {
	tmpDir := t.TempDir()

	_ = os.Chdir(tmpDir)

	creator := NewCreator()

	_, err := creator.CreateProject("myapp", "nonexistent")
	if err == nil {
		t.Fatal("expected error for invalid template name, got nil")
	}
}

func TestCreateProject_UsesCurrentDirIfNameEmpty(t *testing.T) {
	tmpDir := t.TempDir()
	_ = os.Chdir(tmpDir)

	creator := NewCreator()

	result, err := creator.CreateProject("", "default")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expected := filepath.Base(tmpDir)
	if result.Name != expected {
		t.Errorf("expected project name to be %s, got %s", expected, result.Name)
	}
}

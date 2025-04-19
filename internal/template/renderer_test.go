package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRenderFile_CreatesFileWithData(t *testing.T) {
	tmpDir := t.TempDir()

	templateContent := "Hello, {{ .ProjectName }}!"
	templateFile := filepath.Join(tmpDir, "template.txt")
	targetFile := filepath.Join(tmpDir, "result.txt")

	// Write fake template file
	err := os.WriteFile(templateFile, []byte(templateContent), 0644)
	if err != nil {
		t.Fatalf("failed to write template file: %v", err)
	}

	// Setup renderer
	manager := NewManager()
	renderer := NewRenderer(manager)

	data := TemplateData{
		ProjectName: "QuickProject",
	}

	// Run RenderFile
	err = renderer.RenderFile(templateFile, targetFile, data)
	if err != nil {
		t.Fatalf("RenderFile failed: %v", err)
	}

	// Check file exists
	content, err := os.ReadFile(targetFile)
	if err != nil {
		t.Fatalf("failed to read result file: %v", err)
	}

	result := string(content)
	if !strings.Contains(result, "QuickProject") {
		t.Errorf("expected rendered content to include ProjectName, got: %s", result)
	}
}

func TestRenderFile_TemplateNotFound(t *testing.T) {
	tmpDir := t.TempDir()
	renderer := NewRenderer(NewManager())

	err := renderer.RenderFile(filepath.Join(tmpDir, "nonexistent.tmpl"), filepath.Join(tmpDir, "out.txt"), TemplateData{})
	if err == nil {
		t.Error("expected error when template file does not exist, got nil")
	}
}

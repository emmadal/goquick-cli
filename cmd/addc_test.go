package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAddcCommand_CreatesController(t *testing.T) {
	tmp := t.TempDir()

	_ = os.WriteFile(filepath.Join(tmp, "go.mod"), []byte("module fake"), 0644)
	_ = os.Chdir(tmp)

	// Redireciona saída
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{"addc", "user"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	w.Close()
	os.Stdout = oldStdout
	outputBytes, _ := io.ReadAll(r)
	output := string(outputBytes)

	if !strings.Contains(output, "Controller 'user' added successfully!") {
		t.Errorf("expected success message, got: %s", output)
	}
}

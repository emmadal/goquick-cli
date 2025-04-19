package cmd

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestInitCommand_CreatesProject(t *testing.T) {
	tmp := t.TempDir()
	_ = os.Chdir(tmp)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{"init", "mytest"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	w.Close()
	os.Stdout = oldStdout

	outputBytes, _ := io.ReadAll(r)
	output := string(outputBytes)

	if !strings.Contains(output, "\n🎉 Project created successfully! 🎉") {
		t.Errorf("expected success message, got: %s", output)
	}
}

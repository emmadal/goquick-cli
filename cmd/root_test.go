package cmd

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestRootCommand_HelpShownByDefault(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	w.Close()
	os.Stdout = oldStdout
	output, _ := io.ReadAll(r)

	if !strings.Contains(string(output), "Quick CLI is a command-line tool") {
		t.Errorf("expected help output, got: %s", output)
	}
}

func TestRootCommand_VersionFlag(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	rootCmd.SetArgs([]string{"--version"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	w.Close()
	os.Stdout = oldStdout
	output, _ := io.ReadAll(r)

	if !strings.Contains(string(output), "Quick CLI version 1.0.0") {
		t.Errorf("expected version info, got: %s", output)
	}
}

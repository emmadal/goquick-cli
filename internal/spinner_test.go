package internal

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestSpinner(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Create a channel to signal test completion
	done := make(chan bool)

	// Run spinner in a goroutine with small values for testing
	go func() {
		// Use small max value to ensure test completes quickly
		Spinner("Test Loading", 5)
		done <- true
	}()

	// Wait for spinner to finish
	<-done

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Verify output contains expected characters
	if !strings.Contains(output, "-") ||
		!strings.Contains(output, "\\") ||
		!strings.Contains(output, "|") ||
		!strings.Contains(output, "/") ||
		!strings.Contains(output, "Test Loading") {
		t.Errorf("Spinner output didn't contain expected characters. Output: %s", output)
	}
}

func TestSpinnerWithEmptyMessage(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Create a channel to signal test completion
	done := make(chan bool)

	// Run spinner with empty message
	go func() {
		Spinner("", 2)
		done <- true
	}()

	// Wait for spinner to finish
	<-done

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Just verify it completes without error
	var buf bytes.Buffer
	io.Copy(&buf, r)
	// No specific assertions - just testing it doesn't panic
}

package ui

import (
	"os"
	"testing"
)

func TestPrompt_WithInput(t *testing.T) {
	input := "Jeff\n"
	expected := "Jeff"

	// Redireciona stdin
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	w.Write([]byte(input))
	w.Close()
	os.Stdin = r

	result := Prompt("What's your name?", "")
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestPrompt_EmptyInput_ReturnsDefault(t *testing.T) {
	input := "\n"
	defaultValue := "QuickCLI"

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	w.Write([]byte(input))
	w.Close()
	os.Stdin = r

	result := Prompt("What's your project?", defaultValue)
	if result != defaultValue {
		t.Errorf("expected %s, got %s", defaultValue, result)
	}
}

func TestConfirmPrompt_DefaultYes(t *testing.T) {
	input := "\n"

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	r, w, _ := os.Pipe()
	w.Write([]byte(input))
	w.Close()
	os.Stdin = r

	result := ConfirmPrompt("Do you confirm?", true)
	if !result {
		t.Errorf("expected true, got false")
	}
}

func TestConfirmPrompt_No(t *testing.T) {
	input := "n\n"

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	r, w, _ := os.Pipe()
	w.Write([]byte(input))
	w.Close()
	os.Stdin = r

	result := ConfirmPrompt("Do you confirm?", true)
	if result {
		t.Errorf("expected false, got true")
	}
}

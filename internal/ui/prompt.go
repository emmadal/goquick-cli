package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Prompt asks the user for input with an optional default value
func Prompt(message string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)

	if defaultValue != "" {
		fmt.Printf("%s [%s]: ", message, defaultValue)
	} else {
		fmt.Printf("%s: ", message)
	}

	input, err := reader.ReadString('\n')
	if err != nil {
		return defaultValue
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultValue
	}

	return input
}

// ConfirmPrompt asks for user confirmation (yes/no) with default behavior
func ConfirmPrompt(message string, defaultYes bool) bool {
	reader := bufio.NewReader(os.Stdin)

	defaultStr := "n"
	if defaultYes {
		defaultStr = "y"
	}

	fmt.Printf("%s (y/n) [%s]: ", message, defaultStr)

	input, err := reader.ReadString('\n')
	if err != nil {
		return defaultYes
	}

	input = strings.TrimSpace(strings.ToLower(input))
	if input == "" {
		return defaultYes
	}

	return input == "y" || input == "yes"
}

// ShowBootstrapInstructions displays post-creation guidance to the user
func ShowBootstrapInstructions(projectName string) {
	fmt.Println("\n🎉 Project created successfully! 🎉")
	fmt.Println("\nTo get started:")

	fmt.Printf("\n  cd %s\n", projectName)
	fmt.Println("  go mod tidy")
	fmt.Printf("\n  go run cmd/server/main.go\n")

	fmt.Println("\nFor full documentation, visit:")
	fmt.Println("  https://goquick.run/docs")
}

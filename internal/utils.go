package internal

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	MAX_REPLY_LENGTH = 20
	MIN_REPLY_LENGTH = 4
)

// CreateProject creates a new project.
func CreateProject() (string, error) {
	reply := ""
	fmt.Fprint(os.Stdout, "What is your app named? ")
	_, err := fmt.Scanln(&reply)
	if err != nil {
		return "", err
	}
	reply = strings.TrimSpace(reply)

	if len(reply) <= MIN_REPLY_LENGTH {
		return "", fmt.Errorf("project name must be at least %d characters long", MIN_REPLY_LENGTH)
	}

	if len(reply) > MAX_REPLY_LENGTH {
		return "", fmt.Errorf("project name too long. %d characters max allowed", MAX_REPLY_LENGTH)
	}

	if !regexp.MustCompile(`^[a-z]+$`).MatchString(reply) {
		return "", fmt.Errorf("project name can only contain lowercase letters")
	}
	fmt.Fprint(os.Stdout, "\n")

	// Show spinner while creating project
	Spinner("Creating project...", 10)
	return reply, nil
}

// BootstrapInstructions prints instructions for starting the server.
func BootstrapInstructions(projectName string) {
	// Use strings.Builder for efficient string construction
	var sb strings.Builder

	// Pre-allocate an approximate buffer size to avoid reallocations
	sb.Grow(200)

	sb.WriteString(BlackBold)
	sb.WriteString("✅ Your project is ready!\n\n")
	sb.WriteString(ResetColor)
	sb.WriteString("To start the server, navigate to the directory and follow the instructions:\n\n")

	sb.WriteString(BlackBold)
	sb.WriteString("- cd ")
	sb.WriteString(projectName)
	sb.WriteString(ResetColor)
	sb.WriteString("\n")

	sb.WriteString(BlackBold)
	sb.WriteString("- go run main.go")
	sb.WriteString(ResetColor)
	sb.WriteString("\n")

	// Write everything to stdout in a single operation
	fmt.Fprint(os.Stdout, sb.String())
}

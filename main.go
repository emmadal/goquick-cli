package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	BLACK_BOLD = "\033[30;1m"
	RESET      = "\033[0m"
)

type model struct {
	command string
}

// Init handles the initial state of the program.
func (m model) Init() tea.Cmd {
	return nil
}

// Update handles the logic of the program.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, tea.Quit
}

// View handles the display of the program.
func (m model) View() string {
	sb := strings.Builder{}

	// Pre-allocate buffer with approximate size to avoid reallocations
	sb.Grow(50)

	sb.WriteString(`
██████╗ ██╗   ██╗██╗ ██████╗██╗  ██╗
██╔═══██╗██║   ██║██║██╔═══   ██║ ██╔╝
██║   ██║██║   ██║██║██║      █████╔╝
██║▄▄ ██║██║   ██║██║██║      ██╔═██╗
╚██████╔╝╚██████╔╝██║╚██████╔ ██║  ██╗
╚══▀▀═╝  ╚═════╝ ╚═╝ ╚═════╝ ╚═╝  ╚═╝
	`)
	sb.WriteString("\n")
	sb.WriteString("Create fast and reliable projects with quick framework\n")
	sb.WriteString("\n")
	sb.WriteString(BLACK_BOLD + "USAGE:\n" + RESET)
	sb.WriteString("$ quick <command>\n")
	sb.WriteString("\n")
	sb.WriteString(BLACK_BOLD + "COMMANDS:\n" + RESET)
	sb.WriteString("init: \tcreate a new project\n")
	sb.WriteString("addc: \tadd a new controller\n")
	sb.WriteString("\n")
	sb.WriteString(BLACK_BOLD + "EXAMPLES:\n" + RESET)
	sb.WriteString("$ quick init\n")
	sb.WriteString("$ quick addc controller_name\n")
	sb.WriteString("\n")
	sb.WriteString(BLACK_BOLD + "FLAGS:\n" + RESET)
	sb.WriteString("--version: 	show quick version\n")
	sb.WriteString("--help: 	show help for command\n")
	return sb.String()
}

func main() {
	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %s\n", err)
		os.Exit(1)
	}
}

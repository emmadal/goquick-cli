package ui

import (
	"fmt"
	"os"

	"github.com/goquick-run/cli/internal/project"
	"github.com/goquick-run/cli/internal/template"

	tea "github.com/charmbracelet/bubbletea"
)

// Steps in the UI wizard
type step int

const (
	stepSelectTemplate step = iota
	stepEnterProjectName
	stepCreating
	stepDone
)

// Message sent after project creation finishes
type ProjectCreatedMsg struct {
	Success bool
	Error   error
}

// Model holds the Bubble Tea UI state
type Model struct {
	templates     []template.Template
	selectedIndex int
	projectName   string
	step          step
	message       string
}

// RunQuickUI launches the interactive project creation UI
func RunQuickUI() {
	manager := template.NewManager()
	templates, _ := manager.ListTemplates()

	m := Model{
		templates:     templates,
		selectedIndex: 0,
		step:          stepSelectTemplate,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// Init is required by Bubble Tea
func (m Model) Init() tea.Cmd {
	return nil
}

// createProjectAsync wraps the project creation in a tea.Cmd
func createProjectAsync(name, templateName string) tea.Cmd {
	return func() tea.Msg {
		creator := project.NewCreator()
		_, err := creator.CreateProject(name, templateName)
		return ProjectCreatedMsg{
			Success: err == nil,
			Error:   err,
		}
	}
}

// Update reacts to user input and internal messages
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.step {
	case stepSelectTemplate:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				m.step = stepEnterProjectName
			case "up":
				if m.selectedIndex > 0 {
					m.selectedIndex--
				}
			case "down":
				if m.selectedIndex < len(m.templates)-1 {
					m.selectedIndex++
				}
			case "ctrl+c", "q":
				return m, tea.Quit
			}
		}

	case stepEnterProjectName:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				if m.projectName != "" {
					m.step = stepCreating
					return m, createProjectAsync(m.projectName, m.templates[m.selectedIndex].Name)
				}
			case "backspace":
				if len(m.projectName) > 0 {
					m.projectName = m.projectName[:len(m.projectName)-1]
				}
			default:
				m.projectName += msg.String()
			}
		}

	case stepCreating:
		switch msg := msg.(type) {
		case ProjectCreatedMsg:
			m.step = stepDone
			if msg.Success {
				m.message = fmt.Sprintf("✅ Project '%s' created successfully!", m.projectName)
			} else {
				m.message = "❌ Error: " + msg.Error.Error()
			}
			return m, tea.Quit
		}
	}

	return m, nil
}

// View handles the rendering of the UI based on the current step
func (m Model) View() string {
	switch m.step {
	case stepSelectTemplate:
		s := `
  
██████╗ ██╗   ██╗██╗ ██████╗██╗  ██╗
██╔═══██╗██║   ██║██║██╔═══   ██║ ██╔╝
██║   ██║██║   ██║██║██║      █████╔╝
██║▄▄ ██║██║   ██║██║██║      ██╔═██╗
╚██████╔╝╚██████╔╝██║╚██████╔ ██║  ██╗
╚══▀▀═╝  ╚═════╝ ╚═╝ ╚═════╝ ╚═╝  ╚═╝
Welcome to Quick CLI - Let’s build your project 🚀

Select a template:
`
		for i, t := range m.templates {
			cursor := "  "
			if m.selectedIndex == i {
				cursor = "👉"
			}
			s += fmt.Sprintf("%s %s (%s)\n", cursor, t.Name, t.Category)
		}
		s += "\nUse ↑/↓ to navigate, Enter to select."
		return s

	case stepEnterProjectName:
		return fmt.Sprintf("\nEnter project name:\n> %s", m.projectName)

	case stepCreating:
		return fmt.Sprintf("\nCreating project '%s'...\n", m.projectName)

	case stepDone:
		return "\n" + m.message + "\n\n" +
			fmt.Sprintf("cd %s\n", m.projectName) +
			"go mod tidy\n" +
			"go run main.go\n"
	}

	return ""
}

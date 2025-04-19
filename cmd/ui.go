package cmd

import (
	"fmt"

	"github.com/goquick-run/cli/internal/project"
	"github.com/goquick-run/cli/internal/template"
	"github.com/goquick-run/cli/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// Step definitions
type step int

const (
	stepSelectTemplate step = iota
	stepEnterProjectName
	stepCreating
	stepDone
)

// Msg to notify when project creation finishes
type ProjectCreatedMsg struct {
	Success bool
	Error   error
}

type model struct {
	templates     []template.Template
	selectedIndex int
	projectName   string
	step          step
	message       string
}

// quick ui
var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Launch an interactive UI to create a Quick project",
	Run: func(cmd *cobra.Command, args []string) {
		if testMode {
			// Simulates "fake" execution in the test
			return
		}
		ui.RunQuickUI()
	},
}

var testMode bool

func init() {
	uiCmd.Flags().BoolVar(&testMode, "testmode", false, "run UI in test mode (hidden)")
	rootCmd.AddCommand(uiCmd)
}

func (m model) Init() tea.Cmd {
	return nil
}

// Async project creation as tea.Cmd
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

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {
	switch m.step {
	case stepSelectTemplate:
		s := "\nSelect a template:\n\n"
		for i, t := range m.templates {
			cursor := "  "
			if m.selectedIndex == i {
				cursor = "👉"
			}
			s += fmt.Sprintf("%s %s (%s)\n", cursor, t.Name, t.Category)
		}
		s += "\n\nPress Enter to confirm."
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

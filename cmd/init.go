package cmd

import (
	"fmt"
	"os"

	"github.com/goquick-run/cli/internal/project"
	"github.com/goquick-run/cli/internal/ui"

	"github.com/spf13/cobra"
)

var (
	projectName  string
	templateName string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new Quick project",
	Long: `Initialize a new Quick project with the basic structure.
You can specify a name for the project or use the current directory.

Examples:
  quick init my-project
  quick init --template api my-project`,
	Run: func(cmd *cobra.Command, args []string) {
		// If a project name was provided, use it
		if len(args) > 0 {
			projectName = args[0]
		}

		// Start the spinner
		spinner := ui.NewSpinner("Creating project...")
		spinner.Start()

		// Create the project using the provided name and template
		creator := project.NewCreator()
		result, err := creator.CreateProject(projectName, templateName)

		// Stop the spinner
		spinner.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error creating project: %v\n", err)
			os.Exit(1)
		}

		// Show the bootstrap instructions on success
		ui.ShowBootstrapInstructions(result.Name)
	},
}

func init() {
	// Register this command with the root command
	rootCmd.AddCommand(initCmd)

	// Define the --template flag to allow custom project scaffolding
	initCmd.Flags().StringVarP(&templateName, "template", "t", "default", "Template to be used (default, api, web)")
}

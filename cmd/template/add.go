package template

import (
	"fmt"
	"os"

	"github.com/goquick-run/cli/internal/project"
	"github.com/goquick-run/cli/internal/ui"

	"github.com/spf13/cobra"
)

// addTemplateCmd represents the 'quick template add' command
var addcCmd = &cobra.Command{
	Use:   "add [template-name]",
	Short: "Adds a new template to the Quick CLI",
	Long: `Adds a new template to the current Quick CLI environment.
This command must be executed within an existing Quick project.

Examples:
  quick template add api-template
  quick template add ecommerce-template`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templateName := args[0]

		// Start the spinner
		spinner := ui.NewSpinner("Adding template...")
		spinner.Start()

		// Add the template
		creator := project.NewCreator()
		err := creator.AddController(templateName) // TODO: Replace with AddTemplate when available

		// Stop the spinner
		spinner.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error adding template: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Template '%s' added successfully!\n", templateName)
	},
}

func init() {
	TemplateCmd.AddCommand(addcCmd)
}

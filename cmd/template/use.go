package template

import (
	"fmt"

	"github.com/goquick-run/cli/internal/template"
	"github.com/goquick-run/cli/internal/ui"

	"github.com/spf13/cobra"
)

// useTemplateCmd represents the 'quick template use' command
var useCmd = &cobra.Command{
	Use:   "use [template-name]",
	Short: "Use a specific template",
	Long: `Select a specific template to be used when creating a new project or component.

This command verifies the existence of the template and prepares it to be used
with commands such as 'quick init' or 'quick addc'.

Examples:
  quick template use api-rest
  quick template use microservice`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templateName := args[0]

		cmd.SilenceErrors = true
		cmd.SilenceUsage = true

		// Start the spinner
		spinner := ui.NewSpinner("Checking template...")
		spinner.Start()

		// Check if the template exists
		manager := template.NewManager()
		exists, err := manager.TemplateExists(templateName)

		// Stop the spinner
		spinner.Stop()

		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), "❌ Error checking template: %v\n", err)
			return
		}

		if !exists {
			fmt.Fprintf(cmd.ErrOrStderr(), "❌ Template '%s' not found\n", templateName)
			return
		}

		fmt.Fprintf(cmd.OutOrStdout(), "✅ Template '%s' selected successfully!\n", templateName)
		fmt.Fprintf(cmd.OutOrStdout(), "👉 Use it with: quick init --template %s [project-name]\n", templateName)
	},
}

func init() {

	TemplateCmd.AddCommand(useCmd)
}

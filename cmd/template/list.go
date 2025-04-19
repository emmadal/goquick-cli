package template

import (
	"fmt"
	"os"

	"github.com/goquick-run/cli/internal/template"

	"github.com/spf13/cobra"
)

// listTemplatesCmd represents the 'quick template list' command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available templates",
	Long: `Lists all templates available for use with the Quick CLI.

Each template includes a name, description, and category, which help identify its purpose and usage.`,
	Run: func(cmd *cobra.Command, args []string) {
		manager := template.NewManager()
		templates, err := manager.ListTemplates()

		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error listing templates: %v\n", err)
			os.Exit(1)
		}

		cmd.Println("Available Templates:")
		cmd.Println("--------------------")

		for _, tmpl := range templates {
			// TODO: Enhance formatting with colors or table output in the future
			//fmt.Printf("- %s: %s [%s]\n", tmpl.Name, tmpl.Description, tmpl.Category)
			cmd.Printf("- %s: %s [%s]\n", tmpl.Name, tmpl.Description, tmpl.Category)
		}
	},
}

func init() {
	TemplateCmd.AddCommand(listCmd)
}

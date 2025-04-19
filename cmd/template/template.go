package template

import (
	"github.com/spf13/cobra"
)

// TemplateCmd is the root command for managing templates under 'quick template'
var TemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Manage templates for the Quick CLI",
	Long: `The 'template' command allows you to manage reusable project templates.

You can list, add, and use templates to scaffold Quick-based projects and components.

Examples:
  quick template list
  quick template add api-template
  quick template use api-template`,
}

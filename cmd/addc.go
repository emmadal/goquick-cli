package cmd

import (
	"fmt"
	"os"

	"github.com/goquick-run/cli/internal/project"
	"github.com/goquick-run/cli/internal/ui"

	"github.com/spf13/cobra"
)

// addcCmd represents the addc command
var addcCmd = &cobra.Command{
	Use:   "addc [controller-name]",
	Short: "Adds a new controller to the project",
	Long: `Adds a new controller to the current Quick project.
Must be run inside an existing Quick project.

Example:
quick addc user
quick addc product`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		controllerName := args[0]

		// Retrieve the 'force' flag value (not yet used)
		force, _ := cmd.Flags().GetBool("force")

		// TODO: Use the 'force' flag when implementing controller overwrite behavior
		_ = force // placeholder to avoid unused variable warning

		// Start the spinner
		spinner := ui.NewSpinner("Adding controller...")
		spinner.Start()

		// Add the controller
		creator := project.NewCreator()
		err := creator.AddController(controllerName)

		// Stop the spinner
		spinner.Stop()

		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error adding controller: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Controller '%s' added successfully!\n", controllerName)
	},
}

func init() {
	// Register this command with the root command
	rootCmd.AddCommand(addcCmd)

	// Define the '--force' flag to allow controller overwrite in the future
	addcCmd.Flags().BoolP("force", "f", false, "Force overwrite if controller already exists")
}

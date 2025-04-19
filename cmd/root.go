package cmd

import (
	"fmt"
	"os"

	"github.com/goquick-run/cli/cmd/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// CLI version
	Version = "1.0.0"

	// Used for flags
	cfgFile string

	// RootCmd represents the base command
	rootCmd = &cobra.Command{
		Use:   "quick",
		Short: "Quick CLI - Tool to create and manage Quick projects",
		Long: `Quick CLI is a command-line tool to create
and manage projects using the Quick framework.

Create new projects, add controllers, and manage templates
with simple and intuitive commands.`,
		Run: func(cmd *cobra.Command, args []string) {
			versionFlag, _ := cmd.Flags().GetBool("version")
			if versionFlag {
				fmt.Printf("Quick CLI version %s\n", Version)
				return
			}
			// Show help if no subcommands or --version provided
			cmd.Help()
		},
	}
)

// Execute adds all child commands to the root command.
// This is called by main.main().
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	// Initialize config (Viper)
	cobra.OnInitialize(initConfig)

	// Register subcommands
	rootCmd.AddCommand(template.TemplateCmd)

	// Persistent global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (defaults to $HOME/.quick.yaml)")

	// Version flag (global)
	rootCmd.Flags().BoolP("version", "v", false, "Show the Quick CLI version")

	// Bind flag to Viper
	_ = viper.BindPFlag("version", rootCmd.Flags().Lookup("version"))
}

// initConfig sets up the configuration system using Viper.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".quick")
	}

	viper.SetEnvPrefix("QUICK")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using configuration file:", viper.ConfigFileUsed())
	}
}

// Package cmd contains all CLI commands.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	interactive bool
	rootCmdVersion = "v0.1.1"
)

// rootCmd is the root command for dev-helper.
var rootCmd = &cobra.Command{
	Use:   "dev-helper",
	Short: "dev-helper scaffolds backend projects across multiple languages",
	Long: `dev-helper is a CLI tool that generates backend project templates for
Go (Gin/Fiber), Node.js (Express), Python (FastAPI), and Java (Spring Boot).

It supports both interactive TUI mode and non-interactive flag-based mode.`,
	Version: rootCmdVersion,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return nil
	},
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .devhelper.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&interactive, "interactive", "i", true, "use interactive TUI mode")

	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("interactive", rootCmd.PersistentFlags().Lookup("interactive"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("devhelper")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		home, err := os.UserHomeDir()
		if err == nil {
			viper.AddConfigPath(fmt.Sprintf("%s/.config/dev-helper", home))
			viper.AddConfigPath(home)
		}
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// Config loaded successfully (silently)
	}
}

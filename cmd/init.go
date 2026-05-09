// Package cmd contains all CLI commands.
package cmd

import (
	"fmt"
	"strings"

	"github.com/marcosgdz03/dev-helper/core"
	"github.com/marcosgdz03/dev-helper/tui"
	"github.com/spf13/cobra"
)

// initCmd is the "dev-helper init" command.
var initCmd = &cobra.Command{
	Use:   "init [name]",
	Short: "Scaffold a new backend project",
	Long: `Create a new backend project using an interactive wizard or direct flags.

Interactive mode (default):
  dev-helper init

Direct mode:
  dev-helper init --name my-api --lang go --framework gin --output ./my-api`,
	Aliases: []string{"new", "create"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit(args)
	},
}

var (
	initName      string
	initLang      string
	initFramework string
	initOutput    string
)

func init() {
	initCmd.Flags().StringVarP(&initName, "name", "n", "", "project name")
	initCmd.Flags().StringVarP(&initLang, "lang", "l", "", "language (go, node, python, java)")
	initCmd.Flags().StringVarP(&initFramework, "framework", "f", "", "framework (gin, fiber, express, fastapi, springboot)")
	initCmd.Flags().StringVarP(&initOutput, "output", "o", "", "output directory")

	rootCmd.AddCommand(initCmd)
}

func runInit(args []string) error {
	// If interactive mode and no flags provided, launch TUI wizard
	if interactive && initName == "" && initLang == "" && initFramework == "" {
		if len(args) > 0 {
			initName = args[0]
		}
		if initName == "" {
			initName = "my-service"
		}
		return runWizard()
	}

	// Direct mode — validate flags
	if initLang == "" || initFramework == "" {
		return fmt.Errorf("--lang and --framework are required in non-interactive mode")
	}

	// Build scaffolder
	scaff := core.NewScaffolder(cfgFile)

	// Load config for defaults
	if err := scaff.Config.Load(); err != nil {
		// Non-fatal — use defaults
	}

	name := initName
	if name == "" {
		if len(args) > 0 {
			name = args[0]
		} else {
			pc := scaff.Config.GetProjectConfig()
			name = pc.Name
		}
	}
	if name == "" {
		name = "my-service"
	}

	lang := strings.ToLower(initLang)
	framework := strings.ToLower(initFramework)
	output := initOutput

	return scaff.ScaffoldProject(name, lang, framework, output)
}

// runWizard launches the interactive TUI wizard and blocks until the user scaffolds or quits.
func runWizard() error {
	return tui.NewApp().Run()
}

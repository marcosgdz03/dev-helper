package cmd

import (
	"fmt"

	"github.com/dev-helper/dev-helper/core"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate [type] [name]",
	Short: "Generate a component for an existing project",
	Long: `Generate additional components (handlers, models, routes) in a scaffolded project.

Examples:
  dev-helper generate handler user --lang go --framework gin
  dev-helper generate model product --lang python --framework fastapi`,
	Aliases: []string{"gen"},
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGenerate(args)
	},
}

var genLang string
var genFramework string

func init() {
	generateCmd.Flags().StringVarP(&genLang, "lang", "l", "go", "project language")
	generateCmd.Flags().StringVarP(&genFramework, "framework", "f", "gin", "project framework")
	rootCmd.AddCommand(generateCmd)
}

func runGenerate(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("expected at least 2 arguments: type and name (e.g., handler user)")
	}

	compType := args[0]
	compName := args[1]

	// Validate language
	scaff := core.NewScaffolder(cfgFile)
	if !scaff.Config.ValidateLanguage(genLang) {
		return fmt.Errorf("unsupported language: %s", genLang)
	}

	fmt.Printf("Generating %s component: %s (%s/%s)\n", compType, compName, genLang, genFramework)

	// For now, we log this. Full component generation will be added via the plugin.Generate() hook.
	fmt.Printf("Component generation for %s/%s is handled by the %s-%s plugin\n", genLang, genFramework, genLang, genFramework)

	return nil
}

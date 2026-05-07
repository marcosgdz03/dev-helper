package cmd

import (
	"fmt"
	"strings"

	"github.com/dev-helper/dev-helper/core"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Run linters on a generated project",
	Long: `Execute lint commands appropriate for the project language.

Examples:
  dev-helper lint --lang go
  dev-helper lint --lang node --dir ./my-api`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runLint()
	},
}

var lintLang string
var lintDir string

func init() {
	lintCmd.Flags().StringVarP(&lintLang, "lang", "l", "go", "project language")
	lintCmd.Flags().StringVarP(&lintDir, "dir", "d", ".", "project directory")
	rootCmd.AddCommand(lintCmd)
}

func runLint() error {
	lang := strings.ToLower(lintLang)

	scaff := core.NewScaffolder(cfgFile)
	if !scaff.Config.ValidateLanguage(lang) {
		return fmt.Errorf("unsupported language: %s", lang)
	}

	scaff.Executor.WorkingDir = lintDir

	switch lang {
	case "go":
		fmt.Println("Running go vet ./...")
		return scaff.Executor.GoVet()
	case "node":
		if !scaff.Executor.CheckCommand("eslint") {
			fmt.Println("eslint not found — skipping lint")
			return nil
		}
		fmt.Println("Running eslint .")
		_, out, err := scaff.Executor.Run("eslint", ".")
		if err != nil {
			return fmt.Errorf("eslint: %w (%s)", err, out)
		}
		return nil
	case "python":
		if !scaff.Executor.CheckCommand("flake8") {
			fmt.Println("flake8 not found — skipping lint")
			return nil
		}
		fmt.Println("Running flake8 .")
		_, out, err := scaff.Executor.Run("flake8", ".")
		if err != nil {
			return fmt.Errorf("flake8: %w (%s)", err, out)
		}
		return nil
	case "java":
		fmt.Println("Running mvn checkstyle:check")
		_, out, err := scaff.Executor.Run("mvn", "checkstyle:check")
		if err != nil {
			return fmt.Errorf("mvn checkstyle: %w (%s)", err, out)
		}
		return nil
	default:
		return fmt.Errorf("no linter defined for language: %s", lang)
	}
}

package cmd

import (
	"fmt"
	"strings"

	"github.com/marcosgdz03/dev-helper/core"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests for a generated project",
	Long: `Execute test commands appropriate for the project language.

Examples:
  dev-helper test --lang go
  dev-helper test --lang node
  dev-helper test --lang python
  dev-helper test --lang java`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runTest()
	},
}

var testLang string
var testDir string

func init() {
	testCmd.Flags().StringVarP(&testLang, "lang", "l", "go", "project language")
	testCmd.Flags().StringVarP(&testDir, "dir", "d", ".", "project directory")
	rootCmd.AddCommand(testCmd)
}

func runTest() error {
	lang := strings.ToLower(testLang)

	scaff := core.NewScaffolder(cfgFile)
	if !scaff.Config.ValidateLanguage(lang) {
		return fmt.Errorf("unsupported language: %s", lang)
	}

	scaff.Executor.WorkingDir = testDir

	switch lang {
	case "go":
		fmt.Println("Running go test ./...")
		return scaff.Executor.GoTest()
	case "node":
		fmt.Println("Running npm test")
		_, out, err := scaff.Executor.Run("npm", "test")
		if err != nil {
			return fmt.Errorf("npm test: %w (%s)", err, out)
		}
		return nil
	case "python":
		fmt.Println("Running python -m pytest")
		_, out, err := scaff.Executor.Run("python", "-m", "pytest")
		if err != nil {
			return fmt.Errorf("pytest: %w (%s)", err, out)
		}
		return nil
	case "java":
		fmt.Println("Running mvn test")
		_, out, err := scaff.Executor.Run("mvn", "test")
		if err != nil {
			return fmt.Errorf("mvn test: %w (%s)", err, out)
		}
		return nil
	default:
		return fmt.Errorf("no test runner defined for language: %s", lang)
	}
}

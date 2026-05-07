package tui

import (
	"fmt"
	"os"

	"github.com/dev-helper/dev-helper/core"
)

// runScaffold executes the scaffolder using the wizard selections and
// prints a completion signal on success.
func runScaffold(s Selections) error {
	scaff := core.NewScaffolder("")
	if err := scaff.Config.Load(); err != nil {
		// Use defaults
	}

	if err := scaff.ScaffoldProject(s.ProjectName, s.Language, s.Framework, s.OutputDir); err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, "✅ scaffold done")
	return nil
}

// ScaffoldProjectFromSelections is the public entry point used by cmd/init.go
// when non-interactive wiring is needed.
func ScaffoldProjectFromSelections(s Selections) error {
	scaff := core.NewScaffolder("")
	if err := scaff.Config.Load(); err != nil {
		// Use defaults
	}
	return scaff.ScaffoldProject(s.ProjectName, s.Language, s.Framework, s.OutputDir)
}

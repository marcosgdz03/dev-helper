package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// App wraps the Bubble Tea wizard program.
type App struct {
	model Wizard
}

// NewApp creates a new TUI wizard app.
func NewApp() *App {
	return &App{model: NewWizard()}
}

// Run launches the interactive wizard and scaffolds the project on completion.
func (a *App) Run() error {
	p := tea.NewProgram(a.model, tea.WithOutput(os.Stderr))
	result, err := p.Run()
	if err != nil {
		return fmt.Errorf("TUI error: %w", err)
	}

	wiz, ok := result.(Wizard)
	if !ok {
		return fmt.Errorf("unexpected model type from TUI")
	}

	if !wiz.ShouldScaffold {
		return nil
	}

	return runScaffold(wiz.Selections)
}

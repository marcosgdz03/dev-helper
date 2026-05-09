package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// WizardStep represents the current step in the wizard.
type WizardStep int

const (
	StepSelectLanguage WizardStep = iota
	StepSelectFramework
	StepEnterProjectName
	StepSummary
	StepComplete
)

// Item represents a selectable menu item.
type Item struct {
	Title string
	Value string
}

// Wizard is the Bubble Tea model for the scaffolding wizard.
type Wizard struct {
	step             WizardStep
	languages        []Item
	frameworks       []Item
	currentLang      string
	currentFw        string
	projectName      string
	textInput        textinput.Model
	ShouldScaffold   bool
	Selections       Selections
	style           styleModel
	width           int
	height          int
	selectedLangIndex int
	selectedFwIndex   int
}

// Selections holds the final project selections from the wizard.
type Selections struct {
	ProjectName string
	Language   string
	Framework  string
	OutputDir  string
}

// styleModel holds lipgloss styles.
type styleModel struct {
	title      lipgloss.Style
	subtitle   lipgloss.Style
	header     lipgloss.Style
	selected   lipgloss.Style
	deselected lipgloss.Style
	success    lipgloss.Style
	errorStyle lipgloss.Style
	help       lipgloss.Style
	border     lipgloss.Style
}

func NewWizard() Wizard {
	ti := textinput.New()
	ti.Placeholder = "my-service"
	ti.Focus()
	ti.CharLimit = 50

	return Wizard{
		step:            StepSelectLanguage,
		languages:       []Item{
			{Title: "Go", Value: "go"},
			{Title: "Node.js", Value: "node"},
			{Title: "Python", Value: "python"},
			{Title: "Java", Value: "java"},
		},
		textInput:       ti,
		selectedLangIndex: 0,
		selectedFwIndex:    -1,
		style: styleModel{
			title:      lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("105")).MarginBottom(1),
			subtitle:   lipgloss.NewStyle().Foreground(lipgloss.Color("240")).MarginBottom(2),
			header:     lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("208")).Padding(2, 15),
			selected:   lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true),
			deselected: lipgloss.NewStyle().PaddingLeft(1),
			success:    lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Bold(true),
			errorStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true),
			help:       lipgloss.NewStyle().Foreground(lipgloss.Color("241")).MarginTop(2),
			border:     lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).BorderForeground(lipgloss.Color("208")).Padding(1, 2),
		},
	}
}

func (w Wizard) Init() tea.Cmd {
	return textinput.Blink
}

func (w Wizard) renderHeader() string {
	headerText := " ⚡ dev-helper ⚡ "
	return w.style.header.Render(headerText + "\n")
}

func (w Wizard) renderStepIndicator() string {
	stepNum := []string{"1/4", "2/4", "3/4", "4/4", "5/4"}[w.step]
	stepNames := map[WizardStep]string{
		StepSelectLanguage:    "Select Language",
		StepSelectFramework:  "Select Framework",
		StepEnterProjectName: "Enter Project Name",
		StepSummary:          "Summary",
		StepComplete:         "Project Scaffolded",
	}
	step := w.step
	return w.style.subtitle.Render("Step " + stepNum + " — " + stepNames[step] + "\n")
}

func (w Wizard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return w, tea.Quit
		case "enter":
			return w.handleEnter()
		case "up":
			return w.handleUp()
		case "down":
			return w.handleDown()
		}
	case tea.WindowSizeMsg:
		w.width = msg.Width
		w.height = msg.Height
	}

	w.textInput, cmd = w.textInput.Update(msg)
	return w, cmd
}

func (w *Wizard) handleEnter() (tea.Model, tea.Cmd) {
	switch w.step {
	case StepSelectLanguage:
		w.currentLang = w.languages[w.selectedLangIndex].Value
		w.loadFrameworks()
		w.selectedFwIndex = 0
		w.step = StepSelectFramework
	case StepSelectFramework:
		w.currentFw = w.frameworks[w.selectedFwIndex].Value
		w.step = StepEnterProjectName
	case StepEnterProjectName:
		w.projectName = w.textInput.Value()
		w.step = StepSummary
	case StepSummary:
		w.ShouldScaffold = true
		w.Selections.ProjectName = w.projectName
		w.Selections.Language = w.currentLang
		w.Selections.Framework = w.currentFw
		w.step = StepComplete
	case StepComplete:
		return w, tea.Quit
	}
	return w, nil
}

func (w *Wizard) handleUp() (tea.Model, tea.Cmd) {
	switch w.step {
	case StepSelectLanguage:
		if w.selectedLangIndex > 0 {
			w.selectedLangIndex--
		}
	case StepSelectFramework:
		if w.selectedFwIndex > 0 {
			w.selectedFwIndex--
		}
	}
	return w, nil
}

func (w *Wizard) handleDown() (tea.Model, tea.Cmd) {
	switch w.step {
	case StepSelectLanguage:
		if w.selectedLangIndex < len(w.languages)-1 {
			w.selectedLangIndex++
		}
	case StepSelectFramework:
		if w.selectedFwIndex < len(w.frameworks)-1 {
			w.selectedFwIndex++
		}
	}
	return w, nil
}

func (w *Wizard) loadFrameworks() {
	switch w.currentLang {
	case "go":
		w.frameworks = []Item{
			{Title: "Gin", Value: "gin"},
			{Title: "Fiber", Value: "fiber"},
		}
	case "node":
		w.frameworks = []Item{{Title: "Express", Value: "express"}}
	case "python":
		w.frameworks = []Item{{Title: "FastAPI", Value: "fastapi"}}
	case "java":
		w.frameworks = []Item{{Title: "Spring Boot", Value: "springboot"}}
	}
}

func (w Wizard) View() string {
	if w.width == 0 {
		return "Initializing..."
	}

	// Build header and step indicator
	header := w.renderHeader()
	stepIndicator := w.renderStepIndicator()
	separator := "───────────────────────────────────────────────────\n\n"

	// Build the content area based on current step
	var contentArea string
	switch w.step {
	case StepSelectLanguage:
		contentArea = w.renderLanguageSelection()
	case StepSelectFramework:
		contentArea = w.renderFrameworkSelection()
	case StepEnterProjectName:
		contentArea = w.renderNameInput()
	case StepSummary:
		contentArea = w.renderSummary()
	case StepComplete:
		contentArea = w.renderComplete()
	}

	// Combine all elements: header + step indicator + separator + content
	viewContent := header + stepIndicator + separator + contentArea

	// Apply border around the complete view
	return w.style.border.Render(viewContent)
}

func (w Wizard) renderLanguageSelection() string {
	var b string
	b += w.style.title.Render("Choose a Language\n")
	for i, l := range w.languages {
		if i == w.selectedLangIndex {
			b += w.style.selected.Render("▶ " + l.Title + "\n")
		} else {
			b += w.style.deselected.Render("  " + l.Title + "\n")
		}
	}
	b += w.style.help.Render("↑/↓ to select • Enter to continue • q to quit")
	return b
}

func (w Wizard) renderFrameworkSelection() string {
	var b string
	b += w.style.title.Render("Choose a Framework\n")
	for i, f := range w.frameworks {
		if i == w.selectedFwIndex {
			b += w.style.selected.Render("▶ " + f.Title + "\n")
		} else {
			b += w.style.deselected.Render("  " + f.Title + "\n")
		}
	}
	b += w.style.help.Render("↑/↓ to select • Enter to continue • q to quit")
	return b
}

func (w Wizard) renderNameInput() string {
	var b string
	b += w.style.title.Render("Project Name\n")
	b += w.style.subtitle.Render("> ") + w.textInput.View()
	b += w.style.help.Render("Press Enter to continue • q to quit")
	return b
}

func (w Wizard) renderSummary() string {
	var b string
	b += w.style.title.Render("Summary\n\n")
	b += "  Project: " + w.projectName + "\n"
	b += "  Language: " + w.currentLang + "\n"
	b += "  Framework: " + w.currentFw + "\n"
	b += w.style.help.Render("Press Enter to scaffold • q to quit")
	return b
}

func (w Wizard) renderComplete() string {
	return w.style.success.Render("Project scaffolded successfully!\n")
}

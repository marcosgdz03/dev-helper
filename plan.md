# 📋 Build Plan — dev-helper

## Project Summary
Production-grade Go CLI tool with TUI wizard that scaffolds backend service projects across multiple languages (Go, Python, Node.js, Java).

## Architecture
Clean architecture with strict layer separation:
- **cmd/** — Cobra CLI commands
- **core/** — Business logic (Scaffolder, Templater, Executor, Config, Plugin)
- **tui/** — Bubble Tea interactive wizard
- **templates/** — Multi-language project templates
- **plugins/** — Plugin interface definitions

## Execution Strategy
### Phase 1 — Foundation (parallel dispatch)
- Initialize Go module
- Create project directory structure
- Set up dependencies (Cobra, Bubble Tea, Viper)

### Phase 2 — Core Engine
- Implement Templater (text/template wrapper)
- Implement Executor (os/exec wrapper for git, go, npm, pip, docker)
- Implement Config (Viper + YAML)
- Implement Plugin interface + registry
- Implement Scaffolder (orchestrates template + executor)

### Phase 3 — CLI Commands
- Root command
- init (scaffold new project)
- generate (add components to existing project)
- dockerize (create Dockerfile)
- test (run tests)
- lint (run linters)

### Phase 4 — TUI Wizard
- Bubble Tea app initialization
- Language selection widget
- Framework selection widget
- Project summary widget
- Wizard orchestration

### Phase 5 — Templates
- Go Gin + Fiber templates
- Node Express templates
- Python FastAPI templates
- Java Spring Boot templates

### Phase 6 — Validation
- Optimizer review (structure, performance)
- Reviewer validation (code quality, correctness)
- Build verification (`go build`)

### Phase 7 — Git + GitHub
- Initialize repository
- Push to remote

## Dependencies
- github.com/spf13/cobra (CLI framework)
- github.com/spf13/viper (config management)
- github.com/charmbracelet/bubbletea (TUI framework)
- github.com/charmbracelet/lipgloss (TUI styling)
- text/template (stdlib — template engine)
- os/exec (stdlib — process execution)

## Build Command
```bash
go build -o dev-helper .
```

## Risk Profile
- **Low**: Standard Go library usage, well-maintained dependencies
- **Medium**: Template path resolution on Windows
- **Mitigated**: Comprehensive executor error handling

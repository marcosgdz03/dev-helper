# 📋 Tasks — dev-helper

## Task List

### Phase 1 — Foundation
| # | Task | Agent | Dependency |
|---|---|---|---|
| T001 | Initialize Go module (`go mod init`) | code-agent | — |
| T002 | Create directory structure (cmd, core, tui, templates, plugins) | code-agent | — |
| T003 | Download dependencies (Cobra, Bubble Tea, Viper, Lipgloss) | code-agent | T001 |
| T004 | Create .gitignore | code-agent | — |

### Phase 2 — Core Engine
| # | Task | Agent | Dependency |
|---|---|---|---|
| T005 | Implement `core/templater.go` (template engine wrapper) | code-agent | T001 |
| T006 | Implement `core/executor.go` (os/exec wrapper) | code-agent | T001 |
| T007 | Implement `core/config.go` (Viper + YAML) | code-agent | T003 |
| T008 | Implement `plugins/interface.go` (plugin definitions) | code-agent | T001 |
| T009 | Implement `core/plugin.go` (plugin registry) | code-agent | T008 |
| T010 | Implement `core/scaffolder.go` (orchestrator) | code-agent | T005, T006, T007, T009 |

### Phase 3 — CLI Commands
| # | Task | Agent | Dependency |
|---|---|---|---|
| T011 | Implement `cmd/root.go` (Cobra root command + Viper) | code-agent | T007 |
| T012 | Implement `cmd/init.go` (scaffold project) | code-agent | T010 |
| T013 | Implement `cmd/generate.go` (generate components) | code-agent | T010 |
| T014 | Implement `cmd/dockerize.go` (create Dockerfile) | code-agent | T010 |
| T015 | Implement `cmd/test.go` (run tests) | code-agent | T010 |
| T016 | Implement `cmd/lint.go` (run linters) | code-agent | T010 |
| T017 | Create `main.go` (entry point) | code-agent | T011 |

### Phase 4 — TUI Wizard
| # | Task | Agent | Dependency |
|---|---|---|---|
| T018 | Implement `tui/app.go` (Bubble Tea program) | code-agent | T003 |
| T019 | Implement `tui/components/language_select.go` (language picker) | code-agent | T018 |
| T020 | Implement `tui/components/framework_select.go` (framework picker) | code-agent | T019 |
| T021 | Implement `tui/components/summary.go` (project summary) | code-agent | T020 |
| T022 | Implement `tui/wizard.go` (wizard orchestration) | code-agent | T018, T021 |

### Phase 5 — Templates
| # | Task | Agent | Dependency |
|---|---|---|---|
| T023 | Create Go Gin templates | code-agent | T002 |
| T024 | Create Go Fiber templates | code-agent | T002 |
| T025 | Create Node Express templates | code-agent | T002 |
| T026 | Create Python FastAPI templates | code-agent | T002 |
| T027 | Create Java Spring Boot templates | code-agent | T002 |

### Phase 6 — Validation & Build
| # | Task | Agent | Dependency |
|---|---|---|---|
| T028 | Optimizer: structure + performance review | optimizer | All above |
| T029 | Reviewer: strict code review | reviewer | All above |
| T030 | Build binary and verify | code-agent | T028, T029 |

### Phase 7 — CI + Docker
| # | Task | Agent | Dependency |
|---|---|---|---|
| T031 | Create Makefile | code-agent | T030 |
| T032 | Create Dockerfile (distribution) | code-agent | T030 |
| T033 | Create .devhelper.yaml (example config) | code-agent | T007 |

### Phase 8 — Git + GitHub
| # | Task | Agent | Dependency |
|---|---|---|---|
| T034 | git init + initial commit | code-agent | All above |
| T035 | Push to GitHub | code-agent | T034 |

---

## Execution Graph
```
T001 ──→ T003
  │         │
  ├──→ T005 ──┐
  ├──→ T006 ──┤
  ├──→ T007 ──┤  T011 ──→ T012
  │       │       │         ├──→ T013
  │       │       ├──→ T014
  ├──→ T008 ─→ T009 ──→ T010  ├──→ T015
                        │       └──→ T016
                        └──→ T017       │
                                       └──→ T028 ──→ T029 ──→ T030 ──→ T031
                                       T018 ───→ T019 ──→ T020 ──→ T021 ──→ T022
                                       T023 ──+  T024 ──+  T025 ──+  T026 ──+  T027 ──+
```

## Dependencies Matrix
- **Core Engine** must complete before CLI commands
- **CLI Commands** must complete before Validation
- **Validation** must pass before Build
- **Build** must succeed before CI + Docker
- **TUI Wizard** can execute in parallel with CLI commands
- **Templates** can execute in parallel with TUI Wizard

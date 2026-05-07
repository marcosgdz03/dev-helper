# 🤖 Agents — dev-helper

## Agent Registry

| Agent | Role | Tasks | Skill |
|---|---|---|---|
| **code-agent** (primary) | Implementation | T001-T027, T030-T035 | Multiple (see below) |
| **optimizer** | Performance review | T028 | system-architecture-design |
| **reviewer** | Code validation | T029 | testing-universal |

## Skill → Subsystem → Agent Mapping

### Subsystem: Project Foundation
- **Skill:** system-architecture-design
- **Agent:** code-agent
- **Tasks:** T001-T004
- **Input:** project structure specification
- **Output:** initialized Go module, directory tree
- **Success Criteria:** `go mod tidy` succeeds, all directories exist

### Subsystem: Core Engine
- **Skill:** system-architecture-design
- **Agent:** code-agent
- **Tasks:** T005-T010
- **Input:** architecture specification (clean architecture, interface contracts)
- **Output:** 5 production Go files (templater, executor, config, plugin, scaffolder)
- **Success Criteria:** each file compiles, interfaces properly defined, error handling complete

### Subsystem: CLI Commands
- **Skill:** backend-api-universal
- **Agent:** code-agent
- **Tasks:** T011-T017
- **Input:** command specifications (init, generate, dockerize, test, lint)
- **Output:** 7 production Go files (root, init, generate, dockerize, test, lint, main)
- **Success Criteria:** `go build` succeeds, all commands registered, flags functional

### Subsystem: TUI Wizard
- **Skill:** frontend-universal-ui
- **Agent:** code-agent
- **Tasks:** T018-T022
- **Input:** Bubble Tea widget requirements (language select, framework select, summary)
- **Output:** 5 production Go files (app, language_select, framework_select, summary, wizard)
- **Success Criteria:** TUI launches, keyboard navigation works, selections persist

### Subsystem: Template Library
- **Skill:** system-architecture-design
- **Agent:** code-agent
- **Tasks:** T023-T027
- **Input:** 5 technology stacks with file templates
- **Output:** 40+ template files across Go, Node, Python, Java
- **Success Criteria:** each template renders valid project files when executed

### Subsystem: CI + Docker
- **Skill:** devops-docker-ci
- **Agent:** code-agent
- **Tasks:** T031-T032
- **Input:** Go project with Makefile targets
- **Output:** Makefile, Dockerfile
- **Success Criteria:** `make build` succeeds, Dockerfile validates

### Subsystem: Config
- **Skill:** database-schema-design
- **Agent:** code-agent
- **Task:** T033
- **Input:** Viper config schema
- **Output:** .devhelper.yaml example
- **Success Criteria:** Viper loads YAML correctly

### Subsystem: Testing
- **Skill:** testing-universal
- **Agent:** code-agent (test files) + reviewer (validation)
- **Tasks:** T029
- **Input:** complete project
- **Output:** code review report
- **Success Criteria:** no critical issues, PASS verdict

### Subsystem: Performance
- **Skill:** system-architecture-design
- **Agent:** optimizer
- **Tasks:** T028
- **Input:** complete project
- **Output:** optimization report
- **Success Criteria:** acceptable complexity, proper abstractions

## Execution Order
```
Phase 1 (Foundation):
  code-agent ──→ T001, T002, T003, T004

Phase 2 (Core Engine):
  code-agent ──→ T005, T006, T007, T008, T009, T010

Phase 3 (CLI + TUI + Templates — PARALLEL):
  code-agent ──→ T011-T017  (CLI Commands)
  code-agent ──→ T018-T022  (TUI Wizard)
  code-agent ──→ T023-T027  (Templates)

Phase 4 (Validation):
  optimizer ──→ T028        (Structure + Performance)
  reviewer  ──→ T029        (Code Quality)

Phase 5 (Build + Deploy):
  code-agent ──→ T030-T035  (Build, CI, Docker, Git)
```

## Input Contracts
Every agent receives:
1. Task specification (what to build)
2. Input contract (dependencies satisfied)
3. Output contract (expected deliverables)
4. Success criteria (verification conditions)

## Output Contracts
Every agent returns:
1. File list (created/modified files)
2. Build status (compile check)
3. Error list (if any)
4. Completion signal (PASS/FAIL)

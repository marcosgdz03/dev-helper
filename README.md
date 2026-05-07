# dev-helper 🛠️

> Multi-language backend service scaffolding CLI + TUI wizard built in Go.

## Overview

`dev-helper` is a production-grade CLI tool that generates backend service projects across **Go**, **Node.js**, **Python**, and **Java** — complete with an interactive TUI wizard for project configuration.

### Features
- 🚀 **Multi-language scaffolding** — Go (Gin/Fiber), Node.js (Express), Python (FastAPI), Java (Spring Boot)
- 🎨 **Interactive TUI Wizard** — Bubble Tea-driven interactive project setup
- 🐳 **Docker-first** — Auto-generate Dockerfiles and docker-compose files
- 🧪 **Test & Lint** — One-command test and lint execution
- ⚙️ **Configurable** — Viper-powered `.devhelper.yaml` configuration
- 🔌 **Plugin-ready** — Interface-based extensibility for future plugins
- 📦 **Single binary** — Cross-platform Go compilation

## Architecture

Clean architecture with strict separation of concerns:

```
dev-helper/
├── cmd/              # Cobra CLI commands
├── core/             # Business logic engine
├── tui/              # Bubble Tea TUI wizard
│   └── components/   # TUI widgets
├── templates/        # Multi-language project templates
├── plugins/          # Plugin interface definitions
├── main.go           # Entry point
└── .devhelper.yaml   # Configuration file
```

## Quick Start

### Prerequisites
- **Go 1.21+** installed
- Git installed (optional, for initialization)

### Build from Source

```bash
# Clone and enter directory
git clone https://github.com/your-org/dev-helper.git
cd dev-helper

# Download dependencies
go mod download

# Build binary
go build -o dev-helper .

# Run
./dev-helper --help
```

### Using Makefile

```bash
make build      # Build the binary
make clean      # Remove the binary
make test       # Run tests
make install    # Install to GOPATH/bin
```

## CLI Commands

### `dev-helper init`
Scaffold a new backend project.

```bash
# Interactive mode (opens TUI wizard)
./dev-helper init

# Non-interactive mode
./dev-helper init --name my-api --lang go --framework gin
```

**Flags:**

| Flag | Description | Default |
|---|---|---|
| `--name` | Project name | `my-service` |
| `--lang` | Language (go, node, python, java) | required |
| `--framework` | Framework (gin, fiber, express, fastapi, springboot) | required |
| `--output` | Output directory | current directory |
| `--interactive` | Launch TUI wizard | `true` |

### `dev-helper generate`
Generate additional components in an existing project.

```bash
./dev-helper generate --type handler --name user --lang go
```

### `dev-helper dockerize`
Generate a Dockerfile for the project.

```bash
./dev-helper dockerize --lang go --framework gin
```

### `dev-helper test`
Run tests for the generated project.

```bash
./dev-helper test --lang go
```

### `dev-helper lint`
Run linters on the generated project.

```bash
./dev-helper lint --lang go
```

## Configuration

Create `.devhelper.yaml` in your home directory or project root:

```yaml
defaults:
  language: go
  framework: gin
  output_dir: "."
  init_git: true
  generate_docker: true

projects:
  - name: my-api
    language: go
    framework: gin
```

## Interactive TUI Wizard

Run `dev-helper init` without flags to launch the interactive wizard:

```
╭────────────────────────────────────────╮
│           dev-helper wizard            │
│                                        │
│   Select Language:                     │
│   > Go                                 │
│     Node.js                            │
│     Python                             │
│     Java                               │
│                                        │
│   ↑↓ navigate  Enter select            │
╰────────────────────────────────────────╯
```

### Steps
1. Select language
2. Select framework
3. Enter project name
4. Review summary
5. Confirm generation

## Supported Stacks

| Language | Framework | Includes |
|---|---|---|
| Go | Gin | HTTP router, handlers, middleware, tests |
| Go | Fiber | Fast HTTP framework, handlers, middleware |
| Node.js | Express | Express server, routes, middleware |
| Python | FastAPI | Async API, Pydantic models, routers |
| Java | Spring Boot | REST controller, application config |

## Plugin System

dev-helper uses interface-based plugin architecture. Define new plugins by implementing the `Plugin` interface in `core/plugin.go`:

```go
type Plugin interface {
    Name() string
    Version() string
    Generate(config Config) error
}
```

## Workflow

### Agent System
dev-helper is built using an autonomous multi-agent pipeline:

1. **architect** → System design
2. **code-agent** → Implementation
3. **optimizer** → Performance validation
4. **reviewer** → Code review

### Skill Router
Each subsystem is matched to the best-fit skill:

| Subsystem | Skill |
|---|---|
| Core Engine | system-architecture-design |
| CLI Commands | backend-api-universal |
| TUI Wizard | frontend-universal-ui |
| Docker Support | devops-docker-ci |
| Testing | testing-universal |
| Config | database-schema-design |

## Project Structure

See the full file tree below:

```
dev-helper/
├── README.md
├── plan.md
├── tasks.md
├── agents.md
├── Makefile
├── Dockerfile
├── .devhelper.yaml
├── .gitignore
├── go.mod
├── go.sum
├── main.go
├── cmd/
│   ├── root.go
│   ├── init.go
│   ├── generate.go
│   ├── dockerize.go
│   ├── test.go
│   └── lint.go
├── core/
│   ├── scaffolder.go
│   ├── templater.go
│   ├── executor.go
│   ├── config.go
│   └── plugin.go
├── tui/
│   ├── app.go
│   ├── wizard.go
│   └── components/
│       ├── language_select.go
│       ├── framework_select.go
│       └── summary.go
├── templates/
│   ├── go/gin/
│   ├── go/fiber/
│   ├── node/express/
│   ├── python/fastapi/
│   └── java/springboot/
└── plugins/
    └── interface.go
```

## Building and Distributing

### Cross-compile

```bash
# Linux amd64
GOOS=linux GOARCH=amd64 go build -o dev-helper-linux-amd64 .

# macOS arm64
GOOS=darwin GOARCH=arm64 go build -o dev-helper-darwin-arm64 .

# Windows amd64
GOOS=windows GOARCH=amd64 go build -o dev-helper-windows-amd64.exe .
```

### Docker Distribution

```bash
docker build -t dev-helper:latest .
docker run --rm -v ${PWD}:/output dev-helper init --name my-api --lang go --framework gin
```

## License

MIT

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit changes
4. Push to the branch
5. Open a Pull Request

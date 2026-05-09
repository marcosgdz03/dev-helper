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

[![CI](https://github.com/marcosgdz03/dev-helper/actions/workflows/ci.yml/badge.svg)](https://github.com/marcosgdz03/dev-helper/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/marcosgdz03/dev-helper/blob/master/LICENSE.md)
[![Version](https://img.shields.io/github/v/tag/marcosgdz03/dev-helper?label=version)](https://github.com/marcosgdz03/dev-helper/tags)

## Table of Contents
- [Quick Start](#quick-start)
- [CLI Commands](#cli-commands)
- [Configuration](#configuration)
- [Interactive TUI Wizard](#interactive-tui-wizard)
- [Supported Stacks](#supported-stacks)
- [Plugin System](#plugin-system)
- [Project Structure](#project-structure)
- [Building and Distributing](#building-and-distributing)
- [Contributing](#contributing)
- [License](#license)

## Quick Start

### Prerequisites
- **Go 1.21+** installed
- Git installed (optional, for initialization)

### Build from Source

```bash
# Clone and enter directory
git clone https://github.com/marcosgdz03/dev-helper.git
cd dev-helper

# Download dependencies
go mod download

# Build binary
go build -o dev-helper .

# Run
./dev-helper --help
```

### Global Installation (Recommended)

#### Option 1: `go install` (Go 1.18+)
The easiest way to use `dev-helper` from anywhere is to install it globally:

```bash
# Download and install to $GOPATH/bin (or $GOROOT/bin)
go install github.com/dev-helper/dev-helper@latest

# Verify installation
dev-helper --help
```

> **Note:** Ensure `$GOPATH/bin` is in your `PATH` environment variable. By default this is `%USERPROFILE%\go\bin` on Windows or `$HOME/go/bin` on Unix.

#### Option 2: Add binary to PATH manually (Windows)
If you prefer building manually, you can add the `.exe` to your system or user PATH:

```powershell
# 1. Build the binary
go build -o dev-helper.exe .

# 2. Move it to a folder in your PATH (e.g., C:\bin)
Copy-Item .\dev-helper.exe C:\bin\

# 3. Add C:\bin to user PATH (run once)
[user]::SetEnvironmentVariable("Path", "$env:Path;C:\bin", "User")

# 4. Refresh current session
$env:Path = [System.Environment]::GetEnvironmentVariable("Path", "Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path", "User")
```

Now `dev-helper` will be available from any terminal:
```powershell
# From anywhere
dev-helper init --name my-api --lang go --framework gin
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

## Project Structure

```
dev-helper/
├── main.go
├── cmd/              # Cobra CLI commands
├── core/             # Business logic engine
├── tui/              # Bubble Tea TUI wizard
│   └── components/   # TUI widgets
├── templates/        # Multi-language project templates
├── plugins/          # Plugin interface definitions
├── .github/          # GitHub workflows & CODEOWNERS
└── .devhelper.yaml   # Configuration file
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

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License — see the [LICENSE.md](LICENSE.md) file for details.

## Contributors

- Marcos G. D. — Author & Maintainer

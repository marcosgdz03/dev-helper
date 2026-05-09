package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/marcosgdz03/dev-helper/plugins"
)

// Scaffolder orchestrates project generation.
type Scaffolder struct {
	Config    *Config
	Templater *Templater
	Executor  *Executor
	Registry  *plugins.Registry
	Logger    *log.Logger
}

// NewScaffolder builds all subsystems and returns a ready-to-use Scaffolder.
func NewScaffolder(configPath string) *Scaffolder {
	logger := log.New(os.Stderr, "[scaffolder] ", log.LstdFlags)

	cfg := NewConfig()
	if configPath != "" {
		cfg.SetWorkingDir(filepath.Dir(configPath))
	}

	exec := NewExecutor(".")
	tmpl := NewTemplater("templates")
	reg := BuildRegistry(tmpl, exec)

	return &Scaffolder{
		Config:    cfg,
		Templater: tmpl,
		Executor:  exec,
		Registry:  reg,
		Logger:    logger,
	}
}

// ScaffoldProject generates a complete backend project.
func (s *Scaffolder) ScaffoldProject(name, lang, framework, outputDir string) error {
	lang = strings.ToLower(lang)
	framework = strings.ToLower(framework)

	if !s.Config.ValidateLanguage(lang) {
		return fmt.Errorf("unsupported language %q (supported: %v)", lang, s.Config.SupportedLanguages())
	}
	if !s.Config.ValidateFramework(lang, framework) {
		return fmt.Errorf("unsupported framework %q for language %q (supported: %v)",
			framework, lang, s.Config.SupportedFrameworks(lang))
	}

	if outputDir == "" {
		outputDir = filepath.Join(".", name)
	}

	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	s.Executor.WorkingDir = outputDir

	plugin, err := s.Registry.GetByFramework(lang, framework)
	if err != nil {
		return fmt.Errorf("find plugin: %w", err)
	}

	info := plugin.Info()
	s.Logger.Printf("using plugin: %s", info.Name)

	// Generate files
	files, err := plugin.Generate(plugins.PluginConfig{
		ProjectName: name,
		OutputDir:   outputDir,
		Author:      "Developer",
		Version:     "1.0.0",
	})
	if err != nil {
		return fmt.Errorf("generate files: %w", err)
	}
	s.Logger.Printf("generated %d files", len(files))

	// Run post-generation steps
	if err := s.runPostGen(lang, name, outputDir); err != nil {
		s.Logger.Printf("warning: post-gen step failed: %v", err)
	}

	// Git init
	s.Executor.WorkingDir = outputDir
	if err := s.Executor.GitInit(); err != nil {
		s.Logger.Printf("warning: git init failed (this is OK if git is not installed): %v", err)
	}

	s.Logger.Printf("project generated successfully → %s", outputDir)
	return nil
}

// ScaffoldWithConfig loads config and scaffolds.
func (s *Scaffolder) ScaffoldWithConfig() error {
	if err := s.Config.Load(); err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	pc := s.Config.GetProjectConfig()
	if pc.Name == "" {
		return fmt.Errorf("project name is required (set 'name' in config or use --name flag)")
	}

	return s.ScaffoldProject(pc.Name, pc.Language, pc.Framework, pc.OutputDir)
}

func (s *Scaffolder) runPostGen(lang, name, dir string) error {
	s.Executor.WorkingDir = dir
	switch lang {
	case "go":
		if err := s.Executor.GoModInit(name); err != nil {
			return fmt.Errorf("go mod: %w", err)
		}
	case "node":
		return s.Executor.NpmInstall()
	case "python":
		return s.Executor.PipInstall()
	case "java":
		// java: mvn install is optional
		return nil
	}
	return nil
}

// CreateDockerfile generates a Dockerfile for the given language/framework.
func (s *Scaffolder) CreateDockerfile(lang, framework, outputDir string) error {
	var dockerfile string

	switch lang {
	case "go":
		dockerfile = goDockerfile(framework)
	case "node":
		dockerfile = nodeDockerfile()
	case "python":
		dockerfile = pythonDockerfile()
	case "java":
		dockerfile = javaDockerfile()
	default:
		return fmt.Errorf("unsupported language for docker: %s", lang)
	}

	path := filepath.Join(outputDir, "Dockerfile")
	if err := os.WriteFile(path, []byte(dockerfile), 0o644); err != nil {
		return fmt.Errorf("write Dockerfile: %w", err)
	}
	fmt.Printf("Dockerfile written to %s\n", path)
	return nil
}

func goDockerfile(_ string) string {
	return `# ---- Builder ----
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /dev-helper .

# ---- Runner ----
FROM alpine:latest
WORKDIR /app
COPY --from=builder /dev-helper .
EXPOSE 8080
CMD ["/dev-helper"]
`
}

func nodeDockerfile() string {
	return `FROM node:18-alpine
WORKDIR /app
COPY package*.json ./
RUN npm install --production
COPY . .
EXPOSE 3000
CMD ["node", "server.js"]
`
}

func pythonDockerfile() string {
	return `FROM python:3.12-slim
WORKDIR /app
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
EXPOSE 8000
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
`
}

func javaDockerfile() string {
	return `FROM eclipse-temurin:17-jdk-alpine AS builder
WORKDIR /app
COPY pom.xml ./
RUN mvn dependency:go-offline
COPY . .
RUN mvn clean package -DskipTests

FROM eclipse-temurin:17-jre-alpine
WORKDIR /app
COPY --from=builder /app/target/*.jar app.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]
`
}
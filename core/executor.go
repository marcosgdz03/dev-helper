package core

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

// Executor wraps os/exec to run external commands.
type Executor struct {
	Logger     *log.Logger
	WorkingDir string
	Timeout    time.Duration
}

// NewExecutor creates an Executor operating in workingDir.
func NewExecutor(workingDir string) *Executor {
	logger := log.New(os.Stderr, "[executor] ", log.LstdFlags)
	if workingDir == "" {
		workingDir, _ = os.Getwd()
	}
	return &Executor{
		Logger:     logger,
		WorkingDir: workingDir,
		Timeout:    5 * time.Minute,
	}
}

// Run executes command with args, capturing stdout and stderr as strings.
func (e *Executor) Run(command string, args ...string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), e.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Dir = e.WorkingDir

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	e.Logger.Printf("running: %s %s", command, argStr(args))
	if err := cmd.Run(); err != nil {
		return stdout.String(), stderr.String(), fmt.Errorf("%s %s: %w", command, argStr(args), err)
	}
	return stdout.String(), stderr.String(), nil
}

// RunSilent executes command without capturing output.
func (e *Executor) RunSilent(command string, args ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), e.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Dir = e.WorkingDir
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard

	e.Logger.Printf("running (silent): %s %s", command, argStr(args))
	return cmd.Run()
}

// RunInteractive executes command piping output to the terminal.
func (e *Executor) RunInteractive(command string, args ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), e.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Dir = e.WorkingDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	e.Logger.Printf("running (interactive): %s %s", command, argStr(args))
	return cmd.Run()
}

// CheckCommand returns true if the given command is available on PATH.
func (e *Executor) CheckCommand(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err == nil {
		return true
	}
	// Windows: also try with .exe extension
	if _, err := exec.LookPath(cmd + ".exe"); err == nil {
		return true
	}
	return false
}

// GitInit runs "git init" in WorkingDir.
func (e *Executor) GitInit() error {
	if !e.CheckCommand("git") {
		return fmt.Errorf("git not found on PATH")
	}
	_, _, err := e.Run("git", "init")
	return err
}

// GoModInit runs "go mod init <name>" then "go mod tidy".
func (e *Executor) GoModInit(moduleName string) error {
	if !e.CheckCommand("go") {
		return fmt.Errorf("go not found on PATH")
	}
	if _, _, err := e.Run("go", "mod", "init", moduleName); err != nil {
		return fmt.Errorf("go mod init: %w", err)
	}
	if _, out, err := e.Run("go", "mod", "tidy"); err != nil {
		return fmt.Errorf("go mod tidy: %w (%s)", err, out)
	}
	return nil
}

// NpmInstall runs "npm install".
func (e *Executor) NpmInstall() error {
	if !e.CheckCommand("npm") {
		return fmt.Errorf("npm not found on PATH")
	}
	_, out, err := e.Run("npm", "install")
	if err != nil {
		return fmt.Errorf("npm install: %w (%s)", err, out)
	}
	return nil
}

// PipInstall runs "pip install fastapi uvicorn".
func (e *Executor) PipInstall() error {
	cmdName := "pip"
	if !e.CheckCommand("pip") {
		cmdName = "pip3"
	}
	if !e.CheckCommand(cmdName) {
		return fmt.Errorf("pip/pip3 not found on PATH")
	}
	_, out, err := e.Run(cmdName, "install", "fastapi", "uvicorn")
	if err != nil {
		return fmt.Errorf("pip install: %w (%s)", err, out)
	}
	return nil
}

// DockerBuild runs "docker build -t <tag> .".
func (e *Executor) DockerBuild(tag string) error {
	if !e.CheckCommand("docker") {
		return fmt.Errorf("docker not found on PATH")
	}
	_, out, err := e.Run("docker", "build", "-t", tag, ".")
	if err != nil {
		return fmt.Errorf("docker build: %w (%s)", err, out)
	}
	return nil
}

// GoBuild runs "go build -o <name> .".
func (e *Executor) GoBuild(binaryName string) error {
	if !e.CheckCommand("go") {
		return fmt.Errorf("go not found on PATH")
	}
	_, out, err := e.Run("go", "build", "-o", binaryName, ".")
	if err != nil {
		return fmt.Errorf("go build: %w (%s)", err, out)
	}
	return nil
}

// GoTest runs "go test ./...".
func (e *Executor) GoTest() error {
	if !e.CheckCommand("go") {
		return fmt.Errorf("go not found on PATH")
	}
	_, out, err := e.Run("go", "test", "./...")
	if err != nil {
		return fmt.Errorf("go test: %w (%s)", err, out)
	}
	return nil
}

// GoVet runs "go vet ./...".
func (e *Executor) GoVet() error {
	if !e.CheckCommand("go") {
		return fmt.Errorf("go not found on PATH")
	}
	_, out, err := e.Run("go", "vet", "./...")
	if err != nil {
		return fmt.Errorf("go vet: %w (%s)", err, out)
	}
	return nil
}

// argStr joins args with spaces for logging.
func argStr(args []string) string {
	result := ""
	for i, a := range args {
		if i > 0 {
			result += " "
		}
		result += a
	}
	return result
}
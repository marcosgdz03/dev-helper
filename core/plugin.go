package core

import (
	"fmt"
	"os/exec"

	"github.com/dev-helper/dev-helper/plugins"
)

// GoPlugin implements plugins.Plugin for Go stacks (Gin or Fiber).
type GoPlugin struct {
	framework string
	templater *Templater
	executor  *Executor
}

// NewGoPlugin creates a GoPlugin for the given framework.
func NewGoPlugin(framework string, tmpl *Templater, exec *Executor) plugins.Plugin {
	return &GoPlugin{framework: framework, templater: tmpl, executor: exec}
}

func (p *GoPlugin) Info() plugins.PluginInfo {
	return plugins.PluginInfo{
		Name:        fmt.Sprintf("go-%s", p.framework),
		Version:     "1.0.0",
		Description: fmt.Sprintf("Scaffold a Go backend with %s", p.framework),
		Author:      "dev-helper",
		Language:    "go",
		Framework:   p.framework,
		Priority:    1,
	}
}

func (p *GoPlugin) Generate(cfg plugins.PluginConfig) ([]string, error) {
	tmplPath := fmt.Sprintf("templates/go/%s", p.framework)
	files, err := p.templater.ExecuteAllTemplates(tmplPath, cfg.OutputDir, map[string]interface{}{
		"ProjectName": cfg.ProjectName,
		"Author":      cfg.Author,
		"Version":     cfg.Version,
		"Year":        2026,
		"Framework":   p.framework,
	})
	if err != nil {
		return nil, fmt.Errorf("generate go project: %w", err)
	}
	return files, nil
}

func (p *GoPlugin) Validate() error {
	if _, err := exec.LookPath("go"); err != nil {
		return fmt.Errorf("go not found on PATH (required for %s)", p.framework)
	}
	return nil
}

func (p *GoPlugin) Deps() []string {
	return []string{"go mod tidy"}
}

// NodePlugin implements plugins.Plugin for Node.js + Express.
type NodePlugin struct {
	templater *Templater
	executor  *Executor
}

func NewNodePlugin(tmpl *Templater, exec *Executor) plugins.Plugin {
	return &NodePlugin{templater: tmpl, executor: exec}
}

func (p *NodePlugin) Info() plugins.PluginInfo {
	return plugins.PluginInfo{
		Name:        "node-express",
		Version:     "1.0.0",
		Description: "Scaffold a Node.js backend with Express",
		Author:      "dev-helper",
		Language:    "node",
		Framework:   "express",
		Priority:    2,
	}
}

func (p *NodePlugin) Generate(cfg plugins.PluginConfig) ([]string, error) {
	files, err := p.templater.ExecuteAllTemplates("templates/node/express", cfg.OutputDir, map[string]interface{}{
		"ProjectName": cfg.ProjectName,
		"Author":      cfg.Author,
		"Version":     cfg.Version,
		"Year":        2026,
	})
	if err != nil {
		return nil, fmt.Errorf("generate node project: %w", err)
	}
	return files, nil
}

func (p *NodePlugin) Validate() error {
	if _, err := exec.LookPath("node"); err != nil {
		return fmt.Errorf("node not found on PATH")
	}
	return nil
}

func (p *NodePlugin) Deps() []string {
	return []string{"npm install"}
}

// PythonPlugin implements plugins.Plugin for Python + FastAPI.
type PythonPlugin struct {
	templater *Templater
	executor  *Executor
}

func NewPythonPlugin(tmpl *Templater, exec *Executor) plugins.Plugin {
	return &PythonPlugin{templater: tmpl, executor: exec}
}

func (p *PythonPlugin) Info() plugins.PluginInfo {
	return plugins.PluginInfo{
		Name:        "python-fastapi",
		Version:     "1.0.0",
		Description: "Scaffold a Python backend with FastAPI",
		Author:      "dev-helper",
		Language:    "python",
		Framework:   "fastapi",
		Priority:    3,
	}
}

func (p *PythonPlugin) Generate(cfg plugins.PluginConfig) ([]string, error) {
	files, err := p.templater.ExecuteAllTemplates("templates/python/fastapi", cfg.OutputDir, map[string]interface{}{
		"ProjectName": cfg.ProjectName,
		"Author":      cfg.Author,
		"Version":     cfg.Version,
		"Year":        2026,
	})
	if err != nil {
		return nil, fmt.Errorf("generate python project: %w", err)
	}
	return files, nil
}

func (p *PythonPlugin) Validate() error {
	if _, err := exec.LookPath("python3"); err != nil {
		if _, err = exec.LookPath("python"); err != nil {
			return fmt.Errorf("python3 or python not found on PATH")
		}
	}
	return nil
}

func (p *PythonPlugin) Deps() []string {
	return []string{"pip install fastapi uvicorn"}
}

// JavaPlugin implements plugins.Plugin for Java + Spring Boot.
type JavaPlugin struct {
	templater *Templater
	executor  *Executor
}

func NewJavaPlugin(tmpl *Templater, exec *Executor) plugins.Plugin {
	return &JavaPlugin{templater: tmpl, executor: exec}
}

func (p *JavaPlugin) Info() plugins.PluginInfo {
	return plugins.PluginInfo{
		Name:        "java-springboot",
		Version:     "1.0.0",
		Description: "Scaffold a Java backend with Spring Boot",
		Author:      "dev-helper",
		Language:    "java",
		Framework:   "springboot",
		Priority:    4,
	}
}

func (p *JavaPlugin) Generate(cfg plugins.PluginConfig) ([]string, error) {
	files, err := p.templater.ExecuteAllTemplates("templates/java/springboot", cfg.OutputDir, map[string]interface{}{
		"ProjectName": cfg.ProjectName,
		"Author":      cfg.Author,
		"Version":     cfg.Version,
		"Year":        2026,
	})
	if err != nil {
		return nil, fmt.Errorf("generate java project: %w", err)
	}
	return files, nil
}

func (p *JavaPlugin) Validate() error {
	if _, err := exec.LookPath("java"); err != nil {
		return fmt.Errorf("java not found on PATH")
	}
	if _, err := exec.LookPath("mvn"); err != nil {
		return fmt.Errorf("mvn not found on PATH")
	}
	return nil
}

func (p *JavaPlugin) Deps() []string {
	return []string{"mvn clean install"}
}

// BuildRegistry assembles the plugin registry with all built-in plugins.
func BuildRegistry(tmpl *Templater, exec *Executor) *plugins.Registry {
	reg := plugins.NewRegistry()
	reg.Register(NewGoPlugin("gin", tmpl, exec))
	reg.Register(NewGoPlugin("fiber", tmpl, exec))
	reg.Register(NewNodePlugin(tmpl, exec))
	reg.Register(NewPythonPlugin(tmpl, exec))
	reg.Register(NewJavaPlugin(tmpl, exec))
	return reg
}
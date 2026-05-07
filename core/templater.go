// Package core provides the dev-helper core engine.
package core

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

// Templater wraps text/template with utility functions for scaffolding.
type Templater struct {
	TemplateDir string
	Logger      *log.Logger
	funcMap     template.FuncMap
}

// NewTemplater creates a Templater that reads templates from templateDir.
func NewTemplater(templateDir string) *Templater {
	if templateDir == "" {
		templateDir = "templates"
	}
	logger := log.New(os.Stderr, "[templater] ", log.LstdFlags)

	return &Templater{
		TemplateDir: templateDir,
		Logger:      logger,
		funcMap: template.FuncMap{
			"Lower":       strings.ToLower,
			"Upper":       strings.ToUpper,
			"Replace":     strings.Replace,
			"ReplaceAll":  strings.ReplaceAll,
			"ToCamel":     toCamel,
			"ToSnake":     toSnake,
			"CurrentYear": func() int { return time.Now().Year() },
		},
	}
}

func toCamel(s string) string {
	parts := strings.Split(s, "-")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, "")
}

func toSnake(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, "-", "_"))
}

// ParseFile parses a single .tmpl file and returns the parsed template.
func (t *Templater) ParseFile(path string) (*template.Template, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read template %s: %w", path, err)
	}

	name := filepath.Base(path)
	tmpl, err := template.New(name).Funcs(t.funcMap).Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("parse template %s: %w", path, err)
	}

	return tmpl, nil
}

// ExecuteTemplate executes tmpl with data and writes the result to outputPath.
// It creates parent directories as needed.
func (t *Templater) ExecuteTemplate(tmpl *template.Template, outputPath string, data map[string]interface{}) error {
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create dir %s: %w", dir, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	if err := os.WriteFile(outputPath, buf.Bytes(), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", outputPath, err)
	}

	t.Logger.Printf("generated %s", outputPath)
	return nil
}

// ExecuteAllTemplates walks rootDir, executes every .tmpl file found,
// and writes results to the corresponding path under baseOutput.
func (t *Templater) ExecuteAllTemplates(rootDir, baseOutput string, data map[string]interface{}) ([]string, error) {
	var generated []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".tmpl") {
			return nil
		}

		rel, err := filepath.Rel(rootDir, path)
		if err != nil {
			return fmt.Errorf("rel path %s: %w", path, err)
		}

		output := filepath.Join(baseOutput, strings.TrimSuffix(rel, ".tmpl"))
		tmpl, err := t.ParseFile(path)
		if err != nil {
			return err
		}

		if err := t.ExecuteTemplate(tmpl, output, data); err != nil {
			return err
		}
		generated = append(generated, output)
		return nil
	})
	return generated, err
}

// GetTemplateDir returns the configured template directory.
func (t *Templater) GetTemplateDir() string {
	return t.TemplateDir
}
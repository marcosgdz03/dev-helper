package core

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestNewTemplaterDefaults(t *testing.T) {
	tmpl := NewTemplater("")
	if tmpl.GetTemplateDir() != "templates" {
		t.Errorf("NewTemplater(\"\") template dir = %q, want templates", tmpl.GetTemplateDir())
	}
}

func TestFuncRegistry(t *testing.T) {
	tmpl := NewTemplater("")
	if len(tmpl.funcMap) == 0 {
		t.Fatal("NewTemplater did not register template functions")
	}
}

func TestToCamel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"my-api", "MyApi"},
		{"user-management", "UserManagement"},
		{"api-server", "ApiServer"},
		{"SimpleGo", "SimpleGo"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(st *testing.T) {
			got := toCamel(tt.input)
			if got != tt.expected {
				st.Errorf("toCamel(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToSnake(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"my-api", "my_api"},
		{"user_management", "user_management"},
		{"simple", "simple"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(st *testing.T) {
			got := toSnake(tt.input)
			if got != tt.expected {
				st.Errorf("toSnake(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestParseFileAndExecute(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "devhelper-test")
	defer os.RemoveAll(tmpDir)

	tmplPath := filepath.Join(tmpDir, "test.tmpl")
	tmplContent := `Hello {{.Name}}!
API: {{.API | ToCamel}}
Service: {{.API | ToSnake}}
`
	os.WriteFile(tmplPath, []byte(tmplContent), 0644)

	tmpl := NewTemplater("")
	parsed, err := tmpl.ParseFile(tmplPath)
	if err != nil {
		t.Fatalf("ParseFile(%q): %v", tmplPath, err)
	}

	data := map[string]interface{}{
		"Name":  "World",
		"API":  "my-rest-api",
	}

	var buf bytes.Buffer
	err = parsed.Execute(&buf, data)
	if err != nil {
		t.Fatalf("Execute: %v", err)
	}

	expected := "Hello World!\nAPI: MyRestApi\nService: my_rest_api\n"
	if buf.String() != expected {
		t.Errorf("Execute result mismatch:\ngot: %q\nwant: %q", buf.String(), expected)
	}
}

func TestExecuteTemplateCreatesOutputFile(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "devhelper-test")
	defer os.RemoveAll(tmpDir)

	inputPath := filepath.Join(tmpDir, "input.tmpl")
	outputPath := filepath.Join(tmpDir, "output.txt")

	os.WriteFile(inputPath, []byte("Template {{.Value}}"), 0644)

	tmpl := NewTemplater("")
	parsed, _ := tmpl.ParseFile(inputPath)
	tmpl.ExecuteTemplate(parsed, outputPath, map[string]interface{}{"Value": "Hello"})

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Error("ExecuteTemplate did not create output file")
	}
}

func TestExecuteAllTemplatesWalksDirectory(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "devhelper-test")
	defer os.RemoveAll(tmpDir)

	projDir := filepath.Join(tmpDir, "project")
	handlerDir := filepath.Join(projDir, "handlers")
	os.MkdirAll(handlerDir, 0755)

	mainTmpl := `// main.go
package main
func main() {}`
	_ = os.WriteFile(filepath.Join(projDir, "main.tmpl"), []byte(mainTmpl), 0644)

	handlerTmpl := `// {{.Path}}.go
package handlers
func Main() {}`
	os.WriteFile(filepath.Join(handlerDir, "handler.tmpl"), []byte(handlerTmpl), 0644)

	outputDir := filepath.Join(tmpDir, "output")

	tmpl := NewTemplater("")
	files, err := tmpl.ExecuteAllTemplates(projDir, outputDir, map[string]interface{}{"Path": "test"})
	if err != nil {
		t.Fatalf("ExecuteAllTemplates: %v", err)
	}

	if len(files) == 0 {
		t.Error("ExecuteAllTemplates generated no files")
	}
}

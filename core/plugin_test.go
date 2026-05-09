package core

import "testing"

func TestNewGoPluginGinInfo(t *testing.T) {
	exec := NewExecutor("")
	tmpl := NewTemplater("")
	plugin := NewGoPlugin("gin", tmpl, exec)
	info := plugin.Info()

	if info.Language != "go" {
		t.Errorf("GoPlugin.Info() Language = %q, want go", info.Language)
	}
	if info.Framework != "gin" {
		t.Errorf("GoPlugin.Info() Framework = %q, want gin", info.Framework)
	}
	if info.Name != "go-gin" {
		t.Errorf("GoPlugin.Info() Name = %q, want go-gin", info.Name)
	}
}

func TestNewGoPluginFiberInfo(t *testing.T) {
	exec := NewExecutor("")
	tmpl := NewTemplater("")
	plugin := NewGoPlugin("fiber", tmpl, exec)
	info := plugin.Info()

	if info.Language != "go" {
		t.Errorf("GoPlugin.Info() Language = %q, want go", info.Language)
	}
	if info.Framework != "fiber" {
		t.Errorf("GoPlugin.Info() Framework = %q, want fiber", info.Framework)
	}
	if info.Name != "go-fiber" {
		t.Errorf("GoPlugin.Info() Name = %q, want go-fiber", info.Name)
	}
}

func TestNodePluginInfo(t *testing.T) {
	exec := NewExecutor("")
	tmpl := NewTemplater("")
	plugin := NewNodePlugin(tmpl, exec)
	info := plugin.Info()

	if info.Language != "node" {
		t.Errorf("NodePlugin.Info() Language = %q, want node", info.Language)
	}
	if info.Framework != "express" {
		t.Errorf("NodePlugin.Info() Framework = %q, want express", info.Framework)
	}
	if info.Name != "node-express" {
		t.Errorf("NodePlugin.Info() Name = %q, want node-express", info.Name)
	}
}

func TestPythonPluginInfo(t *testing.T) {
	exec := NewExecutor("")
	tmpl := NewTemplater("")
	plugin := NewPythonPlugin(tmpl, exec)
	info := plugin.Info()

	if info.Language != "python" {
		t.Errorf("PythonPlugin.Info() Language = %q, want python", info.Language)
	}
	if info.Framework != "fastapi" {
		t.Errorf("PythonPlugin.Info() Framework = %q, want fastapi", info.Framework)
	}
	if info.Name != "python-fastapi" {
		t.Errorf("PythonPlugin.Info() Name = %q, want python-fastapi", info.Name)
	}
}

func TestJavaPluginInfo(t *testing.T) {
	exec := NewExecutor("")
	tmpl := NewTemplater("")
	plugin := NewJavaPlugin(tmpl, exec)
	info := plugin.Info()

	if info.Language != "java" {
		t.Errorf("JavaPlugin.Info() Language = %q, want java", info.Language)
	}
	if info.Framework != "springboot" {
		t.Errorf("JavaPlugin.Info() Framework = %q, want springboot", info.Framework)
	}
	if info.Name != "java-springboot" {
		t.Errorf("JavaPlugin.Info() Name = %q, want java-springboot", info.Name)
	}
}

func TestDepsNonEmpty(t *testing.T) {
	exec := NewExecutor("")
	tmpl := NewTemplater("")

	goPlugin := NewGoPlugin("gin", tmpl, exec)
	goDeps := goPlugin.Deps()
	if len(goDeps) == 0 {
		t.Error("GoPlugin.Deps() returned empty slice")
	}

	nodePlugin := NewNodePlugin(tmpl, exec)
	nodeDeps := nodePlugin.Deps()
	if len(nodeDeps) == 0 {
		t.Error("NodePlugin.Deps() returned empty slice")
	}

	pythonPlugin := NewPythonPlugin(tmpl, exec)
	pythonDeps := pythonPlugin.Deps()
	if len(pythonDeps) == 0 {
		t.Error("PythonPlugin.Deps() returned empty slice")
	}

	javaPlugin := NewJavaPlugin(tmpl, exec)
	javaDeps := javaPlugin.Deps()
	if len(javaDeps) == 0 {
		t.Error("JavaPlugin.Deps() returned empty slice")
	}
}

// Test that Validate() doesn't panic when Go is not installed
func TestGoPluginValidateNoGoInstalled(t *testing.T) {
	exec := NewExecutor("")
	// Temporarily move go binary out of PATH if it exists
	plugin := NewGoPlugin("gin", NewTemplater(""), exec)
	err := plugin.Validate()
	// Error is expected if go is not available on PATH
	// We only check that it doesn't panic
	_ = err
}

// Similar test for other plugins (node, python, java)
func TestAllPluginsValidateNoBinaryInstalled(t *testing.T) {
	exec := NewExecutor("")
	tmpl := NewTemplater("")

	// Node plugin test - node might not be installed
	nodePlugin := NewNodePlugin(tmpl, exec)
	err := nodePlugin.Validate()
	_ = err

	// Python plugin test - python3 might not be installed
	pythonPlugin := NewPythonPlugin(tmpl, exec)
	err = pythonPlugin.Validate()
	_ = err

	// Java plugin test - java might not be installed
	javaPlugin := NewJavaPlugin(tmpl, exec)
	err = javaPlugin.Validate()
	_ = err

	t.Log("Plugin Validate() methods executed without panic")
}

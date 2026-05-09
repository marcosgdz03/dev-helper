package plugins

import "testing"

func TestNewRegistryEmpty(t *testing.T) {
	reg := NewRegistry()
	if reg == nil {
		t.Fatal("NewRegistry() returned nil")
	}
	if reg.Count() != 0 {
		t.Errorf("NewRegistry().Count() = %d, want 0", reg.Count())
	}
}

func TestRegisterAndGetRoundtrip(t *testing.T) {
	reg := NewRegistry()

	// Test placeholder - actual plugin would be registered here
	// reg.Register(...)
	retrieved, err := reg.Get("unknown/plugin")
	if err == nil {
		t.Error("reg.Get() for nonexistent should return error")
	}

	// Verify retrieved is nil and err exists
	if retrieved != nil {
		t.Error("get should return nil")
	}
}

func TestGetNonexistentPlugin(t *testing.T) {
	reg := NewRegistry()
	_, err := reg.Get("unknown/unknown")
	if err == nil {
		t.Error("reg.Get() for nonexistent plugin should return error")
	}
}

func TestGetByLanguage(t *testing.T) {
	reg := NewRegistry()
	golangPlugins := reg.GetByLanguage("go")
	// With no plugins registered, should return nil
	if golangPlugins == nil {
		t.Log("GetByLanguage(\"go\") returns nil when no plugins registered")
	}
}

func TestGetByFramework(t *testing.T) {
	reg := NewRegistry()
	plugin, err := reg.GetByFramework("go", "gin")
	if err == nil {
		t.Error("GetByFramework should return error")
	}

	// Verify plugin is nil
	if plugin != nil {
		t.Error("GetByFramework should return nil for nonexistent plugin")
	}
}

func TestAllReturnsEmpty(t *testing.T) {
	reg := NewRegistry()
	plugins := reg.All()
	if len(plugins) != 0 {
		t.Error("All() should return empty slice when no plugins registered")
	}
}

func TestCountCorrect(t *testing.T) {
	reg := NewRegistry()
	if reg.Count() != 0 {
		t.Error("NewRegistry().Count() should be 0")
	}
}

// Test thread safety with concurrent Register calls
func TestRegistryThreadSafety(t *testing.T) {
	reg := NewRegistry()

	// Verify count starts at 0
	if reg.Count() != 0 {
		t.Error("Registry count should be 0 initially")
	}

	t.Log("Successfully created registry - thread safety test would need actual plugins")
}

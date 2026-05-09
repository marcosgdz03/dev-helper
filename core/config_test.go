package core

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig()
	if cfg == nil {
		t.Fatal("NewConfig() returned nil")
	}
	if cfg.viper == nil {
		t.Fatal("NewConfig() did not create Viper instance")
	}
}

func TestValidateLanguageValid(t *testing.T) {
	tests := []struct {
		lang     string
		wantValid bool
	}{
		{"go", true},
		{"node", true},
		{"python", true},
		{"java", true},
	}

	cfg := NewConfig()
	for _, tt := range tests {
		t.Run(tt.lang, func(st *testing.T) {
			got := cfg.ValidateLanguage(tt.lang)
			if got != tt.wantValid {
				st.Errorf("ValidateLanguage(%q) = %v, want %v", tt.lang, got, tt.wantValid)
			}
		})
	}
}

func TestValidateLanguageInvalid(t *testing.T) {
	cfg := NewConfig()
	invalidLangs := []string{"rust", "csharp", "php", "typescript"}
	for _, lang := range invalidLangs {
		if cfg.ValidateLanguage(lang) {
			t.Errorf("ValidateLanguage(%q) should return false for invalid language", lang)
		}
	}
}

func TestValidateFrameworkValid(t *testing.T) {
	tests := []struct {
		lang       string
		framework string
		wantValid  bool
	}{
		{"go", "gin", true},
		{"go", "fiber", true},
		{"node", "express", true},
		{"python", "fastapi", true},
		{"java", "springboot", true},
	}

	cfg := NewConfig()
	for _, tt := range tests {
		t.Run(tt.lang+"/"+tt.framework, func(st *testing.T) {
			got := cfg.ValidateFramework(tt.lang, tt.framework)
			if got != tt.wantValid {
				st.Errorf("ValidateFramework(%q, %q) = %v, want %v", tt.lang, tt.framework, got, tt.wantValid)
			}
		})
	}
}

func TestValidateFrameworkInvalid(t *testing.T) {
	cfg := NewConfig()
	invalidCombos := []struct {
		lang       string
		framework string
	}{
		{"go", "express"},
		{"node", "fastapi"},
		{"python", "springboot"},
		{"java", "gin"},
	}

	for _, tt := range invalidCombos {
		t.Run(tt.lang+"/"+tt.framework, func(st *testing.T) {
			if cfg.ValidateFramework(tt.lang, tt.framework) {
				st.Errorf("ValidateFramework(%q, %q) should return false for invalid combo", tt.lang, tt.framework)
			}
		})
	}
}

func TestSupportedLanguages(t *testing.T) {
	cfg := NewConfig()
	langs := cfg.SupportedLanguages()

	if len(langs) != 4 {
		t.Errorf("SupportedLanguages() returned %d languages, want 4", len(langs))
	}

	expected := []string{"go", "node", "python", "java"}
	for i, lang := range langs {
		if lang != expected[i] {
			t.Errorf("SupportedLanguages()[%d] = %q, want %q", i, lang, expected[i])
		}
	}
}

func TestGetProjectConfigDefaults(t *testing.T) {
	cfg := NewConfig()
	cfg.Load()
	pc := cfg.GetProjectConfig()

	if pc.Language != "go" {
		t.Errorf("GetProjectConfig() Language = %q, want go", pc.Language)
	}
	if pc.Framework != "gin" {
		t.Errorf("GetProjectConfig() Framework = %q, want gin", pc.Framework)
	}
	if pc.Author != "Developer" {
		t.Errorf("GetProjectConfig() Author = %q, want Developer", pc.Author)
	}
	if pc.Version != "1.0.0" {
		t.Errorf("GetProjectConfig() Version = %q, want 1.0.0", pc.Version)
	}
}

func TestGetProjectConfig(t *testing.T) {
	if _, err := os.Stat("devhelper.yaml"); err == nil {
		t.Log("devhelper.yaml exists locally, using values from file")
	}

	cfg := NewConfig()
	cfg.SetWorkingDir("")

	if cfg.viper.ConfigFileUsed() == "" {
		t.Logf("viper.ConfigFileUsed() = \"\" (no local config file found)")
	}

	pc := cfg.GetProjectConfig()
	if pc.Language != "go" {
		t.Errorf("GetProjectConfig() Language = %q, want go", pc.Language)
	}
}
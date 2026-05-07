package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// ProjectConfig describes a single generated project.
type ProjectConfig struct {
	Name        string `mapstructure:"name"`
	Language    string `mapstructure:"language"`
	Framework   string `mapstructure:"framework"`
	OutputDir   string `mapstructure:"output_dir"`
	InitGit     bool   `mapstructure:"init_git"`
	GenerateDocker bool `mapstructure:"generate_docker"`
	Author      string `mapstructure:"author"`
	Version     string `mapstructure:"version"`
}

// Config manages Viper-backed configuration.
type Config struct {
	viper      *viper.Viper
	workingDir string
}

var supportedLangs = []string{"go", "node", "python", "java"}

var frameworkMap = map[string][]string{
	"go":     {"gin", "fiber"},
	"node":   {"express"},
	"python": {"fastapi"},
	"java":   {"springboot"},
}

// NewConfig creates a Config with a fresh Viper instance.
func NewConfig() *Config {
	v := viper.New()
	v.SetConfigName("devhelper")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	home, err := os.UserHomeDir()
	if err == nil {
		v.AddConfigPath(filepath.Join(home, ".config", "dev-helper"))
		v.AddConfigPath(home)
	}

	v.SetDefault("name", "")
	v.SetDefault("language", "go")
	v.SetDefault("framework", "gin")
	v.SetDefault("output_dir", ".")
	v.SetDefault("init_git", true)
	v.SetDefault("generate_docker", false)
	v.SetDefault("author", "Developer")
	v.SetDefault("version", "1.0.0")

	return &Config{viper: v}
}

// Load reads the configuration file. Returns nil even if no file is found
// (defaults will be used).
func (c *Config) Load() error {
	err := c.viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("[config] no config file found, using defaults")
			return nil
		}
		return fmt.Errorf("read config: %w", err)
	}
	log.Printf("[config] loaded config from %s", c.viper.ConfigFileUsed())
	return nil
}

// Save writes current defaults to path.
func (c *Config) Save(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}
	return c.viper.SafeWriteConfigAs(path)
}

// GetProjectConfig unmarshals stored + default values into ProjectConfig.
func (c *Config) GetProjectConfig() ProjectConfig {
	var pc ProjectConfig
	if err := c.viper.Unmarshal(&pc); err != nil {
		// fallback to defaults
		pc = ProjectConfig{
			Language:    "go",
			Framework:   "gin",
			OutputDir:   ".",
			InitGit:     true,
			Author:      "Developer",
			Version:     "1.0.0",
		}
	} else {
		// apply defaults for empty fields
		if pc.Language == "" {
			pc.Language = "go"
		}
		if pc.Framework == "" {
			pc.Framework = "gin"
		}
	}
	return pc
}

// SetWorkingDir sets the directory config will be searched in.
func (c *Config) SetWorkingDir(dir string) {
	c.workingDir = dir
	c.viper.AddConfigPath(dir)
}

// ValidateLanguage returns true if lang is supported.
func (c *Config) ValidateLanguage(lang string) bool {
	lang = strings.ToLower(lang)
	for _, l := range supportedLangs {
		if l == lang {
			return true
		}
	}
	return false
}

// ValidateFramework returns true if framework is valid for the given language.
func (c *Config) ValidateFramework(lang, fw string) bool {
	langLower := strings.ToLower(lang)
	supported, ok := frameworkMap[langLower]
	if !ok {
		return false
	}
	for _, f := range supported {
		if strings.EqualFold(f, fw) {
			return true
		}
	}
	return false
}

// SupportedLanguages returns the list of supported languages.
func (c *Config) SupportedLanguages() []string {
	return supportedLangs
}

// SupportedFrameworks returns the available frameworks for a language.
func (c *Config) SupportedFrameworks(lang string) []string {
	return frameworkMap[strings.ToLower(lang)]
}
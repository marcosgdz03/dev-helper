package plugins

import (
	"fmt"
	"sync"
)

// Language constants
const (
	LangGo     = "go"
	LangNode   = "node"
	LangPython = "python"
	LangJava   = "java"
)

// Framework constants
const (
	FrameworkGin        = "gin"
	FrameworkFiber      = "fiber"
	FrameworkExpress    = "express"
	FrameworkFastapi    = "fastapi"
	FrameworkSpringboot = "springboot"
)

// PluginInfo holds metadata about a plugin
type PluginInfo struct {
	Name        string
	Version     string
	Description string
	Author      string
	Language    string
	Framework   string
	Priority    int
}

// Plugin defines the interface all plugins must implement
type Plugin interface {
	Info() PluginInfo
	Generate(config PluginConfig) ([]string, error)
	Validate() error
	Deps() []string
}

// PluginConfig is passed to Generate()
type PluginConfig struct {
	ProjectName string
	OutputDir   string
	Author      string
	Version     string
	Extra       map[string]interface{}
}

// Factory function type for plugin registration
type Factory func() Plugin

// Registry stores and manages plugins
type Registry struct {
	mu        sync.RWMutex
	plugins   []Plugin
	index     map[string]Plugin
	langIndex map[string][]Plugin
}

// NewRegistry creates an empty plugin registry
func NewRegistry() *Registry {
	return &Registry{
		index:     make(map[string]Plugin),
		langIndex: make(map[string][]Plugin),
	}
}

// Register adds a plugin to the registry
func (r *Registry) Register(plugin Plugin) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.plugins = append(r.plugins, plugin)

	info := plugin.Info()
	key := fmt.Sprintf("%s/%s", info.Language, info.Framework)
	r.index[key] = plugin

	r.langIndex[info.Language] = append(r.langIndex[info.Language], plugin)
}

// Get retrieves a plugin by key "language/framework"
func (r *Registry) Get(name string) (Plugin, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.index[name]
	if !ok {
		return nil, fmt.Errorf("plugin not found: %s", name)
	}
	return p, nil
}

// GetByLanguage returns all plugins for a given language
func (r *Registry) GetByLanguage(lang string) []Plugin {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.langIndex[lang]
	if !ok {
		return nil
	}
	result := make([]Plugin, len(p))
	copy(result, p)
	return result
}

// GetByFramework returns the plugin matching language/framework
func (r *Registry) GetByFramework(lang, framework string) (Plugin, error) {
	return r.Get(fmt.Sprintf("%s/%s", lang, framework))
}

// All returns all registered plugins
func (r *Registry) All() []Plugin {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]Plugin, len(r.plugins))
	copy(result, r.plugins)
	return result
}

// Count returns the number of registered plugins
func (r *Registry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.plugins)
}
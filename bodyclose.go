// Package bodyclose provides a linter that checks whether HTTP response bodies are properly closed.
// This linter can be used as a plugin with golangci-lint v2+ or as a standalone tool.
//
// When used with golangci-lint v2+, configure your .golangci.yml like this:
//
//	linters-settings:
//	  plugin:
//	    enable: true
//	    path: github.com/bfirestone/bodyclose
//
//	linters:
//	  enable:
//	    - bodyclose
package bodyclose

import (
	"github.com/bfirestone/bodyclose/passes/bodyclose"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("bodyclose", New)
}

type BodyClosePlugin struct{}

// New creates a new instance of the bodyclose linter plugin
func New(settings any) (register.LinterPlugin, error) {
	return &BodyClosePlugin{}, nil
}

// BuildAnalyzers returns the bodyclose analyzer
func (p *BodyClosePlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		bodyclose.Analyzer,
	}, nil
}

// GetLoadMode specifies what kind of AST loading is needed
func (p *BodyClosePlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo // We need types information for the analyzer
}

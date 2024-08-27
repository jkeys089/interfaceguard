package main

import (
	"github.com/jkeys089/interfaceguard/plugin/golangci/module"
	"golang.org/x/tools/go/analysis"
)

func New(settings any) ([]*analysis.Analyzer, error) {
	plugin, err := module.New(settings)
	if err != nil {
		return nil, err
	}

	return plugin.BuildAnalyzers()
}

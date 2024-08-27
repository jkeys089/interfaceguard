package module

import (
	"fmt"
	"strconv"

	"github.com/golangci/plugin-module-register/register"
	"github.com/jkeys089/interfaceguard"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("interfaceguard", New)
}

type InterfaceguardSettings struct {
	DisableInterfaceComparison bool `json:"disable-interface-to-interface"`
	DisableNilComparison       bool `json:"disable-interface-to-nil"`
}

type InterfaceguardPlugin struct {
	*analysis.Analyzer
}

func New(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[InterfaceguardSettings](settings)
	if err != nil {
		return nil, fmt.Errorf("error decoding interfaceguard settings: %w", err)
	}

	a := interfaceguard.NewAnalyzer(false, false)

	if err = a.Flags.Set("i", strconv.FormatBool(s.DisableInterfaceComparison)); err != nil {
		return nil, fmt.Errorf("error setting interfaceguard 'i' flag: %w", err)
	}

	if err = a.Flags.Set("n", strconv.FormatBool(s.DisableNilComparison)); err != nil {
		return nil, fmt.Errorf("error setting interfaceguard 'n' flag: %w", err)
	}

	return &InterfaceguardPlugin{Analyzer: a}, nil
}

func (p *InterfaceguardPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{p.Analyzer}, nil
}

func (p *InterfaceguardPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

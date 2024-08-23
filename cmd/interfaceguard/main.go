package main

import (
	"github.com/jkeys089/interfaceguard"
	"golang.org/x/tools/go/analysis/singlechecker"
)

const (
	defaultInterfaceCheckDisabled = false
	defaultNilCheckDisabled       = false
)

func main() {
	singlechecker.Main(interfaceguard.NewAnalyzer(defaultNilCheckDisabled, defaultInterfaceCheckDisabled))
}

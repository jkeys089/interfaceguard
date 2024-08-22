package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"interfaceguard"
)

const (
	defaultInterfaceCheckDisabled = false
	defaultNilCheckDisabled       = false
)

func main() {
	singlechecker.Main(interfaceguard.NewAnalyzer(defaultNilCheckDisabled, defaultInterfaceCheckDisabled))
}

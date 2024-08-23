package interfaceguard_test

import (
	"testing"

	"github.com/jkeys089/interfaceguard"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testCases := []struct {
		name                   string
		path                   string
		nilCheckDisabled       bool
		interfaceCheckDisabled bool
	}{
		{
			name: "all enabled",
			path: "a",
		},
		{
			name:                   "disable interface-to-interface comparison checks",
			path:                   "disableiface",
			interfaceCheckDisabled: true,
		},
		{
			name:             "disable interface nil comparison checks",
			path:             "disablenil",
			nilCheckDisabled: true,
		},
		{
			name:                   "all disabled",
			path:                   "disableall",
			interfaceCheckDisabled: true,
			nilCheckDisabled:       true,
		},
	}

	testData := analysistest.TestData()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			analyzer := interfaceguard.NewAnalyzer(tc.nilCheckDisabled, tc.interfaceCheckDisabled)
			analysistest.Run(t, testData, analyzer, tc.path)
		})
	}
}

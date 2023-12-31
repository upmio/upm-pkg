//go:build tools
// +build tools

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import (
	_ "k8s.io/code-generator"
	_ "k8s.io/kubernetes/cmd/preferredimports"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
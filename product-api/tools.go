// +build tools

// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package main

import (
	_ "golang.org/x/tools/cmd/goimports"
)

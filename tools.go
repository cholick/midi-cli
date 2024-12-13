//go:build tools

// Based on https://marcofranssen.nl/manage-go-tools-via-go-modules
//
// Golang in 1.24 is making the tools pattern official, so this is temporary.
// See https://github.com/golang/go/issues/48429

package main

import (
	// goimports for the cleaner import grouping (stdlib, newline, external)
	_ "golang.org/x/tools/cmd/goimports"

	_ "github.com/maxbrunsfeld/counterfeiter/v6"
)
